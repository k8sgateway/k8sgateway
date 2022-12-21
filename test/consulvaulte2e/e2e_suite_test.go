package consulvaulte2e_test

import (
	"os"
	"testing"

	testhelpers "github.com/solo-io/gloo/test/helpers"

	"github.com/onsi/ginkgo/reporters"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/solo-io/gloo/test/services"
	"github.com/solo-io/solo-kit/test/helpers"
)

var (
	envoyFactory  *services.EnvoyFactory
	consulFactory *services.ConsulFactory
	vaultFactory  *services.VaultFactory
)

var _ = BeforeSuite(func() {
	var err error
	envoyFactory, err = services.NewEnvoyFactory()
	Expect(err).NotTo(HaveOccurred())
	consulFactory, err = services.NewConsulFactory()
	Expect(err).NotTo(HaveOccurred())
	vaultFactory, err = services.NewVaultFactory()
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	_ = envoyFactory.Clean()
	_ = consulFactory.Clean()
	_ = vaultFactory.Clean()
})

func TestE2e(t *testing.T) {
	testhelpers.ValidateRequirementsAndNotifyGinkgo(
		testhelpers.TruthyEnv("RUN_VAULT_TESTS"),
		testhelpers.TruthyEnv("RUN_CONSUL_TESTS"),
	)

	// set KUBECONFIG to a nonexistent cfg.
	// this way we are also testing that Gloo can run without a kubeconfig present
	os.Setenv("KUBECONFIG", ".")

	helpers.RegisterCommonFailHandlers()
	helpers.SetupLog()
	junitReporter := reporters.NewJUnitReporter("junit.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "Consul+Vault E2e Suite", []Reporter{junitReporter})
}
