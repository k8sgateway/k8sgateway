package k8sgateway_test

import (
	"context"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/suite"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	"github.com/solo-io/gloo/test/kubernetes/e2e"
	"github.com/solo-io/gloo/test/kubernetes/e2e/features/deployer"
	"github.com/solo-io/gloo/test/kubernetes/testutils/gloogateway"
	"github.com/solo-io/skv2/codegen/util"
)

// TestK8sGateway is the function which executes a series of tests against a given installation
func TestK8sGateway(t *testing.T) {
	RegisterFailHandler(Fail)

	ctx := context.Background()
	testCluster := e2e.NewTestCluster()
	testInstallation := testCluster.RegisterTestInstallation(
		t,
		&gloogateway.Context{
			InstallNamespace:   "k8s-gw-helm-test",
			ValuesManifestFile: filepath.Join(util.MustGetThisDir(), "manifests", "k8s-gateway-helm-test.yaml"),
		},
	)

	// We register the cleanup function _before_ we actually perform the installation.
	// This allows us to uninstall Gloo Gateway, in case the original installation only completed partially
	t.Cleanup(func() {
		if t.Failed() {
			testInstallation.PreFailHandler()
		}

		testInstallation.UninstallGlooGateway(ctx, testInstallation.Actions.Glooctl().NewTestHelperUninstallAction())
		testCluster.UnregisterTestInstallation(testInstallation)
	})

	t.Run("install gateway", func(t *testing.T) {
		testInstallation.InstallGlooGateway(ctx, testInstallation.Actions.Glooctl().NewTestHelperInstallAction())
	})

	t.Run("deployer", func(t *testing.T) {
		suite.Run(t, deployer.NewFeatureSuite(ctx, testInstallation))
	})

	// todo: run RouteOptions
}
