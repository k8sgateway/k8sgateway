package helpers_test

import (
	"testing"

	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHelpers(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("junit.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "Helpers Suite", []Reporter{junitReporter})
}
