package flagutils

import (
	"github.com/solo-io/gloo/pkg/version"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/spf13/pflag"
)

func AddInstallFlags(set *pflag.FlagSet, install *options.Install) {
	set.BoolVarP(&install.DryRun, "dry-run", "d", false, "Dump the raw installation yaml instead of applying it to kubernetes")
	if !version.IsReleaseVersion() {
		set.StringVar(&install.ReleaseVersion, "release", "", "install using this release version. defaults to the latest github release")
	}
	set.StringVarP(&install.GlooManifestOverride, "file", "f", "", "Install Gloo from this kubernetes manifest yaml file rather than from a release")
}

func AddKnativeInstallFlags(set *pflag.FlagSet, knative *options.KnativeInstall) {
	set.StringVar(&knative.CrdManifestOverride, "knative-crds-manifest", "", "Install Knative CRDs from this kubernetes manifest yaml file rather than from a release")
	set.StringVar(&knative.InstallManifestOverride, "knative-install-manifest", "", "Install Knative Serving from this kubernetes manifest yaml file rather than from a release")
}
