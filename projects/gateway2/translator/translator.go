package translator

import (
	"context"

	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/cache"

	"istio.io/istio/pkg/kube/krt"

	envoy_config_endpoint_v3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	"github.com/solo-io/gloo/pkg/utils/statsutils"
	"github.com/solo-io/gloo/projects/gateway2/endpoints"
	"github.com/solo-io/gloo/projects/gateway2/extensions2/common"
	extensionsplug "github.com/solo-io/gloo/projects/gateway2/extensions2/plugin"
	"github.com/solo-io/gloo/projects/gateway2/ir"
	"github.com/solo-io/gloo/projects/gateway2/krtcollections"
	"github.com/solo-io/gloo/projects/gateway2/query"
	"github.com/solo-io/gloo/projects/gateway2/reports"
	gwtranslator "github.com/solo-io/gloo/projects/gateway2/translator/gateway"
	"github.com/solo-io/gloo/projects/gateway2/translator/irtranslator"
	"github.com/solo-io/go-utils/contextutils"
	"sigs.k8s.io/controller-runtime/pkg/client"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
)

// Combines all the translators needed for xDS translation.
type CombinedTranslator struct {
	extensions extensionsplug.Plugin
	commonCols *common.CommonCollections

	waitForSync []cache.InformerSynced

	gwtranslator       extensionsplug.K8sGwTranslator
	irtranslator       *irtranslator.Translator
	upstreamTranslator *irtranslator.UpstreamTranslator
	endpointPlugins    []extensionsplug.EndpointPlugin

	logger *zap.Logger
}

func NewCombinedTranslator(
	ctx context.Context,
	extensions extensionsplug.Plugin,
	commonCols *common.CommonCollections,
) *CombinedTranslator {
	var endpointPlugins []extensionsplug.EndpointPlugin
	for _, ext := range extensions.ContributesPolicies {
		if ext.PerClientProcessEndpoints != nil {
			endpointPlugins = append(endpointPlugins, ext.PerClientProcessEndpoints)
		}
	}
	return &CombinedTranslator{
		commonCols:      commonCols,
		extensions:      extensions,
		endpointPlugins: endpointPlugins,
		logger:          contextutils.LoggerFrom(ctx).Desugar().With(zap.String("component", "translator_syncer")),
		waitForSync:     []cache.InformerSynced{extensions.HasSynced},
	}
}

// Note: isOurGw is shared between us and the deployer.
func (s *CombinedTranslator) Init(ctx context.Context, isOurGw func(gw *gwv1.Gateway) bool) error {
	ctx = contextutils.WithLogger(ctx, "k8s-gw-proxy-syncer")

	nsCol := krtcollections.NewNamespaceCollection(ctx, s.commonCols.Client, s.commonCols.KrtOpts)

	kubeGateways, routes, finalUpstreams, endpointIRs := krtcollections.InitCollections(ctx, s.extensions, s.commonCols.Client, isOurGw, s.commonCols.RefGrants, s.commonCols.KrtOpts)
	queries := query.NewData(
		routes,
		s.commonCols.Secrets,
		nsCol,
	)
	s.gwtranslator = gwtranslator.NewTranslator(queries)
	s.irtranslator = &irtranslator.Translator{
		ContributedPolicies: s.extensions.ContributesPolicies,
	}
	s.upstreamTranslator = &irtranslator.UpstreamTranslator{
		ContributedUpstreams: make(map[schema.GroupKind]ir.UpstreamInit),
		ContributedPolicies:  s.extensions.ContributesPolicies,
	}
	for k, up := range s.extensions.ContributesUpstreams {
		s.upstreamTranslator.ContributedUpstreams[k] = up.UpstreamInit
	}

	s.waitForSync = append(s.waitForSync,
		endpointIRs.Synced().HasSynced,
		endpointIRs.Synced().HasSynced,
		s.commonCols.HasSynced,
		finalUpstreams.Synced().HasSynced,
		kubeGateways.Gateways.Synced().HasSynced,
		s.extensions.HasSynced,
		routes.HasSynced,
	)
	return nil
}
func (s *CombinedTranslator) HasSynced() bool {
	for _, sync := range s.waitForSync {
		if !sync() {
			return false
		}
	}
	return true
}

// buildProxy performs translation of a kube Gateway -> gloov1.Proxy (really a wrapper type)
func (s *CombinedTranslator) buildProxy(kctx krt.HandlerContext, ctx context.Context, gw ir.Gateway, r reports.Reporter) *ir.GatewayIR {
	stopwatch := statsutils.NewTranslatorStopWatch("CombinedTranslator")
	stopwatch.Start()
	var gatewayTranslator extensionsplug.K8sGwTranslator = s.gwtranslator
	if s.extensions.ContributesGwTranslator != nil {
		maybeGatewayTranslator := s.extensions.ContributesGwTranslator(gw.Obj)
		if maybeGatewayTranslator != nil {
			// TODO: need better error handling here
			// and filtering out of our gateway classes, like before
			// contextutils.LoggerFrom(ctx).Errorf("no translator found for Gateway %s (gatewayClass %s)", gw.Name, gw.Obj.Spec.GatewayClassName)
			gatewayTranslator = maybeGatewayTranslator
		}
	} else {

	}
	proxy := gatewayTranslator.Translate(kctx, ctx, &gw, r)
	if proxy == nil {
		return nil
	}

	duration := stopwatch.Stop(ctx)
	contextutils.LoggerFrom(ctx).Debugf("translated proxy %s/%s in %s", gw.Namespace, gw.Name, duration.String())

	// TODO: these are likely unnecessary and should be removed!
	//	applyPostTranslationPlugins(ctx, pluginRegistry, &gwplugins.PostTranslationContext{
	//		TranslatedGateways: translatedGateways,
	//	})

	return proxy
}

func (s *CombinedTranslator) GetUpstreamTranslator() *irtranslator.UpstreamTranslator {
	return s.upstreamTranslator
}

// ctx needed for logging; remove once we refactor logging.
func (s *CombinedTranslator) TranslateGateway(kctx krt.HandlerContext, ctx context.Context, gw ir.Gateway) (*irtranslator.TranslationResult, reports.ReportMap) {
	logger := contextutils.LoggerFrom(ctx)

	logger.Debugf("building proxy for kube gw %s version %s", client.ObjectKeyFromObject(gw.Obj), gw.Obj.GetResourceVersion())
	rm := reports.NewReportMap()
	r := reports.NewReporter(&rm)
	gwir := s.buildProxy(kctx, ctx, gw, r)

	if gwir == nil {
		return nil, reports.ReportMap{}
	}

	// we are recomputing xds snapshots as proxies have changed, signal that we need to sync xds with these new snapshots
	xdsSnap := s.irtranslator.Translate(*gwir, r)

	return &xdsSnap, rm

}

func (s *CombinedTranslator) TranslateEndpoints(kctx krt.HandlerContext, ucc ir.UniqlyConnectedClient, ep ir.EndpointsForUpstream) (*envoy_config_endpoint_v3.ClusterLoadAssignment, uint64) {
	// check if we have a plugin to do it
	cla, additionalHash := proccessWithPlugins(s.endpointPlugins, kctx, context.TODO(), ucc, ep)
	if cla != nil {
		return cla, additionalHash

	}
	return endpoints.PrioritizeEndpoints(s.logger, nil, ep, ucc), 0
}

func proccessWithPlugins(plugins []extensionsplug.EndpointPlugin, kctx krt.HandlerContext, ctx context.Context, ucc ir.UniqlyConnectedClient, in ir.EndpointsForUpstream) (*envoy_config_endpoint_v3.ClusterLoadAssignment, uint64) {
	for _, processEnddpoints := range plugins {
		cla, additionalHash := processEnddpoints(kctx, context.TODO(), ucc, in)
		if cla != nil {
			return cla, additionalHash
		}
	}
	return nil, 0
}
