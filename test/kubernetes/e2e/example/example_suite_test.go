package example_test

import (
	"github.com/solo-io/gloo/test/helpers"
	"github.com/solo-io/gloo/test/kubernetes/testutils/assertions"
	"github.com/solo-io/gloo/test/kubernetes/testutils/operations"
	"github.com/solo-io/gloo/test/kubernetes/testutils/operations/provider"
	"github.com/solo-io/gloo/test/testutils"
	"github.com/solo-io/gloo/test/testutils/kubeutils"
	skhelpers "github.com/solo-io/solo-kit/test/helpers"

	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
)

func TestExampleSuite(t *testing.T) {
	helpers.RegisterGlooDebugLogPrintHandlerAndClearLogs()
	skhelpers.RegisterCommonFailHandlers()

	RunSpecs(t, "Example Suite")
}

var (
	clusterContext     *kubeutils.ClusterContext
	operator           *operations.Operator
	operationsProvider *provider.Provider
	assertionProvider  *assertions.Provider
)

var _ = BeforeSuite(func() {
	clusterContext = kubeutils.MustKindClusterContext(os.Getenv(testutils.ClusterName))

	// Create an operator which is responsible for execution Operation agains the cluster
	operator = operations.NewGinkgoOperator()

	// Set the operations provider to point to the running cluster
	operationsProvider = provider.NewProvider().WithClusterContext(clusterContext)

	// Set the assertion provider to point to the running cluster
	assertionProvider = assertions.NewProvider().WithClusterContext(clusterContext)
})

var _ = AfterSuite(func() {

})
