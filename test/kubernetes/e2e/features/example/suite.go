package example

import (
	"context"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/test/kubernetes/e2e"
	"github.com/stretchr/testify/suite"
)

// FeatureSuite is the entire Suite of tests for the "example" feature
type FeatureSuite struct {
	suite.Suite

	ctx context.Context

	// testInstallation contains all the metadata/utilities necessary to execute a series of tests
	// against an installation of Gloo Gateway
	testInstallation *e2e.TestInstallation
}

func NewFeatureSuite(ctx context.Context, testInst *e2e.TestInstallation) *FeatureSuite {
	return &FeatureSuite{
		ctx:              ctx,
		testInstallation: testInst,
	}
}

func (s *FeatureSuite) SetupSuite() {
}

func (s *FeatureSuite) TearDownSuite() {
}

func (s *FeatureSuite) BeforeTest(suiteName, testName string) {
}

func (s *FeatureSuite) AfterTest(suiteName, testName string) {
}

func (s *FeatureSuite) TestInstallationSuccessful() {
	g := NewWithT(s.T())

	s.testInstallation.Assertions.AssertInstallationWasSuccessful(g, s.ctx)
}

func (s *FeatureSuite) TestFailureAllowed() {
	g := NewWithT(s.T())

	g.Expect(1).NotTo(Equal(2), "1 does not equal 2")
}
