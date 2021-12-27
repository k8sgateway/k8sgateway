// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/grpc/validation/gloo_validation.proto

package validation

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
func (m *GlooValidationServiceRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GlooValidationServiceRequest)
	if !ok {
		that2, ok := that.(GlooValidationServiceRequest)
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

	if h, ok := interface{}(m.GetProxy()).(equality.Equalizer); ok {
		if !h.Equal(target.GetProxy()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetProxy(), target.GetProxy()) {
			return false
		}
	}

	switch m.Resources.(type) {

	case *GlooValidationServiceRequest_ModifiedResources:
		if _, ok := target.Resources.(*GlooValidationServiceRequest_ModifiedResources); !ok {
			return false
		}

		if h, ok := interface{}(m.GetModifiedResources()).(equality.Equalizer); ok {
			if !h.Equal(target.GetModifiedResources()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetModifiedResources(), target.GetModifiedResources()) {
				return false
			}
		}

	case *GlooValidationServiceRequest_DeletedResources:
		if _, ok := target.Resources.(*GlooValidationServiceRequest_DeletedResources); !ok {
			return false
		}

		if h, ok := interface{}(m.GetDeletedResources()).(equality.Equalizer); ok {
			if !h.Equal(target.GetDeletedResources()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetDeletedResources(), target.GetDeletedResources()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Resources != target.Resources {
			return false
		}
	}

	return true
}

// Equal function
func (m *GlooValidationServiceResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GlooValidationServiceResponse)
	if !ok {
		that2, ok := that.(GlooValidationServiceResponse)
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

	if len(m.GetValidationReports()) != len(target.GetValidationReports()) {
		return false
	}
	for idx, v := range m.GetValidationReports() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetValidationReports()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetValidationReports()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *ModifiedResources) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ModifiedResources)
	if !ok {
		that2, ok := that.(ModifiedResources)
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

	if len(m.GetUpstreams()) != len(target.GetUpstreams()) {
		return false
	}
	for idx, v := range m.GetUpstreams() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetUpstreams()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetUpstreams()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *DeletedResources) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*DeletedResources)
	if !ok {
		that2, ok := that.(DeletedResources)
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

	if len(m.GetUpstreamRefs()) != len(target.GetUpstreamRefs()) {
		return false
	}
	for idx, v := range m.GetUpstreamRefs() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetUpstreamRefs()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetUpstreamRefs()[idx]) {
				return false
			}
		}

	}

	if len(m.GetSecretRefs()) != len(target.GetSecretRefs()) {
		return false
	}
	for idx, v := range m.GetSecretRefs() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSecretRefs()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSecretRefs()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *ValidationReport) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ValidationReport)
	if !ok {
		that2, ok := that.(ValidationReport)
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

	if h, ok := interface{}(m.GetProxyReport()).(equality.Equalizer); ok {
		if !h.Equal(target.GetProxyReport()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetProxyReport(), target.GetProxyReport()) {
			return false
		}
	}

	if len(m.GetUpstreamReports()) != len(target.GetUpstreamReports()) {
		return false
	}
	for idx, v := range m.GetUpstreamReports() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetUpstreamReports()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetUpstreamReports()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetProxy()).(equality.Equalizer); ok {
		if !h.Equal(target.GetProxy()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetProxy(), target.GetProxy()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ResourceReport) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ResourceReport)
	if !ok {
		that2, ok := that.(ResourceReport)
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

	if h, ok := interface{}(m.GetResourceRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetResourceRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetResourceRef(), target.GetResourceRef()) {
			return false
		}
	}

	if len(m.GetWarnings()) != len(target.GetWarnings()) {
		return false
	}
	for idx, v := range m.GetWarnings() {

		if strings.Compare(v, target.GetWarnings()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetErrors()) != len(target.GetErrors()) {
		return false
	}
	for idx, v := range m.GetErrors() {

		if strings.Compare(v, target.GetErrors()[idx]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *NotifyOnResyncRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*NotifyOnResyncRequest)
	if !ok {
		that2, ok := that.(NotifyOnResyncRequest)
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

	return true
}

// Equal function
func (m *NotifyOnResyncResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*NotifyOnResyncResponse)
	if !ok {
		that2, ok := that.(NotifyOnResyncResponse)
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

	return true
}

// Equal function
func (m *ProxyReport) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ProxyReport)
	if !ok {
		that2, ok := that.(ProxyReport)
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

	if len(m.GetListenerReports()) != len(target.GetListenerReports()) {
		return false
	}
	for idx, v := range m.GetListenerReports() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetListenerReports()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetListenerReports()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *ListenerReport) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ListenerReport)
	if !ok {
		that2, ok := that.(ListenerReport)
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

	if len(m.GetErrors()) != len(target.GetErrors()) {
		return false
	}
	for idx, v := range m.GetErrors() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetErrors()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetErrors()[idx]) {
				return false
			}
		}

	}

	switch m.ListenerTypeReport.(type) {

	case *ListenerReport_HttpListenerReport:
		if _, ok := target.ListenerTypeReport.(*ListenerReport_HttpListenerReport); !ok {
			return false
		}

		if h, ok := interface{}(m.GetHttpListenerReport()).(equality.Equalizer); ok {
			if !h.Equal(target.GetHttpListenerReport()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetHttpListenerReport(), target.GetHttpListenerReport()) {
				return false
			}
		}

	case *ListenerReport_TcpListenerReport:
		if _, ok := target.ListenerTypeReport.(*ListenerReport_TcpListenerReport); !ok {
			return false
		}

		if h, ok := interface{}(m.GetTcpListenerReport()).(equality.Equalizer); ok {
			if !h.Equal(target.GetTcpListenerReport()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetTcpListenerReport(), target.GetTcpListenerReport()) {
				return false
			}
		}

	case *ListenerReport_HybridListenerReport:
		if _, ok := target.ListenerTypeReport.(*ListenerReport_HybridListenerReport); !ok {
			return false
		}

		if h, ok := interface{}(m.GetHybridListenerReport()).(equality.Equalizer); ok {
			if !h.Equal(target.GetHybridListenerReport()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetHybridListenerReport(), target.GetHybridListenerReport()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.ListenerTypeReport != target.ListenerTypeReport {
			return false
		}
	}

	return true
}

// Equal function
func (m *HttpListenerReport) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HttpListenerReport)
	if !ok {
		that2, ok := that.(HttpListenerReport)
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

	if len(m.GetErrors()) != len(target.GetErrors()) {
		return false
	}
	for idx, v := range m.GetErrors() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetErrors()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetErrors()[idx]) {
				return false
			}
		}

	}

	if len(m.GetVirtualHostReports()) != len(target.GetVirtualHostReports()) {
		return false
	}
	for idx, v := range m.GetVirtualHostReports() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetVirtualHostReports()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetVirtualHostReports()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *VirtualHostReport) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualHostReport)
	if !ok {
		that2, ok := that.(VirtualHostReport)
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

	if len(m.GetErrors()) != len(target.GetErrors()) {
		return false
	}
	for idx, v := range m.GetErrors() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetErrors()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetErrors()[idx]) {
				return false
			}
		}

	}

	if len(m.GetRouteReports()) != len(target.GetRouteReports()) {
		return false
	}
	for idx, v := range m.GetRouteReports() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetRouteReports()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetRouteReports()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *RouteReport) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RouteReport)
	if !ok {
		that2, ok := that.(RouteReport)
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

	if len(m.GetErrors()) != len(target.GetErrors()) {
		return false
	}
	for idx, v := range m.GetErrors() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetErrors()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetErrors()[idx]) {
				return false
			}
		}

	}

	if len(m.GetWarnings()) != len(target.GetWarnings()) {
		return false
	}
	for idx, v := range m.GetWarnings() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetWarnings()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetWarnings()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *TcpListenerReport) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TcpListenerReport)
	if !ok {
		that2, ok := that.(TcpListenerReport)
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

	if len(m.GetErrors()) != len(target.GetErrors()) {
		return false
	}
	for idx, v := range m.GetErrors() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetErrors()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetErrors()[idx]) {
				return false
			}
		}

	}

	if len(m.GetTcpHostReports()) != len(target.GetTcpHostReports()) {
		return false
	}
	for idx, v := range m.GetTcpHostReports() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetTcpHostReports()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetTcpHostReports()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *TcpHostReport) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TcpHostReport)
	if !ok {
		that2, ok := that.(TcpHostReport)
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

	if len(m.GetErrors()) != len(target.GetErrors()) {
		return false
	}
	for idx, v := range m.GetErrors() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetErrors()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetErrors()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *HybridListenerReport) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HybridListenerReport)
	if !ok {
		that2, ok := that.(HybridListenerReport)
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

	if len(m.GetMatchedListenerReports()) != len(target.GetMatchedListenerReports()) {
		return false
	}
	for k, v := range m.GetMatchedListenerReports() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetMatchedListenerReports()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetMatchedListenerReports()[k]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *MatchedListenerReport) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*MatchedListenerReport)
	if !ok {
		that2, ok := that.(MatchedListenerReport)
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

	switch m.ListenerReportType.(type) {

	case *MatchedListenerReport_HttpListenerReport:
		if _, ok := target.ListenerReportType.(*MatchedListenerReport_HttpListenerReport); !ok {
			return false
		}

		if h, ok := interface{}(m.GetHttpListenerReport()).(equality.Equalizer); ok {
			if !h.Equal(target.GetHttpListenerReport()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetHttpListenerReport(), target.GetHttpListenerReport()) {
				return false
			}
		}

	case *MatchedListenerReport_TcpListenerReport:
		if _, ok := target.ListenerReportType.(*MatchedListenerReport_TcpListenerReport); !ok {
			return false
		}

		if h, ok := interface{}(m.GetTcpListenerReport()).(equality.Equalizer); ok {
			if !h.Equal(target.GetTcpListenerReport()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetTcpListenerReport(), target.GetTcpListenerReport()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.ListenerReportType != target.ListenerReportType {
			return false
		}
	}

	return true
}

// Equal function
func (m *ListenerReport_Error) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ListenerReport_Error)
	if !ok {
		that2, ok := that.(ListenerReport_Error)
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

	if m.GetType() != target.GetType() {
		return false
	}

	if strings.Compare(m.GetReason(), target.GetReason()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *HttpListenerReport_Error) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HttpListenerReport_Error)
	if !ok {
		that2, ok := that.(HttpListenerReport_Error)
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

	if m.GetType() != target.GetType() {
		return false
	}

	if strings.Compare(m.GetReason(), target.GetReason()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *VirtualHostReport_Error) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualHostReport_Error)
	if !ok {
		that2, ok := that.(VirtualHostReport_Error)
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

	if m.GetType() != target.GetType() {
		return false
	}

	if strings.Compare(m.GetReason(), target.GetReason()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *RouteReport_Error) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RouteReport_Error)
	if !ok {
		that2, ok := that.(RouteReport_Error)
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

	if m.GetType() != target.GetType() {
		return false
	}

	if strings.Compare(m.GetReason(), target.GetReason()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *RouteReport_Warning) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RouteReport_Warning)
	if !ok {
		that2, ok := that.(RouteReport_Warning)
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

	if m.GetType() != target.GetType() {
		return false
	}

	if strings.Compare(m.GetReason(), target.GetReason()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *TcpListenerReport_Error) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TcpListenerReport_Error)
	if !ok {
		that2, ok := that.(TcpListenerReport_Error)
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

	if m.GetType() != target.GetType() {
		return false
	}

	if strings.Compare(m.GetReason(), target.GetReason()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *TcpHostReport_Error) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TcpHostReport_Error)
	if !ok {
		that2, ok := that.(TcpHostReport_Error)
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

	if m.GetType() != target.GetType() {
		return false
	}

	if strings.Compare(m.GetReason(), target.GetReason()) != 0 {
		return false
	}

	return true
}
