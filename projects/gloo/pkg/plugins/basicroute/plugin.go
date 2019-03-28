package basicroute

import (
	envoyroute "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	"github.com/gogo/protobuf/types"

	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	retries "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/retries"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	"github.com/solo-io/solo-kit/pkg/errors"
)

type Plugin struct{}

var _ plugins.RoutePlugin = NewPlugin()
var _ plugins.VirtualHostPlugin = NewPlugin()

// Handles a RoutePlugin APIs which map directly to basic Envoy config
func NewPlugin() *Plugin {
	return &Plugin{}
}

func (p *Plugin) Init(params plugins.InitParams) error {
	return nil
}

func (p *Plugin) ProcessVirtualHost(params plugins.Params, in *v1.VirtualHost, out *envoyroute.VirtualHost) error {
	if in.VirtualHostPlugins == nil {
		return nil
	}
	return applyRetriesVhost(in, out)
}

func (p *Plugin) ProcessRoute(params plugins.Params, in *v1.Route, out *envoyroute.Route) error {
	if in.RoutePlugins == nil {
		return nil
	}
	if err := applyPrefixRewrite(in, out); err != nil {
		return err
	}
	if err := applyTimeout(in, out); err != nil {
		return err
	}
	if err := applyRetries(in, out); err != nil {
		return err
	}

	return nil
}

func applyPrefixRewrite(in *v1.Route, out *envoyroute.Route) error {
	if in.RoutePlugins.PrefixRewrite == nil {
		return nil
	}
	routeAction, ok := out.Action.(*envoyroute.Route_Route)
	if !ok {
		return errors.Errorf("prefix rewrite is only available for Route Actions")
	}
	if routeAction.Route == nil {
		return errors.Errorf("internal error: route %v specified a prefix, but output Envoy object "+
			"had nil route", in.Action)
	}
	routeAction.Route.PrefixRewrite = in.RoutePlugins.PrefixRewrite.PrefixRewrite
	return nil
}

func applyTimeout(in *v1.Route, out *envoyroute.Route) error {
	if in.RoutePlugins.Timeout == nil {
		return nil
	}
	routeAction, ok := out.Action.(*envoyroute.Route_Route)
	if !ok {
		return errors.Errorf("timeout is only available for Route Actions")
	}
	if routeAction.Route == nil {
		return errors.Errorf("internal error: route %v specified a prefix, but output Envoy object "+
			"had nil route", in.Action)
	}

	routeAction.Route.Timeout = in.RoutePlugins.Timeout
	return nil
}

func applyRetries(in *v1.Route, out *envoyroute.Route) error {
	policy := in.RoutePlugins.Retries
	if policy == nil {
		return nil
	}
	routeAction, ok := out.Action.(*envoyroute.Route_Route)
	if !ok {
		return errors.Errorf("retries is only available for Route Actions")
	}
	if routeAction.Route == nil {
		return errors.Errorf("internal error: route %v specified a prefix, but output Envoy object "+
			"had nil route", in.Action)
	}

	routeAction.Route.RetryPolicy = convertPolicy(policy)
	return nil
}

func applyRetriesVhost(in *v1.VirtualHost, out *envoyroute.VirtualHost) error {
	out.RetryPolicy = convertPolicy(in.VirtualHostPlugins.Retries)
	return nil
}

func convertPolicy(policy *retries.RetryPolicy) *envoyroute.RetryPolicy {
	if policy == nil {
		return nil
	}

	numRetries := policy.NumRetries
	if numRetries == 0 {
		numRetries = 1
	}

	return &envoyroute.RetryPolicy{
		RetryOn:       policy.RetryOn,
		NumRetries:    &types.UInt32Value{Value: numRetries},
		PerTryTimeout: policy.PerTryTimeout,
	}
}
