package add_test

import (
	"log"
	"os"

	"github.com/hashicorp/consul/api"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/testutils"
)

var _ = Describe("Add", func() {
	if os.Getenv("RUN_CONSUL_TESTS") != "1" {
		log.Print("This test downloads and runs consul and is disabled by default. To enable, set RUN_CONSUL_TESTS=1 in your env.")
		return
	}

	BeforeEach(func() {
		helpers.UseDefaultClients()
		var err error
		// Start Consul
		consulInstance, err = consulFactory.NewConsulInstance()
		Expect(err).NotTo(HaveOccurred())
		err = consulInstance.Run()
		Expect(err).NotTo(HaveOccurred())
		// wait for consul to start
		Eventually(func() error {
			_, err := client.KV().Put(&api.KVPair{Key: "test"}, nil)
			return err
		}).ShouldNot(HaveOccurred())
	})

	AfterEach(func() {
		consulInstance.Clean()

		helpers.UseDefaultClients()
	})

	Context("consul storage backend", func() {
		It("works", func() {
			err := testutils.Glooctl("add route " +
				"--path-prefix / " +
				"--dest-name petstore " +
				"--prefix-rewrite /api/pets " +
				"--use-consul")
			Expect(err).NotTo(HaveOccurred())
			kv, _, err := client.KV().Get("gloo/gateway.solo.io/v1/VirtualService/gloo-system/default", nil)
			Expect(err).NotTo(HaveOccurred())
			Expect(kv).NotTo(BeNil())
		})
	})
})
