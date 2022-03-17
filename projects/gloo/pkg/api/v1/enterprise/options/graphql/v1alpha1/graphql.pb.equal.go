// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/graphql/v1alpha1/graphql.proto

package v1alpha1

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
func (m *RequestTemplate) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RequestTemplate)
	if !ok {
		that2, ok := that.(RequestTemplate)
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

	if len(m.GetHeaders()) != len(target.GetHeaders()) {
		return false
	}
	for k, v := range m.GetHeaders() {

		if strings.Compare(v, target.GetHeaders()[k]) != 0 {
			return false
		}

	}

	if len(m.GetQueryParams()) != len(target.GetQueryParams()) {
		return false
	}
	for k, v := range m.GetQueryParams() {

		if strings.Compare(v, target.GetQueryParams()[k]) != 0 {
			return false
		}

	}

	if h, ok := interface{}(m.GetBody()).(equality.Equalizer); ok {
		if !h.Equal(target.GetBody()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetBody(), target.GetBody()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ResponseTemplate) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ResponseTemplate)
	if !ok {
		that2, ok := that.(ResponseTemplate)
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

	if strings.Compare(m.GetResultRoot(), target.GetResultRoot()) != 0 {
		return false
	}

	if len(m.GetSetters()) != len(target.GetSetters()) {
		return false
	}
	for k, v := range m.GetSetters() {

		if strings.Compare(v, target.GetSetters()[k]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *GrpcRequestTemplate) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GrpcRequestTemplate)
	if !ok {
		that2, ok := that.(GrpcRequestTemplate)
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

	if h, ok := interface{}(m.GetOutgoingMessageJson()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOutgoingMessageJson()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOutgoingMessageJson(), target.GetOutgoingMessageJson()) {
			return false
		}
	}

	if strings.Compare(m.GetServiceName(), target.GetServiceName()) != 0 {
		return false
	}

	if strings.Compare(m.GetMethodName(), target.GetMethodName()) != 0 {
		return false
	}

	if len(m.GetRequestMetadata()) != len(target.GetRequestMetadata()) {
		return false
	}
	for k, v := range m.GetRequestMetadata() {

		if strings.Compare(v, target.GetRequestMetadata()[k]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *RESTResolver) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RESTResolver)
	if !ok {
		that2, ok := that.(RESTResolver)
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

	if h, ok := interface{}(m.GetUpstreamRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetUpstreamRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetUpstreamRef(), target.GetUpstreamRef()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetRequest()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRequest()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRequest(), target.GetRequest()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetResponse()).(equality.Equalizer); ok {
		if !h.Equal(target.GetResponse()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetResponse(), target.GetResponse()) {
			return false
		}
	}

	if strings.Compare(m.GetSpanName(), target.GetSpanName()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *GrpcDescriptorRegistry) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GrpcDescriptorRegistry)
	if !ok {
		that2, ok := that.(GrpcDescriptorRegistry)
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

	switch m.DescriptorSet.(type) {

	case *GrpcDescriptorRegistry_ProtoDescriptor:
		if _, ok := target.DescriptorSet.(*GrpcDescriptorRegistry_ProtoDescriptor); !ok {
			return false
		}

		if strings.Compare(m.GetProtoDescriptor(), target.GetProtoDescriptor()) != 0 {
			return false
		}

	case *GrpcDescriptorRegistry_ProtoDescriptorBin:
		if _, ok := target.DescriptorSet.(*GrpcDescriptorRegistry_ProtoDescriptorBin); !ok {
			return false
		}

		if bytes.Compare(m.GetProtoDescriptorBin(), target.GetProtoDescriptorBin()) != 0 {
			return false
		}

	default:
		// m is nil but target is not nil
		if m.DescriptorSet != target.DescriptorSet {
			return false
		}
	}

	return true
}

// Equal function
func (m *GrpcResolver) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GrpcResolver)
	if !ok {
		that2, ok := that.(GrpcResolver)
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

	if h, ok := interface{}(m.GetUpstreamRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetUpstreamRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetUpstreamRef(), target.GetUpstreamRef()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetRequestTransform()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRequestTransform()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRequestTransform(), target.GetRequestTransform()) {
			return false
		}
	}

	if strings.Compare(m.GetSpanName(), target.GetSpanName()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *GatewaySchema) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GatewaySchema)
	if !ok {
		that2, ok := that.(GatewaySchema)
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

	if len(m.GetSubschemas()) != len(target.GetSubschemas()) {
		return false
	}
	for idx, v := range m.GetSubschemas() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSubschemas()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSubschemas()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *Resolution) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Resolution)
	if !ok {
		that2, ok := that.(Resolution)
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

	if h, ok := interface{}(m.GetStatPrefix()).(equality.Equalizer); ok {
		if !h.Equal(target.GetStatPrefix()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetStatPrefix(), target.GetStatPrefix()) {
			return false
		}
	}

	switch m.Resolver.(type) {

	case *Resolution_RestResolver:
		if _, ok := target.Resolver.(*Resolution_RestResolver); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRestResolver()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRestResolver()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRestResolver(), target.GetRestResolver()) {
				return false
			}
		}

	case *Resolution_GrpcResolver:
		if _, ok := target.Resolver.(*Resolution_GrpcResolver); !ok {
			return false
		}

		if h, ok := interface{}(m.GetGrpcResolver()).(equality.Equalizer); ok {
			if !h.Equal(target.GetGrpcResolver()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetGrpcResolver(), target.GetGrpcResolver()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Resolver != target.Resolver {
			return false
		}
	}

	return true
}

// Equal function
func (m *GraphQLApi) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GraphQLApi)
	if !ok {
		that2, ok := that.(GraphQLApi)
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

	if h, ok := interface{}(m.GetNamespacedStatuses()).(equality.Equalizer); ok {
		if !h.Equal(target.GetNamespacedStatuses()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetNamespacedStatuses(), target.GetNamespacedStatuses()) {
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

	if h, ok := interface{}(m.GetStatPrefix()).(equality.Equalizer); ok {
		if !h.Equal(target.GetStatPrefix()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetStatPrefix(), target.GetStatPrefix()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetPersistedQueryCacheConfig()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPersistedQueryCacheConfig()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPersistedQueryCacheConfig(), target.GetPersistedQueryCacheConfig()) {
			return false
		}
	}

	if len(m.GetAllowedQueryHashes()) != len(target.GetAllowedQueryHashes()) {
		return false
	}
	for idx, v := range m.GetAllowedQueryHashes() {

		if strings.Compare(v, target.GetAllowedQueryHashes()[idx]) != 0 {
			return false
		}

	}

	switch m.Schema.(type) {

	case *GraphQLApi_ExecutableSchema:
		if _, ok := target.Schema.(*GraphQLApi_ExecutableSchema); !ok {
			return false
		}

		if h, ok := interface{}(m.GetExecutableSchema()).(equality.Equalizer); ok {
			if !h.Equal(target.GetExecutableSchema()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetExecutableSchema(), target.GetExecutableSchema()) {
				return false
			}
		}

	case *GraphQLApi_GatewaySchema:
		if _, ok := target.Schema.(*GraphQLApi_GatewaySchema); !ok {
			return false
		}

		if h, ok := interface{}(m.GetGatewaySchema()).(equality.Equalizer); ok {
			if !h.Equal(target.GetGatewaySchema()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetGatewaySchema(), target.GetGatewaySchema()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Schema != target.Schema {
			return false
		}
	}

	return true
}

// Equal function
func (m *PersistedQueryCacheConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*PersistedQueryCacheConfig)
	if !ok {
		that2, ok := that.(PersistedQueryCacheConfig)
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

	if m.GetCacheSize() != target.GetCacheSize() {
		return false
	}

	return true
}

// Equal function
func (m *ExecutableSchema) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExecutableSchema)
	if !ok {
		that2, ok := that.(ExecutableSchema)
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

	if strings.Compare(m.GetSchemaDefinition(), target.GetSchemaDefinition()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetExecutor()).(equality.Equalizer); ok {
		if !h.Equal(target.GetExecutor()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetExecutor(), target.GetExecutor()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetGrpcDescriptorRegistry()).(equality.Equalizer); ok {
		if !h.Equal(target.GetGrpcDescriptorRegistry()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetGrpcDescriptorRegistry(), target.GetGrpcDescriptorRegistry()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *Executor) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Executor)
	if !ok {
		that2, ok := that.(Executor)
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

	switch m.Executor.(type) {

	case *Executor_Local_:
		if _, ok := target.Executor.(*Executor_Local_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetLocal()).(equality.Equalizer); ok {
			if !h.Equal(target.GetLocal()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetLocal(), target.GetLocal()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Executor != target.Executor {
			return false
		}
	}

	return true
}

// Equal function
func (m *GatewaySchema_SubschemaConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GatewaySchema_SubschemaConfig)
	if !ok {
		that2, ok := that.(GatewaySchema_SubschemaConfig)
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

	if strings.Compare(m.GetName(), target.GetName()) != 0 {
		return false
	}

	if strings.Compare(m.GetNamespace(), target.GetNamespace()) != 0 {
		return false
	}

	if len(m.GetTypeMerge()) != len(target.GetTypeMerge()) {
		return false
	}
	for k, v := range m.GetTypeMerge() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetTypeMerge()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetTypeMerge()[k]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *GatewaySchema_SubschemaConfig_TypeMergeConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GatewaySchema_SubschemaConfig_TypeMergeConfig)
	if !ok {
		that2, ok := that.(GatewaySchema_SubschemaConfig_TypeMergeConfig)
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

	if strings.Compare(m.GetSelectionSet(), target.GetSelectionSet()) != 0 {
		return false
	}

	if strings.Compare(m.GetQueryName(), target.GetQueryName()) != 0 {
		return false
	}

	if len(m.GetArgs()) != len(target.GetArgs()) {
		return false
	}
	for k, v := range m.GetArgs() {

		if strings.Compare(v, target.GetArgs()[k]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *Executor_Local) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Executor_Local)
	if !ok {
		that2, ok := that.(Executor_Local)
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

	if len(m.GetResolutions()) != len(target.GetResolutions()) {
		return false
	}
	for k, v := range m.GetResolutions() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetResolutions()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetResolutions()[k]) {
				return false
			}
		}

	}

	if m.GetEnableIntrospection() != target.GetEnableIntrospection() {
		return false
	}

	return true
}
