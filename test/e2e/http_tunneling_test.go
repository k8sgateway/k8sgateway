package e2e_test

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"time"

	static_plugin_gloo "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/static"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

	gatewayv1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	gatewaydefaults "github.com/solo-io/gloo/projects/gateway/pkg/defaults"
	gloohelpers "github.com/solo-io/gloo/test/helpers"

	"github.com/elazarl/goproxy"
	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/test/services"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

var _ = FDescribe("tunneling", func() {

	var (
		ctx           context.Context
		cancel        context.CancelFunc
		testClients   services.TestClients
		envoyInstance *services.EnvoyInstance
		up            *gloov1.Upstream

		writeNamespace = defaults.GlooSystem
	)

	BeforeEach(func() {
		var err error
		ctx, cancel = context.WithCancel(context.Background())
		defaults.HttpPort = services.NextBindPort()

		// run gloo
		ro := &services.RunOptions{
			NsToWrite: writeNamespace,
			NsToWatch: []string{"default", writeNamespace},
			WhatToRun: services.What{
				DisableFds: true,
				DisableUds: true,
			},
		}
		testClients = services.RunGlooGatewayUdsFds(ctx, ro)

		// write gateways and wait for them to be created
		err = gloohelpers.WriteDefaultGateways(writeNamespace, testClients.GatewayClient)
		Expect(err).NotTo(HaveOccurred(), "Should be able to write default gateways")
		Eventually(func() (gatewayv1.GatewayList, error) {
			return testClients.GatewayClient.List(writeNamespace, clients.ListOpts{})
		}, "10s", "0.1s").Should(HaveLen(2), "Gateways should be present")

		// run envoy
		envoyInstance, err = envoyFactory.NewEnvoyInstance()
		Expect(err).NotTo(HaveOccurred())
		err = envoyInstance.RunWithRoleAndRestXds(writeNamespace+"~"+gatewaydefaults.GatewayProxyName, testClients.GlooPort, testClients.RestXdsPort)
		Expect(err).NotTo(HaveOccurred())

		// write a test upstream
		// this is the upstream that will handle requests
		proxy := goproxy.NewProxyHttpServer()
		proxy.Verbose = true
		err = setCA([]byte(gloohelpers.Certificate()), []byte(gloohelpers.PrivateKey()))
		Expect(err).ToNot(HaveOccurred())
		proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)

		listener, err := net.Listen("tcp", ":0")
		if err != nil {
			panic(err)
		}

		addr := listener.Addr().String()
		_, portStr, err := net.SplitHostPort(addr)
		if err != nil {
			panic(err)
		}

		port, err := strconv.Atoi(portStr)
		if err != nil {
			panic(err)
		}

		go func() {
			defer GinkgoRecover()
			server := &http.Server{Addr: addr, Handler: proxy}
			server.Serve(listener)

		}()

		secret := &gloov1.Secret{
			Metadata: &core.Metadata{
				Name:      "secret",
				Namespace: "default",
			},
			Kind: &gloov1.Secret_Tls{
				Tls: &gloov1.TlsSecret{
					CertChain:  gloohelpers.Certificate(),
					PrivateKey: gloohelpers.PrivateKey(),
				},
			},
		}

		_, err = testClients.SecretClient.Write(secret, clients.WriteOpts{OverwriteExisting: true})
		Expect(err).NotTo(HaveOccurred())

		up = &gloov1.Upstream{
			Metadata: &core.Metadata{
				Name:      "local-1",
				Namespace: "default",
			},
			UpstreamType: &gloov1.Upstream_Static{
				Static: &static_plugin_gloo.UpstreamSpec{
					Hosts: []*static_plugin_gloo.Host{
						{
							Addr: envoyInstance.LocalAddr(),
							Port: uint32(port),
						},
					},
				},
			},
		}

		up.HttpProxyHostname = "host.com:443" // enable HTTP tunneling
		up.SslConfig = &gloov1.UpstreamSslConfig{
			SslSecrets: &gloov1.UpstreamSslConfig_SecretRef{
				SecretRef: &core.ResourceRef{Name: "secret", Namespace: "default"},
			},
		}
		_, err = testClients.UpstreamClient.Write(up, clients.WriteOpts{OverwriteExisting: true})
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		if envoyInstance != nil {
			_ = envoyInstance.Clean()
		}
		cancel()
	})

	checkProxy := func() {
		// ensure the proxy is created
		Eventually(func() (*gloov1.Proxy, error) {
			return testClients.ProxyClient.Read(writeNamespace, gatewaydefaults.GatewayProxyName, clients.ReadOpts{})
		}, "5s", "0.1s").ShouldNot(BeNil())
	}

	checkVirtualService := func(testVs *gatewayv1.VirtualService) {
		Eventually(func() (*gatewayv1.VirtualService, error) {
			return testClients.VirtualServiceClient.Read(testVs.Metadata.GetNamespace(), testVs.Metadata.GetName(), clients.ReadOpts{})
		}, "5s", "0.1s").ShouldNot(BeNil())
	}

	testRequest := func() string {
		By("Make request")
		responseBody := ""
		EventuallyWithOffset(1, func() error {

			var client http.Client

			scheme := "http"

			//cert := gloohelpers.Certificate()
			//rootca := &cert
			//if rootca != nil {
			//	scheme = "https"
			//	caCertPool := x509.NewCertPool()
			//	ok := caCertPool.AppendCertsFromPEM([]byte(*rootca))
			//	if !ok {
			//		return fmt.Errorf("ca cert is not OK")
			//	}
			//
			//	client.Transport = &http.Transport{
			//		TLSClientConfig: &tls.Config{
			//			RootCAs:            caCertPool,
			//			InsecureSkipVerify: true,
			//		},
			//	}
			//}

			req, err := http.NewRequest("GET", fmt.Sprintf("%s://%s:%d/test", scheme, "localhost", defaults.HttpPort), nil)
			if err != nil {
				return err
			}
			res, err := client.Do(req)
			if err != nil {
				return err
			}
			p := new(bytes.Buffer)
			if _, err := io.Copy(p, res.Body); err != nil {
				return err
			}
			defer res.Body.Close()
			responseBody = p.String()
			return nil
		}, "10s", ".1s").Should(BeNil())
		return responseBody
	}

	Context("with envoy", func() {

		JustBeforeEach(func() {
			// write a virtual service so we have a proxy to our test upstream
			testVs := getTrivialVirtualServiceForUpstream(writeNamespace, up.Metadata.Ref())
			_, err := testClients.VirtualServiceClient.Write(testVs, clients.WriteOpts{})
			Expect(err).NotTo(HaveOccurred())

			checkProxy()
			checkVirtualService(testVs)
		})

		It("should proxy http", func() {
			time.Sleep(1 * time.Second) //TODO(kdorosh) remove
			testReq := testRequest()
			Expect(testReq).Should(ContainSubstring("400 The plain HTTP request was sent to HTTPS port"))
		})

		Context("with SSL", func() {
			It("should proxy HTTPS", func() {
				time.Sleep(1 * time.Second) //TODO(kdorosh) remove
				testReq := testRequest()
				Expect(testReq).Should(ContainSubstring("403 Forbidden")) //TODO(kdorosh) change to real site
			})
		})
	})

})

func setCA(caCert, caKey []byte) error {
	goproxyCa, err := tls.X509KeyPair(caCert, caKey)
	if err != nil {
		return err
	}
	if goproxyCa.Leaf, err = x509.ParseCertificate(goproxyCa.Certificate[0]); err != nil {
		return err
	}
	goproxy.GoproxyCa = goproxyCa
	goproxy.OkConnect = &goproxy.ConnectAction{Action: goproxy.ConnectAccept, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	goproxy.MitmConnect = &goproxy.ConnectAction{Action: goproxy.ConnectMitm, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	goproxy.HTTPMitmConnect = &goproxy.ConnectAction{Action: goproxy.ConnectHTTPMitm, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	goproxy.RejectConnect = &goproxy.ConnectAction{Action: goproxy.ConnectReject, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	return nil
}
