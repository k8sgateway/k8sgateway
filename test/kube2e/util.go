package kube2e

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/test/helpers"
	"github.com/solo-io/go-utils/errors"
	"github.com/solo-io/solo-kit/pkg/utils/log"
	"k8s.io/helm/pkg/repo"
)

const (
	GATEWAY = "gateway"
	INGRESS = "ingress"
	KNATIVE = "knative"
)

func AreTestsDisabled() bool {
	if os.Getenv("RUN_KUBE2E_TESTS") != "1" {
		log.Warnf("This test builds and deploys images to dockerhub and kubernetes, " +
			"and is disabled by default. To enable, set RUN_KUBE2E_TESTS=1 in your env.")
		return true
	}
	return false
}

func InstallGloo(deploymentType string) string {

	version, err := GetTestVersion()
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	log.Debugf("gloo test version is: %s", version)

	namespace := version

	err = glooctlInstall(namespace, version, deploymentType)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())

	err = deployTestRunner(namespace, defaultTestRunnerImage, TestRunnerPort)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	log.Debugf("successfully deployed test runner pod to namespace: %s", namespace)

	return namespace
}

// Returns the version identifier for the current build
func GetTestVersion() (string, error) {

	// Find helm index file in test asset directory
	helmIndex, err := repo.LoadIndexFile(filepath.Join(helpers.GlooDir(), "_test", "index.yaml"))
	if err != nil {
		return "", err
	}

	// Read and return version from helm index file
	if chartVersions, ok := helmIndex.Entries["gloo"]; !ok {
		return "", errors.Errorf("gloo chart not found")
	} else if len(chartVersions) == 0 || len(chartVersions) > 1 {
		return "", errors.Errorf("Expected one chart archive, found: %v", len(chartVersions))
	} else {
		return chartVersions[0].Version, nil
	}
}

func glooctlInstall(namespace, version, deploymentType string) error {
	return helpers.RunCommand(true,
		"_output/glooctl-"+runtime.GOOS+"-amd64",
		"install", deploymentType,
		"-n", namespace,
		"-f", strings.Join([]string{"_test/gloo-", version, ".tgz"}, ""),
	)
}

func GlooctlUninstall(namespace string) error {
	err := helpers.RunCommand(true,
		"_output/glooctl-"+runtime.GOOS+"-amd64",
		"uninstall",
		"-n", namespace,
	)
	if err != nil {
		return err
	}
	// ignore error in case it is already deleted.
	helpers.RunCommand(true,
		"kubectl", "delete", "pod",
		"-n", namespace,
		"testrunner", "--grace-period=0")

	EventuallyWithOffset(1, func() error {
		return helpers.RunCommand(false,
			"kubectl",
			"get",
			"namespace",
			namespace,
		)
	}, "60s", "1s").Should(HaveOccurred())
	return nil
}
