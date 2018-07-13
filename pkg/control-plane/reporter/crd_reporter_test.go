package reporter_test

import (
	"os"
	"path/filepath"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"github.com/solo-io/gloo/pkg/storage"
	"github.com/solo-io/gloo/pkg/storage/crd"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/solo-io/gloo/pkg/api/types/v1"
	. "github.com/solo-io/gloo/pkg/control-plane/reporter"
	"github.com/solo-io/gloo/pkg/log"
	. "github.com/solo-io/gloo/test/helpers"
)

var _ = Describe("CrdReporter", func() {
	if os.Getenv("RUN_KUBE_TESTS") != "1" {
		log.Printf("This test creates kubernetes resources and is disabled by default. To enable, set RUN_KUBE_TESTS=1 in your env.")
		return
	}
	var (
		masterUrl, kubeconfigPath string
		namespace                 string
		rptr                      Interface
	)
	BeforeEach(func() {
		namespace = RandString(8)
		err := SetupKubeForTest(namespace)
		Must(err)
		kubeconfigPath = filepath.Join(os.Getenv("HOME"), ".kube", "config")
		masterUrl = ""
	})
	AfterEach(func() {
		TeardownKube(namespace)
	})
	Describe("writereports", func() {
		var (
			glooClient      storage.Interface
			reports         []ConfigObjectReport
			upstreams       []*v1.Upstream
			virtualServices []*v1.VirtualService
			roles           []*v1.Role
			attributes      []*v1.Attribute
		)
		Context("writes status reports for cfg crds with 0 errors", func() {
			BeforeEach(func() {
				reports = nil
				cfg, err := clientcmd.BuildConfigFromFlags(masterUrl, kubeconfigPath)
				Expect(err).NotTo(HaveOccurred())
				glooClient, err = crd.NewStorage(cfg, namespace, time.Second)
				Expect(err).NotTo(HaveOccurred())
				rptr = NewReporter(glooClient)

				testCfg := NewTestConfig()
				upstreams = testCfg.Upstreams
				var storables []v1.ConfigObject
				for _, us := range upstreams {
					_, err := glooClient.V1().Upstreams().Create(us)
					Expect(err).NotTo(HaveOccurred())
					storables = append(storables, us)
				}
				virtualServices = testCfg.VirtualServices
				for _, vService := range virtualServices {
					_, err := glooClient.V1().VirtualServices().Create(vService)
					Expect(err).NotTo(HaveOccurred())
					storables = append(storables, vService)
				}
				roles = testCfg.Roles
				for _, role := range roles {
					_, err := glooClient.V1().Roles().Create(role)
					Expect(err).NotTo(HaveOccurred())

					storables = append(storables, role)
				}
				attributes = testCfg.Attributes
				for _, attribute := range attributes {
					_, err := glooClient.V1().Attributes().Create(attribute)
					Expect(err).NotTo(HaveOccurred())

					storables = append(storables, attribute)
				}
				for _, storable := range storables {
					reports = append(reports, ConfigObjectReport{
						CfgObject: storable,
						Err:       nil,
					})
				}
			})

			It("writes an acceptance status for each crd", func() {
				err := rptr.WriteReports(reports)
				Expect(err).NotTo(HaveOccurred())
				updatedUpstreams, err := glooClient.V1().Upstreams().List()
				Expect(err).NotTo(HaveOccurred())
				Expect(updatedUpstreams).To(HaveLen(len(upstreams)))
				for _, updatedUpstream := range updatedUpstreams {
					Expect(updatedUpstream.Status.State).To(Equal(v1.Status_Accepted))
				}
				updatedvServices, err := glooClient.V1().VirtualServices().List()
				Expect(err).NotTo(HaveOccurred())
				Expect(updatedvServices).To(HaveLen(len(virtualServices)))
				for _, updatedvService := range updatedvServices {
					Expect(updatedvService.Status.State).To(Equal(v1.Status_Accepted))
				}
				updatedroles, err := glooClient.V1().Roles().List()
				Expect(err).NotTo(HaveOccurred())
				Expect(updatedroles).To(HaveLen(len(roles)))
				for _, updatedrole := range updatedroles {
					Expect(updatedrole.Status.State).To(Equal(v1.Status_Accepted))
				}
				updatedattributes, err := glooClient.V1().Attributes().List()
				Expect(err).NotTo(HaveOccurred())
				Expect(updatedattributes).To(HaveLen(len(attributes)))
				for _, updatedattribute := range updatedattributes {
					Expect(updatedattribute.Status.State).To(Equal(v1.Status_Accepted))
				}
			})
		})
		Context("writes status reports for cfg crds with SOME errors", func() {
			BeforeEach(func() {
				reports = nil
				cfg, err := clientcmd.BuildConfigFromFlags(masterUrl, kubeconfigPath)
				Expect(err).NotTo(HaveOccurred())
				glooClient, err = crd.NewStorage(cfg, namespace, time.Second)
				Expect(err).NotTo(HaveOccurred())
				rptr = NewReporter(glooClient)

				testCfg := NewTestConfig()
				upstreams = testCfg.Upstreams
				var storables []v1.ConfigObject
				for _, us := range upstreams {
					_, err := glooClient.V1().Upstreams().Create(us)
					Expect(err).NotTo(HaveOccurred())
					storables = append(storables, us)
				}
				virtualServices = testCfg.VirtualServices
				for _, vService := range virtualServices {
					_, err := glooClient.V1().VirtualServices().Create(vService)
					Expect(err).NotTo(HaveOccurred())
					storables = append(storables, vService)
				}
				roles = testCfg.Roles
				for _, role := range roles {
					_, err := glooClient.V1().Roles().Create(role)
					Expect(err).NotTo(HaveOccurred())
					storables = append(storables, role)
				}
				attributes = testCfg.Attributes
				for _, attribute := range attributes {
					_, err := glooClient.V1().Attributes().Create(attribute)
					Expect(err).NotTo(HaveOccurred())
					storables = append(storables, attribute)
				}
				for _, storable := range storables {
					reports = append(reports, ConfigObjectReport{
						CfgObject: storable,
						Err:       errors.New("oh no an error what did u do!"),
					})
				}
			})

			It("writes an rejected status for each crd", func() {
				err := rptr.WriteReports(reports)
				Expect(err).NotTo(HaveOccurred())
				updatedUpstreams, err := glooClient.V1().Upstreams().List()
				Expect(err).NotTo(HaveOccurred())
				Expect(updatedUpstreams).To(HaveLen(len(upstreams)))
				for _, updatedUpstream := range updatedUpstreams {
					Expect(updatedUpstream.Status.State).To(Equal(v1.Status_Rejected))
				}
				updatedvServices, err := glooClient.V1().VirtualServices().List()
				Expect(err).NotTo(HaveOccurred())
				Expect(updatedvServices).To(HaveLen(len(virtualServices)))
				for _, updatedvService := range updatedvServices {
					Expect(updatedvService.Status.State).To(Equal(v1.Status_Rejected))
				}
				updatedroles, err := glooClient.V1().Roles().List()
				Expect(err).NotTo(HaveOccurred())
				Expect(updatedroles).To(HaveLen(len(roles)))
				for _, updatedrole := range updatedroles {
					Expect(updatedrole.Status.State).To(Equal(v1.Status_Rejected))
				}
				updatedattributes, err := glooClient.V1().Attributes().List()
				Expect(err).NotTo(HaveOccurred())
				Expect(updatedattributes).To(HaveLen(len(attributes)))
				for _, updatedattribute := range updatedattributes {
					Expect(updatedattribute.Status.State).To(Equal(v1.Status_Rejected))
				}
			})
		})

		Context("creates the role crd if writing a report for a role that doesn't exist", func() {
			BeforeEach(func() {
				reports = nil
				cfg, err := clientcmd.BuildConfigFromFlags(masterUrl, kubeconfigPath)
				Expect(err).NotTo(HaveOccurred())
				glooClient, err = crd.NewStorage(cfg, namespace, time.Second)
				Expect(err).NotTo(HaveOccurred())
				rptr = NewReporter(glooClient)

				var storables []v1.ConfigObject
				roles = NewTestConfig().Roles
				for _, role := range roles {
					storables = append(storables, role)
				}
				for _, storable := range storables {
					reports = append(reports, ConfigObjectReport{
						CfgObject: storable,
						Err:       errors.New("oh no an error what did u do!"),
					})
				}
			})

			It("writes an rejected status for each crd", func() {
				err := rptr.WriteReports(reports)
				Expect(err).NotTo(HaveOccurred())
				roleList, err := glooClient.V1().Roles().List()
				Expect(err).NotTo(HaveOccurred())
				// zero out fields we dont care about & expect to be different
				for _, role := range roleList {
					role.Metadata = &v1.Metadata{}
				}
				for _, role := range roles {
					role.Status = &v1.Status{
						State:  v1.Status_Rejected,
						Reason: "oh no an error what did u do!",
					}
					role.Metadata = &v1.Metadata{}
					Expect(roleList).To(ContainElement(role))
				}
			})
		})
	})
})
