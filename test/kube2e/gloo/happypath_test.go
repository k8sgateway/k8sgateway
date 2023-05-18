package gloo_test

import (
	"context"
	"fmt"
	"net"
	"strings"

	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	"github.com/golang/protobuf/ptypes/wrappers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	gatewaydefaults "github.com/solo-io/gloo/projects/gateway/pkg/defaults"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	testhelpers "github.com/solo-io/gloo/test/helpers"
	"github.com/solo-io/gloo/test/services"
	"github.com/solo-io/gloo/test/v1helpers"
	"github.com/solo-io/k8s-utils/kubeutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	skkubeutils "github.com/solo-io/solo-kit/pkg/utils/kubeutils"
	"github.com/solo-io/solo-kit/test/helpers"
	"github.com/solo-io/solo-kit/test/setup"
	kubev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var _ = Describe("Happy path", func() {

	var (
		ctx           context.Context
		cancel        context.CancelFunc
		testClients   services.TestClients
		envoyInstance *services.EnvoyInstance
		tu            *v1helpers.TestUpstream
		envoyPort     uint32

		testCases = []struct {
			Title               string
			RestEdsEnabled      *wrappers.BoolValue
			TransportApiVersion envoy_config_core_v3.ApiVersion
		}{
			{
				Title: "Rest Eds Enabled",
				RestEdsEnabled: &wrappers.BoolValue{
					Value: true,
				},
				TransportApiVersion: envoy_config_core_v3.ApiVersion_V3,
			},
			{
				Title: "Rest Eds Disabled",
				RestEdsEnabled: &wrappers.BoolValue{
					Value: false,
				},
				TransportApiVersion: envoy_config_core_v3.ApiVersion_V3,
			},
		}
	)

	BeforeEach(func() {
		ctx, cancel = context.WithCancel(context.Background())
		defaults.HttpPort = services.NextBindPort()
		defaults.HttpsPort = services.NextBindPort()

		var err error
		envoyInstance, err = envoyFactory.NewEnvoyInstance()
		Expect(err).NotTo(HaveOccurred())

		tu = v1helpers.NewTestHttpUpstream(ctx, envoyInstance.LocalAddr())
		envoyPort = defaults.HttpPort
	})

	AfterEach(func() {
		envoyInstance.Clean()
		cancel()
	})

	TestUpstreamReachable := func() {
		v1helpers.TestUpstreamReachableWithOffset(3, envoyPort, tu, nil)
	}

	for _, testCase := range testCases {

		Describe(fmt.Sprintf("%s: (%s)", testCase.Title, testCase.TransportApiVersion.String()), func() {

			Describe("kubernetes happy path", func() {

				var (
					namespace      string
					writeNamespace string
					cfg            *rest.Config
					kubeClient     kubernetes.Interface
					svc            *kubev1.Service
				)

				BeforeEach(func() {
					namespace = ""
					writeNamespace = ""
					var err error
					svc = nil
					cfg, err = kubeutils.GetConfig("", "")
					Expect(err).NotTo(HaveOccurred())
					kubeClient, err = kubernetes.NewForConfig(cfg)
					Expect(err).NotTo(HaveOccurred())
				})

				AfterEach(func() {
					if namespace != "" {
						_ = setup.TeardownKube(namespace)
					}
				})

				prepNamespace := func() {
					if namespace == "" {
						namespace = "gloo-e2e-" + helpers.RandString(8)
					}

					_, err := kubeClient.CoreV1().Namespaces().Create(ctx, &kubev1.Namespace{
						ObjectMeta: metav1.ObjectMeta{
							Name: namespace,
						},
					}, metav1.CreateOptions{})
					Expect(err).NotTo(HaveOccurred())

					svc, err = kubeClient.CoreV1().Services(namespace).Create(ctx, &kubev1.Service{
						ObjectMeta: metav1.ObjectMeta{
							Namespace: namespace,
							Name:      "headlessservice",
						},
						Spec: kubev1.ServiceSpec{
							Ports: []kubev1.ServicePort{
								{
									Name: "foo",
									Port: int32(tu.Port),
								},
							},
						},
					}, metav1.CreateOptions{})
					Expect(err).NotTo(HaveOccurred())

					_, err = kubeClient.CoreV1().Endpoints(namespace).Create(ctx, &kubev1.Endpoints{
						ObjectMeta: metav1.ObjectMeta{
							Namespace: namespace,
							Name:      svc.Name,
						},
						Subsets: []kubev1.EndpointSubset{{
							Addresses: []kubev1.EndpointAddress{{
								IP:       getNonSpecialIP(envoyInstance),
								Hostname: "localhost",
							}},
							Ports: []kubev1.EndpointPort{{
								Port: int32(tu.Port),
							}},
						}},
					}, metav1.CreateOptions{})
					Expect(err).NotTo(HaveOccurred())
				}

				getUpstream := func() (*gloov1.Upstream, error) {
					l, err := testClients.UpstreamClient.List(writeNamespace, clients.ListOpts{})
					if err != nil {
						return nil, err
					}
					for _, u := range l {
						if strings.Contains(u.Metadata.Name, svc.Name) && strings.Contains(u.Metadata.Name, svc.Namespace) {
							return u, nil
						}
					}
					return nil, fmt.Errorf("not found")
				}

				Context("specific namespace", func() {

					BeforeEach(func() {
						prepNamespace()
						writeNamespace = namespace
						ro := &services.RunOptions{
							NsToWrite: writeNamespace,
							NsToWatch: []string{namespace},
							WhatToRun: services.What{
								DisableGateway: true,
							},
							KubeClient: kubeClient,
							Settings: &gloov1.Settings{
								Gloo: &gloov1.GlooOptions{
									EnableRestEds: testCase.RestEdsEnabled,
								},
							},
						}

						testClients = services.RunGlooGatewayUdsFds(ctx, ro)
						role := namespace + "~" + gatewaydefaults.GatewayProxyName
						err := envoyInstance.RunWithRole(role, testClients.GlooPort)
						Expect(err).NotTo(HaveOccurred())

						testhelpers.EventuallyResourceAccepted(func() (resources.InputResource, error) {
							return getUpstream()
						}, "20s", ".5s")
					})

					It("should discover service", func() {
						up, err := getUpstream()
						Expect(err).NotTo(HaveOccurred())

						proxycli := testClients.ProxyClient
						proxy := getTrivialProxyForUpstream(namespace, envoyPort, up.Metadata.Ref())
						var opts clients.WriteOpts
						_, err = proxycli.Write(proxy, opts)
						Expect(err).NotTo(HaveOccurred())

						TestUpstreamReachable()
					})

					It("should create appropriate config for discovered service", func() {
						up, err := getUpstream()
						Expect(err).NotTo(HaveOccurred())

						svc.Annotations = map[string]string{
							"gloo.solo.io/upstream_config": "{\"initial_stream_window_size\": 2048}",
						}
						svc, err = kubeClient.CoreV1().Services(namespace).Update(ctx, svc, metav1.UpdateOptions{})
						Expect(err).NotTo(HaveOccurred())

						proxycli := testClients.ProxyClient
						proxy := getTrivialProxyForUpstream(namespace, envoyPort, up.Metadata.Ref())
						var opts clients.WriteOpts
						_, err = proxycli.Write(proxy, opts)
						Expect(err).NotTo(HaveOccurred())

						TestUpstreamReachable()
						upstream, err := getUpstream()
						Expect(err).NotTo(HaveOccurred())
						Expect(int(upstream.GetInitialStreamWindowSize().GetValue())).To(Equal(2048))
					})

					It("correctly routes requests to a service destination", func() {
						svcRef := skkubeutils.FromKubeMeta(svc.ObjectMeta, true).Ref()
						svcPort := svc.Spec.Ports[0].Port
						proxy := getTrivialProxyForService(namespace, envoyPort, svcRef, uint32(svcPort))

						_, err := testClients.ProxyClient.Write(proxy, clients.WriteOpts{})
						Expect(err).NotTo(HaveOccurred())

						TestUpstreamReachable()
					})
				})

				Context("all namespaces", func() {
					BeforeEach(func() {
						namespace = "gloo-e2e-" + helpers.RandString(8)
						prepNamespace()

						writeNamespace = namespace
						ro := &services.RunOptions{
							NsToWrite: writeNamespace,
							NsToWatch: []string{},
							WhatToRun: services.What{
								DisableGateway: true,
							},
							KubeClient: kubeClient,
							Settings: &gloov1.Settings{
								Gloo: &gloov1.GlooOptions{
									EnableRestEds: testCase.RestEdsEnabled,
								},
							},
						}

						testClients = services.RunGlooGatewayUdsFds(ctx, ro)
						role := namespace + "~" + gatewaydefaults.GatewayProxyName
						err := envoyInstance.RunWithRole(role, testClients.GlooPort)
						Expect(err).NotTo(HaveOccurred())

					})

					It("watch all namespaces", func() {
						testhelpers.EventuallyResourceAccepted(func() (resources.InputResource, error) {
							return getUpstream()
						})

						up, err := getUpstream()
						Expect(err).NotTo(HaveOccurred())

						proxycli := testClients.ProxyClient
						proxy := getTrivialProxyForUpstream(namespace, envoyPort, up.Metadata.Ref())
						var opts clients.WriteOpts
						_, err = proxycli.Write(proxy, opts)
						Expect(err).NotTo(HaveOccurred())

						TestUpstreamReachable()
					})
				})
			})
		})

	}

})

func getTrivialProxyForUpstream(ns string, bindPort uint32, upstream *core.ResourceRef) *gloov1.Proxy {
	proxy := getTrivialProxy(ns, bindPort)
	proxy.Listeners[0].ListenerType.(*gloov1.Listener_HttpListener).HttpListener.
		VirtualHosts[0].Routes[0].Action.(*gloov1.Route_RouteAction).RouteAction.
		Destination.(*gloov1.RouteAction_Single).Single.DestinationType =
		&gloov1.Destination_Upstream{Upstream: upstream}
	return proxy
}

func getTrivialProxyForService(ns string, bindPort uint32, service *core.ResourceRef, svcPort uint32) *gloov1.Proxy {
	proxy := getTrivialProxy(ns, bindPort)
	proxy.Listeners[0].ListenerType.(*gloov1.Listener_HttpListener).HttpListener.
		VirtualHosts[0].Routes[0].Action.(*gloov1.Route_RouteAction).RouteAction.
		Destination.(*gloov1.RouteAction_Single).Single.DestinationType =
		&gloov1.Destination_Kube{
			Kube: &gloov1.KubernetesServiceDestination{
				Ref:  service,
				Port: svcPort,
			},
		}
	return proxy
}

func getTrivialProxy(ns string, bindPort uint32) *gloov1.Proxy {
	return &gloov1.Proxy{
		Metadata: &core.Metadata{
			Name:      gatewaydefaults.GatewayProxyName,
			Namespace: ns,
		},
		Listeners: []*gloov1.Listener{{
			Name:        "listener",
			BindAddress: "::",
			BindPort:    bindPort,
			ListenerType: &gloov1.Listener_HttpListener{
				HttpListener: &gloov1.HttpListener{
					VirtualHosts: []*gloov1.VirtualHost{{
						Name:    "virt1",
						Domains: []string{"*"},
						Routes: []*gloov1.Route{{
							Action: &gloov1.Route_RouteAction{
								RouteAction: &gloov1.RouteAction{
									Destination: &gloov1.RouteAction_Single{
										Single: &gloov1.Destination{},
									},
								},
							},
						}},
					}},
				},
			},
		}},
	}
}

// getNonSpecialIP returns a non-special IP that Kubernetes will allow in an endpoint.
func getNonSpecialIP(instance *services.EnvoyInstance) string {
	if instance.UseDocker {
		return instance.LocalAddr()
	}

	ifaces, err := net.Interfaces()
	ExpectWithOffset(1, err).NotTo(HaveOccurred())

	for _, iface := range ifaces {
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			default:
				continue
			}
			if isNonSpecialIP(ip) {
				return ip.String()
			}
		}
	}
	Fail("no ip address available", 1)
	return ""
}

// isNonSpecialIP is adapted from ValidateNonSpecialIP in k8s.io/kubernetes/pkg/apis/core/validation/validation.go
//
// Specifically disallowed are unspecified, loopback addresses, and link-local addresses
// which tend to be used for node-centric purposes (e.g. metadata service).
func isNonSpecialIP(ip net.IP) bool {
	if ip == nil {
		return false // must be a valid IP address
	}
	if ip.IsUnspecified() {
		return false // may not be unspecified
	}
	if ip.IsLoopback() {
		return false // may not be in the loopback range (127.0.0.0/8, ::1/128)
	}
	if ip.IsLinkLocalUnicast() {
		return false // may not be in the link-local range (169.254.0.0/16, fe80::/10)
	}
	if ip.IsLinkLocalMulticast() {
		return false // may not be in the link-local multicast range (224.0.0.0/24, ff02::/10)
	}
	return true
}
