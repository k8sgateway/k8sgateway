// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gateway/api/v1/gateway.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_wrappers "github.com/golang/protobuf/ptypes/wrappers"

	github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_core_v3 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/config/core/v3"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_core_selectors "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/core/selectors"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_hcm "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/hcm"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_ssl "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/ssl"

	github_com_solo_io_solo_kit_pkg_api_v1_resources_core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
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
func (m *Gateway) Clone() proto.Message {
	var target *Gateway
	if m == nil {
		return target
	}
	target = &Gateway{}

	target.Ssl = m.GetSsl()

	target.BindAddress = m.GetBindAddress()

	target.BindPort = m.GetBindPort()

	if h, ok := interface{}(m.GetOptions()).(clone.Cloner); ok {
		target.Options = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1.ListenerOptions)
	} else {
		target.Options = proto.Clone(m.GetOptions()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1.ListenerOptions)
	}

	if h, ok := interface{}(m.GetNamespacedStatuses()).(clone.Cloner); ok {
		target.NamespacedStatuses = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.NamespacedStatuses)
	} else {
		target.NamespacedStatuses = proto.Clone(m.GetNamespacedStatuses()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.NamespacedStatuses)
	}

	if h, ok := interface{}(m.GetMetadata()).(clone.Cloner); ok {
		target.Metadata = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Metadata)
	} else {
		target.Metadata = proto.Clone(m.GetMetadata()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Metadata)
	}

	if h, ok := interface{}(m.GetUseProxyProto()).(clone.Cloner); ok {
		target.UseProxyProto = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.UseProxyProto = proto.Clone(m.GetUseProxyProto()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if m.GetProxyNames() != nil {
		target.ProxyNames = make([]string, len(m.GetProxyNames()))
		for idx, v := range m.GetProxyNames() {

			target.ProxyNames[idx] = v

		}
	}

	if h, ok := interface{}(m.GetRouteOptions()).(clone.Cloner); ok {
		target.RouteOptions = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1.RouteConfigurationOptions)
	} else {
		target.RouteOptions = proto.Clone(m.GetRouteOptions()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1.RouteConfigurationOptions)
	}

	switch m.GatewayType.(type) {

	case *Gateway_HttpGateway:

		if h, ok := interface{}(m.GetHttpGateway()).(clone.Cloner); ok {
			target.GatewayType = &Gateway_HttpGateway{
				HttpGateway: h.Clone().(*HttpGateway),
			}
		} else {
			target.GatewayType = &Gateway_HttpGateway{
				HttpGateway: proto.Clone(m.GetHttpGateway()).(*HttpGateway),
			}
		}

	case *Gateway_TcpGateway:

		if h, ok := interface{}(m.GetTcpGateway()).(clone.Cloner); ok {
			target.GatewayType = &Gateway_TcpGateway{
				TcpGateway: h.Clone().(*TcpGateway),
			}
		} else {
			target.GatewayType = &Gateway_TcpGateway{
				TcpGateway: proto.Clone(m.GetTcpGateway()).(*TcpGateway),
			}
		}

	case *Gateway_HybridGateway:

		if h, ok := interface{}(m.GetHybridGateway()).(clone.Cloner); ok {
			target.GatewayType = &Gateway_HybridGateway{
				HybridGateway: h.Clone().(*HybridGateway),
			}
		} else {
			target.GatewayType = &Gateway_HybridGateway{
				HybridGateway: proto.Clone(m.GetHybridGateway()).(*HybridGateway),
			}
		}

	}

	return target
}

// Clone function
func (m *TcpGateway) Clone() proto.Message {
	var target *TcpGateway
	if m == nil {
		return target
	}
	target = &TcpGateway{}

	if m.GetTcpHosts() != nil {
		target.TcpHosts = make([]*github_com_solo_io_gloo_projects_gloo_pkg_api_v1.TcpHost, len(m.GetTcpHosts()))
		for idx, v := range m.GetTcpHosts() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.TcpHosts[idx] = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1.TcpHost)
			} else {
				target.TcpHosts[idx] = proto.Clone(v).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1.TcpHost)
			}

		}
	}

	if h, ok := interface{}(m.GetOptions()).(clone.Cloner); ok {
		target.Options = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1.TcpListenerOptions)
	} else {
		target.Options = proto.Clone(m.GetOptions()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1.TcpListenerOptions)
	}

	return target
}

// Clone function
func (m *HybridGateway) Clone() proto.Message {
	var target *HybridGateway
	if m == nil {
		return target
	}
	target = &HybridGateway{}

	if m.GetMatchedGateways() != nil {
		target.MatchedGateways = make([]*MatchedGateway, len(m.GetMatchedGateways()))
		for idx, v := range m.GetMatchedGateways() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.MatchedGateways[idx] = h.Clone().(*MatchedGateway)
			} else {
				target.MatchedGateways[idx] = proto.Clone(v).(*MatchedGateway)
			}

		}
	}

	if h, ok := interface{}(m.GetDelegatedHttpGateways()).(clone.Cloner); ok {
		target.DelegatedHttpGateways = h.Clone().(*DelegatedHttpGateway)
	} else {
		target.DelegatedHttpGateways = proto.Clone(m.GetDelegatedHttpGateways()).(*DelegatedHttpGateway)
	}

	if h, ok := interface{}(m.GetDelegatedTcpGateways()).(clone.Cloner); ok {
		target.DelegatedTcpGateways = h.Clone().(*DelegatedTcpGateway)
	} else {
		target.DelegatedTcpGateways = proto.Clone(m.GetDelegatedTcpGateways()).(*DelegatedTcpGateway)
	}

	return target
}

// Clone function
func (m *DelegatedHttpGateway) Clone() proto.Message {
	var target *DelegatedHttpGateway
	if m == nil {
		return target
	}
	target = &DelegatedHttpGateway{}

	target.PreventChildOverrides = m.GetPreventChildOverrides()

	if h, ok := interface{}(m.GetHttpConnectionManagerSettings()).(clone.Cloner); ok {
		target.HttpConnectionManagerSettings = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_hcm.HttpConnectionManagerSettings)
	} else {
		target.HttpConnectionManagerSettings = proto.Clone(m.GetHttpConnectionManagerSettings()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_hcm.HttpConnectionManagerSettings)
	}

	if h, ok := interface{}(m.GetSslConfig()).(clone.Cloner); ok {
		target.SslConfig = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_ssl.SslConfig)
	} else {
		target.SslConfig = proto.Clone(m.GetSslConfig()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_ssl.SslConfig)
	}

	switch m.SelectionType.(type) {

	case *DelegatedHttpGateway_Ref:

		if h, ok := interface{}(m.GetRef()).(clone.Cloner); ok {
			target.SelectionType = &DelegatedHttpGateway_Ref{
				Ref: h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef),
			}
		} else {
			target.SelectionType = &DelegatedHttpGateway_Ref{
				Ref: proto.Clone(m.GetRef()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef),
			}
		}

	case *DelegatedHttpGateway_Selector:

		if h, ok := interface{}(m.GetSelector()).(clone.Cloner); ok {
			target.SelectionType = &DelegatedHttpGateway_Selector{
				Selector: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_core_selectors.Selector),
			}
		} else {
			target.SelectionType = &DelegatedHttpGateway_Selector{
				Selector: proto.Clone(m.GetSelector()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_core_selectors.Selector),
			}
		}

	}

	return target
}

// Clone function
func (m *DelegatedTcpGateway) Clone() proto.Message {
	var target *DelegatedTcpGateway
	if m == nil {
		return target
	}
	target = &DelegatedTcpGateway{}

	switch m.SelectionType.(type) {

	case *DelegatedTcpGateway_Ref:

		if h, ok := interface{}(m.GetRef()).(clone.Cloner); ok {
			target.SelectionType = &DelegatedTcpGateway_Ref{
				Ref: h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef),
			}
		} else {
			target.SelectionType = &DelegatedTcpGateway_Ref{
				Ref: proto.Clone(m.GetRef()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef),
			}
		}

	case *DelegatedTcpGateway_Selector:

		if h, ok := interface{}(m.GetSelector()).(clone.Cloner); ok {
			target.SelectionType = &DelegatedTcpGateway_Selector{
				Selector: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_core_selectors.Selector),
			}
		} else {
			target.SelectionType = &DelegatedTcpGateway_Selector{
				Selector: proto.Clone(m.GetSelector()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_core_selectors.Selector),
			}
		}

	}

	return target
}

// Clone function
func (m *MatchedGateway) Clone() proto.Message {
	var target *MatchedGateway
	if m == nil {
		return target
	}
	target = &MatchedGateway{}

	if h, ok := interface{}(m.GetMatcher()).(clone.Cloner); ok {
		target.Matcher = h.Clone().(*Matcher)
	} else {
		target.Matcher = proto.Clone(m.GetMatcher()).(*Matcher)
	}

	switch m.GatewayType.(type) {

	case *MatchedGateway_HttpGateway:

		if h, ok := interface{}(m.GetHttpGateway()).(clone.Cloner); ok {
			target.GatewayType = &MatchedGateway_HttpGateway{
				HttpGateway: h.Clone().(*HttpGateway),
			}
		} else {
			target.GatewayType = &MatchedGateway_HttpGateway{
				HttpGateway: proto.Clone(m.GetHttpGateway()).(*HttpGateway),
			}
		}

	case *MatchedGateway_TcpGateway:

		if h, ok := interface{}(m.GetTcpGateway()).(clone.Cloner); ok {
			target.GatewayType = &MatchedGateway_TcpGateway{
				TcpGateway: h.Clone().(*TcpGateway),
			}
		} else {
			target.GatewayType = &MatchedGateway_TcpGateway{
				TcpGateway: proto.Clone(m.GetTcpGateway()).(*TcpGateway),
			}
		}

	}

	return target
}

// Clone function
func (m *Matcher) Clone() proto.Message {
	var target *Matcher
	if m == nil {
		return target
	}
	target = &Matcher{}

	if h, ok := interface{}(m.GetSslConfig()).(clone.Cloner); ok {
		target.SslConfig = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_ssl.SslConfig)
	} else {
		target.SslConfig = proto.Clone(m.GetSslConfig()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_ssl.SslConfig)
	}

	if m.GetSourcePrefixRanges() != nil {
		target.SourcePrefixRanges = make([]*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_core_v3.CidrRange, len(m.GetSourcePrefixRanges()))
		for idx, v := range m.GetSourcePrefixRanges() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.SourcePrefixRanges[idx] = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_core_v3.CidrRange)
			} else {
				target.SourcePrefixRanges[idx] = proto.Clone(v).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_core_v3.CidrRange)
			}

		}
	}

	if m.GetPrefixRanges() != nil {
		target.PrefixRanges = make([]*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_core_v3.CidrRange, len(m.GetPrefixRanges()))
		for idx, v := range m.GetPrefixRanges() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.PrefixRanges[idx] = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_core_v3.CidrRange)
			} else {
				target.PrefixRanges[idx] = proto.Clone(v).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_core_v3.CidrRange)
			}

		}
	}

	if h, ok := interface{}(m.GetDestinationPort()).(clone.Cloner); ok {
		target.DestinationPort = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.DestinationPort = proto.Clone(m.GetDestinationPort()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	if m.GetPassthroughCipherSuites() != nil {
		target.PassthroughCipherSuites = make([]string, len(m.GetPassthroughCipherSuites()))
		for idx, v := range m.GetPassthroughCipherSuites() {

			target.PassthroughCipherSuites[idx] = v

		}
	}

	return target
}
