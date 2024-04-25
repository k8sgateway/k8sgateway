package example

import (
	"context"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/test/kubernetes/e2e"
	"github.com/solo-io/gloo/test/kubernetes/testutils/cluster"
	"github.com/solo-io/gloo/test/kubernetes/testutils/runtime"
	"github.com/stretchr/testify/suite"
)

func TestClusterSuite(t *testing.T) {
	RegisterFailHandler(Fail)

	runtimeContext := runtime.NewContext()

	// Construct the cluster.Context for this suite
	clusterContext := cluster.MustKindContext(runtimeContext.ClusterName)

	clusterSuite := NewClusterSuite(context.Background(), &e2e.TestCluster{
		RuntimeContext: runtimeContext,
		ClusterContext: clusterContext,
	})

	t.Run("cluster suite", func(t *testing.T) {
		suite.Run(t, clusterSuite)
	})
}

func NewClusterSuite(ctx context.Context, testCluster *e2e.TestCluster) *ClusterSuite {
	return &ClusterSuite{
		ctx:         ctx,
		testCluster: testCluster,
	}
}

// ClusterSuite is the entire Suite of tests against a cluster
type ClusterSuite struct {
	suite.Suite

	ctx         context.Context
	testCluster *e2e.TestCluster
}

func (s *ClusterSuite) SetupSuite() {}

func (s *ClusterSuite) TearDownSuite() {
	s.testCluster.PreFailHandler()

}
