package install

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	. "github.com/onsi/gomega"

	"github.com/solo-io/gloo/pkg/cliutil"
)

func KubectlApply(manifest []byte, extraArgs ...string) error {
	return Kubectl(bytes.NewBuffer(manifest), append([]string{"apply", "-f", "-"}, extraArgs...)...)
}

func KubectlApplyOut(manifest []byte, extraArgs ...string) ([]byte, error) {
	return KubectlOut(bytes.NewBuffer(manifest), append([]string{"apply", "-f", "-"}, extraArgs...)...)
}

func KubectlDelete(manifest []byte, extraArgs ...string) error {
	return Kubectl(bytes.NewBuffer(manifest), append([]string{"delete", "-f", "-"}, extraArgs...)...)
}

type KubeCli interface {
	Kubectl(stdin io.Reader, args ...string) error
	KubectlOut(stdin io.Reader, args ...string) ([]byte, error)
}

type CmdKubectl struct{}

var _ KubeCli = &CmdKubectl{}

func (k *CmdKubectl) Kubectl(stdin io.Reader, args ...string) error {
	return Kubectl(stdin, args...)
}

func (k *CmdKubectl) KubectlOut(stdin io.Reader, args ...string) ([]byte, error) {
	return KubectlOut(stdin, args...)
}

type MockKubectlCalls struct {
	Mocks []*MockKubectl
	Next            int
}

type MockKubectl struct {
	Expected        []string
	Next            int
	StdoutLines     []string
	StdoutLineIndex int
}

var _ KubeCli = &MockKubectlCalls{}

func NewMockKubectlCalls(mocks []*MockKubectl) *MockKubectlCalls {
	return &MockKubectlCalls{
		Mocks:    mocks,
		Next:        0,
	}
}

func NewMockKubectl(cmds []string, stdoutLines []string) *MockKubectl {
	return &MockKubectl{
		Expected:    cmds,
		Next:        0,
		StdoutLines: stdoutLines,
	}
}

func (kMocks *MockKubectlCalls) Kubectl(stdin io.Reader, args ...string) error {
	k := kMocks.Mocks[kMocks.Next]
	// If this fails then the CLI tried to run commands we didn't account for in the mock
	Expect(k.Next < len(k.Expected)).To(BeTrue())
	Expect(stdin).To(BeNil())
	cmd := strings.Join(args, " ")
	Expect(cmd).To(BeEquivalentTo(k.Expected[k.Next]))
	k.Next = k.Next + 1
	if k.Next == len(k.Expected) {
		kMocks.Next = kMocks.Next + 1
	}
	return nil
}

func (kMocks *MockKubectlCalls) KubectlOut(stdin io.Reader, args ...string) ([]byte, error) {
	k := kMocks.Mocks[kMocks.Next]
	Expect(k.Next < len(k.Expected)).To(BeTrue(), "MockKubectl did not have a next command for KubectlOut")
	Expect(stdin).To(BeNil(), "Should have passed nil to MockKubectl.KubectlOut")
	cmd := strings.Join(args, " ")
	Expect(cmd).To(BeEquivalentTo(k.Expected[k.Next]), "Wrong next command for MockKubectl.KubectlOut")
	k.Next = k.Next + 1
	Expect(k.StdoutLineIndex < len(k.StdoutLines)).To(BeTrue(), "Mock kubectl has run out of stdout lines on command "+cmd)
	stdOutLine := k.StdoutLines[k.StdoutLineIndex]
	k.StdoutLineIndex = k.StdoutLineIndex + 1
	if k.Next == len(k.Expected) {
		kMocks.Next = kMocks.Next + 1
	}
	return []byte(stdOutLine), nil
}

var verbose bool

func SetVerbose(b bool) {
	verbose = b
}

func Kubectl(stdin io.Reader, args ...string) error {
	kubectl := exec.Command("kubectl", args...)
	if stdin != nil {
		kubectl.Stdin = stdin
	}
	if verbose {
		fmt.Fprintf(os.Stderr, "running kubectl command: %v\n", kubectl.Args)
		kubectl.Stdout = os.Stdout
		kubectl.Stderr = os.Stderr
	} else {
		// use logfile
		cliutil.Initialize()
		kubectl.Stdout = cliutil.GetLogger()
		kubectl.Stderr = cliutil.GetLogger()
	}
	return kubectl.Run()
}

func KubectlOut(stdin io.Reader, args ...string) ([]byte, error) {
	kubectl := exec.Command("kubectl", args...)

	if stdin != nil {
		kubectl.Stdin = stdin
	}

	var stdout, stderr io.Writer
	if verbose {
		fmt.Fprintf(os.Stderr, "running kubectl command: %v\n", kubectl.Args)
		stdout = os.Stdout
		stderr = os.Stderr
	} else {
		// use logfile
		cliutil.Initialize()
		stdout = cliutil.GetLogger()
		stderr = cliutil.GetLogger()
	}

	buf := &bytes.Buffer{}

	kubectl.Stdout = io.MultiWriter(stdout, buf)
	kubectl.Stderr = io.MultiWriter(stderr, buf)

	err := kubectl.Run()

	return buf.Bytes(), err
}
