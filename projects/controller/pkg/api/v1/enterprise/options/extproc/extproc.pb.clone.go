// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/v1/enterprise/options/extproc/extproc.proto

package extproc

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_config_common_mutation_rules_v3 "github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/config/common/mutation_rules/v3"

	github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_config_core_v3 "github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/config/core/v3"

	github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_extensions_filters_http_ext_proc_v3 "github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/extensions/filters/http/ext_proc/v3"

	github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_type_matcher_v3 "github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/type/matcher/v3"

	github_com_solo_io_gloo_projects_controller_pkg_api_v1_filters "github.com/solo-io/gloo/projects/controller/pkg/api/v1/filters"

	github_com_solo_io_solo_kit_pkg_api_v1_resources_core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

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
func (m *Settings) Clone() proto.Message {
	var target *Settings
	if m == nil {
		return target
	}
	target = &Settings{}

	if h, ok := interface{}(m.GetGrpcService()).(clone.Cloner); ok {
		target.GrpcService = h.Clone().(*GrpcService)
	} else {
		target.GrpcService = proto.Clone(m.GetGrpcService()).(*GrpcService)
	}

	if h, ok := interface{}(m.GetFilterStage()).(clone.Cloner); ok {
		target.FilterStage = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_filters.FilterStage)
	} else {
		target.FilterStage = proto.Clone(m.GetFilterStage()).(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_filters.FilterStage)
	}

	if h, ok := interface{}(m.GetFailureModeAllow()).(clone.Cloner); ok {
		target.FailureModeAllow = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.FailureModeAllow = proto.Clone(m.GetFailureModeAllow()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	if h, ok := interface{}(m.GetProcessingMode()).(clone.Cloner); ok {
		target.ProcessingMode = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_extensions_filters_http_ext_proc_v3.ProcessingMode)
	} else {
		target.ProcessingMode = proto.Clone(m.GetProcessingMode()).(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_extensions_filters_http_ext_proc_v3.ProcessingMode)
	}

	if h, ok := interface{}(m.GetAsyncMode()).(clone.Cloner); ok {
		target.AsyncMode = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.AsyncMode = proto.Clone(m.GetAsyncMode()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	if m.GetRequestAttributes() != nil {
		target.RequestAttributes = make([]string, len(m.GetRequestAttributes()))
		for idx, v := range m.GetRequestAttributes() {

			target.RequestAttributes[idx] = v

		}
	}

	if m.GetResponseAttributes() != nil {
		target.ResponseAttributes = make([]string, len(m.GetResponseAttributes()))
		for idx, v := range m.GetResponseAttributes() {

			target.ResponseAttributes[idx] = v

		}
	}

	if h, ok := interface{}(m.GetMessageTimeout()).(clone.Cloner); ok {
		target.MessageTimeout = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.MessageTimeout = proto.Clone(m.GetMessageTimeout()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	if h, ok := interface{}(m.GetStatPrefix()).(clone.Cloner); ok {
		target.StatPrefix = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	} else {
		target.StatPrefix = proto.Clone(m.GetStatPrefix()).(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	}

	if h, ok := interface{}(m.GetMutationRules()).(clone.Cloner); ok {
		target.MutationRules = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_config_common_mutation_rules_v3.HeaderMutationRules)
	} else {
		target.MutationRules = proto.Clone(m.GetMutationRules()).(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_config_common_mutation_rules_v3.HeaderMutationRules)
	}

	if h, ok := interface{}(m.GetMaxMessageTimeout()).(clone.Cloner); ok {
		target.MaxMessageTimeout = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.MaxMessageTimeout = proto.Clone(m.GetMaxMessageTimeout()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	if h, ok := interface{}(m.GetDisableClearRouteCache()).(clone.Cloner); ok {
		target.DisableClearRouteCache = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.DisableClearRouteCache = proto.Clone(m.GetDisableClearRouteCache()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	if h, ok := interface{}(m.GetForwardRules()).(clone.Cloner); ok {
		target.ForwardRules = h.Clone().(*HeaderForwardingRules)
	} else {
		target.ForwardRules = proto.Clone(m.GetForwardRules()).(*HeaderForwardingRules)
	}

	if h, ok := interface{}(m.GetFilterMetadata()).(clone.Cloner); ok {
		target.FilterMetadata = h.Clone().(*google_golang_org_protobuf_types_known_structpb.Struct)
	} else {
		target.FilterMetadata = proto.Clone(m.GetFilterMetadata()).(*google_golang_org_protobuf_types_known_structpb.Struct)
	}

	if h, ok := interface{}(m.GetAllowModeOverride()).(clone.Cloner); ok {
		target.AllowModeOverride = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.AllowModeOverride = proto.Clone(m.GetAllowModeOverride()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	if m.GetMetadataContextNamespaces() != nil {
		target.MetadataContextNamespaces = make([]string, len(m.GetMetadataContextNamespaces()))
		for idx, v := range m.GetMetadataContextNamespaces() {

			target.MetadataContextNamespaces[idx] = v

		}
	}

	if m.GetTypedMetadataContextNamespaces() != nil {
		target.TypedMetadataContextNamespaces = make([]string, len(m.GetTypedMetadataContextNamespaces()))
		for idx, v := range m.GetTypedMetadataContextNamespaces() {

			target.TypedMetadataContextNamespaces[idx] = v

		}
	}

	return target
}

// Clone function
func (m *RouteSettings) Clone() proto.Message {
	var target *RouteSettings
	if m == nil {
		return target
	}
	target = &RouteSettings{}

	switch m.Override.(type) {

	case *RouteSettings_Disabled:

		if h, ok := interface{}(m.GetDisabled()).(clone.Cloner); ok {
			target.Override = &RouteSettings_Disabled{
				Disabled: h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue),
			}
		} else {
			target.Override = &RouteSettings_Disabled{
				Disabled: proto.Clone(m.GetDisabled()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue),
			}
		}

	case *RouteSettings_Overrides:

		if h, ok := interface{}(m.GetOverrides()).(clone.Cloner); ok {
			target.Override = &RouteSettings_Overrides{
				Overrides: h.Clone().(*Overrides),
			}
		} else {
			target.Override = &RouteSettings_Overrides{
				Overrides: proto.Clone(m.GetOverrides()).(*Overrides),
			}
		}

	}

	return target
}

// Clone function
func (m *GrpcService) Clone() proto.Message {
	var target *GrpcService
	if m == nil {
		return target
	}
	target = &GrpcService{}

	if h, ok := interface{}(m.GetExtProcServerRef()).(clone.Cloner); ok {
		target.ExtProcServerRef = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	} else {
		target.ExtProcServerRef = proto.Clone(m.GetExtProcServerRef()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	}

	if h, ok := interface{}(m.GetAuthority()).(clone.Cloner); ok {
		target.Authority = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	} else {
		target.Authority = proto.Clone(m.GetAuthority()).(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	}

	if h, ok := interface{}(m.GetRetryPolicy()).(clone.Cloner); ok {
		target.RetryPolicy = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_config_core_v3.RetryPolicy)
	} else {
		target.RetryPolicy = proto.Clone(m.GetRetryPolicy()).(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_config_core_v3.RetryPolicy)
	}

	if h, ok := interface{}(m.GetTimeout()).(clone.Cloner); ok {
		target.Timeout = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.Timeout = proto.Clone(m.GetTimeout()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	if m.GetInitialMetadata() != nil {
		target.InitialMetadata = make([]*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_config_core_v3.HeaderValue, len(m.GetInitialMetadata()))
		for idx, v := range m.GetInitialMetadata() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.InitialMetadata[idx] = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_config_core_v3.HeaderValue)
			} else {
				target.InitialMetadata[idx] = proto.Clone(v).(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_config_core_v3.HeaderValue)
			}

		}
	}

	return target
}

// Clone function
func (m *Overrides) Clone() proto.Message {
	var target *Overrides
	if m == nil {
		return target
	}
	target = &Overrides{}

	if h, ok := interface{}(m.GetProcessingMode()).(clone.Cloner); ok {
		target.ProcessingMode = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_extensions_filters_http_ext_proc_v3.ProcessingMode)
	} else {
		target.ProcessingMode = proto.Clone(m.GetProcessingMode()).(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_extensions_filters_http_ext_proc_v3.ProcessingMode)
	}

	if h, ok := interface{}(m.GetAsyncMode()).(clone.Cloner); ok {
		target.AsyncMode = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.AsyncMode = proto.Clone(m.GetAsyncMode()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	if m.GetRequestAttributes() != nil {
		target.RequestAttributes = make([]string, len(m.GetRequestAttributes()))
		for idx, v := range m.GetRequestAttributes() {

			target.RequestAttributes[idx] = v

		}
	}

	if m.GetResponseAttributes() != nil {
		target.ResponseAttributes = make([]string, len(m.GetResponseAttributes()))
		for idx, v := range m.GetResponseAttributes() {

			target.ResponseAttributes[idx] = v

		}
	}

	if h, ok := interface{}(m.GetGrpcService()).(clone.Cloner); ok {
		target.GrpcService = h.Clone().(*GrpcService)
	} else {
		target.GrpcService = proto.Clone(m.GetGrpcService()).(*GrpcService)
	}

	if m.GetMetadataContextNamespaces() != nil {
		target.MetadataContextNamespaces = make([]string, len(m.GetMetadataContextNamespaces()))
		for idx, v := range m.GetMetadataContextNamespaces() {

			target.MetadataContextNamespaces[idx] = v

		}
	}

	if m.GetTypedMetadataContextNamespaces() != nil {
		target.TypedMetadataContextNamespaces = make([]string, len(m.GetTypedMetadataContextNamespaces()))
		for idx, v := range m.GetTypedMetadataContextNamespaces() {

			target.TypedMetadataContextNamespaces[idx] = v

		}
	}

	return target
}

// Clone function
func (m *HeaderForwardingRules) Clone() proto.Message {
	var target *HeaderForwardingRules
	if m == nil {
		return target
	}
	target = &HeaderForwardingRules{}

	if h, ok := interface{}(m.GetAllowedHeaders()).(clone.Cloner); ok {
		target.AllowedHeaders = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_type_matcher_v3.ListStringMatcher)
	} else {
		target.AllowedHeaders = proto.Clone(m.GetAllowedHeaders()).(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_type_matcher_v3.ListStringMatcher)
	}

	if h, ok := interface{}(m.GetDisallowedHeaders()).(clone.Cloner); ok {
		target.DisallowedHeaders = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_type_matcher_v3.ListStringMatcher)
	} else {
		target.DisallowedHeaders = proto.Clone(m.GetDisallowedHeaders()).(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_type_matcher_v3.ListStringMatcher)
	}

	return target
}
