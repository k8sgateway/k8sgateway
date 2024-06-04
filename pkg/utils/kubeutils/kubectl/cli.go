package kubectl

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"k8s.io/apimachinery/pkg/types"

	"github.com/solo-io/gloo/pkg/utils/cmdutils"
	"github.com/solo-io/gloo/pkg/utils/requestutils/curl"
	"github.com/solo-io/k8s-utils/testutils/kube"

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

// PodExecOptions describes the options used to execute a command in a pod
type PodExecOptions struct {
	Name      string
	Namespace string
	Container string
}

// NewCli returns an implementation of the kubectl.Cli
func NewCli() *Cli {
	return &Cli{
		receiver:    io.Discard,
		kubeContext: "",
	}
}

type CurlResponse struct {
	Headers string
	Body    string
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
	_, err := c.ApplyFileWithOutput(ctx, fileName, extraArgs...)
	return err
}

// ApplyFileWithOutput applies the resources defined in a file,
// if an error occurred, it will be returned along with the output of the command
func (c *Cli) ApplyFileWithOutput(ctx context.Context, fileName string, extraArgs ...string) (string, error) {
	applyArgs := append([]string{"apply", "-f", fileName}, extraArgs...)

	fileInput, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = fileInput.Close()
	}()

	runErr := c.Command(ctx, applyArgs...).
		WithStdin(fileInput).
		Run()
	return runErr.OutputString(), runErr.Cause()
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
	_, err := c.DeleteFileWithOutput(ctx, fileName, extraArgs...)
	return err
}

// DeleteFileWithOutput deletes the resources defined in a file,
// if an error occurred, it will be returned along with the output of the command
func (c *Cli) DeleteFileWithOutput(ctx context.Context, fileName string, extraArgs ...string) (string, error) {
	applyArgs := append([]string{"delete", "-f", fileName}, extraArgs...)

	fileInput, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = fileInput.Close()
	}()

	runErr := c.Command(ctx, applyArgs...).
		WithStdin(fileInput).
		Run()
	return runErr.OutputString(), runErr.Cause()
}

// DeleteFileSafe deletes the resources defined in a file, and returns an error if one occurred
// This differs from DeleteFile in that we always append --ignore-not-found
func (c *Cli) DeleteFileSafe(ctx context.Context, fileName string, extraArgs ...string) error {
	safeArgs := append(extraArgs, "--ignore-not-found")
	return c.DeleteFile(ctx, fileName, safeArgs...)
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

// CurlFromEphemeralPod executes a Curl from a pod, using an ephemeral container to execute the Curl command
func (c *Cli) CurlFromEphemeralPod(ctx context.Context, podMeta types.NamespacedName, options ...curl.Option) string {
	appendOption := func(option curl.Option) {
		options = append(options, option)
	}

	// The e2e test assertions rely on the transforms.WithCurlHttpResponse to validate the response is what
	// we would expect
	// For this transform to behave appropriately, we must execute the request with verbose=true
	appendOption(curl.VerboseOutput())

	curlArgs := curl.BuildArgs(options...)

	return kube.CurlWithEphemeralPodStable(
		ctx,
		c.receiver,
		c.kubeContext,
		podMeta.Namespace,
		podMeta.Name,
		curlArgs...)
}

// CurlFromPod executes a Curl request from the given pod for the given options.
// It differs from CurlFromEphemeralPod in that it does not uses an ephemeral container to execute the Curl command
func (c *Cli) CurlFromPod(ctx context.Context, podOpts PodExecOptions, options ...curl.Option) (*CurlResponse, error) {
	appendOption := func(option curl.Option) {
		options = append(options, option)
	}

	// The e2e test assertions rely on the transforms.WithCurlHttpResponse to validate the response is what
	// we would expect
	// For this transform to behave appropriately, we must execute the request with verbose=true
	appendOption(curl.VerboseOutput())

	curlArgs := curl.BuildArgs(options...)

	container := podOpts.Container
	if container == "" {
		container = "curl"
	}

	args := append([]string{
		"exec",
		"--container=" + container,
		podOpts.Name,
		"-n",
		podOpts.Namespace,
		"--",
		"curl",
		"--connect-timeout",
		"1",
		"--max-time",
		"5",
	}, curlArgs...)

	stdout, stderr, err := c.ExecuteOn(ctx, c.kubeContext, nil, args...)
	return &CurlResponse{Body: stdout, Headers: stderr}, err
}

func (c *Cli) ExecuteOn(ctx context.Context, kubeContext string, stdin *bytes.Buffer, args ...string) (string, string, error) {
	args = append([]string{"--context", kubeContext}, args...)
	return c.Execute(ctx, stdin, args...)
}

func (c *Cli) Execute(ctx context.Context, stdin *bytes.Buffer, args ...string) (string, string, error) {
	stdout := new(strings.Builder)
	stderr := new(strings.Builder)

	err := cmdutils.Command(ctx, "kubectl", args...).
		// For convenience, we set the stdout and stderr to the receiver
		// This can still be overwritten by consumers who use the commands
		WithStdout(stdout).
		WithStderr(stderr).Run().Cause()

	return stdout.String(), stderr.String(), err
}
