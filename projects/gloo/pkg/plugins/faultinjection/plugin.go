package faultinjection

import (
	envoyroute "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	v2 "github.com/envoyproxy/go-control-plane/envoy/config/filter/fault/v2"
	envoyfault "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/fault/v2"
	envoytype "github.com/envoyproxy/go-control-plane/envoy/type"
	"github.com/gogo/protobuf/proto"
	fault "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/faultinjection"

	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins/pluginutils"
)

const (
	FilterName  = "envoy.fault"
	pluginStage = plugins.FaultFilter
)

type Plugin struct {
}

func NewPlugin() *Plugin {
	return &Plugin{}
}

func (p *Plugin) Init(params plugins.InitParams) error {
	return nil
}

func (p *Plugin) HttpFilters(params plugins.Params, listener *v1.HttpListener) ([]plugins.StagedHttpFilter, error) {
	// put the filter in the chain, but the actual faults will be configured on the routes
	return []plugins.StagedHttpFilter{
		plugins.NewStagedFilter(FilterName, pluginStage),
	}, nil
}

func (p *Plugin) ProcessRoute(params plugins.Params, in *v1.Route, out *envoyroute.Route) error {
	markFilterConfigFunc := func(spec *v1.Destination) (proto.Message, error) {
		if in.RoutePlugins == nil {
			return nil, nil
		}
		routeFaults := in.GetRoutePlugins().GetFaults()
		if routeFaults == nil {
			return nil, nil
		}
		routeAbort := routeFaults.GetAbort()
		routeDelay := routeFaults.GetDelay()
		if routeAbort == nil && routeDelay == nil {
			return nil, nil
		}
		return generateEnvoyConfigForHttpFault(routeAbort, routeDelay), nil
	}
	return pluginutils.MarkPerFilterConfig(params.Ctx, params.Snapshot, in, out, FilterName, markFilterConfigFunc)
}

func toEnvoyAbort(abort *fault.RouteAbort) *envoyfault.FaultAbort {
	if abort == nil {
		return nil
	}
	percentage := toEnvoyPercentage(abort.Percentage)
	errorType := &envoyfault.FaultAbort_HttpStatus{
		HttpStatus: uint32(abort.HttpStatus),
	}
	return &envoyfault.FaultAbort{
		Percentage: percentage,
		ErrorType:  errorType,
	}
}

func toEnvoyPercentage(percentage float32) *envoytype.FractionalPercent {
	return &envoytype.FractionalPercent{
		Numerator:   uint32(percentage * 10000),
		Denominator: envoytype.FractionalPercent_MILLION,
	}
}

func toEnvoyDelay(delay *fault.RouteDelay) *v2.FaultDelay {
	if delay == nil {
		return nil
	}
	percentage := toEnvoyPercentage(delay.Percentage)
	delaySpec := &v2.FaultDelay_FixedDelay{
		FixedDelay: delay.FixedDelay,
	}
	return &v2.FaultDelay{
		Percentage:         percentage,
		FaultDelaySecifier: delaySpec,
	}
}

func generateEnvoyConfigForHttpFault(routeAbort *fault.RouteAbort, routeDelay *fault.RouteDelay) *envoyfault.HTTPFault {
	abort := toEnvoyAbort(routeAbort)
	delay := toEnvoyDelay(routeDelay)
	return &envoyfault.HTTPFault{
		Abort: abort,
		Delay: delay,
	}
}
