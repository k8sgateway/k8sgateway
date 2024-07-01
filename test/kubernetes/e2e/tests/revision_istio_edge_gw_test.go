package tests_test

import (
	"context"
	"log"
	"path/filepath"
	"testing"
	"time"

	"github.com/solo-io/gloo/test/kube2e/helper"
	"github.com/solo-io/gloo/test/kubernetes/e2e"
	. "github.com/solo-io/gloo/test/kubernetes/e2e/tests"
	"github.com/solo-io/gloo/test/kubernetes/testutils/gloogateway"

	"github.com/solo-io/skv2/codegen/util"
)

// TestRevisionIstioRegression is the function which executes a series of tests against a given installation where
// the k8s Gateway controller is disabled and the deprecated Istio integration values are used to check for regressions
func TestRevisionIstioRegression(t *testing.T) {
	ctx := context.Background()
	testInstallation := e2e.CreateTestInstallation(
		t,
		&gloogateway.Context{
			InstallNamespace:   "istio-rev-regression-test",
			ValuesManifestFile: filepath.Join(util.MustGetThisDir(), "manifests", "istio-revision-helm.yaml"),
		},
	)

	testHelper := e2e.MustTestHelper(ctx, testInstallation)

	err := testInstallation.AddIstioctl(ctx)
	if err != nil {
		log.Printf("failed to add istioctl: %v\n", err)
		t.Fail()
	}

	// We register the cleanup function _before_ we actually perform the installation.
	// This allows us to uninstall Gloo Gateway, in case the original installation only completed partially
	t.Cleanup(func() {
		if t.Failed() {
			testInstallation.PreFailHandler(ctx)

			// Generate istioctl bug report
			testInstallation.CreateIstioBugReport(ctx)
		}

		testInstallation.UninstallGlooGateway(ctx, func(ctx context.Context) error {
			return testHelper.UninstallGlooAll()
		})

		// Uninstall Istio
		err = testInstallation.UninstallIstio()
		if err != nil {
			log.Printf("failed to uninstall: %v\n", err)
			t.Fail()
		}
	})

	// Install Istio before Gloo Gateway to make sure istiod is present before istio-proxy
	err = testInstallation.InstallRevisionedIstio(ctx)
	if err != nil {
		log.Printf("failed to install: %v\n", err)
		t.Fail()
	}

	// Install Gloo Gateway with only Edge APIs enabled
	testInstallation.InstallGlooGateway(ctx, func(ctx context.Context) error {
		return testHelper.InstallGloo(ctx, helper.GATEWAY, 5*time.Minute, helper.ExtraArgs("--values", testInstallation.Metadata.ValuesManifestFile))
	})

	RevisionIstioEdgeGatewaySuiteRunner().Run(ctx, t, testInstallation)
}
