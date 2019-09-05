package aws

import (
	envoyapi "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoyroute "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	"github.com/envoyproxy/go-control-plane/pkg/util"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/aws"

	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/aws"
	awsapi "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/aws"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

const (
	accessKeyValue = "some acccess value"
	secretKeyValue = "some secret value"
)

var _ = Describe("Plugin", func() {
	var (
		params      plugins.Params
		vhostParams plugins.VirtualHostParams
		plugin      plugins.Plugin
		upstream    *v1.Upstream
		route       *v1.Route
		out         *envoyapi.Cluster
		outroute    *envoyroute.Route
	)
	BeforeEach(func() {
		var b bool
		plugin = NewPlugin(&b)
		plugin.Init(plugins.InitParams{})
		upstreamName := "up"
		clusterName := upstreamName
		funcName := "foo"
		upstream = &v1.Upstream{
			Metadata: core.Metadata{
				Name: upstreamName,
				// TODO(yuval-k): namespace
				Namespace: "",
			},
			UpstreamSpec: &v1.UpstreamSpec{
				UpstreamType: &v1.UpstreamSpec_Aws{
					Aws: &aws.UpstreamSpec{
						LambdaFunctions: []*aws.LambdaFunctionSpec{{
							LogicalName:        funcName,
							LambdaFunctionName: "foo",
							Qualifier:          "v1",
						}},
						Region: "us-east1",
						SecretRef: core.ResourceRef{
							Namespace: "",
							Name:      "secretref",
						},
					},
				},
			},
		}
		route = &v1.Route{
			Action: &v1.Route_RouteAction{
				RouteAction: &v1.RouteAction{
					Destination: &v1.RouteAction_Single{
						Single: &v1.Destination{
							DestinationType: &v1.Destination_Upstream{
								Upstream: &core.ResourceRef{
									Namespace: "",
									Name:      upstreamName,
								},
							},
							DestinationSpec: &v1.DestinationSpec{
								DestinationType: &v1.DestinationSpec_Aws{
									Aws: &awsapi.DestinationSpec{
										LogicalName: funcName,
									},
								},
							},
						},
					},
				},
			},
		}

		out = &envoyapi.Cluster{}
		outroute = &envoyroute.Route{
			Action: &envoyroute.Route_Route{
				Route: &envoyroute.RouteAction{
					ClusterSpecifier: &envoyroute.RouteAction_Cluster{
						Cluster: clusterName,
					},
				},
			},
		}

		params.Snapshot = &v1.ApiSnapshot{
			Secrets: v1.SecretList{{
				Metadata: core.Metadata{
					Name: "secretref",
					// TODO(yuval-k): namespace
					Namespace: "",
				},
				Kind: &v1.Secret_Aws{
					Aws: &v1.AwsSecret{
						AccessKey: accessKeyValue,
						SecretKey: secretKeyValue,
					},
				},
			}},
		}
		vhostParams = plugins.VirtualHostParams{Params: params}

	})
	Context("upstreams", func() {

		It("should process upstream with secrets", func() {
			err := plugin.(plugins.UpstreamPlugin).ProcessUpstream(params, upstream, out)
			Expect(err).NotTo(HaveOccurred())
			Expect(out.ExtensionProtocolOptions).NotTo(BeEmpty())
			Expect(out.ExtensionProtocolOptions).To(HaveKey(filterName))

			lpe := &LambdaProtocolExtension{}
			err = util.StructToMessage(out.ExtensionProtocolOptions[filterName], lpe)
			Expect(err).NotTo(HaveOccurred())

			Expect(lpe.AccessKey).To(Equal(accessKeyValue))
			Expect(lpe.SecretKey).To(Equal(secretKeyValue))
			Expect(lpe.Region).To(Equal("us-east1"))
			Expect(lpe.Host).To(Equal("lambda.us-east1.amazonaws.com"))
		})

		It("should error upstream with no secrets", func() {
			params.Snapshot.Secrets = nil
			err := plugin.(plugins.UpstreamPlugin).ProcessUpstream(params, upstream, out)
			Expect(err).To(HaveOccurred())
		})

		It("should error upstream with no access_key", func() {
			params.Snapshot.Secrets[0].Kind.(*v1.Secret_Aws).Aws.AccessKey = ""
			err := plugin.(plugins.UpstreamPlugin).ProcessUpstream(params, upstream, out)
			Expect(err).To(HaveOccurred())
		})
		It("should error upstream with no secret_key", func() {
			params.Snapshot.Secrets[0].Kind.(*v1.Secret_Aws).Aws.SecretKey = ""
			err := plugin.(plugins.UpstreamPlugin).ProcessUpstream(params, upstream, out)
			Expect(err).To(HaveOccurred())
		})
		It("should error upstream with invalid access_key", func() {
			params.Snapshot.Secrets[0].Kind.(*v1.Secret_Aws).Aws.AccessKey = "\xffbinary!"
			err := plugin.(plugins.UpstreamPlugin).ProcessUpstream(params, upstream, out)
			Expect(err).To(HaveOccurred())
		})
		It("should error upstream with invalid secret_key", func() {
			params.Snapshot.Secrets[0].Kind.(*v1.Secret_Aws).Aws.SecretKey = "\xffbinary!"
			err := plugin.(plugins.UpstreamPlugin).ProcessUpstream(params, upstream, out)
			Expect(err).To(HaveOccurred())
		})

		It("should not process and not error with non aws upstream", func() {
			upstream.UpstreamSpec.UpstreamType = nil
			err := plugin.(plugins.UpstreamPlugin).ProcessUpstream(params, upstream, out)
			Expect(err).NotTo(HaveOccurred())
			Expect(out.ExtensionProtocolOptions).To(BeEmpty())
			Expect(outroute.PerFilterConfig).NotTo(HaveKey(filterName))

		})
	})

	Context("routes", func() {

		var destination *v1.Destination

		BeforeEach(func() {
			err := plugin.(plugins.UpstreamPlugin).ProcessUpstream(params, upstream, out)
			Expect(err).NotTo(HaveOccurred())
			destination = route.Action.(*v1.Route_RouteAction).RouteAction.Destination.(*v1.RouteAction_Single).Single
		})

		It("should process route", func() {
			err := plugin.(plugins.RoutePlugin).ProcessRoute(plugins.RouteParams{VirtualHostParams: vhostParams}, route, outroute)
			Expect(err).NotTo(HaveOccurred())
			Expect(outroute.PerFilterConfig).To(HaveKey(filterName))
		})

		It("should not process with no spec", func() {
			destination.DestinationSpec = nil

			err := plugin.(plugins.RoutePlugin).ProcessRoute(plugins.RouteParams{VirtualHostParams: vhostParams}, route, outroute)
			Expect(err).NotTo(HaveOccurred())
			Expect(outroute.PerFilterConfig).NotTo(HaveKey(filterName))
		})

		It("should not process with a function mismatch", func() {
			destination.DestinationSpec.DestinationType.(*v1.DestinationSpec_Aws).Aws.LogicalName = "somethingelse"

			err := plugin.(plugins.RoutePlugin).ProcessRoute(plugins.RouteParams{VirtualHostParams: vhostParams}, route, outroute)
			Expect(err).To(HaveOccurred())
			Expect(outroute.PerFilterConfig).NotTo(HaveKey(filterName))
		})

		It("should not process with no spec", func() {
			Skip("redo this when we have more destination type")
			// destination.DestinationSpec.DestinationType =

			err := plugin.(plugins.RoutePlugin).ProcessRoute(plugins.RouteParams{VirtualHostParams: vhostParams}, route, outroute)
			Expect(err).NotTo(HaveOccurred())
			Expect(outroute.PerFilterConfig).NotTo(HaveKey(filterName))
		})

	})

	Context("filters", func() {
		It("should produce filters when upstream is present", func() {
			// process upstream
			err := plugin.(plugins.UpstreamPlugin).ProcessUpstream(params, upstream, out)
			Expect(err).NotTo(HaveOccurred())
			err = plugin.(plugins.RoutePlugin).ProcessRoute(plugins.RouteParams{VirtualHostParams: vhostParams}, route, outroute)
			Expect(err).NotTo(HaveOccurred())

			// check that we have filters
			filters, err := plugin.(plugins.HttpFilterPlugin).HttpFilters(params, nil)
			Expect(err).NotTo(HaveOccurred())
			Expect(filters).NotTo(BeEmpty())
		})

		It("should not produce filters when no upstreams are present", func() {
			filters, err := plugin.(plugins.HttpFilterPlugin).HttpFilters(params, nil)
			Expect(err).NotTo(HaveOccurred())
			Expect(filters).To(BeEmpty())
		})
	})
})
