package install_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/install"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/flagutils"
	"github.com/spf13/pflag"
	"io"
	"strings"
)

type MockKubectl struct {
	expected []string
	next     int
}

func NewMockKubectl(cmds ...string) *MockKubectl {
	return &MockKubectl{
		expected: cmds,
		next: 0,
	}
}

func (k *MockKubectl) Kubectl(stdin io.Reader, args ...string) error {
	// If this fails then the CLI tried to run commands we didn't account for in the mock
	Expect(k.next < len(k.expected)).To(BeTrue())
	Expect(stdin).To(BeNil())
	cmd := strings.Join(args, " ")
	Expect(cmd).To(BeEquivalentTo(k.expected[k.next]))
	k.next = k.next + 1
	return nil
}

var _ = Describe("Uninstall", func() {

	const (
		deleteCrds = "delete crd gateways.gateway.solo.io proxies.gateway.solo.io settings.gateway.solo.io upstreams.gateway.solo.io virtualservices.gateway.solo.io"
	)

	var flagSet *pflag.FlagSet
	var opts options.Options

	BeforeEach(func() {
		flagSet = pflag.NewFlagSet("uninstall", pflag.ContinueOnError)
		opts = options.Options{}
		flagutils.AddUninstallFlags(flagSet, &opts.Uninstall)
	})

	uninstall := func(cli *MockKubectl) {
		install.UninstallGloo(&opts, cli)
		// If this fails, then the mock CLI had extra commands that were expected to run but weren't
		Expect(cli.next).To(BeEquivalentTo(len(cli.expected)))
	}

	It("works with no args", func () {
		flagSet.Parse([]string {})
		cli := NewMockKubectl(
			"delete Deployment -l gloo -n gloo-system",
			"delete Service -l gloo -n gloo-system",
			"delete ConfigMap -l gloo -n gloo-system")
		uninstall(cli)
	})

	It("works with namespace", func () {
		flagSet.Parse([]string { "-n", "foo" })
		cli := NewMockKubectl(
			"delete Deployment -l gloo -n foo",
			"delete Service -l gloo -n foo",
			"delete ConfigMap -l gloo -n foo")
		uninstall(cli)
	})

	It("works with delete crds", func () {
		flagSet.Parse([]string {"--delete-crds"})
		cli := NewMockKubectl(
			"delete Deployment -l gloo -n gloo-system",
			"delete Service -l gloo -n gloo-system",
			"delete ConfigMap -l gloo -n gloo-system",
			deleteCrds)
		uninstall(cli)
	})

	It("works with delete crds and namespace", func () {
		flagSet.Parse([]string { "-n", "foo", "--delete-crds" })
		cli := NewMockKubectl(
			"delete Deployment -l gloo -n foo",
			"delete Service -l gloo -n foo",
			"delete ConfigMap -l gloo -n foo",
			deleteCrds)
		uninstall(cli)
	})

	It("works with delete crds and namespace", func () {
		flagSet.Parse([]string { "-n", "foo", "--delete-crds" })
		cli := NewMockKubectl(
			"delete Deployment -l gloo -n foo",
			"delete Service -l gloo -n foo",
			"delete ConfigMap -l gloo -n foo",
			deleteCrds)
		uninstall(cli)
	})

	It("works with delete namespace", func () {
		flagSet.Parse([]string {"--delete-namespace"})
		cli := NewMockKubectl(
			"delete namespace gloo-system")
		uninstall(cli)
	})

	It("works with delete namespace with custom namespace", func () {
		flagSet.Parse([]string {"--delete-namespace", "-n", "foo"})
		cli := NewMockKubectl(
			"delete namespace foo")
		uninstall(cli)
	})

	It("works with delete namespace and crds", func () {
		flagSet.Parse([]string {"--delete-namespace", "--delete-crds"})
		cli := NewMockKubectl(
			"delete namespace gloo-system",
			deleteCrds)
		uninstall(cli)
	})

	It("works with delete crds and namespace with custom namespace", func () {
		flagSet.Parse([]string {"--delete-namespace", "--delete-crds", "-n", "foo"})
		cli := NewMockKubectl(
			"delete namespace foo",
			deleteCrds)
		uninstall(cli)
	})

	It("works with delete all", func() {
		flagSet.Parse([]string {"--all"})
		cli := NewMockKubectl(
			"delete namespace gloo-system",
			deleteCrds,
			"delete ClusterRole -l gloo",
			"delete ClusterRoleBinding -l gloo")
		uninstall(cli)
	})

	It("works with delete all custom namespace", func() {
		flagSet.Parse([]string {"--all", "-n", "foo"})
		cli := NewMockKubectl(
			"delete namespace foo",
			deleteCrds,
			"delete ClusterRole -l gloo",
			"delete ClusterRoleBinding -l gloo")
		uninstall(cli)
	})
})
