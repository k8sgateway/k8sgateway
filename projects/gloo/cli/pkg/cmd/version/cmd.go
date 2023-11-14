package version

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/solo-io/go-utils/contextutils"

	"github.com/ghodss/yaml"
	"github.com/olekukonko/tablewriter"
	"github.com/rotisserie/eris"
	linkedversion "github.com/solo-io/gloo/pkg/version"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/constants"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/flagutils"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/printers"
	"github.com/solo-io/go-utils/cliutils"
	"github.com/spf13/cobra"
)

const (
	undefinedServer = "Server: version undefined, could not find any version of gloo running"
)

var (
	NoNamespaceAllError = eris.New("single namespace must be specified, cannot be namespace all for version command")
)

type ClientVersion struct {
	Version string
}
type Versions struct {
	Server *ServerVersionInfo
	Client ClientVersion
}

func RootCmd(opts *options.Options, optionsFunc ...cliutils.OptionsFunc) *cobra.Command {
	// Default output for version command is JSON
	versionOutput := printers.JSON

	cmd := &cobra.Command{
		Use:     constants.VERSION_COMMAND.Use,
		Aliases: constants.VERSION_COMMAND.Aliases,
		Short:   constants.VERSION_COMMAND.Short,
		Long:    constants.VERSION_COMMAND.Long,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.Top.Output = versionOutput

			if opts.Top.Namespace == "" {
				return NoNamespaceAllError
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return printVersion(NewKube(opts.Top.Namespace, ""), os.Stdout, opts)
		},
	}

	pflags := cmd.PersistentFlags()
	flagutils.AddOutputFlag(pflags, &versionOutput)
	flagutils.AddNamespaceFlag(pflags, &opts.Top.Namespace)

	return cmd
}

func GetClientServerVersions(ctx context.Context, sv ServerVersion) (*Versions, error) {
	v := &Versions{
		Client: getClientVersion(),
	}
	serverVersion, err := sv.Get(ctx)
	if err != nil {
		return v, err
	}
	v.Server = serverVersion
	return v, nil
}

func getClientVersion() ClientVersion {
	return ClientVersion{
		Version: linkedversion.Version,
	}
}

func printVersion(sv ServerVersion, w io.Writer, opts *options.Options) error {
	vrs, _ := GetClientServerVersions(opts.Top.Ctx, sv)
	// ignoring error so we still print client version even if we can't get server versions (e.g., not deployed, no rbac)
	switch opts.Top.Output {
	case printers.JSON:
		clientVersion, err := GetJson(vrs.Client)
		if err != nil {
			return err
		}
		clientVersionStr := string(clientVersion)
		clientVersionStr = strings.ReplaceAll(clientVersionStr, "\n", "")
		fmt.Fprintf(w, "Client: %s\n", clientVersionStr)
		if vrs.Server == nil {
			fmt.Fprintln(w, undefinedServer)
			return nil
		}
		fmt.Fprint(w, "Server: ")
		for _, v := range vrs.Server.Containers {
			serverVersionStr, err := GetJson(v)
			if err != nil {
				return err
			}
			fmt.Fprintf(w, "%s\n", string(serverVersionStr))
		}
	case printers.YAML:
		clientVersion, err := GetYaml(vrs.Client)
		if err != nil {
			return err
		}
		clientVersionStr := string(clientVersion)
		clientVersionStr = strings.ReplaceAll(clientVersionStr, "\n", "")
		fmt.Fprintf(w, "Client: %s\n", clientVersionStr)
		if vrs.Server == nil {
			fmt.Fprintln(w, undefinedServer)
			return nil
		}
		fmt.Fprintln(w, "Server:")
		for _, v := range vrs.Server.Containers {
			serverVersion, err := GetYaml(v)
			if err != nil {
				return err
			}
			serverVersionStr := string(serverVersion)
			clientVersionStr = strings.TrimRight(clientVersionStr, "\n")
			fmt.Fprintf(w, "%s\n", serverVersionStr)
		}
	default:
		fmt.Fprintf(w, "Client: version: %s\n", vrs.Client.Version)
		if vrs.Server == nil {
			fmt.Fprintln(w, undefinedServer)
			return nil
		}
		srv := vrs.Server
		if srv == nil {
			fmt.Println(undefinedServer)
			return nil
		}

		table := tablewriter.NewWriter(w)
		headers := []string{"Namespace", "Deployment-Type", "Containers"}
		var rows [][]string
		content := []string{srv.Namespace}
		for i, container := range srv.Containers {
			name := fmt.Sprintf("%s: %s", container.Repository, container.Tag)
			if i == 0 {
				rows = append(rows, append(content, name))
			} else {
				rows = append(rows, []string{"", "", name})
			}
		}

		table.SetHeader(headers)
		table.AppendBulk(rows)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		fmt.Println("Server:")
		table.Render()
	}
	return nil
}

func getDistributionName(name string, enterprise bool) string {
	if enterprise {
		return name + " Enterprise"
	}
	return name
}

func GetJson(pb any) ([]byte, error) {
	data, err := json.Marshal(pb)
	if err != nil {
		contextutils.LoggerFrom(context.Background()).DPanic(err)
		return nil, err
	}
	return data, nil
}

func GetYaml(pb any) ([]byte, error) {
	jsn, err := GetJson(pb)
	if err != nil {
		contextutils.LoggerFrom(context.Background()).DPanic(err)
		return nil, err
	}
	data, err := yaml.JSONToYAML(jsn)
	if err != nil {
		contextutils.LoggerFrom(context.Background()).DPanic(err)
		return nil, err
	}
	return data, nil
}
