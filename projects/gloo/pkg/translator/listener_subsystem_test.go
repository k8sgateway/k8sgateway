package translator_test

import (
	"context"
	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	envoyhttp "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	errors "github.com/rotisserie/eris"
	gatewaydefaults "github.com/solo-io/gloo/projects/gateway/pkg/defaults"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/grpc/validation"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	gloov1snap "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/gloosnapshot"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/hcm"
	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins/buffer"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins/registry"
	"github.com/solo-io/gloo/projects/gloo/pkg/translator"
	sslutils "github.com/solo-io/gloo/projects/gloo/pkg/utils"
	gloovalidation "github.com/solo-io/gloo/projects/gloo/pkg/utils/validation"
	gloohelpers "github.com/solo-io/gloo/test/helpers"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Allow each test to define its own set of assertions
// based on the envoy types returned by executing the ListenerTranslator and RouteConfigurationTranslator
type ResourceAssertionHandler func(
	listener *envoy_config_listener_v3.Listener,
	routeConfigurations []*envoy_config_route_v3.RouteConfiguration)

type ReportAssertionHandler func(
	proxyReport *validation.ProxyReport)

var _ = Describe("Listener Subsystem", func() {

	// These tests validate that the ListenerSubsystemTranslatorFactory produces Translators
	// which in turn create Envoy Listeners and RouteConfigurations with expected values
	// The tests are non-exhaustive, as we expect each translator to more rigorously test the
	// edge cases. Instead, these tests focus on the high level Envoy resources that are created.

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
			buffer.NewPlugin(),
		})

		// The translatorFactory expects each of the plugins to be initialized
		// Therefore, to test this component we pre-initialize the plugins
		for _, p := range pluginRegistry.GetPlugins() {
			if err := p.Init(plugins.InitParams{
				Ctx:      ctx,
				Settings: &v1.Settings{},
			}); err != nil {
				panic(errors.Wrapf(err, "test is not initialized properly"))
			}
		}

		translatorFactory = translator.NewListenerSubsystemTranslatorFactory(pluginRegistry, sslutils.NewSslConfigTranslator())
	})

	AfterEach(func() {
		cancel()
	})

	DescribeTable("GetAggregateListenerTranslators (success)",
		func(aggregateListener *v1.AggregateListener, assertionHandler ResourceAssertionHandler) {
			listener := &v1.Listener{
				Name:        "aggregate-listener",
				BindAddress: gatewaydefaults.GatewayBindAddress,
				BindPort:    defaults.HttpPort,
				ListenerType: &v1.Listener_AggregateListener{
					AggregateListener: aggregateListener,
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
			envoyRouteConfigs := routeConfigurationTranslator.ComputeRouteConfiguration(params)

			// Validate that no Errors were encountered during translation
			Expect(gloovalidation.GetProxyError(proxyReport)).NotTo(HaveOccurred())

			// Validate the ResourceAssertionHandler defined by each test
			assertionHandler(envoyListener, envoyRouteConfigs)
		},
		Entry(
			"0 filter chains",
			&v1.AggregateListener{
				HttpResources:    &v1.AggregateListener_HttpResources{},
				HttpFilterChains: []*v1.AggregateListener_HttpFilterChain{},
			},
			func(listener *envoy_config_listener_v3.Listener, routeConfigs []*envoy_config_route_v3.RouteConfiguration) {
				ExpectWithOffset(1, listener.GetFilterChains()).To(HaveLen(0))
				ExpectWithOffset(1, routeConfigs).To(HaveLen(0))
			},
		),
		Entry(
			"1 insecure filter chain",
			&v1.AggregateListener{
				HttpResources: &v1.AggregateListener_HttpResources{
					HttpOptions: map[string]*v1.HttpListenerOptions{
						"http-options-ref": {
							HttpConnectionManagerSettings: &hcm.HttpConnectionManagerSettings{},
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
			func(listener *envoy_config_listener_v3.Listener, routeConfigs []*envoy_config_route_v3.RouteConfiguration) {
				By("1 insecure filter chain")
				ExpectWithOffset(1, listener.GetFilterChains()).To(HaveLen(1))
				filterChain := listener.GetFilterChains()[0]
				ExpectWithOffset(1, filterChain.GetFilterChainMatch()).To(BeNil())

				By("with hcm network filter")
				hcmFilter := filterChain.GetFilters()[0]
				typedConfig, err := sslutils.AnyToMessage(hcmFilter.GetConfigType().(*envoy_config_listener_v3.Filter_TypedConfig).TypedConfig)
				ExpectWithOffset(1, err).NotTo(HaveOccurred())
				hcm := typedConfig.(*envoyhttp.HttpConnectionManager)
				hcmRouteConfigName := hcm.GetRds().GetRouteConfigName()

				By("1 route configuration, with name matching HCM")
				ExpectWithOffset(1, routeConfigs).To(HaveLen(1))
				routeConfig := routeConfigs[0]
				ExpectWithOffset(1, routeConfig.GetName()).To(Equal(hcmRouteConfigName))
			},
		),
		Entry(
			"1 secure filter chain",
			&v1.AggregateListener{
				HttpResources: &v1.AggregateListener_HttpResources{
					HttpOptions: map[string]*v1.HttpListenerOptions{
						"http-options-ref": {
							HttpConnectionManagerSettings: &hcm.HttpConnectionManagerSettings{},
						},
					},
					VirtualHosts: map[string]*v1.VirtualHost{
						"vhost-ref": {
							Name: "virtual-host",
						},
					},
				},
				HttpFilterChains: []*v1.AggregateListener_HttpFilterChain{{
					Matcher: &v1.Matcher{
						SslConfig: &v1.SslConfig{
							SniDomains:    []string{"sni-domain"},
							AlpnProtocols: []string{"h2"},
							SslSecrets: &v1.SslConfig_SecretRef{
								SecretRef: createTLSSecret().GetMetadata().Ref(),
							},
						},
					},
					HttpOptionsRef:  "http-options-ref",
					VirtualHostRefs: []string{"vhost-ref"},
				}},
			},
			func(listener *envoy_config_listener_v3.Listener, routeConfigs []*envoy_config_route_v3.RouteConfiguration) {
				By("1 secure filter chain")
				ExpectWithOffset(1, listener.GetFilterChains()).To(HaveLen(1))
				filterChain := listener.GetFilterChains()[0]
				ExpectWithOffset(1, filterChain.GetFilterChainMatch()).To(Equal(&envoy_config_listener_v3.FilterChainMatch{
					ServerNames: []string{"sni-domain"},
				}))

				By("with hcm network filter")
				hcmFilter := filterChain.GetFilters()[0]
				typedConfig, err := sslutils.AnyToMessage(hcmFilter.GetConfigType().(*envoy_config_listener_v3.Filter_TypedConfig).TypedConfig)
				ExpectWithOffset(1, err).NotTo(HaveOccurred())
				hcm := typedConfig.(*envoyhttp.HttpConnectionManager)
				hcmRouteConfigName := hcm.GetRds().GetRouteConfigName()

				By("1 route configuration, with name matching HCM")
				ExpectWithOffset(1, routeConfigs).To(HaveLen(1))
				routeConfig := routeConfigs[0]
				ExpectWithOffset(1, routeConfig.GetName()).To(Equal(hcmRouteConfigName))
			},
		),
		Entry(
			"multiple secure filter chains",
			&v1.AggregateListener{
				HttpResources: &v1.AggregateListener_HttpResources{
					HttpOptions: map[string]*v1.HttpListenerOptions{
						"http-options-ref": {
							HttpConnectionManagerSettings: &hcm.HttpConnectionManagerSettings{},
						},
					},
					VirtualHosts: map[string]*v1.VirtualHost{
						"vhost-ref": {
							Name: "virtual-host",
						},
					},
				},
				HttpFilterChains: []*v1.AggregateListener_HttpFilterChain{
					{
						Matcher: &v1.Matcher{
							SslConfig: &v1.SslConfig{
								SniDomains:    []string{"sni-domain"},
								AlpnProtocols: []string{"h2"},
								SslSecrets: &v1.SslConfig_SecretRef{
									SecretRef: createTLSSecret().GetMetadata().Ref(),
								},
							},
						},
						HttpOptionsRef:  "http-options-ref",
						VirtualHostRefs: []string{"vhost-ref"},
					},
					{
						Matcher: &v1.Matcher{
							SslConfig: &v1.SslConfig{
								SniDomains:    []string{"other-sni-domain"},
								AlpnProtocols: []string{"h2"},
								SslSecrets: &v1.SslConfig_SecretRef{
									SecretRef: createTLSSecret().GetMetadata().Ref(),
								},
							},
						},
						HttpOptionsRef:  "http-options-ref",
						VirtualHostRefs: []string{"vhost-ref"},
					}},
			},
			func(listener *envoy_config_listener_v3.Listener, routeConfigs []*envoy_config_route_v3.RouteConfiguration) {
				By("2 secure filter chains and route configurations")
				ExpectWithOffset(1, listener.GetFilterChains()).To(HaveLen(2))
				ExpectWithOffset(1, routeConfigs).To(HaveLen(2))

				By("first filter chain")
				filterChain := listener.GetFilterChains()[0]
				ExpectWithOffset(1, filterChain.GetFilterChainMatch()).To(Equal(&envoy_config_listener_v3.FilterChainMatch{
					ServerNames: []string{"sni-domain"},
				}))

				By("with hcm network filter")
				hcmFilter := filterChain.GetFilters()[0]
				typedConfig, err := sslutils.AnyToMessage(hcmFilter.GetConfigType().(*envoy_config_listener_v3.Filter_TypedConfig).TypedConfig)
				ExpectWithOffset(1, err).NotTo(HaveOccurred())
				hcm := typedConfig.(*envoyhttp.HttpConnectionManager)
				hcmRouteConfigName := hcm.GetRds().GetRouteConfigName()

				By("route config name matches HCM")
				routeConfig := routeConfigs[0]
				ExpectWithOffset(1, routeConfig.GetName()).To(Equal(hcmRouteConfigName))

				By("second filter chain")
				filterChain = listener.GetFilterChains()[1]
				ExpectWithOffset(1, filterChain.GetFilterChainMatch()).To(Equal(&envoy_config_listener_v3.FilterChainMatch{
					ServerNames: []string{"other-sni-domain"},
				}))

				By("with hcm network filter")
				hcmFilter = filterChain.GetFilters()[0]
				typedConfig, err = sslutils.AnyToMessage(hcmFilter.GetConfigType().(*envoy_config_listener_v3.Filter_TypedConfig).TypedConfig)
				ExpectWithOffset(1, err).NotTo(HaveOccurred())
				hcm = typedConfig.(*envoyhttp.HttpConnectionManager)
				hcmRouteConfigName = hcm.GetRds().GetRouteConfigName()

				By("route config name matches HCM")
				routeConfig = routeConfigs[1]
				ExpectWithOffset(1, routeConfig.GetName()).To(Equal(hcmRouteConfigName))
			},
		),
	)

	DescribeTable("GetAggregateListenerTranslators (failure)",
		func(aggregateListener *v1.AggregateListener, assertionHandler ReportAssertionHandler) {
			listener := &v1.Listener{
				Name:        "aggregate-listener",
				BindAddress: gatewaydefaults.GatewayBindAddress,
				BindPort:    defaults.HttpPort,
				ListenerType: &v1.Listener_AggregateListener{
					AggregateListener: aggregateListener,
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
				Snapshot: &gloov1snap.ApiSnapshot{},
			}
			_ = listenerTranslator.ComputeListener(params)
			_ = routeConfigurationTranslator.ComputeRouteConfiguration(params)

			// Validate the ReportAssertionHandler defined by each test
			assertionHandler(proxyReport)
		},
		Entry(
			"listener error",
			&v1.AggregateListener{
				HttpResources: &v1.AggregateListener_HttpResources{
					HttpOptions: map[string]*v1.HttpListenerOptions{
						"http-options-ref": {
							HttpConnectionManagerSettings: &hcm.HttpConnectionManagerSettings{},
						},
					},
					VirtualHosts: map[string]*v1.VirtualHost{
						"vhost-ref": {
							Name: "virtual-host",
						},
					},
				},
				HttpFilterChains: []*v1.AggregateListener_HttpFilterChain{{
					Matcher: &v1.Matcher{
						SslConfig: &v1.SslConfig{
							SslSecrets: &v1.SslConfig_SecretRef{
								SecretRef: &core.ResourceRef{
									Name: "secret-that-is-not-in-snapshot",
									Namespace: defaults.GlooSystem,
								},
							},
						},
					},
					HttpOptionsRef:  "http-options-ref",
					VirtualHostRefs: []string{"vhost-ref"},
				}},
			},
			func(proxyReport *validation.ProxyReport) {
				proxyErr := gloovalidation.GetProxyError(proxyReport)
				Expect(proxyErr).To(HaveOccurred())
				Expect(proxyErr.Error()).To(ContainSubstring(validation.ListenerReport_Error_SSLConfigError.String()))
			},
		),
	)
})

func createTLSSecret() *v1.Secret {
	return &v1.Secret{
		Metadata: &core.Metadata{
			Name:      "tls",
			Namespace: defaults.GlooSystem,
		},
		Kind: &v1.Secret_Tls{
			Tls: &v1.TlsSecret{
				CertChain:  gloohelpers.Certificate(),
				PrivateKey: gloohelpers.PrivateKey(),
			},
		},
	}
}
