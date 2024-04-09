package xds

import (
	"context"

	"github.com/solo-io/solo-kit/pkg/utils/statusutils"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"

	"github.com/solo-io/gloo/pkg/utils/syncutil"

	"github.com/solo-io/gloo/projects/gateway2/query"

	"github.com/solo-io/gloo/projects/gateway2/extensions"

	gwplugins "github.com/solo-io/gloo/projects/gateway2/translator/plugins"
	"github.com/solo-io/gloo/projects/gateway2/translator/translatorutils"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	"github.com/solo-io/gloo/projects/gateway2/reports"
	gloot "github.com/solo-io/gloo/projects/gateway2/translator"
	"github.com/solo-io/gloo/projects/gateway2/translator/plugins/registry"
	gloo_solo_io "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	v1snap "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/gloosnapshot"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	"github.com/solo-io/gloo/projects/gloo/pkg/syncer/sanitizer"
	syncerstats "github.com/solo-io/gloo/projects/gloo/pkg/syncer/stats"
	"github.com/solo-io/gloo/projects/gloo/pkg/translator"
	"github.com/solo-io/gloo/projects/gloo/pkg/utils"
	"github.com/solo-io/gloo/projects/gloo/pkg/xds"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/go-utils/hashutils"
	envoycache "github.com/solo-io/solo-kit/pkg/api/v1/control-plane/cache"
	"github.com/solo-io/solo-kit/pkg/api/v1/control-plane/types"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/api/v2/reporter"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
	"go.opencensus.io/trace"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	apiv1 "sigs.k8s.io/gateway-api/apis/v1"
)

// empty resources to give to envoy when a proxy was deleted
const emptyVersionKey = "empty"

var (
	emptyResource = envoycache.Resources{
		Version: emptyVersionKey,
		Items:   map[string]envoycache.Resource{},
	}
	emptySnapshot = xds.NewSnapshotFromResources(
		emptyResource,
		emptyResource,
		emptyResource,
		emptyResource,
	)
)

var (
	envoySnapshotOut   = stats.Int64("api.gloo.solo.io/translator/resources", "The number of resources in the snapshot in", "1")
	resourceNameKey, _ = tag.NewKey("resource")

	envoySnapshotOutView = &view.View{
		Name:        "api.gloo.solo.io/translator/resources",
		Measure:     envoySnapshotOut,
		Description: "The number of resources in the snapshot for envoy",
		Aggregation: view.LastValue(),
		TagKeys:     []tag.Key{syncerstats.ProxyNameKey, resourceNameKey},
	}
)

func init() {
	_ = view.Register(envoySnapshotOutView)
}

type XdsSyncer struct {
	translator     translator.Translator
	sanitizer      sanitizer.XdsSanitizer
	xdsCache       envoycache.SnapshotCache
	controllerName string

	// used for debugging purposes only
	latestSnap *v1snap.ApiSnapshot

	xdsGarbageCollection bool

	inputs          *XdsInputChannels
	mgr             manager.Manager
	k8sGwExtensions extensions.K8sGatewayExtensions

	// proxyReconciler wraps the client that writes Proxy resources into an in-memory cache
	// This cache is utilized by the debug.ProxyEndpointServer
	proxyReconciler gloo_solo_io.ProxyReconciler
}

type XdsInputChannels struct {
	genericEvent   AsyncQueue[struct{}]
	discoveryEvent AsyncQueue[DiscoveryInputs]
	secretEvent    AsyncQueue[SecretInputs]
}

func (x *XdsInputChannels) Kick(ctx context.Context) {
	x.genericEvent.Enqueue(struct{}{})
}

func (x *XdsInputChannels) UpdateDiscoveryInputs(ctx context.Context, inputs DiscoveryInputs) {
	x.discoveryEvent.Enqueue(inputs)
}

func (x *XdsInputChannels) UpdateSecretInputs(ctx context.Context, inputs SecretInputs) {
	x.secretEvent.Enqueue(inputs)
}

func NewXdsInputChannels() *XdsInputChannels {
	return &XdsInputChannels{
		genericEvent:   NewAsyncQueue[struct{}](),
		discoveryEvent: NewAsyncQueue[DiscoveryInputs](),
		secretEvent:    NewAsyncQueue[SecretInputs](),
	}
}

func NewXdsSyncer(
	controllerName string,
	translator translator.Translator,
	sanitizer sanitizer.XdsSanitizer,
	xdsCache envoycache.SnapshotCache,
	xdsGarbageCollection bool,
	inputs *XdsInputChannels,
	mgr manager.Manager,
	k8sGwExtensions extensions.K8sGatewayExtensions,
	proxyClient gloo_solo_io.ProxyClient,
) *XdsSyncer {
	return &XdsSyncer{
		controllerName:       controllerName,
		translator:           translator,
		sanitizer:            sanitizer,
		xdsCache:             xdsCache,
		xdsGarbageCollection: xdsGarbageCollection,
		inputs:               inputs,
		mgr:                  mgr,
		k8sGwExtensions:      k8sGwExtensions,
		proxyReconciler:      gloo_solo_io.NewProxyReconciler(proxyClient, statusutils.NewNoOpStatusClient()),
	}
}

func (s *XdsSyncer) Start(ctx context.Context) error {
	proxyApiSnapshot := &v1snap.ApiSnapshot{}
	ctx = contextutils.WithLogger(ctx, "k8s-gw-syncer")

	var (
		discoveryWarmed bool
		secretsWarmed   bool
	)
	resyncXds := func() {
		if !discoveryWarmed || !secretsWarmed {
			return
		}

		var gwl apiv1.GatewayList
		err := s.mgr.GetClient().List(ctx, &gwl)
		if err != nil {
			// This should never happen, try again?
			return
		}

		gatewayQueries := query.NewData(s.mgr.GetClient(), s.mgr.GetScheme())

		pluginRegistry := s.k8sGwExtensions.CreatePluginRegistry(ctx)
		gatewayTranslator := gloot.NewTranslator(gatewayQueries, pluginRegistry)

		proxies := gloo_solo_io.ProxyList{}
		rm := reports.NewReportMap()
		r := reports.NewReporter(&rm)

		var translatedGateways []gwplugins.TranslatedGateway
		for _, gw := range gwl.Items {
			proxy := gatewayTranslator.TranslateProxy(ctx, &gw, r)
			if proxy != nil {
				proxies = append(proxies, proxy)
				translatedGateways = append(translatedGateways, gwplugins.TranslatedGateway{
					Gateway: gw,
				})
				//TODO: handle reports and process statuses
			}
		}
		proxyApiSnapshot.Proxies = proxies

		applyPostTranslationPlugins(ctx, pluginRegistry, &gwplugins.PostTranslationContext{
			TranslatedGateways: translatedGateways,
		})

		proxiesWithReports := s.syncEnvoy(ctx, proxyApiSnapshot)
		s.applyStatusPlugins(ctx, pluginRegistry, proxiesWithReports)
		s.syncStatus(ctx, rm, gwl)
		s.syncRouteStatus(ctx, rm)
		s.syncProxyCache(ctx, proxies)
	}

	for {
		select {
		case <-ctx.Done():
			contextutils.LoggerFrom(ctx).Debug("context done, stopping syncer")
			return nil
		case <-s.inputs.genericEvent.Next():
			resyncXds()
		case discoveryEvent := <-s.inputs.discoveryEvent.Next():
			proxyApiSnapshot.Upstreams = discoveryEvent.Upstreams
			proxyApiSnapshot.Endpoints = discoveryEvent.Endpoints
			discoveryWarmed = true
			resyncXds()
		case secretEvent := <-s.inputs.secretEvent.Next():
			proxyApiSnapshot.Secrets = secretEvent.Secrets
			secretsWarmed = true
			resyncXds()
		}
	}
}

func (s *XdsSyncer) applyStatusPlugins(
	ctx context.Context,
	pluginRegistry registry.PluginRegistry,
	proxiesWithReports []translatorutils.ProxyWithReports,
) {
	statusCtx := &gwplugins.StatusContext{
		ProxiesWithReports: proxiesWithReports,
	}
	for _, plugin := range pluginRegistry.GetStatusPlugins() {
		plugin.ApplyStatusPlugin(ctx, statusCtx)
	}
}

// syncEnvoy will translate, sanatize, and set the snapshot for each of the proxies, all while merging all the reports into allReports.
// NOTE(ilackarms): the below code was copy-pasted (with some deletions) from projects/gloo/pkg/syncer/translator_syncer.go
func (s *XdsSyncer) syncEnvoy(ctx context.Context, snap *v1snap.ApiSnapshot) []translatorutils.ProxyWithReports {
	ctx, span := trace.StartSpan(ctx, "gloo.syncer.Sync")
	defer span.End()

	s.latestSnap = snap
	logger := log.FromContext(ctx, "pkg", "envoyTranslatorSyncer")
	snapHash := hashutils.MustHash(snap)
	logger.Info("begin sync", "snapHash", snapHash,
		"len(proxies)", len(snap.Proxies), "len(upstreams)", len(snap.Upstreams), "len(endpoints)", len(snap.Endpoints), "len(secrets)", len(snap.Secrets), "len(artifacts)", len(snap.Artifacts), "len(authconfigs)", len(snap.AuthConfigs), "len(ratelimits)", len(snap.Ratelimitconfigs), "len(graphqls)", len(snap.GraphqlApis))
	debugLogger := logger.V(1)

	defer logger.Info("end sync", "len(snapHash)", snapHash)

	// stringifying the snapshot may be an expensive operation, so we'd like to avoid building the large
	// string if we're not even going to log it anyway
	if debugLogger.Enabled() {
		debugLogger.Info("snap", "snap", syncutil.StringifySnapshot(snap))
	}

	reportss := make(reporter.ResourceReports)
	reportss.Accept(snap.Upstreams.AsInputResources()...)
	reportss.Accept(snap.Proxies.AsInputResources()...)

	if !s.xdsGarbageCollection {
		allKeys := map[string]bool{
			xds.FallbackNodeCacheKey: true,
		}
		// Get all envoy node ID keys
		for _, key := range s.xdsCache.GetStatusKeys() {
			allKeys[key] = false
		}
		// Get all valid node ID keys for Proxies
		for _, key := range xds.SnapshotCacheKeys(utils.GlooGatewayTranslatorValue, snap.Proxies) {
			allKeys[key] = true
		}

		// preserve keys from the current list of proxies, set previous invalid snapshots to empty snapshot
		for key, valid := range allKeys {
			if !valid && xds.SnapshotBelongsTo(key, utils.GlooGatewayTranslatorValue) {
				s.xdsCache.SetSnapshot(key, emptySnapshot)
			}
		}
	}
	proxiesWithReports := []translatorutils.ProxyWithReports{}
	for _, proxy := range snap.Proxies {
		proxyCtx := ctx
		if ctxWithTags, err := tag.New(proxyCtx, tag.Insert(syncerstats.ProxyNameKey, proxy.GetMetadata().Ref().Key())); err == nil {
			proxyCtx = ctxWithTags
		}

		params := plugins.Params{
			Ctx:      proxyCtx,
			Snapshot: snap,
			Messages: map[*core.ResourceRef][]string{},
		}

		xdsSnapshot, reports, proxyReport := s.translator.Translate(params, proxy)
		proxyWithReport := translatorutils.ProxyWithReports{
			Proxy: proxy,
			Reports: translatorutils.TranslationReports{
				ProxyReport:     proxyReport,
				ResourceReports: reports,
			},
		}
		proxiesWithReports = append(proxiesWithReports, proxyWithReport)

		// if validateErr := reports.ValidateStrict(); validateErr != nil {
		// 	logger.Warnw("Proxy had invalid config", zap.Any("proxy", proxy.GetMetadata().Ref()), zap.Error(validateErr))
		// }

		sanitizedSnapshot := s.sanitizer.SanitizeSnapshot(ctx, snap, xdsSnapshot, reportss)
		// if the snapshot is not consistent, make it so
		xdsSnapshot.MakeConsistent()

		// if validateErr := reports.ValidateStrict(); validateErr != nil {
		// 	logger.Error(validateErr, "Proxy had invalid config after xds sanitization", "proxy", proxy.GetMetadata().Ref())
		// }

		debugLogger.Info("snap", "key", sanitizedSnapshot)

		// Merge reports after sanitization to capture changes made by the sanitizers
		reportss.Merge(reportss)
		key := xds.SnapshotCacheKey(utils.GlooGatewayTranslatorValue, proxy)
		s.xdsCache.SetSnapshot(key, sanitizedSnapshot)

		// Record some metrics
		clustersLen := len(xdsSnapshot.GetResources(types.ClusterTypeV3).Items)
		listenersLen := len(xdsSnapshot.GetResources(types.ListenerTypeV3).Items)
		routesLen := len(xdsSnapshot.GetResources(types.RouteTypeV3).Items)
		endpointsLen := len(xdsSnapshot.GetResources(types.EndpointTypeV3).Items)

		measureResource(proxyCtx, "clusters", clustersLen)
		measureResource(proxyCtx, "listeners", listenersLen)
		measureResource(proxyCtx, "routes", routesLen)
		measureResource(proxyCtx, "endpoints", endpointsLen)

		debugLogger.Info("Setting xDS Snapshot", "key", key,
			"clusters", clustersLen,
			"clustersVersion", xdsSnapshot.GetResources(types.ClusterTypeV3).Version,
			"listeners", listenersLen,
			"listenersVersion", xdsSnapshot.GetResources(types.ListenerTypeV3).Version,
			"routes", routesLen,
			"routesVersion", xdsSnapshot.GetResources(types.RouteTypeV3).Version,
			"endpoints", endpointsLen,
			"endpointsVersion", xdsSnapshot.GetResources(types.EndpointTypeV3).Version)

		debugLogger.Info("Full snapshot for proxy", proxy.GetMetadata().GetName(), xdsSnapshot)
	}

	debugLogger.Info("gloo reports to be written", "reports", reportss)

	return proxiesWithReports
}

func measureResource(ctx context.Context, resource string, length int) {
	if ctxWithTags, err := tag.New(ctx, tag.Insert(resourceNameKey, resource)); err == nil {
		stats.Record(ctxWithTags, envoySnapshotOut.M(int64(length)))
	}
}

func (s *XdsSyncer) syncRouteStatus(ctx context.Context, rm reports.ReportMap) {
	ctx = contextutils.WithLogger(ctx, "routeStatusSyncer")
	logger := contextutils.LoggerFrom(ctx)
	rl := apiv1.HTTPRouteList{}
	err := s.mgr.GetClient().List(ctx, &rl)
	if err != nil {
		logger.Error(err)
		return
	}

	for _, route := range rl.Items {
		route := route // pike
		if status := rm.BuildRouteStatus(ctx, route, s.controllerName); status != nil {
			route.Status = *status
			if err := s.mgr.GetClient().Status().Update(ctx, &route); err != nil {
				logger.Error(err)
			}
		}
	}
}

func (s *XdsSyncer) syncStatus(ctx context.Context, rm reports.ReportMap, gwl apiv1.GatewayList) {
	ctx = contextutils.WithLogger(ctx, "statusSyncer")
	logger := contextutils.LoggerFrom(ctx)
	for _, gw := range gwl.Items {
		gw := gw // pike
		if status := rm.BuildGWStatus(ctx, gw); status != nil {
			gw.Status = *status
			if err := s.mgr.GetClient().Status().Patch(ctx, &gw, client.Merge); err != nil {
				logger.Error(err)
			}
		}
	}
}

// syncProxyCache persists the proxies that were generated during translations and stores them in an in-memory cache
// This cache is utilized by the debug.ProxyEndpointServer
func (s *XdsSyncer) syncProxyCache(ctx context.Context, proxyList gloo_solo_io.ProxyList) {
	ctx = contextutils.WithLogger(ctx, "proxyCache")
	logger := contextutils.LoggerFrom(ctx)

	// Proxy CR is located in the same namespace as the originating Gateway CR
	// As a result, we may have a list of Proxies that are in different namespaces
	// Since the reconciler accepts the namespace as an argument, we need to split
	// the list so we have a lists of proxies, isolated to each namespace
	var proxyListByNamespace = make(map[string]gloo_solo_io.ProxyList)

	for _, proxy := range proxyList {
		proxyNs := proxy.GetMetadata().GetNamespace()
		nsList, ok := proxyListByNamespace[proxyNs]
		if ok {
			nsList = append(nsList, proxy)
			proxyListByNamespace[proxyNs] = nsList
		} else {
			proxyListByNamespace[proxyNs] = gloo_solo_io.ProxyList{
				proxy,
			}
		}
	}

	for ns, nsList := range proxyListByNamespace {
		err := s.proxyReconciler.Reconcile(
			ns,
			nsList,
			func(original, desired *gloo_solo_io.Proxy) (bool, error) {
				// Do nothing to the new proxy, just always overwrite the old one
				return true, nil
			},
			clients.ListOpts{
				Ctx: ctx,
			})
		if err != nil {
			// A write error to our cache should not impact translation
			// We will emit a message, and continue
			logger.Error(err)
		}
	}
}

func applyPostTranslationPlugins(ctx context.Context, pluginRegistry registry.PluginRegistry, translationContext *gwplugins.PostTranslationContext) {
	ctx = contextutils.WithLogger(ctx, "postTranslation")
	logger := contextutils.LoggerFrom(ctx)

	for _, postTranslationPlugin := range pluginRegistry.GetPostTranslationPlugins() {
		err := postTranslationPlugin.ApplyPostTranslationPlugin(ctx, translationContext)
		if err != nil {
			logger.Errorf("Error applying post-translation plugin: %v", err)
			continue
		}
	}
}
