package get_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/testutils"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/static"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

var _ = Describe("Upstream", func() {

	BeforeEach(func() {
		helpers.UseMemoryClients()
	})

	getUpstream := func(name string) *v1.Upstream {
		up, err := helpers.MustUpstreamClient().Read("gloo-system", name, clients.ReadOpts{})
		Expect(err).NotTo(HaveOccurred())

		return up
	}

	tableOutput := `+--------------------+--------+---------+---------------------------------+
|      UPSTREAM      |  TYPE  | STATUS  |             DETAILS             |
+--------------------+--------+---------+---------------------------------+
| jsonplaceholder-80 | Static | Pending | hosts:                          |
|                    |        |         | -                               |
|                    |        |         | jsonplaceholder.typicode.com:80 |
|                    |        |         |                                 |
+--------------------+--------+---------+---------------------------------+`

	kubeYamlOutput := `apiVersion: gloo.solo.io/v1
kind: Upstream
metadata:
  creationTimestamp: null
  name: jsonplaceholder-80
  namespace: gloo-system
  resourceVersion: "2"
spec:
  upstreamSpec:
    static:
      hosts:
      - addr: jsonplaceholder.typicode.com
        port: 80
status: {}
`

	yamlOutput := `---
metadata:
  name: jsonplaceholder-80
  namespace: gloo-system
  resourceVersion: "2"
status: {}
upstreamSpec:
  static:
    hosts:
    - addr: jsonplaceholder.typicode.com
      port: 80
`

	Context("default output should be -o table", func() {
		It("should override/allow -o flags as expected", func() {
			output, err := testutils.GlooctlOut("create upstream static jsonplaceholder-80 --static-hosts jsonplaceholder.typicode.com:80")
			Expect(err).NotTo(HaveOccurred())
			Expect(output).To(Equal(tableOutput))

			// make sure that we created the upstream that we intended
			up := getUpstream("jsonplaceholder-80")
			staticSpec := up.UpstreamSpec.UpstreamType.(*v1.UpstreamSpec_Static).Static
			expectedHosts := []*static.Host{{Addr: "jsonplaceholder.typicode.com", Port: 80}}
			Expect(staticSpec.Hosts).To(Equal(expectedHosts))

			By("should default to -o table")
			output, err = testutils.GlooctlOut("get upstreams")
			Expect(err).NotTo(HaveOccurred())
			Expect(output).To(Equal(tableOutput))

			By("should respect (unnecessary) -o table flag")
			output, err = testutils.GlooctlOut("get upstreams -o table")
			Expect(err).NotTo(HaveOccurred())
			Expect(output).To(Equal(tableOutput))

			By("should respect -o yaml flag")
			output, err = testutils.GlooctlOut("get upstreams -o yaml")
			Expect(err).NotTo(HaveOccurred())
			Expect(output).To(Equal(yamlOutput))

			By("should respect -o kube-yaml flag")
			output, err = testutils.GlooctlOut("get upstreams -o kube-yaml")
			Expect(err).NotTo(HaveOccurred())
			Expect(output).To(Equal(kubeYamlOutput))
		})
	})
})
