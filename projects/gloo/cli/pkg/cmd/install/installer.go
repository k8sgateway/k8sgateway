package install

import (
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/solo-io/gloo/pkg/cliutil/install"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/constants"
	"github.com/solo-io/go-utils/kubeutils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/helm/pkg/chartutil"
	"k8s.io/helm/pkg/manifest"
	"k8s.io/helm/pkg/proto/hapi/chart"
	"k8s.io/helm/pkg/renderutil"
)

type GlooKubeInstallClient interface {
	KubectlApply(manifest []byte) error
	WaitForCrdsToBeRegistered(crds []string, timeout, interval time.Duration) error
	CheckKnativeInstallation() (isInstalled bool, isOurs bool, err error)
}

type DefaultGlooKubeInstallClient struct{}

func (i *DefaultGlooKubeInstallClient) KubectlApply(manifest []byte) error {
	return install.KubectlApply(manifest)
}

func (i *DefaultGlooKubeInstallClient) WaitForCrdsToBeRegistered(crds []string, timeout, interval time.Duration) error {
	if len(crds) == 0 {
		return nil
	}

	// TODO: think about improving
	// Just pick the last crd in the list an wait for it to be created. It is reasonable to assume that by the time we
	// get to applying the manifest the other ones will be ready as well.
	crdName := crds[len(crds)-1]

	elapsed := time.Duration(0)
	for {
		select {
		case <-time.After(interval):
			if err := install.Kubectl(nil, "get", crdName); err == nil {
				return nil
			}
			elapsed += interval
			if elapsed > timeout {
				return errors.Errorf("failed to confirm knative crd registration after %v", timeout)
			}
		}
	}
}

func (i *DefaultGlooKubeInstallClient) CheckKnativeInstallation() (bool, bool, error) {
	restCfg, err := kubeutils.GetConfig("", "")
	if err != nil {
		return false, false, err
	}
	kube, err := kubernetes.NewForConfig(restCfg)
	if err != nil {
		return false, false, err
	}
	namespaces, err := kube.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		return false, false, err
	}
	for _, ns := range namespaces.Items {
		if ns.Name == constants.KnativeServingNamespace {
			ours := ns.Labels != nil && ns.Labels["app"] == "gloo"
			return true, ours, nil
		}
	}
	return false, false, nil
}

type ManifestInstaller interface {
	InstallManifest(manifest []byte) error
	InstallCrds(crdNames []string, manifest []byte) error
}

type GlooKubeManifestInstaller struct {
	GlooKubeInstallClient GlooKubeInstallClient
}

func (i *GlooKubeManifestInstaller) InstallManifest(manifest []byte) error {
	if install.IsEmptyManifest(string(manifest)) {
		return nil
	}
	if err := i.GlooKubeInstallClient.KubectlApply(manifest); err != nil {
		return errors.Wrapf(err, "running kubectl apply on manifest")
	}
	return nil
}

func (i *GlooKubeManifestInstaller) InstallCrds(crdNames []string, manifest []byte) error {
	if err := i.InstallManifest(manifest); err != nil {
		return err
	}
	if err := i.GlooKubeInstallClient.WaitForCrdsToBeRegistered(crdNames, time.Second*5, time.Millisecond*500); err != nil {
		return errors.Wrapf(err, "waiting for crds to be registered")
	}
	return nil
}

type DryRunManifestInstaller struct{}

func (i *DryRunManifestInstaller) InstallManifest(manifest []byte) error {
	manifestString := string(manifest)
	if install.IsEmptyManifest(manifestString) {
		return nil
	}
	fmt.Printf("%s", manifestString)
	// For safety, print a YAML separator so multiple invocations of this function will produce valid output
	fmt.Println("\n---")
	return nil
}

func (i *DryRunManifestInstaller) InstallCrds(crdNames []string, manifest []byte) error {
	return i.InstallManifest(manifest)
}

type KnativeInstallStatus struct {
	isInstalled bool
	isOurs      bool
}

type GlooStagedInstaller interface {
	DoCrdInstall() error
	DoPreInstall() error
	DoInstall() error
	DoKnativeInstall() error
}

type DefaultGlooStagedInstaller struct {
	chart                *chart.Chart
	values               *chart.Config
	renderOpts           renderutil.Options
	knativeInstallStatus KnativeInstallStatus
	excludeResources     install.ResourceMatcherFunc
	manifestInstaller    ManifestInstaller
	dryRun               bool
}

func NewGlooStagedInstaller(opts *options.Options, spec GlooInstallSpec, client GlooKubeInstallClient) (GlooStagedInstaller, error) {
	if path.Ext(spec.HelmArchiveUri) != ".tgz" && !strings.HasSuffix(spec.HelmArchiveUri, ".tar.gz") {
		return nil, errors.Errorf("unsupported file extension for Helm chart URI: [%s]. Extension must either be .tgz or .tar.gz", spec.HelmArchiveUri)
	}

	chart, err := install.GetHelmArchive(spec.HelmArchiveUri)
	if err != nil {
		return nil, errors.Wrapf(err, "retrieving gloo helm chart archive")
	}

	values, err := install.GetValuesFromFileIncludingExtra(chart, spec.ValueFileName, spec.ExtraValues)
	if err != nil {
		return nil, errors.Wrapf(err, "retrieving value file: %s", spec.ValueFileName)
	}

	// These are the .Release.* variables used during rendering
	renderOpts := renderutil.Options{
		ReleaseOptions: chartutil.ReleaseOptions{
			Namespace: opts.Install.Namespace,
			Name:      spec.ProductName,
		},
	}

	isInstalled, isOurs, err := client.CheckKnativeInstallation()
	if err != nil {
		return nil, err
	}
	knativeInstallStatus := KnativeInstallStatus{
		isInstalled: isInstalled,
		isOurs:      isOurs,
	}

	var manifestInstaller ManifestInstaller
	if opts.Install.DryRun {
		manifestInstaller = &DryRunManifestInstaller{}
	} else {
		manifestInstaller = &GlooKubeManifestInstaller{
			GlooKubeInstallClient: client,
		}
	}

	return &DefaultGlooStagedInstaller{
		chart:                chart,
		values:               values,
		renderOpts:           renderOpts,
		knativeInstallStatus: knativeInstallStatus,
		excludeResources:     spec.ExcludeResources,
		manifestInstaller:    manifestInstaller,
		dryRun:               opts.Install.DryRun,
	}, nil
}

func (i *DefaultGlooStagedInstaller) DoCrdInstall() error {

	// Keep only CRDs and collect the names
	var crdNames []string
	excludeNonCrdsAndCollectCrdNames := func(input []manifest.Manifest) ([]manifest.Manifest, error) {
		manifests, resourceNames, err := install.ExcludeNonCrds(input)
		crdNames = resourceNames
		return manifests, err
	}

	// Render and install CRD manifests
	crdManifestBytes, err := install.RenderChart(i.chart, i.values, i.renderOpts,
		install.ExcludeNotes,
		install.KnativeResourceFilterFunction(i.knativeInstallStatus.isInstalled),
		excludeNonCrdsAndCollectCrdNames,
		install.ExcludeEmptyManifests)
	if err != nil {
		return errors.Wrapf(err, "rendering crd manifests")
	}

	if !i.dryRun {
		fmt.Printf("Installing CRDs...\n")
	}

	return i.manifestInstaller.InstallCrds(crdNames, crdManifestBytes)
}

func (i *DefaultGlooStagedInstaller) DoPreInstall() error {
	// Render and install Gloo manifest
	manifestBytes, err := install.RenderChart(i.chart, i.values, i.renderOpts,
		install.ExcludeNotes,
		install.ExcludeKnative,
		install.IncludeOnlyPreInstall,
		install.ExcludeEmptyManifests,
		install.ExcludeMatchingResources(i.excludeResources))
	if err != nil {
		return err
	}
	if !i.dryRun {
		fmt.Printf("Preparing namespace and other pre-install tasks...\n")
	}
	return i.manifestInstaller.InstallManifest(manifestBytes)
}

func (i *DefaultGlooStagedInstaller) DoInstall() error {
	// Render and install Gloo manifest
	manifestBytes, err := install.RenderChart(i.chart, i.values, i.renderOpts,
		install.ExcludeNotes,
		install.ExcludeKnative,
		install.ExcludePreInstall,
		install.ExcludeCrds,
		install.ExcludeEmptyManifests,
		install.ExcludeMatchingResources(i.excludeResources))
	if err != nil {
		return err
	}
	if !i.dryRun {
		fmt.Printf("Installing...\n")
	}
	return i.manifestInstaller.InstallManifest(manifestBytes)
}

// This is a bit tricky. The manifest is already filtered based on the values file. If the values file includes
// knative stuff, then we may want to do a knative install -- if there isn't an install already, or if there is
// an install and it's ours (i.e. an upgrade)
func (i *DefaultGlooStagedInstaller) DoKnativeInstall() error {
	// Exclude everything but knative non-crds
	manifestBytes, err := install.RenderChart(i.chart, i.values, i.renderOpts,
		install.ExcludeNonKnative,
		install.KnativeResourceFilterFunction(i.knativeInstallStatus.isInstalled && !i.knativeInstallStatus.isOurs),
		install.ExcludeCrds)
	if err != nil {
		return err
	}
	return i.manifestInstaller.InstallManifest(manifestBytes)
}
