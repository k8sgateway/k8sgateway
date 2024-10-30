package proxy_syncer

import (
	"sort"
	"strings"

	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	envoy_config_endpoint_v3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/solo-io/gloo/projects/gateway2/krtcollections"
	"istio.io/api/networking/v1alpha3"
	"istio.io/istio/pkg/slices"
)

type LoadBalancingInfo struct {
	// pod info:

	// Augmented proxy pod labels
	PodLabels map[string]string
	// locality info for proxy pod
	PodLocality krtcollections.PodLocality

	// dest rule info:
	PriorityInfo *PriorityInfo
}

type PriorityInfo struct {

	// dest rule info:
	//	Failover []*LocalityLoadBalancerSetting_Failover
	FailoverPriority *Prioritizer
	Failover         []*v1alpha3.LocalityLoadBalancerSetting_Failover
}

type Prioritizer struct {
	proxyLabels            map[string]string
	priorityLabels         []string
	priorityLabelOverrides map[string]string
	lowestPriority         int
}

func NewPriorities(failoverPriority []string) *Prioritizer {
	if len(failoverPriority) == 0 {
		return nil
	}
	lowestPriority := len(failoverPriority)
	priorityLabels, priorityLabelOverrides := priorityLabelOverrides(failoverPriority)
	return &Prioritizer{
		priorityLabels:         priorityLabels,
		priorityLabelOverrides: priorityLabelOverrides,
		lowestPriority:         lowestPriority,
	}
}

func (p *Prioritizer) GetPriority(proxyLabels, upstreamEndpointLabels map[string]string) int {
	for j, label := range p.priorityLabels {
		valueForProxy, ok := p.priorityLabelOverrides[label]
		if !ok {
			valueForProxy = p.proxyLabels[label]
		}

		if valueForProxy != upstreamEndpointLabels[label] {
			return p.lowestPriority - j
		}
	}
	return 0
}

// Returning the label names in a separate array as the iteration of map is not ordered.
func priorityLabelOverrides(labels []string) ([]string, map[string]string) {
	const FailoverPriorityLabelDefaultSeparator = '='
	priorityLabels := make([]string, 0, len(labels))
	overriddenValueByLabel := make(map[string]string, len(labels))
	var tempStrings []string
	for _, labelWithValue := range labels {
		tempStrings = strings.Split(labelWithValue, string(FailoverPriorityLabelDefaultSeparator))
		priorityLabels = append(priorityLabels, tempStrings[0])
		if len(tempStrings) == 2 {
			overriddenValueByLabel[tempStrings[0]] = tempStrings[1]
			continue
		}
	}
	return priorityLabels, overriddenValueByLabel
}

func prioritize2(ep EndpointsForUpstream, lbInfo LoadBalancingInfo) *envoy_config_endpoint_v3.ClusterLoadAssignment {
	cla := &envoy_config_endpoint_v3.ClusterLoadAssignment{
		ClusterName: ep.clusterName,
	}
	for loc, eps := range ep.LbEps {
		var l *envoy_config_core_v3.Locality
		if loc != (krtcollections.PodLocality{}) {
			l = &envoy_config_core_v3.Locality{
				Region:  loc.Region,
				Zone:    loc.Zone,
				SubZone: loc.Subzone,
			}
		}

		endpoints := getEndpoints(eps, lbInfo)
		for _, ep := range endpoints {
			ep.Locality = l
		}

		cla.Endpoints = append(cla.Endpoints, endpoints...)
	}

	if lbInfo.PriorityInfo.FailoverPriority == nil {
		// if no priorities, fallback to failover
		if len(lbInfo.PriorityInfo.Failover) != 0 {
			proxyLocality := envoy_config_core_v3.Locality{
				Region:  lbInfo.PodLocality.Region,
				Zone:    lbInfo.PodLocality.Zone,
				SubZone: lbInfo.PodLocality.Subzone,
			}
			applyLocalityFailover(&proxyLocality, cla, lbInfo.PriorityInfo.Failover)
		}
	}

	// in theory we want to run endpoint plugins here.
	// we only have one endpoint plugin, and it's not clear if it is in use. so
	// consider deprecating the functionality. it's not easy to do as with krt we no longer have gloo 'Endpoint' objects
	return cla
}

func getEndpoints(eps []EndpointWithMd, lbinfo LoadBalancingInfo) []*envoy_config_endpoint_v3.LocalityLbEndpoints {
	if lbinfo.PriorityInfo.FailoverPriority == nil {
		return []*envoy_config_endpoint_v3.LocalityLbEndpoints{{
			LbEndpoints: slices.Map(eps, func(e EndpointWithMd) *envoy_config_endpoint_v3.LbEndpoint { return e.LbEndpoint }),
		}}
	}
	return applyFailoverPriorityPerLocality(eps, lbinfo)
}

func applyFailoverPriorityPerLocality(
	eps []EndpointWithMd, lbinfo LoadBalancingInfo) []*envoy_config_endpoint_v3.LocalityLbEndpoints {
	// key is priority, value is the index of LocalityLbEndpoints.LbEndpoints
	priorityMap := map[int][]int{}
	for i, ep := range eps {
		priority := lbinfo.PriorityInfo.FailoverPriority.GetPriority(lbinfo.PodLabels, ep.EndpointMd.Labels)
		priorityMap[priority] = append(priorityMap[priority], i)
	}

	// sort all priorities in increasing order.
	priorities := []int{}
	for priority := range priorityMap {
		priorities = append(priorities, priority)
	}
	sort.Ints(priorities)

	out := make([]*envoy_config_endpoint_v3.LocalityLbEndpoints, len(priorityMap))
	for i, priority := range priorities {
		out[i].Priority = uint32(priority)
		var weight uint32
		for _, index := range priorityMap[priority] {
			out[i].LbEndpoints = append(out[i].LbEndpoints, eps[index].LbEndpoint)
			weight += eps[index].GetLoadBalancingWeight().GetValue()
		}
		// reset weight
		out[i].LoadBalancingWeight = &wrappers.UInt32Value{
			Value: weight,
		}
	}

	return out
}

// talk about settings doing an internal restart - we may not need it here with krt.
// and if we do, make sure that it works correctly with connected client set
// set locality loadbalancing priority - This is based on Region/Zone/SubZone matching.
func applyLocalityFailover(
	proxyLocality *envoy_config_core_v3.Locality,
	loadAssignment *envoy_config_endpoint_v3.ClusterLoadAssignment,
	failover []*v1alpha3.LocalityLoadBalancerSetting_Failover,
) {
	// key is priority, value is the index of the LocalityLbEndpoints in ClusterLoadAssignment
	priorityMap := map[int][]int{}

	// 1. calculate the LocalityLbEndpoints.Priority compared with proxy locality
	for i, localityEndpoint := range loadAssignment.Endpoints {
		// if region/zone/subZone all match, the priority is 0.
		// if region/zone match, the priority is 1.
		// if region matches, the priority is 2.
		// if locality not match, the priority is 3.
		priority := LbPriority(proxyLocality, localityEndpoint.Locality)
		// region not match, apply failover settings when specified
		// update localityLbEndpoints' priority to 4 if failover not match
		if priority == 3 {
			for _, failoverSetting := range failover {
				if failoverSetting.From == proxyLocality.Region {
					if localityEndpoint.Locality == nil || localityEndpoint.Locality.Region != failoverSetting.To {
						priority = 4
					}
					break
				}
			}
		}
		// priority is calculated using the already assigned priority using failoverPriority.
		// Since there are at most 5 priorities can be assigned using locality failover(0-4),
		// we multiply the priority by 5 for maintaining the priorities already assigned.
		// Afterwards the final priorities can be calculted from 0 (highest) to N (lowest) without skipping.
		priorityInt := int(loadAssignment.Endpoints[i].Priority*5) + priority
		loadAssignment.Endpoints[i].Priority = uint32(priorityInt)
		priorityMap[priorityInt] = append(priorityMap[priorityInt], i)
	}

	// since Priorities should range from 0 (highest) to N (lowest) without skipping.
	// 2. adjust the priorities in order
	// 2.1 sort all priorities in increasing order.
	priorities := []int{}
	for priority := range priorityMap {
		priorities = append(priorities, priority)
	}
	sort.Ints(priorities)
	// 2.2 adjust LocalityLbEndpoints priority
	// if the index and value of priorities array is not equal.
	for i, priority := range priorities {
		if i != priority {
			// the LocalityLbEndpoints index in ClusterLoadAssignment.Endpoints
			for _, index := range priorityMap[priority] {
				loadAssignment.Endpoints[index].Priority = uint32(i)
			}
		}
	}
}
func LbPriority(proxyLocality, endpointsLocality *envoy_config_core_v3.Locality) int {
	if proxyLocality.GetRegion() == endpointsLocality.GetRegion() {
		if proxyLocality.GetZone() == endpointsLocality.GetZone() {
			if proxyLocality.GetSubZone() == endpointsLocality.GetSubZone() {
				return 0
			}
			return 1
		}
		return 2
	}
	return 3
}
