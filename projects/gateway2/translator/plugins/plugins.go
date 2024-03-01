package plugins

import (
	"context"

	"github.com/solo-io/gloo/projects/gateway2/reports"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"

	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
)

// Empty type for base plugins, currently no base methods.
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
	// called for each Match in a given Rule
	ApplyRoutePlugin(
		ctx context.Context,
		routeCtx *RouteContext,
		outputRoute *v1.Route,
	) error
}

type NamespaceContext struct {
	// this is the namespace where the gateway lives
	Namespace string
}

type PostTranslationPlugin interface {
	// called for each Namespace containing a gateway
	ApplyPostTranslationPlugin(
		ctx context.Context,
		namespaceCtx *NamespaceContext,
	) error
}
