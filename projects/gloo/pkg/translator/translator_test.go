package translator_test

import (
	"context"

	envoyapi "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/solo-io/gloo/projects/gloo/pkg/translator"

	envoycluster "github.com/envoyproxy/go-control-plane/envoy/api/v2/cluster"
	envoyhttp "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/http_connection_manager/v2"
	types "github.com/gogo/protobuf/types"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	v1plugins "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins"
	v1kubernetes "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/kubernetes"
	v1static "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/static"
	"github.com/solo-io/gloo/projects/gloo/pkg/bootstrap"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins/registry"
	"github.com/solo-io/gloo/projects/gloo/pkg/xds"
	envoycache "github.com/solo-io/solo-kit/pkg/api/v1/control-plane/cache"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

var _ = Describe("Translator", func() {
	var (
		settings   *v1.Settings
		translator Translator
		upstream   *v1.Upstream
		proxy      *v1.Proxy
		params     plugins.Params
		matcher    *v1.Matcher
		routes     []*v1.Route

		snapshot            envoycache.Snapshot
		cluster             *envoyapi.Cluster
		listener            *envoyapi.Listener
		hcm_cfg             *envoyhttp.HttpConnectionManager
		route_configuration *envoyapi.RouteConfiguration
	)

	BeforeEach(func() {
		cluster = nil
		settings = &v1.Settings{}
		opts := bootstrap.Opts{
			Settings: settings,
		}
		tplugins := registry.Plugins(opts)
		translator = NewTranslator(tplugins, settings)

		upname := core.Metadata{
			Name:      "test",
			Namespace: "gloo-system",
		}
		upstream = &v1.Upstream{
			Metadata: upname,
			UpstreamSpec: &v1.UpstreamSpec{
				UpstreamType: &v1.UpstreamSpec_Static{
					Static: &v1static.UpstreamSpec{
						Hosts: []*v1static.Host{
							{
								Addr: "Test",
								Port: 124,
							},
						},
					},
				},
			},
		}

		params = plugins.Params{
			Ctx: context.Background(),
			Snapshot: &v1.ApiSnapshot{
				Upstreams: v1.UpstreamsByNamespace{
					"gloo-system": v1.UpstreamList{
						upstream,
					},
				},
			},
		}
		matcher = &v1.Matcher{
			PathSpecifier: &v1.Matcher_Prefix{
				Prefix: "/",
			},
		}
		routes = []*v1.Route{{
			Matcher: matcher,
			Action: &v1.Route_RouteAction{
				RouteAction: &v1.RouteAction{
					Destination: &v1.RouteAction_Single{
						Single: &v1.Destination{
							Upstream: upname.Ref(),
						},
					},
				},
			},
		}}
	})
	JustBeforeEach(func() {
		proxy = &v1.Proxy{
			Metadata: core.Metadata{
				Name:      "test",
				Namespace: "gloo-system",
			},
			Listeners: []*v1.Listener{{
				Name:        "listener",
				BindAddress: "127.0.0.1",
				BindPort:    80,
				ListenerType: &v1.Listener_HttpListener{
					HttpListener: &v1.HttpListener{
						VirtualHosts: []*v1.VirtualHost{{
							Name:    "virt1",
							Domains: []string{"*"},
							Routes:  routes,
						}},
					},
				},
			}},
		}
	})
	translate := func() {

		snap, errs, err := translator.Translate(params, proxy)
		Expect(err).NotTo(HaveOccurred())
		Expect(errs.Validate()).NotTo(HaveOccurred())
		Expect(snap).NotTo(BeNil())

		clusters := snap.GetResources(xds.ClusterType)
		clusterResource := clusters.Items[UpstreamToClusterName(upstream.Metadata.Ref())]
		cluster = clusterResource.ResourceProto().(*envoyapi.Cluster)
		Expect(cluster).NotTo(BeNil())

		listeners := snap.GetResources(xds.ListenerType)

		listenerResource := listeners.Items["listener"]
		listener = listenerResource.ResourceProto().(*envoyapi.Listener)
		Expect(listener).NotTo(BeNil())

		hcm_filter := listener.FilterChains[0].Filters[0]
		hcm_cfg = &envoyhttp.HttpConnectionManager{}
		err = ParseConfig(&hcm_filter, hcm_cfg)
		Expect(err).NotTo(HaveOccurred())

		routes := snap.GetResources(xds.RouteType)

		Expect(routes.Items).To(HaveKey("listener-routes"))
		routeResource := routes.Items["listener-routes"]
		route_configuration = routeResource.ResourceProto().(*envoyapi.RouteConfiguration)
		Expect(route_configuration).NotTo(BeNil())

		snapshot = snap

	}

	Context("route header match", func() {
		It("should translate header matcher with no value to a PresentMatch", func() {

			matcher.Headers = []*v1.HeaderMatcher{
				{
					Name: "test",
				},
			}
			translate()
			headermatch := route_configuration.VirtualHosts[0].Routes[0].Match.Headers[0]
			Expect(headermatch.Name).To(Equal("test"))
			presentmatch := headermatch.GetPresentMatch()
			Expect(presentmatch).To(BeTrue())
		})

		It("should translate header matcher with value to exact match", func() {

			matcher.Headers = []*v1.HeaderMatcher{
				{
					Name:  "test",
					Value: "testvalue",
				},
			}
			translate()

			headermatch := route_configuration.VirtualHosts[0].Routes[0].Match.Headers[0]
			Expect(headermatch.Name).To(Equal("test"))
			exactmatch := headermatch.GetExactMatch()
			Expect(exactmatch).To(Equal("testvalue"))
		})

		It("should translate header matcher with regex becomes regex match", func() {

			matcher.Headers = []*v1.HeaderMatcher{
				{
					Name:  "test",
					Value: "testvalue",
					Regex: true,
				},
			}
			translate()

			headermatch := route_configuration.VirtualHosts[0].Routes[0].Match.Headers[0]
			Expect(headermatch.Name).To(Equal("test"))
			regex := headermatch.GetRegexMatch()
			Expect(regex).To(Equal("testvalue"))
		})

	})

	Context("circuit breakers", func() {

		It("should NOT translate circuit breakers on upstream", func() {
			translate()
			Expect(cluster.CircuitBreakers).To(BeNil())
		})

		It("should translate circuit breakers on upstream", func() {

			upstream.UpstreamSpec.CircuitBreakers = &v1.CircuitBreakerConfig{
				MaxConnections:     &types.UInt32Value{Value: 1},
				MaxPendingRequests: &types.UInt32Value{Value: 2},
				MaxRequests:        &types.UInt32Value{Value: 3},
				MaxRetries:         &types.UInt32Value{Value: 4},
			}

			expectedCircuitBreakers := &envoycluster.CircuitBreakers{
				Thresholds: []*envoycluster.CircuitBreakers_Thresholds{
					{
						MaxConnections:     &types.UInt32Value{Value: 1},
						MaxPendingRequests: &types.UInt32Value{Value: 2},
						MaxRequests:        &types.UInt32Value{Value: 3},
						MaxRetries:         &types.UInt32Value{Value: 4},
					},
				},
			}
			translate()

			Expect(cluster.CircuitBreakers).To(BeEquivalentTo(expectedCircuitBreakers))
		})

		It("should translate circuit breakers on settings", func() {

			settings.CircuitBreakers = &v1.CircuitBreakerConfig{
				MaxConnections:     &types.UInt32Value{Value: 1},
				MaxPendingRequests: &types.UInt32Value{Value: 2},
				MaxRequests:        &types.UInt32Value{Value: 3},
				MaxRetries:         &types.UInt32Value{Value: 4},
			}

			expectedCircuitBreakers := &envoycluster.CircuitBreakers{
				Thresholds: []*envoycluster.CircuitBreakers_Thresholds{
					{
						MaxConnections:     &types.UInt32Value{Value: 1},
						MaxPendingRequests: &types.UInt32Value{Value: 2},
						MaxRequests:        &types.UInt32Value{Value: 3},
						MaxRetries:         &types.UInt32Value{Value: 4},
					},
				},
			}
			translate()

			Expect(cluster.CircuitBreakers).To(BeEquivalentTo(expectedCircuitBreakers))
		})

		It("should override circuit breakers on upstream", func() {

			settings.CircuitBreakers = &v1.CircuitBreakerConfig{
				MaxConnections:     &types.UInt32Value{Value: 11},
				MaxPendingRequests: &types.UInt32Value{Value: 12},
				MaxRequests:        &types.UInt32Value{Value: 13},
				MaxRetries:         &types.UInt32Value{Value: 14},
			}

			upstream.UpstreamSpec.CircuitBreakers = &v1.CircuitBreakerConfig{
				MaxConnections:     &types.UInt32Value{Value: 1},
				MaxPendingRequests: &types.UInt32Value{Value: 2},
				MaxRequests:        &types.UInt32Value{Value: 3},
				MaxRetries:         &types.UInt32Value{Value: 4},
			}

			expectedCircuitBreakers := &envoycluster.CircuitBreakers{
				Thresholds: []*envoycluster.CircuitBreakers_Thresholds{
					{
						MaxConnections:     &types.UInt32Value{Value: 1},
						MaxPendingRequests: &types.UInt32Value{Value: 2},
						MaxRequests:        &types.UInt32Value{Value: 3},
						MaxRetries:         &types.UInt32Value{Value: 4},
					},
				},
			}
			translate()

			Expect(cluster.CircuitBreakers).To(BeEquivalentTo(expectedCircuitBreakers))
		})

	})

	Context("when handling upstream groups", func() {

		var (
			upstream2     *v1.Upstream
			upstreamGroup *v1.UpstreamGroup
		)

		BeforeEach(func() {
			upstream2 = &v1.Upstream{
				Metadata: core.Metadata{
					Name:      "test2",
					Namespace: "gloo-system",
				},
				UpstreamSpec: &v1.UpstreamSpec{
					UpstreamType: &v1.UpstreamSpec_Static{
						Static: &v1static.UpstreamSpec{
							Hosts: []*v1static.Host{
								{
									Addr: "Test2",
									Port: 124,
								},
							},
						},
					},
				},
			}
			upstreamGroup = &v1.UpstreamGroup{
				Metadata: core.Metadata{
					Name:      "test",
					Namespace: "gloo-system",
				},
				Destinations: []*v1.WeightedDestination{
					{
						Weight: 1,
						Destination: &v1.Destination{
							Upstream: upstream.Metadata.Ref(),
						},
					},
					{
						Weight: 1,
						Destination: &v1.Destination{
							Upstream: upstream2.Metadata.Ref(),
						},
					},
				},
			}
			params.Snapshot.Upstreams["gloo-system"] = append(params.Snapshot.Upstreams["gloo-system"], upstream2)
			params.Snapshot.Upstreamgroups = v1.UpstreamgroupsByNamespace{
				"gloo-system": v1.UpstreamGroupList{
					upstreamGroup,
				},
			}
			ref := upstreamGroup.Metadata.Ref()
			routes = []*v1.Route{{
				Matcher: matcher,
				Action: &v1.Route_RouteAction{
					RouteAction: &v1.RouteAction{
						Destination: &v1.RouteAction_UpstreamGroup{
							UpstreamGroup: &ref,
						},
					},
				},
			}}
		})

		It("should translate upstream groups", func() {
			translate()

			route := route_configuration.VirtualHosts[0].Routes[0].GetRoute()
			Expect(route).ToNot(BeNil())
			clusters := route.GetWeightedClusters()
			Expect(clusters).ToNot(BeNil())
			Expect(clusters.TotalWeight.Value).To(BeEquivalentTo(2))
			Expect(clusters.Clusters).To(HaveLen(2))
			Expect(clusters.Clusters[0].Name).To(Equal(UpstreamToClusterName(upstream.Metadata.Ref())))
			Expect(clusters.Clusters[1].Name).To(Equal(UpstreamToClusterName(upstream2.Metadata.Ref())))
		})

		It("should error on invalid ref in upstream groups", func() {
			upstreamGroup.Destinations[0].Destination.Upstream.Name = "notexist"

			_, errs, err := translator.Translate(params, proxy)
			Expect(err).NotTo(HaveOccurred())
			err = errs.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("destination # 1: upstream not found: list did not find upstream gloo-system.notexist"))
		})

	})

	Context("when handling subsets", func() {
		var (
			cla_configuration *envoyapi.ClusterLoadAssignment
		)
		BeforeEach(func() {
			cla_configuration = nil

			upstream.UpstreamSpec.UpstreamType = &v1.UpstreamSpec_Kube{
				Kube: &v1kubernetes.UpstreamSpec{
					SubsetSpec: &v1plugins.SubsetSpec{
						Selectors: []*v1plugins.Selector{{
							Keys: []string{
								"testkey",
							},
						}},
					},
				},
			}
			ref := upstream.Metadata.Ref()
			params.Snapshot.Endpoints = v1.EndpointsByNamespace{
				"gloo-system": v1.EndpointList{
					{
						Metadata: core.Metadata{
							Name:      "test",
							Namespace: "gloo-system",
							Labels:    map[string]string{"testkey": "testvalue"},
						},
						Upstreams: []*core.ResourceRef{
							&ref,
						},
						Address: "1.2.3.4",
						Port:    1234,
					},
				},
			}

			routes = []*v1.Route{{
				Matcher: matcher,
				Action: &v1.Route_RouteAction{
					RouteAction: &v1.RouteAction{
						Destination: &v1.RouteAction_Single{
							Single: &v1.Destination{
								Upstream: upstream.Metadata.Ref(),
								Subset: &v1.Subset{
									Values: map[string]string{
										"testkey": "testvalue",
									},
								},
							},
						},
					},
				},
			}}

		})

		translateWithEndpoints := func() {
			translate()

			endpoints := snapshot.GetResources(xds.EndpointType)

			clusterName := UpstreamToClusterName(upstream.Metadata.Ref())
			Expect(endpoints.Items).To(HaveKey(clusterName))
			endpointsResource := endpoints.Items[clusterName]
			cla_configuration = endpointsResource.ResourceProto().(*envoyapi.ClusterLoadAssignment)
			Expect(cla_configuration).NotTo(BeNil())
			Expect(cla_configuration.ClusterName).To(Equal(clusterName))
			Expect(cla_configuration.Endpoints).To(HaveLen(1))
			Expect(cla_configuration.Endpoints[0].LbEndpoints).To(HaveLen(len(params.Snapshot.Endpoints["gloo-system"])))
		}

		Context("when happy path", func() {

			It("should transfer labels to envoy", func() {
				translateWithEndpoints()

				endpointMeta := cla_configuration.Endpoints[0].LbEndpoints[0].Metadata
				fields := endpointMeta.FilterMetadata["envoy.lb"].Fields
				Expect(fields).To(HaveKeyWithValue("testkey", sv("testvalue")))
			})

			It("should add subset to cluster", func() {
				translateWithEndpoints()

				Expect(cluster.LbSubsetConfig).ToNot(BeNil())
				Expect(cluster.LbSubsetConfig.FallbackPolicy).To(Equal(envoyapi.Cluster_LbSubsetConfig_ANY_ENDPOINT))
				Expect(cluster.LbSubsetConfig.SubsetSelectors).To(HaveLen(1))
				Expect(cluster.LbSubsetConfig.SubsetSelectors[0].Keys).To(HaveLen(1))
				Expect(cluster.LbSubsetConfig.SubsetSelectors[0].Keys[0]).To(Equal("testkey"))
			})
			It("should add subset to route", func() {
				translateWithEndpoints()

				metadatamatch := route_configuration.VirtualHosts[0].Routes[0].GetRoute().GetMetadataMatch()
				fields := metadatamatch.FilterMetadata["envoy.lb"].Fields
				Expect(fields).To(HaveKeyWithValue("testkey", sv("testvalue")))
			})
		})

		It("should create empty value if missing labels on the endpoint are provided in the upstream", func() {
			params.Snapshot.Endpoints["gloo-system"][0].Metadata.Labels = nil
			translateWithEndpoints()
			endpointMeta := cla_configuration.Endpoints[0].LbEndpoints[0].Metadata
			Expect(endpointMeta).ToNot(BeNil())
			Expect(endpointMeta.FilterMetadata).To(HaveKey("envoy.lb"))
			fields := endpointMeta.FilterMetadata["envoy.lb"].Fields
			Expect(fields).To(HaveKeyWithValue("testkey", sv("")))
		})

		Context("bad route", func() {
			BeforeEach(func() {
				routes = []*v1.Route{{
					Matcher: matcher,
					Action: &v1.Route_RouteAction{
						RouteAction: &v1.RouteAction{
							Destination: &v1.RouteAction_Single{
								Single: &v1.Destination{
									Upstream: upstream.Metadata.Ref(),
									Subset: &v1.Subset{
										Values: map[string]string{
											"nottestkey": "value",
										},
									},
								},
							},
						},
					},
				}}
			})

			It("should error a route when subset in route doesnt match subset in upstream", func() {
				_, errs, err := translator.Translate(params, proxy)
				Expect(err).NotTo(HaveOccurred())
				Expect(errs.Validate()).To(HaveOccurred())
			})
		})

	})

})

func sv(s string) *types.Value {
	return &types.Value{
		Kind: &types.Value_StringValue{
			StringValue: s,
		},
	}
}
