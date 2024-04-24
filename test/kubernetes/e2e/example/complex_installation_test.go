package example

import (
	"path/filepath"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/test/kubernetes/e2e"
	"github.com/solo-io/gloo/test/kubernetes/e2e/features/example"
	"github.com/solo-io/gloo/test/kubernetes/testutils/gloogateway"
	"github.com/solo-io/skv2/codegen/util"
	"github.com/stretchr/testify/suite"
)

func (s *ClusterSuite) TestComplexInstallation() {

	var testInstallation *e2e.TestInstallation

	s.T().Run("before", func(t *testing.T) {
		testInstallation = s.testCluster.RegisterTestInstallation(
			&gloogateway.Context{
				InstallNamespace:   "complex-example",
				ValuesManifestFile: filepath.Join(util.MustGetThisDir(), "manifests", "complex-example.yaml"),
			},
		)

		err := testInstallation.InstallGlooGateway(s.ctx, testInstallation.Actions.Glooctl().NewTestHelperInstallAction())
		Expect(err).NotTo(HaveOccurred())
	})

	s.T().Run("example feature", func(t *testing.T) {
		suite.Run(t, example.NewFeatureSuite(s.ctx, testInstallation))
	})

	s.T().Run("after", func(t *testing.T) {
		err := testInstallation.UninstallGlooGateway(s.ctx, testInstallation.Actions.Glooctl().NewTestHelperUninstallAction())
		Expect(err).NotTo(HaveOccurred())

		s.testCluster.UnregisterTestInstallation(testInstallation)
	})

}
