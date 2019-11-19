package install_test

import (
	"fmt"
	"path/filepath"

	"github.com/solo-io/go-utils/testutils/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/testutils"
)

var _ = Describe("Install", func() {

	It("shouldn't get errors for gateway dry run", func() {
		_, err := testutils.GlooctlOut(fmt.Sprintf("install gateway --file %s --dry-run", file))
		Expect(err).NotTo(HaveOccurred())
	})

	It("shouldn't get errors for gateway upgrade dry run", func() {
		_, err := testutils.GlooctlOut(fmt.Sprintf("install gateway --file %s --dry-run --upgrade", file))
		Expect(err).NotTo(HaveOccurred())
	})

	It("shouldn't get errors for gateway dry run with multiple values", func() {
		outputYaml, err := testutils.GlooctlOut(fmt.Sprintf("install gateway --file %s --dry-run --values %s,%s ", file, values1, values2))
		Expect(err).NotTo(HaveOccurred())
		// Test that the values are being merged as we expect
		Expect(outputYaml).To(ContainSubstring("test-namespace-2\n"))
	})

	const licenseKey = "--license-key=fake-license-key"

	It("shouldn't get errors for enterprise dry run", func() {
		_, err := testutils.GlooctlOut(fmt.Sprintf("install gateway enterprise --file %s --dry-run %s", file, licenseKey))
		Expect(err).NotTo(HaveOccurred())
	})

	It("shouldn't get errors for enterprise upgrade dry run", func() {
		_, err := testutils.GlooctlOut(fmt.Sprintf("install gateway enterprise --file %s --dry-run --upgrade", file))
		Expect(err).NotTo(HaveOccurred())
	})

	It("shouldn't get errors for knative dry run", func() {
		_, err := testutils.GlooctlOut(fmt.Sprintf("install knative --file %s --dry-run", file))
		Expect(err).NotTo(HaveOccurred())
	})

	It("shouldn't get errors for ingress dry run", func() {
		_, err := testutils.GlooctlOut(fmt.Sprintf("install ingress --file %s --dry-run", file))
		Expect(err).NotTo(HaveOccurred())
	})

	It("should error when not overriding helm chart in dev mode", func() {
		_, err := testutils.GlooctlOut("install ingress")
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("installing gloo in ingress mode: you must provide a Gloo Helm chart URI via the 'file' option when running an unreleased version of glooctl"))
	})

	It("should error when not providing file with valid extension", func() {
		_, err := testutils.GlooctlOut("install gateway --file foo")
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("installing gloo in gateway mode: unsupported file extension for Helm chart URI: [foo]. Extension must either be .tgz or .tar.gz"))
	})

	It("should error when not providing valid file", func() {
		_, err := testutils.GlooctlOut("install gateway --file foo.tgz")
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("installing gloo in gateway mode: retrieving gloo helm chart archive: opening file"))
	})

	It("should not error when providing the admin console flag", func() {
		// This test fetches the corresponding GlooE helm chart, thus it needs the version that gets linked
		// into the glooctl binary at build time
		out, err := exec.RunCommandOutput(RootDir, true, filepath.Join("_output", "glooctl"), "install", "gateway", "--dry-run", "--with-admin-console")
		Expect(err).NotTo(HaveOccurred())
		Expect(out).To(ContainSubstring("kind: Namespace"))
	})

})
