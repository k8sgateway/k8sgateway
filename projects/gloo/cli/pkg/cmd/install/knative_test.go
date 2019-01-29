package install_test

import (
	"path/filepath"
	"time"

	"k8s.io/client-go/kubernetes"

	helpers2 "github.com/solo-io/gloo/test/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/testutils"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"

	"github.com/solo-io/go-utils/kubeutils"
	kubeerrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("Knative", func() {
	It("should install gloo and knative", func() {
		err := testutils.Glooctl("install knative " +
			"--file " + filepath.Join(helpers2.GlooInstallDir(), "gloo-knative.yaml") + " " +
			"--knative-crds-manifest " + filepath.Join(helpers2.GlooInstallDir(), "integrations", "knative-crds-0.3.0.yaml") + " " +
			"--knative-install-manifest " + filepath.Join(helpers2.GlooInstallDir(), "integrations", "knative-no-istio-0.3.0.yaml") + " ",
		)
		Expect(err).NotTo(HaveOccurred())

		// when we see that discovery has created an upstream for gateway-proxy, we're good
		var us *v1.Upstream
		Eventually(func() (*v1.Upstream, error) {
			u, err := helpers.MustUpstreamClient().Read("gloo-system", "gloo-system-clusteringress-proxy-80", clients.ReadOpts{})
			us = u
			return us, err
		}, time.Minute).Should(Not(BeNil()))
		Eventually(func() (*v1.Upstream, error) {
			u, err := helpers.MustUpstreamClient().Read("gloo-system", "knative-serving-controller-9090", clients.ReadOpts{})
			us = u
			return us, err
		}, time.Minute).Should(Not(BeNil()))

		err = testutils.Glooctl("uninstall")
		Expect(err).NotTo(HaveOccurred())

		// uninstall should delete both namespaces
		Eventually(func() bool {
			cfg, err := kubeutils.GetConfig("", "")
			Expect(err).NotTo(HaveOccurred())
			kube, err := kubernetes.NewForConfig(cfg)
			Expect(err).NotTo(HaveOccurred())
			_, err = kube.CoreV1().Namespaces().Get("gloo-system", metav1.GetOptions{})
			return err != nil && kubeerrs.IsNotFound(err)
		}, time.Minute).Should(BeTrue())
		Eventually(func() bool {
			cfg, err := kubeutils.GetConfig("", "")
			Expect(err).NotTo(HaveOccurred())
			kube, err := kubernetes.NewForConfig(cfg)
			Expect(err).NotTo(HaveOccurred())
			_, err = kube.CoreV1().Namespaces().Get("knative-serving", metav1.GetOptions{})
			return err != nil && kubeerrs.IsNotFound(err)
		}, time.Minute).Should(BeTrue())
	})
})
