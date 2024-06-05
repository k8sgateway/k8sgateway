package translator

// The tests in this file are private to our translator package, but warrant their own tests
// To avoid exporting methods unnecessarily, just for testing, we define these tests in the same package

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rotisserie/eris"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins/pluginutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reporting is a critical aspect of our translation engine that has minimal test coverage
// We split off this functionality into a standalone file to make it easier to make changes and test
var _ = Describe("Reporting", func() {

	DescribeTable("reportPluginProcessingErrorOrWarning",
		func(err error, reportWarning bool) {
			var (
				errCount, warningCount int
			)

			reportPluginProcessingErrorOrWarning(
				err,
				func() { errCount++ },
				func() { warningCount++ })

			if reportWarning {
				Expect(errCount).To(Equal(0))
				Expect(warningCount).To(Equal(1))
			} else {
				Expect(errCount).To(Equal(1))
				Expect(warningCount).To(Equal(0))
			}
		},
		Entry("generic error",
			eris.New("generic error"),
			false,
		),
		Entry("isWarningErr",
			pluginutils.NewUpstreamNotFoundErr(core.ResourceRef{}),
			true,
		),
		Entry("ConfigurationError with a warning",
			plugins.NewWarningConfigurationError("configuration-error"),
			true,
		),
		Entry("ConfigurationError without a warning",
			plugins.NewConfigurationError("configuration-error"),
			false,
		),
	)
})
