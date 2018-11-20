package kubernetes

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/kubeutils"
	"github.com/solo-io/solo-projects/projects/gloo/pkg/api/v1"
	kubeplugin "github.com/solo-io/solo-projects/projects/gloo/pkg/api/v1/plugins/kubernetes"
	"github.com/solo-io/solo-projects/projects/gloo/pkg/discovery"
	kubev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
)

const (
	discoveryAnnotationKey  = "gloo.solo.io/discover"
	discoveryAnnotationTrue = "true"
)

func (p *plugin) DiscoverUpstreams(watchNamespaces []string, writeNamespace string, opts clients.WatchOpts, discOpts discovery.Opts) (chan v1.UpstreamList, chan error, error) {
	if p.kubeShareFactory == nil {
		p.kubeShareFactory = getInformerFactory(p.kube)
	}

	watch := p.kubeShareFactory.Subscribe()

	opts = opts.WithDefaults()
	upstreamsChan := make(chan v1.UpstreamList)
	errs := make(chan error)
	discoverUpstreams := func() {
		services, err := p.kubeShareFactory.ServicesLister().List(labels.SelectorFromSet(opts.Selector))
		if err != nil {
			errs <- err
			return
		}
		pods, err := p.kubeShareFactory.PodsLister().List(labels.SelectorFromSet(opts.Selector))
		if err != nil {
			errs <- err
			return
		}

		upstreamsChan <- convertServices(watchNamespaces, services, pods, discOpts, writeNamespace)
	}

	go func() {
		defer p.kubeShareFactory.Unsubscribe(watch)
		defer close(upstreamsChan)
		defer close(errs)
		// watch should open up with an initial read
		discoverUpstreams()
		for {
			select {
			case _, ok := <-watch:
				if !ok {
					return
				}
				discoverUpstreams()
			case <-opts.Ctx.Done():
				return
			}
		}
	}()
	return upstreamsChan, errs, nil
}

func convertServices(watchNamespaces []string, services []*kubev1.Service, pods []*kubev1.Pod, opts discovery.Opts, writeNamespace string) v1.UpstreamList {
	var upstreams v1.UpstreamList
	for _, svc := range services {
		if skip(svc, opts) {
			continue
		}

		if !containsString(svc.Namespace, watchNamespaces) {
			continue
		}

		upstreams = append(upstreams, upstreamsForService(svc, pods, writeNamespace)...)
	}
	return upstreams
}


func upstreamsForService(svc *kubev1.Service, pods []*kubev1.Pod, writeNamespace string) v1.UpstreamList {
	var upstreams v1.UpstreamList
	for _, port := range svc.Spec.Ports {
		var extendedLabelSets []map[string]string
		for _, pod := range pods {
			if pod.Namespace != svc.Namespace {
				continue
			}
			if !labels.AreLabelsInWhiteList(svc.Spec.Selector, pod.Labels) {
				continue
			}
			if reflect.DeepEqual(svc.Spec.Selector, pod.Labels) {
				continue
			}

			// create upstreams for the extra labels beyond the selector
			extendedLabels := make(map[string]string)
			for k, v := range pod.Labels {
				// special case we ignore
				// it is common and provides nothing useful for discovery
				if k == "pod-template-hash" {
					continue
				}
				extendedLabels[k] = v
			}
			if len(extendedLabels) > 0 {
				extendedLabelSets = append(extendedLabelSets, extendedLabels)
			}
		}
		if len(extendedLabelSets) > 0 {
			extendedLabelSets = uniqueLabelSets(extendedLabelSets)
			for _, extendedLabels := range extendedLabelSets {
				upstreams = append(upstreams, createUpstream(writeNamespace, svc, port, extendedLabels))
			}
		}
		upstreams = append(upstreams, createUpstream(writeNamespace, svc, port, svc.Spec.Selector))
	}
	return upstreams
}

func createUpstream(writeNamespace string, svc *kubev1.Service, port kubev1.ServicePort,
	labels map[string]string) *v1.Upstream {
	meta := svc.ObjectMeta
	coremeta := kubeutils.FromKubeMeta(meta)
	coremeta.ResourceVersion = ""
	extraLabels := make(map[string]string)
	// find extra keys not present in the service selector
	for k, v := range labels {
		if _, ok := svc.Spec.Selector[k]; ok {
			continue
		}
		extraLabels[k] = v
	}
	coremeta.Name = upstreamName(meta.Namespace, meta.Name, port.Port, extraLabels)
	coremeta.Namespace = writeNamespace
	servicePort := port.TargetPort.IntVal
	if servicePort == 0 {
		servicePort = port.Port
	}
	return &v1.Upstream{
		Metadata: coremeta,
		UpstreamSpec: &v1.UpstreamSpec{
			UpstreamType: &v1.UpstreamSpec_Kube{
				Kube: &kubeplugin.UpstreamSpec{
					ServiceName:      meta.Name,
					ServiceNamespace: meta.Namespace,
					ServicePort:      uint32(servicePort),
					Selector:         labels,
				},
			},
		},
		DiscoveryMetadata: &v1.DiscoveryMetadata{},
	}
}

func upstreamName(serviceNamespace, serviceName string, servicePort int32, extraLabels map[string]string) string {
	var labelsTag string
	if len(extraLabels) > 0 {
		_, values := keysAndValues(extraLabels)
		labelsTag = fmt.Sprintf("-%v", strings.Join(values, "-"))
	}
	name := fmt.Sprintf("%s-%s%s-%v", serviceNamespace, serviceName, labelsTag, servicePort)
	if len(name) > 63 {
		// todo: ilackarms: handle potential collisions
		name = name[:63]
	}
	return name
}

// TODO: move to a utils package
//guaranteed to be same length

func containsString(s string, slice []string) bool {
	for _, s2 := range slice {
		if s2 == s {
			return true
		}
	}
	return false
}

func mapLess(m1, m2 map[string]string) bool {
	if len(m1) != len(m2) {
		return len(m1) < len(m2)
	}
	var keys1, keys2 []string
	for k := range m1 {
		keys1 = append(keys1, k)
	}
	sort.SliceStable(keys1, func(i, j int) bool {
		return keys1[i] < keys1[j]
	})
	for k := range m2 {
		keys2 = append(keys2, k)
	}
	sort.SliceStable(keys2, func(i, j int) bool {
		return keys2[i] < keys2[j]
	})
	for i := range keys1 {
		if keys1[i] != keys2[i] {
			return keys1[i] < keys2[i]
		}
		if m1[keys1[i]] != m2[keys2[i]] {
			return m1[keys1[i]] < m2[keys2[i]]
		}
	}
	return false
}

func uniqueLabelSets(in []map[string]string) []map[string]string {
	var out []map[string]string
	for _, set := range in {
		var found bool
		for _, outSet := range out {
			if reflect.DeepEqual(set, outSet) {
				found = true
				break
			}
		}
		if found {
			continue
		}
		out = append(out, set)
	}
	sort.SliceStable(out, func(i, j int) bool {
		return mapLess(out[i], out[j])
	})
	return in
}

func keysAndValues(m map[string]string) ([]string, []string) {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var values [] string
	for _, k := range keys {
		values = append(values, m[k])
	}
	return keys, values
}

func skip(svc *kubev1.Service, opts discovery.Opts) bool {
	// ilackarms: allow user to override the skip with an annotation
	// force discovery for a service with no selector
	if svc.ObjectMeta.Annotations[discoveryAnnotationKey] == discoveryAnnotationTrue {
		return false
	}
	if len(svc.Spec.Selector) == 0 {
		return true
	}
	// note: ilackarms: IgnoredServices is not set anywhere
	for _, name := range opts.KubeOpts.IgnoredServices {
		if svc.Name == name {
			return true
		}
	}
	return false
}

func (p *plugin) UpdateUpstream(original, desired *v1.Upstream) (bool, error) {
	originalSpec, ok := original.UpstreamSpec.UpstreamType.(*v1.UpstreamSpec_Kube)
	if !ok {
		return false, errors.Errorf("internal error: expected *v1.UpstreamSpec_Kube, got %v", reflect.TypeOf(original.UpstreamSpec.UpstreamType).Name())
	}
	desiredSpec, ok := desired.UpstreamSpec.UpstreamType.(*v1.UpstreamSpec_Kube)
	if !ok {
		return false, errors.Errorf("internal error: expected *v1.UpstreamSpec_Kube, got %v", reflect.TypeOf(original.UpstreamSpec.UpstreamType).Name())
	}
	// copy service spec, we don't want to overwrite that
	desiredSpec.Kube.ServiceSpec = originalSpec.Kube.ServiceSpec
	// copy labels; user may have written them over. cannot be auto-discovered
	desiredSpec.Kube.Selector = originalSpec.Kube.Selector

	if originalSpec.Equal(desiredSpec) {
		return false, nil
	}

	return true, nil
}
