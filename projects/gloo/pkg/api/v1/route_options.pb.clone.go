// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/route_options.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_filters_http_buffer_v3 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/filters/http/buffer/v3"

	github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_filters_http_csrf_v3 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/filters/http/csrf/v3"

	github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_type_matcher_v3 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/type/matcher/v3"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ai "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_dlp "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/dlp"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_extauth_v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/extauth/v1"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_extproc "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/extproc"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_jwt "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/jwt"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ratelimit"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_rbac "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/rbac"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_waf "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/waf"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_cors "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/cors"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_faultinjection "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/faultinjection"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_headers "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/headers"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_lbhash "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/lbhash"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_protocol_upgrade "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/protocol_upgrade"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_retries "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/retries"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_shadowing "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/shadowing"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_tracing "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/tracing"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_transformation "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/transformation"

	google_golang_org_protobuf_types_known_durationpb "google.golang.org/protobuf/types/known/durationpb"

	google_golang_org_protobuf_types_known_structpb "google.golang.org/protobuf/types/known/structpb"

	google_golang_org_protobuf_types_known_wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = clone.Cloner(nil)
	_ = proto.Message(nil)
)

// Clone function
func (m *RouteOptions) Clone() proto.Message {
	var target *RouteOptions
	if m == nil {
		return target
	}
	target = &RouteOptions{}

	if h, ok := interface{}(m.GetTransformations()).(clone.Cloner); ok {
		target.Transformations = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_transformation.Transformations)
	} else {
		target.Transformations = proto.Clone(m.GetTransformations()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_transformation.Transformations)
	}

	if h, ok := interface{}(m.GetFaults()).(clone.Cloner); ok {
		target.Faults = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_faultinjection.RouteFaults)
	} else {
		target.Faults = proto.Clone(m.GetFaults()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_faultinjection.RouteFaults)
	}

	if h, ok := interface{}(m.GetPrefixRewrite()).(clone.Cloner); ok {
		target.PrefixRewrite = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	} else {
		target.PrefixRewrite = proto.Clone(m.GetPrefixRewrite()).(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	}

	if h, ok := interface{}(m.GetTimeout()).(clone.Cloner); ok {
		target.Timeout = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.Timeout = proto.Clone(m.GetTimeout()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	if h, ok := interface{}(m.GetRetries()).(clone.Cloner); ok {
		target.Retries = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_retries.RetryPolicy)
	} else {
		target.Retries = proto.Clone(m.GetRetries()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_retries.RetryPolicy)
	}

	if h, ok := interface{}(m.GetExtensions()).(clone.Cloner); ok {
		target.Extensions = h.Clone().(*Extensions)
	} else {
		target.Extensions = proto.Clone(m.GetExtensions()).(*Extensions)
	}

	if h, ok := interface{}(m.GetTracing()).(clone.Cloner); ok {
		target.Tracing = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_tracing.RouteTracingSettings)
	} else {
		target.Tracing = proto.Clone(m.GetTracing()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_tracing.RouteTracingSettings)
	}

	if h, ok := interface{}(m.GetShadowing()).(clone.Cloner); ok {
		target.Shadowing = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_shadowing.RouteShadowing)
	} else {
		target.Shadowing = proto.Clone(m.GetShadowing()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_shadowing.RouteShadowing)
	}

	if h, ok := interface{}(m.GetHeaderManipulation()).(clone.Cloner); ok {
		target.HeaderManipulation = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_headers.HeaderManipulation)
	} else {
		target.HeaderManipulation = proto.Clone(m.GetHeaderManipulation()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_headers.HeaderManipulation)
	}

	if h, ok := interface{}(m.GetAppendXForwardedHost()).(clone.Cloner); ok {
		target.AppendXForwardedHost = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.AppendXForwardedHost = proto.Clone(m.GetAppendXForwardedHost()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	if h, ok := interface{}(m.GetCors()).(clone.Cloner); ok {
		target.Cors = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_cors.CorsPolicy)
	} else {
		target.Cors = proto.Clone(m.GetCors()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_cors.CorsPolicy)
	}

	if h, ok := interface{}(m.GetLbHash()).(clone.Cloner); ok {
		target.LbHash = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_lbhash.RouteActionHashConfig)
	} else {
		target.LbHash = proto.Clone(m.GetLbHash()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_lbhash.RouteActionHashConfig)
	}

	if m.GetUpgrades() != nil {
		target.Upgrades = make([]*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_protocol_upgrade.ProtocolUpgradeConfig, len(m.GetUpgrades()))
		for idx, v := range m.GetUpgrades() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Upgrades[idx] = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_protocol_upgrade.ProtocolUpgradeConfig)
			} else {
				target.Upgrades[idx] = proto.Clone(v).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_protocol_upgrade.ProtocolUpgradeConfig)
			}

		}
	}

	if h, ok := interface{}(m.GetRatelimitBasic()).(clone.Cloner); ok {
		target.RatelimitBasic = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.IngressRateLimit)
	} else {
		target.RatelimitBasic = proto.Clone(m.GetRatelimitBasic()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.IngressRateLimit)
	}

	if h, ok := interface{}(m.GetWaf()).(clone.Cloner); ok {
		target.Waf = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_waf.Settings)
	} else {
		target.Waf = proto.Clone(m.GetWaf()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_waf.Settings)
	}

	if h, ok := interface{}(m.GetRbac()).(clone.Cloner); ok {
		target.Rbac = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_rbac.ExtensionSettings)
	} else {
		target.Rbac = proto.Clone(m.GetRbac()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_rbac.ExtensionSettings)
	}

	if h, ok := interface{}(m.GetExtauth()).(clone.Cloner); ok {
		target.Extauth = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_extauth_v1.ExtAuthExtension)
	} else {
		target.Extauth = proto.Clone(m.GetExtauth()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_extauth_v1.ExtAuthExtension)
	}

	if h, ok := interface{}(m.GetDlp()).(clone.Cloner); ok {
		target.Dlp = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_dlp.Config)
	} else {
		target.Dlp = proto.Clone(m.GetDlp()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_dlp.Config)
	}

	if h, ok := interface{}(m.GetBufferPerRoute()).(clone.Cloner); ok {
		target.BufferPerRoute = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_filters_http_buffer_v3.BufferPerRoute)
	} else {
		target.BufferPerRoute = proto.Clone(m.GetBufferPerRoute()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_filters_http_buffer_v3.BufferPerRoute)
	}

	if h, ok := interface{}(m.GetCsrf()).(clone.Cloner); ok {
		target.Csrf = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_filters_http_csrf_v3.CsrfPolicy)
	} else {
		target.Csrf = proto.Clone(m.GetCsrf()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_filters_http_csrf_v3.CsrfPolicy)
	}

	if h, ok := interface{}(m.GetStagedTransformations()).(clone.Cloner); ok {
		target.StagedTransformations = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_transformation.TransformationStages)
	} else {
		target.StagedTransformations = proto.Clone(m.GetStagedTransformations()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_transformation.TransformationStages)
	}

	if m.GetEnvoyMetadata() != nil {
		target.EnvoyMetadata = make(map[string]*google_golang_org_protobuf_types_known_structpb.Struct, len(m.GetEnvoyMetadata()))
		for k, v := range m.GetEnvoyMetadata() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.EnvoyMetadata[k] = h.Clone().(*google_golang_org_protobuf_types_known_structpb.Struct)
			} else {
				target.EnvoyMetadata[k] = proto.Clone(v).(*google_golang_org_protobuf_types_known_structpb.Struct)
			}

		}
	}

	if h, ok := interface{}(m.GetRegexRewrite()).(clone.Cloner); ok {
		target.RegexRewrite = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_type_matcher_v3.RegexMatchAndSubstitute)
	} else {
		target.RegexRewrite = proto.Clone(m.GetRegexRewrite()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_type_matcher_v3.RegexMatchAndSubstitute)
	}

	if h, ok := interface{}(m.GetMaxStreamDuration()).(clone.Cloner); ok {
		target.MaxStreamDuration = h.Clone().(*RouteOptions_MaxStreamDuration)
	} else {
		target.MaxStreamDuration = proto.Clone(m.GetMaxStreamDuration()).(*RouteOptions_MaxStreamDuration)
	}

	if h, ok := interface{}(m.GetIdleTimeout()).(clone.Cloner); ok {
		target.IdleTimeout = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.IdleTimeout = proto.Clone(m.GetIdleTimeout()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	if h, ok := interface{}(m.GetExtProc()).(clone.Cloner); ok {
		target.ExtProc = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_extproc.RouteSettings)
	} else {
		target.ExtProc = proto.Clone(m.GetExtProc()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_extproc.RouteSettings)
	}

	if h, ok := interface{}(m.GetAi()).(clone.Cloner); ok {
		target.Ai = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ai.RouteSettings)
	} else {
		target.Ai = proto.Clone(m.GetAi()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ai.RouteSettings)
	}

	switch m.HostRewriteType.(type) {

	case *RouteOptions_HostRewrite:

		target.HostRewriteType = &RouteOptions_HostRewrite{
			HostRewrite: m.GetHostRewrite(),
		}

	case *RouteOptions_AutoHostRewrite:

		if h, ok := interface{}(m.GetAutoHostRewrite()).(clone.Cloner); ok {
			target.HostRewriteType = &RouteOptions_AutoHostRewrite{
				AutoHostRewrite: h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue),
			}
		} else {
			target.HostRewriteType = &RouteOptions_AutoHostRewrite{
				AutoHostRewrite: proto.Clone(m.GetAutoHostRewrite()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue),
			}
		}

	case *RouteOptions_HostRewritePathRegex:

		if h, ok := interface{}(m.GetHostRewritePathRegex()).(clone.Cloner); ok {
			target.HostRewriteType = &RouteOptions_HostRewritePathRegex{
				HostRewritePathRegex: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_type_matcher_v3.RegexMatchAndSubstitute),
			}
		} else {
			target.HostRewriteType = &RouteOptions_HostRewritePathRegex{
				HostRewritePathRegex: proto.Clone(m.GetHostRewritePathRegex()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_type_matcher_v3.RegexMatchAndSubstitute),
			}
		}

	case *RouteOptions_HostRewriteHeader:

		if h, ok := interface{}(m.GetHostRewriteHeader()).(clone.Cloner); ok {
			target.HostRewriteType = &RouteOptions_HostRewriteHeader{
				HostRewriteHeader: h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.StringValue),
			}
		} else {
			target.HostRewriteType = &RouteOptions_HostRewriteHeader{
				HostRewriteHeader: proto.Clone(m.GetHostRewriteHeader()).(*google_golang_org_protobuf_types_known_wrapperspb.StringValue),
			}
		}

	}

	switch m.RateLimitEarlyConfigType.(type) {

	case *RouteOptions_RatelimitEarly:

		if h, ok := interface{}(m.GetRatelimitEarly()).(clone.Cloner); ok {
			target.RateLimitEarlyConfigType = &RouteOptions_RatelimitEarly{
				RatelimitEarly: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.RateLimitRouteExtension),
			}
		} else {
			target.RateLimitEarlyConfigType = &RouteOptions_RatelimitEarly{
				RatelimitEarly: proto.Clone(m.GetRatelimitEarly()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.RateLimitRouteExtension),
			}
		}

	case *RouteOptions_RateLimitEarlyConfigs:

		if h, ok := interface{}(m.GetRateLimitEarlyConfigs()).(clone.Cloner); ok {
			target.RateLimitEarlyConfigType = &RouteOptions_RateLimitEarlyConfigs{
				RateLimitEarlyConfigs: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.RateLimitConfigRefs),
			}
		} else {
			target.RateLimitEarlyConfigType = &RouteOptions_RateLimitEarlyConfigs{
				RateLimitEarlyConfigs: proto.Clone(m.GetRateLimitEarlyConfigs()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.RateLimitConfigRefs),
			}
		}

	}

	switch m.RateLimitConfigType.(type) {

	case *RouteOptions_Ratelimit:

		if h, ok := interface{}(m.GetRatelimit()).(clone.Cloner); ok {
			target.RateLimitConfigType = &RouteOptions_Ratelimit{
				Ratelimit: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.RateLimitRouteExtension),
			}
		} else {
			target.RateLimitConfigType = &RouteOptions_Ratelimit{
				Ratelimit: proto.Clone(m.GetRatelimit()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.RateLimitRouteExtension),
			}
		}

	case *RouteOptions_RateLimitConfigs:

		if h, ok := interface{}(m.GetRateLimitConfigs()).(clone.Cloner); ok {
			target.RateLimitConfigType = &RouteOptions_RateLimitConfigs{
				RateLimitConfigs: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.RateLimitConfigRefs),
			}
		} else {
			target.RateLimitConfigType = &RouteOptions_RateLimitConfigs{
				RateLimitConfigs: proto.Clone(m.GetRateLimitConfigs()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.RateLimitConfigRefs),
			}
		}

	}

	switch m.RateLimitRegularConfigType.(type) {

	case *RouteOptions_RatelimitRegular:

		if h, ok := interface{}(m.GetRatelimitRegular()).(clone.Cloner); ok {
			target.RateLimitRegularConfigType = &RouteOptions_RatelimitRegular{
				RatelimitRegular: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.RateLimitRouteExtension),
			}
		} else {
			target.RateLimitRegularConfigType = &RouteOptions_RatelimitRegular{
				RatelimitRegular: proto.Clone(m.GetRatelimitRegular()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.RateLimitRouteExtension),
			}
		}

	case *RouteOptions_RateLimitRegularConfigs:

		if h, ok := interface{}(m.GetRateLimitRegularConfigs()).(clone.Cloner); ok {
			target.RateLimitRegularConfigType = &RouteOptions_RateLimitRegularConfigs{
				RateLimitRegularConfigs: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.RateLimitConfigRefs),
			}
		} else {
			target.RateLimitRegularConfigType = &RouteOptions_RateLimitRegularConfigs{
				RateLimitRegularConfigs: proto.Clone(m.GetRateLimitRegularConfigs()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.RateLimitConfigRefs),
			}
		}

	}

	switch m.JwtConfig.(type) {

	case *RouteOptions_Jwt:

		if h, ok := interface{}(m.GetJwt()).(clone.Cloner); ok {
			target.JwtConfig = &RouteOptions_Jwt{
				Jwt: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_jwt.RouteExtension),
			}
		} else {
			target.JwtConfig = &RouteOptions_Jwt{
				Jwt: proto.Clone(m.GetJwt()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_jwt.RouteExtension),
			}
		}

	case *RouteOptions_JwtStaged:

		if h, ok := interface{}(m.GetJwtStaged()).(clone.Cloner); ok {
			target.JwtConfig = &RouteOptions_JwtStaged{
				JwtStaged: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_jwt.JwtStagedRouteExtension),
			}
		} else {
			target.JwtConfig = &RouteOptions_JwtStaged{
				JwtStaged: proto.Clone(m.GetJwtStaged()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_jwt.JwtStagedRouteExtension),
			}
		}

	case *RouteOptions_JwtProvidersStaged:

		if h, ok := interface{}(m.GetJwtProvidersStaged()).(clone.Cloner); ok {
			target.JwtConfig = &RouteOptions_JwtProvidersStaged{
				JwtProvidersStaged: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_jwt.JwtStagedRouteProvidersExtension),
			}
		} else {
			target.JwtConfig = &RouteOptions_JwtProvidersStaged{
				JwtProvidersStaged: proto.Clone(m.GetJwtProvidersStaged()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_jwt.JwtStagedRouteProvidersExtension),
			}
		}

	}

	return target
}

// Clone function
func (m *RouteOptions_MaxStreamDuration) Clone() proto.Message {
	var target *RouteOptions_MaxStreamDuration
	if m == nil {
		return target
	}
	target = &RouteOptions_MaxStreamDuration{}

	if h, ok := interface{}(m.GetMaxStreamDuration()).(clone.Cloner); ok {
		target.MaxStreamDuration = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.MaxStreamDuration = proto.Clone(m.GetMaxStreamDuration()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	if h, ok := interface{}(m.GetGrpcTimeoutHeaderMax()).(clone.Cloner); ok {
		target.GrpcTimeoutHeaderMax = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.GrpcTimeoutHeaderMax = proto.Clone(m.GetGrpcTimeoutHeaderMax()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	if h, ok := interface{}(m.GetGrpcTimeoutHeaderOffset()).(clone.Cloner); ok {
		target.GrpcTimeoutHeaderOffset = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.GrpcTimeoutHeaderOffset = proto.Clone(m.GetGrpcTimeoutHeaderOffset()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	return target
}
