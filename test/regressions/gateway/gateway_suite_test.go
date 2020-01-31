package gateway_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/avast/retry-go"
	"github.com/gogo/protobuf/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	errors "github.com/rotisserie/eris"
	"github.com/solo-io/gloo/pkg/cliutil"
	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/check"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/core/matchers"
	"github.com/solo-io/gloo/test/helpers"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/go-utils/kubeutils"
	"github.com/solo-io/go-utils/testutils"
	"github.com/solo-io/go-utils/testutils/clusterlock"
	"github.com/solo-io/go-utils/testutils/exec"
	"github.com/solo-io/go-utils/testutils/helper"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	skhelpers "github.com/solo-io/solo-kit/test/helpers"
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func TestGateway(t *testing.T) {
	if testutils.AreTestsDisabled() {
		return
	}
	skhelpers.RegisterCommonFailHandlers()
	skhelpers.SetupLog()
	_ = os.Remove(cliutil.GetLogsPath())
	skhelpers.RegisterPreFailHandler(printGlooDebugLogs)
	RunSpecs(t, "Gateway Suite")
}

func printGlooDebugLogs() {
	logs, _ := ioutil.ReadFile(cliutil.GetLogsPath())
	fmt.Println("*** Gloo debug logs ***")
	fmt.Println(string(logs))
	fmt.Println("*** End Gloo debug logs ***")
}

const (
	testMatcherPrefix = "/test"

	ldapAssetDir               = "./../assets/ldap"
	ldapServerConfigDirName    = "ldif"
	ldapServerManifestFilename = "ldap-server-manifest.yaml"
)

var (
	testHelper *helper.SoloTestHelper
	locker     *clusterlock.TestClusterLocker
)

var _ = BeforeSuite(func() {
	cwd, err := os.Getwd()
	Expect(err).NotTo(HaveOccurred())

	testHelper, err = helper.NewSoloTestHelper(func(defaults helper.TestConfig) helper.TestConfig {
		defaults.RootDir = filepath.Join(cwd, "../../..")
		defaults.HelmChartName = "gloo-ee"
		defaults.LicenseKey = "eyJleHAiOjM4Nzk1MTY3ODYsImlhdCI6MTU1NDk0MDM0OCwiayI6IkJ3ZXZQQSJ9.tbJ9I9AUltZ-iMmHBertugI2YIg1Z8Q0v6anRjc66Jo"
		defaults.InstallNamespace = "gateway-test-" + fmt.Sprintf("%d-%d", time.Now().Unix()%10000, GinkgoParallelNode())
		return defaults
	})
	Expect(err).NotTo(HaveOccurred())

	skhelpers.RegisterPreFailHandler(helpers.KubeDumpOnFail(GinkgoWriter, testHelper.InstallNamespace))

	locker, err = clusterlock.NewKubeClusterLocker(mustKubeClient(), clusterlock.Options{})
	Expect(err).NotTo(HaveOccurred())
	Expect(locker.AcquireLock(retry.Attempts(40))).NotTo(HaveOccurred())

	// Install Gloo
	values, cleanup := getHelmOverrides()
	defer cleanup()

	err = testHelper.InstallGloo(helper.GATEWAY, 5*time.Minute, helper.ExtraArgs("--values", values))
	Expect(err).NotTo(HaveOccurred())
	Eventually(func() error {
		opts := &options.Options{
			Metadata: core.Metadata{
				Namespace: testHelper.InstallNamespace,
			},
		}
		ok, err := check.CheckResources(opts)
		if err != nil {
			return errors.Wrapf(err, "unable to run glooctl check")
		}
		if ok {
			return nil
		}
		return errors.New("glooctl check detected a problem with the installation")
	}, 2*time.Minute, "5s").Should(BeNil())

	// Print out the versions of CLI and server components
	glooctlVersionCommand := []string{
		filepath.Join(testHelper.BuildAssetDir, testHelper.GlooctlExecName),
		"version", "-n", testHelper.InstallNamespace}
	output, err := exec.RunCommandOutput(testHelper.RootDir, true, glooctlVersionCommand...)
	Expect(err).NotTo(HaveOccurred())
	fmt.Println(output)

	// TODO(marco): explicitly enable strict validation, this can be removed once we enable validation by default
	// See https://github.com/solo-io/gloo/issues/1374
	enableStrictValidation()

	// This should not interfere with any test that is not LDAP related.
	// If it does, we are doing something wrong
	deployLdapServer(mustKubeClient(), testHelper)

})

var _ = AfterSuite(func() {
	defer locker.ReleaseLock()

	cleanupLdapServer(mustKubeClient())

	err := testHelper.UninstallGloo()
	Expect(err).NotTo(HaveOccurred())

	// glooctl should delete the namespace. we do it again just in case it failed
	// ignore errors
	_ = testutils.Kubectl("delete", "namespace", testHelper.InstallNamespace)

	EventuallyWithOffset(1, func() error {
		return testutils.Kubectl("get", "namespace", testHelper.InstallNamespace)
	}, "60s", "1s").Should(HaveOccurred())
})

func mustKubeClient() kubernetes.Interface {
	restConfig, err := kubeutils.GetConfig("", "")
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	kubeClient, err := kubernetes.NewForConfig(restConfig)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	return kubeClient
}

func getHelmOverrides() (filename string, cleanup func()) {
	values, err := ioutil.TempFile("", "*.yaml")
	Expect(err).NotTo(HaveOccurred())
	_, err = values.Write([]byte(`
gloo:
  rbac:    
    namespaced: true
    nameSuffix: e2e-test-rbac-suffix
settings:
  singleNamespace: true
  create: true
prometheus:
  podSecurityPolicy:
    enabled: true
grafana:
  testFramework:
    enabled: false
global:
  extensions:
    extAuth:
      # we want to deploy extauth as both a standalone deployment (the default) and as a sidecar in the envoy pod, so we can test both
      envoySidecar: true
`))
	Expect(err).NotTo(HaveOccurred())
	err = values.Close()
	Expect(err).NotTo(HaveOccurred())

	return values.Name(), func() {
		_ = os.Remove(values.Name())
	}
}

func enableStrictValidation() {
	// enable strict validation
	// this can be removed once we enable validation by default
	// set projects/gateway/pkg/syncer.AcceptAllResourcesByDefault is set to false
	cfg, err := kubeutils.GetConfig("", "")
	Expect(err).NotTo(HaveOccurred())

	kubeCache := kube.NewKubeCache(context.Background())
	settingsClientFactory := &factory.KubeResourceClientFactory{
		Crd:         gloov1.SettingsCrd,
		Cfg:         cfg,
		SharedCache: kubeCache,
	}

	settingsClient, err := gloov1.NewSettingsClient(settingsClientFactory)
	Expect(err).NotTo(HaveOccurred())

	settings, err := settingsClient.Read(testHelper.InstallNamespace, "default", clients.ReadOpts{})
	Expect(err).NotTo(HaveOccurred())

	Expect(settings.Gateway).NotTo(BeNil())
	Expect(settings.Gateway.Validation).NotTo(BeNil())
	settings.Gateway.Validation.AlwaysAccept = &types.BoolValue{Value: false}

	_, err = settingsClient.Write(settings, clients.WriteOpts{OverwriteExisting: true})
	Expect(err).NotTo(HaveOccurred())
}

func deployLdapServer(kubeClient kubernetes.Interface, soloTestHelper *helper.SoloTestHelper) {

	By("create a config map containing the bootstrap configuration for the LDAP server", func() {
		err := testutils.Kubectl(
			"create", "configmap", "ldap", "-n", soloTestHelper.InstallNamespace, "--from-file", filepath.Join(ldapAssetDir, ldapServerConfigDirName))
		Expect(err).NotTo(HaveOccurred())

		Eventually(func() error {
			_, err := kubeClient.CoreV1().ConfigMaps(soloTestHelper.InstallNamespace).Get("ldap", metav1.GetOptions{})
			return err
		}, "15s", "0.5s").Should(BeNil())
	})

	By("deploy an the LDAP server with the correspondent service", func() {
		err := testutils.Kubectl("apply", "-n", soloTestHelper.InstallNamespace, "-f", filepath.Join(ldapAssetDir, ldapServerManifestFilename))
		Expect(err).NotTo(HaveOccurred())

		Eventually(func() error {
			_, err := kubeClient.CoreV1().Services(soloTestHelper.InstallNamespace).Get("ldap", metav1.GetOptions{})
			return err
		}, "15s", "0.5s").Should(BeNil())

		Eventually(func() error {
			deployment, err := kubeClient.AppsV1().Deployments(soloTestHelper.InstallNamespace).Get("ldap", metav1.GetOptions{})
			if err != nil {
				return err
			}
			if deployment.Status.AvailableReplicas == 0 {
				return errors.New("no available replicas for LDAP server deployment")
			}
			return nil
		}, time.Minute, "0.5s").Should(BeNil())

		// Make sure we can query the LDAP server
		soloTestHelper.CurlEventuallyShouldRespond(helper.CurlOpts{
			Protocol:          "ldap",
			Path:              "/",
			Method:            "GET",
			Service:           fmt.Sprintf("ldap.%s.svc.cluster.local", soloTestHelper.InstallNamespace),
			Port:              389,
			ConnectionTimeout: 3,
			Verbose:           true,
		}, "OpenLDAProotDSE", 1, time.Minute)
	})
}

func cleanupLdapServer(kubeClient kubernetes.Interface) {

	// Delete config map
	// Ignore the error on deletion (as it might have never been created if something went wrong in the suite setup),
	// all we care about is that the config map does not exist
	_ = kubeClient.CoreV1().ConfigMaps(testHelper.InstallNamespace).Delete("ldap", &metav1.DeleteOptions{})
	Eventually(func() bool {
		_, err := kubeClient.CoreV1().ConfigMaps(testHelper.InstallNamespace).Get("ldap", metav1.GetOptions{})
		return isNotFound(err)
	}, "15s", "0.5s").Should(BeTrue())

	// Delete LDAP server deployment and service
	// We ignore the error on the deletion call for the same reason as above
	_ = testutils.Kubectl("delete", "-n", testHelper.InstallNamespace, "-f", filepath.Join(ldapAssetDir, ldapServerManifestFilename))
	Eventually(func() bool {
		_, err := kubeClient.CoreV1().Services(testHelper.InstallNamespace).Get("ldap", metav1.GetOptions{})
		return isNotFound(err)
	}, "15s", "0.5s").Should(BeTrue())
	Eventually(func() bool {
		_, err := kubeClient.AppsV1().Deployments(testHelper.InstallNamespace).Get("ldap", metav1.GetOptions{})
		return isNotFound(err)
	}, "15s", "0.5s").Should(BeTrue())
}

func writeVirtualService(ctx context.Context, vsClient v1.VirtualServiceClient,
	virtualHostPlugins *gloov1.VirtualHostOptions, routePlugins *gloov1.RouteOptions,
	sslConfig *gloov1.SslConfig) {

	upstreamRef := &core.ResourceRef{
		Namespace: testHelper.InstallNamespace,
		Name:      fmt.Sprintf("%s-%s-%v", testHelper.InstallNamespace, "testrunner", helper.TestRunnerPort),
	}
	writeCustomVirtualService(ctx, vsClient, virtualHostPlugins, routePlugins, sslConfig, upstreamRef)
}

func writeCustomVirtualService(ctx context.Context, vsClient v1.VirtualServiceClient,
	virtualHostPlugins *gloov1.VirtualHostOptions, routePlugins *gloov1.RouteOptions,
	sslConfig *gloov1.SslConfig, upstreamRef *core.ResourceRef) {

	if routePlugins.GetPrefixRewrite() == nil {
		if routePlugins == nil {
			routePlugins = &gloov1.RouteOptions{}
		}
		routePlugins.PrefixRewrite = &types.StringValue{
			Value: "/",
		}
	}

	// We wrap this in a eventually because the validating webhook may reject the virtual service if one of the
	// resources the VS depends on is not yet available.
	EventuallyWithOffset(1, func() error {
		_, err := vsClient.Write(&v1.VirtualService{

			Metadata: core.Metadata{
				Name:      "vs",
				Namespace: testHelper.InstallNamespace,
			},
			SslConfig: sslConfig,
			VirtualHost: &v1.VirtualHost{
				Options: virtualHostPlugins,
				Domains: []string{"*"},
				Routes: []*v1.Route{{
					Options: routePlugins,
					Matchers: []*matchers.Matcher{{
						PathSpecifier: &matchers.Matcher_Prefix{
							Prefix: testMatcherPrefix,
						},
					}},
					Action: &v1.Route_RouteAction{
						RouteAction: &gloov1.RouteAction{
							Destination: &gloov1.RouteAction_Single{
								Single: &gloov1.Destination{
									DestinationType: &gloov1.Destination_Upstream{
										Upstream: upstreamRef,
									},
								},
							},
						},
					},
				}},
			},
		}, clients.WriteOpts{Ctx: ctx})

		if err != nil {
			contextutils.LoggerFrom(ctx).Warnw("failed to create virtual service", zap.Error(err))
		}

		return err
	}, time.Minute, "5s").Should(BeNil())
}

func deleteVirtualService(vsClient v1.VirtualServiceClient, ns, name string, opts clients.DeleteOpts) {
	// We wrap this in a eventually because the validating webhook may reject the virtual service if one of the
	// resources the VS depends on is not yet available.
	EventuallyWithOffset(1, func() error {
		return vsClient.Delete(ns, name, opts)
	}, time.Minute, "5s").Should(BeNil())
}
