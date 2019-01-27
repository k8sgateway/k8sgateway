package install

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/pkg/errors"
	"github.com/solo-io/gloo/pkg/version"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/flagutils"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube"
	"github.com/solo-io/solo-kit/pkg/utils/kubeutils"
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/spf13/cobra"
)

// TODO: support configuring install namespace
// requires changing a few places in the yaml as well
const (
	glooUrlTemplate = "https://github.com/solo-io/gloo/releases/download/v%s/gloo.yaml"
)

func KubeCmd(opts *options.Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kube",
		Short: "install Gloo on kubernetes",
		Long:  "requires kubectl to be installed",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := registerSettingsCrd(); err != nil {
				return errors.Wrapf(err, "registering settings crd")
			}
			glooManifestBytes, err := readGlooManifest(opts, glooUrlTemplate)
			if err != nil {
				return errors.Wrapf(err, "reading gloo manifest")
			}
			if opts.Install.DryRun {
				fmt.Printf("%s", glooManifestBytes)
				return nil
			}
			return applyManifest(glooManifestBytes)
		},
	}
	pflags := cmd.PersistentFlags()
	flagutils.AddInstallFlags(pflags, &opts.Install)
	return cmd
}

func readGlooManifest(opts *options.Options, urlTemplate string) ([]byte, error) {
	if opts.Install.File != "" {
		return readManifestFromFile(opts.Install.File)
	}
	if version.Version == version.UndefinedVersion || version.Version == version.DevVersion {
		return nil, errors.Errorf("You must provide a file containing the gloo manifest when running an unreleased version of glooctl.")
	}
	return readGlooManifestFromRelease(version.Version, urlTemplate)
}

func readManifestFromFile(path string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, "Error reading file %s", path)
	}
	return bytes, nil
}

func readGlooManifestFromRelease(version, urlTemplate string) ([]byte, error) {
	return readManifest(version, urlTemplate)
}

func applyManifest(manifest []byte) error {
	kubectl := exec.Command("kubectl", "apply", "-f", "-")
	kubectl.Stdin = bytes.NewBuffer(manifest)
	kubectl.Stdout = os.Stdout
	kubectl.Stderr = os.Stderr
	return kubectl.Run()
}

func registerSettingsCrd() error {
	cfg, err := kubeutils.GetConfig("", os.Getenv("KUBECONFIG"))
	if err != nil {
		return err
	}

	settingsClient, err := gloov1.NewSettingsClient(&factory.KubeResourceClientFactory{
		Crd:         gloov1.SettingsCrd,
		Cfg:         cfg,
		SharedCache: kube.NewKubeCache(),
	})

	return settingsClient.Register()
}
