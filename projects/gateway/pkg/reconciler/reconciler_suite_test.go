package reconciler_test

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
	"github.com/solo-io/solo-kit/pkg/utils/envutils"
)

var (
	test      *testing.T
	namespace = "reconciler-test-ns"

	_ = BeforeSuite(func() {
		err := os.Setenv(envutils.PodNamespaceEnvName, namespace)
		Expect(err).NotTo(HaveOccurred())
	})

	_ = AfterSuite(func() {
		err := os.Unsetenv(envutils.PodNamespaceEnvName)
		Expect(err).NotTo(HaveOccurred())
	})
)

func TestReconciler(t *testing.T) {
	test = t
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("junit.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "Reconciler Suite", []Reporter{junitReporter})
}
