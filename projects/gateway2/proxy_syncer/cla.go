package proxy_syncer

import (
	"fmt"
	"hash/fnv"

	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	envoy_config_endpoint_v3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	"github.com/solo-io/gloo/projects/gateway2/krtcollections"
	envoycache "github.com/solo-io/solo-kit/pkg/api/v1/control-plane/cache"
	"github.com/solo-io/solo-kit/pkg/api/v1/control-plane/resource"
	"go.uber.org/zap"
	"istio.io/api/networking/v1alpha3"
	"istio.io/istio/pkg/kube/krt"
	"k8s.io/apimachinery/pkg/types"
)

type EndpointResources struct {
	Endpoints        envoycache.Resource
	EndpointsVersion uint64
	UpstreamRef      types.NamespacedName
}

func (c EndpointResources) ResourceName() string {
	return c.UpstreamRef.String()
}

func (c EndpointResources) Equals(in EndpointResources) bool {
	return c.UpstreamRef == in.UpstreamRef && c.EndpointsVersion == in.EndpointsVersion
}

// TODO: this is needed temporary while we don't have the per-upstream translation done.
// once the plugins are fixed to support it, we can have the proxy translation skip upstreams/endpoints and remove this collection
func newEnvoyEndpoints(glooEndpoints krt.Collection[EndpointsForUpstream]) krt.Collection[EndpointResources] {

	clas := krt.NewCollection(glooEndpoints, func(_ krt.HandlerContext, ep EndpointsForUpstream) *EndpointResources {
		return TransformEndpointToResources(ep)
	})
	return clas
}

func TransformEndpointToResources(ep EndpointsForUpstream) *EndpointResources {
	cla := prioritize(ep)
	return &EndpointResources{
		Endpoints:        resource.NewEnvoyResource(cla),
		EndpointsVersion: ep.lbEpsEqualityHash,
		UpstreamRef:      ep.UpstreamRef,
	}
}

func prioritize(ep EndpointsForUpstream) *envoy_config_endpoint_v3.ClusterLoadAssignment {
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

		lbeps := make([]*envoy_config_endpoint_v3.LbEndpoint, 0, len(eps))
		for _, ep := range eps {
			lbeps = append(lbeps, ep.LbEndpoint)
		}

		endpoint := &envoy_config_endpoint_v3.LocalityLbEndpoints{
			LbEndpoints: lbeps,
			Locality:    l,
		}

		cla.Endpoints = append(cla.GetEndpoints(), endpoint)
	}

	// In theory we want to run endpoint plugins here.
	// we only have one endpoint plugin - and it also does failover... so might be simpler to not support it in ggv2 and
	// deprecating the functionality. it's not easy to do as with krt we no longer have gloo 'Endpoint' objects
	return cla
}

type UccWithEndpoints struct {
	Client        krtcollections.UniqlyConnectedClient
	Endpoints     envoycache.Resource
	EndpointsHash uint64
	endpointsName string
}

func (c UccWithEndpoints) ResourceName() string {
	return fmt.Sprintf("%s/%s", c.Client.ResourceName(), c.endpointsName)
}

func (c UccWithEndpoints) Equals(in UccWithEndpoints) bool {
	return c.Client.Equals(in.Client) && c.EndpointsHash == in.EndpointsHash
}

type PerClientEnvoyEndpoints struct {
	endpoints krt.Collection[UccWithEndpoints]
	index     krt.Index[string, UccWithEndpoints]
}

func (ie *PerClientEnvoyEndpoints) FetchEndpointsForClient(kctx krt.HandlerContext, ucc krtcollections.UniqlyConnectedClient) []UccWithEndpoints {
	return krt.Fetch(kctx, ie.endpoints, krt.FilterIndex(ie.index, ucc.ResourceName()))
}

func NewPerClientEnvoyEndpoints(logger *zap.Logger, uccs krt.Collection[krtcollections.UniqlyConnectedClient],
	glooEndpoints krt.Collection[EndpointsForUpstream],
	destinationRulesIndex DestinationRuleIndex) PerClientEnvoyEndpoints {

	clas := krt.NewManyCollection(glooEndpoints, func(kctx krt.HandlerContext, ep EndpointsForUpstream) []UccWithEndpoints {
		uccs := krt.Fetch(kctx, uccs)
		uccWithEndpointsRet := make([]UccWithEndpoints, 0, len(uccs))
		for _, ucc := range uccs {
			destrule := destinationRulesIndex.FetchDestRulesFor(kctx, ucc.Namespace, ep.Hostname, ucc.Labels)
			uccWithEp := PrioritizeEndpoints(logger, destrule, ep, ucc)
			uccWithEndpointsRet = append(uccWithEndpointsRet, uccWithEp)
		}
		return uccWithEndpointsRet
	})
	idx := krt.NewIndex(clas, func(ucc UccWithEndpoints) []string {
		return []string{ucc.Client.ResourceName()}
	})

	return PerClientEnvoyEndpoints{
		endpoints: clas,
		index:     idx,
	}
}

func PrioritizeEndpoints(logger *zap.Logger, destrule *DestinationRuleWrapper, ep EndpointsForUpstream, ucc krtcollections.UniqlyConnectedClient) UccWithEndpoints {
	var additionalHash uint64
	var priorityInfo *PriorityInfo
	if destrule != nil {
		enabled := destrule.Spec.GetTrafficPolicy().GetLoadBalancer().GetLocalityLbSetting().GetEnabled()
		if enabled == nil || enabled.Value {
			localityLb := destrule.Spec.GetTrafficPolicy().GetLoadBalancer().GetLocalityLbSetting()
			for _, portlevel := range destrule.Spec.GetTrafficPolicy().GetPortLevelSettings() {
				if portlevel.GetPort().GetNumber() == ep.Port {
					localityLb = portlevel.GetLoadBalancer().GetLocalityLbSetting()
					break
				}
			}

			if localityLb != nil {
				priorityInfo = getPriorityInfoFromDestrule(localityLb)
				hasher := fnv.New64()
				hasher.Write([]byte(destrule.UID))
				hasher.Write([]byte(fmt.Sprintf("%v", destrule.Generation)))
				additionalHash = hasher.Sum64()
			}
		}
	}
	lbInfo := LoadBalancingInfo{
		PodLabels:    ucc.Labels,
		PodLocality:  ucc.Locality,
		PriorityInfo: priorityInfo,
	}

	cla := prioritizeWithLbInfo(logger, ep, lbInfo)
	return UccWithEndpoints{
		Client:        ucc,
		Endpoints:     resource.NewEnvoyResource(cla),
		EndpointsHash: ep.lbEpsEqualityHash ^ additionalHash,
		endpointsName: ep.ResourceName(),
	}
}

func getPriorityInfoFromDestrule(localityLb *v1alpha3.LocalityLoadBalancerSetting) *PriorityInfo {
	return &PriorityInfo{
		FailoverPriority: NewPriorities(localityLb.GetFailoverPriority()),
		Failover:         localityLb.GetFailover(),
	}
}
