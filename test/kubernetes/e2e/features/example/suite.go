package example

import (
	"context"

	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/test/kubernetes/e2e"
	"github.com/stretchr/testify/suite"
)

// testingSuite is the entire Suite of tests for the "example" feature
// Typically, we would include a link to the feature code here
type testingSuite struct {
	suite.Suite

	ctx context.Context

	// testInstallation contains all the metadata/utilities necessary to execute a series of tests
	// against an installation of Gloo Gateway
	testInstallation *e2e.TestInstallation
}

func NewTestingSuite(ctx context.Context, testInst *e2e.TestInstallation) suite.TestingSuite {
	return &testingSuite{
		ctx:              ctx,
		testInstallation: testInst,
	}
}

func (s *testingSuite) SetupSuite() {
}

func (s *testingSuite) TearDownSuite() {
}

func (s *testingSuite) BeforeTest(suiteName, testName string) {
}

func (s *testingSuite) AfterTest(suiteName, testName string) {
}

func (s *testingSuite) TestExampleAssertion() {
	// Testify assertion
	s.Assert().NotEqual(1, 2, "1 does not equal 2")

	// Testify assertion, using the TestInstallation to provide it
	s.testInstallation.Assertions.NotEqual(1, 2, "1 does not equal 2")

	// Gomega assertion, using the TestInstallation to provide it
	s.testInstallation.Assertions.Expect(1).NotTo(Equal(2), "1 does not equal 2")
}
