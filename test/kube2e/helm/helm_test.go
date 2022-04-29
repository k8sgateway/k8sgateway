package helm_test

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/ghodss/yaml"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/version"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"
	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	"github.com/solo-io/gloo/test/kube2e"
	exec_utils "github.com/solo-io/go-utils/testutils/exec"
	"github.com/solo-io/k8s-utils/kubeutils"
	"github.com/solo-io/k8s-utils/testutils/helper"
	"github.com/solo-io/skv2/codegen/util"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/code-generator/schemagen"
	admission_v1_types "k8s.io/api/admissionregistration/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

// now that we run CI on a kube 1.22 cluster, we must ensure that we install versions of gloo with v1 CRDs
// Per https://github.com/solo-io/gloo/issues/4543: CRDs were migrated from v1beta1 -> v1 in Gloo 1.9.0
const earliestVersionWithV1CRDs = "1.9.0"

const namespace = defaults.GlooSystem

var _ = Describe("Kube2e: helm", func() {

	var (
		crdDir   string
		chartUri string

		ctx    context.Context
		cancel context.CancelFunc

		testHelper *helper.SoloTestHelper

		// whether to install from a released version of the helm chart
		fromRelease bool
		// whether to set validation webhook's failurePolicy=Fail
		strictValidation bool
	)

	BeforeEach(func() {
		ctx, cancel = context.WithCancel(context.Background())

		cwd, err := os.Getwd()
		Expect(err).NotTo(HaveOccurred())
		testHelper, err = helper.NewSoloTestHelper(func(defaults helper.TestConfig) helper.TestConfig {
			defaults.RootDir = filepath.Join(cwd, "../../..")
			defaults.HelmChartName = "gloo"
			defaults.InstallNamespace = namespace
			defaults.Verbose = true
			return defaults
		})
		Expect(err).NotTo(HaveOccurred())

		crdDir = filepath.Join(util.GetModuleRoot(), "install", "helm", "gloo", "crds")
		chartUri = filepath.Join(testHelper.RootDir, testHelper.TestAssetDir, testHelper.HelmChartName+"-"+testHelper.ChartVersion()+".tgz")

		fromRelease = false
		strictValidation = false
	})

	JustBeforeEach(func() {
		installGloo(testHelper, chartUri, fromRelease, strictValidation)
	})

	AfterEach(func() {
		uninstallGloo(testHelper, ctx, cancel)
	})

	Context("upgrades", func() {
		BeforeEach(func() {
			fromRelease = true
		})

		It("uses helm to upgrade to this gloo version without errors", func() {

			By("should start with gloo version 1.9.0")
			Expect(getGlooServerVersion(ctx, testHelper.InstallNamespace)).To(Equal(earliestVersionWithV1CRDs))

			// CRDs are applied to a cluster when performing a `helm install` operation
			// However, `helm upgrade` intentionally does not apply CRDs (https://helm.sh/docs/topics/charts/#limitations-on-crds)
			// Before performing the upgrade, we must manually apply any CRDs that were introduced since v1.9.0
			type crd struct{ name, file string }
			crdsToManuallyApply := []crd{
				{
					name: "graphqlapis.graphql.gloo.solo.io",
					file: filepath.Join(crdDir, "graphql.gloo.solo.io_v1beta1_GraphQLApi.yaml"),
				},
				{
					name: "httpgateways.gateway.solo.io",
					file: filepath.Join(crdDir, "gateway.solo.io_v1_MatchableHttpGateway.yaml"),
				},
				{
					name: "settings.gloo.solo.io",
					file: filepath.Join(crdDir, "gloo.solo.io_v1_Settings.yaml"),
				},
			}

			for _, crd := range crdsToManuallyApply {
				By(fmt.Sprintf("apply new %s CRD", crd.file))
				crdContents, _ := ioutil.ReadFile(crd.file)
				fmt.Println(string(crdContents))
				// Apply the CRD and ensure it is eventually accepted
				runAndCleanCommand("kubectl", "apply", "-f", crd.file)
				Eventually(func() string {
					return string(runAndCleanCommand("kubectl", "get", "crd", crd.name))
				}, "5s", "1s").Should(ContainSubstring(crd.name))
			}

			// upgrade to the gloo version being tested
			upgradeGloo(testHelper, chartUri, strictValidation, nil)

			By("should have upgraded to the gloo version being tested")
			Expect(getGlooServerVersion(ctx, testHelper.InstallNamespace)).To(Equal(testHelper.ChartVersion()))
		})

		It("uses helm to update the settings without errors", func() {

			By("should start with the settings.invalidConfigPolicy.invalidRouteResponseCode=404")
			client := helpers.MustSettingsClient(ctx)
			settings, err := client.Read(testHelper.InstallNamespace, defaults.SettingsName, clients.ReadOpts{})
			Expect(err).To(BeNil())
			Expect(settings.GetGloo().GetInvalidConfigPolicy().GetInvalidRouteResponseCode()).To(Equal(uint32(404)))

			upgradeGloo(testHelper, chartUri, strictValidation, []string{
				"--set", "settings.replaceInvalidRoutes=true",
				"--set", "settings.invalidConfigPolicy.invalidRouteResponseCode=400",
			})

			By("should have updated to settings.invalidConfigPolicy.invalidRouteResponseCode=400")
			settings, err = client.Read(testHelper.InstallNamespace, defaults.SettingsName, clients.ReadOpts{})
			Expect(err).To(BeNil())
			Expect(settings.GetGloo().GetInvalidConfigPolicy().GetInvalidRouteResponseCode()).To(Equal(uint32(400)))
		})

		It("uses helm to update the validationServerGrpcMaxSizeBytes without errors", func() {

			// this is the default value from the 1.9.0 chart
			By("should start with the gateway.validation.validationServerGrpcMaxSizeBytes=4000000 (4MB)")
			client := helpers.MustSettingsClient(ctx)
			settings, err := client.Read(testHelper.InstallNamespace, defaults.SettingsName, clients.ReadOpts{})
			Expect(err).To(BeNil())
			Expect(settings.GetGateway().GetValidation().GetValidationServerGrpcMaxSizeBytes().GetValue()).To(Equal(int32(4000000)))

			upgradeGloo(testHelper, chartUri, strictValidation, []string{
				"--set", "gateway.validation.validationServerGrpcMaxSizeBytes=5000000",
			})

			By("should have updated to gateway.validation.validationServerGrpcMaxSizeBytes=5000000 (5MB)")
			settings, err = client.Read(testHelper.InstallNamespace, defaults.SettingsName, clients.ReadOpts{})
			Expect(err).To(BeNil())
			Expect(settings.GetGateway().GetValidation().GetValidationServerGrpcMaxSizeBytes().GetValue()).To(Equal(int32(5000000)))
		})
	})

	Context("validation webhook", func() {
		var kubeClientset kubernetes.Interface

		BeforeEach(func() {
			cfg, err := kubeutils.GetConfig("", "")
			Expect(err).NotTo(HaveOccurred())
			kubeClientset, err = kubernetes.NewForConfig(cfg)
			Expect(err).NotTo(HaveOccurred())

			strictValidation = true
		})

		It("sets validation webhook caBundle on install and upgrade", func() {
			webhookConfigClient := kubeClientset.AdmissionregistrationV1().ValidatingWebhookConfigurations()
			secretClient := kubeClientset.CoreV1().Secrets(testHelper.InstallNamespace)

			By("the webhook caBundle should be the same as the secret's root ca value")
			webhookConfig, err := webhookConfigClient.Get(ctx, "gloo-gateway-validation-webhook-"+testHelper.InstallNamespace, metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())
			secret, err := secretClient.Get(ctx, "gateway-validation-certs", metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())
			Expect(webhookConfig.Webhooks[0].ClientConfig.CABundle).To(Equal(secret.Data[corev1.ServiceAccountRootCAKey]))

			// do an upgrade
			upgradeGloo(testHelper, chartUri, strictValidation, nil)

			By("the webhook caBundle and secret's root ca value should still match after upgrade")
			webhookConfig, err = webhookConfigClient.Get(ctx, "gloo-gateway-validation-webhook-"+testHelper.InstallNamespace, metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())
			secret, err = secretClient.Get(ctx, "gateway-validation-certs", metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())
			Expect(webhookConfig.Webhooks[0].ClientConfig.CABundle).To(Equal(secret.Data[corev1.ServiceAccountRootCAKey]))
		})

		It("can change failurePolicy from Fail to Ignore and back again", func() {
			webhookConfigClient := kubeClientset.AdmissionregistrationV1().ValidatingWebhookConfigurations()

			By("should start with gateway.validation.failurePolicy=Fail")
			webhookConfig, err := webhookConfigClient.Get(ctx, "gloo-gateway-validation-webhook-"+testHelper.InstallNamespace, metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())
			Expect(*webhookConfig.Webhooks[0].FailurePolicy).To(Equal(admission_v1_types.Fail))

			// upgrade and change to Ignore
			upgradeGloo(testHelper, chartUri, false, nil)

			By("should have updated to gateway.validation.failurePolicy=Ignore")
			webhookConfig, err = webhookConfigClient.Get(ctx, "gloo-gateway-validation-webhook-"+testHelper.InstallNamespace, metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())
			Expect(*webhookConfig.Webhooks[0].FailurePolicy).To(Equal(admission_v1_types.Ignore))

			// upgrade and change to Fail, and set some arbitrary value on the gateway (which should trigger validation webhook)
			upgradeGloo(testHelper, chartUri, true, []string{
				"--set", "gatewayProxies.gatewayProxy.gatewaySettings.useProxyProto=true",
			})

			By("should have updated to gateway.validation.failurePolicy=Fail")
			webhookConfig, err = webhookConfigClient.Get(ctx, "gloo-gateway-validation-webhook-"+testHelper.InstallNamespace, metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())
			Expect(*webhookConfig.Webhooks[0].FailurePolicy).To(Equal(admission_v1_types.Fail))
		})
	})

	Context("applies all CRD manifests without an error", func() {

		var crdsByFileName = map[string]v1.CustomResourceDefinition{}

		BeforeEach(func() {
			err := filepath.Walk(crdDir, func(crdFile string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.IsDir() {
					return nil
				}

				// Parse the file, and extract the CRD
				crd, err := schemagen.GetCRDFromFile(crdFile)
				if err != nil {
					return err
				}
				crdsByFileName[crdFile] = crd

				// continue traversing
				return nil
			})
			Expect(err).NotTo(HaveOccurred())
		})

		It("works using kubectl apply", func() {
			for crdFile, crd := range crdsByFileName {
				// Apply the CRD
				err := exec_utils.RunCommand(testHelper.RootDir, false, "kubectl", "apply", "-f", crdFile)
				Expect(err).NotTo(HaveOccurred(), "should be able to kubectl apply -f %s", crdFile)

				// Ensure the CRD is eventually accepted
				Eventually(func() (string, error) {
					return exec_utils.RunCommandOutput(testHelper.RootDir, false, "kubectl", "get", "crd", crd.GetName())
				}, "10s", "1s").Should(ContainSubstring(crd.GetName()))
			}
		})
	})

	Context("applies settings manifests used in helm unit tests (install/test/fixtures/settings)", func() {
		// The local helm tests involve templating settings with various values set
		// and then validating that the templated data matches fixture data.
		// The tests assume that the fixture data we have defined is valid yaml that
		// will be accepted by a cluster. However, this has not always been the case
		// and it's important that we validate the settings end to end
		//
		// This solution may not be the best way to validate settings, but it
		// attempts to avoid re-running all the helm template tests against a live cluster
		var settingsFixturesFolder string

		BeforeEach(func() {
			settingsFixturesFolder = filepath.Join(util.GetModuleRoot(), "install", "test", "fixtures", "settings")

			// Apply the Settings CRD to ensure it is the most up to date version
			// this ensures that any new fields that have been added are included in the CRD validation schemas
			settingsCrdFilePath := filepath.Join(crdDir, "gloo.solo.io_v1_Settings.yaml")
			runAndCleanCommand("kubectl", "apply", "-f", settingsCrdFilePath)
		})

		It("works using kubectl apply", func() {
			err := filepath.Walk(settingsFixturesFolder, func(settingsFixtureFile string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.IsDir() {
					return nil
				}

				templatedSettings := makeUnstructuredFromTemplateFile(settingsFixtureFile, namespace)
				settingsBytes, err := templatedSettings.MarshalJSON()

				// Apply the fixture
				err = exec_utils.RunCommandInput(string(settingsBytes), testHelper.RootDir, false, "kubectl", "apply", "-f", "-")
				Expect(err).NotTo(HaveOccurred(), "should be able to kubectl apply -f %s", settingsFixtureFile)

				// continue traversing
				return nil
			})
			Expect(err).NotTo(HaveOccurred())
		})

	})
})

func getGlooServerVersion(ctx context.Context, namespace string) (v string) {
	glooVersion, err := version.GetClientServerVersions(ctx, version.NewKube(namespace))
	Expect(err).To(BeNil())
	Expect(len(glooVersion.GetServer())).To(Equal(1))
	for _, container := range glooVersion.GetServer()[0].GetKubernetes().GetContainers() {
		if v == "" {
			v = container.Tag
		} else {
			Expect(container.Tag).To(Equal(v))
		}
	}
	return v
}

func makeUnstructured(yam string) *unstructured.Unstructured {
	jsn, err := yaml.YAMLToJSON([]byte(yam))
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	runtimeObj, err := runtime.Decode(unstructured.UnstructuredJSONScheme, jsn)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	return runtimeObj.(*unstructured.Unstructured)
}

func makeUnstructuredFromTemplateFile(fixtureName string, values interface{}) *unstructured.Unstructured {
	tmpl, err := template.ParseFiles(fixtureName)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	var b bytes.Buffer
	err = tmpl.Execute(&b, values)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	return makeUnstructured(b.String())
}

func installGloo(testHelper *helper.SoloTestHelper, chartUri string, fromRelease bool, strictValidation bool) {
	valueOverrideFile, cleanupFunc := kube2e.GetHelmValuesOverrideFile()
	defer cleanupFunc()

	// construct helm args
	var args = []string{"install", testHelper.HelmChartName}
	if fromRelease {
		args = append(args, "gloo/gloo",
			"--version", fmt.Sprintf("v%s", earliestVersionWithV1CRDs))
	} else {
		args = append(args, chartUri)
	}
	args = append(args, "-n", testHelper.InstallNamespace,
		"--create-namespace",
		"--values", valueOverrideFile)
	if strictValidation {
		args = append(args, strictValidationArgs...)
	}

	if fromRelease {
		runAndCleanCommand("helm", "repo", "add", testHelper.HelmChartName, "https://storage.googleapis.com/solo-public-helm",
			"--force-update")
		runAndCleanCommand("helm", "repo", "update")
	}
	fmt.Printf("running helm with args: %v\n", args)
	runAndCleanCommand("helm", args...)

	// Check that everything is OK
	kube2e.GlooctlCheckEventuallyHealthy(1, testHelper, "90s")
}

func upgradeGloo(testHelper *helper.SoloTestHelper, chartUri string, strictValidation bool, additionalArgs []string) {
	valueOverrideFile, cleanupFunc := kube2e.GetHelmValuesOverrideFile()
	defer cleanupFunc()

	var args = []string{"upgrade", testHelper.HelmChartName, chartUri,
		"-n", testHelper.InstallNamespace,
		"--values", valueOverrideFile}
	if strictValidation {
		args = append(args, strictValidationArgs...)
	}
	args = append(args, additionalArgs...)

	fmt.Printf("running helm with args: %v\n", args)
	runAndCleanCommand("helm", args...)

	// Check that everything is OK
	kube2e.GlooctlCheckEventuallyHealthy(1, testHelper, "90s")
}

func uninstallGloo(testHelper *helper.SoloTestHelper, ctx context.Context, cancel context.CancelFunc) {
	Expect(testHelper).ToNot(BeNil())
	err := testHelper.UninstallGloo()
	Expect(err).NotTo(HaveOccurred())
	_, err = kube2e.MustKubeClient().CoreV1().Namespaces().Get(ctx, testHelper.InstallNamespace, metav1.GetOptions{})
	Expect(apierrors.IsNotFound(err)).To(BeTrue())
	cancel()
}

var strictValidationArgs = []string{
	"--set", "gateway.validation.failurePolicy=Fail",
	"--set", "gateway.validation.allowWarnings=false",
	"--set", "gateway.validation.alwaysAcceptResources=false",
}

func runAndCleanCommand(name string, arg ...string) []byte {
	cmd := exec.Command(name, arg...)
	b, err := cmd.Output()
	// for debugging in Cloud Build
	if err != nil {
		if v, ok := err.(*exec.ExitError); ok {
			fmt.Println("ExitError: ", string(v.Stderr))
		}
	}
	Expect(err).To(BeNil())
	cmd.Process.Kill()
	cmd.Process.Release()
	return b
}
