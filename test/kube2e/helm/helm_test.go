package helm_test

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/ghodss/yaml"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/version"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"
	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	"github.com/solo-io/gloo/test/kube2e"
	"github.com/solo-io/go-utils/testutils/exec"
	"github.com/solo-io/k8s-utils/kubeutils"
	"github.com/solo-io/skv2/codegen/util"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/code-generator/schemagen"
	admission_v1_types "k8s.io/api/admissionregistration/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

var _ = Describe("Kube2e: helm", func() {

	var (
		crdDir   string
		chartUri string

		ctx    context.Context
		cancel context.CancelFunc
	)

	BeforeEach(func() {
		ctx, cancel = context.WithCancel(context.Background())

		crdDir = filepath.Join(util.GetModuleRoot(), "install", "helm", "gloo", "crds")
		chartUri = filepath.Join(testHelper.RootDir, testHelper.TestAssetDir, testHelper.HelmChartName+"-"+testHelper.ChartVersion()+".tgz")
	})

	AfterEach(func() {
		cancel()
	})

	It("uses helm to upgrade to this gloo version without errors", func() {
		if os.Getenv("STRICT_VALIDATION") == "true" {
			Skip("skipping test that only passes with strict validation disabled")
		}

		By("should start with gloo version 1.9.0")
		Expect(GetGlooServerVersion(ctx, testHelper.InstallNamespace)).To(Equal(earliestVersionWithV1CRDs))

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

		//Helm upgrade expects the same values overrides as installs
		valueOverrideFile, cleanupFunc := kube2e.GetHelmValuesOverrideFile()
		defer cleanupFunc()
		// upgrade to the gloo version being tested
		// Using the flag --disable-openapi-validation although helm upgrade works without it everywhere except for CI
		runAndCleanCommand("helm", "upgrade", "--disable-openapi-validation", "gloo", chartUri, "-n", testHelper.InstallNamespace, "--values", valueOverrideFile)
		By("should have upgraded to the gloo version being tested")
		Expect(GetGlooServerVersion(ctx, testHelper.InstallNamespace)).To(Equal(testHelper.ChartVersion()))

		kube2e.GlooctlCheckEventuallyHealthy(1, testHelper, "180s")
	})

	It("uses helm to update the settings without errors", func() {

		By("should start with the settings.invalidConfigPolicy.invalidRouteResponseCode=404")
		client := helpers.MustSettingsClient(ctx)
		settings, err := client.Read(testHelper.InstallNamespace, defaults.SettingsName, clients.ReadOpts{})
		Expect(err).To(BeNil())
		Expect(settings.GetGloo().GetInvalidConfigPolicy().GetInvalidRouteResponseCode()).To(Equal(uint32(404)))

		var args []string
		// following logic handles chartUri for focused test
		// update the settings with `helm upgrade` (without updating the gloo version)
		if chartUri == "" { // hasn't yet upgraded to the chart being tested- use regular gloo/gloo chart
			args = []string{"upgrade", "gloo", "gloo/gloo",
				"-n", testHelper.InstallNamespace,
				"--set", "settings.replaceInvalidRoutes=true",
				"--set", "settings.invalidConfigPolicy.invalidRouteResponseCode=400",
				"--version", GetGlooServerVersion(ctx, testHelper.InstallNamespace)}
		} else { // has already upgraded to the chart being tested- use it
			args = []string{"upgrade", "gloo", chartUri,
				"-n", testHelper.InstallNamespace,
				"--set", "settings.replaceInvalidRoutes=true",
				"--set", "settings.invalidConfigPolicy.invalidRouteResponseCode=400"}
		}
		if os.Getenv("STRICT_VALIDATION") == "true" {
			// in the strict validation tests, make sure we retain the failurePolicy=Fail on upgrades
			args = append(args, "--set", "gateway.validation.failurePolicy=Fail")
		}
		runAndCleanCommand("helm", args...)

		By("should have updated to settings.invalidConfigPolicy.invalidRouteResponseCode=400")
		settings, err = client.Read(testHelper.InstallNamespace, defaults.SettingsName, clients.ReadOpts{})
		Expect(err).To(BeNil())
		Expect(settings.GetGloo().GetInvalidConfigPolicy().GetInvalidRouteResponseCode()).To(Equal(uint32(400)))

		kube2e.GlooctlCheckEventuallyHealthy(1, testHelper, "90s")
	})

	It("uses helm to update the validationServerGrpcMaxSizeBytes without errors", func() {

		By("should start with the gateway.validation.validationServerGrpcMaxSizeBytes=104857600 (100MiB)")
		client := helpers.MustSettingsClient(ctx)
		settings, err := client.Read(testHelper.InstallNamespace, defaults.SettingsName, clients.ReadOpts{})
		Expect(err).To(BeNil())
		Expect(settings.GetGateway().GetValidation().GetValidationServerGrpcMaxSizeBytes().GetValue()).To(Equal(int32(104857600)))

		var args []string
		// following logic handles chartUri for focused test
		// update the settings with `helm upgrade` (without updating the gloo version)
		if chartUri == "" { // hasn't yet upgraded to the chart being tested- use regular gloo/gloo chart
			args = []string{"upgrade", "gloo", "gloo/gloo",
				"-n", testHelper.InstallNamespace,
				"--set", "gateway.validation.validationServerGrpcMaxSizeBytes=5000000",
				"--version", GetGlooServerVersion(ctx, testHelper.InstallNamespace)}
		} else { // has already upgraded to the chart being tested- use it
			args = []string{"upgrade", "gloo", chartUri,
				"-n", testHelper.InstallNamespace,
				"--set", "gateway.validation.validationServerGrpcMaxSizeBytes=5000000"}
		}
		if os.Getenv("STRICT_VALIDATION") == "true" {
			// in the strict validation tests, make sure we retain the failurePolicy=Fail on upgrades
			args = append(args, "--set", "gateway.validation.failurePolicy=Fail")
		}
		runAndCleanCommand("helm", args...)

		By("should have updated to gateway.validation.validationServerGrpcMaxSizeBytes=5000000 (5MB)")
		settings, err = client.Read(testHelper.InstallNamespace, defaults.SettingsName, clients.ReadOpts{})
		Expect(err).To(BeNil())
		Expect(settings.GetGateway().GetValidation().GetValidationServerGrpcMaxSizeBytes().GetValue()).To(Equal(int32(5000000)))

		kube2e.GlooctlCheckEventuallyHealthy(1, testHelper, "90s")
	})

	Context("validation webhook", func() {
		var kubeClientset kubernetes.Interface

		BeforeEach(func() {
			cfg, err := kubeutils.GetConfig("", "")
			Expect(err).NotTo(HaveOccurred())
			kubeClientset, err = kubernetes.NewForConfig(cfg)
			Expect(err).NotTo(HaveOccurred())
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
			runAndCleanCommand("helm", "upgrade", "gloo", chartUri,
				"-n", testHelper.InstallNamespace)
			kube2e.GlooctlCheckEventuallyHealthy(1, testHelper, "90s")

			By("the webhook caBundle and secret's root ca value should still match after upgrade")
			webhookConfig, err = webhookConfigClient.Get(ctx, "gloo-gateway-validation-webhook-"+testHelper.InstallNamespace, metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())
			secret, err = secretClient.Get(ctx, "gateway-validation-certs", metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())
			Expect(webhookConfig.Webhooks[0].ClientConfig.CABundle).To(Equal(secret.Data[corev1.ServiceAccountRootCAKey]))
		})

		It("can change failurePolicy from Fail to Ignore and back again", func() {
			if os.Getenv("STRICT_VALIDATION") != "true" {
				Skip("skipping test that only passes with strict validation enabled")
			}

			webhookConfigClient := kubeClientset.AdmissionregistrationV1().ValidatingWebhookConfigurations()

			By("should start with gateway.validation.failurePolicy=Fail")
			webhookConfig, err := webhookConfigClient.Get(ctx, "gloo-gateway-validation-webhook-"+testHelper.InstallNamespace, metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())
			Expect(*webhookConfig.Webhooks[0].FailurePolicy).To(Equal(admission_v1_types.Fail))

			// upgrade and change to Ignore
			runAndCleanCommand("helm", "upgrade", "gloo", chartUri,
				"-n", testHelper.InstallNamespace,
				"--set", "gateway.validation.failurePolicy=Ignore")
			kube2e.GlooctlCheckEventuallyHealthy(1, testHelper, "90s")

			By("should have updated to gateway.validation.failurePolicy=Ignore")
			webhookConfig, err = webhookConfigClient.Get(ctx, "gloo-gateway-validation-webhook-"+testHelper.InstallNamespace, metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())
			Expect(*webhookConfig.Webhooks[0].FailurePolicy).To(Equal(admission_v1_types.Ignore))

			// upgrade and change to Fail, and set some arbitrary value on the gateway (which should trigger validation webhook)
			runAndCleanCommand("helm", "upgrade", "gloo", chartUri,
				"-n", testHelper.InstallNamespace,
				"--set", "gateway.validation.failurePolicy=Fail",
				"--set", "gatewayProxies.gatewayProxy.gatewaySettings.useProxyProto=true")
			kube2e.GlooctlCheckEventuallyHealthy(1, testHelper, "90s")

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
				err := exec.RunCommand(testHelper.RootDir, false, "kubectl", "apply", "-f", crdFile)
				Expect(err).NotTo(HaveOccurred(), "should be able to kubectl apply -f %s", crdFile)

				// Ensure the CRD is eventually accepted
				Eventually(func() (string, error) {
					return exec.RunCommandOutput(testHelper.RootDir, false, "kubectl", "get", "crd", crd.GetName())
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
				err = exec.RunCommandInput(string(settingsBytes), testHelper.RootDir, false, "kubectl", "apply", "-f", "-")
				Expect(err).NotTo(HaveOccurred(), "should be able to kubectl apply -f %s", settingsFixtureFile)

				// continue traversing
				return nil
			})
			Expect(err).NotTo(HaveOccurred())
		})

	})
})

func GetGlooServerVersion(ctx context.Context, namespace string) (v string) {
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
