package azure_test

import (
	"context"

	envoy_config_cluster_v3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	"github.com/solo-io/gloo/projects/gloo/pkg/utils"

	envoyauth "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/azure"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	azureplugin "github.com/solo-io/gloo/projects/gloo/pkg/plugins/azure"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Plugin", func() {
	var (
		p      plugins.Plugin
		out    *envoy_config_cluster_v3.Cluster
		params plugins.Params
	)

	BeforeEach(func() {
		var b bool
		p = azureplugin.NewPlugin(&b)
		p.Init(plugins.InitParams{Ctx: context.TODO()})
		out = &envoy_config_cluster_v3.Cluster{}
		params = plugins.Params{}
	})

	Context("with valid upstream spec", func() {
		var (
			err      error
			upstream *v1.Upstream
		)

		BeforeEach(func() {
			upstream = &v1.Upstream{
				Metadata: &core.Metadata{
					Name: "test",
					// TODO(yuval-k): namespace
					Namespace: "",
				},
				UpstreamType: &v1.Upstream_Azure{
					Azure: &azure.UpstreamSpec{
						FunctionAppName: "my-appwhos",
					},
				},
			}
		})
		Context("with secrets", func() {

			BeforeEach(func() {
				upstream.UpstreamType.(*v1.Upstream_Azure).Azure.SecretRef = &core.ResourceRef{
					Namespace: "",
					Name:      "azure-secret1",
				}

				params.Snapshot = &v1.ApiSnapshot{
					Secrets: v1.SecretList{{
						Metadata: &core.Metadata{
							Name: "azure-secret1",
							// TODO(yuval-k): namespace
							Namespace: "",
						},
						Kind: &v1.Secret_Azure{
							Azure: &v1.AzureSecret{
								ApiKeys: map[string]string{"_master": "key1", "foo": "key1", "bar": "key2"},
							},
						},
					}},
				}

				err = p.(plugins.UpstreamPlugin).ProcessUpstream(params, upstream, out)
			})
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("should have the correct output", func() {
				Expect(out.LoadAssignment.Endpoints).Should(HaveLen(1))

				tlsContext := utils.MustAnyToMessage(out.TransportSocket.GetTypedConfig()).(*envoyauth.UpstreamTlsContext)
				Expect(tlsContext.Sni).To(Equal("my-appwhos.azurewebsites.net"))
				Expect(out.GetType()).To(Equal(envoy_config_cluster_v3.Cluster_LOGICAL_DNS))
				Expect(out.DnsLookupFamily).To(Equal(envoy_config_cluster_v3.Cluster_V4_ONLY))
			})

		})
		Context("without secrets", func() {
			BeforeEach(func() {
				err = p.(plugins.UpstreamPlugin).ProcessUpstream(params, upstream, out)
			})
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
			It("should have the correct output", func() {
				Expect(out.LoadAssignment.Endpoints).Should(HaveLen(1))
				tlsContext := utils.MustAnyToMessage(out.TransportSocket.GetTypedConfig()).(*envoyauth.UpstreamTlsContext)
				Expect(tlsContext.Sni).To(Equal("my-appwhos.azurewebsites.net"))
				Expect(out.GetType()).To(Equal(envoy_config_cluster_v3.Cluster_LOGICAL_DNS))
				Expect(out.DnsLookupFamily).To(Equal(envoy_config_cluster_v3.Cluster_V4_ONLY))
			})
		})
	})
})
