package syncer

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	runner2 "github.com/solo-io/gloo/projects/discovery/pkg/uds/runner"
	"github.com/solo-io/gloo/projects/gloo/pkg/runner"

	"github.com/golang/protobuf/ptypes/wrappers"

	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

var _ = Describe("UDS setup syncer tests", func() {
	Context("StartUDS", func() {
		It("returns an error when both UDS and FDS are disabled", func() {
			opts := runner.StartOpts{
				Settings: &v1.Settings{
					Metadata: &core.Metadata{
						Name:      "test-settings",
						Namespace: "gloo-system",
					},
					Discovery: &v1.Settings_DiscoveryOptions{
						UdsOptions: &v1.Settings_DiscoveryOptions_UdsOptions{
							Enabled: &wrappers.BoolValue{Value: false},
						},
						FdsMode: v1.Settings_DiscoveryOptions_DISABLED,
					},
				},
			}
			Expect(runner2.StartUDS(opts)).To(HaveOccurred())
		})

		It("Does not return an error when WatchLabels are set", func() {
			memcache := memory.NewInMemoryResourceCache()
			f := &factory.MemoryResourceClientFactory{
				Cache: memcache,
			}

			opts := runner.StartOpts{
				Settings: &v1.Settings{
					Metadata: &core.Metadata{
						Name:      "test-settings",
						Namespace: "gloo-system",
					},
					Discovery: &v1.Settings_DiscoveryOptions{
						UdsOptions: &v1.Settings_DiscoveryOptions_UdsOptions{
							Enabled:     &wrappers.BoolValue{Value: true},
							WatchLabels: map[string]string{"A": "B"},
						},
						FdsMode: v1.Settings_DiscoveryOptions_DISABLED,
					},
				},
				Upstreams: f,
				Secrets:   f,
			}
			Expect(runner2.StartUDS(opts)).NotTo(HaveOccurred())
		})
	})

})
