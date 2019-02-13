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
		emitter        SetupEmitter
		settingsClient SettingsClient
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
		// Settings Constructor
		settingsClientFactory := &factory.KubeResourceClientFactory{
			Crd:         SettingsCrd,
			Cfg:         cfg,
			SharedCache: kuberc.NewKubeCache(context.TODO()),
		}
		settingsClient, err = NewSettingsClient(settingsClientFactory)
		Expect(err).NotTo(HaveOccurred())
		emitter = NewSetupEmitter(settingsClient)
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

		var snap *SetupSnapshot

		/*
			Settings
		*/

		assertSnapshotSettings := func(expectSettings SettingsList, unexpectSettings SettingsList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectSettings {
						if _, err := snap.Settings.List().Find(expected.Metadata.Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectSettings {
						if _, err := snap.Settings.List().Find(unexpected.Metadata.Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := settingsClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := settingsClient.List(namespace2, clients.ListOpts{})
					combined := SettingsByNamespace{
						namespace1: nsList1,
						namespace2: nsList2,
					}
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		settings1a, err := settingsClient.Write(NewSettings(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		settings1b, err := settingsClient.Write(NewSettings(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotSettings(SettingsList{settings1a, settings1b}, nil)
		settings2a, err := settingsClient.Write(NewSettings(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		settings2b, err := settingsClient.Write(NewSettings(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotSettings(SettingsList{settings1a, settings1b, settings2a, settings2b}, nil)

		err = settingsClient.Delete(settings2a.Metadata.Namespace, settings2a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = settingsClient.Delete(settings2b.Metadata.Namespace, settings2b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotSettings(SettingsList{settings1a, settings1b}, SettingsList{settings2a, settings2b})

		err = settingsClient.Delete(settings1a.Metadata.Namespace, settings1a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = settingsClient.Delete(settings1b.Metadata.Namespace, settings1b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotSettings(nil, SettingsList{settings1a, settings1b, settings2a, settings2b})
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

		var snap *SetupSnapshot

		/*
			Settings
		*/

		assertSnapshotSettings := func(expectSettings SettingsList, unexpectSettings SettingsList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectSettings {
						if _, err := snap.Settings.List().Find(expected.Metadata.Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectSettings {
						if _, err := snap.Settings.List().Find(unexpected.Metadata.Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := settingsClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := settingsClient.List(namespace2, clients.ListOpts{})
					combined := SettingsByNamespace{
						namespace1: nsList1,
						namespace2: nsList2,
					}
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		settings1a, err := settingsClient.Write(NewSettings(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		settings1b, err := settingsClient.Write(NewSettings(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotSettings(SettingsList{settings1a, settings1b}, nil)
		settings2a, err := settingsClient.Write(NewSettings(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		settings2b, err := settingsClient.Write(NewSettings(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotSettings(SettingsList{settings1a, settings1b, settings2a, settings2b}, nil)

		err = settingsClient.Delete(settings2a.Metadata.Namespace, settings2a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = settingsClient.Delete(settings2b.Metadata.Namespace, settings2b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotSettings(SettingsList{settings1a, settings1b}, SettingsList{settings2a, settings2b})

		err = settingsClient.Delete(settings1a.Metadata.Namespace, settings1a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = settingsClient.Delete(settings1b.Metadata.Namespace, settings1b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotSettings(nil, SettingsList{settings1a, settings1b, settings2a, settings2b})
	})
})
