// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/graphql/v1alpha1/graphql.proto

package v1alpha1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_struct "github.com/golang/protobuf/ptypes/struct"

	github_com_golang_protobuf_ptypes_wrappers "github.com/golang/protobuf/ptypes/wrappers"

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
func (m *RequestTemplate) Clone() proto.Message {
	var target *RequestTemplate
	if m == nil {
		return target
	}
	target = &RequestTemplate{}

	if m.GetHeaders() != nil {
		target.Headers = make(map[string]string, len(m.GetHeaders()))
		for k, v := range m.GetHeaders() {

			target.Headers[k] = v

		}
	}

	if m.GetQueryParams() != nil {
		target.QueryParams = make(map[string]string, len(m.GetQueryParams()))
		for k, v := range m.GetQueryParams() {

			target.QueryParams[k] = v

		}
	}

	if h, ok := interface{}(m.GetBody()).(clone.Cloner); ok {
		target.Body = h.Clone().(*github_com_golang_protobuf_ptypes_struct.Value)
	} else {
		target.Body = proto.Clone(m.GetBody()).(*github_com_golang_protobuf_ptypes_struct.Value)
	}

	return target
}

// Clone function
func (m *ResponseTemplate) Clone() proto.Message {
	var target *ResponseTemplate
	if m == nil {
		return target
	}
	target = &ResponseTemplate{}

	target.ResultRoot = m.GetResultRoot()

	if m.GetSetters() != nil {
		target.Setters = make(map[string]string, len(m.GetSetters()))
		for k, v := range m.GetSetters() {

			target.Setters[k] = v

		}
	}

	return target
}

// Clone function
func (m *GrpcRequestTemplate) Clone() proto.Message {
	var target *GrpcRequestTemplate
	if m == nil {
		return target
	}
	target = &GrpcRequestTemplate{}

	if h, ok := interface{}(m.GetOutgoingMessageJson()).(clone.Cloner); ok {
		target.OutgoingMessageJson = h.Clone().(*github_com_golang_protobuf_ptypes_struct.Value)
	} else {
		target.OutgoingMessageJson = proto.Clone(m.GetOutgoingMessageJson()).(*github_com_golang_protobuf_ptypes_struct.Value)
	}

	target.ServiceName = m.GetServiceName()

	target.MethodName = m.GetMethodName()

	if m.GetRequestMetadata() != nil {
		target.RequestMetadata = make(map[string]string, len(m.GetRequestMetadata()))
		for k, v := range m.GetRequestMetadata() {

			target.RequestMetadata[k] = v

		}
	}

	return target
}

// Clone function
func (m *RESTResolver) Clone() proto.Message {
	var target *RESTResolver
	if m == nil {
		return target
	}
	target = &RESTResolver{}

	if h, ok := interface{}(m.GetUpstreamRef()).(clone.Cloner); ok {
		target.UpstreamRef = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	} else {
		target.UpstreamRef = proto.Clone(m.GetUpstreamRef()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	}

	if h, ok := interface{}(m.GetRequest()).(clone.Cloner); ok {
		target.Request = h.Clone().(*RequestTemplate)
	} else {
		target.Request = proto.Clone(m.GetRequest()).(*RequestTemplate)
	}

	if h, ok := interface{}(m.GetResponse()).(clone.Cloner); ok {
		target.Response = h.Clone().(*ResponseTemplate)
	} else {
		target.Response = proto.Clone(m.GetResponse()).(*ResponseTemplate)
	}

	target.SpanName = m.GetSpanName()

	return target
}

// Clone function
func (m *GrpcDescriptorRegistry) Clone() proto.Message {
	var target *GrpcDescriptorRegistry
	if m == nil {
		return target
	}
	target = &GrpcDescriptorRegistry{}

	switch m.DescriptorSet.(type) {

	case *GrpcDescriptorRegistry_ProtoDescriptor:

		target.DescriptorSet = &GrpcDescriptorRegistry_ProtoDescriptor{
			ProtoDescriptor: m.GetProtoDescriptor(),
		}

	case *GrpcDescriptorRegistry_ProtoDescriptorBin:

		if m.GetProtoDescriptorBin() != nil {
			newArr := make([]byte, len(m.GetProtoDescriptorBin()))
			copy(newArr, m.GetProtoDescriptorBin())
			target.DescriptorSet = &GrpcDescriptorRegistry_ProtoDescriptorBin{
				ProtoDescriptorBin: newArr,
			}
		} else {
			target.DescriptorSet = &GrpcDescriptorRegistry_ProtoDescriptorBin{
				ProtoDescriptorBin: nil,
			}
		}

	}

	return target
}

// Clone function
func (m *GrpcResolver) Clone() proto.Message {
	var target *GrpcResolver
	if m == nil {
		return target
	}
	target = &GrpcResolver{}

	if h, ok := interface{}(m.GetUpstreamRef()).(clone.Cloner); ok {
		target.UpstreamRef = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	} else {
		target.UpstreamRef = proto.Clone(m.GetUpstreamRef()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	}

	if h, ok := interface{}(m.GetRequestTransform()).(clone.Cloner); ok {
		target.RequestTransform = h.Clone().(*GrpcRequestTemplate)
	} else {
		target.RequestTransform = proto.Clone(m.GetRequestTransform()).(*GrpcRequestTemplate)
	}

	target.SpanName = m.GetSpanName()

	return target
}

// Clone function
func (m *GatewaySchema) Clone() proto.Message {
	var target *GatewaySchema
	if m == nil {
		return target
	}
	target = &GatewaySchema{}

	if m.GetSubschemas() != nil {
		target.Subschemas = make([]*GatewaySchema_SubschemaConfig, len(m.GetSubschemas()))
		for idx, v := range m.GetSubschemas() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Subschemas[idx] = h.Clone().(*GatewaySchema_SubschemaConfig)
			} else {
				target.Subschemas[idx] = proto.Clone(v).(*GatewaySchema_SubschemaConfig)
			}

		}
	}

	return target
}

// Clone function
func (m *Resolution) Clone() proto.Message {
	var target *Resolution
	if m == nil {
		return target
	}
	target = &Resolution{}

	if h, ok := interface{}(m.GetStatPrefix()).(clone.Cloner); ok {
		target.StatPrefix = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.StringValue)
	} else {
		target.StatPrefix = proto.Clone(m.GetStatPrefix()).(*github_com_golang_protobuf_ptypes_wrappers.StringValue)
	}

	switch m.Resolver.(type) {

	case *Resolution_RestResolver:

		if h, ok := interface{}(m.GetRestResolver()).(clone.Cloner); ok {
			target.Resolver = &Resolution_RestResolver{
				RestResolver: h.Clone().(*RESTResolver),
			}
		} else {
			target.Resolver = &Resolution_RestResolver{
				RestResolver: proto.Clone(m.GetRestResolver()).(*RESTResolver),
			}
		}

	case *Resolution_GrpcResolver:

		if h, ok := interface{}(m.GetGrpcResolver()).(clone.Cloner); ok {
			target.Resolver = &Resolution_GrpcResolver{
				GrpcResolver: h.Clone().(*GrpcResolver),
			}
		} else {
			target.Resolver = &Resolution_GrpcResolver{
				GrpcResolver: proto.Clone(m.GetGrpcResolver()).(*GrpcResolver),
			}
		}

	}

	return target
}

// Clone function
func (m *GraphQLApi) Clone() proto.Message {
	var target *GraphQLApi
	if m == nil {
		return target
	}
	target = &GraphQLApi{}

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

	if h, ok := interface{}(m.GetStatPrefix()).(clone.Cloner); ok {
		target.StatPrefix = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.StringValue)
	} else {
		target.StatPrefix = proto.Clone(m.GetStatPrefix()).(*github_com_golang_protobuf_ptypes_wrappers.StringValue)
	}

	if h, ok := interface{}(m.GetPersistedQueryCacheConfig()).(clone.Cloner); ok {
		target.PersistedQueryCacheConfig = h.Clone().(*PersistedQueryCacheConfig)
	} else {
		target.PersistedQueryCacheConfig = proto.Clone(m.GetPersistedQueryCacheConfig()).(*PersistedQueryCacheConfig)
	}

	if m.GetAllowedQueryHashes() != nil {
		target.AllowedQueryHashes = make([]string, len(m.GetAllowedQueryHashes()))
		for idx, v := range m.GetAllowedQueryHashes() {

			target.AllowedQueryHashes[idx] = v

		}
	}

	switch m.Schema.(type) {

	case *GraphQLApi_ExecutableSchema:

		if h, ok := interface{}(m.GetExecutableSchema()).(clone.Cloner); ok {
			target.Schema = &GraphQLApi_ExecutableSchema{
				ExecutableSchema: h.Clone().(*ExecutableSchema),
			}
		} else {
			target.Schema = &GraphQLApi_ExecutableSchema{
				ExecutableSchema: proto.Clone(m.GetExecutableSchema()).(*ExecutableSchema),
			}
		}

	case *GraphQLApi_GatewaySchema:

		if h, ok := interface{}(m.GetGatewaySchema()).(clone.Cloner); ok {
			target.Schema = &GraphQLApi_GatewaySchema{
				GatewaySchema: h.Clone().(*GatewaySchema),
			}
		} else {
			target.Schema = &GraphQLApi_GatewaySchema{
				GatewaySchema: proto.Clone(m.GetGatewaySchema()).(*GatewaySchema),
			}
		}

	}

	return target
}

// Clone function
func (m *PersistedQueryCacheConfig) Clone() proto.Message {
	var target *PersistedQueryCacheConfig
	if m == nil {
		return target
	}
	target = &PersistedQueryCacheConfig{}

	target.CacheSize = m.GetCacheSize()

	return target
}

// Clone function
func (m *ExecutableSchema) Clone() proto.Message {
	var target *ExecutableSchema
	if m == nil {
		return target
	}
	target = &ExecutableSchema{}

	target.SchemaDefinition = m.GetSchemaDefinition()

	if h, ok := interface{}(m.GetExecutor()).(clone.Cloner); ok {
		target.Executor = h.Clone().(*Executor)
	} else {
		target.Executor = proto.Clone(m.GetExecutor()).(*Executor)
	}

	if h, ok := interface{}(m.GetGrpcDescriptorRegistry()).(clone.Cloner); ok {
		target.GrpcDescriptorRegistry = h.Clone().(*GrpcDescriptorRegistry)
	} else {
		target.GrpcDescriptorRegistry = proto.Clone(m.GetGrpcDescriptorRegistry()).(*GrpcDescriptorRegistry)
	}

	return target
}

// Clone function
func (m *Executor) Clone() proto.Message {
	var target *Executor
	if m == nil {
		return target
	}
	target = &Executor{}

	switch m.Executor.(type) {

	case *Executor_Local_:

		if h, ok := interface{}(m.GetLocal()).(clone.Cloner); ok {
			target.Executor = &Executor_Local_{
				Local: h.Clone().(*Executor_Local),
			}
		} else {
			target.Executor = &Executor_Local_{
				Local: proto.Clone(m.GetLocal()).(*Executor_Local),
			}
		}

	}

	return target
}

// Clone function
func (m *GatewaySchema_SubschemaConfig) Clone() proto.Message {
	var target *GatewaySchema_SubschemaConfig
	if m == nil {
		return target
	}
	target = &GatewaySchema_SubschemaConfig{}

	target.Name = m.GetName()

	target.Namespace = m.GetNamespace()

	if m.GetTypeMerge() != nil {
		target.TypeMerge = make(map[string]*GatewaySchema_SubschemaConfig_TypeMergeConfig, len(m.GetTypeMerge()))
		for k, v := range m.GetTypeMerge() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.TypeMerge[k] = h.Clone().(*GatewaySchema_SubschemaConfig_TypeMergeConfig)
			} else {
				target.TypeMerge[k] = proto.Clone(v).(*GatewaySchema_SubschemaConfig_TypeMergeConfig)
			}

		}
	}

	return target
}

// Clone function
func (m *GatewaySchema_SubschemaConfig_TypeMergeConfig) Clone() proto.Message {
	var target *GatewaySchema_SubschemaConfig_TypeMergeConfig
	if m == nil {
		return target
	}
	target = &GatewaySchema_SubschemaConfig_TypeMergeConfig{}

	target.SelectionSet = m.GetSelectionSet()

	target.QueryName = m.GetQueryName()

	if m.GetArgs() != nil {
		target.Args = make(map[string]string, len(m.GetArgs()))
		for k, v := range m.GetArgs() {

			target.Args[k] = v

		}
	}

	return target
}

// Clone function
func (m *Executor_Local) Clone() proto.Message {
	var target *Executor_Local
	if m == nil {
		return target
	}
	target = &Executor_Local{}

	if m.GetResolutions() != nil {
		target.Resolutions = make(map[string]*Resolution, len(m.GetResolutions()))
		for k, v := range m.GetResolutions() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Resolutions[k] = h.Clone().(*Resolution)
			} else {
				target.Resolutions[k] = proto.Clone(v).(*Resolution)
			}

		}
	}

	target.EnableIntrospection = m.GetEnableIntrospection()

	return target
}
