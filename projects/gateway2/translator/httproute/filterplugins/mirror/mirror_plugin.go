package mirror

import (
	"context"

	"github.com/pkg/errors"
	"github.com/solo-io/gloo/projects/gateway2/query"
	"github.com/solo-io/gloo/projects/gateway2/translator/plugins"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/shadowing"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
)

type plugin struct {
	queries query.GatewayQueries
}

func NewPlugin(queries query.GatewayQueries) *plugin {
	return &plugin{
		queries,
	}
}

func (p *plugin) ApplyFilter(
	ctx context.Context,
	routeCtx *plugins.RouteContext,
	filter gwv1.HTTPRouteFilter,
	outputRoute *v1.Route,
) error {
	config := filter.RequestMirror
	if config == nil {
		return errors.Errorf("RequestMirror filter supplied does not define requestMirror config")
	}

	routeAction := outputRoute.GetAction()
	if routeAction == nil {
		return errors.Errorf("RequestMirror must have destinations")
	}

	obj, err := p.queries.GetBackendForRef(ctx, p.queries.ObjToFrom(routeCtx.Route), &config.BackendRef)
	clusterName := query.ProcessBackendRef(
		obj,
		err,
		routeCtx.Reporter,
		config.BackendRef,
	)
	if clusterName == nil {
		return nil //TODO https://github.com/solo-io/gloo/pull/8890/files#r1391523183
	}

	outputRoute.Options.Shadowing = &shadowing.RouteShadowing{
		Upstream: &core.ResourceRef{
			Name:      *clusterName,
			Namespace: obj.GetNamespace(),
		},
		Percentage: 100.0,
	}

	return nil
}
