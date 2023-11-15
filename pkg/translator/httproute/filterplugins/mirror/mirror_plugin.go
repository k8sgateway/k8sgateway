package mirror

import (
	routev3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"github.com/pkg/errors"
	"github.com/solo-io/gloo/v2/pkg/query"
	"github.com/solo-io/gloo/v2/pkg/translator/httproute/filterplugins"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
)

type Plugin struct{}

func NewPlugin() *Plugin {
	return &Plugin{}
}

func (p *Plugin) ApplyFilter(
	ctx *filterplugins.RouteContext,
	filter gwv1.HTTPRouteFilter,
	outputRoute *routev3.Route,
) error {
	config := filter.RequestMirror
	if config == nil {
		return errors.Errorf("RequestMirror filter supplied does not define requestMirror config")
	}

	routeAction := outputRoute.GetRoute()

	if routeAction == nil {
		return errors.Errorf("RequestMirror must have destinations")
	}

	cli, err := ctx.Queries.GetBackendForRef(ctx.Ctx, ctx.Queries.ObjToFrom(ctx.Route), &config.BackendRef)
	clusterName := query.ProcessBackendRef(
		cli,
		err,
		ctx.Reporter,
		config.BackendRef,
	)
	if clusterName == nil {
		return nil
	}

	routeAction.RequestMirrorPolicies = append(routeAction.GetRequestMirrorPolicies(), &routev3.RouteAction_RequestMirrorPolicy{
		Cluster: *clusterName,
	})

	return nil
}
