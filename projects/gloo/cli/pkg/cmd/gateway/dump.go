package gateway

import (
	"archive/zip"
	"fmt"
	"os"
	"time"

	"github.com/solo-io/go-utils/cliutils"

	"github.com/solo-io/gloo/pkg/utils/envoyutils/admincli"

	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	"github.com/spf13/cobra"
)

func dumpCfgCmd(opts *options.Options, optionsFunc ...cliutils.OptionsFunc) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dump",
		Short: "dump Envoy config from one of the proxy instances",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getEnvoyCfgDump(opts)
		},
	}
	cliutils.ApplyOptions(cmd, optionsFunc)
	return cmd
}

func dumpStatsCmd(opts *options.Options, optionsFunc ...cliutils.OptionsFunc) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stats",
		Short: "stats for one of the proxy instances",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getEnvoyStatsDump(opts)
		},
	}
	cliutils.ApplyOptions(cmd, optionsFunc)
	return cmd
}

func writeSnapshotCmd(opts *options.Options, optionsFunc ...cliutils.OptionsFunc) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "snapshot",
		Short: "snapshot complete proxy state for the given instance to an archive",
		RunE: func(cmd *cobra.Command, args []string) error {
			dumpFile, err := getEnvoyFullDumpToDisk(opts)
			if err != nil {
				// If we have an error writing zip (or fetching dump)
				// delete the file after it's flushed to clean up.
				_ = os.Remove(dumpFile)
				return err
			}
			return nil
		},
	}
	cliutils.ApplyOptions(cmd, optionsFunc)
	return cmd
}

func getEnvoyCfgDump(opts *options.Options) error {
	adminCli, shutdownFunc, err := admincli.NewPortForwardedClient(opts.Top.Ctx, opts.Proxy.Name, opts.Metadata.GetNamespace())
	if err != nil {
		return err
	}

	defer shutdownFunc()

	return adminCli.ConfigDumpCmd(opts.Top.Ctx, nil).WithStdout(os.Stdout).Run().Cause()
}

func getEnvoyStatsDump(opts *options.Options) error {
	adminCli, shutdownFunc, err := admincli.NewPortForwardedClient(opts.Top.Ctx, opts.Proxy.Name, opts.Metadata.GetNamespace())
	if err != nil {
		return err
	}

	defer shutdownFunc()

	return adminCli.StatsCmd(opts.Top.Ctx).WithStdout(os.Stdout).Run().Cause()
}

func getEnvoyFullDumpToDisk(opts *options.Options) (string, error) {
	proxyOutArchiveFile, err := createArchiveFile()
	if err != nil {
		return proxyOutArchiveFile.Name(), err
	}
	proxyOutArchive := zip.NewWriter(proxyOutArchiveFile)
	defer proxyOutArchiveFile.Close()
	defer proxyOutArchive.Close()

	proxyNamespace := opts.Metadata.GetNamespace()
	if proxyNamespace == "" {
		proxyNamespace = defaults.GlooSystem
	}

	adminCli, shutdownFunc, err := admincli.NewPortForwardedClient(opts.Top.Ctx, opts.Proxy.Name, opts.Metadata.GetNamespace())
	if err != nil {
		return proxyOutArchiveFile.Name(), err
	}

	defer shutdownFunc()

	writeErr := adminCli.WriteEnvoyDumpToZip(opts.Top.Ctx, proxyOutArchive)

	if writeErr == nil {
		fmt.Println("proxy snapshot written to " + proxyOutArchiveFile.Name())
	} else {
		fmt.Printf("Error writing proxy snapshot: %s", writeErr)
	}

	return proxyOutArchiveFile.Name(), writeErr
}

// createArchive forcibly deletes/creates the output directory
func createArchiveFile() (*os.File, error) {
	f, err := os.Create(fmt.Sprintf("glooctl-proxy-snapshot-%s.zip", time.Now().Format("2006-01-02-T15.04.05")))
	if err != nil {
		fmt.Printf("error creating proxy snapshot archive: %f\n", err)
	}
	return f, err
}
