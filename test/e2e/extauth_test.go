package e2e_test

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os/exec"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/solo-io/ext-auth-service/pkg/server"

	"github.com/solo-io/ext-auth-service/pkg/config/oauth/test_utils"
	"github.com/solo-io/ext-auth-service/pkg/config/oauth/user_info"
	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/gogo/protobuf/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	. "github.com/onsi/gomega/gstruct"

	"github.com/solo-io/ext-auth-service/pkg/config/oidc"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	extauthrunner "github.com/solo-io/solo-projects/projects/extauth/pkg/runner"
	"github.com/solo-io/solo-projects/test/services"

	extauth "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/extauth/v1"

	"github.com/dgrijalva/jwt-go"
	"github.com/fgrosse/zaptest"
	"github.com/solo-io/gloo/pkg/utils"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	gloov1static "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/static"
	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-projects/test/v1helpers"
)

var (
	baseExtauthPort = uint32(27000)
)

var _ = Describe("External auth", func() {

	var (
		ctx         context.Context
		cancel      context.CancelFunc
		testClients services.TestClients
		settings    extauthrunner.Settings
		cache       memory.InMemoryResourceCache
	)

	BeforeEach(func() {
		extAuthPort := atomic.AddUint32(&baseExtauthPort, 1) + uint32(config.GinkgoConfig.ParallelNode*1000)

		logger := zaptest.LoggerWriter(GinkgoWriter)
		contextutils.SetFallbackLogger(logger.Sugar())

		ctx, cancel = context.WithCancel(context.Background())
		cache = memory.NewInMemoryResourceCache()

		testClients = services.GetTestClients(ctx, cache)
		testClients.GlooPort = int(services.AllocateGlooPort())

		extauthAddr := "localhost"
		if runtime.GOOS == "darwin" {
			extauthAddr = "host.docker.internal"
		}

		extAuthServer := &gloov1.Upstream{
			Metadata: core.Metadata{
				Name:      "extauth-server",
				Namespace: "default",
			},
			UseHttp2: &types.BoolValue{Value: true},
			UpstreamType: &gloov1.Upstream_Static{
				Static: &gloov1static.UpstreamSpec{
					Hosts: []*gloov1static.Host{{
						Addr: extauthAddr,
						Port: extAuthPort,
					}},
				},
			},
		}

		_, err := testClients.UpstreamClient.Write(extAuthServer, clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())

		ref := extAuthServer.Metadata.Ref()
		extauthSettings := &extauth.Settings{
			ExtauthzServerRef: &ref,
		}

		settings = extauthrunner.Settings{
			GlooAddress: fmt.Sprintf("localhost:%d", testClients.GlooPort),
			ExtAuthSettings: server.Settings{
				DebugPort:              0,
				ServerPort:             int(extAuthPort),
				SigningKey:             "hello",
				UserIdHeader:           "X-User-Id",
				HealthCheckFailTimeout: 2, // seconds
			},
		}
		glooSettings := &gloov1.Settings{Extauth: extauthSettings}

		what := services.What{
			DisableGateway: true,
			DisableUds:     true,
			DisableFds:     true,
		}

		services.RunGlooGatewayUdsFdsOnPort(ctx, cache, int32(testClients.GlooPort), what, defaults.GlooSystem, nil, nil, glooSettings)
		go func(testCtx context.Context) {
			defer GinkgoRecover()
			err := extauthrunner.RunWithSettings(testCtx, settings)
			if testCtx.Err() == nil {
				Expect(err).NotTo(HaveOccurred())
			}
		}(ctx)
	})

	AfterEach(func() {
		cancel()
	})

	Context("With envoy", func() {

		var (
			envoyInstance *services.EnvoyInstance
			testUpstream  *v1helpers.TestUpstream
			envoyPort     = uint32(8080)
		)

		BeforeEach(func() {
			var err error
			envoyInstance, err = envoyFactory.NewEnvoyInstance()
			Expect(err).NotTo(HaveOccurred())

			err = envoyInstance.Run(testClients.GlooPort)
			Expect(err).NotTo(HaveOccurred())

			testUpstream = v1helpers.NewTestHttpUpstream(ctx, envoyInstance.LocalAddr())

			var opts clients.WriteOpts
			up := testUpstream.Upstream
			_, err = testClients.UpstreamClient.Write(up, opts)
			Expect(err).NotTo(HaveOccurred())

		})

		AfterEach(func() {
			if envoyInstance != nil {
				_ = envoyInstance.Clean()
			}
		})

		var basicConfigSetup = func() {
			_, err := testClients.AuthConfigClient.Write(&extauth.AuthConfig{
				Metadata: core.Metadata{
					Name:      GetBasicAuthExtension().GetConfigRef().Name,
					Namespace: GetBasicAuthExtension().GetConfigRef().Namespace,
				},
				Configs: []*extauth.AuthConfig_Config{{
					AuthConfig: &extauth.AuthConfig_Config_BasicAuth{
						BasicAuth: getBasicAuthConfig(),
					},
				}},
			}, clients.WriteOpts{Ctx: ctx})
			ExpectWithOffset(1, err).NotTo(HaveOccurred())

			proxy := getProxyExtAuthBasicAuth(envoyPort, testUpstream.Upstream.Metadata.Ref())

			_, err = testClients.ProxyClient.Write(proxy, clients.WriteOpts{Ctx: ctx})
			ExpectWithOffset(1, err).NotTo(HaveOccurred())

			EventuallyWithOffset(1, func() (core.Status, error) {
				proxy, err := testClients.ProxyClient.Read(proxy.Metadata.Namespace, proxy.Metadata.Name, clients.ReadOpts{})
				if err != nil {
					return core.Status{}, err
				}

				return proxy.Status, nil
			}, "5s", "0.1s").Should(MatchFields(IgnoreExtras, Fields{
				"Reason": BeEmpty(),
				"State":  Equal(core.Status_Accepted),
			}))
		}

		Context("using new config format", func() {

			Context("basic auth sanity tests", func() {

				BeforeEach(func() {

					// drain channel as we dont care about it
					go func(testUpstream v1helpers.TestUpstream) {
						for range testUpstream.C {
						}
					}(*testUpstream)

					basicConfigSetup()
				})

				It("should deny ext auth envoy", func() {
					Eventually(func() (int, error) {
						resp, err := http.Get(fmt.Sprintf("http://%s:%d/1", "localhost", envoyPort))
						if err != nil {
							return 0, err
						}
						return resp.StatusCode, nil
					}, "5s", "0.5s").Should(Equal(http.StatusUnauthorized))
				})

				It("should allow ext auth envoy", func() {
					Eventually(func() (int, error) {
						resp, err := http.Get(fmt.Sprintf("http://user:password@%s:%d/1", "localhost", envoyPort))
						if err != nil {
							return 0, err
						}
						return resp.StatusCode, nil
					}, "5s", "0.5s").Should(Equal(http.StatusOK))
				})

				It("should deny ext auth with wrong password", func() {
					Eventually(func() (int, error) {
						resp, err := http.Get(fmt.Sprintf("http://user:password2@%s:%d/1", "localhost", envoyPort))
						if err != nil {
							return 0, err
						}
						return resp.StatusCode, nil
					}, "5s", "0.5s").Should(Equal(http.StatusUnauthorized))
				})
			})

			Context("oidc sanity", func() {
				var (
					authConfig      *extauth.AuthConfig
					oauth2          *extauth.OAuth2_OidcAuthorizationCode
					privateKey      *rsa.PrivateKey
					discoveryServer fakeDiscoveryServer
					secret          *gloov1.Secret
					proxy           *gloov1.Proxy
					token           string
					cookies         []*http.Cookie
				)
				BeforeEach(func() {
					discoveryServer = fakeDiscoveryServer{}

					privateKey = discoveryServer.Start()

					clientSecret := &extauth.OauthSecret{
						ClientSecret: "test",
					}

					secret = &gloov1.Secret{
						Metadata: core.Metadata{
							Name:      "secret",
							Namespace: "default",
						},
						Kind: &gloov1.Secret_Oauth{
							Oauth: clientSecret,
						},
					}
					_, err := testClients.SecretClient.Write(secret, clients.WriteOpts{})
					Expect(err).NotTo(HaveOccurred())
					oauth2 = getOidcAuthCodeConfig(envoyPort, secret.Metadata.Ref())
					authConfig = &extauth.AuthConfig{
						Metadata: core.Metadata{
							Name:      getOidcExtAuthExtension().GetConfigRef().Name,
							Namespace: getOidcExtAuthExtension().GetConfigRef().Namespace,
						},
						Configs: []*extauth.AuthConfig_Config{{
							AuthConfig: &extauth.AuthConfig_Config_Oauth2{
								Oauth2: &extauth.OAuth2{
									OauthType: oauth2,
								},
							},
						}},
					}

					proxy = getProxyExtAuthOIDC(envoyPort, testUpstream.Upstream.Metadata.Ref())

					// create an id token
					tokenToSign := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
						"foo": "bar",
						"aud": "test-clientid",
						"sub": "user",
						"iss": "http://localhost:5556",
					})
					tokenToSign.Header["kid"] = "test-123"
					token, err = tokenToSign.SignedString(privateKey)
					Expect(err).NotTo(HaveOccurred())
				})

				JustBeforeEach(func() {
					_, err := testClients.AuthConfigClient.Write(authConfig, clients.WriteOpts{Ctx: ctx})
					Expect(err).NotTo(HaveOccurred())

					_, err = testClients.ProxyClient.Write(proxy, clients.WriteOpts{})
					Expect(err).NotTo(HaveOccurred())

					Eventually(func() (core.Status, error) {
						proxy, err := testClients.ProxyClient.Read(proxy.Metadata.Namespace, proxy.Metadata.Name, clients.ReadOpts{})
						if err != nil {
							return core.Status{}, err
						}

						return proxy.Status, nil
					}, "5s", "0.1s").Should(MatchFields(IgnoreExtras, Fields{
						"Reason": BeEmpty(),
						"State":  Equal(core.Status_Accepted),
					}))
				})

				AfterEach(func() {
					discoveryServer.Stop()
				})

				ExpectHappyPathToWork := func(loginSuccessExpectation func()) {
					// do auth flow and make sure we have a cookie named cookie:
					appPage, err := url.Parse(fmt.Sprintf("http://%s:%d/", "localhost", envoyPort))
					Expect(err).NotTo(HaveOccurred())

					var finalurl *url.URL
					jar, err := cookiejar.New(nil)
					Expect(err).NotTo(HaveOccurred())
					client := &http.Client{
						Jar: &unsecureCookieJar{CookieJar: jar},
						CheckRedirect: func(req *http.Request, via []*http.Request) error {
							finalurl = req.URL
							return nil
						},
					}

					Eventually(func() (http.Response, error) {
						req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/success?foo=bar", "localhost", envoyPort), nil)
						Expect(err).NotTo(HaveOccurred())
						r, err := client.Do(req)
						if err != nil {
							return http.Response{}, err
						}
						return *r, err
					}, "5s", "0.5s").Should(MatchFields(IgnoreExtras, Fields{
						"StatusCode": Equal(http.StatusOK),
					}))

					Expect(finalurl).NotTo(BeNil())
					Expect(finalurl.Path).To(Equal("/success"))
					// make sure query is passed through as well
					Expect(finalurl.RawQuery).To(Equal("foo=bar"))

					// check the cookie jar
					cookies = jar.Cookies(appPage)
					Expect(cookies).NotTo(BeEmpty())

					// make sure login is successful
					loginSuccessExpectation()

					// try to logout:

					logout := fmt.Sprintf("http://%s:%d/logout", "localhost", envoyPort)
					req, err := http.NewRequest("GET", logout, nil)
					Expect(err).NotTo(HaveOccurred())
					resp, err := client.Do(req)
					Expect(err).NotTo(HaveOccurred())
					Expect(resp.StatusCode).To(Equal(http.StatusOK))

					cookies = jar.Cookies(appPage)
					Expect(cookies).To(BeEmpty())
				}

				Context("redis for session store", func() {

					const (
						redisaddr  = "127.0.0.1"
						redisport  = uint32(6379)
						cookieName = "cookie"
					)
					var (
						redisSession *gexec.Session
					)
					BeforeEach(func() {
						// update the config to use redis
						oauth2.OidcAuthorizationCode.Session = &extauth.UserSession{
							FailOnFetchFailure: true,
							Session: &extauth.UserSession_Redis{
								Redis: &extauth.UserSession_RedisSession{
									Options: &extauth.RedisOptions{
										Host: fmt.Sprintf("%s:%d", redisaddr, redisport),
									},
									KeyPrefix:  "key",
									CookieName: cookieName,
								},
							},
						}

						command := exec.Command(getRedisPath(), "--port", fmt.Sprintf("%d", redisport))
						var err error
						redisSession, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
						Expect(err).NotTo(HaveOccurred())
						// give redis a chance to start
						Eventually(redisSession.Out, "5s").Should(gbytes.Say("Ready to accept connections"))
					})

					AfterEach(func() {
						redisSession.Kill()
					})

					It("should work", func() {
						ExpectHappyPathToWork(func() {
							Expect(cookies[0].Name).To(Equal(cookieName))
						})
					})
				})
				Context("forward id token", func() {

					BeforeEach(func() {
						// update the config to use redis
						oauth2.OidcAuthorizationCode.Headers = &extauth.HeaderConfiguration{
							IdTokenHeader: "foo",
						}
					})

					It("should work", func() {
						ExpectHappyPathToWork(func() {})

						select {
						case r := <-testUpstream.C:
							Expect(r.Headers.Get("foo")).NotTo(BeEmpty())
						case <-time.After(time.Second):
							Fail("timedout")
						}
					})
				})

				Context("discovery override", func() {

					BeforeEach(func() {
						oauth2.OidcAuthorizationCode.DiscoveryOverride = &extauth.DiscoveryOverride{
							AuthEndpoint: "http://localhost:5556/alternate-auth",
						}
					})

					It("should redirect to different auth endpoint with auth override", func() {
						client := &http.Client{
							CheckRedirect: func(req *http.Request, via []*http.Request) error {
								// stop at the auth point
								if req.Response != nil && req.Response.Header.Get("x-auth") != "" {
									return http.ErrUseLastResponse
								}
								return nil
							},
						}
						// Confirm that the response matches the one set by the /alternate-auth endpoint
						Eventually(func() (string, error) {
							req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/1", "localhost", envoyPort), nil)
							Expect(err).NotTo(HaveOccurred())
							resp, err := client.Do(req)
							if err != nil {
								return "", err
							}
							body, err := ioutil.ReadAll(resp.Body)
							if err != nil {
								return "", err
							}
							fmt.Fprintf(GinkgoWriter, "headers are %v \n", resp.Header)
							return string(body), nil
						}, "5s", "0.5s").Should(Equal("alternate-auth"))
					})
				})

				Context("happy path with default settings (no redis)", func() {
					It("should work", func() {
						ExpectHappyPathToWork(func() {
							Expect(cookies).ToNot(BeEmpty())
							var cookienames []string
							for _, c := range cookies {
								cookienames = append(cookienames, c.Name)
							}
							Expect(cookienames).To(ConsistOf("id_token", "access_token"))
						})
					})
				})

				Context("Oidc tests that don't forward to upstream", func() {
					BeforeEach(func() {
						// drain channel as we dont care about it
						go func(testUpstream v1helpers.TestUpstream) {
							for range testUpstream.C {
							}
						}(*testUpstream)
					})

					It("should redirect to auth page", func() {
						client := &http.Client{
							CheckRedirect: func(req *http.Request, via []*http.Request) error {
								// stop at the auth point
								if req.Response != nil && req.Response.Header.Get("x-auth") != "" {
									return http.ErrUseLastResponse
								}
								return nil
							},
						}
						Eventually(func() (string, error) {
							req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/1", "localhost", envoyPort), nil)
							Expect(err).NotTo(HaveOccurred())
							resp, err := client.Do(req)
							if err != nil {
								return "", err
							}
							body, err := ioutil.ReadAll(resp.Body)
							if err != nil {
								return "", err
							}
							fmt.Fprintf(GinkgoWriter, "headers are %v \n", resp.Header)
							return string(body), nil
						}, "5s", "0.5s").Should(Equal("auth"))
					})

					It("should include email scope in url", func() {
						client := &http.Client{
							CheckRedirect: func(req *http.Request, via []*http.Request) error {
								return http.ErrUseLastResponse
							},
						}
						req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/1", "localhost", envoyPort), nil)
						Expect(err).NotTo(HaveOccurred())

						Eventually(func() (http.Response, error) {
							r, err := client.Do(req)
							if err != nil {
								return http.Response{}, err
							}
							return *r, err
						}, "5s", "0.5s").Should(MatchFields(IgnoreExtras, Fields{
							"StatusCode": Equal(http.StatusFound),
							"Header":     HaveKeyWithValue("Location", ContainElement(ContainSubstring("email"))),
						}))
					})

					It("should exchange token", func() {
						finalpage := fmt.Sprintf("http://%s:%d/success", "localhost", envoyPort)
						client := &http.Client{
							CheckRedirect: func(req *http.Request, via []*http.Request) error {
								return http.ErrUseLastResponse
							},
						}

						st := oidc.NewStateSigner([]byte(settings.ExtAuthSettings.SigningKey))
						signedState, err := st.Sign(finalpage)
						Expect(err).NotTo(HaveOccurred())
						req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/callback?code=1234&state="+string(signedState), "localhost", envoyPort), nil)
						Expect(err).NotTo(HaveOccurred())

						Eventually(func() (http.Response, error) {
							r, err := client.Do(req)
							if err != nil {
								return http.Response{}, err
							}
							return *r, err
						}, "5s", "0.5s").Should(MatchFields(IgnoreExtras, Fields{
							"StatusCode": Equal(http.StatusFound),
							"Header":     HaveKeyWithValue("Location", []string{finalpage}),
						}))
					})

					Context("oidc + opa sanity", func() {
						BeforeEach(func() {
							policy := &gloov1.Artifact{
								Metadata: core.Metadata{
									Name:      "jwt",
									Namespace: "default",
									Labels:    map[string]string{"team": "infrastructure"},
								},
								Data: map[string]string{
									"jwt.rego": `package test
	
				default allow = false
				allow {
					[header, payload, signature] = io.jwt.decode(input.state.jwt)
					payload["foo"] = "not-bar"
				}
				`}}
							modules := []*core.ResourceRef{{Name: policy.Metadata.Name}}

							_, err := testClients.AuthConfigClient.Write(&extauth.AuthConfig{
								Metadata: core.Metadata{
									Name:      getOidcAndOpaExtAuthExtension().GetConfigRef().Name,
									Namespace: getOidcAndOpaExtAuthExtension().GetConfigRef().Namespace,
								},
								Configs: []*extauth.AuthConfig_Config{
									{
										AuthConfig: &extauth.AuthConfig_Config_Oauth2{
											Oauth2: &extauth.OAuth2{
												OauthType: getOidcAuthCodeConfig(envoyPort, secret.Metadata.Ref()),
											},
										},
									},
									{
										AuthConfig: &extauth.AuthConfig_Config_OpaAuth{
											OpaAuth: getOpaConfig(modules),
										},
									},
								},
							}, clients.WriteOpts{Ctx: ctx})
							Expect(err).NotTo(HaveOccurred())

							proxy = getProxyExtAuthOIDCAndOpa(envoyPort, secret.Metadata.Ref(), testUpstream.Upstream.Metadata.Ref(), modules)

							_, err = testClients.ArtifactClient.Write(policy, clients.WriteOpts{})
							Expect(err).ToNot(HaveOccurred())
						})

						It("should NOT allow access", func() {
							EventuallyWithOffset(1, func() (int, error) {
								req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/1", "localhost", envoyPort), nil)
								req.Header.Add("Authorization", "Bearer "+token)

								resp, err := http.DefaultClient.Do(req)
								if err != nil {
									return 0, err
								}
								return resp.StatusCode, nil
							}, "5s", "0.5s").Should(Equal(http.StatusForbidden))

						})

					})
				})

				ExpectUpstreamRequest := func() {
					EventuallyWithOffset(1, func() (int, error) {
						req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/1", "localhost", envoyPort), nil)
						req.Header.Add("Authorization", "Bearer "+token)

						resp, err := http.DefaultClient.Do(req)
						if err != nil {
							return 0, err
						}
						return resp.StatusCode, nil
					}, "5s", "0.5s").Should(Equal(http.StatusOK))

					select {
					case r := <-testUpstream.C:
						ExpectWithOffset(1, r.Headers["X-User-Id"]).To(HaveLen(1))
						ExpectWithOffset(1, r.Headers["X-User-Id"][0]).To(Equal("http://localhost:5556;user"))
					case <-time.After(time.Second):
						Fail("expected a message to be received")
					}
				}

				Context("Oidc tests that do forward to upstream", func() {
					It("should allow access with proper jwt token", func() {
						ExpectUpstreamRequest()
					})
				})

				Context("oidc + opa sanity", func() {
					BeforeEach(func() {
						policy := &gloov1.Artifact{
							Metadata: core.Metadata{
								Name:      "jwt",
								Namespace: "default",
								Labels:    map[string]string{"team": "infrastructure"},
							},
							Data: map[string]string{
								"jwt.rego": `package test

			default allow = false
			allow {
				[header, payload, signature] = io.jwt.decode(input.state.jwt)
				payload["foo"] = "bar"
			}
			`}}
						modules := []*core.ResourceRef{{Name: policy.Metadata.Name, Namespace: policy.Metadata.Namespace}}
						_, err := testClients.AuthConfigClient.Write(&extauth.AuthConfig{
							Metadata: core.Metadata{
								Name:      getOidcAndOpaExtAuthExtension().GetConfigRef().Name,
								Namespace: getOidcAndOpaExtAuthExtension().GetConfigRef().Namespace,
							},
							Configs: []*extauth.AuthConfig_Config{
								{
									AuthConfig: &extauth.AuthConfig_Config_Oauth2{
										Oauth2: &extauth.OAuth2{
											OauthType: getOidcAuthCodeConfig(envoyPort, secret.Metadata.Ref()),
										},
									},
								},
								{
									AuthConfig: &extauth.AuthConfig_Config_OpaAuth{
										OpaAuth: getOpaConfig(modules),
									},
								},
							},
						}, clients.WriteOpts{Ctx: ctx})
						Expect(err).NotTo(HaveOccurred())
						proxy = getProxyExtAuthOIDCAndOpa(envoyPort, secret.Metadata.Ref(), testUpstream.Upstream.Metadata.Ref(), modules)

						_, err = testClients.ArtifactClient.Write(policy, clients.WriteOpts{})
						Expect(err).ToNot(HaveOccurred())
					})
					It("should allow access", func() {
						ExpectUpstreamRequest()
					})
				})

			})

			Context("oauth2 token introspection sanity", func() {
				var (
					proxy  *gloov1.Proxy
					server *test_utils.AuthServer
				)
				BeforeEach(func() {

					server = test_utils.NewAuthServer(fmt.Sprintf(":%d", 5556), &test_utils.AuthEndpoints{
						TokenIntrospectionEndpoint: "/introspection",
						UserInfoEndpoint:           "/userinfo",
					}, sets.NewString("valid-access-token"), map[string]user_info.UserInfo{})
					server.Start()

					_, err := testClients.AuthConfigClient.Write(&extauth.AuthConfig{
						Metadata: core.Metadata{
							Name:      getOauthTokenIntrospectionExtAuthExtension().GetConfigRef().Name,
							Namespace: getOauthTokenIntrospectionExtAuthExtension().GetConfigRef().Namespace,
						},
						Configs: []*extauth.AuthConfig_Config{{
							AuthConfig: &extauth.AuthConfig_Config_Oauth2{
								Oauth2: &extauth.OAuth2{
									OauthType: getOauthTokenIntrospectionConfig(),
								},
							},
						}},
					}, clients.WriteOpts{Ctx: ctx})
					Expect(err).NotTo(HaveOccurred())

					proxy = getProxyExtAuthOauthTokenIntrospection(envoyPort, testUpstream.Upstream.Metadata.Ref())
				})

				JustBeforeEach(func() {
					_, err := testClients.ProxyClient.Write(proxy, clients.WriteOpts{})
					Expect(err).NotTo(HaveOccurred())

					Eventually(func() (core.Status, error) {
						proxy, err := testClients.ProxyClient.Read(proxy.Metadata.Namespace, proxy.Metadata.Name, clients.ReadOpts{})
						if err != nil {
							return core.Status{}, err
						}

						return proxy.Status, nil
					}, "5s", "0.1s").Should(MatchFields(IgnoreExtras, Fields{
						"Reason": BeEmpty(),
						"State":  Equal(core.Status_Accepted),
					}))
				})

				AfterEach(func() {
					server.Stop()
				})

				BeforeEach(func() {
					// drain channel as we dont care about it
					go func(testUpstream v1helpers.TestUpstream) {
						for range testUpstream.C {
						}
					}(*testUpstream)
				})

				It("should accept extauth oauth2 introspection with valid access token", func() {
					Eventually(func() (int, error) {
						getReq, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/1", "localhost", envoyPort), nil)
						Expect(err).ToNot(HaveOccurred())
						getReq.Header.Set("authorization", "Bearer valid-access-token")

						client := &http.Client{
							CheckRedirect: func(req *http.Request, via []*http.Request) error {
								return http.ErrUseLastResponse
							},
						}
						var resp *http.Response
						resp, err = client.Do(getReq)
						if err != nil {
							return 0, err
						}
						return resp.StatusCode, nil
					}, "5s", "0.5s").Should(Equal(http.StatusOK))
				})

				It("should deny extauth oauth2 introspection with invalid access token", func() {
					Eventually(func() (int, error) {
						getReq, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/1", "localhost", envoyPort), nil)
						Expect(err).ToNot(HaveOccurred())
						getReq.Header.Set("authorization", "Bearer invalid-access-token")

						client := &http.Client{
							CheckRedirect: func(req *http.Request, via []*http.Request) error {
								return http.ErrUseLastResponse
							},
						}
						var resp *http.Response
						resp, err = client.Do(getReq)
						if err != nil {
							return 0, err
						}
						return resp.StatusCode, nil
					}, "5s", "0.5s").Should(Equal(http.StatusForbidden))
				})
			})

			Context("api key sanity tests", func() {
				BeforeEach(func() {

					// drain channel as we dont care about it
					go func(testUpstream v1helpers.TestUpstream) {
						for range testUpstream.C {
						}
					}(*testUpstream)

					_, err := testClients.AuthConfigClient.Write(&extauth.AuthConfig{
						Metadata: core.Metadata{
							Name:      getApiKeyExtAuthExtension().GetConfigRef().Name,
							Namespace: getApiKeyExtAuthExtension().GetConfigRef().Namespace,
						},
						Configs: []*extauth.AuthConfig_Config{{
							AuthConfig: &extauth.AuthConfig_Config_ApiKeyAuth{
								ApiKeyAuth: getApiKeyAuthConfig(),
							},
						}},
					}, clients.WriteOpts{Ctx: ctx})
					Expect(err).NotTo(HaveOccurred())

					apiKeySecret1 := &extauth.ApiKeySecret{
						ApiKey: "secretApiKey1",
					}

					secret1 := &gloov1.Secret{
						Metadata: core.Metadata{
							Name:      "secret1",
							Namespace: "default",
						},
						Kind: &gloov1.Secret_ApiKey{
							ApiKey: apiKeySecret1,
						},
					}

					apiKeySecret2 := &extauth.ApiKeySecret{
						ApiKey: "secretApiKey2",
					}

					secret2 := &gloov1.Secret{
						Metadata: core.Metadata{
							Name:      "secret2",
							Namespace: "default",
							Labels:    map[string]string{"team": "infrastructure"},
						},
						Kind: &gloov1.Secret_ApiKey{
							ApiKey: apiKeySecret2,
						},
					}

					_, err = testClients.SecretClient.Write(secret1, clients.WriteOpts{})
					Expect(err).ToNot(HaveOccurred())

					_, err = testClients.SecretClient.Write(secret2, clients.WriteOpts{})
					Expect(err).ToNot(HaveOccurred())

					proxy := getProxyExtAuthApiKeyAuth(envoyPort, testUpstream.Upstream.Metadata.Ref())

					_, err = testClients.ProxyClient.Write(proxy, clients.WriteOpts{})
					Expect(err).NotTo(HaveOccurred())

					Eventually(func() (core.Status, error) {
						proxy, err := testClients.ProxyClient.Read(proxy.Metadata.Namespace, proxy.Metadata.Name, clients.ReadOpts{})
						if err != nil {
							return core.Status{}, err
						}

						return proxy.Status, nil
					}, "5s", "0.1s").Should(MatchFields(IgnoreExtras, Fields{
						"Reason": BeEmpty(),
						"State":  Equal(core.Status_Accepted),
					}))
				})

				It("should deny ext auth envoy without apikey", func() {
					Eventually(func() (int, error) {
						resp, err := http.Get(fmt.Sprintf("http://%s:%d/1", "localhost", envoyPort))
						if err != nil {
							return 0, err
						}
						return resp.StatusCode, nil
					}, "5s", "0.5s").Should(Equal(http.StatusUnauthorized))
				})

				It("should deny ext auth envoy with incorrect apikey", func() {
					Eventually(func() (int, error) {
						req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/1", "localhost", envoyPort), nil)
						req.Header.Add("api-key", "badApiKey")
						resp, err := http.DefaultClient.Do(req)

						if err != nil {
							return 0, err
						}
						return resp.StatusCode, nil
					}, "5s", "0.5s").Should(Equal(http.StatusUnauthorized))
				})

				It("should accept ext auth envoy with correct apikey -- secret ref match", func() {
					Eventually(func() (int, error) {
						req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/1", "localhost", envoyPort), nil)
						req.Header.Add("api-key", "secretApiKey1")
						resp, err := http.DefaultClient.Do(req)

						if err != nil {
							return 0, err
						}
						return resp.StatusCode, nil
					}, "5s", "0.5s").Should(Equal(http.StatusOK))
				})

				It("should accept ext auth envoy with correct apikey -- label match", func() {
					Eventually(func() (int, error) {
						req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/1", "localhost", envoyPort), nil)
						req.Header.Add("api-key", "secretApiKey2")
						resp, err := http.DefaultClient.Do(req)

						if err != nil {
							return 0, err
						}
						return resp.StatusCode, nil
					}, "5s", "0.5s").Should(Equal(http.StatusOK))
				})
			})
		})

		Context("using old config format", func() {

			Context("oidc sanity", func() {
				var (
					privateKey      *rsa.PrivateKey
					discoveryServer fakeDiscoveryServer
					secret          *gloov1.Secret
					proxy           *gloov1.Proxy
					token           string
				)
				BeforeEach(func() {
					discoveryServer = fakeDiscoveryServer{}

					privateKey = discoveryServer.Start()

					clientSecret := &extauth.OauthSecret{
						ClientSecret: "test",
					}

					secret = &gloov1.Secret{
						Metadata: core.Metadata{
							Name:      "secret",
							Namespace: "default",
						},
						Kind: &gloov1.Secret_Oauth{
							Oauth: clientSecret,
						},
					}
					_, err := testClients.SecretClient.Write(secret, clients.WriteOpts{})
					Expect(err).NotTo(HaveOccurred())

					_, err = testClients.AuthConfigClient.Write(&extauth.AuthConfig{
						Metadata: core.Metadata{
							Name:      getOidcExtAuthExtension().GetConfigRef().Name,
							Namespace: getOidcExtAuthExtension().GetConfigRef().Namespace,
						},
						Configs: []*extauth.AuthConfig_Config{{
							AuthConfig: &extauth.AuthConfig_Config_Oauth{
								Oauth: getOauthConfig(envoyPort, secret.Metadata.Ref()),
							},
						}},
					}, clients.WriteOpts{Ctx: ctx})
					Expect(err).NotTo(HaveOccurred())

					proxy = getProxyExtAuthOIDC(envoyPort, testUpstream.Upstream.Metadata.Ref())

					// create an id token
					tokenToSign := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
						"foo": "bar",
						"aud": "test-clientid",
						"sub": "user",
						"iss": "http://localhost:5556",
					})
					tokenToSign.Header["kid"] = "test-123"
					token, err = tokenToSign.SignedString(privateKey)
					Expect(err).NotTo(HaveOccurred())
				})

				JustBeforeEach(func() {
					_, err := testClients.ProxyClient.Write(proxy, clients.WriteOpts{})
					Expect(err).NotTo(HaveOccurred())

					Eventually(func() (core.Status, error) {
						proxy, err := testClients.ProxyClient.Read(proxy.Metadata.Namespace, proxy.Metadata.Name, clients.ReadOpts{})
						if err != nil {
							return core.Status{}, err
						}

						return proxy.Status, nil
					}, "5s", "0.1s").Should(MatchFields(IgnoreExtras, Fields{
						"Reason": BeEmpty(),
						"State":  Equal(core.Status_Accepted),
					}))
				})

				AfterEach(func() {
					discoveryServer.Stop()
				})

				Context("Oidc tests that don't forward to upstream", func() {
					BeforeEach(func() {
						// drain channel as we dont care about it
						go func(testUpstream v1helpers.TestUpstream) {
							for range testUpstream.C {
							}
						}(*testUpstream)
					})

					It("should redirect to auth page", func() {
						client := &http.Client{
							CheckRedirect: func(req *http.Request, via []*http.Request) error {
								// stop at the auth point
								if req.Response != nil && req.Response.Header.Get("x-auth") != "" {
									return http.ErrUseLastResponse
								}
								return nil
							},
						}
						Eventually(func() (string, error) {
							req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/1", "localhost", envoyPort), nil)
							Expect(err).NotTo(HaveOccurred())
							resp, err := client.Do(req)
							if err != nil {
								return "", err
							}
							body, err := ioutil.ReadAll(resp.Body)
							if err != nil {
								return "", err
							}
							return string(body), nil
						}, "5s", "0.5s").Should(Equal("auth"))
					})

					It("should include email scope in url", func() {
						client := &http.Client{
							CheckRedirect: func(req *http.Request, via []*http.Request) error {
								return http.ErrUseLastResponse
							},
						}
						req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/1", "localhost", envoyPort), nil)
						Expect(err).NotTo(HaveOccurred())

						Eventually(func() (http.Response, error) {
							r, err := client.Do(req)
							if err != nil {
								return http.Response{}, err
							}
							return *r, err
						}, "5s", "0.5s").Should(MatchFields(IgnoreExtras, Fields{
							"StatusCode": Equal(http.StatusFound),
							"Header":     HaveKeyWithValue("Location", ContainElement(ContainSubstring("email"))),
						}))
					})

					It("should exchange token", func() {
						finalpage := fmt.Sprintf("http://%s:%d/success", "localhost", envoyPort)
						client := &http.Client{
							CheckRedirect: func(req *http.Request, via []*http.Request) error {
								return http.ErrUseLastResponse
							},
						}

						st := oidc.NewStateSigner([]byte(settings.ExtAuthSettings.SigningKey))
						signedState, err := st.Sign(finalpage)
						Expect(err).NotTo(HaveOccurred())
						req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/callback?code=1234&state="+string(signedState), "localhost", envoyPort), nil)
						Expect(err).NotTo(HaveOccurred())

						Eventually(func() (http.Response, error) {
							r, err := client.Do(req)
							if err != nil {
								return http.Response{}, err
							}
							return *r, err
						}, "5s", "0.5s").Should(MatchFields(IgnoreExtras, Fields{
							"StatusCode": Equal(http.StatusFound),
							"Header":     HaveKeyWithValue("Location", []string{finalpage}),
						}))
					})

					Context("oidc + opa sanity", func() {
						BeforeEach(func() {
							policy := &gloov1.Artifact{
								Metadata: core.Metadata{
									Name:      "jwt",
									Namespace: "default",
									Labels:    map[string]string{"team": "infrastructure"},
								},
								Data: map[string]string{
									"jwt.rego": `package test
	
				default allow = false
				allow {
					[header, payload, signature] = io.jwt.decode(input.state.jwt)
					payload["foo"] = "not-bar"
				}
				`}}
							modules := []*core.ResourceRef{{Name: policy.Metadata.Name}}

							_, err := testClients.AuthConfigClient.Write(&extauth.AuthConfig{
								Metadata: core.Metadata{
									Name:      getOidcAndOpaExtAuthExtension().GetConfigRef().Name,
									Namespace: getOidcAndOpaExtAuthExtension().GetConfigRef().Namespace,
								},
								Configs: []*extauth.AuthConfig_Config{
									{
										AuthConfig: &extauth.AuthConfig_Config_Oauth{
											Oauth: getOauthConfig(envoyPort, secret.Metadata.Ref()),
										},
									},
									{
										AuthConfig: &extauth.AuthConfig_Config_OpaAuth{
											OpaAuth: getOpaConfig(modules),
										},
									},
								},
							}, clients.WriteOpts{Ctx: ctx})
							Expect(err).NotTo(HaveOccurred())

							proxy = getProxyExtAuthOIDCAndOpa(envoyPort, secret.Metadata.Ref(), testUpstream.Upstream.Metadata.Ref(), modules)

							_, err = testClients.ArtifactClient.Write(policy, clients.WriteOpts{})
							Expect(err).ToNot(HaveOccurred())
						})

						It("should NOT allow access", func() {
							EventuallyWithOffset(1, func() (int, error) {
								req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/1", "localhost", envoyPort), nil)
								req.Header.Add("Authorization", "Bearer "+token)

								resp, err := http.DefaultClient.Do(req)
								if err != nil {
									return 0, err
								}
								return resp.StatusCode, nil
							}, "5s", "0.5s").Should(Equal(http.StatusForbidden))

						})

					})
				})

				ExpectUpstreamRequest := func() {
					EventuallyWithOffset(1, func() (int, error) {
						req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/1", "localhost", envoyPort), nil)
						req.Header.Add("Authorization", "Bearer "+token)

						resp, err := http.DefaultClient.Do(req)
						if err != nil {
							return 0, err
						}
						return resp.StatusCode, nil
					}, "5s", "0.5s").Should(Equal(http.StatusOK))

					select {
					case r := <-testUpstream.C:
						ExpectWithOffset(1, r.Headers["X-User-Id"]).To(HaveLen(1))
						ExpectWithOffset(1, r.Headers["X-User-Id"][0]).To(Equal("http://localhost:5556;user"))
					case <-time.After(time.Second):
						Fail("expected a message to be received")
					}
				}

				Context("Oidc tests that do forward to upstream", func() {
					It("should allow access with proper jwt token", func() {
						ExpectUpstreamRequest()
					})
				})

				Context("oidc + opa sanity", func() {
					BeforeEach(func() {
						policy := &gloov1.Artifact{
							Metadata: core.Metadata{
								Name:      "jwt",
								Namespace: "default",
								Labels:    map[string]string{"team": "infrastructure"},
							},
							Data: map[string]string{
								"jwt.rego": `package test

			default allow = false
			allow {
				[header, payload, signature] = io.jwt.decode(input.state.jwt)
				payload["foo"] = "bar"
			}
			`}}
						modules := []*core.ResourceRef{{Name: policy.Metadata.Name, Namespace: policy.Metadata.Namespace}}
						_, err := testClients.AuthConfigClient.Write(&extauth.AuthConfig{
							Metadata: core.Metadata{
								Name:      getOidcAndOpaExtAuthExtension().GetConfigRef().Name,
								Namespace: getOidcAndOpaExtAuthExtension().GetConfigRef().Namespace,
							},
							Configs: []*extauth.AuthConfig_Config{
								{
									AuthConfig: &extauth.AuthConfig_Config_Oauth{
										Oauth: getOauthConfig(envoyPort, secret.Metadata.Ref()),
									},
								},
								{
									AuthConfig: &extauth.AuthConfig_Config_OpaAuth{
										OpaAuth: getOpaConfig(modules),
									},
								},
							},
						}, clients.WriteOpts{Ctx: ctx})
						Expect(err).NotTo(HaveOccurred())
						proxy = getProxyExtAuthOIDCAndOpa(envoyPort, secret.Metadata.Ref(), testUpstream.Upstream.Metadata.Ref(), modules)

						_, err = testClients.ArtifactClient.Write(policy, clients.WriteOpts{})
						Expect(err).ToNot(HaveOccurred())
					})
					It("should allow access", func() {
						ExpectUpstreamRequest()
					})
				})

			})

		})

		Context("health checker", func() {

			// NOTE: This test MUST run last, since it runs cancel()
			It("should fail healthcheck immediately on shutdown", func() {
				// Need to create some generic config so that extauth starts passing health checks
				basicConfigSetup()

				// Connects to the extauth service's health check
				conn, err := grpc.Dial("localhost:"+strconv.Itoa(settings.ExtAuthSettings.ServerPort), grpc.WithInsecure())
				Expect(err).To(BeNil())
				defer conn.Close()
				healthCheckClient := grpc_health_v1.NewHealthClient(conn)
				Eventually(func() bool { // Wait for the extauth server to start up
					var header metadata.MD
					resp, err := healthCheckClient.Check(ctx, &grpc_health_v1.HealthCheckRequest{
						Service: settings.ExtAuthSettings.ServiceName,
					}, grpc.Header(&header))
					if err != nil {
						return false
					}

					return resp.Status == grpc_health_v1.HealthCheckResponse_SERVING
				}, "5m", ".1s").Should(BeTrue())

				// Start sending health checking requests continuously
				waitForHealthcheck := make(chan struct{})
				go func(waitForHealthcheck chan struct{}) {
					defer GinkgoRecover()
					Eventually(func() bool {
						ctx = context.Background()
						var header metadata.MD
						healthCheckClient.Check(ctx, &grpc_health_v1.HealthCheckRequest{
							Service: settings.ExtAuthSettings.ServiceName,
						}, grpc.Header(&header))
						return len(header.Get("x-envoy-immediate-health-check-fail")) == 1
					}, "5s", ".1s").Should(BeTrue())
					waitForHealthcheck <- struct{}{}
				}(waitForHealthcheck)

				// Start the health checker first, then cancel
				time.Sleep(200 * time.Millisecond)
				cancel()
				Eventually(waitForHealthcheck).Should(Receive(), "5s", ".1s")
			})
		})
	})

})

var startDiscoveryServerOnce sync.Once
var cachedPrivateKey *rsa.PrivateKey

type fakeDiscoveryServer struct {
	s http.Server
}

func (f *fakeDiscoveryServer) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_ = f.s.Shutdown(ctx)
}

func (f *fakeDiscoveryServer) Start() *rsa.PrivateKey {
	startDiscoveryServerOnce.Do(func() {
		var err error
		cachedPrivateKey, err = rsa.GenerateKey(rand.Reader, 512)
		Expect(err).NotTo(HaveOccurred())
	})

	n := base64.RawURLEncoding.EncodeToString(cachedPrivateKey.N.Bytes())
	e := base64.RawURLEncoding.EncodeToString(big.NewInt(0).SetUint64(uint64(cachedPrivateKey.E)).Bytes())

	tokenToSign := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"foo": "bar",
		"aud": "test-clientid",
		"sub": "user",
		"iss": "http://localhost:5556",
	})
	tokenToSign.Header["kid"] = "test-123"
	token, err := tokenToSign.SignedString(cachedPrivateKey)
	Expect(err).NotTo(HaveOccurred())

	f.s = http.Server{
		Addr: ":5556",
	}

	f.s.Handler = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		defer GinkgoRecover()
		rw.Header().Set("content-type", "application/json")

		switch r.URL.Path {
		case "/auth":
			// redirect back immediately. This simulates a user that's already logged in by the IDP.
			redirect_uri := r.URL.Query().Get("redirect_uri")
			state := r.URL.Query().Get("state")
			u, err := url.Parse(redirect_uri)
			Expect(err).NotTo(HaveOccurred())

			u.RawQuery = "code=1234&state=" + state
			fmt.Fprintf(GinkgoWriter, "redirecting to %s\n", u.String())
			rw.Header().Add("Location", u.String())
			rw.Header().Add("x-auth", "auth")
			rw.WriteHeader(http.StatusFound)

			_, _ = rw.Write([]byte(`auth`))
		case "/alternate-auth":
			// redirect back immediately. This simulates a user that's already logged in by the IDP.
			redirect_uri := r.URL.Query().Get("redirect_uri")
			state := r.URL.Query().Get("state")
			u, err := url.Parse(redirect_uri)
			Expect(err).NotTo(HaveOccurred())

			u.RawQuery = "code=9876&state=" + state
			fmt.Fprintf(GinkgoWriter, "redirecting to %s\n", u.String())
			rw.Header().Add("Location", u.String())
			rw.Header().Add("x-auth", "alternate-auth")
			rw.WriteHeader(http.StatusFound)

			_, _ = rw.Write([]byte(`alternate-auth`))
		case "/.well-known/openid-configuration":
			_, _ = rw.Write([]byte(`
		{
			"issuer": "http://localhost:5556",
			"authorization_endpoint": "http://localhost:5556/auth",
			"token_endpoint": "http://localhost:5556/token",
			"jwks_uri": "http://localhost:5556/keys",
			"response_types_supported": [
			  "code"
			],
			"subject_types_supported": [
			  "public"
			],
			"id_token_signing_alg_values_supported": [
			  "RS256"
			],
			"scopes_supported": [
			  "openid",
			  "email",
			  "profile"
			]
		  }
		`))
		case "/token":
			_, _ = rw.Write([]byte(`
			{
				"access_token": "SlAV32hkKG",
				"token_type": "Bearer",
				"refresh_token": "8xLOxBtZp8",
				"expires_in": 3600,
				"id_token": "` + token + `"
			 }
	`))
		case "/keys":
			_, _ = rw.Write([]byte(`
		{
			"keys": [
			  {
				"use": "sig",
				"kty": "RSA",
				"kid": "test-123",
				"alg": "RS256",
				"n": "` + n + `",
				"e": "` + e + `"
			  }
			]
		  }
		`))
		}
	})

	go func() {
		defer GinkgoRecover()
		err := f.s.ListenAndServe()
		if err != http.ErrServerClosed {
			Expect(err).NotTo(HaveOccurred())
		}
	}()

	return cachedPrivateKey
}

func getOauthTokenIntrospectionConfig() *extauth.OAuth2_AccessTokenValidation {
	return &extauth.OAuth2_AccessTokenValidation{
		AccessTokenValidation: &extauth.AccessTokenValidation{
			ValidationType: &extauth.AccessTokenValidation_IntrospectionUrl{
				IntrospectionUrl: "http://localhost:5556/introspection",
			},
			UserinfoUrl:  "http://localhost:5556/userinfo",
			CacheTimeout: nil,
		},
	}
}

func getOauthTokenIntrospectionExtAuthExtension() *extauth.ExtAuthExtension {
	return &extauth.ExtAuthExtension{
		Spec: &extauth.ExtAuthExtension_ConfigRef{
			ConfigRef: &core.ResourceRef{
				Name:      "oauth-token-introspection",
				Namespace: defaults.GlooSystem,
			},
		},
	}
}

func getProxyExtAuthOauthTokenIntrospection(envoyPort uint32, upstream core.ResourceRef) *gloov1.Proxy {
	return getProxyExtAuth(envoyPort, upstream, getOauthTokenIntrospectionExtAuthExtension())
}

func getOauthConfig(envoyPort uint32, secretRef core.ResourceRef) *extauth.OAuth {
	return &extauth.OAuth{
		ClientId:        "test-clientid",
		ClientSecretRef: &secretRef,
		IssuerUrl:       "http://localhost:5556/",
		AppUrl:          fmt.Sprintf("http://localhost:%d", envoyPort),
		CallbackPath:    "/callback",
		Scopes:          []string{"email"},
	}
}

func getOidcAuthCodeConfig(envoyPort uint32, secretRef core.ResourceRef) *extauth.OAuth2_OidcAuthorizationCode {
	return &extauth.OAuth2_OidcAuthorizationCode{
		OidcAuthorizationCode: &extauth.OidcAuthorizationCode{
			ClientId:        "test-clientid",
			ClientSecretRef: &secretRef,
			IssuerUrl:       "http://localhost:5556/",
			AppUrl:          fmt.Sprintf("http://localhost:%d", envoyPort),
			CallbackPath:    "/callback",
			LogoutPath:      "/logout",
			Scopes:          []string{"email"},
		},
	}
}

func getProxyExtAuthOIDC(envoyPort uint32, upstream core.ResourceRef) *gloov1.Proxy {
	return getProxyExtAuth(envoyPort, upstream, getOidcExtAuthExtension())
}

func getOidcExtAuthExtension() *extauth.ExtAuthExtension {
	return &extauth.ExtAuthExtension{
		Spec: &extauth.ExtAuthExtension_ConfigRef{
			ConfigRef: &core.ResourceRef{
				Name:      "oidc-auth",
				Namespace: defaults.GlooSystem,
			},
		},
	}
}

func getProxyExtAuthOIDCAndOpa(envoyPort uint32, secretRef, upstream core.ResourceRef, modules []*core.ResourceRef) *gloov1.Proxy {
	return getProxyExtAuth(envoyPort, upstream, getOidcAndOpaExtAuthExtension())
}

func getOidcAndOpaExtAuthExtension() *extauth.ExtAuthExtension {
	return &extauth.ExtAuthExtension{
		Spec: &extauth.ExtAuthExtension_ConfigRef{
			ConfigRef: &core.ResourceRef{
				Name:      "oidcand-opa-auth",
				Namespace: defaults.GlooSystem,
			},
		},
	}
}

func getOpaConfig(modules []*core.ResourceRef) *extauth.OpaAuth {
	return &extauth.OpaAuth{
		Modules: modules,
		Query:   "data.test.allow == true",
	}
}

func getProxyExtAuthBasicAuth(envoyPort uint32, upstream core.ResourceRef) *gloov1.Proxy {
	return getProxyExtAuth(envoyPort, upstream, GetBasicAuthExtension())
}

func GetBasicAuthExtension() *extauth.ExtAuthExtension {
	return &extauth.ExtAuthExtension{
		Spec: &extauth.ExtAuthExtension_ConfigRef{
			ConfigRef: &core.ResourceRef{
				Name:      "basic-auth",
				Namespace: defaults.GlooSystem,
			},
		},
	}
}

func getBasicAuthConfig() *extauth.BasicAuth {
	return &extauth.BasicAuth{
		Realm: "gloo",
		Apr: &extauth.BasicAuth_Apr{
			Users: map[string]*extauth.BasicAuth_Apr_SaltedHashedPassword{
				"user": {
					// Password is password
					Salt:           "0adzfifo",
					HashedPassword: "14o4fMw/Pm2L34SvyyA2r.",
				},
			},
		},
	}
}

func getProxyExtAuthApiKeyAuth(envoyPort uint32, upstream core.ResourceRef) *gloov1.Proxy {
	return getProxyExtAuth(envoyPort, upstream, getApiKeyExtAuthExtension())
}

func getApiKeyAuthConfig() *extauth.ApiKeyAuth {
	return &extauth.ApiKeyAuth{
		ApiKeySecretRefs: []*core.ResourceRef{
			{
				Namespace: "default",
				Name:      "secret1",
			},
		},
		LabelSelector: map[string]string{"team": "infrastructure"},
	}
}

func getApiKeyExtAuthExtension() *extauth.ExtAuthExtension {
	return &extauth.ExtAuthExtension{
		Spec: &extauth.ExtAuthExtension_ConfigRef{
			ConfigRef: &core.ResourceRef{
				Name:      "apikey-auth",
				Namespace: defaults.GlooSystem,
			},
		},
	}
}

func getProxyExtAuth(envoyPort uint32, upstream core.ResourceRef, extauthCfg *extauth.ExtAuthExtension) *gloov1.Proxy {
	var vhosts []*gloov1.VirtualHost

	vhost := &gloov1.VirtualHost{
		Name:    "gloo-system.virt1",
		Domains: []string{"*"},
		Options: &gloov1.VirtualHostOptions{
			Extauth: extauthCfg,
		},
		Routes: []*gloov1.Route{{
			Action: &gloov1.Route_RouteAction{
				RouteAction: &gloov1.RouteAction{
					Destination: &gloov1.RouteAction_Single{
						Single: &gloov1.Destination{
							DestinationType: &gloov1.Destination_Upstream{
								Upstream: utils.ResourceRefPtr(upstream),
							},
						},
					},
				},
			},
		}},
	}

	vhosts = append(vhosts, vhost)

	p := &gloov1.Proxy{
		Metadata: core.Metadata{
			Name:      "proxy",
			Namespace: "default",
		},
		Listeners: []*gloov1.Listener{{
			Name:        "listener",
			BindAddress: "0.0.0.0",
			BindPort:    envoyPort,
			ListenerType: &gloov1.Listener_HttpListener{
				HttpListener: &gloov1.HttpListener{
					VirtualHosts: vhosts,
				},
			},
		}},
	}

	return p
}

type unsecureCookieJar struct {
	http.CookieJar
}

func (j *unsecureCookieJar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	for _, c := range cookies {
		// hack to work around go client impl that doesn't consider localhost secure.
		c.Secure = false
	}
	j.CookieJar.SetCookies(u, cookies)
}
