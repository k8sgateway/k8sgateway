package httproute

import (
	"container/list"
	"context"

	"github.com/rotisserie/eris"
	"github.com/solo-io/gloo/projects/gateway2/ir"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"

	"github.com/solo-io/gloo/projects/gateway2/query"
	"github.com/solo-io/gloo/projects/gateway2/reports"
)

var (
	awsMissingFuncRefError                = eris.New("upstreams must have a logical name specified in the backend ref via the parameters extensionref")
	azureMissingFuncRefError              = eris.New("upstreams must have a function name specified in the backend ref via the parameters extensionref")
	nonFunctionUpstreamWithParameterError = eris.New("parameters extensionref is only supported for aws and azure upstreams")
)

func TranslateGatewayHTTPRouteRules(
	ctx context.Context,
	gwListener gwv1.Listener,
	routeInfo *query.RouteInfo,
	reporter reports.ParentRefReporter,
	baseReporter reports.Reporter,
) []ir.HttpRouteRuleMatchIR {
	var finalRoutes []ir.HttpRouteRuleMatchIR
	routesVisited := sets.New[types.NamespacedName]()

	// Only HTTPRoute types should be translated.
	route, ok := routeInfo.Object.(*ir.HttpRouteIR)
	if !ok {
		return finalRoutes
	}

	// Hostnames need to be explicitly passed to the plugins since they
	// are required by delegatee (child) routes of delegated routes that
	// won't have spec.Hostnames set.
	hostnames := route.Hostnames

	delegationChain := list.New()

	translateGatewayHTTPRouteRulesUtil(
		ctx, gwListener, routeInfo, reporter, baseReporter, &finalRoutes, routesVisited, hostnames, delegationChain)
	return finalRoutes
}

// translateGatewayHTTPRouteRulesUtil is a helper to translate an HTTPRoute.
// In case of route delegation, this function is recursively invoked to flatten the delegated route tree.
func translateGatewayHTTPRouteRulesUtil(
	ctx context.Context,
	gwListener gwv1.Listener,
	routeInfo *query.RouteInfo,
	reporter reports.ParentRefReporter,
	baseReporter reports.Reporter,
	outputs *[]ir.HttpRouteRuleMatchIR,
	routesVisited sets.Set[types.NamespacedName],
	hostnames []string,
	delegationChain *list.List,
) {
	// Only HTTPRoute types should be translated.
	route, ok := routeInfo.Object.(*ir.HttpRouteIR)
	if !ok {
		return
	}

	for ruleIdx, rule := range route.Rules {
		rule := rule
		if rule.Matches == nil {
			// from the spec:
			// If no matches are specified, the default is a prefix path match on “/”, which has the effect of matching every HTTP request.
			rule.Matches = []gwv1.HTTPRouteMatch{{}}
		}

		outputRoutes := translateGatewayHTTPRouteRule(
			ctx,
			gwListener,
			routeInfo,
			rule,
			ruleIdx,
			reporter,
			baseReporter,
			outputs,
			routesVisited,
			hostnames,
			delegationChain,
		)
		for _, outputRoute := range outputRoutes {
			// The above function will return a nil route if a matcher fails to apply plugins
			// properly. This is a signal to the caller that the route should be dropped.
			//		if outputRoute == nil {
			//			continue
			//		}

			*outputs = append(*outputs, outputRoute)
		}
	}
}

// MARK: translate rules
func translateGatewayHTTPRouteRule(
	ctx context.Context,
	gwListener gwv1.Listener,
	gwroute *query.RouteInfo,
	rule ir.HttpRouteRuleIR,
	ruleIdx int,
	reporter reports.ParentRefReporter,
	baseReporter reports.Reporter,
	outputs *[]ir.HttpRouteRuleMatchIR,
	routesVisited sets.Set[types.NamespacedName],
	hostnames []string,
	delegationChain *list.List,
) []ir.HttpRouteRuleMatchIR {
	routes := make([]ir.HttpRouteRuleMatchIR, 0, len(rule.Matches))

	for idx, match := range rule.Matches {
		match := match // pike
		// HTTPRoute names are being introduced to upstream as part of https://github.com/kubernetes-sigs/gateway-api/issues/995
		// For now, the HTTPRoute needs a unique name for each Route to support features that require the route name
		// set (basic ratelimit, route-level jwt, etc.). The unique name is generated by appending the index of the route to the
		// HTTPRoute name.namespace.
		uniqueRouteName := gwroute.UniqueRouteName(ruleIdx, idx)

		outputRoute := ir.HttpRouteRuleMatchIR{
			HttpRouteRuleCommonIR: rule.HttpRouteRuleCommonIR,
			ParentRef:             gwroute.ListenerParentRef,
			Name:                  uniqueRouteName,
			Backends:              nil,
			MatchIndex:            idx,
			Match:                 match,
		}

		var delegatedRoutes []ir.HttpRouteRuleMatchIR
		var delegates bool
		if len(rule.Backends) > 0 {
			delegates = setRouteAction(
				ctx,
				gwroute,
				rule,
				outputRoute,
				reporter,
				baseReporter,
				gwListener,
				match,
				&delegatedRoutes,
				routesVisited,
				delegationChain,
			)
		}
		/*
			plugins now happen in the IR translator, and policy attachment happens when the IR is
			constructed.. so this can probably be removed
					rtCtx := &plugins.RouteContext{
						Listener:        &gwListener,
						HTTPRoute:       route,
						Hostnames:       hostnames,
						DelegationChain: delegationChain,
						Rule:            &rule,
						Match:           &match,
						Reporter:        reporter,
					}

					// Apply the plugins for this route
					for _, plugin := range pluginRegistry.GetRoutePlugins() {
						err := plugin.ApplyRoutePlugin(ctx, rtCtx, outputRoute)
						if err != nil {
							contextutils.LoggerFrom(ctx).Errorf("error in RoutePlugin: %v", err)
						}

						// If this parent route has delegatee routes, override any applied policies
						// that are on the child with the parent's policies.
						// When a plugin is invoked on a route, it must override the existing route.
						for _, child := range delegatedRoutes {
							err := plugin.ApplyRoutePlugin(ctx, rtCtx, child)
							if err != nil {
								contextutils.LoggerFrom(ctx).Errorf("error applying RoutePlugin to child route %s: %v", child.GetName(), err)
							}
						}
					}
		*/

		// Add the delegatee output routes to the final output list
		*outputs = append(*outputs, delegatedRoutes...)

		// It is possible for a parent route to not produce an output route action
		// if it only delegates and does not directly route to a backend.
		// We should only set a direct response action when there is no output action
		// for a parent rule and when there are no delegated routes because this would
		// otherwise result in a top level matcher with a direct response action for the
		// path that the parent is delegating for.

		if len(outputRoute.Backends) == 0 && !delegates {
			/* i don't think we need this as round without backends will error regardless.
			outputRoute.Action = &v1.Route_DirectResponseAction{
				DirectResponseAction: &v1.DirectResponseAction{
					Status: http.StatusInternalServerError,
				},
			}
			*/
		}
		// A parent route that delegates to a child route should not have an output route
		// action (outputRoute.Action) as the routes are derived from the child route.
		// So this conditional ensures that we do not create a top level route matcher
		// for the parent route when it delegates to a child route.

		if len(outputRoute.Backends) > 0 {
			routes = append(routes, outputRoute)
		}
	}
	return routes
}

func setRouteAction(
	ctx context.Context,
	gwroute *query.RouteInfo,
	rule ir.HttpRouteRuleIR,
	outputRoute ir.HttpRouteRuleMatchIR,
	reporter reports.ParentRefReporter,
	baseReporter reports.Reporter,
	gwListener gwv1.Listener,
	match gwv1.HTTPRouteMatch,
	outputs *[]ir.HttpRouteRuleMatchIR,
	routesVisited sets.Set[types.NamespacedName],
	delegationChain *list.List,
) bool {
	backends := rule.Backends
	delegates := false

	for _, backend := range backends {
		// If the backend is an HTTPRoute, it implies route delegation
		// for which delegated routes are recursively flattened and translated
		if backend.Delegate != nil {
			delegates = true
			// Flatten delegated HTTPRoute references
			err := flattenDelegatedRoutes(
				ctx,
				gwroute,
				backend,
				reporter,
				baseReporter,
				gwListener,
				match,
				outputs,
				routesVisited,
				delegationChain,
			)
			if err != nil {
				query.ProcessBackendError(err, reporter)
			}
			continue
		}

		httpBackend := ir.HttpBackend{
			Backend:          *backend.Backend, // TODO: Nil check?
			AttachedPolicies: backend.AttachedPolicies,
		}
		outputRoute.Backends = append(outputRoute.Backends, httpBackend)
	}

	return delegates
}

/* TODO: demonstrate that we can replace this with 'virtual' GKs
func applyBackendPlugins(
	obj client.Object,
	backendRef gwv1.BackendObjectReference,
	plugins registry.PluginRegistry,
) (*v1.Destination, bool) {
	for _, bp := range plugins.GetBackendPlugins() {
		if dest, ok := bp.ApplyBackendPlugin(obj, backendRef); ok {
			return dest, true
		}
	}
	return nil, false
}
*/
