package gloo_test

import (
	"time"

	"github.com/hashicorp/consul/api"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/k8s-utils/kubeutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	corecache "github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/cache"
	"github.com/solo-io/solo-kit/test/helpers"
	kubev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/solo-io/gloo/projects/gloo/pkg/bootstrap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// Kubernetes tests for artifact client from projects/gloo/pkg/bootstrap
var _ = Describe("Artifact Client", func() {

	var (
		testNamespace string
		cfg           *rest.Config

		kubeClient    kubernetes.Interface
		kubeCoreCache corecache.KubeCoreCache
	)

	BeforeEach(func() {
		var err error

		testNamespace = helpers.RandString(8)
		kubeClient = resourceClientset.KubeClients()

		cfg, err = kubeutils.GetConfig("", "")
		Expect(err).NotTo(HaveOccurred())

		_, err = kubeClient.CoreV1().Namespaces().Create(ctx, &kubev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: testNamespace,
			},
		}, metav1.CreateOptions{})
		Expect(err).NotTo(HaveOccurred())
		kubeCoreCache, err = corecache.NewKubeCoreCacheWithOptions(ctx, kubeClient, time.Hour, []string{testNamespace})
		Expect(err).NotTo(HaveOccurred())

	})
	AfterEach(func() {
		err := kubeClient.CoreV1().Namespaces().Delete(ctx, testNamespace, metav1.DeleteOptions{})
		Expect(err).NotTo(HaveOccurred())
	})

	Context("artifacts as config maps", func() {

		var (
			artifactClient v1.ArtifactClient
		)

		BeforeEach(func() {
			_, err := kubeClient.CoreV1().ConfigMaps(testNamespace).Create(ctx,
				&kubev1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Name: "cfg"},
					Data: map[string]string{
						"test": "data",
					},
				}, metav1.CreateOptions{})
			Expect(err).NotTo(HaveOccurred())

			settings := &v1.Settings{
				ArtifactSource:  &v1.Settings_KubernetesArtifactSource{},
				WatchNamespaces: []string{testNamespace},
			}

			factory, err := ArtifactFactoryForSettings(ctx,
				settings,
				nil,
				&cfg,
				&kubeClient,
				&kubeCoreCache,
				&api.Client{},
				"artifacts")
			Expect(err).NotTo(HaveOccurred())
			artifactClient, err = v1.NewArtifactClient(ctx, factory)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should work with artifacts", func() {
			artifact, err := artifactClient.Read(testNamespace, "cfg", clients.ReadOpts{Ctx: ctx})
			Expect(err).NotTo(HaveOccurred())
			Expect(artifact.GetMetadata().Name).To(Equal("cfg"))
			Expect(artifact.Data["test"]).To(Equal("data"))
		})
	})
})
