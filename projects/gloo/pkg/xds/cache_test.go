package xds_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/xds"
)

var _ = Describe("Cache", func() {

	It("SnapshotCacheKeys returns the keys formatted correctly", func() {
		namespace1, namespace2, name1, name2 := "namespace1", "namespace2", "name1", "name2"
		proxies := []*v1.Proxy{
			v1.NewProxy(namespace1, name1),
			v1.NewProxy(namespace2, name2),
		}
		expectedKeys := []string{fmt.Sprintf("%v~%v", namespace1, name1), fmt.Sprintf("%v~%v", namespace2, name2)}
		actualKeys := xds.SnapshotCacheKeys(proxies)
		Expect(actualKeys).To(BeEquivalentTo(expectedKeys))
	})
})
