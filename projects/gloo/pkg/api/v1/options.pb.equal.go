// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *ListenerOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ListenerOptions)
	if !ok {
		that2, ok := that.(ListenerOptions)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetAccessLoggingService()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAccessLoggingService()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAccessLoggingService(), target.GetAccessLoggingService()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetExtensions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetExtensions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetExtensions(), target.GetExtensions()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetPerConnectionBufferLimitBytes()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPerConnectionBufferLimitBytes()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPerConnectionBufferLimitBytes(), target.GetPerConnectionBufferLimitBytes()) {
			return false
		}
	}

	if len(m.GetSocketOptions()) != len(target.GetSocketOptions()) {
		return false
	}
	for idx, v := range m.GetSocketOptions() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSocketOptions()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSocketOptions()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetProxyProtocol()).(equality.Equalizer); ok {
		if !h.Equal(target.GetProxyProtocol()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetProxyProtocol(), target.GetProxyProtocol()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *RouteConfigurationOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RouteConfigurationOptions)
	if !ok {
		that2, ok := that.(RouteConfigurationOptions)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetMaxDirectResponseBodySizeBytes()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMaxDirectResponseBodySizeBytes()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMaxDirectResponseBodySizeBytes(), target.GetMaxDirectResponseBodySizeBytes()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *HttpListenerOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HttpListenerOptions)
	if !ok {
		that2, ok := that.(HttpListenerOptions)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetGrpcWeb()).(equality.Equalizer); ok {
		if !h.Equal(target.GetGrpcWeb()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetGrpcWeb(), target.GetGrpcWeb()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetHttpConnectionManagerSettings()).(equality.Equalizer); ok {
		if !h.Equal(target.GetHttpConnectionManagerSettings()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetHttpConnectionManagerSettings(), target.GetHttpConnectionManagerSettings()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetHealthCheck()).(equality.Equalizer); ok {
		if !h.Equal(target.GetHealthCheck()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetHealthCheck(), target.GetHealthCheck()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetExtensions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetExtensions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetExtensions(), target.GetExtensions()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetWaf()).(equality.Equalizer); ok {
		if !h.Equal(target.GetWaf()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetWaf(), target.GetWaf()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetDlp()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDlp()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDlp(), target.GetDlp()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetWasm()).(equality.Equalizer); ok {
		if !h.Equal(target.GetWasm()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetWasm(), target.GetWasm()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetExtauth()).(equality.Equalizer); ok {
		if !h.Equal(target.GetExtauth()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetExtauth(), target.GetExtauth()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetRatelimitServer()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRatelimitServer()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRatelimitServer(), target.GetRatelimitServer()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetGzip()).(equality.Equalizer); ok {
		if !h.Equal(target.GetGzip()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetGzip(), target.GetGzip()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetProxyLatency()).(equality.Equalizer); ok {
		if !h.Equal(target.GetProxyLatency()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetProxyLatency(), target.GetProxyLatency()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetBuffer()).(equality.Equalizer); ok {
		if !h.Equal(target.GetBuffer()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetBuffer(), target.GetBuffer()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetCsrf()).(equality.Equalizer); ok {
		if !h.Equal(target.GetCsrf()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetCsrf(), target.GetCsrf()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetGrpcJsonTranscoder()).(equality.Equalizer); ok {
		if !h.Equal(target.GetGrpcJsonTranscoder()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetGrpcJsonTranscoder(), target.GetGrpcJsonTranscoder()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetSanitizeClusterHeader()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSanitizeClusterHeader()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSanitizeClusterHeader(), target.GetSanitizeClusterHeader()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetLeftmostXffAddress()).(equality.Equalizer); ok {
		if !h.Equal(target.GetLeftmostXffAddress()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetLeftmostXffAddress(), target.GetLeftmostXffAddress()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetDynamicForwardProxy()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDynamicForwardProxy()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDynamicForwardProxy(), target.GetDynamicForwardProxy()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *TcpListenerOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TcpListenerOptions)
	if !ok {
		that2, ok := that.(TcpListenerOptions)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetTcpProxySettings()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTcpProxySettings()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTcpProxySettings(), target.GetTcpProxySettings()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *VirtualHostOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualHostOptions)
	if !ok {
		that2, ok := that.(VirtualHostOptions)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetExtensions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetExtensions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetExtensions(), target.GetExtensions()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetRetries()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRetries()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRetries(), target.GetRetries()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetStats()).(equality.Equalizer); ok {
		if !h.Equal(target.GetStats()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetStats(), target.GetStats()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetHeaderManipulation()).(equality.Equalizer); ok {
		if !h.Equal(target.GetHeaderManipulation()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetHeaderManipulation(), target.GetHeaderManipulation()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetCors()).(equality.Equalizer); ok {
		if !h.Equal(target.GetCors()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetCors(), target.GetCors()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetTransformations()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTransformations()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTransformations(), target.GetTransformations()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetRatelimitBasic()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRatelimitBasic()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRatelimitBasic(), target.GetRatelimitBasic()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetWaf()).(equality.Equalizer); ok {
		if !h.Equal(target.GetWaf()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetWaf(), target.GetWaf()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetRbac()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRbac()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRbac(), target.GetRbac()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetExtauth()).(equality.Equalizer); ok {
		if !h.Equal(target.GetExtauth()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetExtauth(), target.GetExtauth()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetDlp()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDlp()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDlp(), target.GetDlp()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetBufferPerRoute()).(equality.Equalizer); ok {
		if !h.Equal(target.GetBufferPerRoute()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetBufferPerRoute(), target.GetBufferPerRoute()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetCsrf()).(equality.Equalizer); ok {
		if !h.Equal(target.GetCsrf()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetCsrf(), target.GetCsrf()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetIncludeRequestAttemptCount()).(equality.Equalizer); ok {
		if !h.Equal(target.GetIncludeRequestAttemptCount()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetIncludeRequestAttemptCount(), target.GetIncludeRequestAttemptCount()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetIncludeAttemptCountInResponse()).(equality.Equalizer); ok {
		if !h.Equal(target.GetIncludeAttemptCountInResponse()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetIncludeAttemptCountInResponse(), target.GetIncludeAttemptCountInResponse()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetStagedTransformations()).(equality.Equalizer); ok {
		if !h.Equal(target.GetStagedTransformations()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetStagedTransformations(), target.GetStagedTransformations()) {
			return false
		}
	}

	switch m.RateLimitConfigType.(type) {

	case *VirtualHostOptions_Ratelimit:
		if _, ok := target.RateLimitConfigType.(*VirtualHostOptions_Ratelimit); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRatelimit()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRatelimit()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRatelimit(), target.GetRatelimit()) {
				return false
			}
		}

	case *VirtualHostOptions_RateLimitConfigs:
		if _, ok := target.RateLimitConfigType.(*VirtualHostOptions_RateLimitConfigs); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRateLimitConfigs()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRateLimitConfigs()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRateLimitConfigs(), target.GetRateLimitConfigs()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.RateLimitConfigType != target.RateLimitConfigType {
			return false
		}
	}

	switch m.JwtConfig.(type) {

	case *VirtualHostOptions_Jwt:
		if _, ok := target.JwtConfig.(*VirtualHostOptions_Jwt); !ok {
			return false
		}

		if h, ok := interface{}(m.GetJwt()).(equality.Equalizer); ok {
			if !h.Equal(target.GetJwt()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetJwt(), target.GetJwt()) {
				return false
			}
		}

	case *VirtualHostOptions_JwtStaged:
		if _, ok := target.JwtConfig.(*VirtualHostOptions_JwtStaged); !ok {
			return false
		}

		if h, ok := interface{}(m.GetJwtStaged()).(equality.Equalizer); ok {
			if !h.Equal(target.GetJwtStaged()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetJwtStaged(), target.GetJwtStaged()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.JwtConfig != target.JwtConfig {
			return false
		}
	}

	return true
}

// Equal function
func (m *RouteOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RouteOptions)
	if !ok {
		that2, ok := that.(RouteOptions)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetTransformations()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTransformations()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTransformations(), target.GetTransformations()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetFaults()).(equality.Equalizer); ok {
		if !h.Equal(target.GetFaults()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetFaults(), target.GetFaults()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetPrefixRewrite()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPrefixRewrite()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPrefixRewrite(), target.GetPrefixRewrite()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetTimeout()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTimeout()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTimeout(), target.GetTimeout()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetRetries()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRetries()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRetries(), target.GetRetries()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetExtensions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetExtensions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetExtensions(), target.GetExtensions()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetTracing()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTracing()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTracing(), target.GetTracing()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetShadowing()).(equality.Equalizer); ok {
		if !h.Equal(target.GetShadowing()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetShadowing(), target.GetShadowing()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetHeaderManipulation()).(equality.Equalizer); ok {
		if !h.Equal(target.GetHeaderManipulation()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetHeaderManipulation(), target.GetHeaderManipulation()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetCors()).(equality.Equalizer); ok {
		if !h.Equal(target.GetCors()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetCors(), target.GetCors()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetLbHash()).(equality.Equalizer); ok {
		if !h.Equal(target.GetLbHash()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetLbHash(), target.GetLbHash()) {
			return false
		}
	}

	if len(m.GetUpgrades()) != len(target.GetUpgrades()) {
		return false
	}
	for idx, v := range m.GetUpgrades() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetUpgrades()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetUpgrades()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetRatelimitBasic()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRatelimitBasic()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRatelimitBasic(), target.GetRatelimitBasic()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetWaf()).(equality.Equalizer); ok {
		if !h.Equal(target.GetWaf()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetWaf(), target.GetWaf()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetRbac()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRbac()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRbac(), target.GetRbac()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetExtauth()).(equality.Equalizer); ok {
		if !h.Equal(target.GetExtauth()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetExtauth(), target.GetExtauth()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetDlp()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDlp()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDlp(), target.GetDlp()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetBufferPerRoute()).(equality.Equalizer); ok {
		if !h.Equal(target.GetBufferPerRoute()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetBufferPerRoute(), target.GetBufferPerRoute()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetCsrf()).(equality.Equalizer); ok {
		if !h.Equal(target.GetCsrf()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetCsrf(), target.GetCsrf()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetStagedTransformations()).(equality.Equalizer); ok {
		if !h.Equal(target.GetStagedTransformations()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetStagedTransformations(), target.GetStagedTransformations()) {
			return false
		}
	}

	if len(m.GetEnvoyMetadata()) != len(target.GetEnvoyMetadata()) {
		return false
	}
	for k, v := range m.GetEnvoyMetadata() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetEnvoyMetadata()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetEnvoyMetadata()[k]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetRegexRewrite()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRegexRewrite()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRegexRewrite(), target.GetRegexRewrite()) {
			return false
		}
	}

	switch m.HostRewriteType.(type) {

	case *RouteOptions_HostRewrite:
		if _, ok := target.HostRewriteType.(*RouteOptions_HostRewrite); !ok {
			return false
		}

		if strings.Compare(m.GetHostRewrite(), target.GetHostRewrite()) != 0 {
			return false
		}

	case *RouteOptions_AutoHostRewrite:
		if _, ok := target.HostRewriteType.(*RouteOptions_AutoHostRewrite); !ok {
			return false
		}

		if h, ok := interface{}(m.GetAutoHostRewrite()).(equality.Equalizer); ok {
			if !h.Equal(target.GetAutoHostRewrite()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetAutoHostRewrite(), target.GetAutoHostRewrite()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.HostRewriteType != target.HostRewriteType {
			return false
		}
	}

	switch m.RateLimitConfigType.(type) {

	case *RouteOptions_Ratelimit:
		if _, ok := target.RateLimitConfigType.(*RouteOptions_Ratelimit); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRatelimit()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRatelimit()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRatelimit(), target.GetRatelimit()) {
				return false
			}
		}

	case *RouteOptions_RateLimitConfigs:
		if _, ok := target.RateLimitConfigType.(*RouteOptions_RateLimitConfigs); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRateLimitConfigs()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRateLimitConfigs()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRateLimitConfigs(), target.GetRateLimitConfigs()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.RateLimitConfigType != target.RateLimitConfigType {
			return false
		}
	}

	switch m.JwtConfig.(type) {

	case *RouteOptions_Jwt:
		if _, ok := target.JwtConfig.(*RouteOptions_Jwt); !ok {
			return false
		}

		if h, ok := interface{}(m.GetJwt()).(equality.Equalizer); ok {
			if !h.Equal(target.GetJwt()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetJwt(), target.GetJwt()) {
				return false
			}
		}

	case *RouteOptions_JwtStaged:
		if _, ok := target.JwtConfig.(*RouteOptions_JwtStaged); !ok {
			return false
		}

		if h, ok := interface{}(m.GetJwtStaged()).(equality.Equalizer); ok {
			if !h.Equal(target.GetJwtStaged()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetJwtStaged(), target.GetJwtStaged()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.JwtConfig != target.JwtConfig {
			return false
		}
	}

	return true
}

// Equal function
func (m *DestinationSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*DestinationSpec)
	if !ok {
		that2, ok := that.(DestinationSpec)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	switch m.DestinationType.(type) {

	case *DestinationSpec_Aws:
		if _, ok := target.DestinationType.(*DestinationSpec_Aws); !ok {
			return false
		}

		if h, ok := interface{}(m.GetAws()).(equality.Equalizer); ok {
			if !h.Equal(target.GetAws()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetAws(), target.GetAws()) {
				return false
			}
		}

	case *DestinationSpec_Azure:
		if _, ok := target.DestinationType.(*DestinationSpec_Azure); !ok {
			return false
		}

		if h, ok := interface{}(m.GetAzure()).(equality.Equalizer); ok {
			if !h.Equal(target.GetAzure()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetAzure(), target.GetAzure()) {
				return false
			}
		}

	case *DestinationSpec_Rest:
		if _, ok := target.DestinationType.(*DestinationSpec_Rest); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRest()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRest()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRest(), target.GetRest()) {
				return false
			}
		}

	case *DestinationSpec_Grpc:
		if _, ok := target.DestinationType.(*DestinationSpec_Grpc); !ok {
			return false
		}

		if h, ok := interface{}(m.GetGrpc()).(equality.Equalizer); ok {
			if !h.Equal(target.GetGrpc()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetGrpc(), target.GetGrpc()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.DestinationType != target.DestinationType {
			return false
		}
	}

	return true
}

// Equal function
func (m *WeightedDestinationOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WeightedDestinationOptions)
	if !ok {
		that2, ok := that.(WeightedDestinationOptions)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetHeaderManipulation()).(equality.Equalizer); ok {
		if !h.Equal(target.GetHeaderManipulation()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetHeaderManipulation(), target.GetHeaderManipulation()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetTransformations()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTransformations()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTransformations(), target.GetTransformations()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetExtensions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetExtensions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetExtensions(), target.GetExtensions()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetExtauth()).(equality.Equalizer); ok {
		if !h.Equal(target.GetExtauth()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetExtauth(), target.GetExtauth()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetBufferPerRoute()).(equality.Equalizer); ok {
		if !h.Equal(target.GetBufferPerRoute()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetBufferPerRoute(), target.GetBufferPerRoute()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetCsrf()).(equality.Equalizer); ok {
		if !h.Equal(target.GetCsrf()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetCsrf(), target.GetCsrf()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetStagedTransformations()).(equality.Equalizer); ok {
		if !h.Equal(target.GetStagedTransformations()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetStagedTransformations(), target.GetStagedTransformations()) {
			return false
		}
	}

	return true
}
