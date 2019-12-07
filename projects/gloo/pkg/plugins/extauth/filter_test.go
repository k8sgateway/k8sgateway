package extauth_test

import (
	"time"

	envoycore "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	envoyauth "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/ext_authz/v2"
	envoytype "github.com/envoyproxy/go-control-plane/envoy/type"
	envoymatcher "github.com/envoyproxy/go-control-plane/envoy/type/matcher"
	"github.com/golang/protobuf/ptypes/duration"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/pkg/utils/gogoutils"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	extauthv1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/extauth/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/static"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	. "github.com/solo-io/gloo/projects/gloo/pkg/plugins/extauth"
	"github.com/solo-io/gloo/projects/gloo/pkg/translator"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

var _ = Describe("Extauth Http filter builder function", func() {

	When("no extauth settings are provided", func() {
		It("does not return any filter", func() {
			filters, err := BuildHttpFilters(nil, nil)
			Expect(err).NotTo(HaveOccurred())
			Expect(filters).To(HaveLen(0))
		})
	})

	When("settings do not contain ext auth server ref", func() {
		It("returns an error", func() {
			_, err := BuildHttpFilters(&extauthv1.Settings{}, nil)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(NoServerRefErr))
		})
	})

	When("settings contain incorrect ext auth server ref", func() {
		It("returns an error", func() {
			invalidUs := &core.ResourceRef{
				Name:      "non",
				Namespace: "existent",
			}
			_, err := BuildHttpFilters(&extauthv1.Settings{ExtauthzServerRef: invalidUs}, nil)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(ServerNotFound(invalidUs)))
		})
	})

	Context("settings point to a valid ext auth server", func() {

		var (
			upstream       *gloov1.Upstream
			settings       *extauthv1.Settings
			expectedConfig *envoyauth.ExtAuthz
		)

		getExtAuthz := func(extAuthFilter plugins.StagedHttpFilter) *envoyauth.ExtAuthz {
			ExpectWithOffset(1, extAuthFilter).NotTo(BeNil())
			ExpectWithOffset(1, extAuthFilter.HttpFilter.Name).To(Equal(FilterName))

			filterConfig := &envoyauth.ExtAuthz{}
			err := translator.ParseConfig(extAuthFilter.HttpFilter, filterConfig)
			ExpectWithOffset(1, err).NotTo(HaveOccurred())

			return filterConfig
		}

		BeforeEach(func() {

			upstream = &gloov1.Upstream{
				Metadata: core.Metadata{
					Name:      "extauth",
					Namespace: "default",
				},
				UpstreamType: &gloov1.Upstream_Static{
					Static: &static.UpstreamSpec{
						Hosts: []*static.Host{{
							Addr: "test",
							Port: 1234,
						}},
					},
				},
			}
		})

		When("minimal settings are provided", func() {

			BeforeEach(func() {
				usRef := upstream.Metadata.Ref()

				settings = &extauthv1.Settings{
					ExtauthzServerRef: &usRef,
				}

				expectedConfig = &envoyauth.ExtAuthz{
					Services: &envoyauth.ExtAuthz_GrpcService{
						GrpcService: &envoycore.GrpcService{
							Timeout: &duration.Duration{
								Nanos: int32(DefaultTimeout),
							},
							TargetSpecifier: &envoycore.GrpcService_EnvoyGrpc_{
								EnvoyGrpc: &envoycore.GrpcService_EnvoyGrpc{
									ClusterName: translator.UpstreamToClusterName(usRef),
								},
							},
						},
					},
				}
			})

			It("uses the expected defaults", func() {
				filters, err := BuildHttpFilters(settings, gloov1.UpstreamList{upstream})
				Expect(err).NotTo(HaveOccurred())
				Expect(filters).To(HaveLen(1))

				actualFilterConfig := getExtAuthz(filters[0])
				Expect(actualFilterConfig).To(Equal(expectedConfig))
			})
		})

		When("complete settings are provided", func() {

			BeforeEach(func() {
				usRef := upstream.Metadata.Ref()

				customTimeout := 500 * time.Millisecond

				settings = &extauthv1.Settings{
					ExtauthzServerRef: &usRef,
					RequestTimeout:    &customTimeout,
					FailureModeAllow:  true,
					RequestBody: &extauthv1.BufferSettings{
						AllowPartialMessage: true,
						MaxRequestBytes:     54,
					},
					ClearRouteCache: true,
					StatusOnError:   400,
				}

				expectedConfig = &envoyauth.ExtAuthz{
					Services: &envoyauth.ExtAuthz_GrpcService{
						GrpcService: &envoycore.GrpcService{
							Timeout: &duration.Duration{
								Nanos: int32(customTimeout),
							},
							TargetSpecifier: &envoycore.GrpcService_EnvoyGrpc_{
								EnvoyGrpc: &envoycore.GrpcService_EnvoyGrpc{
									ClusterName: translator.UpstreamToClusterName(usRef),
								},
							},
						},
					},
					FailureModeAllow: true,
					WithRequestBody: &envoyauth.BufferSettings{
						AllowPartialMessage: true,
						MaxRequestBytes:     54,
					},
					ClearRouteCache: true,
					StatusOnError:   &envoytype.HttpStatus{Code: envoytype.StatusCode_BadRequest},
				}
			})

			It("generates the expected configuration", func() {
				filters, err := BuildHttpFilters(settings, gloov1.UpstreamList{upstream})
				Expect(err).NotTo(HaveOccurred())
				Expect(filters).To(HaveLen(1))

				actualFilterConfig := getExtAuthz(filters[0])
				Expect(actualFilterConfig).To(Equal(expectedConfig))
			})
		})

		When("invalid settings are provided", func() {

			BeforeEach(func() {
				usRef := upstream.Metadata.Ref()

				settings = &extauthv1.Settings{
					ExtauthzServerRef: &usRef,
					// This is the only thing that can go wrong in the BuildHttpFilters function
					StatusOnError: 999,
				}
			})

			It("returns an error", func() {
				_, err := BuildHttpFilters(settings, gloov1.UpstreamList{upstream})
				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(InvalidStatusOnErrorErr(999)))
			})
		})

		When("an HTTP service is configured", func() {

			BeforeEach(func() {
				usRef := upstream.Metadata.Ref()

				settings = &extauthv1.Settings{
					ExtauthzServerRef: &usRef,
					HttpService: &extauthv1.HttpService{
						PathPrefix: "/foo",
						Request: &extauthv1.HttpService_Request{
							AllowedHeaders: []string{"allowed-header"},
							HeadersToAdd:   map[string]string{"header": "add"},
						},
						Response: &extauthv1.HttpService_Response{
							AllowedClientHeaders:   []string{"allowed-client-header"},
							AllowedUpstreamHeaders: []string{"allowed-upstream-header"},
						},
					},
				}

				expectedConfig = &envoyauth.ExtAuthz{
					Services: &envoyauth.ExtAuthz_HttpService{
						HttpService: &envoyauth.HttpService{
							AuthorizationRequest: &envoyauth.AuthorizationRequest{
								AllowedHeaders: &envoymatcher.ListStringMatcher{
									Patterns: []*envoymatcher.StringMatcher{{
										MatchPattern: &envoymatcher.StringMatcher_Exact{Exact: "allowed-header"},
									}},
								},
								HeadersToAdd: []*envoycore.HeaderValue{{
									Key:   "header",
									Value: "add",
								}},
							},
							AuthorizationResponse: &envoyauth.AuthorizationResponse{
								AllowedClientHeaders: &envoymatcher.ListStringMatcher{
									Patterns: []*envoymatcher.StringMatcher{{
										MatchPattern: &envoymatcher.StringMatcher_Exact{Exact: "allowed-client-header"},
									}},
								},
								AllowedUpstreamHeaders: &envoymatcher.ListStringMatcher{
									Patterns: []*envoymatcher.StringMatcher{{
										MatchPattern: &envoymatcher.StringMatcher_Exact{Exact: "allowed-upstream-header"},
									}},
								},
							},
							PathPrefix: "/foo",
							ServerUri: &envoycore.HttpUri{
								Timeout: gogoutils.DurationStdToProto(&DefaultTimeout),
								Uri:     HttpServerUri,
								HttpUpstreamType: &envoycore.HttpUri_Cluster{
									Cluster: translator.UpstreamToClusterName(usRef),
								},
							},
						},
					},
				}
			})

			It("uses the expected defaults", func() {
				filters, err := BuildHttpFilters(settings, gloov1.UpstreamList{upstream})
				Expect(err).NotTo(HaveOccurred())
				Expect(filters).To(HaveLen(1))

				actualFilterConfig := getExtAuthz(filters[0])
				Expect(actualFilterConfig).To(Equal(expectedConfig))
			})
		})
	})
})
