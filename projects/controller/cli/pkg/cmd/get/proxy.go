package get

import (
	"github.com/solo-io/gloo/projects/controller/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/controllerrollerroller/cli/pkg/common"
	"github.com/solo-io/gloo/projects/controllerrollerroller/cli/pkg/constants"
	"github.com/solo-io/gloo/projects/controllerrollerroller/cli/pkg/flagutils"
	"github.com/solo-io/gloo/projects/controllerrollerroller/cli/pkg/printers"
	"github.com/spf13/cobra"
)

func Proxy(opts *options.Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:     constants.PROXY_COMMAND.Use,
		Aliases: constants.PROXY_COMMAND.Aliases,
		Short:   "read a proxy or list proxies in a namespace",
		Long:    "usage: glooctl get proxy",
		RunE: func(cmd *cobra.Command, args []string) error {
			proxyList, err := common.GetProxies(common.GetName(args, opts), opts)
			if err != nil {
				return err
			}
			printers.PrintProxies(proxyList, opts.Top.Output)
			return nil
		},
	}

	flagutils.AddGetProxyFlags(cmd.Flags(), &opts.Get.Proxy)

	return cmd
}
