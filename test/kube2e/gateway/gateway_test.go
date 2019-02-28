package gateway_test

import (
	"context"
	"fmt"
	"github.com/solo-io/gloo/test/kube2e"
	"io/ioutil"
	"k8s.io/client-go/kubernetes"
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gateway/pkg/defaults"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/test/helpers"
	"github.com/solo-io/go-utils/kubeutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/test/setup"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"k8s.io/client-go/rest"
)

var _ = Describe("Kube2e: gateway", func() {

	var (
		ctx        context.Context
		cancel     context.CancelFunc
		cfg        *rest.Config
		kubeClient kubernetes.Interface

		gatewayClient        v1.GatewayClient
		virtualServiceClient v1.VirtualServiceClient
	)

	BeforeEach(func() {
		ctx, cancel = context.WithCancel(context.Background())

		var err error
		cfg, err = kubeutils.GetConfig("", "")
		Expect(err).NotTo(HaveOccurred())

		kubeClient, err = kubernetes.NewForConfig(cfg)
		Expect(err).NotTo(HaveOccurred())

		cache := kube.NewKubeCache(ctx)
		gatewayClientFactory := &factory.KubeResourceClientFactory{
			Crd:         v1.GatewayCrd,
			Cfg:         cfg,
			SharedCache: cache,
		}
		virtualServiceClientFactory := &factory.KubeResourceClientFactory{
			Crd:         v1.VirtualServiceCrd,
			Cfg:         cfg,
			SharedCache: cache,
		}

		gatewayClient, err = v1.NewGatewayClient(gatewayClientFactory)
		Expect(err).NotTo(HaveOccurred())

		virtualServiceClient, err = v1.NewVirtualServiceClient(virtualServiceClientFactory)
		Expect(err).NotTo(HaveOccurred())

	})

	AfterEach(func() {
		cancel()
		err := virtualServiceClient.Delete(namespace, "vs", clients.DeleteOpts{})
		Expect(err).NotTo(HaveOccurred())
	})

	It("works", func() {

		_, err := virtualServiceClient.Write(&v1.VirtualService{

			Metadata: core.Metadata{
				Name:      "vs",
				Namespace: namespace,
			},
			VirtualHost: &gloov1.VirtualHost{
				Name:    "default",
				Domains: []string{"*"},
				Routes: []*gloov1.Route{{
					Matcher: &gloov1.Matcher{
						PathSpecifier: &gloov1.Matcher_Prefix{
							Prefix: "/",
						},
					},
					Action: &gloov1.Route_RouteAction{
						RouteAction: &gloov1.RouteAction{
							Destination: &gloov1.RouteAction_Single{
								Single: &gloov1.Destination{
									Upstream: core.ResourceRef{Namespace: namespace, Name: fmt.Sprintf("%s-%s-%v", namespace, "testrunner", kube2e.TestRunnerPort)},
								},
							},
						},
					},
				}},
			},
		}, clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())

		fmt.Println("wait for gateway to be created")

		defaultGateway := defaults.DefaultGateway(namespace)
		// wait for default gateway to be created
		Eventually(func() (*v1.Gateway, error) {
			return gatewayClient.Read(namespace, defaultGateway.Metadata.Name, clients.ReadOpts{})
		}, "5s", "0.5s").Should(Not(BeNil()))

		gatewayProxy := "gateway-proxy"
		gatewayPort := int(80)
		setup.CurlEventuallyShouldRespond(setup.CurlOpts{
			Protocol: "http",
			Path:     "/",
			Method:   "GET",
			Host:     gatewayProxy,
			Service:  gatewayProxy,
			Port:     gatewayPort,
		}, namespace, kube2e.SimpleHttpResponse, 1, time.Minute)
	})

	Context("native ssl ", func() {

		BeforeEach(func() {
			// get the certificate so it is generated in the background
			go helpers.Certificate()
		})

		AfterEach(func() {
			err := kubeClient.CoreV1().Secrets(namespace).Delete("secret", nil)
			Expect(err).NotTo(HaveOccurred())
		})

		It("works with ssl", func() {
			createdSecret, err := kubeClient.CoreV1().Secrets(namespace).Create(helpers.GetKubeSecret("secret", namespace))
			Expect(err).NotTo(HaveOccurred())

			_, err = virtualServiceClient.Write(&v1.VirtualService{

				Metadata: core.Metadata{
					Name:      "vs",
					Namespace: namespace,
				},
				SslConfig: &gloov1.SslConfig{
					SslSecrets: &gloov1.SslConfig_SecretRef{
						SecretRef: &core.ResourceRef{
							Name:      createdSecret.ObjectMeta.Name,
							Namespace: createdSecret.ObjectMeta.Namespace,
						},
					},
				},
				VirtualHost: &gloov1.VirtualHost{
					Name:    "default",
					Domains: []string{"*"},
					Routes: []*gloov1.Route{{
						Matcher: &gloov1.Matcher{
							PathSpecifier: &gloov1.Matcher_Prefix{
								Prefix: "/",
							},
						},
						Action: &gloov1.Route_RouteAction{
							RouteAction: &gloov1.RouteAction{
								Destination: &gloov1.RouteAction_Single{
									Single: &gloov1.Destination{
										Upstream: core.ResourceRef{Namespace: namespace, Name: fmt.Sprintf("%s-%s-%v", namespace, "testrunner", kube2e.TestRunnerPort)},
									},
								},
							},
						},
					}},
				},
			}, clients.WriteOpts{})
			Expect(err).NotTo(HaveOccurred())

			defaultGateway := defaults.DefaultGateway(namespace)
			// wait for default gateway to be created
			Eventually(func() (*v1.Gateway, error) {
				return gatewayClient.Read(namespace, defaultGateway.Metadata.Name, clients.ReadOpts{})
			}, "5s", "0.5s").Should(Not(BeNil()))

			gatewayProxy := "gateway-proxy"
			gatewayPort := int(443)
			caFile := ToFile(helpers.Certificate())
			//noinspection GoUnhandledErrorResult
			defer os.Remove(caFile)

			err = setup.Kubectl("cp", caFile, namespace+"/testrunner:/tmp/ca.crt")
			Expect(err).NotTo(HaveOccurred())

			setup.CurlEventuallyShouldRespond(setup.CurlOpts{
				Protocol: "https",
				Path:     "/",
				Method:   "GET",
				Host:     gatewayProxy,
				Service:  gatewayProxy,
				Port:     gatewayPort,
				CaFile:   "/tmp/ca.crt",
			}, namespace, kube2e.SimpleHttpResponse, 1, time.Minute)
		})
	})
})

func ToFile(content string) string {
	f, err := ioutil.TempFile("", "")
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	n, err := f.WriteString(content)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	ExpectWithOffset(1, n).To(Equal(len(content)))
	_ = f.Close()
	return f.Name()
}
