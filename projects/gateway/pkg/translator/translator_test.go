package translator_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/projects/gateway/pkg/api/v2alpha1"

	. "github.com/solo-io/gloo/projects/gateway/pkg/translator"

	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

const (
	ns  = "gloo-system"
	ns2 = "gloo-system2"
)

var _ = Describe("Translator", func() {
	var (
		snap       *v2alpha1.ApiSnapshot
		translator Translator
	)

	Context("translator", func() {
		BeforeEach(func() {
			translator = NewTranslator([]ListenerFactory{&HttpTranslator{}, &TcpTranslator{}})
			snap = &v2alpha1.ApiSnapshot{
				Gateways: v2alpha1.GatewayList{
					{
						Metadata: core.Metadata{Namespace: ns, Name: "name"},
						GatewayType: &v2alpha1.Gateway_HttpGateway{
							HttpGateway: &v2alpha1.HttpGateway{},
						},
						BindPort: 2,
					},
					{
						Metadata: core.Metadata{Namespace: ns2, Name: "name2"},
						GatewayType: &v2alpha1.Gateway_HttpGateway{
							HttpGateway: &v2alpha1.HttpGateway{},
						},
						BindPort: 2,
					},
				},
				VirtualServices: v1.VirtualServiceList{
					{
						Metadata: core.Metadata{Namespace: ns, Name: "name1"},
						VirtualHost: &gloov1.VirtualHost{
							Domains: []string{"d1.com"},
							Routes: []*gloov1.Route{
								{
									Matcher: &gloov1.Matcher{
										PathSpecifier: &gloov1.Matcher_Prefix{
											Prefix: "/1",
										},
									},
								},
							},
						},
					},
					{
						Metadata: core.Metadata{Namespace: ns, Name: "name2"},
						VirtualHost: &gloov1.VirtualHost{
							Domains: []string{"d2.com"},
							Routes: []*gloov1.Route{
								{
									Matcher: &gloov1.Matcher{
										PathSpecifier: &gloov1.Matcher_Prefix{
											Prefix: "/2",
										},
									},
								},
							},
						},
					},
				},
			}
		})

		It("should translate proxy with default name", func() {

			proxy, errs := translator.Translate(context.Background(), ns, snap)

			Expect(errs).To(HaveLen(3))
			Expect(errs.Validate()).NotTo(HaveOccurred())
			Expect(proxy.Metadata.Name).To(Equal(GatewayProxyName))
			Expect(proxy.Metadata.Namespace).To(Equal(ns))
		})

		It("should translate two gateways with same name (different types) to one proxy with the same name", func() {
			snap.Gateways = append(
				snap.Gateways,
				&v2alpha1.Gateway{
					Metadata: core.Metadata{Namespace: ns, Name: "name2"},
					GatewayType: &v2alpha1.Gateway_TcpGateway{
						TcpGateway: &v2alpha1.TcpGateway{},
					},
				},
			)

			proxy, errs := translator.Translate(context.Background(), ns, snap)

			Expect(errs.Validate()).NotTo(HaveOccurred())
			Expect(proxy.Metadata.Name).To(Equal(GatewayProxyName))
			Expect(proxy.Metadata.Namespace).To(Equal(ns))
			Expect(proxy.Listeners).To(HaveLen(2))
		})

		It("should translate two gateways with same name (and types) to one proxy with the same name", func() {
			snap.Gateways = append(
				snap.Gateways,
				&v2alpha1.Gateway{
					Metadata: core.Metadata{Namespace: ns, Name: "name2"},
					GatewayType: &v2alpha1.Gateway_HttpGateway{
						HttpGateway: &v2alpha1.HttpGateway{},
					},
				},
			)

			proxy, errs := translator.Translate(context.Background(), ns, snap)

			Expect(errs.Validate()).NotTo(HaveOccurred())
			Expect(proxy.Metadata.Name).To(Equal(GatewayProxyName))
			Expect(proxy.Metadata.Namespace).To(Equal(ns))
			Expect(proxy.Listeners).To(HaveLen(2))
		})

		It("should error on two gateways with the same port in the same namespace", func() {
			dupeGateway := v2alpha1.Gateway{
				Metadata: core.Metadata{Namespace: ns, Name: "name2"},
				BindPort: 2,
			}
			snap.Gateways = append(snap.Gateways, &dupeGateway)

			_, errs := translator.Translate(context.Background(), ns, snap)
			err := errs.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("bind-address :2 is not unique in a proxy. gateways: gloo-system.name,gloo-system.name2"))
		})

	})

	Context("http", func() {
		var (
			factory *HttpTranslator
		)

		BeforeEach(func() {
			factory = &HttpTranslator{}
			translator = NewTranslator([]ListenerFactory{factory})
			snap = &v2alpha1.ApiSnapshot{
				Gateways: v2alpha1.GatewayList{
					{
						Metadata: core.Metadata{Namespace: ns, Name: "name"},
						GatewayType: &v2alpha1.Gateway_HttpGateway{
							HttpGateway: &v2alpha1.HttpGateway{},
						},
						BindPort: 2,
					},
					{
						Metadata: core.Metadata{Namespace: ns2, Name: "name2"},
						GatewayType: &v2alpha1.Gateway_HttpGateway{
							HttpGateway: &v2alpha1.HttpGateway{},
						},
						BindPort: 2,
					},
				},
				VirtualServices: v1.VirtualServiceList{
					{
						Metadata: core.Metadata{Namespace: ns, Name: "name1"},
						VirtualHost: &gloov1.VirtualHost{
							Domains: []string{"d1.com"},
							Routes: []*gloov1.Route{
								{
									Matcher: &gloov1.Matcher{
										PathSpecifier: &gloov1.Matcher_Prefix{
											Prefix: "/1",
										},
									},
								},
							},
						},
					},
					{
						Metadata: core.Metadata{Namespace: ns, Name: "name2"},
						VirtualHost: &gloov1.VirtualHost{
							Domains: []string{"d2.com"},
							Routes: []*gloov1.Route{
								{
									Matcher: &gloov1.Matcher{
										PathSpecifier: &gloov1.Matcher_Prefix{
											Prefix: "/2",
										},
									},
								},
							},
						},
					},
				},
			}
		})

		It("should translate an empty gateway to have all vservices", func() {

			proxy, _ := translator.Translate(context.Background(), ns, snap)

			Expect(proxy.Listeners).To(HaveLen(1))
			listener := proxy.Listeners[0].ListenerType.(*gloov1.Listener_HttpListener).HttpListener
			Expect(listener.VirtualHosts).To(HaveLen(2))
		})

		It("should have no ssl config", func() {
			proxy, _ := translator.Translate(context.Background(), ns, snap)

			Expect(proxy.Listeners).To(HaveLen(1))
			Expect(proxy.Listeners[0].SslConfigurations).To(BeEmpty())
		})

		It("should translate a gateway to only have its vservices", func() {
			snap.Gateways[0].GatewayType = &v2alpha1.Gateway_HttpGateway{
				HttpGateway: &v2alpha1.HttpGateway{
					VirtualServices: []core.ResourceRef{snap.VirtualServices[0].Metadata.Ref()},
				},
			}

			proxy, errs := translator.Translate(context.Background(), ns, snap)

			Expect(errs.Validate()).NotTo(HaveOccurred())
			Expect(proxy).NotTo(BeNil())
			Expect(proxy.Listeners).To(HaveLen(1))
			listener := proxy.Listeners[0].ListenerType.(*gloov1.Listener_HttpListener).HttpListener
			Expect(listener.VirtualHosts).To(HaveLen(1))
		})

		It("should not have vhosts with ssl", func() {
			snap.VirtualServices[0].SslConfig = new(gloov1.SslConfig)

			proxy, errs := translator.Translate(context.Background(), ns, snap)

			Expect(errs.Validate()).NotTo(HaveOccurred())

			Expect(proxy.Listeners).To(HaveLen(1))
			listener := proxy.Listeners[0].ListenerType.(*gloov1.Listener_HttpListener).HttpListener
			Expect(listener.VirtualHosts).To(HaveLen(1))
			Expect(listener.VirtualHosts[0].Name).To(ContainSubstring("name2"))
		})

		It("should not have vhosts without ssl", func() {
			snap.Gateways[0].Ssl = true
			snap.VirtualServices[0].SslConfig = new(gloov1.SslConfig)

			proxy, errs := translator.Translate(context.Background(), ns, snap)

			Expect(errs.Validate()).NotTo(HaveOccurred())

			Expect(proxy.Listeners).To(HaveLen(1))
			listener := proxy.Listeners[0].ListenerType.(*gloov1.Listener_HttpListener).HttpListener
			Expect(listener.VirtualHosts).To(HaveLen(1))
			Expect(listener.VirtualHosts[0].Name).To(ContainSubstring("name1"))
		})

		Context("merge", func() {
			BeforeEach(func() {
				snap.VirtualServices[1].VirtualHost.Domains = snap.VirtualServices[0].VirtualHost.Domains
			})

			It("should translate 2 virtual services with the same domains to 1 virtual service", func() {

				proxy, errs := translator.Translate(context.Background(), ns, snap)

				Expect(errs.Validate()).NotTo(HaveOccurred())
				Expect(proxy.Metadata.Name).To(Equal(GatewayProxyName))
				Expect(proxy.Metadata.Namespace).To(Equal(ns))
				Expect(proxy.Listeners).To(HaveLen(1))
				listener := proxy.Listeners[0].ListenerType.(*gloov1.Listener_HttpListener).HttpListener
				Expect(listener.VirtualHosts).To(HaveLen(1))
			})

			It("should translate 2 virtual services with the empty domains", func() {
				snap.VirtualServices[1].VirtualHost.Domains = nil
				snap.VirtualServices[0].VirtualHost.Domains = nil

				proxy, errs := translator.Translate(context.Background(), ns, snap)

				Expect(errs.Validate()).NotTo(HaveOccurred())
				Expect(proxy.Listeners).To(HaveLen(1))
				listener := proxy.Listeners[0].ListenerType.(*gloov1.Listener_HttpListener).HttpListener
				Expect(listener.VirtualHosts).To(HaveLen(1))
				Expect(listener.VirtualHosts[0].Name).NotTo(BeEmpty())
				Expect(listener.VirtualHosts[0].Name).NotTo(Equal(ns + "."))
			})

			It("should not error with one contains plugins", func() {
				snap.VirtualServices[0].VirtualHost.VirtualHostPlugins = new(gloov1.VirtualHostPlugins)

				_, errs := translator.Translate(context.Background(), ns, snap)

				Expect(errs.Validate()).NotTo(HaveOccurred())
			})

			It("should error with both having plugins", func() {
				snap.VirtualServices[0].VirtualHost.VirtualHostPlugins = new(gloov1.VirtualHostPlugins)
				snap.VirtualServices[1].VirtualHost.VirtualHostPlugins = new(gloov1.VirtualHostPlugins)

				_, errs := translator.Translate(context.Background(), ns, snap)

				Expect(errs.Validate()).To(HaveOccurred())
			})

			It("should not error with one contains ssl config", func() {
				snap.VirtualServices[0].SslConfig = new(gloov1.SslConfig)

				proxy, errs := translator.Translate(context.Background(), ns, snap)

				Expect(errs.Validate()).NotTo(HaveOccurred())
				listener := proxy.Listeners[0].ListenerType.(*gloov1.Listener_HttpListener).HttpListener
				Expect(listener.VirtualHosts).To(HaveLen(1))
			})

			It("should not error with one contains ssl config", func() {
				snap.Gateways[0].Ssl = true
				snap.VirtualServices[0].SslConfig = new(gloov1.SslConfig)

				proxy, errs := translator.Translate(context.Background(), ns, snap)

				Expect(errs.Validate()).NotTo(HaveOccurred())
				listener := proxy.Listeners[0].ListenerType.(*gloov1.Listener_HttpListener).HttpListener
				Expect(listener.VirtualHosts).To(HaveLen(1))
				Expect(listener.VirtualHosts[0].Routes).To(HaveLen(1))
			})

			It("should error when two virtual services conflict", func() {
				snap.Gateways[0].Ssl = true
				snap.VirtualServices[0].SslConfig = new(gloov1.SslConfig)
				snap.VirtualServices[1].SslConfig = new(gloov1.SslConfig)
				snap.VirtualServices[0].SslConfig.SniDomains = []string{"bar"}
				snap.VirtualServices[1].SslConfig.SniDomains = []string{"foo"}

				_, errs := translator.Translate(context.Background(), ns, snap)

				Expect(errs.Validate()).To(HaveOccurred())
			})

			It("should error when two virtual services conflict", func() {
				snap.Gateways[0].Ssl = true
				snap.VirtualServices[0].SslConfig = new(gloov1.SslConfig)
				snap.VirtualServices[1].SslConfig = new(gloov1.SslConfig)
				snap.VirtualServices[0].SslConfig.SniDomains = []string{"foo"}
				snap.VirtualServices[1].SslConfig.SniDomains = []string{"foo"}

				_, errs := translator.Translate(context.Background(), ns, snap)

				Expect(errs.Validate()).NotTo(HaveOccurred())
			})

		})
	})

	Context("tcp", func() {
		var (
			factory *TcpTranslator
		)
		BeforeEach(func() {
			factory = &TcpTranslator{}
			translator = NewTranslator([]ListenerFactory{factory})
			snap = &v2alpha1.ApiSnapshot{
				Gateways: v2alpha1.GatewayList{
					{
						Metadata: core.Metadata{Namespace: ns, Name: "name"},
						GatewayType: &v2alpha1.Gateway_TcpGateway{
							TcpGateway: &v2alpha1.TcpGateway{},
						},
						BindPort: 2,
					},
				},
			}
		})

		It("can properly translate a tcp proxy", func() {

		})

	})

})
