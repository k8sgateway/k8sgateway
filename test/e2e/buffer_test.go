package e2e_test

import (
	"bytes"
	"fmt"
	"github.com/solo-io/gloo/test/e2e"
	"github.com/solo-io/gloo/test/helpers"
	matchers2 "github.com/solo-io/gloo/test/matchers"
	"net/http"
	"time"

	"github.com/golang/protobuf/ptypes/wrappers"

	gatewayv1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	gatewaydefaults "github.com/solo-io/gloo/projects/gateway/pkg/defaults"
	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	buffer "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/filters/http/buffer/v3"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
)

var _ = Describe("buffer", func() {

	var (
		testContext *e2e.TestContext
	)

	BeforeEach(func() {
		testContext = testContextFactory.NewTestContext()
		testContext.BeforeEach()
	})

	AfterEach(func() {
		testContext.AfterEach()
	})

	JustBeforeEach(func() {
		testContext.JustBeforeEach()
	})

	JustBeforeEach(func() {
		testContext.JustAfterEach()
	})

	testRequest := func() func() (*http.Response, error) {
		return func() (*http.Response, error) {
			var json = []byte(`{"value":"test"}`)
			req, err := http.NewRequest("POST", fmt.Sprintf("http://%s:%d/test", "localhost", defaults.HttpPort), bytes.NewBuffer(json))
			//req.Header.Add("Content-Length",size)
			req.Header.Set("Content-Type", "application/json")
			req.Host = "test.com"
			if err != nil {
				return nil, err
			}

			return http.DefaultClient.Do(req)
		}
	}

	Context("filter defined on listener", func() {

		Context("Large buffer ", func() {

			BeforeEach(func() {
				gw := gatewaydefaults.DefaultGateway(writeNamespace)
				gw.GetHttpGateway().Options = &gloov1.HttpListenerOptions{
					Buffer: &buffer.Buffer{
						MaxRequestBytes: &wrappers.UInt32Value{
							Value: 4098, // max size
						},
					},
				}

				testContext.ResourcesToCreate().Gateways = gatewayv1.GatewayList{
					gw,
				}
			})

			It("valid buffer size should succeed", func() {
				Eventually(testRequest(), 10*time.Second, 1*time.Second).Should(matchers2.MatchHttpResponse(&matchers2.HttpResponse{
					Body: "{\"value\":\"test\"}",
					Response: &http.Response{
						StatusCode: http.StatusOK,
					},
				}))

			})

		})

		Context("Small buffer ", func() {

			BeforeEach(func() {
				gw := gatewaydefaults.DefaultGateway(writeNamespace)
				gw.GetHttpGateway().Options = &gloov1.HttpListenerOptions{
					Buffer: &buffer.Buffer{
						MaxRequestBytes: &wrappers.UInt32Value{
							Value: 1,
						},
					},
				}

				testContext.ResourcesToCreate().Gateways = gatewayv1.GatewayList{
					gw,
				}
			})

			It("empty buffer should fail", func() {
				Eventually(testRequest(), 10*time.Second, 1*time.Second).Should(matchers2.MatchHttpResponse(&matchers2.HttpResponse{
					Body: "Payload Too Large",
					Response: &http.Response{
						StatusCode: http.StatusRequestEntityTooLarge,
					},
				}))
			})
		})
	})

	Context("filter defined on listener and vhost", func() {

		Context("Large buffer ", func() {
			BeforeEach(func() {
				gw := gatewaydefaults.DefaultGateway(writeNamespace)
				gw.GetHttpGateway().Options = &gloov1.HttpListenerOptions{
					Buffer: &buffer.Buffer{
						MaxRequestBytes: &wrappers.UInt32Value{
							Value: 1,
						},
					},
				}
				vsToTestUpstream := helpers.NewVirtualServiceBuilder().
					WithName("vs-test").
					WithNamespace(writeNamespace).
					WithDomain("test.com").
					WithVirtualHostOptions(&gloov1.VirtualHostOptions{
						BufferPerRoute: &buffer.BufferPerRoute{
							Override: &buffer.BufferPerRoute_Buffer{
								Buffer: &buffer.Buffer{
									MaxRequestBytes: &wrappers.UInt32Value{
										Value: 4098, // max size
									},
								},
							},
						},
					}).
					WithRoutePrefixMatcher("test", "/").
					WithRouteActionToUpstream("test", testContext.TestUpstream().Upstream).
					Build()

				testContext.ResourcesToCreate().Gateways = gatewayv1.GatewayList{
					gw,
				}
				testContext.ResourcesToCreate().VirtualServices = gatewayv1.VirtualServiceList{
					vsToTestUpstream,
				}
			})

			It("valid buffer size should succeed", func() {
				Eventually(testRequest(), 10*time.Second, 1*time.Second).Should(matchers2.MatchHttpResponse(&matchers2.HttpResponse{
					Body: "{\"value\":\"test\"}",
					Response: &http.Response{
						StatusCode: http.StatusOK,
					},
				}))
			})

		})

		Context("Small buffer ", func() {

			BeforeEach(func() {
				gw := gatewaydefaults.DefaultGateway(writeNamespace)
				gw.GetHttpGateway().Options = &gloov1.HttpListenerOptions{
					Buffer: &buffer.Buffer{
						MaxRequestBytes: &wrappers.UInt32Value{
							Value: 4098,
						},
					},
				}
				vsToTestUpstream := helpers.NewVirtualServiceBuilder().
					WithName("vs-test").
					WithNamespace(writeNamespace).
					WithDomain("test.com").
					WithVirtualHostOptions(&gloov1.VirtualHostOptions{
						BufferPerRoute: &buffer.BufferPerRoute{
							Override: &buffer.BufferPerRoute_Buffer{
								Buffer: &buffer.Buffer{
									MaxRequestBytes: &wrappers.UInt32Value{
										Value: 1,
									},
								},
							},
						},
					}).
					WithRoutePrefixMatcher("test", "/").
					WithRouteActionToUpstream("test", testContext.TestUpstream().Upstream).
					Build()

				testContext.ResourcesToCreate().Gateways = gatewayv1.GatewayList{
					gw,
				}
				testContext.ResourcesToCreate().VirtualServices = gatewayv1.VirtualServiceList{
					vsToTestUpstream,
				}
			})

			It("empty buffer should fail", func() {
				Eventually(testRequest(), 10*time.Second, 1*time.Second).Should(matchers2.MatchHttpResponse(&matchers2.HttpResponse{
					Body: "Payload Too Large",
					Response: &http.Response{
						StatusCode: http.StatusRequestEntityTooLarge,
					},
				}))
			})
		})
	})

	Context("filter defined on listener and route", func() {

		Context("Large buffer ", func() {

			BeforeEach(func() {
				gw := gatewaydefaults.DefaultGateway(writeNamespace)
				gw.GetHttpGateway().Options = &gloov1.HttpListenerOptions{
					Buffer: &buffer.Buffer{
						MaxRequestBytes: &wrappers.UInt32Value{
							Value: 1,
						},
					},
				}
				vsToTestUpstream := helpers.NewVirtualServiceBuilder().
					WithName("vs-test").
					WithNamespace(writeNamespace).
					WithDomain("test.com").
					WithRoutePrefixMatcher("test", "/").
					WithRouteActionToUpstream("test", testContext.TestUpstream().Upstream).
					WithRouteOptions("test", &gloov1.RouteOptions{
						BufferPerRoute: &buffer.BufferPerRoute{
							Override: &buffer.BufferPerRoute_Buffer{
								Buffer: &buffer.Buffer{
									MaxRequestBytes: &wrappers.UInt32Value{
										Value: 4098, // max size
									},
								},
							},
						},
					}).
					Build()

				testContext.ResourcesToCreate().Gateways = gatewayv1.GatewayList{
					gw,
				}
				testContext.ResourcesToCreate().VirtualServices = gatewayv1.VirtualServiceList{
					vsToTestUpstream,
				}
			})

			It("valid buffer size should succeed", func() {
				Eventually(testRequest(), 10*time.Second, 1*time.Second).Should(matchers2.MatchHttpResponse(&matchers2.HttpResponse{
					Body: "{\"value\":\"test\"}",
					Response: &http.Response{
						StatusCode: http.StatusOK,
					},
				}))
			})

		})

		Context("Small buffer ", func() {

			BeforeEach(func() {
				gw := gatewaydefaults.DefaultGateway(writeNamespace)
				gw.GetHttpGateway().Options = &gloov1.HttpListenerOptions{
					Buffer: &buffer.Buffer{
						MaxRequestBytes: &wrappers.UInt32Value{
							Value: 4098,
						},
					},
				}
				vsToTestUpstream := helpers.NewVirtualServiceBuilder().
					WithName("vs-test").
					WithNamespace(writeNamespace).
					WithDomain("test.com").
					WithRoutePrefixMatcher("test", "/").
					WithRouteActionToUpstream("test", testContext.TestUpstream().Upstream).
					WithRouteOptions("test", &gloov1.RouteOptions{
						BufferPerRoute: &buffer.BufferPerRoute{
							Override: &buffer.BufferPerRoute_Buffer{
								Buffer: &buffer.Buffer{
									MaxRequestBytes: &wrappers.UInt32Value{
										Value: 1,
									},
								},
							},
						},
					}).
					Build()

				testContext.ResourcesToCreate().Gateways = gatewayv1.GatewayList{
					gw,
				}
				testContext.ResourcesToCreate().VirtualServices = gatewayv1.VirtualServiceList{
					vsToTestUpstream,
				}
			})

			It("empty buffer should fail", func() {
				Eventually(testRequest(), 10*time.Second, 1*time.Second).Should(matchers2.MatchHttpResponse(&matchers2.HttpResponse{
					Body: "Payload Too Large",
					Response: &http.Response{
						StatusCode: http.StatusRequestEntityTooLarge,
					},
				}))
			})
		})
	})

})
