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

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"

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

	}

	return target
}

// Clone function
func (m *HttpGateway) Clone() proto.Message {
	var target *HttpGateway
	if m == nil {
		return target
	}
	target = &HttpGateway{}

	if m.GetVirtualServices() != nil {
		target.VirtualServices = make([]*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef, len(m.GetVirtualServices()))
		for idx, v := range m.GetVirtualServices() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.VirtualServices[idx] = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
			} else {
				target.VirtualServices[idx] = proto.Clone(v).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
			}

		}
	}

	if m.GetVirtualServiceSelector() != nil {
		target.VirtualServiceSelector = make(map[string]string, len(m.GetVirtualServiceSelector()))
		for k, v := range m.GetVirtualServiceSelector() {

			target.VirtualServiceSelector[k] = v

		}
	}

	if h, ok := interface{}(m.GetVirtualServiceExpressions()).(clone.Cloner); ok {
		target.VirtualServiceExpressions = h.Clone().(*VirtualServiceSelectorExpressions)
	} else {
		target.VirtualServiceExpressions = proto.Clone(m.GetVirtualServiceExpressions()).(*VirtualServiceSelectorExpressions)
	}

	if m.GetVirtualServiceNamespaces() != nil {
		target.VirtualServiceNamespaces = make([]string, len(m.GetVirtualServiceNamespaces()))
		for idx, v := range m.GetVirtualServiceNamespaces() {

			target.VirtualServiceNamespaces[idx] = v

		}
	}

	if h, ok := interface{}(m.GetOptions()).(clone.Cloner); ok {
		target.Options = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1.HttpListenerOptions)
	} else {
		target.Options = proto.Clone(m.GetOptions()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1.HttpListenerOptions)
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
func (m *VirtualServiceSelectorExpressions) Clone() proto.Message {
	var target *VirtualServiceSelectorExpressions
	if m == nil {
		return target
	}
	target = &VirtualServiceSelectorExpressions{}

	if m.GetExpressions() != nil {
		target.Expressions = make([]*VirtualServiceSelectorExpressions_Expression, len(m.GetExpressions()))
		for idx, v := range m.GetExpressions() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Expressions[idx] = h.Clone().(*VirtualServiceSelectorExpressions_Expression)
			} else {
				target.Expressions[idx] = proto.Clone(v).(*VirtualServiceSelectorExpressions_Expression)
			}

		}
	}

	return target
}

// Clone function
func (m *VirtualServiceSelectorExpressions_Expression) Clone() proto.Message {
	var target *VirtualServiceSelectorExpressions_Expression
	if m == nil {
		return target
	}
	target = &VirtualServiceSelectorExpressions_Expression{}

	target.Key = m.GetKey()

	target.Operator = m.GetOperator()

	if m.GetValues() != nil {
		target.Values = make([]string, len(m.GetValues()))
		for idx, v := range m.GetValues() {

			target.Values[idx] = v

		}
	}

	return target
}
