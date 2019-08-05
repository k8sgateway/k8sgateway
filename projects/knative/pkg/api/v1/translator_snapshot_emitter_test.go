// Code generated by solo-kit. DO NOT EDIT.

// +build solokit

package v1

import (
	"context"
	"os"
	"time"

	gloo_solo_io "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	github_com_solo_io_gloo_projects_knative_pkg_api_external_knative "github.com/solo-io/gloo/projects/knative/pkg/api/external/knative"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/go-utils/kubeutils"
	"github.com/solo-io/go-utils/log"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"
	"github.com/solo-io/solo-kit/test/helpers"
	"k8s.io/client-go/kubernetes"

	// Needed to run tests in GKE
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	// From https://github.com/kubernetes/client-go/blob/53c7adfd0294caa142d961e1f780f74081d5b15f/examples/out-of-cluster-client-configuration/main.go#L31
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

var _ = Describe("V1Emitter", func() {
	if os.Getenv("RUN_KUBE_TESTS") != "1" {
		log.Printf("This test creates kubernetes resources and is disabled by default. To enable, set RUN_KUBE_TESTS=1 in your env.")
		return
	}
	var (
		namespace1    string
		namespace2    string
		name1, name2  = "angela" + helpers.RandString(3), "bob" + helpers.RandString(3)
		kube          kubernetes.Interface
		emitter       TranslatorEmitter
		secretClient  gloo_solo_io.SecretClient
		ingressClient github_com_solo_io_gloo_projects_knative_pkg_api_external_knative.IngressClient
	)

	BeforeEach(func() {
		namespace1 = helpers.RandString(8)
		namespace2 = helpers.RandString(8)
		kube = helpers.MustKubeClient()
		err := kubeutils.CreateNamespacesInParallel(kube, namespace1, namespace2)
		Expect(err).NotTo(HaveOccurred())
		// Secret Constructor
		secretClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}

		secretClient, err = gloo_solo_io.NewSecretClient(secretClientFactory)
		Expect(err).NotTo(HaveOccurred())
		// Ingress Constructor
		ingressClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}

		ingressClient, err = github_com_solo_io_gloo_projects_knative_pkg_api_external_knative.NewIngressClient(ingressClientFactory)
		Expect(err).NotTo(HaveOccurred())
		emitter = NewTranslatorEmitter(secretClient, ingressClient)
	})
	AfterEach(func() {
		err := kubeutils.DeleteNamespacesInParallelBlocking(kube, namespace1, namespace2)
		Expect(err).NotTo(HaveOccurred())
	})
	It("tracks snapshots on changes to any resource", func() {
		ctx := context.Background()
		err := emitter.Register()
		Expect(err).NotTo(HaveOccurred())

		snapshots, errs, err := emitter.Snapshots([]string{namespace1, namespace2}, clients.WatchOpts{
			Ctx:         ctx,
			RefreshRate: time.Second,
		})
		Expect(err).NotTo(HaveOccurred())

		var snap *TranslatorSnapshot

		/*
			Secret
		*/

		assertSnapshotSecrets := func(expectSecrets gloo_solo_io.SecretList, unexpectSecrets gloo_solo_io.SecretList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectSecrets {
						if _, err := snap.Secrets.Find(expected.GetMetadata().Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectSecrets {
						if _, err := snap.Secrets.Find(unexpected.GetMetadata().Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := secretClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := secretClient.List(namespace2, clients.ListOpts{})
					combined := append(nsList1, nsList2...)
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		secret1a, err := secretClient.Write(gloo_solo_io.NewSecret(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		secret1b, err := secretClient.Write(gloo_solo_io.NewSecret(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotSecrets(gloo_solo_io.SecretList{secret1a, secret1b}, nil)
		secret2a, err := secretClient.Write(gloo_solo_io.NewSecret(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		secret2b, err := secretClient.Write(gloo_solo_io.NewSecret(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotSecrets(gloo_solo_io.SecretList{secret1a, secret1b, secret2a, secret2b}, nil)

		err = secretClient.Delete(secret2a.GetMetadata().Namespace, secret2a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = secretClient.Delete(secret2b.GetMetadata().Namespace, secret2b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotSecrets(gloo_solo_io.SecretList{secret1a, secret1b}, gloo_solo_io.SecretList{secret2a, secret2b})

		err = secretClient.Delete(secret1a.GetMetadata().Namespace, secret1a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = secretClient.Delete(secret1b.GetMetadata().Namespace, secret1b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotSecrets(nil, gloo_solo_io.SecretList{secret1a, secret1b, secret2a, secret2b})

		/*
			Ingress
		*/

		assertSnapshotingresses := func(expectingresses github_com_solo_io_gloo_projects_knative_pkg_api_external_knative.IngressList, unexpectingresses github_com_solo_io_gloo_projects_knative_pkg_api_external_knative.IngressList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectingresses {
						if _, err := snap.Ingresses.Find(expected.GetMetadata().Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectingresses {
						if _, err := snap.Ingresses.Find(unexpected.GetMetadata().Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := ingressClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := ingressClient.List(namespace2, clients.ListOpts{})
					combined := append(nsList1, nsList2...)
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		ingress1a, err := ingressClient.Write(github_com_solo_io_gloo_projects_knative_pkg_api_external_knative.NewIngress(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		ingress1b, err := ingressClient.Write(github_com_solo_io_gloo_projects_knative_pkg_api_external_knative.NewIngress(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotingresses(github_com_solo_io_gloo_projects_knative_pkg_api_external_knative.IngressList{ingress1a, ingress1b}, nil)
		ingress2a, err := ingressClient.Write(github_com_solo_io_gloo_projects_knative_pkg_api_external_knative.NewIngress(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		ingress2b, err := ingressClient.Write(github_com_solo_io_gloo_projects_knative_pkg_api_external_knative.NewIngress(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotingresses(github_com_solo_io_gloo_projects_knative_pkg_api_external_knative.IngressList{ingress1a, ingress1b, ingress2a, ingress2b}, nil)

		err = ingressClient.Delete(ingress2a.GetMetadata().Namespace, ingress2a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = ingressClient.Delete(ingress2b.GetMetadata().Namespace, ingress2b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotingresses(github_com_solo_io_gloo_projects_knative_pkg_api_external_knative.IngressList{ingress1a, ingress1b}, github_com_solo_io_gloo_projects_knative_pkg_api_external_knative.IngressList{ingress2a, ingress2b})

		err = ingressClient.Delete(ingress1a.GetMetadata().Namespace, ingress1a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = ingressClient.Delete(ingress1b.GetMetadata().Namespace, ingress1b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotingresses(nil, github_com_solo_io_gloo_projects_knative_pkg_api_external_knative.IngressList{ingress1a, ingress1b, ingress2a, ingress2b})
	})
	It("tracks snapshots on changes to any resource using AllNamespace", func() {
		ctx := context.Background()
		err := emitter.Register()
		Expect(err).NotTo(HaveOccurred())

		snapshots, errs, err := emitter.Snapshots([]string{""}, clients.WatchOpts{
			Ctx:         ctx,
			RefreshRate: time.Second,
		})
		Expect(err).NotTo(HaveOccurred())

		var snap *TranslatorSnapshot

		/*
			Secret
		*/

		assertSnapshotSecrets := func(expectSecrets gloo_solo_io.SecretList, unexpectSecrets gloo_solo_io.SecretList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectSecrets {
						if _, err := snap.Secrets.Find(expected.GetMetadata().Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectSecrets {
						if _, err := snap.Secrets.Find(unexpected.GetMetadata().Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := secretClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := secretClient.List(namespace2, clients.ListOpts{})
					combined := append(nsList1, nsList2...)
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		secret1a, err := secretClient.Write(gloo_solo_io.NewSecret(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		secret1b, err := secretClient.Write(gloo_solo_io.NewSecret(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotSecrets(gloo_solo_io.SecretList{secret1a, secret1b}, nil)
		secret2a, err := secretClient.Write(gloo_solo_io.NewSecret(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		secret2b, err := secretClient.Write(gloo_solo_io.NewSecret(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotSecrets(gloo_solo_io.SecretList{secret1a, secret1b, secret2a, secret2b}, nil)

		err = secretClient.Delete(secret2a.GetMetadata().Namespace, secret2a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = secretClient.Delete(secret2b.GetMetadata().Namespace, secret2b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotSecrets(gloo_solo_io.SecretList{secret1a, secret1b}, gloo_solo_io.SecretList{secret2a, secret2b})

		err = secretClient.Delete(secret1a.GetMetadata().Namespace, secret1a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = secretClient.Delete(secret1b.GetMetadata().Namespace, secret1b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotSecrets(nil, gloo_solo_io.SecretList{secret1a, secret1b, secret2a, secret2b})

		/*
			Ingress
		*/

		assertSnapshotingresses := func(expectingresses github_com_solo_io_gloo_projects_knative_pkg_api_external_knative.IngressList, unexpectingresses github_com_solo_io_gloo_projects_knative_pkg_api_external_knative.IngressList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectingresses {
						if _, err := snap.Ingresses.Find(expected.GetMetadata().Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectingresses {
						if _, err := snap.Ingresses.Find(unexpected.GetMetadata().Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := ingressClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := ingressClient.List(namespace2, clients.ListOpts{})
					combined := append(nsList1, nsList2...)
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		ingress1a, err := ingressClient.Write(github_com_solo_io_gloo_projects_knative_pkg_api_external_knative.NewIngress(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		ingress1b, err := ingressClient.Write(github_com_solo_io_gloo_projects_knative_pkg_api_external_knative.NewIngress(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotingresses(github_com_solo_io_gloo_projects_knative_pkg_api_external_knative.IngressList{ingress1a, ingress1b}, nil)
		ingress2a, err := ingressClient.Write(github_com_solo_io_gloo_projects_knative_pkg_api_external_knative.NewIngress(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		ingress2b, err := ingressClient.Write(github_com_solo_io_gloo_projects_knative_pkg_api_external_knative.NewIngress(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotingresses(github_com_solo_io_gloo_projects_knative_pkg_api_external_knative.IngressList{ingress1a, ingress1b, ingress2a, ingress2b}, nil)

		err = ingressClient.Delete(ingress2a.GetMetadata().Namespace, ingress2a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = ingressClient.Delete(ingress2b.GetMetadata().Namespace, ingress2b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotingresses(github_com_solo_io_gloo_projects_knative_pkg_api_external_knative.IngressList{ingress1a, ingress1b}, github_com_solo_io_gloo_projects_knative_pkg_api_external_knative.IngressList{ingress2a, ingress2b})

		err = ingressClient.Delete(ingress1a.GetMetadata().Namespace, ingress1a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = ingressClient.Delete(ingress1b.GetMetadata().Namespace, ingress1b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotingresses(nil, github_com_solo_io_gloo_projects_knative_pkg_api_external_knative.IngressList{ingress1a, ingress1b, ingress2a, ingress2b})
	})
})
