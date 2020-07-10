package install

import (
	"github.com/rotisserie/eris"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/flagutils"
	"github.com/spf13/cobra"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func gatewayCmd(opts *options.Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "gateway",
		Short:  "install the Gloo Gateway on Kubernetes",
		Long:   "requires kubectl to be installed",
		PreRun: setVerboseMode(opts),

		RunE: func(cmd *cobra.Command, args []string) error {
			helmClient := DefaultHelmClient()
			installer := NewInstaller(helmClient)
			if err := installer.Install(&InstallerConfig{
				InstallCliArgs: &opts.Install,
				Verbose:        opts.Top.Verbose,
			}); err != nil {
				return eris.Wrapf(err, "installing gloo in gateway mode")
			}
			return nil
		},
	}

	cmd.AddCommand(enterpriseCmd(opts))
	pFlags := cmd.PersistentFlags()
	flagutils.AddGlooInstallFlags(pFlags, &opts.Install)

	return cmd
}
