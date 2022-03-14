// Code generated by solo-kit. DO NOT EDIT.

package gloosnapshot

import (
	"sync"
	"time"

	gateway_solo_io "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	github_com_solo_io_gloo_projects_gloo_pkg_api_external_solo_ratelimit "github.com/solo-io/gloo/projects/gloo/pkg/api/external/solo/ratelimit"
	gloo_solo_io "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	enterprise_gloo_solo_io "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/extauth/v1"
	graphql_gloo_solo_io "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/graphql/v1alpha1"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
	"go.uber.org/zap"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/errors"
	skstats "github.com/solo-io/solo-kit/pkg/stats"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/go-utils/errutils"
)

var (
	// Deprecated. See mApiResourcesIn
	mApiSnapshotIn = stats.Int64("api.gloosnapshot.gloo.solo.io/emitter/snap_in", "Deprecated. Use api.gloosnapshot.gloo.solo.io/emitter/resources_in. The number of snapshots in", "1")

	// metrics for emitter
	mApiResourcesIn    = stats.Int64("api.gloosnapshot.gloo.solo.io/emitter/resources_in", "The number of resource lists received on open watch channels", "1")
	mApiSnapshotOut    = stats.Int64("api.gloosnapshot.gloo.solo.io/emitter/snap_out", "The number of snapshots out", "1")
	mApiSnapshotMissed = stats.Int64("api.gloosnapshot.gloo.solo.io/emitter/snap_missed", "The number of snapshots missed", "1")

	// views for emitter
	// deprecated: see apiResourcesInView
	apisnapshotInView = &view.View{
		Name:        "api.gloosnapshot.gloo.solo.io/emitter/snap_in",
		Measure:     mApiSnapshotIn,
		Description: "Deprecated. Use api.gloosnapshot.gloo.solo.io/emitter/resources_in. The number of snapshots updates coming in.",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}

	apiResourcesInView = &view.View{
		Name:        "api.gloosnapshot.gloo.solo.io/emitter/resources_in",
		Measure:     mApiResourcesIn,
		Description: "The number of resource lists received on open watch channels",
		Aggregation: view.Count(),
		TagKeys: []tag.Key{
			skstats.NamespaceKey,
			skstats.ResourceKey,
		},
	}
	apisnapshotOutView = &view.View{
		Name:        "api.gloosnapshot.gloo.solo.io/emitter/snap_out",
		Measure:     mApiSnapshotOut,
		Description: "The number of snapshots updates going out",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
	apisnapshotMissedView = &view.View{
		Name:        "api.gloosnapshot.gloo.solo.io/emitter/snap_missed",
		Measure:     mApiSnapshotMissed,
		Description: "The number of snapshots updates going missed. this can happen in heavy load. missed snapshot will be re-tried after a second.",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
)

func init() {
	view.Register(
		apisnapshotInView,
		apisnapshotOutView,
		apisnapshotMissedView,
		apiResourcesInView,
	)
}

type ApiSnapshotEmitter interface {
	Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *ApiSnapshot, <-chan error, error)
}

type ApiEmitter interface {
	ApiSnapshotEmitter
	Register() error
	Artifact() gloo_solo_io.ArtifactClient
	Endpoint() gloo_solo_io.EndpointClient
	Proxy() gloo_solo_io.ProxyClient
	UpstreamGroup() gloo_solo_io.UpstreamGroupClient
	Secret() gloo_solo_io.SecretClient
	Upstream() gloo_solo_io.UpstreamClient
	AuthConfig() enterprise_gloo_solo_io.AuthConfigClient
	RateLimitConfig() github_com_solo_io_gloo_projects_gloo_pkg_api_external_solo_ratelimit.RateLimitConfigClient
	VirtualService() gateway_solo_io.VirtualServiceClient
	RouteTable() gateway_solo_io.RouteTableClient
	Gateway() gateway_solo_io.GatewayClient
	VirtualHostOption() gateway_solo_io.VirtualHostOptionClient
	RouteOption() gateway_solo_io.RouteOptionClient
	GraphQLApi() graphql_gloo_solo_io.GraphQLApiClient
}

func NewApiEmitter(artifactClient gloo_solo_io.ArtifactClient, endpointClient gloo_solo_io.EndpointClient, proxyClient gloo_solo_io.ProxyClient, upstreamGroupClient gloo_solo_io.UpstreamGroupClient, secretClient gloo_solo_io.SecretClient, upstreamClient gloo_solo_io.UpstreamClient, authConfigClient enterprise_gloo_solo_io.AuthConfigClient, rateLimitConfigClient github_com_solo_io_gloo_projects_gloo_pkg_api_external_solo_ratelimit.RateLimitConfigClient, virtualServiceClient gateway_solo_io.VirtualServiceClient, routeTableClient gateway_solo_io.RouteTableClient, gatewayClient gateway_solo_io.GatewayClient, virtualHostOptionClient gateway_solo_io.VirtualHostOptionClient, routeOptionClient gateway_solo_io.RouteOptionClient, graphQLApiClient graphql_gloo_solo_io.GraphQLApiClient) ApiEmitter {
	return NewApiEmitterWithEmit(artifactClient, endpointClient, proxyClient, upstreamGroupClient, secretClient, upstreamClient, authConfigClient, rateLimitConfigClient, virtualServiceClient, routeTableClient, gatewayClient, virtualHostOptionClient, routeOptionClient, graphQLApiClient, make(chan struct{}))
}

func NewApiEmitterWithEmit(artifactClient gloo_solo_io.ArtifactClient, endpointClient gloo_solo_io.EndpointClient, proxyClient gloo_solo_io.ProxyClient, upstreamGroupClient gloo_solo_io.UpstreamGroupClient, secretClient gloo_solo_io.SecretClient, upstreamClient gloo_solo_io.UpstreamClient, authConfigClient enterprise_gloo_solo_io.AuthConfigClient, rateLimitConfigClient github_com_solo_io_gloo_projects_gloo_pkg_api_external_solo_ratelimit.RateLimitConfigClient, virtualServiceClient gateway_solo_io.VirtualServiceClient, routeTableClient gateway_solo_io.RouteTableClient, gatewayClient gateway_solo_io.GatewayClient, virtualHostOptionClient gateway_solo_io.VirtualHostOptionClient, routeOptionClient gateway_solo_io.RouteOptionClient, graphQLApiClient graphql_gloo_solo_io.GraphQLApiClient, emit <-chan struct{}) ApiEmitter {
	return &apiEmitter{
		artifact:          artifactClient,
		endpoint:          endpointClient,
		proxy:             proxyClient,
		upstreamGroup:     upstreamGroupClient,
		secret:            secretClient,
		upstream:          upstreamClient,
		authConfig:        authConfigClient,
		rateLimitConfig:   rateLimitConfigClient,
		virtualService:    virtualServiceClient,
		routeTable:        routeTableClient,
		gateway:           gatewayClient,
		virtualHostOption: virtualHostOptionClient,
		routeOption:       routeOptionClient,
		graphQLApi:        graphQLApiClient,
		forceEmit:         emit,
	}
}

type apiEmitter struct {
	forceEmit         <-chan struct{}
	artifact          gloo_solo_io.ArtifactClient
	endpoint          gloo_solo_io.EndpointClient
	proxy             gloo_solo_io.ProxyClient
	upstreamGroup     gloo_solo_io.UpstreamGroupClient
	secret            gloo_solo_io.SecretClient
	upstream          gloo_solo_io.UpstreamClient
	authConfig        enterprise_gloo_solo_io.AuthConfigClient
	rateLimitConfig   github_com_solo_io_gloo_projects_gloo_pkg_api_external_solo_ratelimit.RateLimitConfigClient
	virtualService    gateway_solo_io.VirtualServiceClient
	routeTable        gateway_solo_io.RouteTableClient
	gateway           gateway_solo_io.GatewayClient
	virtualHostOption gateway_solo_io.VirtualHostOptionClient
	routeOption       gateway_solo_io.RouteOptionClient
	graphQLApi        graphql_gloo_solo_io.GraphQLApiClient
}

func (c *apiEmitter) Register() error {
	if err := c.artifact.Register(); err != nil {
		return err
	}
	if err := c.endpoint.Register(); err != nil {
		return err
	}
	if err := c.proxy.Register(); err != nil {
		return err
	}
	if err := c.upstreamGroup.Register(); err != nil {
		return err
	}
	if err := c.secret.Register(); err != nil {
		return err
	}
	if err := c.upstream.Register(); err != nil {
		return err
	}
	if err := c.authConfig.Register(); err != nil {
		return err
	}
	if err := c.rateLimitConfig.Register(); err != nil {
		return err
	}
	if err := c.virtualService.Register(); err != nil {
		return err
	}
	if err := c.routeTable.Register(); err != nil {
		return err
	}
	if err := c.gateway.Register(); err != nil {
		return err
	}
	if err := c.virtualHostOption.Register(); err != nil {
		return err
	}
	if err := c.routeOption.Register(); err != nil {
		return err
	}
	if err := c.graphQLApi.Register(); err != nil {
		return err
	}
	return nil
}

func (c *apiEmitter) Artifact() gloo_solo_io.ArtifactClient {
	return c.artifact
}

func (c *apiEmitter) Endpoint() gloo_solo_io.EndpointClient {
	return c.endpoint
}

func (c *apiEmitter) Proxy() gloo_solo_io.ProxyClient {
	return c.proxy
}

func (c *apiEmitter) UpstreamGroup() gloo_solo_io.UpstreamGroupClient {
	return c.upstreamGroup
}

func (c *apiEmitter) Secret() gloo_solo_io.SecretClient {
	return c.secret
}

func (c *apiEmitter) Upstream() gloo_solo_io.UpstreamClient {
	return c.upstream
}

func (c *apiEmitter) AuthConfig() enterprise_gloo_solo_io.AuthConfigClient {
	return c.authConfig
}

func (c *apiEmitter) RateLimitConfig() github_com_solo_io_gloo_projects_gloo_pkg_api_external_solo_ratelimit.RateLimitConfigClient {
	return c.rateLimitConfig
}

func (c *apiEmitter) VirtualService() gateway_solo_io.VirtualServiceClient {
	return c.virtualService
}

func (c *apiEmitter) RouteTable() gateway_solo_io.RouteTableClient {
	return c.routeTable
}

func (c *apiEmitter) Gateway() gateway_solo_io.GatewayClient {
	return c.gateway
}

func (c *apiEmitter) VirtualHostOption() gateway_solo_io.VirtualHostOptionClient {
	return c.virtualHostOption
}

func (c *apiEmitter) RouteOption() gateway_solo_io.RouteOptionClient {
	return c.routeOption
}

func (c *apiEmitter) GraphQLApi() graphql_gloo_solo_io.GraphQLApiClient {
	return c.graphQLApi
}

func (c *apiEmitter) Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *ApiSnapshot, <-chan error, error) {

	if len(watchNamespaces) == 0 {
		watchNamespaces = []string{""}
	}

	for _, ns := range watchNamespaces {
		if ns == "" && len(watchNamespaces) > 1 {
			return nil, nil, errors.Errorf("the \"\" namespace is used to watch all namespaces. Snapshots can either be tracked for " +
				"specific namespaces or \"\" AllNamespaces, but not both.")
		}
	}

	errs := make(chan error)
	var done sync.WaitGroup
	ctx := opts.Ctx
	/* Create channel for Artifact */
	type artifactListWithNamespace struct {
		list      gloo_solo_io.ArtifactList
		namespace string
	}
	artifactChan := make(chan artifactListWithNamespace)

	var initialArtifactList gloo_solo_io.ArtifactList
	/* Create channel for Endpoint */
	type endpointListWithNamespace struct {
		list      gloo_solo_io.EndpointList
		namespace string
	}
	endpointChan := make(chan endpointListWithNamespace)

	var initialEndpointList gloo_solo_io.EndpointList
	/* Create channel for Proxy */
	type proxyListWithNamespace struct {
		list      gloo_solo_io.ProxyList
		namespace string
	}
	proxyChan := make(chan proxyListWithNamespace)

	var initialProxyList gloo_solo_io.ProxyList
	/* Create channel for UpstreamGroup */
	type upstreamGroupListWithNamespace struct {
		list      gloo_solo_io.UpstreamGroupList
		namespace string
	}
	upstreamGroupChan := make(chan upstreamGroupListWithNamespace)

	var initialUpstreamGroupList gloo_solo_io.UpstreamGroupList
	/* Create channel for Secret */
	type secretListWithNamespace struct {
		list      gloo_solo_io.SecretList
		namespace string
	}
	secretChan := make(chan secretListWithNamespace)

	var initialSecretList gloo_solo_io.SecretList
	/* Create channel for Upstream */
	type upstreamListWithNamespace struct {
		list      gloo_solo_io.UpstreamList
		namespace string
	}
	upstreamChan := make(chan upstreamListWithNamespace)

	var initialUpstreamList gloo_solo_io.UpstreamList
	/* Create channel for AuthConfig */
	type authConfigListWithNamespace struct {
		list      enterprise_gloo_solo_io.AuthConfigList
		namespace string
	}
	authConfigChan := make(chan authConfigListWithNamespace)

	var initialAuthConfigList enterprise_gloo_solo_io.AuthConfigList
	/* Create channel for RateLimitConfig */
	type rateLimitConfigListWithNamespace struct {
		list      github_com_solo_io_gloo_projects_gloo_pkg_api_external_solo_ratelimit.RateLimitConfigList
		namespace string
	}
	rateLimitConfigChan := make(chan rateLimitConfigListWithNamespace)

	var initialRateLimitConfigList github_com_solo_io_gloo_projects_gloo_pkg_api_external_solo_ratelimit.RateLimitConfigList
	/* Create channel for VirtualService */
	type virtualServiceListWithNamespace struct {
		list      gateway_solo_io.VirtualServiceList
		namespace string
	}
	virtualServiceChan := make(chan virtualServiceListWithNamespace)

	var initialVirtualServiceList gateway_solo_io.VirtualServiceList
	/* Create channel for RouteTable */
	type routeTableListWithNamespace struct {
		list      gateway_solo_io.RouteTableList
		namespace string
	}
	routeTableChan := make(chan routeTableListWithNamespace)

	var initialRouteTableList gateway_solo_io.RouteTableList
	/* Create channel for Gateway */
	type gatewayListWithNamespace struct {
		list      gateway_solo_io.GatewayList
		namespace string
	}
	gatewayChan := make(chan gatewayListWithNamespace)

	var initialGatewayList gateway_solo_io.GatewayList
	/* Create channel for VirtualHostOption */
	type virtualHostOptionListWithNamespace struct {
		list      gateway_solo_io.VirtualHostOptionList
		namespace string
	}
	virtualHostOptionChan := make(chan virtualHostOptionListWithNamespace)

	var initialVirtualHostOptionList gateway_solo_io.VirtualHostOptionList
	/* Create channel for RouteOption */
	type routeOptionListWithNamespace struct {
		list      gateway_solo_io.RouteOptionList
		namespace string
	}
	routeOptionChan := make(chan routeOptionListWithNamespace)

	var initialRouteOptionList gateway_solo_io.RouteOptionList
	/* Create channel for GraphQLApi */
	type graphQLApiListWithNamespace struct {
		list      graphql_gloo_solo_io.GraphQLApiList
		namespace string
	}
	graphQLApiChan := make(chan graphQLApiListWithNamespace)

	var initialGraphQLApiList graphql_gloo_solo_io.GraphQLApiList

	currentSnapshot := ApiSnapshot{}

	for _, namespace := range watchNamespaces {
		/* Setup namespaced watch for Artifact */
		{
			artifacts, err := c.artifact.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial Artifact list")
			}
			initialArtifactList = append(initialArtifactList, artifacts...)
		}
		artifactNamespacesChan, artifactErrs, err := c.artifact.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting Artifact watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, artifactErrs, namespace+"-artifacts")
		}(namespace)
		/* Setup namespaced watch for Endpoint */
		{
			endpoints, err := c.endpoint.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial Endpoint list")
			}
			initialEndpointList = append(initialEndpointList, endpoints...)
		}
		endpointNamespacesChan, endpointErrs, err := c.endpoint.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting Endpoint watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, endpointErrs, namespace+"-endpoints")
		}(namespace)
		/* Setup namespaced watch for Proxy */
		{
			proxies, err := c.proxy.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial Proxy list")
			}
			initialProxyList = append(initialProxyList, proxies...)
		}
		proxyNamespacesChan, proxyErrs, err := c.proxy.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting Proxy watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, proxyErrs, namespace+"-proxies")
		}(namespace)
		/* Setup namespaced watch for UpstreamGroup */
		{
			upstreamGroups, err := c.upstreamGroup.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial UpstreamGroup list")
			}
			initialUpstreamGroupList = append(initialUpstreamGroupList, upstreamGroups...)
		}
		upstreamGroupNamespacesChan, upstreamGroupErrs, err := c.upstreamGroup.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting UpstreamGroup watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, upstreamGroupErrs, namespace+"-upstreamGroups")
		}(namespace)
		/* Setup namespaced watch for Secret */
		{
			secrets, err := c.secret.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial Secret list")
			}
			initialSecretList = append(initialSecretList, secrets...)
		}
		secretNamespacesChan, secretErrs, err := c.secret.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting Secret watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, secretErrs, namespace+"-secrets")
		}(namespace)
		/* Setup namespaced watch for Upstream */
		{
			upstreams, err := c.upstream.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial Upstream list")
			}
			initialUpstreamList = append(initialUpstreamList, upstreams...)
		}
		upstreamNamespacesChan, upstreamErrs, err := c.upstream.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting Upstream watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, upstreamErrs, namespace+"-upstreams")
		}(namespace)
		/* Setup namespaced watch for AuthConfig */
		{
			authConfigs, err := c.authConfig.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial AuthConfig list")
			}
			initialAuthConfigList = append(initialAuthConfigList, authConfigs...)
		}
		authConfigNamespacesChan, authConfigErrs, err := c.authConfig.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting AuthConfig watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, authConfigErrs, namespace+"-authConfigs")
		}(namespace)
		/* Setup namespaced watch for RateLimitConfig */
		{
			ratelimitconfigs, err := c.rateLimitConfig.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial RateLimitConfig list")
			}
			initialRateLimitConfigList = append(initialRateLimitConfigList, ratelimitconfigs...)
		}
		rateLimitConfigNamespacesChan, rateLimitConfigErrs, err := c.rateLimitConfig.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting RateLimitConfig watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, rateLimitConfigErrs, namespace+"-ratelimitconfigs")
		}(namespace)
		/* Setup namespaced watch for VirtualService */
		{
			virtualServices, err := c.virtualService.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial VirtualService list")
			}
			initialVirtualServiceList = append(initialVirtualServiceList, virtualServices...)
		}
		virtualServiceNamespacesChan, virtualServiceErrs, err := c.virtualService.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting VirtualService watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, virtualServiceErrs, namespace+"-virtualServices")
		}(namespace)
		/* Setup namespaced watch for RouteTable */
		{
			routeTables, err := c.routeTable.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial RouteTable list")
			}
			initialRouteTableList = append(initialRouteTableList, routeTables...)
		}
		routeTableNamespacesChan, routeTableErrs, err := c.routeTable.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting RouteTable watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, routeTableErrs, namespace+"-routeTables")
		}(namespace)
		/* Setup namespaced watch for Gateway */
		{
			gateways, err := c.gateway.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial Gateway list")
			}
			initialGatewayList = append(initialGatewayList, gateways...)
		}
		gatewayNamespacesChan, gatewayErrs, err := c.gateway.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting Gateway watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, gatewayErrs, namespace+"-gateways")
		}(namespace)
		/* Setup namespaced watch for VirtualHostOption */
		{
			virtualHostOptions, err := c.virtualHostOption.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial VirtualHostOption list")
			}
			initialVirtualHostOptionList = append(initialVirtualHostOptionList, virtualHostOptions...)
		}
		virtualHostOptionNamespacesChan, virtualHostOptionErrs, err := c.virtualHostOption.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting VirtualHostOption watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, virtualHostOptionErrs, namespace+"-virtualHostOptions")
		}(namespace)
		/* Setup namespaced watch for RouteOption */
		{
			routeOptions, err := c.routeOption.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial RouteOption list")
			}
			initialRouteOptionList = append(initialRouteOptionList, routeOptions...)
		}
		routeOptionNamespacesChan, routeOptionErrs, err := c.routeOption.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting RouteOption watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, routeOptionErrs, namespace+"-routeOptions")
		}(namespace)
		/* Setup namespaced watch for GraphQLApi */
		{
			graphqlApis, err := c.graphQLApi.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial GraphQLApi list")
			}
			initialGraphQLApiList = append(initialGraphQLApiList, graphqlApis...)
		}
		graphQLApiNamespacesChan, graphQLApiErrs, err := c.graphQLApi.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting GraphQLApi watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, graphQLApiErrs, namespace+"-graphqlApis")
		}(namespace)

		/* Watch for changes and update snapshot */
		go func(namespace string) {
			for {
				select {
				case <-ctx.Done():
					return
				case artifactList, ok := <-artifactNamespacesChan:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case artifactChan <- artifactListWithNamespace{list: artifactList, namespace: namespace}:
					}
				case endpointList, ok := <-endpointNamespacesChan:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case endpointChan <- endpointListWithNamespace{list: endpointList, namespace: namespace}:
					}
				case proxyList, ok := <-proxyNamespacesChan:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case proxyChan <- proxyListWithNamespace{list: proxyList, namespace: namespace}:
					}
				case upstreamGroupList, ok := <-upstreamGroupNamespacesChan:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case upstreamGroupChan <- upstreamGroupListWithNamespace{list: upstreamGroupList, namespace: namespace}:
					}
				case secretList, ok := <-secretNamespacesChan:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case secretChan <- secretListWithNamespace{list: secretList, namespace: namespace}:
					}
				case upstreamList, ok := <-upstreamNamespacesChan:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case upstreamChan <- upstreamListWithNamespace{list: upstreamList, namespace: namespace}:
					}
				case authConfigList, ok := <-authConfigNamespacesChan:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case authConfigChan <- authConfigListWithNamespace{list: authConfigList, namespace: namespace}:
					}
				case rateLimitConfigList, ok := <-rateLimitConfigNamespacesChan:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case rateLimitConfigChan <- rateLimitConfigListWithNamespace{list: rateLimitConfigList, namespace: namespace}:
					}
				case virtualServiceList, ok := <-virtualServiceNamespacesChan:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case virtualServiceChan <- virtualServiceListWithNamespace{list: virtualServiceList, namespace: namespace}:
					}
				case routeTableList, ok := <-routeTableNamespacesChan:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case routeTableChan <- routeTableListWithNamespace{list: routeTableList, namespace: namespace}:
					}
				case gatewayList, ok := <-gatewayNamespacesChan:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case gatewayChan <- gatewayListWithNamespace{list: gatewayList, namespace: namespace}:
					}
				case virtualHostOptionList, ok := <-virtualHostOptionNamespacesChan:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case virtualHostOptionChan <- virtualHostOptionListWithNamespace{list: virtualHostOptionList, namespace: namespace}:
					}
				case routeOptionList, ok := <-routeOptionNamespacesChan:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case routeOptionChan <- routeOptionListWithNamespace{list: routeOptionList, namespace: namespace}:
					}
				case graphQLApiList, ok := <-graphQLApiNamespacesChan:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case graphQLApiChan <- graphQLApiListWithNamespace{list: graphQLApiList, namespace: namespace}:
					}
				}
			}
		}(namespace)
	}
	/* Initialize snapshot for Artifacts */
	currentSnapshot.Artifacts = initialArtifactList.Sort()
	/* Initialize snapshot for Endpoints */
	currentSnapshot.Endpoints = initialEndpointList.Sort()
	/* Initialize snapshot for Proxies */
	currentSnapshot.Proxies = initialProxyList.Sort()
	/* Initialize snapshot for UpstreamGroups */
	currentSnapshot.UpstreamGroups = initialUpstreamGroupList.Sort()
	/* Initialize snapshot for Secrets */
	currentSnapshot.Secrets = initialSecretList.Sort()
	/* Initialize snapshot for Upstreams */
	currentSnapshot.Upstreams = initialUpstreamList.Sort()
	/* Initialize snapshot for AuthConfigs */
	currentSnapshot.AuthConfigs = initialAuthConfigList.Sort()
	/* Initialize snapshot for Ratelimitconfigs */
	currentSnapshot.Ratelimitconfigs = initialRateLimitConfigList.Sort()
	/* Initialize snapshot for VirtualServices */
	currentSnapshot.VirtualServices = initialVirtualServiceList.Sort()
	/* Initialize snapshot for RouteTables */
	currentSnapshot.RouteTables = initialRouteTableList.Sort()
	/* Initialize snapshot for Gateways */
	currentSnapshot.Gateways = initialGatewayList.Sort()
	/* Initialize snapshot for VirtualHostOptions */
	currentSnapshot.VirtualHostOptions = initialVirtualHostOptionList.Sort()
	/* Initialize snapshot for RouteOptions */
	currentSnapshot.RouteOptions = initialRouteOptionList.Sort()
	/* Initialize snapshot for GraphqlApis */
	currentSnapshot.GraphqlApis = initialGraphQLApiList.Sort()

	snapshots := make(chan *ApiSnapshot)
	go func() {
		// sent initial snapshot to kick off the watch
		initialSnapshot := currentSnapshot.Clone()
		snapshots <- &initialSnapshot

		timer := time.NewTicker(time.Second * 1)
		previousHash, err := currentSnapshot.Hash(nil)
		if err != nil {
			contextutils.LoggerFrom(ctx).Panicw("error while hashing, this should never happen", zap.Error(err))
		}
		sync := func() {
			currentHash, err := currentSnapshot.Hash(nil)
			// this should never happen, so panic if it does
			if err != nil {
				contextutils.LoggerFrom(ctx).Panicw("error while hashing, this should never happen", zap.Error(err))
			}
			if previousHash == currentHash {
				return
			}

			sentSnapshot := currentSnapshot.Clone()
			select {
			case snapshots <- &sentSnapshot:
				stats.Record(ctx, mApiSnapshotOut.M(1))
				previousHash = currentHash
			default:
				stats.Record(ctx, mApiSnapshotMissed.M(1))
			}
		}
		artifactsByNamespace := make(map[string]gloo_solo_io.ArtifactList)
		endpointsByNamespace := make(map[string]gloo_solo_io.EndpointList)
		proxiesByNamespace := make(map[string]gloo_solo_io.ProxyList)
		upstreamGroupsByNamespace := make(map[string]gloo_solo_io.UpstreamGroupList)
		secretsByNamespace := make(map[string]gloo_solo_io.SecretList)
		upstreamsByNamespace := make(map[string]gloo_solo_io.UpstreamList)
		authConfigsByNamespace := make(map[string]enterprise_gloo_solo_io.AuthConfigList)
		ratelimitconfigsByNamespace := make(map[string]github_com_solo_io_gloo_projects_gloo_pkg_api_external_solo_ratelimit.RateLimitConfigList)
		virtualServicesByNamespace := make(map[string]gateway_solo_io.VirtualServiceList)
		routeTablesByNamespace := make(map[string]gateway_solo_io.RouteTableList)
		gatewaysByNamespace := make(map[string]gateway_solo_io.GatewayList)
		virtualHostOptionsByNamespace := make(map[string]gateway_solo_io.VirtualHostOptionList)
		routeOptionsByNamespace := make(map[string]gateway_solo_io.RouteOptionList)
		graphqlApisByNamespace := make(map[string]graphql_gloo_solo_io.GraphQLApiList)
		defer func() {
			close(snapshots)
			// we must wait for done before closing the error chan,
			// to avoid sending on close channel.
			done.Wait()
			close(errs)
		}()
		for {
			record := func() { stats.Record(ctx, mApiSnapshotIn.M(1)) }

			select {
			case <-timer.C:
				sync()
			case <-ctx.Done():
				return
			case <-c.forceEmit:
				sentSnapshot := currentSnapshot.Clone()
				snapshots <- &sentSnapshot
			case artifactNamespacedList, ok := <-artifactChan:
				if !ok {
					return
				}
				record()

				namespace := artifactNamespacedList.namespace

				skstats.IncrementResourceCount(
					ctx,
					namespace,
					"artifact",
					mApiResourcesIn,
				)

				// merge lists by namespace
				artifactsByNamespace[namespace] = artifactNamespacedList.list
				var artifactList gloo_solo_io.ArtifactList
				for _, artifacts := range artifactsByNamespace {
					artifactList = append(artifactList, artifacts...)
				}
				currentSnapshot.Artifacts = artifactList.Sort()
			case endpointNamespacedList, ok := <-endpointChan:
				if !ok {
					return
				}
				record()

				namespace := endpointNamespacedList.namespace

				skstats.IncrementResourceCount(
					ctx,
					namespace,
					"endpoint",
					mApiResourcesIn,
				)

				// merge lists by namespace
				endpointsByNamespace[namespace] = endpointNamespacedList.list
				var endpointList gloo_solo_io.EndpointList
				for _, endpoints := range endpointsByNamespace {
					endpointList = append(endpointList, endpoints...)
				}
				currentSnapshot.Endpoints = endpointList.Sort()
			case proxyNamespacedList, ok := <-proxyChan:
				if !ok {
					return
				}
				record()

				namespace := proxyNamespacedList.namespace

				skstats.IncrementResourceCount(
					ctx,
					namespace,
					"proxy",
					mApiResourcesIn,
				)

				// merge lists by namespace
				proxiesByNamespace[namespace] = proxyNamespacedList.list
				var proxyList gloo_solo_io.ProxyList
				for _, proxies := range proxiesByNamespace {
					proxyList = append(proxyList, proxies...)
				}
				currentSnapshot.Proxies = proxyList.Sort()
			case upstreamGroupNamespacedList, ok := <-upstreamGroupChan:
				if !ok {
					return
				}
				record()

				namespace := upstreamGroupNamespacedList.namespace

				skstats.IncrementResourceCount(
					ctx,
					namespace,
					"upstream_group",
					mApiResourcesIn,
				)

				// merge lists by namespace
				upstreamGroupsByNamespace[namespace] = upstreamGroupNamespacedList.list
				var upstreamGroupList gloo_solo_io.UpstreamGroupList
				for _, upstreamGroups := range upstreamGroupsByNamespace {
					upstreamGroupList = append(upstreamGroupList, upstreamGroups...)
				}
				currentSnapshot.UpstreamGroups = upstreamGroupList.Sort()
			case secretNamespacedList, ok := <-secretChan:
				if !ok {
					return
				}
				record()

				namespace := secretNamespacedList.namespace

				skstats.IncrementResourceCount(
					ctx,
					namespace,
					"secret",
					mApiResourcesIn,
				)

				// merge lists by namespace
				secretsByNamespace[namespace] = secretNamespacedList.list
				var secretList gloo_solo_io.SecretList
				for _, secrets := range secretsByNamespace {
					secretList = append(secretList, secrets...)
				}
				currentSnapshot.Secrets = secretList.Sort()
			case upstreamNamespacedList, ok := <-upstreamChan:
				if !ok {
					return
				}
				record()

				namespace := upstreamNamespacedList.namespace

				skstats.IncrementResourceCount(
					ctx,
					namespace,
					"upstream",
					mApiResourcesIn,
				)

				// merge lists by namespace
				upstreamsByNamespace[namespace] = upstreamNamespacedList.list
				var upstreamList gloo_solo_io.UpstreamList
				for _, upstreams := range upstreamsByNamespace {
					upstreamList = append(upstreamList, upstreams...)
				}
				currentSnapshot.Upstreams = upstreamList.Sort()
			case authConfigNamespacedList, ok := <-authConfigChan:
				if !ok {
					return
				}
				record()

				namespace := authConfigNamespacedList.namespace

				skstats.IncrementResourceCount(
					ctx,
					namespace,
					"auth_config",
					mApiResourcesIn,
				)

				// merge lists by namespace
				authConfigsByNamespace[namespace] = authConfigNamespacedList.list
				var authConfigList enterprise_gloo_solo_io.AuthConfigList
				for _, authConfigs := range authConfigsByNamespace {
					authConfigList = append(authConfigList, authConfigs...)
				}
				currentSnapshot.AuthConfigs = authConfigList.Sort()
			case rateLimitConfigNamespacedList, ok := <-rateLimitConfigChan:
				if !ok {
					return
				}
				record()

				namespace := rateLimitConfigNamespacedList.namespace

				skstats.IncrementResourceCount(
					ctx,
					namespace,
					"rate_limit_config",
					mApiResourcesIn,
				)

				// merge lists by namespace
				ratelimitconfigsByNamespace[namespace] = rateLimitConfigNamespacedList.list
				var rateLimitConfigList github_com_solo_io_gloo_projects_gloo_pkg_api_external_solo_ratelimit.RateLimitConfigList
				for _, ratelimitconfigs := range ratelimitconfigsByNamespace {
					rateLimitConfigList = append(rateLimitConfigList, ratelimitconfigs...)
				}
				currentSnapshot.Ratelimitconfigs = rateLimitConfigList.Sort()
			case virtualServiceNamespacedList, ok := <-virtualServiceChan:
				if !ok {
					return
				}
				record()

				namespace := virtualServiceNamespacedList.namespace

				skstats.IncrementResourceCount(
					ctx,
					namespace,
					"virtual_service",
					mApiResourcesIn,
				)

				// merge lists by namespace
				virtualServicesByNamespace[namespace] = virtualServiceNamespacedList.list
				var virtualServiceList gateway_solo_io.VirtualServiceList
				for _, virtualServices := range virtualServicesByNamespace {
					virtualServiceList = append(virtualServiceList, virtualServices...)
				}
				currentSnapshot.VirtualServices = virtualServiceList.Sort()
			case routeTableNamespacedList, ok := <-routeTableChan:
				if !ok {
					return
				}
				record()

				namespace := routeTableNamespacedList.namespace

				skstats.IncrementResourceCount(
					ctx,
					namespace,
					"route_table",
					mApiResourcesIn,
				)

				// merge lists by namespace
				routeTablesByNamespace[namespace] = routeTableNamespacedList.list
				var routeTableList gateway_solo_io.RouteTableList
				for _, routeTables := range routeTablesByNamespace {
					routeTableList = append(routeTableList, routeTables...)
				}
				currentSnapshot.RouteTables = routeTableList.Sort()
			case gatewayNamespacedList, ok := <-gatewayChan:
				if !ok {
					return
				}
				record()

				namespace := gatewayNamespacedList.namespace

				skstats.IncrementResourceCount(
					ctx,
					namespace,
					"gateway",
					mApiResourcesIn,
				)

				// merge lists by namespace
				gatewaysByNamespace[namespace] = gatewayNamespacedList.list
				var gatewayList gateway_solo_io.GatewayList
				for _, gateways := range gatewaysByNamespace {
					gatewayList = append(gatewayList, gateways...)
				}
				currentSnapshot.Gateways = gatewayList.Sort()
			case virtualHostOptionNamespacedList, ok := <-virtualHostOptionChan:
				if !ok {
					return
				}
				record()

				namespace := virtualHostOptionNamespacedList.namespace

				skstats.IncrementResourceCount(
					ctx,
					namespace,
					"virtual_host_option",
					mApiResourcesIn,
				)

				// merge lists by namespace
				virtualHostOptionsByNamespace[namespace] = virtualHostOptionNamespacedList.list
				var virtualHostOptionList gateway_solo_io.VirtualHostOptionList
				for _, virtualHostOptions := range virtualHostOptionsByNamespace {
					virtualHostOptionList = append(virtualHostOptionList, virtualHostOptions...)
				}
				currentSnapshot.VirtualHostOptions = virtualHostOptionList.Sort()
			case routeOptionNamespacedList, ok := <-routeOptionChan:
				if !ok {
					return
				}
				record()

				namespace := routeOptionNamespacedList.namespace

				skstats.IncrementResourceCount(
					ctx,
					namespace,
					"route_option",
					mApiResourcesIn,
				)

				// merge lists by namespace
				routeOptionsByNamespace[namespace] = routeOptionNamespacedList.list
				var routeOptionList gateway_solo_io.RouteOptionList
				for _, routeOptions := range routeOptionsByNamespace {
					routeOptionList = append(routeOptionList, routeOptions...)
				}
				currentSnapshot.RouteOptions = routeOptionList.Sort()
			case graphQLApiNamespacedList, ok := <-graphQLApiChan:
				if !ok {
					return
				}
				record()

				namespace := graphQLApiNamespacedList.namespace

				skstats.IncrementResourceCount(
					ctx,
					namespace,
					"graph_ql_api",
					mApiResourcesIn,
				)

				// merge lists by namespace
				graphqlApisByNamespace[namespace] = graphQLApiNamespacedList.list
				var graphQLApiList graphql_gloo_solo_io.GraphQLApiList
				for _, graphqlApis := range graphqlApisByNamespace {
					graphQLApiList = append(graphQLApiList, graphqlApis...)
				}
				currentSnapshot.GraphqlApis = graphQLApiList.Sort()
			}
		}
	}()
	return snapshots, errs, nil
}
