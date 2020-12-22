package csrf_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCSRF(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CSRF Policy Suite")
}
