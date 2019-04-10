// Code generated by solo-kit. DO NOT EDIT.

// +build solokit

package v1

import (
	"context"
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/go-utils/kubeutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	kuberc "github.com/solo-io/solo-kit/pkg/api/v1/clients/kube"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"
	"github.com/solo-io/solo-kit/pkg/utils/log"
	"github.com/solo-io/solo-kit/test/helpers"
	"github.com/solo-io/solo-kit/test/setup"
	"k8s.io/client-go/rest"

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
		namespace1     string
		namespace2     string
		name1, name2   = "angela" + helpers.RandString(3), "bob" + helpers.RandString(3)
		cfg            *rest.Config
		emitter        DiscoveryEmitter
		upstreamClient UpstreamClient
		secretClient   SecretClient
	)

	BeforeEach(func() {
		namespace1 = helpers.RandString(8)
		namespace2 = helpers.RandString(8)
		var err error
		cfg, err = kubeutils.GetConfig("", "")
		Expect(err).NotTo(HaveOccurred())
		err = setup.SetupKubeForTest(namespace1)
		Expect(err).NotTo(HaveOccurred())
		err = setup.SetupKubeForTest(namespace2)
		Expect(err).NotTo(HaveOccurred())
		// Upstream Constructor
		upstreamClientFactory := &factory.KubeResourceClientFactory{
			Crd:         UpstreamCrd,
			Cfg:         cfg,
			SharedCache: kuberc.NewKubeCache(context.TODO()),
		}

		upstreamClient, err = NewUpstreamClient(upstreamClientFactory)
		Expect(err).NotTo(HaveOccurred())
		// Secret Constructor
		secretClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}

		secretClient, err = NewSecretClient(secretClientFactory)
		Expect(err).NotTo(HaveOccurred())
		emitter = NewDiscoveryEmitter(upstreamClient, secretClient)
	})
	AfterEach(func() {
		setup.TeardownKube(namespace1)
		setup.TeardownKube(namespace2)
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

		var snap *DiscoverySnapshot

		/*
			Upstream
		*/

		assertSnapshotUpstreams := func(expectUpstreams UpstreamList, unexpectUpstreams UpstreamList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectUpstreams {
						if _, err := snap.Upstreams.List().Find(expected.GetMetadata().Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectUpstreams {
						if _, err := snap.Upstreams.List().Find(unexpected.GetMetadata().Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := upstreamClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := upstreamClient.List(namespace2, clients.ListOpts{})
					combined := UpstreamsByNamespace{
						namespace1: nsList1,
						namespace2: nsList2,
					}
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		upstream1a, err := upstreamClient.Write(NewUpstream(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		upstream1b, err := upstreamClient.Write(NewUpstream(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreams(UpstreamList{upstream1a, upstream1b}, nil)
		upstream2a, err := upstreamClient.Write(NewUpstream(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		upstream2b, err := upstreamClient.Write(NewUpstream(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreams(UpstreamList{upstream1a, upstream1b, upstream2a, upstream2b}, nil)

		err = upstreamClient.Delete(upstream2a.GetMetadata().Namespace, upstream2a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = upstreamClient.Delete(upstream2b.GetMetadata().Namespace, upstream2b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreams(UpstreamList{upstream1a, upstream1b}, UpstreamList{upstream2a, upstream2b})

		err = upstreamClient.Delete(upstream1a.GetMetadata().Namespace, upstream1a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = upstreamClient.Delete(upstream1b.GetMetadata().Namespace, upstream1b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreams(nil, UpstreamList{upstream1a, upstream1b, upstream2a, upstream2b})

		/*
			Secret
		*/

		assertSnapshotSecrets := func(expectSecrets SecretList, unexpectSecrets SecretList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectSecrets {
						if _, err := snap.Secrets.List().Find(expected.GetMetadata().Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectSecrets {
						if _, err := snap.Secrets.List().Find(unexpected.GetMetadata().Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := secretClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := secretClient.List(namespace2, clients.ListOpts{})
					combined := SecretsByNamespace{
						namespace1: nsList1,
						namespace2: nsList2,
					}
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		secret1a, err := secretClient.Write(NewSecret(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		secret1b, err := secretClient.Write(NewSecret(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotSecrets(SecretList{secret1a, secret1b}, nil)
		secret2a, err := secretClient.Write(NewSecret(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		secret2b, err := secretClient.Write(NewSecret(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotSecrets(SecretList{secret1a, secret1b, secret2a, secret2b}, nil)

		err = secretClient.Delete(secret2a.GetMetadata().Namespace, secret2a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = secretClient.Delete(secret2b.GetMetadata().Namespace, secret2b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotSecrets(SecretList{secret1a, secret1b}, SecretList{secret2a, secret2b})

		err = secretClient.Delete(secret1a.GetMetadata().Namespace, secret1a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = secretClient.Delete(secret1b.GetMetadata().Namespace, secret1b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotSecrets(nil, SecretList{secret1a, secret1b, secret2a, secret2b})
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

		var snap *DiscoverySnapshot

		/*
			Upstream
		*/

		assertSnapshotUpstreams := func(expectUpstreams UpstreamList, unexpectUpstreams UpstreamList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectUpstreams {
						if _, err := snap.Upstreams.List().Find(expected.GetMetadata().Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectUpstreams {
						if _, err := snap.Upstreams.List().Find(unexpected.GetMetadata().Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := upstreamClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := upstreamClient.List(namespace2, clients.ListOpts{})
					combined := UpstreamsByNamespace{
						namespace1: nsList1,
						namespace2: nsList2,
					}
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		upstream1a, err := upstreamClient.Write(NewUpstream(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		upstream1b, err := upstreamClient.Write(NewUpstream(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreams(UpstreamList{upstream1a, upstream1b}, nil)
		upstream2a, err := upstreamClient.Write(NewUpstream(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		upstream2b, err := upstreamClient.Write(NewUpstream(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreams(UpstreamList{upstream1a, upstream1b, upstream2a, upstream2b}, nil)

		err = upstreamClient.Delete(upstream2a.GetMetadata().Namespace, upstream2a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = upstreamClient.Delete(upstream2b.GetMetadata().Namespace, upstream2b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreams(UpstreamList{upstream1a, upstream1b}, UpstreamList{upstream2a, upstream2b})

		err = upstreamClient.Delete(upstream1a.GetMetadata().Namespace, upstream1a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = upstreamClient.Delete(upstream1b.GetMetadata().Namespace, upstream1b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreams(nil, UpstreamList{upstream1a, upstream1b, upstream2a, upstream2b})

		/*
			Secret
		*/

		assertSnapshotSecrets := func(expectSecrets SecretList, unexpectSecrets SecretList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectSecrets {
						if _, err := snap.Secrets.List().Find(expected.GetMetadata().Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectSecrets {
						if _, err := snap.Secrets.List().Find(unexpected.GetMetadata().Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := secretClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := secretClient.List(namespace2, clients.ListOpts{})
					combined := SecretsByNamespace{
						namespace1: nsList1,
						namespace2: nsList2,
					}
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		secret1a, err := secretClient.Write(NewSecret(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		secret1b, err := secretClient.Write(NewSecret(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotSecrets(SecretList{secret1a, secret1b}, nil)
		secret2a, err := secretClient.Write(NewSecret(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		secret2b, err := secretClient.Write(NewSecret(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotSecrets(SecretList{secret1a, secret1b, secret2a, secret2b}, nil)

		err = secretClient.Delete(secret2a.GetMetadata().Namespace, secret2a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = secretClient.Delete(secret2b.GetMetadata().Namespace, secret2b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotSecrets(SecretList{secret1a, secret1b}, SecretList{secret2a, secret2b})

		err = secretClient.Delete(secret1a.GetMetadata().Namespace, secret1a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = secretClient.Delete(secret1b.GetMetadata().Namespace, secret1b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotSecrets(nil, SecretList{secret1a, secret1b, secret2a, secret2b})
	})
})
