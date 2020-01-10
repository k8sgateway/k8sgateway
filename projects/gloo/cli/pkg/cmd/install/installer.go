package install

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"reflect"
	"strings"

	"github.com/solo-io/gloo/install/helm/gloo/generate"

	"github.com/solo-io/gloo/projects/gloo/cli/pkg/constants"

	"github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"

	"github.com/solo-io/gloo/pkg/cliutil/helm"

	"github.com/solo-io/gloo/pkg/cliutil"
	"github.com/solo-io/gloo/pkg/version"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/go-utils/errors"
	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/cli/values"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/release"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"sigs.k8s.io/yaml"
)

var (
	ChartAndReleaseFlagErr = func(chartOverride, versionOverride string) error {
		return errors.Errorf("you may not specify both a chart with -f and a release version with --version. Received: %s and %s", chartOverride, versionOverride)
	}
)

type Installer interface {
	Install(installerConfig *InstallerConfig) error
}

type InstallerConfig struct {
	InstallCliArgs *options.Install
	ExtraValues    map[string]interface{}
	Enterprise     bool
	Verbose        bool
}

func NewInstaller(helmClient HelmClient) Installer {
	client := helpers.MustKubeClient()
	return NewInstallerWithWriter(helmClient, client.CoreV1().Namespaces(), os.Stdout)
}

// visible for testing
func NewInstallerWithWriter(helmClient HelmClient, kubeNsClient v1.NamespaceInterface, outputWriter io.Writer) Installer {
	return &installer{
		helmClient:         helmClient,
		kubeNsClient:       kubeNsClient,
		dryRunOutputWriter: outputWriter,
	}
}

func (i *installer) Install(installerConfig *InstallerConfig) error {
	namespace := installerConfig.InstallCliArgs.Namespace
	releaseName := installerConfig.InstallCliArgs.HelmReleaseName
	if !installerConfig.InstallCliArgs.DryRun {
		if releaseExists, err := i.helmClient.ReleaseExists(namespace, releaseName); err != nil {
			return err
		} else if releaseExists {
			return GlooAlreadyInstalled(namespace)
		}
		if installerConfig.InstallCliArgs.CreateNamespace {
			// Create the namespace if it doesn't exist. Helm3 no longer does this.
			i.createNamespace(namespace)
		}
	}

	preInstallMessage(installerConfig.InstallCliArgs, installerConfig.Enterprise)

	helmInstall, helmEnv, err := i.helmClient.NewInstall(namespace, releaseName, installerConfig.InstallCliArgs.DryRun)
	if err != nil {
		return err
	}

	chartUri, err := getChartUri(installerConfig.InstallCliArgs.HelmChartOverride,
		strings.TrimPrefix(installerConfig.InstallCliArgs.Version, "v"),
		installerConfig.InstallCliArgs.WithUi,
		installerConfig.Enterprise)
	if err != nil {
		return err
	}
	if installerConfig.Verbose {
		fmt.Printf("Looking for chart at %s\n", chartUri)
	}

	chartObj, err := i.helmClient.DownloadChart(chartUri)
	if err != nil {
		return err
	}

	// determine if it's an enterprise chart by checking if has gloo as a dependency
	installerConfig.Enterprise = false
	for _, dependency := range chartObj.Dependencies() {
		if dependency.Metadata.Name == constants.GlooReleaseName {
			installerConfig.Enterprise = true
			break
		}
	}

	setExtraValues(installerConfig)

	// Merge values provided via the '--values' flag
	valueOpts := &values.Options{
		ValueFiles: installerConfig.InstallCliArgs.HelmChartValueFileNames,
	}
	cliValues, err := valueOpts.MergeValues(getter.All(helmEnv))
	if err != nil {
		return err
	}

	// We need this to avoid rendering the CRDs we include in the /templates directory
	// for backwards-compatibility with Helm 2.
	setCrdCreateToFalse(installerConfig)

	// Merge the CLI flag values into the extra values, giving the latter higher precedence.
	// (The first argument to CoalesceTables has higher priority)
	completeValues := chartutil.CoalesceTables(installerConfig.ExtraValues, cliValues)
	if installerConfig.Verbose {
		b, err := json.Marshal(completeValues)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		y, err := yaml.JSONToYAML(b)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		fmt.Printf("Installing the %s chart with the following value overrides:\n%s\n", chartObj.Metadata.Name, string(y))
	}

	rel, err := helmInstall.Run(chartObj, completeValues)
	if err != nil {
		// TODO: verify whether we actually log something there after these changes
		_, _ = fmt.Fprintf(os.Stderr, "\nGloo failed to install! Detailed logs available at %s.\n", cliutil.GetLogsPath())
		return err
	}
	if installerConfig.Verbose {
		fmt.Printf("Successfully ran helm install with release %s\n", releaseName)
	}

	if installerConfig.InstallCliArgs.DryRun {
		if err := i.printReleaseManifest(rel); err != nil {
			return err
		}
	}

	postInstallMessage(installerConfig.InstallCliArgs, installerConfig.Enterprise)

	return nil
}

func (i *installer) createNamespace(namespace string) {
	_, err := i.kubeNsClient.Get(namespace, metav1.GetOptions{})
	if apierrors.IsNotFound(err) {
		fmt.Printf("Creating namespace %s... ", namespace)
		if _, err := i.kubeNsClient.Create(&corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: namespace,
			},
		}); err != nil {
			fmt.Printf("\nUnable to create namespace %s. Continuing...\n", namespace)
		} else {
			fmt.Printf("Done.\n")
		}
	} else {
		fmt.Printf("\nUnable to check if namespace %s exists. Continuing...\n", namespace)
	}

}

// if enterprise, nest any gloo helm values under "gloo" heading
func setExtraValues(config *InstallerConfig) error {
	if config.ExtraValues == nil || !config.Enterprise {
		return nil
	}

	var glooHelmConfigEmpty generate.HelmConfig
	for k, v := range config.ExtraValues {

		var glooHelmConfigValue generate.HelmConfig

		// use json as a middleman between map and struct
		valueBytes, err := json.Marshal(map[string]interface{}{k:v})
		if err != nil {
			return err
		}
		err = json.Unmarshal(valueBytes, &glooHelmConfigValue)
		if err != nil {
			return err
		}

		// if the chart with the value isn't the same as the empty one, value is gloo value that needs to be nested
		if !reflect.DeepEqual(glooHelmConfigValue, glooHelmConfigEmpty) {
			delete(config.ExtraValues, k)
			if _, ok := config.ExtraValues[constants.GlooReleaseName]; !ok {
				config.ExtraValues[constants.GlooReleaseName] = map[string]interface{}{}
			}
			config.ExtraValues[constants.GlooReleaseName].(map[string]interface{})[k] = v
		}
	}
	return nil
}

// Note: can be removed if we add {"gloo":{"crds":{"create":false}}} to default enterprise chart
func setCrdCreateToFalse(config *InstallerConfig) {
	if config.ExtraValues == nil {
		config.ExtraValues = map[string]interface{}{}
	}

	mapWithCrdValueToOverride := config.ExtraValues

	// If this is an enterprise install, `crds.create` is nested under the `gloo` field
	if config.Enterprise {
		if _, ok := config.ExtraValues[constants.GlooReleaseName]; !ok {
			config.ExtraValues[constants.GlooReleaseName] = map[string]interface{}{}
		}
		mapWithCrdValueToOverride = config.ExtraValues[constants.GlooReleaseName].(map[string]interface{})
	}

	mapWithCrdValueToOverride["crds"] = map[string]interface{}{
		"create": false,
	}
}

func (i *installer) printReleaseManifest(release *release.Release) error {
	// Print CRDs
	for _, crdFile := range release.Chart.CRDs() {
		_, _ = fmt.Fprintln(i.dryRunOutputWriter, string(crdFile.Data))
		_, _ = fmt.Fprintln(i.dryRunOutputWriter, "---")
	}

	// Print hook resources
	nonCleanupHooks, err := helm.GetNonCleanupHooks(release.Hooks)
	if err != nil {
		return err
	}
	for _, hook := range nonCleanupHooks {
		_, _ = fmt.Fprintln(i.dryRunOutputWriter, hook.Manifest)
		_, _ = fmt.Fprintln(i.dryRunOutputWriter, "---")
	}

	// Print the actual release resources
	_, _ = fmt.Fprintf(i.dryRunOutputWriter, "%s", release.Manifest)

	// For safety, print a YAML separator so multiple invocations of this function will produce valid output
	_, _ = fmt.Fprintln(i.dryRunOutputWriter, "---")
	return nil
}

// The resulting URI can be either a URL or a local file path.
func getChartUri(chartOverride, versionOverride string, withUi, enterprise bool) (string, error) {

	if chartOverride != "" && versionOverride != "" {
		return "", ChartAndReleaseFlagErr(chartOverride, versionOverride)
	}

	var helmChartRepoTemplate, helmChartVersion string
	if enterprise {
		helmChartRepoTemplate = GlooEHelmRepoTemplate
	} else if withUi {
		helmChartRepoTemplate = constants.GlooWithUiHelmRepoTemplate
	} else {
		helmChartRepoTemplate = constants.GlooHelmRepoTemplate
	}

	if versionOverride != "" {
		helmChartVersion = versionOverride
	} else if enterprise || withUi {
		enterpriseVersion, err := version.GetLatestEnterpriseVersion(true)
		if err != nil {
			return "", err
		}
		helmChartVersion = enterpriseVersion
	} else {
		glooOsVersion, err := getDefaultGlooInstallVersion(chartOverride)
		if err != nil {
			return "", err
		}
		helmChartVersion = glooOsVersion
	}

	helmChartArchiveUri := fmt.Sprintf(helmChartRepoTemplate, helmChartVersion)

	if chartOverride != "" {
		helmChartArchiveUri = chartOverride
	}

	if path.Ext(helmChartArchiveUri) != ".tgz" && !strings.HasSuffix(helmChartArchiveUri, ".tar.gz") {
		return "", errors.Errorf("unsupported file extension for Helm chart URI: [%s]. Extension must either be .tgz or .tar.gz", helmChartArchiveUri)
	}
	return helmChartArchiveUri, nil
}

func getDefaultGlooInstallVersion(chartOverride string) (string, error) {
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
		fmt.Println("\nGloo Enterprise was successfully installed!")
	} else {
		fmt.Println("\nGloo was successfully installed!")
	}

}

type installer struct {
	helmClient         HelmClient
	kubeNsClient       v1.NamespaceInterface
	dryRunOutputWriter io.Writer
}
