package full_envoy_validation

import (
	"context"

	"github.com/solo-io/gloo/test/kubernetes/e2e"
	"github.com/solo-io/gloo/test/kubernetes/e2e/features/validation"
	"github.com/stretchr/testify/suite"
)

var _ e2e.NewSuiteFunc = NewTestingSuite

// testingSuite is the entire Suite of tests for the webhook validation fullEnvoyValidation=true feature
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

// TestRejectInvalidTransformation checks webhook rejects invalid transformation when fullEnvoyValidation=true
func (s *testingSuite) TestRejectInvalidTransformation() {
	// rejects invalid inja template in transformation
	output, err := s.testInstallation.Actions.Kubectl().ApplyFileWithOutput(s.ctx, validation.VSTransformationHeaderText, "-n", s.testInstallation.Metadata.InstallNamespace)
	s.Assert().Error(err)
	s.Assert().Contains(output, "Failed to parse response template: Failed to parse "+
		"header template ':status': [inja.exception.parser_error] (at 1:92) expected statement close, got '%'")
}
