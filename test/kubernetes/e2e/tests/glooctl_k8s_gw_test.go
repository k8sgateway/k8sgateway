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

// TestGlooctlK8sGateway is the function which executes a series of glooctl tests against a given installation with
// the k8s gateway controller enabled
func TestGlooctlK8sGateway(t *testing.T) {
	ctx := context.Background()
	testInstallation := e2e.CreateTestInstallation(
		t,
		&gloogateway.Context{
			InstallNamespace:       env.GetOrDefault(testutils.InstallNamespace, "glooctl-k8s-gw-test"),
			ValuesManifestFile:     filepath.Join(util.MustGetThisDir(), "manifests", "glooctl-k8s-gateway-test-helm.yaml"),
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

	GlooctlKubeGatewaySuiteRunner().Run(ctx, t, testInstallation)
}
