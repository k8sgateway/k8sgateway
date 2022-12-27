package e2e_test

import (
	"bytes"
	"encoding/base64"
	"fmt"

	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	"github.com/solo-io/gloo/test/helpers"

	"net/http"

	"github.com/solo-io/gloo/test/e2e"
	testmatchers "github.com/solo-io/gloo/test/matchers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	envoytransformation "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/transformation"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/core/matchers"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/transformation"
	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
)

var _ = Describe("Staged Transformation", func() {

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

	JustAfterEach(func() {
		testContext.JustAfterEach()
	})

	newRequestWithBody := func(body string) *http.Request {
		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost:%d/", defaults.HttpPort), bytes.NewBufferString(body))
		ExpectWithOffset(1, err).NotTo(HaveOccurred())
		req.Header.Set("Content-Type", "application/octet-stream")
		req.Host = e2e.DefaultHost

		return req
	}

	Context("no auth", func() {

		It("should transform response", func() {
			testContext.PatchDefaultVirtualService(func(vs *v1.VirtualService) *v1.VirtualService {
				vsBuilder := helpers.BuilderFromVirtualService(vs)
				vsBuilder.WithVirtualHostOptions(&gloov1.VirtualHostOptions{
					StagedTransformations: &transformation.TransformationStages{
						Early: &transformation.RequestResponseTransformations{
							ResponseTransforms: []*transformation.ResponseMatch{{
								Matchers: []*matchers.HeaderMatcher{
									{
										Name:  ":status",
										Value: "200",
									},
								},
								ResponseTransformation: &transformation.Transformation{
									TransformationType: &transformation.Transformation_TransformationTemplate{
										TransformationTemplate: &envoytransformation.TransformationTemplate{
											ParseBodyBehavior: envoytransformation.TransformationTemplate_DontParse,
											BodyTransformation: &envoytransformation.TransformationTemplate_Body{
												Body: &envoytransformation.InjaTemplate{
													Text: "early-transformed",
												},
											},
										},
									},
								},
							}},
						},
						// add regular response to see that the early one overrides it
						Regular: &transformation.RequestResponseTransformations{
							ResponseTransforms: []*transformation.ResponseMatch{{
								Matchers: []*matchers.HeaderMatcher{
									{
										Name:  ":status",
										Value: "200",
									},
								},
								ResponseTransformation: &transformation.Transformation{
									TransformationType: &transformation.Transformation_TransformationTemplate{
										TransformationTemplate: &envoytransformation.TransformationTemplate{
											ParseBodyBehavior: envoytransformation.TransformationTemplate_DontParse,
											BodyTransformation: &envoytransformation.TransformationTemplate_Body{
												Body: &envoytransformation.InjaTemplate{
													Text: "regular-transformed",
												},
											},
										},
									},
								},
							}},
						},
					},
				})
				return vsBuilder.Build()
			})

			req := newRequestWithBody("test")
			Eventually(func(g Gomega) {
				res, err := http.DefaultClient.Do(req)
				g.Expect(err).NotTo(HaveOccurred())
				g.Expect(res).Should(testmatchers.HaveHttpResponse(&testmatchers.HttpResponse{
					StatusCode: http.StatusOK,
					Body:       "early-transformed",
				}))
			}, "15s", ".5s").Should(Succeed())
		})

		It("should allow multiple header values for the same header when using HeadersToAppend", func() {
			testContext.PatchDefaultVirtualService(func(vs *v1.VirtualService) *v1.VirtualService {
				vsBuilder := helpers.BuilderFromVirtualService(vs)
				vsBuilder.WithVirtualHostOptions(&gloov1.VirtualHostOptions{
					StagedTransformations: &transformation.TransformationStages{
						Regular: &transformation.RequestResponseTransformations{
							ResponseTransforms: []*transformation.ResponseMatch{{
								ResponseTransformation: &transformation.Transformation{
									TransformationType: &transformation.Transformation_TransformationTemplate{
										TransformationTemplate: &envoytransformation.TransformationTemplate{
											ParseBodyBehavior: envoytransformation.TransformationTemplate_DontParse,
											Headers: map[string]*envoytransformation.InjaTemplate{
												"x-custom-header": {Text: "original header"},
											},
											HeadersToAppend: []*envoytransformation.TransformationTemplate_HeaderToAppend{
												{
													Key:   "x-custom-header",
													Value: &envoytransformation.InjaTemplate{Text: "{{upper(\"appended header 1\")}}"},
												},
												{
													Key:   "x-custom-header",
													Value: &envoytransformation.InjaTemplate{Text: "{{upper(\"appended header 2\")}}"},
												},
											},
										},
									},
								},
							}},
						},
					},
				})
				return vsBuilder.Build()
			})

			req := newRequestWithBody("")
			Eventually(func(g Gomega) {
				res, err := http.DefaultClient.Do(req)
				g.Expect(err).NotTo(HaveOccurred())
				g.Expect(res).Should(testmatchers.HaveHttpResponse(&testmatchers.HttpResponse{
					StatusCode: http.StatusOK,
					Body:       BeEmpty(),
					// The default Header matcher only works with single headers, so we supply a custom matcher in this case
					Custom: WithTransform(testmatchers.WithHeaderValues("X-Custom-Header"), ContainElements("original header", "APPENDED HEADER 1", "APPENDED HEADER 2")),
				}))
			}, "15s", ".5s").Should(Succeed())
		})

		It("Should be able to base64 encode the body", func() {
			testContext.PatchDefaultVirtualService(func(vs *v1.VirtualService) *v1.VirtualService {
				vsBuilder := helpers.BuilderFromVirtualService(vs)
				vsBuilder.WithVirtualHostOptions(&gloov1.VirtualHostOptions{
					StagedTransformations: &transformation.TransformationStages{
						Regular: &transformation.RequestResponseTransformations{
							ResponseTransforms: []*transformation.ResponseMatch{{
								ResponseTransformation: &transformation.Transformation{
									TransformationType: &transformation.Transformation_TransformationTemplate{
										TransformationTemplate: &envoytransformation.TransformationTemplate{
											ParseBodyBehavior: envoytransformation.TransformationTemplate_DontParse,
											BodyTransformation: &envoytransformation.TransformationTemplate_Body{
												Body: &envoytransformation.InjaTemplate{
													Text: "{{base64_encode(body())}}",
												},
											},
										},
									},
								},
							}},
						},
					},
				})
				return vsBuilder.Build()
			})

			req := newRequestWithBody("test")
			// send a request, expect that the response body is base64 encoded
			Eventually(func(g Gomega) {
				res, err := http.DefaultClient.Do(req)
				g.Expect(err).NotTo(HaveOccurred())
				g.Expect(res).Should(testmatchers.HaveHttpResponse(&testmatchers.HttpResponse{
					StatusCode: http.StatusOK,
					Body:       base64.StdEncoding.EncodeToString([]byte("test")),
				}))
			}, "15s", ".5s").Should(Succeed())
		})

		It("Should be able to base64 decode the body", func() {
			testContext.PatchDefaultVirtualService(func(vs *v1.VirtualService) *v1.VirtualService {
				vsBuilder := helpers.BuilderFromVirtualService(vs)
				vsBuilder.WithVirtualHostOptions(&gloov1.VirtualHostOptions{
					StagedTransformations: &transformation.TransformationStages{
						Regular: &transformation.RequestResponseTransformations{
							ResponseTransforms: []*transformation.ResponseMatch{{
								ResponseTransformation: &transformation.Transformation{
									TransformationType: &transformation.Transformation_TransformationTemplate{
										TransformationTemplate: &envoytransformation.TransformationTemplate{
											ParseBodyBehavior: envoytransformation.TransformationTemplate_DontParse,
											BodyTransformation: &envoytransformation.TransformationTemplate_Body{
												Body: &envoytransformation.InjaTemplate{
													Text: "{{base64_decode(body())}}",
												},
											},
										},
									},
								},
							}},
						},
					},
				})
				return vsBuilder.Build()
			})

			// send a request, expect that the response body is base64 decoded
			body := "test"
			encodedBody := base64.StdEncoding.EncodeToString([]byte(body))
			req := newRequestWithBody(encodedBody)
			Eventually(func(g Gomega) {
				res, err := http.DefaultClient.Do(req)
				g.Expect(err).NotTo(HaveOccurred())
				g.Expect(res).Should(testmatchers.HaveHttpResponse(&testmatchers.HttpResponse{
					StatusCode: http.StatusOK,
					Body:       body,
				}))
			}, "15s", ".5s").Should(Succeed())
		})

		It("Can extract a substring from the body", func() {
			testContext.PatchDefaultVirtualService(func(vs *v1.VirtualService) *v1.VirtualService {
				vsBuilder := helpers.BuilderFromVirtualService(vs)
				vsBuilder.WithVirtualHostOptions(&gloov1.VirtualHostOptions{
					StagedTransformations: &transformation.TransformationStages{
						Regular: &transformation.RequestResponseTransformations{
							ResponseTransforms: []*transformation.ResponseMatch{{
								ResponseTransformation: &transformation.Transformation{
									TransformationType: &transformation.Transformation_TransformationTemplate{
										TransformationTemplate: &envoytransformation.TransformationTemplate{
											ParseBodyBehavior: envoytransformation.TransformationTemplate_DontParse,
											BodyTransformation: &envoytransformation.TransformationTemplate_Body{
												Body: &envoytransformation.InjaTemplate{
													Text: "{{substring(body(), 0, 4)}}",
												},
											},
										},
									},
								},
							}},
						},
					},
				})
				return vsBuilder.Build()
			})

			req := newRequestWithBody("123456789")
			// send a request, expect that the response body contains only the first 4 characters
			Eventually(func(g Gomega) {
				res, err := http.DefaultClient.Do(req)
				g.Expect(err).NotTo(HaveOccurred())
				g.Expect(res).Should(testmatchers.HaveHttpResponse(&testmatchers.HttpResponse{
					StatusCode: http.StatusOK,
					Body:       "1234",
				}))
			}, "15s", ".5s").Should(Succeed())
		})

		It("Can base64 decode and transform headers", func() {
			testContext.PatchDefaultVirtualService(func(vs *v1.VirtualService) *v1.VirtualService {
				vsBuilder := helpers.BuilderFromVirtualService(vs)
				vsBuilder.WithVirtualHostOptions(&gloov1.VirtualHostOptions{
					StagedTransformations: &transformation.TransformationStages{
						Regular: &transformation.RequestResponseTransformations{
							ResponseTransforms: []*transformation.ResponseMatch{{
								ResponseTransformation: &transformation.Transformation{
									TransformationType: &transformation.Transformation_TransformationTemplate{
										TransformationTemplate: &envoytransformation.TransformationTemplate{
											ParseBodyBehavior: envoytransformation.TransformationTemplate_DontParse,
											Headers: map[string]*envoytransformation.InjaTemplate{
												// decode the x-custom-header header and then extract a substring
												"x-new-custom-header": {Text: `{{substring(base64_decode(request_header("x-custom-header")), 6, 5)}}`},
											},
										},
									},
								},
							}},
						},
					},
				})
				return vsBuilder.Build()
			})

			req := newRequestWithBody("")
			req.Header.Add("x-custom-header", base64.StdEncoding.EncodeToString([]byte("test1.test2")))

			Eventually(func(g Gomega) {
				res, err := http.DefaultClient.Do(req)
				g.Expect(err).NotTo(HaveOccurred())
				g.Expect(res).To(testmatchers.HaveOkResponseWithHeaders(map[string]interface{}{
					"X-New-Custom-Header": ContainSubstring("test2"),
				}))
			}, "15s", ".5s").Should(Succeed())
		})

		It("should apply transforms from most specific level only", func() {
			testContext.PatchDefaultVirtualService(func(vs *v1.VirtualService) *v1.VirtualService {
				vsBuilder := helpers.BuilderFromVirtualService(vs)
				vsBuilder.WithVirtualHostOptions(&gloov1.VirtualHostOptions{
					StagedTransformations: &transformation.TransformationStages{
						Regular: &transformation.RequestResponseTransformations{
							ResponseTransforms: []*transformation.ResponseMatch{{
								ResponseTransformation: &transformation.Transformation{
									TransformationType: &transformation.Transformation_TransformationTemplate{
										TransformationTemplate: &envoytransformation.TransformationTemplate{
											ParseBodyBehavior: envoytransformation.TransformationTemplate_DontParse,
											Headers: map[string]*envoytransformation.InjaTemplate{
												"x-solo-1": {Text: "vhost header"},
											},
										},
									},
								},
							}},
						},
					},
				})
				vsBuilder.WithRouteOptions("test", &gloov1.RouteOptions{
					StagedTransformations: &transformation.TransformationStages{
						Regular: &transformation.RequestResponseTransformations{
							ResponseTransforms: []*transformation.ResponseMatch{{
								ResponseTransformation: &transformation.Transformation{
									TransformationType: &transformation.Transformation_TransformationTemplate{
										TransformationTemplate: &envoytransformation.TransformationTemplate{
											ParseBodyBehavior: envoytransformation.TransformationTemplate_DontParse,
											Headers: map[string]*envoytransformation.InjaTemplate{
												"x-solo-2": {Text: "route header"},
											},
										},
									},
								},
							}},
						},
					},
				})
				return vsBuilder.Build()
			})

			req := newRequestWithBody("")
			Eventually(func(g Gomega) {
				response, err := http.DefaultClient.Do(req)
				g.Expect(err).NotTo(HaveOccurred())
				// Only route level transformations should be applied here due to the nature of envoy choosing
				// the most specific config (weighted cluster > route > vhost)
				// This behaviour can be overridden (in the control plane) by using `inheritableTransformations` to merge
				// transformations down to the route level.
				g.Expect(response).To(testmatchers.HaveOkResponseWithHeaders(map[string]interface{}{
					"x-solo-2": Equal("route header"),
					"x-solo-1": BeEmpty(),
				}))
			}).Should(Succeed())
		})
	})

	/*
		Context("with auth", func() {

			BeforeEach(func() {
				// this upstream doesn't need to exist - in fact, we want ext auth to fail.
				extAuthUpstream := &gloov1.Upstream{
					Metadata: &core.Metadata{
						Name:      "extauth-server",
						Namespace: "default",
					},
					UseHttp2: &wrappers.BoolValue{Value: true},
					UpstreamType: &gloov1.Upstream_Static{
						Static: &gloov1static.UpstreamSpec{
							Hosts: []*gloov1static.Host{{
								Addr: "127.2.3.4",
								Port: 1234,
							}},
						},
					},
				}

				testContext.ResourcesToCreate().Upstreams = append(testContext.ResourcesToCreate().Upstreams, extAuthUpstream)

				testContext.SetRunSettings(&gloov1.Settings{Extauth: &extauthv1.Settings{
					ExtauthzServerRef: extAuthUpstream.GetMetadata().Ref(),
				}})
			})

			Context("disabled", func() {


			})

			Context("enabled", func() {

			})

			TestUpstreamReachable := func() {
				Eventually(func() error {
					resp, err := http.DefaultClient.Get(fmt.Sprintf("http://localhost:%d/1", envoyPort))
					if resp != nil && resp.StatusCode != 403 {
						return errors.New("Expected status 403")
					}
					return err
				}, "30s", "1s").ShouldNot(HaveOccurred())
			}

			It("should transform response code details", func() {
				setProxyWithModifier(&transformation.TransformationStages{
					Early: &transformation.RequestResponseTransformations{
						ResponseTransforms: []*transformation.ResponseMatch{{
							ResponseCodeDetails: "ext_authz_error",
							ResponseTransformation: &transformation.Transformation{
								TransformationType: &transformation.Transformation_TransformationTemplate{
									TransformationTemplate: &envoytransformation.TransformationTemplate{
										ParseBodyBehavior: envoytransformation.TransformationTemplate_DontParse,
										BodyTransformation: &envoytransformation.TransformationTemplate_Body{
											Body: &envoytransformation.InjaTemplate{
												Text: "early-transformed",
											},
										},
									},
								},
							},
						}},
					},
				}, func(vs *gloov1.VirtualHost) {
					vs.Options.Extauth = &extauthv1.ExtAuthExtension{
						Spec: &extauthv1.ExtAuthExtension_CustomAuth{
							CustomAuth: &extauthv1.CustomAuth{},
						},
					}
				})
				TestUpstreamReachable()
				// send a request and expect it transformed!
				res, err := http.DefaultClient.Get(fmt.Sprintf("http://localhost:%d/1", envoyPort))
				Expect(err).NotTo(HaveOccurred())

				body, err := ioutil.ReadAll(res.Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(body)).To(Equal("early-transformed"))
			})

			It("should not transform when auth succeeds", func() {
				testContext.PatchDefaultVirtualService(func(vs *v1.VirtualService) {
					vs.GetVirtualHost().Options = &gloov1.VirtualHostOptions{
						StagedTransformations: &transformation.TransformationStages{
							Early: &transformation.RequestResponseTransformations{
								ResponseTransforms: []*transformation.ResponseMatch{{
									ResponseCodeDetails: "ext_authz_error",
									ResponseTransformation: &transformation.Transformation{
										TransformationType: &transformation.Transformation_TransformationTemplate{
											TransformationTemplate: &envoytransformation.TransformationTemplate{
												ParseBodyBehavior: envoytransformation.TransformationTemplate_DontParse,
												BodyTransformation: &envoytransformation.TransformationTemplate_Body{
													Body: &envoytransformation.InjaTemplate{
														Text: "early-transformed",
													},
												},
											},
										},
									},
								}},
							},
						},
					}
				})

				// send a request and expect it not transformed!
				eventuallyRequestMatches("test", testmatchers.HaveExactResponseBody("test"))
			})
		})

	*/

})
