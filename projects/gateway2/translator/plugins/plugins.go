package plugins

import (
	"context"

	"github.com/solo-io/gloo/projects/gateway2/reports"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"

	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
)

// Plugin is an empty type for base plugins, currently no base methods.
type Plugin interface{}

type RouteContext struct {
	// top-level gw Listener
	Listener *gwv1.Listener
	// top-level HTTPRoute
	Route *gwv1.HTTPRoute
	// specific Rule of the HTTPRoute being processed
	Rule *gwv1.HTTPRouteRule
	// specific Match of the Rule being processed (as there may be multiple Matches per Rule)
	Match *gwv1.HTTPRouteMatch
	// Reporter for the correct ParentRef associated with this HTTPRoute
	Reporter reports.ParentRefReporter
}

type RoutePlugin interface {
	// ApplyRoutePlugin is called for each Match in a given Rule
	ApplyRoutePlugin(
		ctx context.Context,
		routeCtx *RouteContext,
		outputRoute *v1.Route,
	) error
}

type PostTranslationContext struct {
	// TranslatedProxies is the list of Proxies that were generated in a single translation run
	TranslatedProxies []TranslatedProxy
}

type TranslatedProxy struct {
	// Gateway is the input object that produced the Proxy
	Gateway gwv1.Gateway

	// Proxy is the output object, that was created by the Gateway
	Proxy *v1.Proxy
}

type PostTranslationPlugin interface {
	// ApplyPostTranslationPlugin is executed once at the end of a translation run
	ApplyPostTranslationPlugin(
		ctx context.Context,
		postTranslationContext *PostTranslationContext,
	) error
}
