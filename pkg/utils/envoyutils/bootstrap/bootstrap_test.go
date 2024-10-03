package bootstrap

import (
	"log"

	envoy_config_bootstrap_v3 "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v3"
	envoy_config_cluster_v3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	envoy_config_endpoint_v3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	envoy_extensions_filters_network_http_connection_manager_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	"github.com/envoyproxy/go-control-plane/pkg/wellknown"
	anypb "github.com/golang/protobuf/ptypes/any"
	"github.com/solo-io/gloo/pkg/utils/protoutils"
	envoytransformation "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/transformation"
	"github.com/solo-io/gloo/projects/gloo/pkg/utils"
	"google.golang.org/protobuf/proto"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	hcmTypeUrl = "type.googleapis.com/envoy.extensions.filters.network.http_connection_manager.v3.HttpConnectionManager"
)

var _ = Describe("Static bootstrap generation", func() {
	Context("Util functions", func() {
		var (
			routedCluster map[string]struct{}
			listeners     []*envoy_config_listener_v3.Listener
			clusters      []*envoy_config_cluster_v3.Cluster
			routes        []*envoy_config_route_v3.RouteConfiguration
			endpoints     []*envoy_config_endpoint_v3.ClusterLoadAssignment
		)
		BeforeEach(func() {
			routedCluster = make(map[string]struct{})
			listeners = []*envoy_config_listener_v3.Listener{}
			clusters = []*envoy_config_cluster_v3.Cluster{}
			endpoints = []*envoy_config_endpoint_v3.ClusterLoadAssignment{}
			routes = []*envoy_config_route_v3.RouteConfiguration{{
				Name: "foo-routes",
				VirtualHosts: []*envoy_config_route_v3.VirtualHost{
					{
						Name:    "placeholder_host",
						Domains: []string{"*"},
						Routes: []*envoy_config_route_v3.Route{
							{
								Action: &envoy_config_route_v3.Route_Route{
									Route: &envoy_config_route_v3.RouteAction{
										ClusterSpecifier: &envoy_config_route_v3.RouteAction_Cluster{
											Cluster: "foo",
										},
									},
								},
								Name: "foo-route",
							},
							{
								Action: &envoy_config_route_v3.Route_Route{
									Route: &envoy_config_route_v3.RouteAction{
										ClusterSpecifier: &envoy_config_route_v3.RouteAction_Cluster{
											Cluster: "bar",
										},
									},
								},
								Name: "bar-route",
							},
						},
					},
				},
			},
			}
		})
		Context("extractRoutedClustersFromListeners", func() {
			It("does not error if no hcm", func() {
				l := &envoy_config_listener_v3.Listener{
					Name:    "fake-listener",
					Address: &envoy_config_core_v3.Address{},
					FilterChains: []*envoy_config_listener_v3.FilterChain{{
						FilterChainMatch: &envoy_config_listener_v3.FilterChainMatch{},
						Filters:          []*envoy_config_listener_v3.Filter{},
					}},
				}
				listeners = append(listeners, l)
				Expect(extractRoutedClustersFromListeners(routedCluster, listeners, routes)).NotTo(HaveOccurred())
				Expect(routedCluster).To(BeEmpty())
			})
			It("extracts a single happy cluster", func() {
				hcmAny, err := utils.MessageToAny(&envoy_extensions_filters_network_http_connection_manager_v3.HttpConnectionManager{
					StatPrefix: "placeholder",
					RouteSpecifier: &envoy_extensions_filters_network_http_connection_manager_v3.HttpConnectionManager_Rds{
						Rds: &envoy_extensions_filters_network_http_connection_manager_v3.Rds{
							RouteConfigName: "foo-routes",
						},
					},
				})
				Expect(err).NotTo(HaveOccurred())
				l := &envoy_config_listener_v3.Listener{
					Name:    "fake-listener",
					Address: &envoy_config_core_v3.Address{},
					FilterChains: []*envoy_config_listener_v3.FilterChain{{
						FilterChainMatch: &envoy_config_listener_v3.FilterChainMatch{},
						Filters: []*envoy_config_listener_v3.Filter{{
							Name: wellknown.HTTPConnectionManager,
							ConfigType: &envoy_config_listener_v3.Filter_TypedConfig{
								TypedConfig: hcmAny,
							},
						}},
					}},
				}
				listeners = append(listeners, l)
				Expect(extractRoutedClustersFromListeners(routedCluster, listeners, routes)).NotTo(HaveOccurred())
				Expect(routedCluster).To(HaveKey("foo"))
			})
		})
		Context("convertToStaticClusters", func() {
			BeforeEach(func() {
				routedCluster = map[string]struct{}{"foo": struct{}{}, "bar": struct{}{}}
				clusters = []*envoy_config_cluster_v3.Cluster{{
					Name: "foo",
					EdsClusterConfig: &envoy_config_cluster_v3.Cluster_EdsClusterConfig{
						ServiceName: "foo-eds",
					},
				}}
				endpoints = []*envoy_config_endpoint_v3.ClusterLoadAssignment{{
					ClusterName: "foo-eds",
				}}
			})
			It("converts and removes from routedCluster", func() {
				Expect(clusters[0].GetLoadAssignment()).To(BeNil())
				convertToStaticClusters(routedCluster, clusters, endpoints)
				Expect(routedCluster).To(HaveKey("bar"))
				Expect(routedCluster).NotTo(HaveKey("foo"))
				Expect(clusters).To(HaveLen(1))
				Expect(clusters[0].GetLoadAssignment()).NotTo(BeNil())
				Expect(clusters[0].GetLoadAssignment().GetClusterName()).To(Equal("foo-eds"))
			})
		})
		Context("addBlackholeClusters", func() {
			BeforeEach(func() {
				routedCluster = map[string]struct{}{"bar": struct{}{}}
				clusters = []*envoy_config_cluster_v3.Cluster{{
					Name: "foo",
					EdsClusterConfig: &envoy_config_cluster_v3.Cluster_EdsClusterConfig{
						ServiceName: "foo-eds",
					},
				}}
				endpoints = []*envoy_config_endpoint_v3.ClusterLoadAssignment{{
					ClusterName: "foo-eds",
				}}
			})
			It("adds blackhole clusters for missing values", func() {
				clusters = addBlackholeClusters(routedCluster, clusters)
				Expect(clusters).To(HaveLen(2))
				Expect(clusters[1].GetName()).To(Equal("bar"))
				Expect(clusters[1].GetLoadAssignment().GetClusterName()).To(Equal("bar"))
			})
		})
		Context("getHcmForFilterChain", func() {
			It("gets the HCM", func() {
				hcmAny, err := utils.MessageToAny(&envoy_extensions_filters_network_http_connection_manager_v3.HttpConnectionManager{
					StatPrefix: "placeholder",
					RouteSpecifier: &envoy_extensions_filters_network_http_connection_manager_v3.HttpConnectionManager_Rds{
						Rds: &envoy_extensions_filters_network_http_connection_manager_v3.Rds{
							RouteConfigName: "foo-routes",
						},
					},
				})
				Expect(err).NotTo(HaveOccurred())
				fc := &envoy_config_listener_v3.FilterChain{
					FilterChainMatch: &envoy_config_listener_v3.FilterChainMatch{},
					Filters: []*envoy_config_listener_v3.Filter{{
						Name: wellknown.HTTPConnectionManager,
						ConfigType: &envoy_config_listener_v3.Filter_TypedConfig{
							TypedConfig: hcmAny,
						},
					}},
				}
				hcm, filter, err := getHcmForFilterChain(fc)
				Expect(err).NotTo(HaveOccurred())
				Expect(hcm.StatPrefix).To(Equal("placeholder"))
				Expect(filter.GetTypedConfig().GetTypeUrl()).To(Equal(hcmTypeUrl))
			})
		})
		Context("findTargetedClusters", func() {
			It("finds clusters targeted by routes", func() {
				findTargetedClusters(routes[0], routedCluster)
				Expect(routedCluster).To(HaveLen(2))
				Expect(routedCluster).To(HaveKey("foo"))
				Expect(routedCluster).To(HaveKey("bar"))
			})
		})
		Context("setStaticRouteConfig", func() {
			It("sets the route config as static and mutates the filter", func() {
				f := &envoy_config_listener_v3.Filter{}
				hcm := &envoy_extensions_filters_network_http_connection_manager_v3.HttpConnectionManager{}
				Expect(setStaticRouteConfig(f, hcm, routes[0])).NotTo(HaveOccurred())
				Expect(hcm.GetRouteConfig().GetName()).To(Equal(routes[0].GetName()))
				Expect(f.GetTypedConfig().GetTypeUrl()).To(Equal(hcmTypeUrl))
			})
		})
	})
	Context("From Filter", func() {
		It("produces correct bootstrap", func() {
			Skip("TODO")
			inTransformation := &envoytransformation.RouteTransformations{
				ClearRouteCache: true,
				Transformations: []*envoytransformation.RouteTransformations_RouteTransformation{
					{
						Match: &envoytransformation.RouteTransformations_RouteTransformation_RequestMatch_{
							RequestMatch: &envoytransformation.RouteTransformations_RouteTransformation_RequestMatch{ClearRouteCache: true},
						},
					},
				},
			}

			filterName := "transformation"
			actual, err := FromFilter(filterName, inTransformation)
			Expect(err).NotTo(HaveOccurred())

			expectedBootstrap := &envoy_config_bootstrap_v3.Bootstrap{
				Node: &envoy_config_core_v3.Node{
					Id:      "validation-node-id",
					Cluster: "validation-cluster",
				},
				StaticResources: &envoy_config_bootstrap_v3.Bootstrap_StaticResources{
					Listeners: []*envoy_config_listener_v3.Listener{{
						Name: "placeholder_listener",
						Address: &envoy_config_core_v3.Address{
							Address: &envoy_config_core_v3.Address_SocketAddress{SocketAddress: &envoy_config_core_v3.SocketAddress{
								Address:       "0.0.0.0",
								PortSpecifier: &envoy_config_core_v3.SocketAddress_PortValue{PortValue: 8081},
							}},
						},
						FilterChains: []*envoy_config_listener_v3.FilterChain{
							{
								Name: "placeholder_filter_chain",
								Filters: []*envoy_config_listener_v3.Filter{
									{
										Name: wellknown.HTTPConnectionManager,
										ConfigType: &envoy_config_listener_v3.Filter_TypedConfig{
											TypedConfig: func() *anypb.Any {
												hcmAny, err := utils.MessageToAny(&envoy_extensions_filters_network_http_connection_manager_v3.HttpConnectionManager{
													StatPrefix: "placeholder",
													RouteSpecifier: &envoy_extensions_filters_network_http_connection_manager_v3.HttpConnectionManager_RouteConfig{
														RouteConfig: &envoy_config_route_v3.RouteConfiguration{
															VirtualHosts: []*envoy_config_route_v3.VirtualHost{
																{
																	Name:    "placeholder_host",
																	Domains: []string{"*"},
																	TypedPerFilterConfig: map[string]*anypb.Any{
																		filterName: {
																			TypeUrl: "type.googleapis.com/envoy.api.v2.filter.http.RouteTransformations",
																			Value: func() []byte {
																				tformany, err := utils.MessageToAny(inTransformation)
																				Expect(err).NotTo(HaveOccurred())
																				return tformany.GetValue()
																			}(),
																		},
																	},
																},
															},
														},
													},
												})
												Expect(err).NotTo(HaveOccurred())
												return hcmAny
											}(),
										},
									},
								},
							},
						},
					}},
				},
			}

			var actualBootstrap *envoy_config_bootstrap_v3.Bootstrap

			log.Println(actual)
			err = protoutils.UnmarshalBytesAllowUnknown([]byte(actual), actualBootstrap)
			// err = (&jsonpb.Unmarshaler{
			// 	AllowUnknownFields: true,
			// 	AnyResolver:        nil,
			// }).UnmarshalString(actual, actualBootstrap)
			// err = jsonpb.Unmarshal(bytes.NewBuffer([]byte(actual)), actualBootstrap)
			Expect(err).NotTo(HaveOccurred())

			Expect(proto.Equal(expectedBootstrap, actualBootstrap)).To(BeTrue())
		})
	})
})
