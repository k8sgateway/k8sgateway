package kubernetes

import (
	"context"
	"encoding/binary"
	"fmt"
	"hash/fnv"
	"os"
	"sort"
	"strings"
	"time"

	errors "github.com/rotisserie/eris"

	"github.com/solo-io/gloo/pkg/utils/settingsutil"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	kubeplugin "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/kubernetes"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	corecache "github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/cache"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	kubev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

func (p *plugin) WatchEndpoints(writeNamespace string, upstreamsToTrack v1.UpstreamList, opts clients.WatchOpts) (<-chan v1.EndpointList, <-chan error, error) {

	kubeFactory := func(namespaces []string) KubePluginSharedFactory {
		return getInformerFactory(opts.Ctx, p.kube, namespaces)
	}
	watcher, err := newEndpointWatcherForUpstreams(kubeFactory, p.kubeCoreCache, upstreamsToTrack, opts)
	if err != nil {
		return nil, nil, err
	}
	return watcher.watch(writeNamespace, opts)
}

func newEndpointWatcherForUpstreams(kubeFactoryFactory func(ns []string) KubePluginSharedFactory, kubeCoreCache corecache.KubeCoreCache, upstreamsToTrack v1.UpstreamList, opts clients.WatchOpts) (*edsWatcher, error) {
	var namespaces []string

	settings := settingsutil.FromContext(opts.Ctx)
	if settingsutil.IsAllNamespacesFromSettings(settings) {
		namespaces = []string{metav1.NamespaceAll}
	} else {
		nsSet := map[string]bool{}
		for _, upstream := range upstreamsToTrack {
			svcNs := upstream.GetKube().GetServiceNamespace()
			// only care about kube upstreams
			if svcNs == "" {
				continue
			}
			nsSet[svcNs] = true
		}
		for ns := range nsSet {
			namespaces = append(namespaces, ns)
		}
	}

	kubeFactory := kubeFactoryFactory(namespaces)
	// this can take a bit of time some make sure we are still in business
	if opts.Ctx.Err() != nil {
		return nil, opts.Ctx.Err()
	}
	opts = opts.WithDefaults()

	return newEndpointsWatcher(kubeCoreCache, namespaces, kubeFactory, upstreamsToTrack), nil
}

type edsWatcher struct {
	upstreams         map[*core.ResourceRef]*kubeplugin.UpstreamSpec
	kubeShareFactory  KubePluginSharedFactory
	kubeCoreCache     corecache.KubeCoreCache
	namespaces        []string
	lastEndpointsHash uint64
}

func newEndpointsWatcher(kubeCoreCache corecache.KubeCoreCache, namespaces []string, kubeShareFactory KubePluginSharedFactory, upstreams v1.UpstreamList) *edsWatcher {
	upstreamSpecs := make(map[*core.ResourceRef]*kubeplugin.UpstreamSpec)
	for _, us := range upstreams {
		kubeUpstream, ok := us.GetUpstreamType().(*v1.Upstream_Kube)
		// only care about kube upstreams
		if !ok {
			continue
		}
		upstreamSpecs[us.GetMetadata().Ref()] = kubeUpstream.Kube
	}
	return &edsWatcher{
		upstreams:        upstreamSpecs,
		kubeShareFactory: kubeShareFactory,
		kubeCoreCache:    kubeCoreCache,
		namespaces:       namespaces,
	}
}

func (c *edsWatcher) List(writeNamespace string, opts clients.ListOpts) (v1.EndpointList, error) {
	var endpointList []*kubev1.Endpoints
	var serviceList []*kubev1.Service
	var podList []*kubev1.Pod
	ctx := contextutils.WithLogger(opts.Ctx, "kubernetes_eds")
	var warnsToLog []string

	for _, ns := range c.namespaces {
		if c.kubeCoreCache.NamespacedServiceLister(ns) == nil {
			// this namespace is not watched, ignore it.
			warnsToLog = append(warnsToLog, fmt.Sprintf("namespace %v is not watched, and has upstreams pointing to it", ns))
			continue
		}
		services, err := c.kubeCoreCache.NamespacedServiceLister(ns).List(labels.SelectorFromSet(opts.Selector))
		if err != nil {
			return nil, err
		}
		serviceList = append(serviceList, services...)
		pods, err := c.kubeCoreCache.NamespacedPodLister(ns).List(labels.SelectorFromSet(opts.Selector))
		if err != nil {
			return nil, err
		}
		podList = append(podList, pods...)

		endpoints, err := c.kubeShareFactory.EndpointsLister(ns).List(labels.SelectorFromSet(opts.Selector))
		if err != nil {
			return nil, err
		}
		endpointList = append(endpointList, endpoints...)
	}

	// TODO: Switch to filterServiceEndpoints when Gloo configured to user Kube services as endpoints (i.e. for Istio compatibility)
	var eps v1.EndpointList
	var warns, errsToLog []string //, []string
	// TODO: I need to add the new Helm value and use it here
	istioIntegration := os.Getenv("GLOO_INTEGRATION_ISTIO") == "true"
	eps, warns, errsToLog = filterEndpoints(ctx, writeNamespace, endpointList, serviceList, podList, c.upstreams, istioIntegration)

	warnsToLog = append(warnsToLog, warns...)

	hasher := fnv.New64()
	for _, ep := range eps {
		result, err := ep.Hash(hasher)
		if err != nil {
			return nil, err
		}
		err = binary.Write(hasher, binary.LittleEndian, result)
		if err != nil {
			return nil, err
		}
	}
	hash := hasher.Sum64()

	if c.lastEndpointsHash == hash {
		return eps, nil
	}
	c.lastEndpointsHash = hash

	// We return and log any messages here later because endpoints are constantly being updated / triggering the
	// EDS loop by the endpoint shared informer. By ensuring the hash is different, we know the endpoints are
	// functionally different and logging any new information might mean different / new information.
	//
	// This change helps avoid disk filling up from endless logs and cluttering the logs with the same message.
	logger := contextutils.LoggerFrom(ctx)
	for _, warnToLog := range warnsToLog {
		logger.Warn(warnToLog)
	}
	for _, errToLog := range errsToLog {
		logger.Error(errToLog)
	}

	return eps, nil
}

func (c *edsWatcher) watch(writeNamespace string, opts clients.WatchOpts) (<-chan v1.EndpointList, <-chan error, error) {
	watch := c.kubeShareFactory.Subscribe()

	endpointsChan := make(chan v1.EndpointList)
	errs := make(chan error)

	updateResourceList := func() {
		list, err := c.List(writeNamespace, clients.ListOpts{
			Ctx:      opts.Ctx,
			Selector: opts.Selector,
		})
		if err != nil {
			errs <- err
			return
		}
		select {
		case <-opts.Ctx.Done():
			return
		case endpointsChan <- list:
		}
	}

	go func() {
		defer c.kubeShareFactory.Unsubscribe(watch)
		defer close(endpointsChan)
		defer close(errs)

		updateResourceList()
		for {
			select {
			case _, ok := <-watch:
				if !ok {
					return
				}
				updateResourceList()
			case <-opts.Ctx.Done():
				return
			}
		}

	}()
	return endpointsChan, errs, nil
}

type Epkey struct {
	Address     string
	Port        uint32
	Name        string
	Namespace   string
	UpstreamRef *core.ResourceRef
}

func findPortForService(services []*kubev1.Service, spec *kubeplugin.UpstreamSpec) (*kubev1.ServicePort, bool) {
	for _, svc := range services {
		if svc.Namespace != spec.GetServiceNamespace() || svc.Name != spec.GetServiceName() {
			continue
		}
		if len(svc.Spec.Ports) == 1 {
			if spec.GetServicePort() == uint32(svc.Spec.Ports[0].Port) {
				return &svc.Spec.Ports[0], true
			}
		}
		for _, port := range svc.Spec.Ports {
			if spec.GetServicePort() == uint32(port.Port) {
				return &port, false
			}
		}
	}
	return nil, false
}

// filterToIpAddresses - Set to true for Istio integration where endpoints must
// be defined by IP address rather than hostnames. For details, see:
// * https://github.com/murphye/demo-gloo-edge-istio
func filterEndpoints(
	_ context.Context, // do not use for logging! return logging messages as strings and log them after hashing (see https://github.com/solo-io/gloo/issues/3761)
	writeNamespace string,
	kubeEndpoints []*kubev1.Endpoints,
	services []*kubev1.Service,
	pods []*kubev1.Pod,
	upstreams map[*core.ResourceRef]*kubeplugin.UpstreamSpec,
	filterToIpAddresses bool,
) (v1.EndpointList, []string, []string) {
	var endpoints v1.EndpointList

	var warnsToLog, errorsToLog []string

	endpointsMap := make(map[Epkey][]*core.ResourceRef)

	// for each upstream
	for usRef, spec := range upstreams {
		kubeServicePort, singlePortService := findPortForService(services, spec)
		if kubeServicePort == nil {
			errorsToLog = append(errorsToLog, fmt.Sprintf("upstream %v: port %v not found for service %v", usRef.Key(), spec.GetServicePort(), spec.GetServiceName()))
			continue
		}
		// find each matching endpoint
		for _, eps := range kubeEndpoints {
			if eps.Namespace != spec.GetServiceNamespace() || eps.Name != spec.GetServiceName() {
				continue
			}
			for _, subset := range eps.Subsets {
				port := findPortForEndpointSubset(subset, singlePortService, kubeServicePort)
				if port == 0 {
					warnsToLog = append(warnsToLog, fmt.Sprintf("upstream %v: port %v not found for service %v in endpoint %v", usRef.Key(), spec.GetServicePort(), spec.GetServiceName(), subset))
					continue
				}

				if filterToIpAddresses {
					hostname := fmt.Sprintf("%v.%v", spec.GetServiceName(), spec.GetServiceNamespace())
					key := Epkey{hostname, port, spec.GetServiceName(), spec.GetServiceNamespace(), usRef}
					copyRef := *usRef
					endpointsMap[key] = append(endpointsMap[key], &copyRef)
				} else {
					warnings := processSubsetAddresses(subset, spec, pods, usRef, port, endpointsMap)
					warnsToLog = append(warnsToLog, warnings...)
				}
			}
		}
	}

	endpoints = generateFilteredEndpointList(endpointsMap, services, pods, writeNamespace, endpoints, filterToIpAddresses)

	return endpoints, warnsToLog, errorsToLog
}

func processSubsetAddresses(subset kubev1.EndpointSubset, spec *kubeplugin.UpstreamSpec, pods []*kubev1.Pod, usRef *core.ResourceRef, port uint32, endpointsMap map[Epkey][]*core.ResourceRef) []string {
	var warnings []string
	for _, addr := range subset.Addresses {
		var podName, podNamespace string
		targetRef := addr.TargetRef
		if targetRef != nil {
			if targetRef.Kind == "Pod" {
				podName = targetRef.Name
				podNamespace = targetRef.Namespace
			}
		}
		if len(spec.GetSelector()) != 0 {
			// determine whether labels for the owner of this ip (pod) matches the spec
			podLabels, err := getPodLabelsForIp(addr.IP, podName, podNamespace, pods)
			if err != nil {
				// pod not found for IP? what's that about?
				warnings = append(warnings, fmt.Sprintf("error for upstream %v service %v: %v", usRef.Key(), spec.GetServiceName(), err))
				continue
			}
			if !labels.SelectorFromSet(spec.GetSelector()).Matches(labels.Set(podLabels)) {
				continue
			}
			// pod hasn't been assigned address yet
			if addr.IP == "" {
				continue
			}
		}
		key := Epkey{addr.IP, port, podName, podNamespace, usRef}
		copyRef := *usRef
		endpointsMap[key] = append(endpointsMap[key], &copyRef)
	}
	return warnings
}

func findPortForEndpointSubset(subset kubev1.EndpointSubset, singlePortService bool, kubeServicePort *kubev1.ServicePort) uint32 {
	var port uint32
	for _, p := range subset.Ports {
		// if the endpoint port is not named, it implies that
		// the kube service only has a single unnamed port as well.
		switch {
		case singlePortService:
			port = uint32(p.Port)
		case p.Name == kubeServicePort.Name:
			port = uint32(p.Port)
			break
		}
	}
	return port
}

func generateFilteredEndpointList(
	endpointsMap map[Epkey][]*core.ResourceRef,
	services []*kubev1.Service,
	pods []*kubev1.Pod,
	writeNamespace string,
	endpoints v1.EndpointList,
	filterToIpAddresses bool) v1.EndpointList {
	for addr, refs := range endpointsMap {

		// sort refs for idempotency
		sort.Slice(refs, func(i, j int) bool { return refs[i].Key() < refs[j].Key() })
		hasher := fnv.New64()
		hasher.Write([]byte(fmt.Sprintf("%+v", addr)))
		dnsname := strings.Map(func(r rune) rune {
			if '0' <= r && r <= '9' {
				return r
			}
			if 'a' <= r && r <= 'z' {
				return r
			}
			return '-'
		}, addr.Address)
		endpointName := fmt.Sprintf("ep-%v-%v-%x", dnsname, addr.Port, hasher.Sum64())

		var ep *v1.Endpoint
		if filterToIpAddresses {
			service, _ := getServiceForHostname(addr.Address, addr.Name, addr.Namespace, services)
			ep = createEndpoint(writeNamespace, endpointName, refs, service.Spec.ClusterIP, addr.Port, service.GetObjectMeta().GetLabels()) // TODO: labels may be nil
		} else {
			if pod, _ := getPodForIp(addr.Address, addr.Name, addr.Namespace, pods); pod == nil {
				ep = createEndpoint(writeNamespace, endpointName, refs, addr.Address, addr.Port, nil)
			} else {
				ep = createEndpoint(writeNamespace, endpointName, refs, addr.Address, addr.Port, pod.GetObjectMeta().GetLabels())
			}
		}
		endpoints = append(endpoints, ep)
	}

	// sort refs for idempotency
	sort.Slice(endpoints, func(i, j int) bool {
		return endpoints[i].GetMetadata().GetName() < endpoints[j].GetMetadata().GetName()
	})
	return endpoints
}

func createEndpoint(namespace, name string, upstreams []*core.ResourceRef, address string, port uint32, labels map[string]string) *v1.Endpoint {
	ep := &v1.Endpoint{
		Metadata: &core.Metadata{
			Namespace: namespace,
			Name:      name,
			Labels:    labels,
		},
		Upstreams: upstreams,
		Address:   address,
		Port:      port,
		// TODO: add locality info

	}
	return ep
}

func getPodLabelsForIp(ip string, podName, podNamespace string, pods []*kubev1.Pod) (map[string]string, error) {
	pod, err := getPodForIp(ip, podName, podNamespace, pods)
	if err != nil {
		return nil, err
	}
	return pod.Labels, nil
}

func getPodForIp(ip string, podName, podNamespace string, pods []*kubev1.Pod) (*kubev1.Pod, error) {

	for _, pod := range pods {
		if podName != "" && podNamespace != "" {
			// no need for heuristics!
			if podName == pod.Name && podNamespace == pod.Namespace {
				return pod, nil
			}
		}

		if pod.Status.PodIP != ip {
			continue
		}
		if pod.Status.Phase != kubev1.PodRunning {
			continue
		}
		if pod.Spec.HostNetwork {
			// we can't tell pods apart if they are all on the host network.
			continue
		}
		return pod, nil
	}

	return nil, errors.Errorf("running pod not found with IP %v", ip)
}

func getServiceForHostname(hostname string, serviceName, serviceNamespace string, services []*kubev1.Service) (*kubev1.Service, error) {

	for _, service := range services {
		if serviceName != "" && serviceNamespace != "" {
			// no need for heuristics!
			if serviceName == service.Name && serviceNamespace == service.Namespace {
				return service, nil
			}
		}

		serviceHostname := fmt.Sprintf("%v.%v", service.Name, service.Namespace)

		if serviceHostname != hostname {
			continue
		}

		return service, nil
	}

	return nil, errors.Errorf("service not found with hostname %v", hostname)
}
