package kubectl

import (
	"bytes"
	"context"
	"fmt"
	"github.com/solo-io/gloo/pkg/utils/requestutils/curl"
	"github.com/solo-io/go-utils/threadsafe"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"

	"github.com/solo-io/gloo/pkg/utils/cmdutils"

	"io"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/solo-io/gloo/pkg/utils/kubeutils/portforward"
)

// Cli is a utility for executing `kubectl` commands
type Cli struct {
	// receiver is the default destination for the kubectl stdout and stderr
	receiver io.Writer

	// kubeContext is the optional value of the context for a given kubernetes cluster
	// If it is not supplied, no context will be included in the command
	kubeContext string
}

// NewCli returns an implementation of the kubectl.Cli
func NewCli() *Cli {
	return &Cli{
		receiver:    io.Discard,
		kubeContext: "",
	}
}

// WithReceiver sets the io.Writer that will be used by default for the stdout and stderr
// of cmdutils.Cmd created by the Cli
func (c *Cli) WithReceiver(receiver io.Writer) *Cli {
	c.receiver = receiver
	return c
}

// WithKubeContext sets the --context for the kubectl command invoked by the Cli
func (c *Cli) WithKubeContext(kubeContext string) *Cli {
	c.kubeContext = kubeContext
	return c
}

// Command returns a Cmd that executes kubectl command, including the --context if it is defined
// The Cmd sets the Stdout and Stderr to the receiver of the Cli
func (c *Cli) Command(ctx context.Context, args ...string) cmdutils.Cmd {
	if c.kubeContext != "" {
		args = append([]string{"--context", c.kubeContext}, args...)
	}

	return cmdutils.Command(ctx, "kubectl", args...).
		// For convenience, we set the stdout and stderr to the receiver
		// This can still be overwritten by consumers who use the commands
		WithStdout(c.receiver).
		WithStderr(c.receiver)
}

// RunCommand creates a Cmd and then runs it
func (c *Cli) RunCommand(ctx context.Context, args ...string) error {
	return c.Command(ctx, args...).Run().Cause()
}

// Apply applies the resources defined in the bytes, and returns an error if one occurred
func (c *Cli) Apply(ctx context.Context, content []byte, extraArgs ...string) error {
	args := append([]string{"apply", "-f", "-"}, extraArgs...)
	return c.Command(ctx, args...).
		WithStdin(bytes.NewBuffer(content)).
		Run().
		Cause()
}

// ApplyFile applies the resources defined in a file, and returns an error if one occurred
func (c *Cli) ApplyFile(ctx context.Context, fileName string, extraArgs ...string) error {
	applyArgs := append([]string{"apply", "-f", fileName}, extraArgs...)

	fileInput, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer func() {
		_ = fileInput.Close()
	}()

	return c.Command(ctx, applyArgs...).
		WithStdin(fileInput).
		Run().
		Cause()
}

// Delete deletes the resources defined in the bytes, and returns an error if one occurred
func (c *Cli) Delete(ctx context.Context, content []byte, extraArgs ...string) error {
	args := append([]string{"delete", "-f", "-"}, extraArgs...)
	return c.Command(ctx, args...).
		WithStdin(bytes.NewBuffer(content)).
		Run().
		Cause()
}

// DeleteFile deletes the resources defined in a file, and returns an error if one occurred
func (c *Cli) DeleteFile(ctx context.Context, fileName string, extraArgs ...string) error {
	applyArgs := append([]string{"delete", "-f", fileName}, extraArgs...)

	fileInput, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer func() {
		_ = fileInput.Close()
	}()

	return c.Command(ctx, applyArgs...).
		WithStdin(fileInput).
		Run().
		Cause()
}

// Copy copies a file from one location to another
func (c *Cli) Copy(ctx context.Context, from, to string) error {
	return c.RunCommand(ctx, "cp", from, to)
}

// DeploymentRolloutStatus waits for the deployment to complete rolling out
func (c *Cli) DeploymentRolloutStatus(ctx context.Context, deployment string, extraArgs ...string) error {
	rolloutArgs := append([]string{
		"rollout",
		"status",
		fmt.Sprintf("deployment/%s", deployment),
	}, extraArgs...)
	return c.RunCommand(ctx, rolloutArgs...)
}

// StartPortForward creates a PortForwarder based on the provides options, starts it, and returns the PortForwarder
// If an error was encountered while starting the PortForwarder, it is returned as well
// NOTE: It is the callers responsibility to close this port-forward
func (c *Cli) StartPortForward(ctx context.Context, options ...portforward.Option) (portforward.PortForwarder, error) {
	options = append([]portforward.Option{
		// We define some default values, which users can then override
		portforward.WithWriters(c.receiver, c.receiver),
		portforward.WithKubeContext(c.kubeContext),
	}, options...)

	portForwarder := portforward.NewCliPortForwarder(options...)
	err := portForwarder.Start(
		ctx,
		retry.LastErrorOnly(true),
		retry.Delay(100*time.Millisecond),
		retry.DelayType(retry.BackOffDelay),
		retry.Attempts(5),
	)
	return portForwarder, err
}

// CurlFromDeployment executes a curl from a deployment
// The assumption is that the deployment has a container named `curl` that will be used to invoke the request
func (c *Cli) CurlFromDeployment(ctx context.Context, deploymentMeta v1.ObjectMeta, options ...curl.Option) (string, error) {
	var outLocation threadsafe.Buffer

	args := []string{
		"exec",
		"-n", deploymentMeta.GetNamespace(),
		fmt.Sprintf("deployment/%s", deploymentMeta.GetName()),
		"-c",
		"curl",
		"--",
		"curl",
	}
	args = append(args, curl.BuildArgs(options...)...)

	err := c.Command(ctx, args...).WithStdin(&outLocation).Run().Cause()
	if err != nil {
		return "", err
	}

	return outLocation.String(), nil
}
