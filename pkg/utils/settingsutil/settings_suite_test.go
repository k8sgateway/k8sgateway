package settingsutil_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSettings(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Settings Suite")
}
