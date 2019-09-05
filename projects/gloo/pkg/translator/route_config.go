package translator

import (
	"strings"

	"github.com/gogo/protobuf/proto"

	usconversion "github.com/solo-io/gloo/projects/gloo/pkg/upstreams"

	envoyapi "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoycore "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	envoyroute "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	"github.com/gogo/protobuf/types"
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	v1plugins "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	"github.com/solo-io/gloo/projects/gloo/pkg/utils"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

var (
	NoDestinationSpecifiedError = errors.New("must specify at least one weighted destination for multi destination routes")
)

type reportFunc func(error error, format string, args ...interface{})

func (t *translator) computeRouteConfig(params plugins.Params, proxy *v1.Proxy, listener *v1.Listener, routeCfgName string, reportFn reportFunc) *envoyapi.RouteConfiguration {
	if listener.GetHttpListener() == nil {
		return nil
	}
	report := func(err error, format string, args ...interface{}) {
		reportFn(err, "route_config."+format, args...)
	}
	params.Ctx = contextutils.WithLogger(params.Ctx, "compute_route_config."+routeCfgName)

	virtualHosts := t.computeVirtualHosts(params, proxy, listener, report)

	// validate ssl config if the listener specifies any
	if err := validateListenerSslConfig(params, listener); err != nil {
		report(err, "invalid listener %v", listener.Name)
	}

	return &envoyapi.RouteConfiguration{
		Name:         routeCfgName,
		VirtualHosts: virtualHosts,
	}
}

func (t *translator) computeVirtualHosts(params plugins.Params, proxy *v1.Proxy, listener *v1.Listener, report reportFunc) []envoyroute.VirtualHost {
	httpListener, ok := listener.ListenerType.(*v1.Listener_HttpListener)
	if !ok {
		return nil
	}
	virtualHosts := httpListener.HttpListener.VirtualHosts
	if err := validateVirtualHostDomains(virtualHosts); err != nil {
		report(err, "invalid listener %v", listener.Name)
	}
	requireTls := len(listener.SslConfigurations) > 0
	var envoyVirtualHosts []envoyroute.VirtualHost
	for _, virtualHost := range virtualHosts {
		vhostParams := plugins.VirtualHostParams{
			Params:   params,
			Listener: listener,
			Proxy:    proxy,
		}
		envoyVirtualHosts = append(envoyVirtualHosts, t.computeVirtualHost(vhostParams, virtualHost, requireTls, report))
	}
	return envoyVirtualHosts
}

func (t *translator) computeVirtualHost(params plugins.VirtualHostParams, virtualHost *v1.VirtualHost, requireTls bool, report reportFunc) envoyroute.VirtualHost {

	// Make copy to avoid modifying the snapshot
	virtualHost = proto.Clone(virtualHost).(*v1.VirtualHost)
	virtualHost.Name = utils.SanitizeForEnvoy(params.Ctx, virtualHost.Name, "virtual host")

	var envoyRoutes []envoyroute.Route
	for _, route := range virtualHost.Routes {
		routeParams := plugins.RouteParams{
			VirtualHostParams: params,
			VirtualHost:       virtualHost,
		}
		envoyRoute := t.envoyRoute(routeParams, report, route)
		envoyRoutes = append(envoyRoutes, envoyRoute)
	}
	domains := virtualHost.Domains
	if len(domains) == 0 || (len(domains) == 1 && domains[0] == "") {
		domains = []string{"*"}
	}
	var envoyRequireTls envoyroute.VirtualHost_TlsRequirementType
	if requireTls {
		// TODO (ilackarms): support external-only TLS
		envoyRequireTls = envoyroute.VirtualHost_ALL
	}

	out := envoyroute.VirtualHost{
		Name:       virtualHost.Name,
		Domains:    domains,
		Routes:     envoyRoutes,
		RequireTls: envoyRequireTls,
	}

	// run the plugins
	for _, plug := range t.plugins {
		virtualHostPlugin, ok := plug.(plugins.VirtualHostPlugin)
		if !ok {
			continue
		}
		if err := virtualHostPlugin.ProcessVirtualHost(params, virtualHost, &out); err != nil {
			report(err, "invalid virtual host [%s]", virtualHost.Name)
		}
	}
	return out
}

func (t *translator) envoyRoute(params plugins.RouteParams, report reportFunc, in *v1.Route) envoyroute.Route {
	out := &envoyroute.Route{}

	setMatch(in, report, out)

	t.setAction(params, report, in, out)

	return *out
}

func setMatch(in *v1.Route, report reportFunc, out *envoyroute.Route) {

	if in.Matcher.PathSpecifier == nil {
		report(errors.New("no path specifier provided"), "invalid route")
	}
	match := envoyroute.RouteMatch{
		Headers:         envoyHeaderMatcher(in.Matcher.Headers),
		QueryParameters: envoyQueryMatcher(in.Matcher.QueryParameters),
	}
	if len(in.Matcher.Methods) > 0 {
		match.Headers = append(match.Headers, &envoyroute.HeaderMatcher{
			Name: ":method",
			HeaderMatchSpecifier: &envoyroute.HeaderMatcher_RegexMatch{
				RegexMatch: strings.Join(in.Matcher.Methods, "|"),
			},
		})
	}
	// need to do this because Go's proto implementation makes oneofs private
	// which genius thought of that?
	setEnvoyPathMatcher(in.Matcher, &match)

	out.Match = match
}

func (t *translator) setAction(params plugins.RouteParams, report reportFunc, in *v1.Route, out *envoyroute.Route) {
	switch action := in.Action.(type) {
	case *v1.Route_RouteAction:
		if err := ValidateRouteDestinations(params.Snapshot, action.RouteAction); err != nil {
			report(err, "invalid route")
		}

		out.Action = &envoyroute.Route_Route{
			Route: &envoyroute.RouteAction{},
		}
		if err := t.setRouteAction(params, action.RouteAction, out.Action.(*envoyroute.Route_Route).Route, report); err != nil {
			report(err, "translator error on route")
		}

		// run the plugins for RoutePlugin
		for _, plug := range t.plugins {
			routePlugin, ok := plug.(plugins.RoutePlugin)
			if !ok {
				continue
			}
			if err := routePlugin.ProcessRoute(params, in, out); err != nil {
				report(err, "plugin error on ProcessRoute")
			}
		}

		// run the plugins for RouteActionPlugin
		for _, plug := range t.plugins {
			routePlugin, ok := plug.(plugins.RouteActionPlugin)
			if !ok || in.GetRouteAction() == nil || out.GetRoute() == nil {
				continue
			}
			raParams := plugins.RouteActionParams{
				RouteParams: params,
				Route:       in,
			}
			if err := routePlugin.ProcessRouteAction(raParams, in.GetRouteAction(), out.GetRoute()); err != nil {
				report(err, "plugin error on ProcessRouteAction")
			}
		}

	case *v1.Route_DirectResponseAction:
		out.Action = &envoyroute.Route_DirectResponse{
			DirectResponse: &envoyroute.DirectResponseAction{
				Status: action.DirectResponseAction.Status,
				Body:   DataSourceFromString(action.DirectResponseAction.Body),
			},
		}

	case *v1.Route_RedirectAction:
		out.Action = &envoyroute.Route_Redirect{
			Redirect: &envoyroute.RedirectAction{
				HostRedirect:           action.RedirectAction.HostRedirect,
				ResponseCode:           envoyroute.RedirectAction_RedirectResponseCode(action.RedirectAction.ResponseCode),
				SchemeRewriteSpecifier: &envoyroute.RedirectAction_HttpsRedirect{HttpsRedirect: action.RedirectAction.HttpsRedirect},
				StripQuery:             action.RedirectAction.StripQuery,
			},
		}

		switch pathRewrite := action.RedirectAction.PathRewriteSpecifier.(type) {
		case *v1.RedirectAction_PathRedirect:
			out.Action.(*envoyroute.Route_Redirect).Redirect.PathRewriteSpecifier = &envoyroute.RedirectAction_PathRedirect{
				PathRedirect: pathRewrite.PathRedirect,
			}
		case *v1.RedirectAction_PrefixRewrite:
			out.Action.(*envoyroute.Route_Redirect).Redirect.PathRewriteSpecifier = &envoyroute.RedirectAction_PrefixRewrite{
				PrefixRewrite: pathRewrite.PrefixRewrite,
			}
		}
	}
}

func (t *translator) setRouteAction(params plugins.RouteParams, in *v1.RouteAction, out *envoyroute.RouteAction, report reportFunc) error {
	switch dest := in.Destination.(type) {
	case *v1.RouteAction_Single:
		usRef, err := usconversion.DestinationToUpstreamRef(dest.Single)
		if err != nil {
			return err
		}
		out.ClusterSpecifier = &envoyroute.RouteAction_Cluster{
			Cluster: UpstreamToClusterName(*usRef),
		}

		out.MetadataMatch = getSubsetMatch(dest.Single)

		return checkThatSubsetMatchesUpstream(params.Params, dest.Single)
	case *v1.RouteAction_Multi:
		return t.setWeightedClusters(params, dest.Multi, out, report)
	case *v1.RouteAction_UpstreamGroup:
		upstreamGroupRef := dest.UpstreamGroup
		upstreamGroup, err := params.Snapshot.UpstreamGroups.Find(upstreamGroupRef.Namespace, upstreamGroupRef.Name)
		if err != nil {
			return err
		}
		md := &v1.MultiDestination{
			Destinations: upstreamGroup.Destinations,
		}
		return t.setWeightedClusters(params, md, out, report)
	}
	return errors.Errorf("unknown upstream destination type")
}

func (t *translator) setWeightedClusters(params plugins.RouteParams, multiDest *v1.MultiDestination, out *envoyroute.RouteAction, report reportFunc) error {
	if len(multiDest.Destinations) == 0 {
		return NoDestinationSpecifiedError
	}

	clusterSpecifier := &envoyroute.RouteAction_WeightedClusters{
		WeightedClusters: &envoyroute.WeightedCluster{},
	}

	var totalWeight uint32
	for _, weightedDest := range multiDest.Destinations {

		usRef, err := usconversion.DestinationToUpstreamRef(weightedDest.Destination)
		if err != nil {
			return err
		}

		totalWeight += weightedDest.Weight

		weightedCluster := &envoyroute.WeightedCluster_ClusterWeight{
			Name:          UpstreamToClusterName(*usRef),
			Weight:        &types.UInt32Value{Value: weightedDest.Weight},
			MetadataMatch: getSubsetMatch(weightedDest.Destination),
		}

		// run the plugins for Weighted Destinations
		for _, plug := range t.plugins {
			weightedDestinationPlugin, ok := plug.(plugins.WeightedDestinationPlugin)
			if !ok {
				continue
			}
			if err := weightedDestinationPlugin.ProcessWeightedDestination(params, weightedDest, weightedCluster); err != nil {
				report(err, "plugin error on ProcessRoute")
			}
		}

		clusterSpecifier.WeightedClusters.Clusters = append(clusterSpecifier.WeightedClusters.Clusters, weightedCluster)

		if err = checkThatSubsetMatchesUpstream(params.Params, weightedDest.Destination); err != nil {
			return err
		}
	}

	clusterSpecifier.WeightedClusters.TotalWeight = &types.UInt32Value{Value: totalWeight}

	out.ClusterSpecifier = clusterSpecifier
	return nil
}

// TODO(marco): when we update the routing API we should move this to a RouteActionPlugin
func getSubsetMatch(destination *v1.Destination) *envoycore.Metadata {
	var routeMetadata *envoycore.Metadata

	// TODO(yuval-k): should we add validation that the route subset indeed exists in the upstream?
	// First convert the subset information on the base destination, if present
	if destination.Subset != nil {
		routeMetadata = getLbMetadata(nil, destination.Subset.Values, "")
	}
	return routeMetadata
}

func checkThatSubsetMatchesUpstream(params plugins.Params, dest *v1.Destination) error {

	// make sure we have a subset config on the route
	if dest.Subset == nil {
		return nil
	}
	if len(dest.Subset.Values) == 0 {
		return nil
	}
	routeSubset := dest.Subset.Values

	ref, err := usconversion.DestinationToUpstreamRef(dest)
	if err != nil {
		return err
	}

	upstream, err := params.Snapshot.Upstreams.Find(ref.Namespace, ref.Name)
	if err != nil {
		return err
	}

	subsetConfig := getSubsets(upstream)

	// if a route has a subset config, and an upstream doesnt - its an error
	if subsetConfig == nil {
		return errors.Errorf("route has a subset config, but the upstream does not.")
	}

	// make sure that the subset on the route will match a subset on the upstream.
	found := false
Outerloop:
	for _, subset := range subsetConfig.Selectors {
		keys := subset.Keys
		if len(keys) != len(routeSubset) {
			continue
		}
		for _, k := range keys {
			if _, ok := routeSubset[k]; !ok {
				continue Outerloop
			}
		}
		found = true
		break
	}

	if !found {
		return errors.Errorf("route has a subset config, but none of the subsets in the upstream match it.")

	}
	return nil
}

func getSubsets(upstream *v1.Upstream) *v1plugins.SubsetSpec {

	specGetter, ok := upstream.UpstreamSpec.UpstreamType.(v1.SubsetSpecGetter)
	if !ok {
		return nil
	}
	glooSubsetConfig := specGetter.GetSubsetSpec()

	return glooSubsetConfig

}

func setEnvoyPathMatcher(in *v1.Matcher, out *envoyroute.RouteMatch) {
	switch path := in.PathSpecifier.(type) {
	case *v1.Matcher_Exact:
		out.PathSpecifier = &envoyroute.RouteMatch_Path{
			Path: path.Exact,
		}
	case *v1.Matcher_Regex:
		out.PathSpecifier = &envoyroute.RouteMatch_Regex{
			Regex: path.Regex,
		}
	case *v1.Matcher_Prefix:
		out.PathSpecifier = &envoyroute.RouteMatch_Prefix{
			Prefix: path.Prefix,
		}
	}
}

func envoyHeaderMatcher(in []*v1.HeaderMatcher) []*envoyroute.HeaderMatcher {
	var out []*envoyroute.HeaderMatcher
	for _, matcher := range in {

		envoyMatch := &envoyroute.HeaderMatcher{
			Name: matcher.Name,
		}
		if matcher.Value == "" {
			envoyMatch.HeaderMatchSpecifier = &envoyroute.HeaderMatcher_PresentMatch{
				PresentMatch: true,
			}
		} else {

			envoyMatch.HeaderMatchSpecifier = &envoyroute.HeaderMatcher_ExactMatch{
				ExactMatch: matcher.Value,
			}
			if matcher.Regex {
				envoyMatch.HeaderMatchSpecifier = &envoyroute.HeaderMatcher_RegexMatch{
					RegexMatch: matcher.Value,
				}
			} else {
				envoyMatch.HeaderMatchSpecifier = &envoyroute.HeaderMatcher_ExactMatch{
					ExactMatch: matcher.Value,
				}
			}
		}
		out = append(out, envoyMatch)
	}
	return out
}

func envoyQueryMatcher(in []*v1.QueryParameterMatcher) []*envoyroute.QueryParameterMatcher {
	var out []*envoyroute.QueryParameterMatcher
	for _, matcher := range in {
		envoyMatch := &envoyroute.QueryParameterMatcher{
			Name:  matcher.Name,
			Value: matcher.Value,
			Regex: &types.BoolValue{
				Value: matcher.Regex,
			},
		}
		out = append(out, envoyMatch)
	}
	return out
}

// returns an error if any of the virtualhost domains overlap
func validateVirtualHostDomains(virtualHosts []*v1.VirtualHost) error {
	// this shouldbe a 1-1 mapping
	// if len(domainsToVirtualHosts[domain]) > 1, it's an error
	domainsToVirtualHosts := make(map[string][]string)
	for _, vHost := range virtualHosts {
		if len(vHost.Domains) == 0 {
			// default virtualhost
			domainsToVirtualHosts["*"] = append(domainsToVirtualHosts["*"], vHost.Name)
		}
		for _, domain := range vHost.Domains {
			// default virtualhost can be specified with empty string
			if domain == "" {
				domain = "*"
			}
			domainsToVirtualHosts[domain] = append(domainsToVirtualHosts[domain], vHost.Name)
		}
	}
	var domainErrors error
	// see if we found any conflicts, if so, write reports
	for domain, vHosts := range domainsToVirtualHosts {
		if len(vHosts) > 1 {
			domainErrors = multierror.Append(domainErrors, errors.Errorf("domain %v is "+
				"shared by the following virtual hosts: %v", domain, vHosts))
		}
	}
	return domainErrors
}

func ValidateRouteDestinations(snap *v1.ApiSnapshot, action *v1.RouteAction) error {
	upstreams := snap.Upstreams
	// make sure the destination itself has the right structure
	switch dest := action.Destination.(type) {
	case *v1.RouteAction_Single:
		return validateSingleDestination(upstreams, dest.Single)
	case *v1.RouteAction_Multi:
		return validateMultiDestination(upstreams, dest.Multi.Destinations)
	case *v1.RouteAction_UpstreamGroup:
		return validateUpstreamGroup(snap, dest.UpstreamGroup)
	}
	return errors.Errorf("must specify either 'singleDestination', 'multipleDestinations' or 'upstreamGroup' for action")
}

func validateUpstreamGroup(snap *v1.ApiSnapshot, ref *core.ResourceRef) error {

	upstreamGroup, err := snap.UpstreamGroups.Find(ref.Namespace, ref.Name)
	if err != nil {
		return errors.Wrap(err, "invalid destination for upstream group")
	}
	upstreams := snap.Upstreams

	err = validateMultiDestination(upstreams, upstreamGroup.Destinations)
	if err != nil {
		return errors.Wrap(err, "invalid destination in weighted destination in upstream group")
	}
	return nil
}

func validateMultiDestination(upstreams []*v1.Upstream, destinations []*v1.WeightedDestination) error {
	for _, dest := range destinations {
		if err := validateSingleDestination(upstreams, dest.Destination); err != nil {
			return errors.Wrap(err, "invalid destination in weighted destination list")
		}
	}
	return nil
}

func validateSingleDestination(upstreams v1.UpstreamList, destination *v1.Destination) error {
	upstreamRef, err := usconversion.DestinationToUpstreamRef(destination)
	if err != nil {
		return err
	}
	_, err = upstreams.Find(upstreamRef.Strings())
	return err
}

func validateListenerSslConfig(params plugins.Params, listener *v1.Listener) error {
	sslCfgTranslator := utils.NewSslConfigTranslator()
	for _, ssl := range listener.SslConfigurations {
		if _, err := sslCfgTranslator.ResolveDownstreamSslConfig(params.Snapshot.Secrets, ssl); err != nil {
			return err
		}
	}
	return nil
}

func DataSourceFromString(str string) *envoycore.DataSource {
	return &envoycore.DataSource{
		Specifier: &envoycore.DataSource_InlineString{
			InlineString: str,
		},
	}
}
