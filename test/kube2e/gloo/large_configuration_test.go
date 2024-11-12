package gloo_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/gloosnapshot"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

var _ = Describe("Large configuration", func() {
	var (
		glooResources *gloosnapshot.ApiSnapshot
	)

	BeforeEach(func() {
		glooResources = &gloosnapshot.ApiSnapshot{}
	})

	JustBeforeEach(func() {
		err := snapshotWriter.WriteSnapshot(glooResources, clients.WriteOpts{
			Ctx:               ctx,
			OverwriteExisting: false,
		})
		Expect(err).NotTo(HaveOccurred())
	})

	JustAfterEach(func() {
		err := snapshotWriter.DeleteSnapshot(glooResources, clients.DeleteOpts{
			Ctx:            ctx,
			IgnoreNotExist: true,
		})
		Expect(err).NotTo(HaveOccurred())
	})

	// We have a customer that was experiencing issue with full validation of a large configuration.
	// They were seeing an `argument list too long` error.
	// This test is to ensure that we can create a large configuration that is representative of their use case.
	It("should be able to create a large configuration", func() {
		err := testHelper.ApplyFile(ctx, testHelper.RootDir+"/test/kube2e/gloo/artifacts/large-configuration.yaml")
		Expect(err).NotTo(HaveOccurred())
	})
})
