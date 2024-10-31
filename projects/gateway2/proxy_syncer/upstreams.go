package proxy_syncer

import (
	"context"
	"fmt"

	envoy_config_cluster_v3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	"github.com/solo-io/gloo/pkg/utils/settingsutil"
	"github.com/solo-io/gloo/projects/gateway2/krtcollections"
	ggv2utils "github.com/solo-io/gloo/projects/gateway2/utils"
	cluster "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/api/v2/cluster"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/gloosnapshot"
	glookubev1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/kube/apis/gloo.solo.io/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	"github.com/solo-io/gloo/projects/gloo/pkg/syncer/setup"
	"github.com/solo-io/go-utils/contextutils"
	envoycache "github.com/solo-io/solo-kit/pkg/api/v1/control-plane/cache"
	"github.com/solo-io/solo-kit/pkg/api/v1/control-plane/resource"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"istio.io/istio/pkg/kube/krt"
)

type uccWithCluster struct {
	Client         krtcollections.UniqlyConnectedClient
	Cluster        envoycache.Resource
	ClusterVersion uint64
	upstreamName   string
}

func (c uccWithCluster) ResourceName() string {
	return fmt.Sprintf("%s/%s", c.Client.ResourceName(), c.upstreamName)
}

func (c uccWithCluster) Equals(in uccWithCluster) bool {
	return c.Client.Equals(in.Client) && c.ClusterVersion == in.ClusterVersion
}

type IndexedUpstreams struct {
	clusters krt.Collection[uccWithCluster]
	index    krt.Index[string, uccWithCluster]
}

func (iu *IndexedUpstreams) FetchClustersForClient(kctx krt.HandlerContext, ucc krtcollections.UniqlyConnectedClient) []uccWithCluster {
	return krt.Fetch(kctx, iu.clusters, krt.FilterIndex(iu.index, ucc.ResourceName()))
}

func NewIndexedUpstreams(
	ctx context.Context,
	translator setup.TranslatorFactory,
	upstreams krt.Collection[UpstreamWrapper],
	uccs krt.Collection[krtcollections.UniqlyConnectedClient],
	ks krt.Collection[krtcollections.ResourceWrapper[*gloov1.Secret]],
	settings krt.Singleton[glookubev1.Settings],
	destinationRulesIndex DestinationRuleIndex) IndexedUpstreams {
	logger := contextutils.LoggerFrom(ctx).Desugar()

	clusters := krt.NewManyCollection(upstreams, func(kctx krt.HandlerContext, up UpstreamWrapper) []uccWithCluster {
		uccs := krt.Fetch(kctx, uccs)
		uccWithClusterRet := make([]uccWithCluster, 0, len(uccs))
		secrets := krt.Fetch(kctx, ks)
		ksettings := krt.FetchOne(kctx, settings.AsCollection())
		settings := &ksettings.Spec

		for _, ucc := range uccs {
			upstream, name := applyDestRulesForUpstream(logger, kctx, destinationRulesIndex, ucc.Namespace, up, ucc)

			latestSnap := &gloosnapshot.ApiSnapshot{}
			latestSnap.Secrets = make([]*gloov1.Secret, 0, len(secrets))
			for _, s := range secrets {
				latestSnap.Secrets = append(latestSnap.Secrets, s.Inner)
			}

			c, version := translate(ctx, settings, translator, latestSnap, upstream)
			if c == nil {
				continue
			}
			if name != "" && c.GetEdsClusterConfig() != nil {
				c.GetEdsClusterConfig().ServiceName = name
			}

			uccWithClusterRet = append(uccWithClusterRet, uccWithCluster{
				Client:         ucc,
				Cluster:        resource.NewEnvoyResource(c),
				ClusterVersion: version,
				upstreamName:   up.ResourceName(),
			})
		}
		return uccWithClusterRet
	})
	idx := krt.NewIndex(clusters, func(ucc uccWithCluster) []string {
		return []string{ucc.Client.ResourceName()}
	})

	return IndexedUpstreams{
		clusters: clusters,
		index:    idx,
	}
}

func translate(ctx context.Context, settings *gloov1.Settings, translator setup.TranslatorFactory, snap *gloosnapshot.ApiSnapshot, up *gloov1.Upstream) (*envoy_config_cluster_v3.Cluster, uint64) {

	ctx = settingsutil.WithSettings(ctx, settings)

	params := plugins.Params{
		Ctx:      ctx,
		Settings: settings,
		Snapshot: snap,
		Messages: map[*core.ResourceRef][]string{},
	}

	// false here should be ok - plugins should set eds on eds clusters.
	cluster, _ := translator.NewTranslator(ctx, settings).TranslateCluster(params, up, false)
	if cluster == nil {
		return nil, 0
	}

	return cluster, ggv2utils.HashProto(cluster)
}

func applyDestRulesForUpstream(logger *zap.Logger, kctx krt.HandlerContext, destinationRulesIndex DestinationRuleIndex, workloadNs string, u UpstreamWrapper, c krtcollections.UniqlyConnectedClient) (*gloov1.Upstream, string) {
	// host that would match the dest rule from the endpoints.
	// get the matching dest rule
	// get the lb info from the dest rules and call prioritize
	hostname := ggv2utils.GetHostnameForUpstream(u.Inner)

	destrule := destinationRulesIndex.FetchDestRulesFor(kctx, workloadNs, hostname, c.Labels)
	if destrule != nil {

		if outlier := destrule.Spec.GetTrafficPolicy().GetOutlierDetection(); outlier != nil {
			name := getEndpointClusterName(u.Inner)
			// do not mutate the original upstream
			up := *u.Inner

			out := &cluster.OutlierDetection{
				Consecutive_5Xx:  outlier.GetConsecutive_5XxErrors(),
				Interval:         outlier.GetInterval(),
				BaseEjectionTime: outlier.GetBaseEjectionTime(),
				// TODO: do the rest of them
			}
			if outlier.MaxEjectionPercent > 0 {
				out.MaxEjectionPercent = &wrapperspb.UInt32Value{Value: uint32(outlier.MaxEjectionPercent)}
			}

			up.OutlierDetection = out

			return &up, name
		}
	}

	return u.Inner, ""
}
