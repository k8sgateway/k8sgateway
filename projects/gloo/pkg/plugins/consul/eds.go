package consul

import (
	"context"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/solo-io/go-utils/contextutils"

	"github.com/solo-io/gloo/pkg/utils"

	consulapi "github.com/hashicorp/consul/api"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/go-utils/errutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"golang.org/x/sync/errgroup"
)

const (
	EndpointTagMatchTrue  = "1"
	EndpointTagMatchFalse = "0"
	DataCenterLabelKey    = "data_center"
)

// Starts a watch on the Consul service metadata endpoint for all the services associated with the tracked upstreams.
// Whenever it detects an update to said services, it fetches the complete specs for the tracked services,
// converts them to endpoints, and sends the result on the returned channel.
func (p *plugin) WatchEndpoints(_ string, upstreamsToTrack v1.UpstreamList, opts clients.WatchOpts) (<-chan v1.EndpointList, <-chan error, error) {
	endpointsChan := make(chan v1.EndpointList)
	errChan := make(chan error)

	logger := contextutils.LoggerFrom(contextutils.WithLogger(opts.Ctx, "consul_eds"))

	// Filter out non-consul upstreams
	trackedServices := make(map[string][]*v1.Upstream)
	for _, us := range upstreamsToTrack {
		if consulUsSpec := us.UpstreamSpec.GetConsul(); consulUsSpec != nil {
			// We generate one upstream for every Consul service name, so this should never happen.
			trackedServices[consulUsSpec.ServiceName] = append(trackedServices[consulUsSpec.ServiceName], us)
		}
	}

	dataCenters, err := p.client.DataCenters()
	if err != nil {
		return nil, nil, err
	}

	serviceMetaChan, servicesWatchErrChan := p.client.WatchServices(opts.Ctx, dataCenters)

	var errAggregator sync.WaitGroup
	errAggregator.Add(1)
	go func() {
		defer errAggregator.Done()
		errutils.AggregateErrs(opts.Ctx, errChan, servicesWatchErrChan, "consul eds")
	}()

	go func() {

		// Create a new context for each loop, cancel it before each loop
		var cancel context.CancelFunc = func() {}
		// Use closure to allow cancel function to be updated as context changes
		defer func() { cancel() }()

		for {
			select {
			case serviceMeta, ok := <-serviceMetaChan:
				if !ok {
					return
				}

				// Cancel any running requests from previous iteration and set new context/cancel
				cancel()
				ctx, newCancel := context.WithCancel(opts.Ctx)
				cancel = newCancel

				specs := newSpecCollector()

				// Get complete service information for every dataCenter:service tuple in separate goroutines
				var eg errgroup.Group
				for _, service := range serviceMeta {
					for _, dataCenter := range service.DataCenters {

						// Copy iterator variables before passing them to goroutines!
						svc := service
						dcName := dataCenter

						// Get complete spec for each service in parallel
						eg.Go(func() error {
							queryOpts := &consulapi.QueryOptions{Datacenter: dcName, RequireConsistent: true}

							services, _, err := p.client.Service(svc.Name, "", queryOpts.WithContext(ctx))
							if err != nil {
								return err
							}

							specs.Append(services)

							return nil
						})
					}
				}

				// Wait for all requests to complete, an error to occur, or for the underlying context to be cancelled.
				//
				// Don't return if an error occurred. We still want to propagate the endpoints  for the requests that
				// succeeded. Any inconsistencies will be caught by the Gloo translator.
				if err := eg.Wait(); err != nil {
					select {
					case errChan <- err:
					default:
						logger.Errorf("write error channel is full! could not propagate err: %v", err)
					}
				}

				var endpoints v1.EndpointList
				for _, spec := range specs.Get() {
					if upstreams, ok := trackedServices[spec.ServiceName]; ok {
						endpoints = append(endpoints, createEndpoint(spec, upstreams))
					}
				}

				// Sort by name in ascending order for idempotency
				sort.SliceStable(endpoints, func(i, j int) bool {
					return endpoints[i].Metadata.Name < endpoints[j].Metadata.Name
				})

				endpointsChan <- endpoints

			case <-opts.Ctx.Done():
				close(endpointsChan)

				// Wait for error aggregation routing to complete to avoid writing to closed errChan
				errAggregator.Wait()
				close(errChan)
				return
			}
		}
	}()

	return endpointsChan, errChan, nil
}

func createEndpoint(service *consulapi.CatalogService, upstreams []*v1.Upstream) *v1.Endpoint {
	ep := &v1.Endpoint{
		Metadata: core.Metadata{
			Namespace:       "", // no namespace
			Name:            buildEndpointName(service),
			Labels:          buildLabels(service, upstreams),
			ResourceVersion: strconv.FormatUint(service.ModifyIndex, 10),
		},
		Upstreams: toResourceRefs(upstreams),
		Address:   service.Address,
		Port:      uint32(service.ServicePort),
	}
	return ep
}

func buildEndpointName(service *consulapi.CatalogService) string {
	parts := []string{service.ServiceName}
	if service.ServiceID != "" {
		parts = append(parts, service.ServiceID)
	}
	return strings.Join(parts, "_")
}

// The labels will be used by to match the endpoint to the subsets of the cluster represented by the upstream.
func buildLabels(service *consulapi.CatalogService, upstreams []*v1.Upstream) map[string]string {
	svcTags := make(map[string]bool)
	for _, tag := range service.ServiceTags {
		svcTags[tag] = true
	}

	// The ServiceTags on the Consul Upstream(s) represent all tags for Consul services with the given ServiceName across
	// data centers. We create an endpoint label for each of these tags, where the label key is the name of the tag and
	// the label value is "1" if the current service contains the same tag, else "0".
	labels := make(map[string]string)
	for _, usTag := range getUniqueUpstreamTags(upstreams) {
		if _, ok := svcTags[usTag]; ok {
			labels[usTag] = EndpointTagMatchTrue
		} else {
			labels[usTag] = EndpointTagMatchFalse
		}
	}

	// Create a label to associate the endpoint with a data center
	labels[DataCenterLabelKey] = service.Datacenter

	return labels
}

func toResourceRefs(upstreams []*v1.Upstream) (out []*core.ResourceRef) {
	for _, us := range upstreams {
		out = append(out, utils.ResourceRefPtr(us.Metadata.Ref()))
	}
	return
}

func getUniqueUpstreamTags(upstreams []*v1.Upstream) (tags []string) {
	tagMap := make(map[string]bool)
	for _, us := range upstreams {
		for _, tag := range us.UpstreamSpec.GetConsul().ServiceTags {
			tagMap[tag] = true
		}
	}
	for tag := range tagMap {
		tags = append(tags, tag)
	}
	return
}

func newSpecCollector() specCollector {
	return &collector{}
}

type specCollector interface {
	Append([]*consulapi.CatalogService)
	Get() []*consulapi.CatalogService
}

type collector struct {
	mutex sync.Mutex
	specs []*consulapi.CatalogService
}

func (c *collector) Append(specs []*consulapi.CatalogService) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.specs = append(c.specs, specs...)
}

func (c *collector) Get() []*consulapi.CatalogService {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.specs
}
