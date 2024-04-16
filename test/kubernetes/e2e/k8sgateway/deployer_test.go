package k8sgateway_test

import (
	"context"
	"path/filepath"

	"github.com/solo-io/gloo/test/kubernetes/e2e/features/deployer"
	"github.com/solo-io/skv2/codegen/util"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/test/kubernetes/e2e"
	"github.com/solo-io/gloo/test/kubernetes/testutils/gloogateway"
)

var _ = Describe("Deployer Test", Ordered, func() {

	// An entire file is meant to capture the behaviors that we want to test for a given installation of Gloo Gateway

	var (
		ctx context.Context

		// testInstallation contains all the metadata/utilities necessary to execute a series of tests
		// against an installation of Gloo Gateway
		testInstallation *e2e.TestInstallation
	)

	BeforeAll(func() {
		ctx = context.Background()

		testInstallation = testSuite.RegisterTestInstallation(
			"k8s-gw-deployer-test",
			&gloogateway.Context{
				InstallNamespace:   "k8s-gw-deployer-test",
				ValuesManifestFile: filepath.Join(util.MustGetThisDir(), "manifests", "k8s-gateway-test-helm.yaml"),
			},
		)

		err := testInstallation.InstallGlooGateway(ctx, testInstallation.Actions.GlooCtl().NewTestHelperInstallAction())
		Expect(err).NotTo(HaveOccurred())
	})

	AfterAll(func() {
		err := testInstallation.UninstallGlooGateway(ctx, testInstallation.Actions.GlooCtl().NewTestHelperUninstallAction())
		Expect(err).NotTo(HaveOccurred())

		testSuite.UnregisterTestInstallation(testInstallation)
	})

	Context("Deployer", func() {

		It("provisions resources appropriately", func() {
			testInstallation.RunTest(ctx, deployer.ProvisionDeploymentAndService)
		})

		It("configures proxies from the GatewayParameters CR", func() {
			testInstallation.RunTest(ctx, deployer.ConfigureProxiesFromGatewayParameters)
		})

	})

})
