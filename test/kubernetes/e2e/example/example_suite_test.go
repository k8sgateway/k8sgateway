package example_test

import (
	"context"

	"github.com/solo-io/gloo/test/helpers"
	"github.com/solo-io/gloo/test/kubernetes/e2e"
	"github.com/solo-io/gloo/test/kubernetes/testutils/assertions"
	"github.com/solo-io/gloo/test/kubernetes/testutils/cluster"
	"github.com/solo-io/gloo/test/kubernetes/testutils/operations"
	"github.com/solo-io/gloo/test/kubernetes/testutils/operations/provider"
	"github.com/solo-io/gloo/test/kubernetes/testutils/runtime"
	skhelpers "github.com/solo-io/solo-kit/test/helpers"

	"testing"

	. "github.com/onsi/ginkgo/v2"
)

func TestExampleSuite(t *testing.T) {
	helpers.RegisterGlooDebugLogPrintHandlerAndClearLogs()
	skhelpers.RegisterCommonFailHandlers()

	RunSpecs(t, "Example Suite")
}

var (
	testSuite *e2e.TestSuite
)

var _ = BeforeSuite(func(ctx context.Context) {
	runtimeContext := runtime.NewContext()

	// We try to isolate the usage of Ginkgo to only where are tests are invoked
	testingFramework := GinkgoTB()

	// Construct the cluster.Context for this suite
	clusterContext := cluster.MustKindContext(testingFramework, runtimeContext.ClusterName)

	testSuite = &e2e.TestSuite{
		TestingFramework: testingFramework,

		// Create an operator which is responsible for executing operations against the cluster
		Operator: operations.NewGinkgoOperator(),

		// Create an operations provider, and point it to the running cluster
		OperationsProvider: provider.NewOperationProvider().WithClusterContext(clusterContext),

		// Create an assertions provider, and point it to the running cluster
		AssertionsProvider: assertions.NewProvider().
			WithClusterContext(clusterContext).
			WithTestingFramework(testingFramework),
	}
})
