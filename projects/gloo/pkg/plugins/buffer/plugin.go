package buffer

import (
	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	envoybuffer "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/buffer/v3"
	"github.com/envoyproxy/go-control-plane/pkg/wellknown"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/rotisserie/eris"
	buffer "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/filters/http/buffer/v3"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins/pluginutils"
)

// filter should be called after routing decision has been made
var pluginStage = plugins.DuringStage(plugins.RouteStage)

func NewPlugin() *Plugin {
	return &Plugin{}
}

var _ plugins.Plugin = new(Plugin)
var _ plugins.HttpFilterPlugin = new(Plugin)
var _ plugins.RoutePlugin = new(Plugin)
var _ plugins.VirtualHostPlugin = new(Plugin)
var _ plugins.WeightedDestinationPlugin = new(Plugin)

type Plugin struct {
	present bool
}

func (p *Plugin) Init(params plugins.InitParams) error {
	return nil
}

func (p *Plugin) HttpFilters(_ plugins.Params, listener *v1.HttpListener) ([]plugins.StagedHttpFilter, error) {

	bufferConfig, err := p.translateBufferFilter(listener.GetOptions().GetBuffer())
	if err != nil {
		return nil, err
	}

	if bufferConfig == nil {
		if !p.present {
			return nil, nil
		}
		// put the filter in the chain, actual buffer will be configured on route, vhost, etc.
		bufferConfig = &envoybuffer.Buffer{
			MaxRequestBytes: &wrappers.UInt32Value{
				Value: 1,
			},
		}

	}

	// put the filter in the chain, actual buffer will be configured on route, vhost, etc.
	bufferFilter, err := plugins.NewStagedFilterWithConfig(wellknown.Buffer, bufferConfig, pluginStage)
	if err != nil {
		return nil, eris.Wrapf(err, "generating filter config")
	}

	return []plugins.StagedHttpFilter{bufferFilter}, nil
}

func (p *Plugin) translateBufferFilter(buf *buffer.Buffer) (*envoybuffer.Buffer, error) {
	if buf == nil {
		return nil, nil
	}

	envoyConfig := &envoybuffer.Buffer{
		MaxRequestBytes: buf.GetMaxRequestBytes(),
	}

	return envoyConfig, envoyConfig.Validate()
}

func (p *Plugin) ProcessRoute(params plugins.RouteParams, in *v1.Route, out *envoy_config_route_v3.Route) error {
	bufPerRoute := in.Options.GetBufferPerRoute()
	if bufPerRoute == nil {
		return nil
	}

	if bufPerRoute.GetDisabled() {
		p.present = true
		return pluginutils.SetRoutePerFilterConfig(out, wellknown.Buffer, getNoBufferConfig())
	}

	if bufPerRoute.GetBuffer() != nil {
		config, err := getBufferConfig(bufPerRoute)
		if err != nil {
			return err
		}
		p.present = true
		return pluginutils.SetRoutePerFilterConfig(out, wellknown.Buffer, config)
	}

	return nil
}

func (p *Plugin) ProcessVirtualHost(
	params plugins.VirtualHostParams,
	in *v1.VirtualHost,
	out *envoy_config_route_v3.VirtualHost,
) error {
	bufPerRoute := in.GetOptions().GetBufferPerRoute()
	if bufPerRoute == nil {
		return nil
	}

	if bufPerRoute.GetDisabled() {
		p.present = true
		return pluginutils.SetVhostPerFilterConfig(out, wellknown.Buffer, getNoBufferConfig())
	}

	if bufPerRoute.GetBuffer() != nil {
		config, err := getBufferConfig(bufPerRoute)
		if err != nil {
			return err
		}
		p.present = true
		return pluginutils.SetVhostPerFilterConfig(out, wellknown.Buffer, config)
	}

	return nil
}

func (p *Plugin) ProcessWeightedDestination(
	params plugins.RouteParams,
	in *v1.WeightedDestination,
	out *envoy_config_route_v3.WeightedCluster_ClusterWeight,
) error {
	bufPerRoute := in.GetOptions().GetBufferPerRoute()
	if bufPerRoute == nil {
		return nil
	}

	if bufPerRoute.GetDisabled() {
		p.present = true
		return pluginutils.SetWeightedClusterPerFilterConfig(out, wellknown.Buffer, getNoBufferConfig())
	}

	if bufPerRoute.GetBuffer() != nil {
		config, err := getBufferConfig(bufPerRoute)
		if err != nil {
			return err
		}
		p.present = true
		return pluginutils.SetWeightedClusterPerFilterConfig(out, wellknown.Buffer, config)
	}

	return nil
}

func getNoBufferConfig() *envoybuffer.BufferPerRoute {
	return &envoybuffer.BufferPerRoute{
		Override: &envoybuffer.BufferPerRoute_Disabled{
			Disabled: true,
		},
	}
}

func getBufferConfig(bufPerRoute *buffer.BufferPerRoute) (*envoybuffer.BufferPerRoute, error) {
	envoyConfig := &envoybuffer.BufferPerRoute{
		Override: &envoybuffer.BufferPerRoute_Buffer{
			Buffer: &envoybuffer.Buffer{
				MaxRequestBytes: bufPerRoute.GetBuffer().GetMaxRequestBytes(),
			},
		},
	}
	return envoyConfig, envoyConfig.Validate()
}
