package deployer_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDeployer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Deployer Suite")
}
