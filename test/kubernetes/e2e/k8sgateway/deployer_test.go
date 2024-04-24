package k8sgateway_test

import (
	"context"
	"path/filepath"

	"github.com/onsi/ginkgo"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/test/kubernetes/e2e"
	"github.com/solo-io/gloo/test/kubernetes/e2e/features/deployer"
	"github.com/solo-io/gloo/test/kubernetes/e2e/features/route_options"
	"github.com/solo-io/gloo/test/kubernetes/testutils/gloogateway"
	"github.com/solo-io/skv2/codegen/util"
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

		testInstallation = testCluster.RegisterTestInstallation(
			&gloogateway.Context{
				InstallNamespace:   "k8s-gw-deployer-test",
				ValuesManifestFile: filepath.Join(util.MustGetThisDir(), "manifests", "k8s-gateway-test-helm.yaml"),
			},
		)

		testInstallation.InstallGlooGateway(NewGomega(ginkgo.Fail), ctx, testInstallation.Actions.Glooctl().NewTestHelperInstallAction())
	})

	AfterAll(func() {
		testInstallation.UninstallGlooGateway(NewGomega(ginkgo.Fail), ctx, testInstallation.Actions.Glooctl().NewTestHelperUninstallAction())

		testCluster.UnregisterTestInstallation(testInstallation)
	})

	Context("Deployer", func() {

		It("provisions resources appropriately", func() {
			testInstallation.RunTest(ctx, deployer.ProvisionDeploymentAndService)
		})

		It("configures proxies from the GatewayParameters CR", func() {
			testInstallation.RunTest(ctx, deployer.ConfigureProxiesFromGatewayParameters)
		})

	})

	Context("RouteOptions", func() {

		It("Apply fault injection using targetRef RouteOption", func() {
			testInstallation.RunTest(ctx, route_options.ConfigureRouteOptionsWithTargetRef)
		})

		It("Apply fault injection using filter extension RouteOption", func() {
			testInstallation.RunTest(ctx, route_options.ConfigureRouteOptionsWithFilterExtenstion)
		})

	})

})
