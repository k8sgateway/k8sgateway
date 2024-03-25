package translator_test

import (
	"context"

	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	routerv3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/router/v3"
	envoy_http_connection_manager_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	"github.com/golang/protobuf/ptypes/wrappers"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	gatewaydefaults "github.com/solo-io/gloo/projects/gateway/pkg/defaults"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	gloov1snap "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/gloosnapshot"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/hcm"
	routerV1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/router"
	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	hcmplugin "github.com/solo-io/gloo/projects/gloo/pkg/plugins/hcm"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins/registry"
	"github.com/solo-io/gloo/projects/gloo/pkg/translator"
	sslutils "github.com/solo-io/gloo/projects/gloo/pkg/utils"
	gloovalidation "github.com/solo-io/gloo/projects/gloo/pkg/utils/validation"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

var _ = Describe("Router filter test", func() {
	// These tests validate the router filter that's generated from the network_filters translator. It
	// would be ideal if that filter could be broken out into its own separate plugin, but for now
	// it's a bit shoehorned into the HTTP connection manager translator

	var (
		ctx    context.Context
		cancel context.CancelFunc

		translatorFactory *translator.ListenerSubsystemTranslatorFactory
	)

	BeforeEach(func() {
		ctx, cancel = context.WithCancel(context.Background())

		// Create a pluginRegistry with a minimal number of plugins
		// This test is not concerned with the functionality of individual plugins
		pluginRegistry := registry.NewPluginRegistry([]plugins.Plugin{
			hcmplugin.NewPlugin(),
		})

		// The translatorFactory expects each of the plugins to be initialized
		// Therefore, to test this component we pre-initialize the plugins
		for _, p := range pluginRegistry.GetPlugins() {
			p.Init(plugins.InitParams{
				Ctx:      ctx,
				Settings: &v1.Settings{},
			})
		}

		translatorFactory = translator.NewListenerSubsystemTranslatorFactory(pluginRegistry, sslutils.NewSslConfigTranslator())
	})

	AfterEach(func() {
		cancel()
	})

	// FIXME remove focus
	FDescribeTable("GetAggregateListenerTranslators (success)",
		func(router *routerV1.Router, assertionHandler func(*routerv3.Router)) {
			listener := &v1.Listener{
				Name:        "aggregate-listener",
				BindAddress: gatewaydefaults.GatewayBindAddress,
				BindPort:    defaults.HttpPort,
				ListenerType: &v1.Listener_AggregateListener{
					AggregateListener: &v1.AggregateListener{
						HttpResources: &v1.AggregateListener_HttpResources{
							HttpOptions: map[string]*v1.HttpListenerOptions{
								"http-options-ref": {
									HttpConnectionManagerSettings: &hcm.HttpConnectionManagerSettings{},
									Router:                        router,
								},
							},
							VirtualHosts: map[string]*v1.VirtualHost{
								"vhost-ref": {
									Name: "virtual-host",
								},
							},
						},
						HttpFilterChains: []*v1.AggregateListener_HttpFilterChain{{
							Matcher:         nil,
							HttpOptionsRef:  "http-options-ref",
							VirtualHostRefs: []string{"vhost-ref"},
						}},
					},
				},
			}
			proxy := &v1.Proxy{
				Metadata: &core.Metadata{
					Name:      "proxy",
					Namespace: defaults.GlooSystem,
				},
				Listeners: []*v1.Listener{listener},
			}

			proxyReport := gloovalidation.MakeReport(proxy)
			listenerReport := proxyReport.GetListenerReports()[0] // 1 Listener -> 1 ListenerReport

			listenerTranslator, routeConfigurationTranslator := translatorFactory.GetAggregateListenerTranslators(
				ctx,
				proxy,
				listener,
				listenerReport)

			params := plugins.Params{
				Ctx: ctx,
				Snapshot: &gloov1snap.ApiSnapshot{
					// To support ssl filter chain
					Secrets: v1.SecretList{createTLSSecret()},
				},
			}
			envoyListener := listenerTranslator.ComputeListener(params)
			_ = routeConfigurationTranslator.ComputeRouteConfiguration(params)

			// Validate that no Errors were encountered during translation
			Expect(gloovalidation.GetProxyError(proxyReport)).NotTo(HaveOccurred())

			By("configuring the envoy router from gloo settings")
			filterChain := envoyListener.GetFilterChains()[0]
			hcmFilter := filterChain.GetFilters()[0]
			_, err := sslutils.AnyToMessage(hcmFilter.GetConfigType().(*envoy_config_listener_v3.Filter_TypedConfig).TypedConfig)
			Expect(err).NotTo(HaveOccurred())

			hcm := &envoy_http_connection_manager_v3.HttpConnectionManager{}
			err = translator.ParseTypedConfig(hcmFilter, hcm)
			Expect(err).NotTo(HaveOccurred())
			Expect(hcm.HttpFilters).To(HaveLen(1))

			routeFilter := hcm.GetHttpFilters()[0]
			typedRouterFilter := routerv3.Router{}
			routeFilter.GetTypedConfig().UnmarshalTo(&typedRouterFilter)
			// Perform assertions on generated Envoy router filter
			assertionHandler(&typedRouterFilter)
		},

		Entry(
			"Set suppress_envoy_headers to true",
			&routerV1.Router{
				SuppressEnvoyHeaders: &wrappers.BoolValue{
					Value: true,
				},
			},
			func(typedRouterFilter *routerv3.Router) {
				Expect(typedRouterFilter.GetSuppressEnvoyHeaders()).To(BeTrue())
			},
		),
		Entry(
			"Set dynamic_stats to false",
			&routerV1.Router{
				DynamicStats: &wrappers.BoolValue{
					Value: false,
				},
			},
			func(typedRouterFilter *routerv3.Router) {
				Expect(typedRouterFilter.GetDynamicStats().GetValue()).To(BeFalse())
			},
		),

		Entry(
			"Leave envoy's dynamic_stats as nil if not specified in gloo",
			&routerV1.Router{},
			func(typedRouterFilter *routerv3.Router) {
				Expect(typedRouterFilter.GetDynamicStats()).To(BeNil())
			},
		),
	)
})
