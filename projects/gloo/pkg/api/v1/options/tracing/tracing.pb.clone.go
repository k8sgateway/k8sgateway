// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/tracing/tracing.proto

package tracing

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_trace_v3 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/config/trace/v3"

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
func (m *ListenerTracingSettings) Clone() proto.Message {
	var target *ListenerTracingSettings
	if m == nil {
		return target
	}
	target = &ListenerTracingSettings{}

	if m.GetRequestHeadersForTags() != nil {
		target.RequestHeadersForTags = make([]*google_golang_org_protobuf_types_known_wrapperspb.StringValue, len(m.GetRequestHeadersForTags()))
		for idx, v := range m.GetRequestHeadersForTags() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.RequestHeadersForTags[idx] = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
			} else {
				target.RequestHeadersForTags[idx] = proto.Clone(v).(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
			}

		}
	}

	if h, ok := interface{}(m.GetVerbose()).(clone.Cloner); ok {
		target.Verbose = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.Verbose = proto.Clone(m.GetVerbose()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	if h, ok := interface{}(m.GetTracePercentages()).(clone.Cloner); ok {
		target.TracePercentages = h.Clone().(*TracePercentages)
	} else {
		target.TracePercentages = proto.Clone(m.GetTracePercentages()).(*TracePercentages)
	}

	if m.GetEnvironmentVariablesForTags() != nil {
		target.EnvironmentVariablesForTags = make([]*TracingTagEnvironmentVariable, len(m.GetEnvironmentVariablesForTags()))
		for idx, v := range m.GetEnvironmentVariablesForTags() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.EnvironmentVariablesForTags[idx] = h.Clone().(*TracingTagEnvironmentVariable)
			} else {
				target.EnvironmentVariablesForTags[idx] = proto.Clone(v).(*TracingTagEnvironmentVariable)
			}

		}
	}

	if m.GetLiteralsForTags() != nil {
		target.LiteralsForTags = make([]*TracingTagLiteral, len(m.GetLiteralsForTags()))
		for idx, v := range m.GetLiteralsForTags() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.LiteralsForTags[idx] = h.Clone().(*TracingTagLiteral)
			} else {
				target.LiteralsForTags[idx] = proto.Clone(v).(*TracingTagLiteral)
			}

		}
	}

	if m.GetMetadataForTags() != nil {
		target.MetadataForTags = make([]*TracingTagMetadata, len(m.GetMetadataForTags()))
		for idx, v := range m.GetMetadataForTags() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.MetadataForTags[idx] = h.Clone().(*TracingTagMetadata)
			} else {
				target.MetadataForTags[idx] = proto.Clone(v).(*TracingTagMetadata)
			}

		}
	}

	target.SpawnUpstreamSpan = m.GetSpawnUpstreamSpan()

	switch m.ProviderConfig.(type) {

	case *ListenerTracingSettings_ZipkinConfig:

		if h, ok := interface{}(m.GetZipkinConfig()).(clone.Cloner); ok {
			target.ProviderConfig = &ListenerTracingSettings_ZipkinConfig{
				ZipkinConfig: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_trace_v3.ZipkinConfig),
			}
		} else {
			target.ProviderConfig = &ListenerTracingSettings_ZipkinConfig{
				ZipkinConfig: proto.Clone(m.GetZipkinConfig()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_trace_v3.ZipkinConfig),
			}
		}

	case *ListenerTracingSettings_DatadogConfig:

		if h, ok := interface{}(m.GetDatadogConfig()).(clone.Cloner); ok {
			target.ProviderConfig = &ListenerTracingSettings_DatadogConfig{
				DatadogConfig: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_trace_v3.DatadogConfig),
			}
		} else {
			target.ProviderConfig = &ListenerTracingSettings_DatadogConfig{
				DatadogConfig: proto.Clone(m.GetDatadogConfig()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_trace_v3.DatadogConfig),
			}
		}

	case *ListenerTracingSettings_OpenTelemetryConfig:

		if h, ok := interface{}(m.GetOpenTelemetryConfig()).(clone.Cloner); ok {
			target.ProviderConfig = &ListenerTracingSettings_OpenTelemetryConfig{
				OpenTelemetryConfig: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_trace_v3.OpenTelemetryConfig),
			}
		} else {
			target.ProviderConfig = &ListenerTracingSettings_OpenTelemetryConfig{
				OpenTelemetryConfig: proto.Clone(m.GetOpenTelemetryConfig()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_trace_v3.OpenTelemetryConfig),
			}
		}

	case *ListenerTracingSettings_OpenCensusConfig:

		if h, ok := interface{}(m.GetOpenCensusConfig()).(clone.Cloner); ok {
			target.ProviderConfig = &ListenerTracingSettings_OpenCensusConfig{
				OpenCensusConfig: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_trace_v3.OpenCensusConfig),
			}
		} else {
			target.ProviderConfig = &ListenerTracingSettings_OpenCensusConfig{
				OpenCensusConfig: proto.Clone(m.GetOpenCensusConfig()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_trace_v3.OpenCensusConfig),
			}
		}

	}

	return target
}

// Clone function
func (m *RouteTracingSettings) Clone() proto.Message {
	var target *RouteTracingSettings
	if m == nil {
		return target
	}
	target = &RouteTracingSettings{}

	target.RouteDescriptor = m.GetRouteDescriptor()

	if h, ok := interface{}(m.GetTracePercentages()).(clone.Cloner); ok {
		target.TracePercentages = h.Clone().(*TracePercentages)
	} else {
		target.TracePercentages = proto.Clone(m.GetTracePercentages()).(*TracePercentages)
	}

	if h, ok := interface{}(m.GetPropagate()).(clone.Cloner); ok {
		target.Propagate = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.Propagate = proto.Clone(m.GetPropagate()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	return target
}

// Clone function
func (m *TracePercentages) Clone() proto.Message {
	var target *TracePercentages
	if m == nil {
		return target
	}
	target = &TracePercentages{}

	if h, ok := interface{}(m.GetClientSamplePercentage()).(clone.Cloner); ok {
		target.ClientSamplePercentage = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.FloatValue)
	} else {
		target.ClientSamplePercentage = proto.Clone(m.GetClientSamplePercentage()).(*google_golang_org_protobuf_types_known_wrapperspb.FloatValue)
	}

	if h, ok := interface{}(m.GetRandomSamplePercentage()).(clone.Cloner); ok {
		target.RandomSamplePercentage = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.FloatValue)
	} else {
		target.RandomSamplePercentage = proto.Clone(m.GetRandomSamplePercentage()).(*google_golang_org_protobuf_types_known_wrapperspb.FloatValue)
	}

	if h, ok := interface{}(m.GetOverallSamplePercentage()).(clone.Cloner); ok {
		target.OverallSamplePercentage = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.FloatValue)
	} else {
		target.OverallSamplePercentage = proto.Clone(m.GetOverallSamplePercentage()).(*google_golang_org_protobuf_types_known_wrapperspb.FloatValue)
	}

	return target
}

// Clone function
func (m *TracingTagEnvironmentVariable) Clone() proto.Message {
	var target *TracingTagEnvironmentVariable
	if m == nil {
		return target
	}
	target = &TracingTagEnvironmentVariable{}

	if h, ok := interface{}(m.GetTag()).(clone.Cloner); ok {
		target.Tag = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	} else {
		target.Tag = proto.Clone(m.GetTag()).(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	}

	if h, ok := interface{}(m.GetName()).(clone.Cloner); ok {
		target.Name = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	} else {
		target.Name = proto.Clone(m.GetName()).(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	}

	if h, ok := interface{}(m.GetDefaultValue()).(clone.Cloner); ok {
		target.DefaultValue = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	} else {
		target.DefaultValue = proto.Clone(m.GetDefaultValue()).(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	}

	return target
}

// Clone function
func (m *TracingTagLiteral) Clone() proto.Message {
	var target *TracingTagLiteral
	if m == nil {
		return target
	}
	target = &TracingTagLiteral{}

	if h, ok := interface{}(m.GetTag()).(clone.Cloner); ok {
		target.Tag = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	} else {
		target.Tag = proto.Clone(m.GetTag()).(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	}

	if h, ok := interface{}(m.GetValue()).(clone.Cloner); ok {
		target.Value = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	} else {
		target.Value = proto.Clone(m.GetValue()).(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	}

	return target
}

// Clone function
func (m *TracingTagMetadata) Clone() proto.Message {
	var target *TracingTagMetadata
	if m == nil {
		return target
	}
	target = &TracingTagMetadata{}

	target.Tag = m.GetTag()

	target.Kind = m.GetKind()

	if h, ok := interface{}(m.GetValue()).(clone.Cloner); ok {
		target.Value = h.Clone().(*TracingTagMetadata_MetadataValue)
	} else {
		target.Value = proto.Clone(m.GetValue()).(*TracingTagMetadata_MetadataValue)
	}

	target.DefaultValue = m.GetDefaultValue()

	return target
}

// Clone function
func (m *TracingTagMetadata_MetadataValue) Clone() proto.Message {
	var target *TracingTagMetadata_MetadataValue
	if m == nil {
		return target
	}
	target = &TracingTagMetadata_MetadataValue{}

	target.Namespace = m.GetNamespace()

	target.Key = m.GetKey()

	target.NestedFieldDelimiter = m.GetNestedFieldDelimiter()

	return target
}
