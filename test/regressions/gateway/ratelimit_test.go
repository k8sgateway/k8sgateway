package gateway_test

import (
	"context"
	"fmt"
	"time"

	ratelimit2 "github.com/solo-io/gloo/projects/gloo/api/external/solo/ratelimit"
	v1alpha1 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/solo/ratelimit"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-projects/test/regressions"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	rlv1alpha1 "github.com/solo-io/solo-apis/pkg/api/ratelimit.solo.io/v1alpha1"

	extauthv1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/extauth/v1"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

	v2 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"

	"github.com/solo-io/go-utils/testutils/helper"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gateway/pkg/defaults"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"

	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ratelimit"
	"github.com/solo-io/go-utils/kubeutils"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	ratelimitpb "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ratelimit"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube"

	"k8s.io/client-go/rest"
)

var _ = Describe("RateLimit tests", func() {

	var (
		ctx    context.Context
		cancel context.CancelFunc
		cfg    *rest.Config

		cache                 kube.SharedCache
		gatewayClient         v2.GatewayClient
		virtualServiceClient  v1.VirtualServiceClient
		rateLimitConfigClient v1alpha1.RateLimitConfigClient
		settingsClient        gloov1.SettingsClient
		origSettings          *gloov1.Settings // used to capture & restore initial Settings so each test can modify them
		uniqueDescriptorValue string
	)

	const (
		response401 = "HTTP/1.1 401 Unauthorized"
		response429 = "HTTP/1.1 429 Too Many Requests"
	)

	BeforeEach(func() {
		ctx, cancel = context.WithCancel(context.Background())
		var err error
		cfg, err = kubeutils.GetConfig("", "")
		Expect(err).NotTo(HaveOccurred())
		cache = kube.NewKubeCache(ctx)
		settingsClientFactory := &factory.KubeResourceClientFactory{
			Crd:         gloov1.SettingsCrd,
			Cfg:         cfg,
			SharedCache: cache,
		}
		settingsClient, err = gloov1.NewSettingsClient(settingsClientFactory)
		Expect(err).NotTo(HaveOccurred())
		uniqueDescriptorValue = uniqueDescriptorValue + "1"

		gatewayClientFactory := &factory.KubeResourceClientFactory{
			Crd:         v2.GatewayCrd,
			Cfg:         cfg,
			SharedCache: cache,
		}
		gatewayClient, err = v2.NewGatewayClient(gatewayClientFactory)
		Expect(err).NotTo(HaveOccurred())

		virtualServiceClientFactory := &factory.KubeResourceClientFactory{
			Crd:         v1.VirtualServiceCrd,
			Cfg:         cfg,
			SharedCache: cache,
		}
		virtualServiceClient, err = v1.NewVirtualServiceClient(virtualServiceClientFactory)
		Expect(err).NotTo(HaveOccurred())

		rateLimitConfigClientFactory := &factory.KubeResourceClientFactory{
			Crd:         v1alpha1.RateLimitConfigCrd,
			Cfg:         cfg,
			SharedCache: cache,
		}

		rateLimitConfigClient, err = v1alpha1.NewRateLimitConfigClient(rateLimitConfigClientFactory)
		Expect(err).NotTo(HaveOccurred())
	})

	BeforeEach(func() {
		var err error
		origSettings, err = settingsClient.Read(testHelper.InstallNamespace, "default", clients.ReadOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred(), "Should be able to read initial settings")
	})

	AfterEach(func() {
		regressions.DeleteVirtualService(virtualServiceClient, testHelper.InstallNamespace, "vs", clients.DeleteOpts{Ctx: ctx, IgnoreNotExist: true})

		currentSettings, err := settingsClient.Read(testHelper.InstallNamespace, "default", clients.ReadOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred(), "Should be able to read current settings")

		if origSettings.Metadata.ResourceVersion != currentSettings.Metadata.ResourceVersion {
			origSettings.Metadata.ResourceVersion = currentSettings.Metadata.ResourceVersion // so we can overwrite settings
			_, err = settingsClient.Write(origSettings, clients.WriteOpts{Ctx: ctx, OverwriteExisting: true})
			Expect(err).ToNot(HaveOccurred())
		}
		cancel()
	})

	waitForGateway := func() {
		defaultGateway := defaults.DefaultGateway(testHelper.InstallNamespace)
		// wait for default gateway to be created
		EventuallyWithOffset(2, func() (*v2.Gateway, error) {
			return gatewayClient.Read(testHelper.InstallNamespace, defaultGateway.Metadata.Name, clients.ReadOpts{})
		}, "15s", "0.5s").Should(Not(BeNil()))
	}

	checkAuthDenied := func() {
		waitForGateway()

		gatewayPort := 80
		testHelper.CurlEventuallyShouldRespond(helper.CurlOpts{
			Protocol:          "http",
			Path:              regressions.TestMatcherPrefix,
			Method:            "GET",
			Host:              defaults.GatewayProxyName,
			Service:           defaults.GatewayProxyName,
			Port:              gatewayPort,
			ConnectionTimeout: 10, // this is important, as the first curl call sometimes hangs indefinitely
			Verbose:           true,
		}, response401, 1, time.Minute*5)
	}

	checkRateLimited := func() {
		waitForGateway()

		gatewayPort := 80
		testHelper.CurlEventuallyShouldRespond(helper.CurlOpts{
			Protocol:          "http",
			Path:              regressions.TestMatcherPrefix,
			Method:            "GET",
			Host:              defaults.GatewayProxyName,
			Service:           defaults.GatewayProxyName,
			Port:              gatewayPort,
			ConnectionTimeout: 10, // this is important, as the first curl call sometimes hangs indefinitely
			Verbose:           true,
		}, response429, 1, time.Minute*5)
	}

	Context("simple rate limiting", func() {
		var (
			ingressRateLimit = &ratelimit.IngressRateLimit{
				AnonymousLimits: &rlv1alpha1.RateLimit{
					RequestsPerUnit: 1,
					Unit:            rlv1alpha1.RateLimit_HOUR,
				},
			}
			virtualHostPlugins = &gloov1.VirtualHostOptions{
				RatelimitBasic: ingressRateLimit,
			}
		)

		It("can rate limit to upstream", func() {
			regressions.WriteVirtualService(ctx, testHelper, virtualServiceClient, virtualHostPlugins, nil, nil)
			checkRateLimited()
		})
	})

	Context("raw rate limit", func() {
		BeforeEach(func() {
			// Write rate limit service config to settings
			settings, err := settingsClient.Read(testHelper.InstallNamespace, "default", clients.ReadOpts{})
			Expect(err).NotTo(HaveOccurred())

			rlSettings := ratelimitpb.ServiceSettings{
				Descriptors: []*rlv1alpha1.Descriptor{{
					Key:   "generic_key",
					Value: uniqueDescriptorValue,
					RateLimit: &rlv1alpha1.RateLimit{
						RequestsPerUnit: 0,
						Unit:            rlv1alpha1.RateLimit_SECOND,
					},
				}},
			}

			settings.Ratelimit = &rlSettings
			_, err = settingsClient.Write(settings, clients.WriteOpts{OverwriteExisting: true})

		})

		Context("with ext auth also configured", func() {
			BeforeEach(func() {
				kubeCache := kube.NewKubeCache(ctx)
				authConfigClientFactory := &factory.KubeResourceClientFactory{
					Crd:         extauthv1.AuthConfigCrd,
					Cfg:         cfg,
					SharedCache: kubeCache,
				}
				authConfigClient, err := extauthv1.NewAuthConfigClient(authConfigClientFactory)
				Expect(err).NotTo(HaveOccurred(), "Should create auth config client")
				authConfig, err := authConfigClient.Write(&extauthv1.AuthConfig{
					Metadata: core.Metadata{
						Name:      "basic-auth",
						Namespace: testHelper.InstallNamespace,
					},
					Configs: []*extauthv1.AuthConfig_Config{{
						AuthConfig: &extauthv1.AuthConfig_Config_BasicAuth{
							BasicAuth: &extauthv1.BasicAuth{
								Realm: "",
								Apr: &extauthv1.BasicAuth_Apr{
									Users: map[string]*extauthv1.BasicAuth_Apr_SaltedHashedPassword{
										"user": {
											// garbage salt and hash- we want all requests to come back as unauthorized when they hit extauth
											Salt:           "intentionally-garbage-password-salt",
											HashedPassword: "intentionally-garbage-password-hash",
										},
									},
								},
							},
						},
					}},
				}, clients.WriteOpts{Ctx: ctx})
				Expect(err).NotTo(HaveOccurred(), "Should write auth config")

				authConfigRef := authConfig.Metadata.Ref()
				extAuthConfigProto := &extauthv1.ExtAuthExtension{
					Spec: &extauthv1.ExtAuthExtension_ConfigRef{
						ConfigRef: &authConfigRef,
					},
				}

				ratelimitExtension := &ratelimitpb.RateLimitVhostExtension{
					RateLimits: []*rlv1alpha1.RateLimitActions{{
						Actions: []*rlv1alpha1.Action{{
							ActionSpecifier: &rlv1alpha1.Action_GenericKey_{
								GenericKey: &rlv1alpha1.Action_GenericKey{
									DescriptorValue: uniqueDescriptorValue,
								},
							},
						}},
					}},
				}

				virtualHostPlugins := &gloov1.VirtualHostOptions{
					RateLimitConfigType: &gloov1.VirtualHostOptions_Ratelimit{
						Ratelimit: ratelimitExtension,
					},
					Extauth: extAuthConfigProto,
				}

				settings, err := settingsClient.Read(testHelper.InstallNamespace, "default", clients.ReadOpts{})
				Expect(err).NotTo(HaveOccurred(), "Should read settings")

				timeout := time.Second
				settings.RatelimitServer = &ratelimit.Settings{
					RatelimitServerRef: &core.ResourceRef{
						Name:      "rate-limit",
						Namespace: testHelper.InstallNamespace,
					},
					RequestTimeout:      &timeout,
					RateLimitBeforeAuth: false, // start as false to make sure that we correctly get denied by authZ before rate limited
				}
				settings.Extauth = &extauthv1.Settings{
					ExtauthzServerRef: &core.ResourceRef{
						Name:      "extauth",
						Namespace: testHelper.InstallNamespace,
					},
				}
				_, err = settingsClient.Write(settings, clients.WriteOpts{OverwriteExisting: true})
				Expect(err).NotTo(HaveOccurred(), "Should write settings")
				regressions.WriteVirtualService(ctx, testHelper, virtualServiceClient, virtualHostPlugins, nil, nil)

				// should hit auth before getting rate limited by default
				checkAuthDenied()

				settings, err = settingsClient.Read(testHelper.InstallNamespace, "default", clients.ReadOpts{})
				Expect(err).NotTo(HaveOccurred(), "Should read settings to set RateLimitBeforeAuth")

				settings.RatelimitServer.RateLimitBeforeAuth = true

				_, err = settingsClient.Write(settings, clients.WriteOpts{OverwriteExisting: true})
				Expect(err).NotTo(HaveOccurred(), "Should write settings with RateLimitBeforeAuth set")
			})

			It("can rate limit before hitting the auth server when so configured", func() {
				// normally, ext auth runs before rate limiting. So since we've set up ext auth to block every request that comes in,
				// we would normally expect all requests to come back with a 401. But we've *also* set `RateLimitBeforeAuth` on the rate
				// limit settings, which means that now we expect rate limit to run before ext auth. So eventually, this next function
				// call will result in curl eventually NOT receiving a 401 and instead receiving a 429, as expected
				checkRateLimited()
			})
		})

		Context("using the inline config format", func() {

			It("can rate limit to upstream vhost", func() {

				ratelimitExtension := &ratelimitpb.RateLimitVhostExtension{
					RateLimits: []*rlv1alpha1.RateLimitActions{{
						Actions: []*rlv1alpha1.Action{{
							ActionSpecifier: &rlv1alpha1.Action_GenericKey_{
								GenericKey: &rlv1alpha1.Action_GenericKey{
									DescriptorValue: uniqueDescriptorValue,
								},
							},
						}},
					}},
				}

				virtualHostPlugins := &gloov1.VirtualHostOptions{
					RateLimitConfigType: &gloov1.VirtualHostOptions_Ratelimit{
						Ratelimit: ratelimitExtension,
					},
				}

				regressions.WriteVirtualService(ctx, testHelper, virtualServiceClient, virtualHostPlugins, nil, nil)
				checkRateLimited()
			})

			It("can rate limit to upstream route", func() {

				ratelimitExtension := &ratelimitpb.RateLimitRouteExtension{
					RateLimits: []*rlv1alpha1.RateLimitActions{{
						Actions: []*rlv1alpha1.Action{{
							ActionSpecifier: &rlv1alpha1.Action_GenericKey_{
								GenericKey: &rlv1alpha1.Action_GenericKey{
									DescriptorValue: uniqueDescriptorValue,
								},
							},
						}},
					}},
				}

				routePlugins := &gloov1.RouteOptions{
					RateLimitConfigType: &gloov1.RouteOptions_Ratelimit{
						Ratelimit: ratelimitExtension,
					},
				}

				regressions.WriteVirtualService(ctx, testHelper, virtualServiceClient, nil, routePlugins, nil)
				checkRateLimited()
			})

			It("can rate limit to upstream route when config is inherited by parent virtual host", func() {

				vhostRateLimitExtension := &ratelimitpb.RateLimitVhostExtension{
					RateLimits: []*rlv1alpha1.RateLimitActions{{
						Actions: []*rlv1alpha1.Action{{
							ActionSpecifier: &rlv1alpha1.Action_GenericKey_{
								GenericKey: &rlv1alpha1.Action_GenericKey{
									DescriptorValue: uniqueDescriptorValue,
								},
							},
						}},
					}},
				}

				virtualHostPlugins := &gloov1.VirtualHostOptions{
					RateLimitConfigType: &gloov1.VirtualHostOptions_Ratelimit{
						Ratelimit: vhostRateLimitExtension,
					},
				}

				routeRateLimitExtension := &ratelimitpb.RateLimitRouteExtension{
					IncludeVhRateLimits: true,
				}

				routePlugins := &gloov1.RouteOptions{
					RateLimitConfigType: &gloov1.RouteOptions_Ratelimit{
						Ratelimit: routeRateLimitExtension,
					},
				}

				regressions.WriteVirtualService(ctx, testHelper, virtualServiceClient, virtualHostPlugins, routePlugins, nil)
				checkRateLimited()
			})
		})

	})

	Context("using a RateLimitConfig resource", func() {

		var configRef core.ResourceRef

		BeforeEach(func() {
			configRef = core.ResourceRef{
				Name:      fmt.Sprintf("%s-%s-%v", testHelper.InstallNamespace, "testrunner", helper.TestRunnerPort),
				Namespace: testHelper.InstallNamespace,
			}

			_, err := rateLimitConfigClient.Read(configRef.Namespace, configRef.Name, clients.ReadOpts{Ctx: ctx})
			if err == nil {
				return // already exists nothing to do
			}
			if !errors.IsNotExist(err) {
				Expect(err).NotTo(HaveOccurred())
			}

			rateLimitConfig := &v1alpha1.RateLimitConfig{
				RateLimitConfig: ratelimit2.RateLimitConfig{
					ObjectMeta: metav1.ObjectMeta{
						Name:      configRef.Name,
						Namespace: configRef.Namespace,
					},
					Spec: rlv1alpha1.RateLimitConfigSpec{
						ConfigType: &rlv1alpha1.RateLimitConfigSpec_Raw_{
							Raw: &rlv1alpha1.RateLimitConfigSpec_Raw{
								Descriptors: []*rlv1alpha1.Descriptor{{
									Key:   "generic_key",
									Value: "foo",
									RateLimit: &rlv1alpha1.RateLimit{
										Unit:            rlv1alpha1.RateLimit_MINUTE,
										RequestsPerUnit: 1,
									},
								}},
								RateLimits: []*rlv1alpha1.RateLimitActions{{
									Actions: []*rlv1alpha1.Action{{
										ActionSpecifier: &rlv1alpha1.Action_GenericKey_{
											GenericKey: &rlv1alpha1.Action_GenericKey{
												DescriptorValue: "foo",
											},
										},
									}},
								}},
							},
						},
					},
				},
			}
			_, err = rateLimitConfigClient.Write(rateLimitConfig, clients.WriteOpts{OverwriteExisting: false})
			Expect(err).NotTo(HaveOccurred())

		})

		It("works when the resource is referenced from a virtual host", func() {
			virtualHostPlugins := &gloov1.VirtualHostOptions{
				RateLimitConfigType: &gloov1.VirtualHostOptions_RateLimitConfigs{
					RateLimitConfigs: &ratelimitpb.RateLimitConfigRefs{
						Refs: []*ratelimitpb.RateLimitConfigRef{{
							Namespace: configRef.Namespace,
							Name:      configRef.Name,
						}},
					},
				},
			}

			regressions.WriteVirtualService(ctx, testHelper, virtualServiceClient, virtualHostPlugins, nil, nil)
			checkRateLimited()
		})

		It("works when the resource is referenced from a route", func() {
			routePlugins := &gloov1.RouteOptions{
				RateLimitConfigType: &gloov1.RouteOptions_RateLimitConfigs{
					RateLimitConfigs: &ratelimitpb.RateLimitConfigRefs{
						Refs: []*ratelimitpb.RateLimitConfigRef{{
							Namespace: configRef.Namespace,
							Name:      configRef.Name,
						}},
					},
				},
			}

			regressions.WriteVirtualService(ctx, testHelper, virtualServiceClient, nil, routePlugins, nil)
			checkRateLimited()
		})
	})
})
