// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gateway/api/v1/gateway.proto

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
func (m *Gateway) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Gateway)
	if !ok {
		that2, ok := that.(Gateway)
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

	if m.GetSsl() != target.GetSsl() {
		return false
	}

	if strings.Compare(m.GetBindAddress(), target.GetBindAddress()) != 0 {
		return false
	}

	if m.GetBindPort() != target.GetBindPort() {
		return false
	}

	if h, ok := interface{}(m.GetOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOptions(), target.GetOptions()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetStatus()).(equality.Equalizer); ok {
		if !h.Equal(target.GetStatus()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetStatus(), target.GetStatus()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetMetadata()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMetadata()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMetadata(), target.GetMetadata()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetUseProxyProto()).(equality.Equalizer); ok {
		if !h.Equal(target.GetUseProxyProto()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetUseProxyProto(), target.GetUseProxyProto()) {
			return false
		}
	}

	if len(m.GetProxyNames()) != len(target.GetProxyNames()) {
		return false
	}
	for idx, v := range m.GetProxyNames() {

		if strings.Compare(v, target.GetProxyNames()[idx]) != 0 {
			return false
		}

	}

	if h, ok := interface{}(m.GetRouteOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRouteOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRouteOptions(), target.GetRouteOptions()) {
			return false
		}
	}

	switch m.GatewayType.(type) {

	case *Gateway_HttpGateway:
		if _, ok := target.GatewayType.(*Gateway_HttpGateway); !ok {
			return false
		}

		if h, ok := interface{}(m.GetHttpGateway()).(equality.Equalizer); ok {
			if !h.Equal(target.GetHttpGateway()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetHttpGateway(), target.GetHttpGateway()) {
				return false
			}
		}

	case *Gateway_TcpGateway:
		if _, ok := target.GatewayType.(*Gateway_TcpGateway); !ok {
			return false
		}

		if h, ok := interface{}(m.GetTcpGateway()).(equality.Equalizer); ok {
			if !h.Equal(target.GetTcpGateway()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetTcpGateway(), target.GetTcpGateway()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.GatewayType != target.GatewayType {
			return false
		}
	}

	return true
}

// Equal function
func (m *HttpGateway) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HttpGateway)
	if !ok {
		that2, ok := that.(HttpGateway)
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

	if len(m.GetVirtualServices()) != len(target.GetVirtualServices()) {
		return false
	}
	for idx, v := range m.GetVirtualServices() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetVirtualServices()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetVirtualServices()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetVirtualServiceSelector()).(equality.Equalizer); ok {
		if !h.Equal(target.GetVirtualServiceSelector()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetVirtualServiceSelector(), target.GetVirtualServiceSelector()) {
			return false
		}
	}

	if len(m.GetVirtualServiceNamespaces()) != len(target.GetVirtualServiceNamespaces()) {
		return false
	}
	for idx, v := range m.GetVirtualServiceNamespaces() {

		if strings.Compare(v, target.GetVirtualServiceNamespaces()[idx]) != 0 {
			return false
		}

	}

	if h, ok := interface{}(m.GetOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOptions(), target.GetOptions()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *TcpGateway) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TcpGateway)
	if !ok {
		that2, ok := that.(TcpGateway)
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

	if len(m.GetTcpHosts()) != len(target.GetTcpHosts()) {
		return false
	}
	for idx, v := range m.GetTcpHosts() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetTcpHosts()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetTcpHosts()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOptions(), target.GetOptions()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *VirtualServiceSelector) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualServiceSelector)
	if !ok {
		that2, ok := that.(VirtualServiceSelector)
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

	if len(m.GetNamespaces()) != len(target.GetNamespaces()) {
		return false
	}
	for idx, v := range m.GetNamespaces() {

		if strings.Compare(v, target.GetNamespaces()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetLabels()) != len(target.GetLabels()) {
		return false
	}
	for k, v := range m.GetLabels() {

		if strings.Compare(v, target.GetLabels()[k]) != 0 {
			return false
		}

	}

	if len(m.GetExpressions()) != len(target.GetExpressions()) {
		return false
	}
	for idx, v := range m.GetExpressions() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetExpressions()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetExpressions()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *VirtualServiceSelector_Expression) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualServiceSelector_Expression)
	if !ok {
		that2, ok := that.(VirtualServiceSelector_Expression)
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

	if strings.Compare(m.GetKey(), target.GetKey()) != 0 {
		return false
	}

	if m.GetOperator() != target.GetOperator() {
		return false
	}

	if len(m.GetValues()) != len(target.GetValues()) {
		return false
	}
	for idx, v := range m.GetValues() {

		if strings.Compare(v, target.GetValues()[idx]) != 0 {
			return false
		}

	}

	return true
}
