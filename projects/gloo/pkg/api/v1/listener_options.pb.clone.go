// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/listener_options.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_als "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/als"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_proxy_protocol "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/proxy_protocol"

	github_com_solo_io_solo_kit_pkg_api_external_envoy_api_v2_core "github.com/solo-io/solo-kit/pkg/api/external/envoy/api/v2/core"

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
func (m *ListenerOptions) Clone() proto.Message {
	var target *ListenerOptions
	if m == nil {
		return target
	}
	target = &ListenerOptions{}

	if h, ok := interface{}(m.GetAccessLoggingService()).(clone.Cloner); ok {
		target.AccessLoggingService = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_als.AccessLoggingService)
	} else {
		target.AccessLoggingService = proto.Clone(m.GetAccessLoggingService()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_als.AccessLoggingService)
	}

	if h, ok := interface{}(m.GetExtensions()).(clone.Cloner); ok {
		target.Extensions = h.Clone().(*Extensions)
	} else {
		target.Extensions = proto.Clone(m.GetExtensions()).(*Extensions)
	}

	if h, ok := interface{}(m.GetPerConnectionBufferLimitBytes()).(clone.Cloner); ok {
		target.PerConnectionBufferLimitBytes = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.PerConnectionBufferLimitBytes = proto.Clone(m.GetPerConnectionBufferLimitBytes()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	if m.GetSocketOptions() != nil {
		target.SocketOptions = make([]*github_com_solo_io_solo_kit_pkg_api_external_envoy_api_v2_core.SocketOption, len(m.GetSocketOptions()))
		for idx, v := range m.GetSocketOptions() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.SocketOptions[idx] = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_external_envoy_api_v2_core.SocketOption)
			} else {
				target.SocketOptions[idx] = proto.Clone(v).(*github_com_solo_io_solo_kit_pkg_api_external_envoy_api_v2_core.SocketOption)
			}

		}
	}

	if h, ok := interface{}(m.GetProxyProtocol()).(clone.Cloner); ok {
		target.ProxyProtocol = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_proxy_protocol.ProxyProtocol)
	} else {
		target.ProxyProtocol = proto.Clone(m.GetProxyProtocol()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_proxy_protocol.ProxyProtocol)
	}

	if h, ok := interface{}(m.GetConnectionBalanceConfig()).(clone.Cloner); ok {
		target.ConnectionBalanceConfig = h.Clone().(*ConnectionBalanceConfig)
	} else {
		target.ConnectionBalanceConfig = proto.Clone(m.GetConnectionBalanceConfig()).(*ConnectionBalanceConfig)
	}

	if h, ok := interface{}(m.GetListenerAccessLoggingService()).(clone.Cloner); ok {
		target.ListenerAccessLoggingService = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_als.AccessLoggingService)
	} else {
		target.ListenerAccessLoggingService = proto.Clone(m.GetListenerAccessLoggingService()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_als.AccessLoggingService)
	}

	if h, ok := interface{}(m.GetTcpStats()).(clone.Cloner); ok {
		target.TcpStats = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.TcpStats = proto.Clone(m.GetTcpStats()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	return target
}

// Clone function
func (m *ConnectionBalanceConfig) Clone() proto.Message {
	var target *ConnectionBalanceConfig
	if m == nil {
		return target
	}
	target = &ConnectionBalanceConfig{}

	if h, ok := interface{}(m.GetExactBalance()).(clone.Cloner); ok {
		target.ExactBalance = h.Clone().(*ConnectionBalanceConfig_ExactBalance)
	} else {
		target.ExactBalance = proto.Clone(m.GetExactBalance()).(*ConnectionBalanceConfig_ExactBalance)
	}

	return target
}

// Clone function
func (m *ConnectionBalanceConfig_ExactBalance) Clone() proto.Message {
	var target *ConnectionBalanceConfig_ExactBalance
	if m == nil {
		return target
	}
	target = &ConnectionBalanceConfig_ExactBalance{}

	return target
}
