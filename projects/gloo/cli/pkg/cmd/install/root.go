package install

import (
	"fmt"

	"github.com/solo-io/gloo/pkg/cliutil/install"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/constants"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/flagutils"
	"github.com/solo-io/go-utils/cliutils"
	"github.com/spf13/cobra"
)

func InstallCmd(opts *options.Options, optionsFunc ...cliutils.OptionsFunc) *cobra.Command {
	cmd := &cobra.Command{
		Use:   constants.INSTALL_COMMAND.Use,
		Short: constants.INSTALL_COMMAND.Short,
		Long:  constants.INSTALL_COMMAND.Long,
	}
	cmd.AddCommand(gatewayCmd(opts))
	cmd.AddCommand(ingressCmd(opts))
	cmd.AddCommand(knativeCmd(opts))
	cliutils.ApplyOptions(cmd, optionsFunc)
	return cmd
}

func UninstallCmd(opts *options.Options, optionsFunc ...cliutils.OptionsFunc) *cobra.Command {
	cmd := &cobra.Command{
		Use:   constants.UNINSTALL_COMMAND.Use,
		Short: constants.UNINSTALL_COMMAND.Short,
		Long:  constants.UNINSTALL_COMMAND.Long,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("Uninstalling Gloo...\n")
			if err := UninstallGloo(opts, &install.CmdKubectl{}); err != nil {
				return err
			}
			fmt.Printf("\nGloo was successfully uninstalled.\n")
			return nil
		},
	}

	pFlags := cmd.PersistentFlags()
	flagutils.AddUninstallFlags(pFlags, &opts.Uninstall)

	cliutils.ApplyOptions(cmd, optionsFunc)
	return cmd
}
