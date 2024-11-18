// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/external/envoy/config/core/v3/health_check.proto

package v3

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_type_matcher_v3 "github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/type/matcher/v3"

	github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_type_v3 "github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/type/v3"

	github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_advanced_http "github.com/solo-io/gloo/projects/controller/pkg/api/v1/options/advanced_http"

	google_golang_org_protobuf_types_known_anypb "google.golang.org/protobuf/types/known/anypb"

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
func (m *HealthCheck) Clone() proto.Message {
	var target *HealthCheck
	if m == nil {
		return target
	}
	target = &HealthCheck{}

	if h, ok := interface{}(m.GetTimeout()).(clone.Cloner); ok {
		target.Timeout = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.Timeout = proto.Clone(m.GetTimeout()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	if h, ok := interface{}(m.GetInterval()).(clone.Cloner); ok {
		target.Interval = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.Interval = proto.Clone(m.GetInterval()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	if h, ok := interface{}(m.GetInitialJitter()).(clone.Cloner); ok {
		target.InitialJitter = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.InitialJitter = proto.Clone(m.GetInitialJitter()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	if h, ok := interface{}(m.GetIntervalJitter()).(clone.Cloner); ok {
		target.IntervalJitter = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.IntervalJitter = proto.Clone(m.GetIntervalJitter()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	target.IntervalJitterPercent = m.GetIntervalJitterPercent()

	if h, ok := interface{}(m.GetUnhealthyThreshold()).(clone.Cloner); ok {
		target.UnhealthyThreshold = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.UnhealthyThreshold = proto.Clone(m.GetUnhealthyThreshold()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	if h, ok := interface{}(m.GetHealthyThreshold()).(clone.Cloner); ok {
		target.HealthyThreshold = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.HealthyThreshold = proto.Clone(m.GetHealthyThreshold()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	if h, ok := interface{}(m.GetAltPort()).(clone.Cloner); ok {
		target.AltPort = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.AltPort = proto.Clone(m.GetAltPort()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	if h, ok := interface{}(m.GetReuseConnection()).(clone.Cloner); ok {
		target.ReuseConnection = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.ReuseConnection = proto.Clone(m.GetReuseConnection()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	if h, ok := interface{}(m.GetNoTrafficInterval()).(clone.Cloner); ok {
		target.NoTrafficInterval = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.NoTrafficInterval = proto.Clone(m.GetNoTrafficInterval()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	if h, ok := interface{}(m.GetUnhealthyInterval()).(clone.Cloner); ok {
		target.UnhealthyInterval = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.UnhealthyInterval = proto.Clone(m.GetUnhealthyInterval()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	if h, ok := interface{}(m.GetUnhealthyEdgeInterval()).(clone.Cloner); ok {
		target.UnhealthyEdgeInterval = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.UnhealthyEdgeInterval = proto.Clone(m.GetUnhealthyEdgeInterval()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	if h, ok := interface{}(m.GetHealthyEdgeInterval()).(clone.Cloner); ok {
		target.HealthyEdgeInterval = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.HealthyEdgeInterval = proto.Clone(m.GetHealthyEdgeInterval()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	target.EventLogPath = m.GetEventLogPath()

	if h, ok := interface{}(m.GetEventService()).(clone.Cloner); ok {
		target.EventService = h.Clone().(*EventServiceConfig)
	} else {
		target.EventService = proto.Clone(m.GetEventService()).(*EventServiceConfig)
	}

	target.AlwaysLogHealthCheckFailures = m.GetAlwaysLogHealthCheckFailures()

	if h, ok := interface{}(m.GetTlsOptions()).(clone.Cloner); ok {
		target.TlsOptions = h.Clone().(*HealthCheck_TlsOptions)
	} else {
		target.TlsOptions = proto.Clone(m.GetTlsOptions()).(*HealthCheck_TlsOptions)
	}

	if h, ok := interface{}(m.GetTransportSocketMatchCriteria()).(clone.Cloner); ok {
		target.TransportSocketMatchCriteria = h.Clone().(*google_golang_org_protobuf_types_known_structpb.Struct)
	} else {
		target.TransportSocketMatchCriteria = proto.Clone(m.GetTransportSocketMatchCriteria()).(*google_golang_org_protobuf_types_known_structpb.Struct)
	}

	switch m.HealthChecker.(type) {

	case *HealthCheck_HttpHealthCheck_:

		if h, ok := interface{}(m.GetHttpHealthCheck()).(clone.Cloner); ok {
			target.HealthChecker = &HealthCheck_HttpHealthCheck_{
				HttpHealthCheck: h.Clone().(*HealthCheck_HttpHealthCheck),
			}
		} else {
			target.HealthChecker = &HealthCheck_HttpHealthCheck_{
				HttpHealthCheck: proto.Clone(m.GetHttpHealthCheck()).(*HealthCheck_HttpHealthCheck),
			}
		}

	case *HealthCheck_TcpHealthCheck_:

		if h, ok := interface{}(m.GetTcpHealthCheck()).(clone.Cloner); ok {
			target.HealthChecker = &HealthCheck_TcpHealthCheck_{
				TcpHealthCheck: h.Clone().(*HealthCheck_TcpHealthCheck),
			}
		} else {
			target.HealthChecker = &HealthCheck_TcpHealthCheck_{
				TcpHealthCheck: proto.Clone(m.GetTcpHealthCheck()).(*HealthCheck_TcpHealthCheck),
			}
		}

	case *HealthCheck_GrpcHealthCheck_:

		if h, ok := interface{}(m.GetGrpcHealthCheck()).(clone.Cloner); ok {
			target.HealthChecker = &HealthCheck_GrpcHealthCheck_{
				GrpcHealthCheck: h.Clone().(*HealthCheck_GrpcHealthCheck),
			}
		} else {
			target.HealthChecker = &HealthCheck_GrpcHealthCheck_{
				GrpcHealthCheck: proto.Clone(m.GetGrpcHealthCheck()).(*HealthCheck_GrpcHealthCheck),
			}
		}

	case *HealthCheck_CustomHealthCheck_:

		if h, ok := interface{}(m.GetCustomHealthCheck()).(clone.Cloner); ok {
			target.HealthChecker = &HealthCheck_CustomHealthCheck_{
				CustomHealthCheck: h.Clone().(*HealthCheck_CustomHealthCheck),
			}
		} else {
			target.HealthChecker = &HealthCheck_CustomHealthCheck_{
				CustomHealthCheck: proto.Clone(m.GetCustomHealthCheck()).(*HealthCheck_CustomHealthCheck),
			}
		}

	}

	return target
}

// Clone function
func (m *HealthCheck_Payload) Clone() proto.Message {
	var target *HealthCheck_Payload
	if m == nil {
		return target
	}
	target = &HealthCheck_Payload{}

	switch m.Payload.(type) {

	case *HealthCheck_Payload_Text:

		target.Payload = &HealthCheck_Payload_Text{
			Text: m.GetText(),
		}

	case *HealthCheck_Payload_Binary:

		if m.GetBinary() != nil {
			newArr := make([]byte, len(m.GetBinary()))
			copy(newArr, m.GetBinary())
			target.Payload = &HealthCheck_Payload_Binary{
				Binary: newArr,
			}
		} else {
			target.Payload = &HealthCheck_Payload_Binary{
				Binary: nil,
			}
		}

	}

	return target
}

// Clone function
func (m *HealthCheck_HttpHealthCheck) Clone() proto.Message {
	var target *HealthCheck_HttpHealthCheck
	if m == nil {
		return target
	}
	target = &HealthCheck_HttpHealthCheck{}

	target.Host = m.GetHost()

	target.Path = m.GetPath()

	if h, ok := interface{}(m.GetSend()).(clone.Cloner); ok {
		target.Send = h.Clone().(*HealthCheck_Payload)
	} else {
		target.Send = proto.Clone(m.GetSend()).(*HealthCheck_Payload)
	}

	if h, ok := interface{}(m.GetReceive()).(clone.Cloner); ok {
		target.Receive = h.Clone().(*HealthCheck_Payload)
	} else {
		target.Receive = proto.Clone(m.GetReceive()).(*HealthCheck_Payload)
	}

	if m.GetRequestHeadersToAdd() != nil {
		target.RequestHeadersToAdd = make([]*HeaderValueOption, len(m.GetRequestHeadersToAdd()))
		for idx, v := range m.GetRequestHeadersToAdd() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.RequestHeadersToAdd[idx] = h.Clone().(*HeaderValueOption)
			} else {
				target.RequestHeadersToAdd[idx] = proto.Clone(v).(*HeaderValueOption)
			}

		}
	}

	if m.GetRequestHeadersToRemove() != nil {
		target.RequestHeadersToRemove = make([]string, len(m.GetRequestHeadersToRemove()))
		for idx, v := range m.GetRequestHeadersToRemove() {

			target.RequestHeadersToRemove[idx] = v

		}
	}

	if m.GetExpectedStatuses() != nil {
		target.ExpectedStatuses = make([]*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_type_v3.Int64Range, len(m.GetExpectedStatuses()))
		for idx, v := range m.GetExpectedStatuses() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.ExpectedStatuses[idx] = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_type_v3.Int64Range)
			} else {
				target.ExpectedStatuses[idx] = proto.Clone(v).(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_type_v3.Int64Range)
			}

		}
	}

	target.CodecClientType = m.GetCodecClientType()

	if h, ok := interface{}(m.GetServiceNameMatcher()).(clone.Cloner); ok {
		target.ServiceNameMatcher = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_type_matcher_v3.StringMatcher)
	} else {
		target.ServiceNameMatcher = proto.Clone(m.GetServiceNameMatcher()).(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_type_matcher_v3.StringMatcher)
	}

	if h, ok := interface{}(m.GetResponseAssertions()).(clone.Cloner); ok {
		target.ResponseAssertions = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_advanced_http.ResponseAssertions)
	} else {
		target.ResponseAssertions = proto.Clone(m.GetResponseAssertions()).(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_advanced_http.ResponseAssertions)
	}

	return target
}

// Clone function
func (m *HealthCheck_TcpHealthCheck) Clone() proto.Message {
	var target *HealthCheck_TcpHealthCheck
	if m == nil {
		return target
	}
	target = &HealthCheck_TcpHealthCheck{}

	if h, ok := interface{}(m.GetSend()).(clone.Cloner); ok {
		target.Send = h.Clone().(*HealthCheck_Payload)
	} else {
		target.Send = proto.Clone(m.GetSend()).(*HealthCheck_Payload)
	}

	if m.GetReceive() != nil {
		target.Receive = make([]*HealthCheck_Payload, len(m.GetReceive()))
		for idx, v := range m.GetReceive() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Receive[idx] = h.Clone().(*HealthCheck_Payload)
			} else {
				target.Receive[idx] = proto.Clone(v).(*HealthCheck_Payload)
			}

		}
	}

	return target
}

// Clone function
func (m *HealthCheck_RedisHealthCheck) Clone() proto.Message {
	var target *HealthCheck_RedisHealthCheck
	if m == nil {
		return target
	}
	target = &HealthCheck_RedisHealthCheck{}

	target.Key = m.GetKey()

	return target
}

// Clone function
func (m *HealthCheck_GrpcHealthCheck) Clone() proto.Message {
	var target *HealthCheck_GrpcHealthCheck
	if m == nil {
		return target
	}
	target = &HealthCheck_GrpcHealthCheck{}

	target.ServiceName = m.GetServiceName()

	target.Authority = m.GetAuthority()

	return target
}

// Clone function
func (m *HealthCheck_CustomHealthCheck) Clone() proto.Message {
	var target *HealthCheck_CustomHealthCheck
	if m == nil {
		return target
	}
	target = &HealthCheck_CustomHealthCheck{}

	target.Name = m.GetName()

	switch m.ConfigType.(type) {

	case *HealthCheck_CustomHealthCheck_TypedConfig:

		if h, ok := interface{}(m.GetTypedConfig()).(clone.Cloner); ok {
			target.ConfigType = &HealthCheck_CustomHealthCheck_TypedConfig{
				TypedConfig: h.Clone().(*google_golang_org_protobuf_types_known_anypb.Any),
			}
		} else {
			target.ConfigType = &HealthCheck_CustomHealthCheck_TypedConfig{
				TypedConfig: proto.Clone(m.GetTypedConfig()).(*google_golang_org_protobuf_types_known_anypb.Any),
			}
		}

	}

	return target
}

// Clone function
func (m *HealthCheck_TlsOptions) Clone() proto.Message {
	var target *HealthCheck_TlsOptions
	if m == nil {
		return target
	}
	target = &HealthCheck_TlsOptions{}

	if m.GetAlpnProtocols() != nil {
		target.AlpnProtocols = make([]string, len(m.GetAlpnProtocols()))
		for idx, v := range m.GetAlpnProtocols() {

			target.AlpnProtocols[idx] = v

		}
	}

	return target
}
