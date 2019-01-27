package pluginutils_test

import (
	"context"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	envoyroute "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	. "github.com/solo-io/gloo/projects/gloo/pkg/plugins/pluginutils"
)

var _ = Describe("PerFilterConfig", func() {
	var (
		in   *v1.Route
		out  *envoyroute.Route
		msg  *types.Struct
		name string
	)
	BeforeEach(func() {
		msg = &types.Struct{
			Fields: map[string]*types.Value{
				"test": &types.Value{Kind: &types.Value_BoolValue{
					BoolValue: true,
				}},
			},
		}
		name = "fakename"

	})
	Context("single dests", func() {

		BeforeEach(func() {
			in = &v1.Route{
				Action: &v1.Route_RouteAction{
					RouteAction: &v1.RouteAction{
						Destination: &v1.RouteAction_Single{
							Single: &v1.Destination{
								Upstream: core.ResourceRef{
									Name:      "test",
									Namespace: "",
								}},
						},
					},
				},
			}
			out = &envoyroute.Route{
				Action: &envoyroute.Route_Route{
					Route: &envoyroute.RouteAction{
						ClusterSpecifier: &envoyroute.RouteAction_Cluster{
							Cluster: "test",
						},
					},
				},
			}
		})

		It("should add per filter config to upstream", func() {

			err := MarkPerFilterConfig(context.TODO(), in, out, name, func(spec *v1.Destination) (proto.Message, error) {
				return msg, nil
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(out.PerFilterConfig).To(HaveKeyWithValue(name, BeEquivalentTo(msg)))
		})

		It("should add per filter config only to relevant upstream", func() {

			err := MarkPerFilterConfig(context.TODO(), in, out, name, func(spec *v1.Destination) (proto.Message, error) {
				return nil, nil
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(out.PerFilterConfig).ToNot(HaveKey(name))
		})
	})
	Context("multiple dests", func() {
		var (
			yescluster *envoyroute.WeightedCluster_ClusterWeight
			nocluster  *envoyroute.WeightedCluster_ClusterWeight
		)

		BeforeEach(func() {
			in = &v1.Route{
				Action: &v1.Route_RouteAction{
					RouteAction: &v1.RouteAction{
						Destination: &v1.RouteAction_Multi{
							Multi: &v1.MultiDestination{
								Destinations: []*v1.WeightedDestination{{
									Destination: &v1.Destination{
										Upstream: core.ResourceRef{
											Name:      "yes",
											Namespace: "",
										},
									},
								}, {
									Destination: &v1.Destination{
										Upstream: core.ResourceRef{
											Name:      "no",
											Namespace: "",
										},
									},
								}},
							},
						},
					},
				},
			}

			yescluster = &envoyroute.WeightedCluster_ClusterWeight{
				Name: "yes",
			}
			nocluster = &envoyroute.WeightedCluster_ClusterWeight{
				Name: "no",
			}
			out = &envoyroute.Route{
				Action: &envoyroute.Route_Route{
					Route: &envoyroute.RouteAction{
						ClusterSpecifier: &envoyroute.RouteAction_WeightedClusters{
							WeightedClusters: &envoyroute.WeightedCluster{
								Clusters: []*envoyroute.WeightedCluster_ClusterWeight{yescluster, nocluster},
							},
						},
					},
				},
			}
		})
		It("should add per filter config only to relevant upstream in mutiple dest", func() {

			err := MarkPerFilterConfig(context.TODO(), in, out, name, func(spec *v1.Destination) (proto.Message, error) {
				if spec.Upstream.Name == "yes" {
					return msg, nil
				}
				return nil, nil
			})
			Expect(err).NotTo(HaveOccurred())

			Expect(yescluster.PerFilterConfig).To(HaveKeyWithValue(name, BeEquivalentTo(msg)))
			Expect(nocluster.PerFilterConfig).ToNot(HaveKey(name))
			Expect(out.PerFilterConfig).ToNot(HaveKey(name))

		})
	})
})
