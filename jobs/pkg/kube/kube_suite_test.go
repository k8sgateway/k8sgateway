package kube_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/solo-io/gloo/test/gomega"
)

func TestKube(t *testing.T) {
	SetAsyncAssertionDefaults(AsyncAssertionDefaults{})
	RegisterFailHandler(Fail)

	RunSpecs(t, "Kube Suite")
}
