package tests_test

import (
	"context"
	"github.com/solo-io/gloo/pkg/utils/envutils"
	"github.com/solo-io/gloo/test/kubernetes/e2e"
	. "github.com/solo-io/gloo/test/kubernetes/e2e/tests"
	"github.com/solo-io/gloo/test/kubernetes/testutils/gloogateway"
	"github.com/solo-io/gloo/test/testutils"
	"os"
	"testing"
	"time"
)

// TestK8sGatewayIstioAutoMtls is the function which executes a series of tests against a given installation
func TestK8sGatewayIstioAutoMtls(t *testing.T) {
	ctx := context.Background()
	installNs, nsEnvPredefined := envutils.LookupOrDefault(testutils.InstallNamespace, "automtls-istio-k8s-gw-test")
	testInstallation := e2e.CreateTestInstallation(
		t,
		&gloogateway.Context{
			InstallNamespace:          installNs,
			ProfileValuesManifestFile: EmptyProfilePath,
			ValuesManifestFile:        ManifestPath("istio-automtls-k8s-gateway-test-helm.yaml"),
		},
	)

	testHelper := e2e.MustTestHelper(ctx, testInstallation)
	err := testInstallation.AddIstioctl(ctx)
	if err != nil {
		t.Errorf("failed to add istioctl: %v\n", err)
	}

	// Set the env to the install namespace if it is not already set
	if !nsEnvPredefined {
		os.Setenv(testutils.InstallNamespace, installNs)
	}

	// We register the cleanup function _before_ we actually perform the installation.
	// This allows us to uninstall Gloo Gateway, in case the original installation only completed partially
	t.Cleanup(func() {
		if !nsEnvPredefined {
			os.Unsetenv(testutils.InstallNamespace)
		}
		if t.Failed() {
			testInstallation.PreFailHandler(ctx)

			// Generate istioctl bug report
			testInstallation.CreateIstioBugReport(ctx)
		}

		testInstallation.UninstallGlooGatewayWithTestHelper(ctx, testHelper)

		// Uninstall Istio
		err = testInstallation.UninstallIstio()
		if err != nil {
			t.Errorf("failed to add istioctl: %v\n", err)
		}
	})

	// Install Istio before Gloo Gateway to make sure istiod is present before istio-proxy
	err = testInstallation.InstallMinimalIstio(ctx)
	if err != nil {
		t.Errorf("failed to add istioctl: %v\n", err)
	}

	// Install Gloo Gateway
	// istio proxy and sds are added to gateway and take a little longer to start up
	testInstallation.InstallGlooGatewayWithTestHelper(ctx, testHelper, 10*time.Minute)

	AutomtlsIstioSuiteRunner().Run(ctx, t, testInstallation)
}
