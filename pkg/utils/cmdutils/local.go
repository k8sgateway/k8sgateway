package cmdutils

import (
	"context"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
	"github.com/solo-io/go-utils/threadsafe"
)

var (
	_            Cmd   = &LocalCmd{}
	_            Cmder = &LocalCmder{}
	DefaultCmder       = &LocalCmder{}
)

// Command is a convenience wrapper over DefaultCmder.Command
func Command(ctx context.Context, command string, args ...string) Cmd {
	return DefaultCmder.Command(ctx, command, args...).
		WithStdout(io.Discard).
		WithStderr(io.Discard)
}

// LocalCmder is a factory for LocalCmd, implementing Cmder
type LocalCmder struct{}

// Command returns a Cmd which includes the running process's `Environment`
func (c *LocalCmder) Command(ctx context.Context, name string, arg ...string) Cmd {
	cmd := &LocalCmd{
		Cmd: exec.CommandContext(ctx, name, arg...),
	}

	// By default, assign the env variables for the command
	// Consumers of this Cmd can then override it, if they want
	return cmd.WithEnv(os.Environ()...)
}

// LocalCmd wraps os/exec.Cmd, implementing the cmdutils.Cmd interface
type LocalCmd struct {
	*exec.Cmd
}

// WithEnv sets env
func (cmd *LocalCmd) WithEnv(env ...string) Cmd {
	//disable DEBUG=1 from getting through to command
	for i, pair := range env {
		if strings.HasPrefix(pair, "DEBUG") {
			env = append(env[:i], env[i+1:]...)
			break
		}
	}

	cmd.Env = env
	return cmd
}

// WithStdin sets stdin
func (cmd *LocalCmd) WithStdin(r io.Reader) Cmd {
	cmd.Stdin = r
	return cmd
}

// WithStdout set stdout
func (cmd *LocalCmd) WithStdout(w io.Writer) Cmd {
	cmd.Stdout = w
	return cmd
}

// WithStderr sets stderr
func (cmd *LocalCmd) WithStderr(w io.Writer) Cmd {
	cmd.Stderr = w
	return cmd
}

// Run runs the command
// If the returned error is non-nil, it should be of type *RunError
func (cmd *LocalCmd) Run() *RunError {
	var combinedOutput threadsafe.Buffer

	cmd.Stdout = io.MultiWriter(cmd.Stdout, &combinedOutput)
	cmd.Stderr = io.MultiWriter(cmd.Stderr, &combinedOutput)

	if err := cmd.Cmd.Run(); err != nil {
		return &RunError{
			command:    cmd.Args,
			output:     combinedOutput.Bytes(),
			inner:      err,
			stackTrace: errors.WithStack(err),
		}
	}
	return nil
}
