package hcm

import (
	envoyapi "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoycore "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	envoyhttp "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/http_connection_manager/v2"
	envoyutil "github.com/envoyproxy/go-control-plane/pkg/util"
	"github.com/pkg/errors"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/hcm"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	translatorutil "github.com/solo-io/gloo/projects/gloo/pkg/translator"
)

func NewPlugin(registryPlugins *[]plugins.Plugin) *Plugin {
	return &Plugin{
		allPlugins: registryPlugins,
	}
}

var _ plugins.Plugin = new(Plugin)
var _ plugins.ListenerPlugin = new(Plugin)

type Plugin struct {
	// assign this during construction
	allPlugins *[]plugins.Plugin
	// assemble this from the allPlugins list during Init to make sure that all plugins passed to the registry are
	// present (in case any get appended after our constructor is called)
	hcmPlugins []HcmPlugin
}

func (p *Plugin) Init(params plugins.InitParams) error {
	if p.allPlugins == nil {
		return initError()
	}
	// gather the plugins we care about
	for _, plugin := range *p.allPlugins {
		if hp, ok := plugin.(HcmPlugin); ok {
			p.hcmPlugins = append(p.hcmPlugins, hp)
		}
	}
	return nil
}

// ProcessListener has two responsibilities:
// 1. apply the core HCM settings from the HCM plugin to the listener
// 2. call each of the HCM plugins to make sure that they have a chance to apply their modifications to the listener
func (p *Plugin) ProcessListener(params plugins.Params, in *v1.Listener, out *envoyapi.Listener) error {
	hl, ok := in.ListenerType.(*v1.Listener_HttpListener)
	if !ok {
		return nil
	}
	if hl.HttpListener == nil {
		return nil
	}
	var hcmSettings *hcm.HttpConnectionManagerSettings
	if hl.HttpListener.GetListenerPlugins() != nil {
		hcmSettings = hl.HttpListener.GetListenerPlugins().HttpConnectionManagerSettings
	}
	if hcmSettings == nil && len(p.hcmPlugins) == 0 {
		// special case where we have nothing to do
		return nil
	}

	for _, f := range out.FilterChains {
		for i, filter := range f.Filters {
			if filter.Name == envoyutil.HTTPConnectionManager {
				// get config
				var cfg envoyhttp.HttpConnectionManager
				err := translatorutil.ParseConfig(&filter, &cfg)
				// this should never error
				if err != nil {
					return err
				}

				// first apply the core HCM settings, if any
				if hcmSettings != nil {
					copyCoreHcmSettings(&cfg, hcmSettings)
				}

				// then allow any HCM plugins to make their changes, with respect to any changes the core plugin made
				for _, hp := range p.hcmPlugins {
					if err := hp.ProcessHcmSettings(&cfg, hcmSettings); err != nil {
						return hcmPluginError(err)
					}
				}

				f.Filters[i], err = translatorutil.NewFilterWithConfig(envoyutil.HTTPConnectionManager, &cfg)
				// this should never error
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func copyCoreHcmSettings(cfg *envoyhttp.HttpConnectionManager, hcmSettings *hcm.HttpConnectionManagerSettings) {
	cfg.UseRemoteAddress = hcmSettings.UseRemoteAddress
	cfg.XffNumTrustedHops = hcmSettings.XffNumTrustedHops
	cfg.SkipXffAppend = hcmSettings.SkipXffAppend
	cfg.Via = hcmSettings.Via
	cfg.GenerateRequestId = hcmSettings.GenerateRequestId
	cfg.Proxy_100Continue = hcmSettings.Proxy_100Continue
	cfg.StreamIdleTimeout = hcmSettings.StreamIdleTimeout
	cfg.IdleTimeout = hcmSettings.IdleTimeout
	cfg.MaxRequestHeadersKb = hcmSettings.MaxRequestHeadersKb
	cfg.RequestTimeout = hcmSettings.RequestTimeout
	cfg.DrainTimeout = hcmSettings.DrainTimeout
	cfg.DelayedCloseTimeout = hcmSettings.DelayedCloseTimeout
	cfg.ServerName = hcmSettings.ServerName

	if hcmSettings.AcceptHttp_10 {
		cfg.HttpProtocolOptions = &envoycore.Http1ProtocolOptions{
			AcceptHttp_10:         true,
			DefaultHostForHttp_10: hcmSettings.DefaultHostForHttp_10,
		}
	}
}

var (
	initError = func() error {
		return errors.New("no plugins available at time of initialization")
	}
	hcmPluginError = func(err error) error {
		return errors.Wrapf(err, "error while running hcm plugin")
	}
)
