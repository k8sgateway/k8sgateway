package directresponse_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDirectResponseRoute(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DirectResponseRoute Suite")
}
