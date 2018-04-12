package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/solo-io/gloo/internal/control-plane/eventloop"
	"github.com/solo-io/gloo/pkg/bootstrap"
	"github.com/solo-io/gloo/pkg/bootstrap/flags"
	"github.com/solo-io/gloo/pkg/signals"

	//register plugins
	_ "github.com/solo-io/gloo/internal/control-plane/install"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var (
	opts    bootstrap.Options
	xdsPort int
)

var rootCmd = &cobra.Command{
	Use:   "gloo",
	Short: "runs the gloo control plane to manage Envoy as a Function Gateway",
	RunE: func(cmd *cobra.Command, args []string) error {
		stop := signals.SetupSignalHandler()
		eventLoop, err := eventloop.Setup(opts, xdsPort, stop)
		if err != nil {
			return errors.Wrap(err, "setting up event loop")
		}
		if err := eventLoop.Run(stop); err != nil {
			return errors.Wrap(err, "failed running event loop")
		}
		return nil
	},
}

func init() {
	flags.AddConfigStorageOptionFlags(rootCmd, &opts)
	flags.AddSecretStorageOptionFlags(rootCmd, &opts)
	flags.AddFileStorageOptionFlags(rootCmd, &opts)
	flags.AddFileFlags(rootCmd, &opts)
	flags.AddKubernetesFlags(rootCmd, &opts)
	flags.AddConsulFlags(rootCmd, &opts)
	flags.AddVaultFlags(rootCmd, &opts)

	// xds port
	rootCmd.PersistentFlags().IntVar(&xdsPort, "xds.port", 8081, "port to serve envoy xDS services. this port should be specified in your envoy's static config")
}
