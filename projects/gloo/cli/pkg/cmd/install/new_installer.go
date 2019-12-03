package install

import (
	"fmt"
	"github.com/solo-io/gloo/pkg/cliutil"
	"github.com/solo-io/gloo/pkg/cliutil/helm"
	"github.com/solo-io/gloo/pkg/version"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/constants"
	"github.com/solo-io/go-utils/errors"
	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/cli/values"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/release"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"os"
	"path"
	"strings"
)

func Install(installOpts *options.Install, extraValues map[string]interface{}, enterprise bool) error {

	if !installOpts.DryRun {
		if releaseExists, err := ReleaseExists(installOpts.Namespace); err != nil {
			return err
		} else if releaseExists {
			// TODO(helm3): improve error message
			return errors.New("Gloo already installed")
		}
	}

	preInstallMessage(installOpts, enterprise)

	helmInstall, helmEnv, err := helm.NewInstall(installOpts.Namespace, constants.GlooReleaseName, installOpts.DryRun)
	if err != nil {
		return err
	}

	chartUri, err := getChartUri(installOpts.HelmChartOverride, installOpts.WithUi, enterprise)
	if err != nil {
		return err
	}

	chartObj, err := helm.DownloadChart(chartUri)
	if err != nil {
		return err
	}

	// Merge values provided via the '--values' flag
	valueOpts := &values.Options{
		ValueFiles: installOpts.HelmChartValueFileNames,
	}
	cliValues, err := valueOpts.MergeValues(getter.All(helmEnv))
	if err != nil {
		return err
	}

	// We need to make sure the namespace is created when installing UI
	if installOpts.WithUi {
		createNsValues := map[string]interface{}{
			"gloo": map[string]interface{}{
				"namespace": map[string]interface{}{
					"create": "true",
				},
			},
		}
		extraValues = chartutil.CoalesceTables(createNsValues, extraValues)
	}

	// Merge the CLI flag values into the extra values, giving the latter higher precedence.
	// (The first argument to CoalesceTables has higher priority)
	completeValues := chartutil.CoalesceTables(extraValues, cliValues)

	rel, err := helmInstall.Run(chartObj, completeValues)
	if err != nil {
		// TODO: verify whether we actually log something there after these changes
		_, _ = fmt.Fprintf(os.Stderr, "\nGloo failed to install! Detailed logs available at %s.\n", cliutil.GetLogsPath())
		return err
	}

	if installOpts.DryRun {
		PrintReleaseManifest(rel)
	}

	postInstallMessage(installOpts, enterprise)

	return nil
}

func ReleaseExists(namespace string) (bool, error) {
	list, err := helm.NewList(namespace)
	if err != nil {
		return false, err
	}
	list.Filter = constants.GlooReleaseName

	releases, err := list.Run()
	if err != nil {
		return false, err
	}

	return len(releases) > 0, nil
}

func PrintReleaseManifest(release *release.Release) {

	// Print CRDs
	for _, crdFile := range release.Chart.CRDs() {
		fmt.Printf("%s", string(crdFile.Data))
		fmt.Println("\n---")
	}

	// TODO: print hooks, filtering out the duplicated hook resources

	// Print the actual release resources
	fmt.Printf("%s", release.Manifest)

	// For safety, print a YAML separator so multiple invocations of this function will produce valid output
	fmt.Println("\n---")
}

// The resulting URI can be either a URL or a local file path.
func getChartUri(chartOverride string, withUi bool, enterprise bool) (string, error) {
	var helmChartArchiveUri string

	if enterprise {
		helmChartArchiveUri = fmt.Sprintf(GlooEHelmRepoTemplate, version.EnterpriseTag)
	} else if withUi {
		helmChartArchiveUri = fmt.Sprintf(constants.GlooWithUiHelmRepoTemplate, version.EnterpriseTag)
	} else {
		glooOsVersion, err := getGlooVersion(chartOverride)
		if err != nil {
			return "", err
		}
		helmChartArchiveUri = fmt.Sprintf(constants.GlooHelmRepoTemplate, glooOsVersion)
	}

	if chartOverride != "" {
		helmChartArchiveUri = chartOverride
	}

	if path.Ext(helmChartArchiveUri) != ".tgz" && !strings.HasSuffix(helmChartArchiveUri, ".tar.gz") {
		return "", errors.Errorf("unsupported file extension for Helm chart URI: [%s]. Extension must either be .tgz or .tar.gz", helmChartArchiveUri)
	}
	return helmChartArchiveUri, nil
}

func getGlooVersion(chartOverride string) (string, error) {
	if !version.IsReleaseVersion() && chartOverride == "" {
		return "", errors.Errorf("you must provide a Gloo Helm chart URI via the 'file' option " +
			"when running an unreleased version of glooctl")
	}
	return version.Version, nil
}

func preInstallMessage(installOpts *options.Install, enterprise bool) {
	if installOpts.DryRun {
		return
	}
	if enterprise {
		fmt.Println("Starting Gloo Enterprise installation...")
	} else {
		fmt.Println("Starting Gloo installation...")
	}
}
func postInstallMessage(installOpts *options.Install, enterprise bool) {
	if installOpts.DryRun {
		return
	}
	if enterprise {
		fmt.Println("Gloo Enterprise was successfully installed!")
	} else {
		fmt.Println("Gloo was successfully installed!")
	}

}
