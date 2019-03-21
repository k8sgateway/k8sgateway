package translator

import (
	"strings"

	envoyapi "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoycore "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	envoyroute "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	"github.com/gogo/protobuf/types"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	"github.com/solo-io/gloo/projects/gloo/pkg/utils"
	"github.com/solo-io/solo-kit/pkg/utils/contextutils"
)

type reportFunc func(error error, format string, args ...interface{})

func (t *translator) computeRouteConfig(params plugins.Params, proxy *v1.Proxy, listener *v1.Listener, routeCfgName string, reportFn reportFunc) *envoyapi.RouteConfiguration {
	report := func(err error, format string, args ...interface{}) {
		reportFn(err, "route_config."+format, args...)
	}
	params.Ctx = contextutils.WithLogger(params.Ctx, "compute_route_config."+routeCfgName)

	virtualHosts := t.computeVirtualHosts(params, listener, report)

	// validate ssl config if the listener specifies any
	if err := validateListenerSslConfig(listener, params.Snapshot.Secrets.List()); err != nil {
		report(err, "invalid listener %v", listener.Name)
	}

	return &envoyapi.RouteConfiguration{
		Name:         routeCfgName,
		VirtualHosts: virtualHosts,
	}
}

func (t *translator) computeVirtualHosts(params plugins.Params, listener *v1.Listener, report reportFunc) []envoyroute.VirtualHost {
	httpListener, ok := listener.ListenerType.(*v1.Listener_HttpListener)
	if !ok {
		panic("non-HTTP listeners are not currently supported in Gloo")
	}
	virtualHosts := httpListener.HttpListener.VirtualHosts
	if err := validateVirtualHostDomains(virtualHosts); err != nil {
		report(err, "invalid listener %v", listener.Name)
	}
	requireTls := len(listener.SslConfiguations) > 0
	var envoyVirtualHosts []envoyroute.VirtualHost
	for _, virtualHost := range virtualHosts {
		envoyVirtualHosts = append(envoyVirtualHosts, t.computeVirtualHost(params, virtualHost, requireTls, report))
	}
	return envoyVirtualHosts
}

func (t *translator) computeVirtualHost(params plugins.Params, virtualHost *v1.VirtualHost, requireTls bool, report reportFunc) envoyroute.VirtualHost {
	var envoyRoutes []envoyroute.Route
	for _, route := range virtualHost.Routes {
		envoyRoute := t.envoyRoute(params, report, route)
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
		// TODO (ilackarms): plugins for these
		// VirtualClusters: nil,
		// RateLimits: nil,
		// RequestHeadersToAdd: nil,
		// ResponseHeadersToRemove: nil,
		// Auth: nil,
	}

	// run the plugins
	for _, plug := range t.plugins {
		virtualHostPlugin, ok := plug.(plugins.VirtualHostPlugin)
		if !ok {
			continue
		}
		if err := virtualHostPlugin.ProcessVirtualHost(params, virtualHost, &out); err != nil {
			report(err, "invalid virtual host")
		}
	}
	return out
}

func (t *translator) envoyRoute(params plugins.Params, report reportFunc, in *v1.Route) envoyroute.Route {
	out := &envoyroute.Route{}

	setMatch(in, out)

	t.setAction(params, report, in, out)

	return *out
}

func setMatch(in *v1.Route, out *envoyroute.Route) {
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

func (t *translator) setAction(params plugins.Params, report reportFunc, in *v1.Route, out *envoyroute.Route) {
	switch action := in.Action.(type) {
	case *v1.Route_RouteAction:
		if err := validateRouteDestinations(params.Snapshot.Upstreams.List(), action.RouteAction); err != nil {
			report(err, "invalid route")
		}

		out.Action = &envoyroute.Route_Route{
			Route: &envoyroute.RouteAction{},
		}
		if err := setRouteAction(action.RouteAction, out.Action.(*envoyroute.Route_Route).Route); err != nil {
			report(err, "translator error on route")
		}

		// run the plugins for RoutePlugin
		for _, plug := range t.plugins {
			routePlugin, ok := plug.(plugins.RoutePlugin)
			if !ok {
				continue
			}
			if err := routePlugin.ProcessRoute(params, in, out); err != nil {
				report(err, "plugin error on route")
			}
		}
		// run the plugins for RoutePlugin
		for _, plug := range t.plugins {
			routePlugin, ok := plug.(plugins.RouteActionPlugin)
			if !ok || in.GetRouteAction() == nil || out.GetRoute() == nil {
				continue
			}
			if err := routePlugin.ProcessRouteAction(params, in.GetRouteAction(), nil, out.GetRoute()); err != nil {
				report(err, "plugin error on process route action")
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

func setRouteAction(in *v1.RouteAction, out *envoyroute.RouteAction) error {
	switch dest := in.Destination.(type) {
	case *v1.RouteAction_Single:
		out.ClusterSpecifier = &envoyroute.RouteAction_Cluster{
			Cluster: UpstreamToClusterName(dest.Single.Upstream),
		}
	case *v1.RouteAction_Multi:
		return setWeightedClusters(dest.Multi, out)
	}
	return nil
}

func setWeightedClusters(multiDest *v1.MultiDestination, out *envoyroute.RouteAction) error {
	if len(multiDest.Destinations) == 0 {
		return errors.Errorf("must specify at least one weighted destination for multi destination routes")
	}

	clusterSpecifier := &envoyroute.RouteAction_WeightedClusters{
		WeightedClusters: &envoyroute.WeightedCluster{},
	}

	var totalWeight uint32
	for _, weightedDest := range multiDest.Destinations {
		totalWeight += weightedDest.Weight
		clusterSpecifier.WeightedClusters.Clusters = append(clusterSpecifier.WeightedClusters.Clusters, &envoyroute.WeightedCluster_ClusterWeight{
			Name:   UpstreamToClusterName(weightedDest.Destination.Upstream),
			Weight: &types.UInt32Value{Value: weightedDest.Weight},
		})
	}

	clusterSpecifier.WeightedClusters.TotalWeight = &types.UInt32Value{Value: totalWeight}

	out.ClusterSpecifier = clusterSpecifier
	return nil
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
			HeaderMatchSpecifier: &envoyroute.HeaderMatcher_ExactMatch{
				ExactMatch: matcher.Value,
			},
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

func validateRouteDestinations(upstreams []*v1.Upstream, action *v1.RouteAction) error {
	// make sure the destination itself has the right structure
	switch dest := action.Destination.(type) {
	case *v1.RouteAction_Single:
		return validateSingleDestination(upstreams, dest.Single)
	case *v1.RouteAction_Multi:
		return validateMultiDestination(upstreams, dest.Multi.Destinations)
	}
	return errors.Errorf("must specify either 'single_destination' or 'multiple_destinations' for action")
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
	_, err := upstreams.Find(destination.Upstream.Strings())
	return err
}

func validateListenerSslConfig(listener *v1.Listener, secrets []*v1.Secret) error {
	sslCfgTranslator := utils.NewSslConfigTranslator(secrets)
	for _, ssl := range listener.SslConfiguations {
		if _, err := sslCfgTranslator.ResolveDownstreamSslConfig(ssl); err != nil {
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
