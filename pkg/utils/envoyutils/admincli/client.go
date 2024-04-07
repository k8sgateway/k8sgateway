package admincli

import (
	"context"
	"fmt"
	adminv3 "github.com/envoyproxy/go-control-plane/envoy/admin/v3"
	"github.com/golang/protobuf/jsonpb"
	"github.com/solo-io/gloo/pkg/utils/cmdutils"
	"github.com/solo-io/gloo/pkg/utils/requestutils/curl"
	"github.com/solo-io/go-utils/threadsafe"
	"io"
	"net/http"
)

const (
	ConfigDumpPath     = "config_dump"
	StatsPath          = "stats"
	ClustersPath       = "clusters"
	ListenersPath      = "listeners"
	ModifyRuntimePath  = "runtime_modify"
	ShutdownServerPath = "quitquitquit"
)

// Client is a utility for executing requests against the Envoy Admin API
type Client struct {
	// receiver is the default destination for the curl stdout and stderr
	receiver io.Writer

	// curlOptions is the set of default Option that the Client will use for curl commands
	curlOptions []curl.Option
}

// NewClient returns an implementation of the admincli.Client
func NewClient(receiver io.Writer, curlOptions []curl.Option) *Client {
	defaultCurlOptions := []curl.Option{
		curl.WithScheme("http"),
		// 5 retries, exponential back-off, 10 second max
		curl.WithRetries(5, 0, 10),
	}

	return &Client{
		receiver:    receiver,
		curlOptions: append(defaultCurlOptions, curlOptions...),
	}
}

func (c *Client) Command(ctx context.Context, options ...curl.Option) cmdutils.Cmd {
	commandCurlOptions := append(
		c.curlOptions,
		// Ensure any options defined for this command can override any defaults that the Client has defined
		options...)
	curlArgs := curl.BuildArgs(ctx, commandCurlOptions...)

	return cmdutils.Command(ctx, "curl", curlArgs...).
		// For convenience, we set the stdout and stderr to the receiver
		// This can still be overwritten by consumers who use the commands
		WithStdout(c.receiver).
		WithStderr(c.receiver)
}

func (c *Client) RunCommand(ctx context.Context, options ...curl.Option) error {
	return c.Command(ctx, options...).Run().Cause()
}

func (c *Client) RequestPathCmd(ctx context.Context, path string) cmdutils.Cmd {
	return c.Command(ctx, curl.WithPath(path))
}

func (c *Client) StatsCmd(ctx context.Context) cmdutils.Cmd {
	return c.RequestPathCmd(ctx, StatsPath)
}

func (c *Client) GetStats(ctx context.Context) (string, error) {
	var outLocation threadsafe.Buffer
	
	err := c.StatsCmd(ctx).WithStdout(&outLocation).Run().Cause()
	if err != nil {
		return "", err
	}

	return outLocation.String(), nil
}

func (c *Client) ClustersCmd(ctx context.Context) cmdutils.Cmd {
	return c.RequestPathCmd(ctx, ClustersPath)
}

func (c *Client) ListenersCmd(ctx context.Context) cmdutils.Cmd {
	return c.RequestPathCmd(ctx, ListenersPath)
}

func (c *Client) ConfigDumpCmd(ctx context.Context) cmdutils.Cmd {
	return c.RequestPathCmd(ctx, ConfigDumpPath)
}

func (c *Client) GetConfigDump(ctx context.Context) (*adminv3.ConfigDump, error) {
	var (
		cfgDump     adminv3.ConfigDump
		outLocation threadsafe.Buffer
	)

	err := c.ConfigDumpCmd(ctx).WithStdout(&outLocation).Run().Cause()
	if err != nil {
		return nil, err
	}

	jsonpbMarshaler := &jsonpb.Unmarshaler{
		// Ever since upgrading the go-control-plane to v0.10.1 this test fails with the following error:
		// unknown field \"hidden_envoy_deprecated_build_version\" in envoy.config.core.v3.Node"
		// Set AllowUnknownFields to true to get around this
		AllowUnknownFields: true,
	}

	if err = jsonpbMarshaler.Unmarshal(&outLocation, &cfgDump); err != nil {
		return nil, err
	}

	return &cfgDump, nil
}

func (c *Client) ModifyRuntimeConfiguration(ctx context.Context, queryParameters string) error {
	return c.RunCommand(ctx,
		curl.WithPath(fmt.Sprintf("%s?%s", ModifyRuntimePath, queryParameters)),
		curl.WithMethod(http.MethodPost))
}

func (c *Client) ShutdownServer(ctx context.Context) error {
	return c.RunCommand(ctx,
		curl.WithPath(ShutdownServerPath),
		curl.WithMethod(http.MethodPost))
}
