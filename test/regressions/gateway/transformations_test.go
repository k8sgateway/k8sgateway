package gateway_test

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/solo-io/solo-projects/test/services"

	"github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/transformers/xslt"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/core/matchers"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/dlp"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/transformation"
	"github.com/solo-io/go-utils/testutils"
	envoy_type "github.com/solo-io/solo-kit/pkg/api/external/envoy/type"
	"github.com/solo-io/solo-projects/test/e2e/transformation_helpers"

	v2 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-projects/test/regressions"

	"github.com/solo-io/k8s-utils/testutils/helper"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gateway/pkg/defaults"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/k8s-utils/kubeutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube"
	"k8s.io/client-go/rest"
)

var _ = Describe("dlp tests", func() {

	var (
		ctx    context.Context
		cancel context.CancelFunc
		cfg    *rest.Config

		cache                kube.SharedCache
		gatewayClient        v2.GatewayClient
		virtualServiceClient v1.VirtualServiceClient

		httpEcho helper.TestRunner
	)

	BeforeEach(func() {
		ctx, cancel = context.WithCancel(context.Background())

		var err error
		cfg, err = kubeutils.GetConfig("", "")
		Expect(err).NotTo(HaveOccurred())

		cache = kube.NewKubeCache(ctx)
		gatewayClientFactory := &factory.KubeResourceClientFactory{
			Crd:         v2.GatewayCrd,
			Cfg:         cfg,
			SharedCache: cache,
		}
		virtualServiceClientFactory := &factory.KubeResourceClientFactory{
			Crd:         v1.VirtualServiceCrd,
			Cfg:         cfg,
			SharedCache: cache,
		}
		gatewayClient, err = v2.NewGatewayClient(ctx, gatewayClientFactory)
		Expect(err).NotTo(HaveOccurred())

		virtualServiceClient, err = v1.NewVirtualServiceClient(ctx, virtualServiceClientFactory)
		Expect(err).NotTo(HaveOccurred())

		httpEcho, err = helper.NewEchoHttp(testHelper.InstallNamespace)
		Expect(err).NotTo(HaveOccurred())

		err = httpEcho.Deploy(2 * time.Minute)
		Expect(err).NotTo(HaveOccurred())

		// bounce envoy, get a clean state (draining listener can break this test). see https://github.com/solo-io/solo-projects/issues/2921 for more.
		out, err := services.KubectlOut(strings.Split("rollout restart -n "+testHelper.InstallNamespace+" deploy/gateway-proxy", " ")...)
		fmt.Println(out)
		Expect(err).ToNot(HaveOccurred())
		out, err = services.KubectlOut(strings.Split("rollout status -n "+testHelper.InstallNamespace+" deploy/gateway-proxy", " ")...)
		fmt.Println(out)
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		regressions.DeleteVirtualService(virtualServiceClient, testHelper.InstallNamespace, "vs", clients.DeleteOpts{Ctx: ctx, IgnoreNotExist: true})
		err := httpEcho.Terminate()
		Expect(err).NotTo(HaveOccurred())
		// Delete http echo service
		err = testutils.Kubectl("delete", "service", "-n", testHelper.InstallNamespace, helper.HttpEchoName, "--grace-period=0")
		Expect(err).NotTo(HaveOccurred())
		cancel()
	})

	waitForGateway := func() {
		defaultGateway := defaults.DefaultGateway(testHelper.InstallNamespace)
		// wait for default gateway to be created
		EventuallyWithOffset(2, func() (*v2.Gateway, error) {
			return gatewayClient.Read(testHelper.InstallNamespace, defaultGateway.Metadata.Name, clients.ReadOpts{})
		}, "15s", "0.5s").Should(Not(BeNil()))
	}

	checkConnection := func(body string) {
		waitForGateway()

		gatewayPort := 80
		testHelper.CurlEventuallyShouldRespond(helper.CurlOpts{
			Protocol:          "http",
			Path:              regressions.TestMatcherPrefix,
			Method:            "GET",
			Headers:           map[string]string{"hello": "world"},
			Host:              defaults.GatewayProxyName,
			Service:           defaults.GatewayProxyName,
			Port:              gatewayPort,
			ConnectionTimeout: 10, // this is important, as the first curl call sometimes hangs indefinitely
			Verbose:           true,
		}, body, 1, time.Minute*5)
	}

	Context("data loss prevention", func() {

		It("will mask regex matches", func() {

			dlpVhost := &dlp.Config{
				Actions: []*dlp.Action{
					{
						ActionType: dlp.Action_CUSTOM,
						CustomAction: &dlp.CustomAction{
							Name:     "test",
							Regex:    []string{"hello", "world"},
							MaskChar: "Y",
							Percent: &envoy_type.Percent{
								Value: 60,
							},
						},
					},
				},
			}

			virtualHostPlugins := &gloov1.VirtualHostOptions{
				Dlp: dlpVhost,
			}

			httpEchoRef := &core.ResourceRef{
				Namespace: testHelper.InstallNamespace,
				Name:      fmt.Sprintf("%s-%s-%v", testHelper.InstallNamespace, helper.HttpEchoName, helper.HttpEchoPort),
			}
			regressions.WriteCustomVirtualService(ctx, 1, testHelper, virtualServiceClient, virtualHostPlugins, nil, nil, httpEchoRef, regressions.TestMatcherPrefix)
			checkConnection(`"YYYlo":"YYYld"`)
		})
	})

	Context("xslt transformer", func() {
		expectBody := func(body, expectedBody string) {
			waitForGateway()
			expectedString := regexp.MustCompile("[\\r\\n\\s]+").ReplaceAllString(expectedBody, "")
			gatewayPort := 80
			testHelper.CurlEventuallyShouldRespond(helper.CurlOpts{
				Protocol:          "http",
				Path:              regressions.TestMatcherPrefix,
				Headers:           map[string]string{"hello": "world"},
				Host:              defaults.GatewayProxyName,
				Service:           defaults.GatewayProxyName,
				Port:              gatewayPort,
				ConnectionTimeout: 10, // this is important, as the first curl call sometimes hangs indefinitely
				Verbose:           true,
				Body:              body,
			}, expectedString, 1, time.Second*20)
		}

		It("will transform xml -> json", func() {

			virtualHostPlugins := &gloov1.VirtualHostOptions{
				StagedTransformations: &transformation.TransformationStages{
					Early: &transformation.RequestResponseTransformations{
						RequestTransforms: []*transformation.RequestMatch{
							{
								Matcher: &matchers.Matcher{
									PathSpecifier: &matchers.Matcher_Prefix{Prefix: "/"},
								},
								RequestTransformation: &transformation.Transformation{
									TransformationType: &transformation.Transformation_XsltTransformation{
										XsltTransformation: &xslt.XsltTransformation{
											Xslt: transformation_helpers.XmlToJsonTransform,
										},
									},
								},
							},
						},
					},
				},
			}

			httpEchoRef := &core.ResourceRef{
				Namespace: testHelper.InstallNamespace,
				Name:      fmt.Sprintf("%s-%s-%v", testHelper.InstallNamespace, helper.HttpEchoName, helper.HttpEchoPort),
			}
			regressions.WriteCustomVirtualService(ctx, 1, testHelper, virtualServiceClient, virtualHostPlugins, nil, nil, httpEchoRef, regressions.TestMatcherPrefix)
			expectBody(transformation_helpers.CarsXml, transformation_helpers.CarsJson)
		})
	})
})
