package tracing

import (
	envoyroute "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	envoy_config_trace_v3 "github.com/envoyproxy/go-control-plane/envoy/config/trace/v3"
	envoyhttp "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	envoytracing "github.com/envoyproxy/go-control-plane/envoy/type/tracing/v3"
	envoy_type "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/solo-io/gloo/pkg/utils/gogoutils"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/hcm"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/tracing"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	hcmp "github.com/solo-io/gloo/projects/gloo/pkg/plugins/hcm"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins/internal/common"
	translatorutil "github.com/solo-io/gloo/projects/gloo/pkg/translator"
)

// default all tracing percentages to 100%
const oneHundredPercent float32 = 100.0

func NewPlugin() *Plugin {
	return &Plugin{}
}

var _ plugins.Plugin = new(Plugin)
var _ hcmp.HcmPlugin = new(Plugin)
var _ plugins.RoutePlugin = new(Plugin)

type Plugin struct {
}

func (p *Plugin) Init(params plugins.InitParams) error {
	return nil
}

// Manage the tracing portion of the HCM settings
func (p *Plugin) ProcessHcmSettings(cfg *envoyhttp.HttpConnectionManager, hcmSettings *hcm.HttpConnectionManagerSettings) error {

	// only apply tracing config to the listener is using the HCM plugin
	if hcmSettings == nil {
		return nil
	}

	tracingSettings := hcmSettings.Tracing
	if tracingSettings == nil {
		return nil
	}

	// this plugin will overwrite any prior tracing config
	trCfg := &envoyhttp.HttpConnectionManager_Tracing{}

	var customTags []*envoytracing.CustomTag
	for _, h := range tracingSettings.RequestHeadersForTags {
		tag := &envoytracing.CustomTag{
			Tag: h,
			Type: &envoytracing.CustomTag_RequestHeader{
				RequestHeader: &envoytracing.CustomTag_Header{
					Name: h,
				},
			},
		}
		customTags = append(customTags, tag)
	}
	trCfg.CustomTags = customTags
	trCfg.Verbose = tracingSettings.Verbose
	trCfg.Provider = envoyTracingProvider(tracingSettings)

	// Gloo configures envoy as an ingress, rather than an egress
	// 06/2020 removing below- OperationName field is being deprecated, and we set it to the default value anyway
	// trCfg.OperationName = envoyhttp.HttpConnectionManager_Tracing_INGRESS
	if percentages := tracingSettings.GetTracePercentages(); percentages != nil {
		trCfg.ClientSampling = envoySimplePercentWithDefault(percentages.GetClientSamplePercentage(), oneHundredPercent)
		trCfg.RandomSampling = envoySimplePercentWithDefault(percentages.GetRandomSamplePercentage(), oneHundredPercent)
		trCfg.OverallSampling = envoySimplePercentWithDefault(percentages.GetOverallSamplePercentage(), oneHundredPercent)
	} else {
		trCfg.ClientSampling = envoySimplePercent(oneHundredPercent)
		trCfg.RandomSampling = envoySimplePercent(oneHundredPercent)
		trCfg.OverallSampling = envoySimplePercent(oneHundredPercent)
	}
	cfg.Tracing = trCfg
	return nil
}

func envoyTracingProvider(tracingSettings *tracing.ListenerTracingSettings) *envoy_config_trace_v3.Tracing_Http {
	if tracingSettings.Provider == nil {
		return nil
	}

	us := tracingSettings.Provider.UpstreamRef
	clusterName := translatorutil.UpstreamToClusterName(*us)

	// Todo - How to verify that this is a static upstream
	// Todo - Under what circumstances should this error? I assume if we do, we want to do it loudly

	switch typed := tracingSettings.Provider.GetTypedConfig().(type) {
	case *tracing.Provider_ZipkinConfig:
		converted, err := gogoutils.ToEnvoyZipkinTracingProvider(typed.ZipkinConfig, clusterName)
		if err != nil {
			return nil
		}

		marshalled, err := proto.Marshal(converted)
		if err != nil {
			return nil
		}

		return &envoy_config_trace_v3.Tracing_Http{
			Name: "envoy.tracers.zipkin",
			ConfigType: &envoy_config_trace_v3.Tracing_Http_TypedConfig{
				TypedConfig: &any.Any{
					TypeUrl: "type.googleapis.com/envoy.config.trace.v3.ZipkinConfig",
					Value:   marshalled,
				},
			},
		}

	case *tracing.Provider_DatadogConfig:
		converted, err := gogoutils.ToEnvoyDatadogTracingProvider(typed.DatadogConfig, clusterName)
		if err != nil {
			return nil
		}

		marshalled, err := proto.Marshal(converted)
		if err != nil {
			return nil
		}

		return &envoy_config_trace_v3.Tracing_Http{
			Name: "envoy.tracers.datadog",
			ConfigType: &envoy_config_trace_v3.Tracing_Http_TypedConfig{
				TypedConfig: &any.Any{
					TypeUrl: "type.googleapis.com/envoy.config.trace.v3.DatadogConfig",
					Value:   marshalled,
				},
			},
		}
	}
	return nil
}

func envoySimplePercent(numerator float32) *envoy_type.Percent {
	return &envoy_type.Percent{Value: float64(numerator)}
}

// use FloatValue to detect when nil (avoids error-prone float comparisons)
func envoySimplePercentWithDefault(numerator *types.FloatValue, defaultValue float32) *envoy_type.Percent {
	if numerator == nil {
		return envoySimplePercent(defaultValue)
	}
	return envoySimplePercent(numerator.Value)
}

func (p *Plugin) ProcessRoute(params plugins.RouteParams, in *v1.Route, out *envoyroute.Route) error {
	if in.Options == nil || in.Options.Tracing == nil {
		return nil
	}
	if percentages := in.GetOptions().GetTracing().TracePercentages; percentages != nil {
		out.Tracing = &envoyroute.Tracing{
			ClientSampling:  common.ToEnvoyPercentageWithDefault(percentages.GetClientSamplePercentage(), oneHundredPercent),
			RandomSampling:  common.ToEnvoyPercentageWithDefault(percentages.GetRandomSamplePercentage(), oneHundredPercent),
			OverallSampling: common.ToEnvoyPercentageWithDefault(percentages.GetOverallSamplePercentage(), oneHundredPercent),
		}
	} else {
		out.Tracing = &envoyroute.Tracing{
			ClientSampling:  common.ToEnvoyv2Percentage(oneHundredPercent),
			RandomSampling:  common.ToEnvoyv2Percentage(oneHundredPercent),
			OverallSampling: common.ToEnvoyv2Percentage(oneHundredPercent),
		}
	}
	descriptor := in.Options.Tracing.RouteDescriptor
	if descriptor != "" {
		out.Decorator = &envoyroute.Decorator{
			Operation: descriptor,
		}
	}
	return nil
}
