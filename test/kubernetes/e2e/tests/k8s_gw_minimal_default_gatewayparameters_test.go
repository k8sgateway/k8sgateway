package tests_test

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/solo-io/gloo/pkg/utils/env"
	"github.com/solo-io/gloo/test/kubernetes/e2e"
	. "github.com/solo-io/gloo/test/kubernetes/e2e/tests"
	"github.com/solo-io/gloo/test/kubernetes/testutils/gloogateway"
	"github.com/solo-io/gloo/test/kubernetes/testutils/helper"
	"github.com/solo-io/gloo/test/testutils"
	"github.com/solo-io/skv2/codegen/util"
)

// TestK8sGatewayMinimalDefaultGatewayParameters is the function which executes a series of tests against a given installation
// which is expected to have all user-facing options set to null in helm values
func TestK8sGatewayMinimalDefaultGatewayParameters(t *testing.T) {
	ctx := context.Background()
	testInstallation := e2e.CreateTestInstallation(
		t,
		&gloogateway.Context{
			InstallNamespace:       env.GetOrDefault(testutils.InstallNamespace, "k8s-gateway-minimal-default-gatewayparameters-test"),
			ValuesManifestFile:     filepath.Join(util.MustGetThisDir(), "manifests", "k8s-gateway-minimal-default-gatewayparameters-test-helm.yaml"),
			ValidationAlwaysAccept: false,
			K8sGatewayEnabled:      true,
		},
	)

	testHelper := e2e.MustTestHelper(ctx, testInstallation)

	// We register the cleanup function _before_ we actually perform the installation.
	// This allows us to uninstall Gloo Gateway, in case the original installation only completed partially
	t.Cleanup(func() {
		if t.Failed() {
			testInstallation.PreFailHandler(ctx)
		}

		testInstallation.UninstallGlooGateway(ctx, func(ctx context.Context) error {
			return testHelper.UninstallGlooAll()
		})
	})

	// Install Gloo Gateway
	testInstallation.InstallGlooGateway(ctx, func(ctx context.Context) error {
		return testHelper.InstallGloo(ctx, 5*time.Minute, helper.WithExtraArgs("--values", testInstallation.Metadata.ValuesManifestFile))
	})

	KubeGatewayMinimalDefaultGatewayParametersSuiteRunner().Run(ctx, t, testInstallation)
}
