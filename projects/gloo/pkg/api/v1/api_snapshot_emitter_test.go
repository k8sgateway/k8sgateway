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
	"github.com/solo-io/solo-kit/pkg/utils/log"
	"github.com/solo-io/solo-kit/test/helpers"
	"github.com/solo-io/solo-kit/test/setup"
	"k8s.io/client-go/rest"

	// Needed to run tests in GKE
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	// From https://github.com/kubernetes/client-go/blob/53c7adfd0294caa142d961e1f780f74081d5b15f/examples/out-of-cluster-client-configuration/main.go#L31
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/cache"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

var _ = Describe("V1Emitter", func() {
	if os.Getenv("RUN_KUBE_TESTS") != "1" {
		log.Printf("This test creates kubernetes resources and is disabled by default. To enable, set RUN_KUBE_TESTS=1 in your env.")
		return
	}
	var (
		namespace1          string
		namespace2          string
		name1, name2        = "angela" + helpers.RandString(3), "bob" + helpers.RandString(3)
		cfg                 *rest.Config
		emitter             ApiEmitter
		artifactClient      ArtifactClient
		endpointClient      EndpointClient
		proxyClient         ProxyClient
		upstreamGroupClient UpstreamGroupClient
		secretClient        SecretClient
		upstreamClient      UpstreamClient
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
		var kube kubernetes.Interface
		// Artifact Constructor
		kube, err = kubernetes.NewForConfig(cfg)
		Expect(err).NotTo(HaveOccurred())

		kcache, err := cache.NewKubeCoreCache(context.TODO(), kube)
		Expect(err).NotTo(HaveOccurred())
		artifactClientFactory := &factory.KubeConfigMapClientFactory{
			Clientset: kube,
			Cache:     kcache,
		}
		artifactClient, err = NewArtifactClient(artifactClientFactory)
		Expect(err).NotTo(HaveOccurred())
		// Endpoint Constructor
		kube, err = kubernetes.NewForConfig(cfg)
		Expect(err).NotTo(HaveOccurred())

		kcache, err := cache.NewKubeCoreCache(context.TODO(), kube)
		Expect(err).NotTo(HaveOccurred())
		endpointClientFactory := &factory.KubeConfigMapClientFactory{
			Clientset: kube,
			Cache:     kcache,
		}
		endpointClient, err = NewEndpointClient(endpointClientFactory)
		Expect(err).NotTo(HaveOccurred())
		// Proxy Constructor
		proxyClientFactory := &factory.KubeResourceClientFactory{
			Crd:         ProxyCrd,
			Cfg:         cfg,
			SharedCache: kuberc.NewKubeCache(context.TODO()),
		}
		proxyClient, err = NewProxyClient(proxyClientFactory)
		Expect(err).NotTo(HaveOccurred())
		// UpstreamGroup Constructor
		upstreamGroupClientFactory := &factory.KubeResourceClientFactory{
			Crd:         UpstreamGroupCrd,
			Cfg:         cfg,
			SharedCache: kuberc.NewKubeCache(context.TODO()),
		}
		upstreamGroupClient, err = NewUpstreamGroupClient(upstreamGroupClientFactory)
		Expect(err).NotTo(HaveOccurred())
		// Secret Constructor
		kube, err = kubernetes.NewForConfig(cfg)
		Expect(err).NotTo(HaveOccurred())

		kcache, err := cache.NewKubeCoreCache(context.TODO(), kube)
		Expect(err).NotTo(HaveOccurred())
		secretClientFactory := &factory.KubeSecretClientFactory{
			Clientset: kube,
			Cache:     kcache,
		}
		secretClient, err = NewSecretClient(secretClientFactory)
		Expect(err).NotTo(HaveOccurred())
		// Upstream Constructor
		upstreamClientFactory := &factory.KubeResourceClientFactory{
			Crd:         UpstreamCrd,
			Cfg:         cfg,
			SharedCache: kuberc.NewKubeCache(context.TODO()),
		}
		upstreamClient, err = NewUpstreamClient(upstreamClientFactory)
		Expect(err).NotTo(HaveOccurred())
		emitter = NewApiEmitter(artifactClient, endpointClient, proxyClient, upstreamGroupClient, secretClient, upstreamClient)
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

		var snap *ApiSnapshot

		/*
			Artifact
		*/

		assertSnapshotArtifacts := func(expectArtifacts ArtifactList, unexpectArtifacts ArtifactList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectArtifacts {
						if _, err := snap.Artifacts.List().Find(expected.Metadata.Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectArtifacts {
						if _, err := snap.Artifacts.List().Find(unexpected.Metadata.Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := artifactClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := artifactClient.List(namespace2, clients.ListOpts{})
					combined := ArtifactsByNamespace{
						namespace1: nsList1,
						namespace2: nsList2,
					}
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		artifact1a, err := artifactClient.Write(NewArtifact(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		artifact1b, err := artifactClient.Write(NewArtifact(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotArtifacts(ArtifactList{artifact1a, artifact1b}, nil)
		artifact2a, err := artifactClient.Write(NewArtifact(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		artifact2b, err := artifactClient.Write(NewArtifact(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotArtifacts(ArtifactList{artifact1a, artifact1b, artifact2a, artifact2b}, nil)

		err = artifactClient.Delete(artifact2a.Metadata.Namespace, artifact2a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = artifactClient.Delete(artifact2b.Metadata.Namespace, artifact2b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotArtifacts(ArtifactList{artifact1a, artifact1b}, ArtifactList{artifact2a, artifact2b})

		err = artifactClient.Delete(artifact1a.Metadata.Namespace, artifact1a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = artifactClient.Delete(artifact1b.Metadata.Namespace, artifact1b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotArtifacts(nil, ArtifactList{artifact1a, artifact1b, artifact2a, artifact2b})

		/*
			Endpoint
		*/

		assertSnapshotEndpoints := func(expectEndpoints EndpointList, unexpectEndpoints EndpointList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectEndpoints {
						if _, err := snap.Endpoints.List().Find(expected.Metadata.Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectEndpoints {
						if _, err := snap.Endpoints.List().Find(unexpected.Metadata.Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := endpointClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := endpointClient.List(namespace2, clients.ListOpts{})
					combined := EndpointsByNamespace{
						namespace1: nsList1,
						namespace2: nsList2,
					}
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		endpoint1a, err := endpointClient.Write(NewEndpoint(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		endpoint1b, err := endpointClient.Write(NewEndpoint(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotEndpoints(EndpointList{endpoint1a, endpoint1b}, nil)
		endpoint2a, err := endpointClient.Write(NewEndpoint(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		endpoint2b, err := endpointClient.Write(NewEndpoint(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotEndpoints(EndpointList{endpoint1a, endpoint1b, endpoint2a, endpoint2b}, nil)

		err = endpointClient.Delete(endpoint2a.Metadata.Namespace, endpoint2a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = endpointClient.Delete(endpoint2b.Metadata.Namespace, endpoint2b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotEndpoints(EndpointList{endpoint1a, endpoint1b}, EndpointList{endpoint2a, endpoint2b})

		err = endpointClient.Delete(endpoint1a.Metadata.Namespace, endpoint1a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = endpointClient.Delete(endpoint1b.Metadata.Namespace, endpoint1b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotEndpoints(nil, EndpointList{endpoint1a, endpoint1b, endpoint2a, endpoint2b})

		/*
			Proxy
		*/

		assertSnapshotProxies := func(expectProxies ProxyList, unexpectProxies ProxyList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectProxies {
						if _, err := snap.Proxies.List().Find(expected.Metadata.Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectProxies {
						if _, err := snap.Proxies.List().Find(unexpected.Metadata.Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := proxyClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := proxyClient.List(namespace2, clients.ListOpts{})
					combined := ProxiesByNamespace{
						namespace1: nsList1,
						namespace2: nsList2,
					}
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		proxy1a, err := proxyClient.Write(NewProxy(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		proxy1b, err := proxyClient.Write(NewProxy(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotProxies(ProxyList{proxy1a, proxy1b}, nil)
		proxy2a, err := proxyClient.Write(NewProxy(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		proxy2b, err := proxyClient.Write(NewProxy(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotProxies(ProxyList{proxy1a, proxy1b, proxy2a, proxy2b}, nil)

		err = proxyClient.Delete(proxy2a.Metadata.Namespace, proxy2a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = proxyClient.Delete(proxy2b.Metadata.Namespace, proxy2b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotProxies(ProxyList{proxy1a, proxy1b}, ProxyList{proxy2a, proxy2b})

		err = proxyClient.Delete(proxy1a.Metadata.Namespace, proxy1a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = proxyClient.Delete(proxy1b.Metadata.Namespace, proxy1b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotProxies(nil, ProxyList{proxy1a, proxy1b, proxy2a, proxy2b})

		/*
			UpstreamGroup
		*/

		assertSnapshotUpstreamgroups := func(expectUpstreamgroups UpstreamGroupList, unexpectUpstreamgroups UpstreamGroupList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectUpstreamgroups {
						if _, err := snap.Upstreamgroups.List().Find(expected.Metadata.Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectUpstreamgroups {
						if _, err := snap.Upstreamgroups.List().Find(unexpected.Metadata.Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := upstreamGroupClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := upstreamGroupClient.List(namespace2, clients.ListOpts{})
					combined := UpstreamgroupsByNamespace{
						namespace1: nsList1,
						namespace2: nsList2,
					}
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		upstreamGroup1a, err := upstreamGroupClient.Write(NewUpstreamGroup(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		upstreamGroup1b, err := upstreamGroupClient.Write(NewUpstreamGroup(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreamgroups(UpstreamGroupList{upstreamGroup1a, upstreamGroup1b}, nil)
		upstreamGroup2a, err := upstreamGroupClient.Write(NewUpstreamGroup(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		upstreamGroup2b, err := upstreamGroupClient.Write(NewUpstreamGroup(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreamgroups(UpstreamGroupList{upstreamGroup1a, upstreamGroup1b, upstreamGroup2a, upstreamGroup2b}, nil)

		err = upstreamGroupClient.Delete(upstreamGroup2a.Metadata.Namespace, upstreamGroup2a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = upstreamGroupClient.Delete(upstreamGroup2b.Metadata.Namespace, upstreamGroup2b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreamgroups(UpstreamGroupList{upstreamGroup1a, upstreamGroup1b}, UpstreamGroupList{upstreamGroup2a, upstreamGroup2b})

		err = upstreamGroupClient.Delete(upstreamGroup1a.Metadata.Namespace, upstreamGroup1a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = upstreamGroupClient.Delete(upstreamGroup1b.Metadata.Namespace, upstreamGroup1b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreamgroups(nil, UpstreamGroupList{upstreamGroup1a, upstreamGroup1b, upstreamGroup2a, upstreamGroup2b})

		/*
			Secret
		*/

		assertSnapshotSecrets := func(expectSecrets SecretList, unexpectSecrets SecretList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectSecrets {
						if _, err := snap.Secrets.List().Find(expected.Metadata.Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectSecrets {
						if _, err := snap.Secrets.List().Find(unexpected.Metadata.Ref().Strings()); err == nil {
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

		err = secretClient.Delete(secret2a.Metadata.Namespace, secret2a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = secretClient.Delete(secret2b.Metadata.Namespace, secret2b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotSecrets(SecretList{secret1a, secret1b}, SecretList{secret2a, secret2b})

		err = secretClient.Delete(secret1a.Metadata.Namespace, secret1a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = secretClient.Delete(secret1b.Metadata.Namespace, secret1b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotSecrets(nil, SecretList{secret1a, secret1b, secret2a, secret2b})

		/*
			Upstream
		*/

		assertSnapshotUpstreams := func(expectUpstreams UpstreamList, unexpectUpstreams UpstreamList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectUpstreams {
						if _, err := snap.Upstreams.List().Find(expected.Metadata.Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectUpstreams {
						if _, err := snap.Upstreams.List().Find(unexpected.Metadata.Ref().Strings()); err == nil {
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

		err = upstreamClient.Delete(upstream2a.Metadata.Namespace, upstream2a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = upstreamClient.Delete(upstream2b.Metadata.Namespace, upstream2b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreams(UpstreamList{upstream1a, upstream1b}, UpstreamList{upstream2a, upstream2b})

		err = upstreamClient.Delete(upstream1a.Metadata.Namespace, upstream1a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = upstreamClient.Delete(upstream1b.Metadata.Namespace, upstream1b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreams(nil, UpstreamList{upstream1a, upstream1b, upstream2a, upstream2b})
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

		var snap *ApiSnapshot

		/*
			Artifact
		*/

		assertSnapshotArtifacts := func(expectArtifacts ArtifactList, unexpectArtifacts ArtifactList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectArtifacts {
						if _, err := snap.Artifacts.List().Find(expected.Metadata.Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectArtifacts {
						if _, err := snap.Artifacts.List().Find(unexpected.Metadata.Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := artifactClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := artifactClient.List(namespace2, clients.ListOpts{})
					combined := ArtifactsByNamespace{
						namespace1: nsList1,
						namespace2: nsList2,
					}
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		artifact1a, err := artifactClient.Write(NewArtifact(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		artifact1b, err := artifactClient.Write(NewArtifact(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotArtifacts(ArtifactList{artifact1a, artifact1b}, nil)
		artifact2a, err := artifactClient.Write(NewArtifact(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		artifact2b, err := artifactClient.Write(NewArtifact(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotArtifacts(ArtifactList{artifact1a, artifact1b, artifact2a, artifact2b}, nil)

		err = artifactClient.Delete(artifact2a.Metadata.Namespace, artifact2a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = artifactClient.Delete(artifact2b.Metadata.Namespace, artifact2b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotArtifacts(ArtifactList{artifact1a, artifact1b}, ArtifactList{artifact2a, artifact2b})

		err = artifactClient.Delete(artifact1a.Metadata.Namespace, artifact1a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = artifactClient.Delete(artifact1b.Metadata.Namespace, artifact1b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotArtifacts(nil, ArtifactList{artifact1a, artifact1b, artifact2a, artifact2b})

		/*
			Endpoint
		*/

		assertSnapshotEndpoints := func(expectEndpoints EndpointList, unexpectEndpoints EndpointList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectEndpoints {
						if _, err := snap.Endpoints.List().Find(expected.Metadata.Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectEndpoints {
						if _, err := snap.Endpoints.List().Find(unexpected.Metadata.Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := endpointClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := endpointClient.List(namespace2, clients.ListOpts{})
					combined := EndpointsByNamespace{
						namespace1: nsList1,
						namespace2: nsList2,
					}
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		endpoint1a, err := endpointClient.Write(NewEndpoint(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		endpoint1b, err := endpointClient.Write(NewEndpoint(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotEndpoints(EndpointList{endpoint1a, endpoint1b}, nil)
		endpoint2a, err := endpointClient.Write(NewEndpoint(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		endpoint2b, err := endpointClient.Write(NewEndpoint(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotEndpoints(EndpointList{endpoint1a, endpoint1b, endpoint2a, endpoint2b}, nil)

		err = endpointClient.Delete(endpoint2a.Metadata.Namespace, endpoint2a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = endpointClient.Delete(endpoint2b.Metadata.Namespace, endpoint2b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotEndpoints(EndpointList{endpoint1a, endpoint1b}, EndpointList{endpoint2a, endpoint2b})

		err = endpointClient.Delete(endpoint1a.Metadata.Namespace, endpoint1a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = endpointClient.Delete(endpoint1b.Metadata.Namespace, endpoint1b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotEndpoints(nil, EndpointList{endpoint1a, endpoint1b, endpoint2a, endpoint2b})

		/*
			Proxy
		*/

		assertSnapshotProxies := func(expectProxies ProxyList, unexpectProxies ProxyList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectProxies {
						if _, err := snap.Proxies.List().Find(expected.Metadata.Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectProxies {
						if _, err := snap.Proxies.List().Find(unexpected.Metadata.Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := proxyClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := proxyClient.List(namespace2, clients.ListOpts{})
					combined := ProxiesByNamespace{
						namespace1: nsList1,
						namespace2: nsList2,
					}
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		proxy1a, err := proxyClient.Write(NewProxy(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		proxy1b, err := proxyClient.Write(NewProxy(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotProxies(ProxyList{proxy1a, proxy1b}, nil)
		proxy2a, err := proxyClient.Write(NewProxy(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		proxy2b, err := proxyClient.Write(NewProxy(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotProxies(ProxyList{proxy1a, proxy1b, proxy2a, proxy2b}, nil)

		err = proxyClient.Delete(proxy2a.Metadata.Namespace, proxy2a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = proxyClient.Delete(proxy2b.Metadata.Namespace, proxy2b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotProxies(ProxyList{proxy1a, proxy1b}, ProxyList{proxy2a, proxy2b})

		err = proxyClient.Delete(proxy1a.Metadata.Namespace, proxy1a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = proxyClient.Delete(proxy1b.Metadata.Namespace, proxy1b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotProxies(nil, ProxyList{proxy1a, proxy1b, proxy2a, proxy2b})

		/*
			UpstreamGroup
		*/

		assertSnapshotUpstreamgroups := func(expectUpstreamgroups UpstreamGroupList, unexpectUpstreamgroups UpstreamGroupList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectUpstreamgroups {
						if _, err := snap.Upstreamgroups.List().Find(expected.Metadata.Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectUpstreamgroups {
						if _, err := snap.Upstreamgroups.List().Find(unexpected.Metadata.Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := upstreamGroupClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := upstreamGroupClient.List(namespace2, clients.ListOpts{})
					combined := UpstreamgroupsByNamespace{
						namespace1: nsList1,
						namespace2: nsList2,
					}
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		upstreamGroup1a, err := upstreamGroupClient.Write(NewUpstreamGroup(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		upstreamGroup1b, err := upstreamGroupClient.Write(NewUpstreamGroup(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreamgroups(UpstreamGroupList{upstreamGroup1a, upstreamGroup1b}, nil)
		upstreamGroup2a, err := upstreamGroupClient.Write(NewUpstreamGroup(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		upstreamGroup2b, err := upstreamGroupClient.Write(NewUpstreamGroup(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreamgroups(UpstreamGroupList{upstreamGroup1a, upstreamGroup1b, upstreamGroup2a, upstreamGroup2b}, nil)

		err = upstreamGroupClient.Delete(upstreamGroup2a.Metadata.Namespace, upstreamGroup2a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = upstreamGroupClient.Delete(upstreamGroup2b.Metadata.Namespace, upstreamGroup2b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreamgroups(UpstreamGroupList{upstreamGroup1a, upstreamGroup1b}, UpstreamGroupList{upstreamGroup2a, upstreamGroup2b})

		err = upstreamGroupClient.Delete(upstreamGroup1a.Metadata.Namespace, upstreamGroup1a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = upstreamGroupClient.Delete(upstreamGroup1b.Metadata.Namespace, upstreamGroup1b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreamgroups(nil, UpstreamGroupList{upstreamGroup1a, upstreamGroup1b, upstreamGroup2a, upstreamGroup2b})

		/*
			Secret
		*/

		assertSnapshotSecrets := func(expectSecrets SecretList, unexpectSecrets SecretList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectSecrets {
						if _, err := snap.Secrets.List().Find(expected.Metadata.Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectSecrets {
						if _, err := snap.Secrets.List().Find(unexpected.Metadata.Ref().Strings()); err == nil {
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

		err = secretClient.Delete(secret2a.Metadata.Namespace, secret2a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = secretClient.Delete(secret2b.Metadata.Namespace, secret2b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotSecrets(SecretList{secret1a, secret1b}, SecretList{secret2a, secret2b})

		err = secretClient.Delete(secret1a.Metadata.Namespace, secret1a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = secretClient.Delete(secret1b.Metadata.Namespace, secret1b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotSecrets(nil, SecretList{secret1a, secret1b, secret2a, secret2b})

		/*
			Upstream
		*/

		assertSnapshotUpstreams := func(expectUpstreams UpstreamList, unexpectUpstreams UpstreamList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectUpstreams {
						if _, err := snap.Upstreams.List().Find(expected.Metadata.Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectUpstreams {
						if _, err := snap.Upstreams.List().Find(unexpected.Metadata.Ref().Strings()); err == nil {
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

		err = upstreamClient.Delete(upstream2a.Metadata.Namespace, upstream2a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = upstreamClient.Delete(upstream2b.Metadata.Namespace, upstream2b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreams(UpstreamList{upstream1a, upstream1b}, UpstreamList{upstream2a, upstream2b})

		err = upstreamClient.Delete(upstream1a.Metadata.Namespace, upstream1a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = upstreamClient.Delete(upstream1b.Metadata.Namespace, upstream1b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreams(nil, UpstreamList{upstream1a, upstream1b, upstream2a, upstream2b})
	})
})
