package ratelimit_test

import (
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ratelimit"
	"github.com/solo-io/gloo/projects/gloo/pkg/syncer"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"
	envoycache "github.com/solo-io/solo-kit/pkg/api/v1/control-plane/cache"
	skcore "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/api/v2/reporter"

	. "github.com/solo-io/gloo/projects/gloo/pkg/syncer/ratelimit"
)

var _ = Describe("RatelimitTranslatorSyncer", func() {
	var (
		ctx              context.Context
		cancel           context.CancelFunc
		proxy            *gloov1.Proxy
		params           syncer.TranslatorSyncerExtensionParams
		translator       syncer.TranslatorSyncerExtension
		apiSnapshot      *gloov1.ApiSnapshot
		proxyClient      clients.ResourceClient
		snapCache        *mockSetSnapshot
	)
	JustBeforeEach(func() {
		ctx, cancel = context.WithCancel(context.Background())
		var err error
		helpers.UseMemoryClients()
		resourceClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}

		proxyClient, err = resourceClientFactory.NewResourceClient(ctx, factory.NewResourceClientParams{ResourceType: &gloov1.Proxy{}})
		Expect(err).NotTo(HaveOccurred())

		params.Reports = make(reporter.ResourceReports)
		translator, err = NewTranslatorSyncerExtension(ctx, params)
		Expect(err).NotTo(HaveOccurred())

		config := &ratelimit.IngressRateLimit{
			AuthorizedLimits: nil,
			AnonymousLimits:  nil,
		}

		proxy = &gloov1.Proxy{
			Metadata: &skcore.Metadata{
				Name:      "proxy",
				Namespace: "gloo-system",
			},
			Listeners: []*gloov1.Listener{{
				Name: "listener-::-8080",
				ListenerType: &gloov1.Listener_HttpListener{
					HttpListener: &gloov1.HttpListener{
						VirtualHosts: []*gloov1.VirtualHost{
							&gloov1.VirtualHost{
								Name: "gloo-system.default",
								Options: &gloov1.VirtualHostOptions{
									RatelimitBasic: config,
								},
							},
						},
					},
				},
			}},
		}

		proxyClient.Write(proxy, clients.WriteOpts{})

		apiSnapshot = &gloov1.ApiSnapshot{
			Proxies:     []*gloov1.Proxy{proxy},
		}
	})

	AfterEach(func() {
		cancel()
	})

	Context("strongly typed config", func() {

		Context("config that needs to be translated (non-custom)", func() {
			It("should work with one listener", func() {
				_, err := translator.Sync(ctx, apiSnapshot, snapCache)
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(ErrEnterpriseOnly))
			})

		})
	})
})

type mockSetSnapshot struct {
	Snapshots map[string]envoycache.Snapshot
}

func (m *mockSetSnapshot) CreateWatch(request envoycache.Request) (value chan envoycache.Response, cancel func()) {
	// Dummy method
	return nil, nil
}

func (m *mockSetSnapshot) Fetch(ctx context.Context, request envoycache.Request) (*envoycache.Response, error) {
	// Dummy method
	return nil, nil
}

func (m *mockSetSnapshot) GetStatusInfo(s string) envoycache.StatusInfo {
	// Dummy method
	return nil
}

func (m *mockSetSnapshot) GetStatusKeys() []string {
	// Dummy method
	return []string{}
}

func (m *mockSetSnapshot) GetSnapshot(node string) (envoycache.Snapshot, error) {
	// Dummy method
	return m.Snapshots[node], nil
}

func (m *mockSetSnapshot) ClearSnapshot(node string) {
	// Dummy method
	m.Snapshots[node] = nil
}

func (m *mockSetSnapshot) SetSnapshot(node string, snapshot envoycache.Snapshot) error {
	if m.Snapshots == nil {
		m.Snapshots = make(map[string]envoycache.Snapshot)
	}

	m.Snapshots[node] = snapshot
	return nil
}



