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
func (m *PathSegment) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*PathSegment)
	if !ok {
		that2, ok := that.(PathSegment)
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

	switch m.Segment.(type) {

	case *PathSegment_Key:
		if _, ok := target.Segment.(*PathSegment_Key); !ok {
			return false
		}

		if strings.Compare(m.GetKey(), target.GetKey()) != 0 {
			return false
		}

	case *PathSegment_Index:
		if _, ok := target.Segment.(*PathSegment_Index); !ok {
			return false
		}

		if m.GetIndex() != target.GetIndex() {
			return false
		}

	default:
		// m is nil but target is not nil
		if m.Segment != target.Segment {
			return false
		}
	}

	return true
}

// Equal function
func (m *ValueProvider) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ValueProvider)
	if !ok {
		that2, ok := that.(ValueProvider)
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

	switch m.Provider.(type) {

	case *ValueProvider_GraphqlArg:
		if _, ok := target.Provider.(*ValueProvider_GraphqlArg); !ok {
			return false
		}

		if h, ok := interface{}(m.GetGraphqlArg()).(equality.Equalizer); ok {
			if !h.Equal(target.GetGraphqlArg()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetGraphqlArg(), target.GetGraphqlArg()) {
				return false
			}
		}

	case *ValueProvider_TypedProvider:
		if _, ok := target.Provider.(*ValueProvider_TypedProvider); !ok {
			return false
		}

		if h, ok := interface{}(m.GetTypedProvider()).(equality.Equalizer); ok {
			if !h.Equal(target.GetTypedProvider()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetTypedProvider(), target.GetTypedProvider()) {
				return false
			}
		}

	case *ValueProvider_GraphqlParent:
		if _, ok := target.Provider.(*ValueProvider_GraphqlParent); !ok {
			return false
		}

		if h, ok := interface{}(m.GetGraphqlParent()).(equality.Equalizer); ok {
			if !h.Equal(target.GetGraphqlParent()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetGraphqlParent(), target.GetGraphqlParent()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Provider != target.Provider {
			return false
		}
	}

	return true
}

// Equal function
func (m *JsonKeyValue) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*JsonKeyValue)
	if !ok {
		that2, ok := that.(JsonKeyValue)
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

	if h, ok := interface{}(m.GetValue()).(equality.Equalizer); ok {
		if !h.Equal(target.GetValue()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetValue(), target.GetValue()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *JsonNode) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*JsonNode)
	if !ok {
		that2, ok := that.(JsonNode)
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

	if len(m.GetKeyValues()) != len(target.GetKeyValues()) {
		return false
	}
	for idx, v := range m.GetKeyValues() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetKeyValues()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetKeyValues()[idx]) {
				return false
			}
		}

	}

	return true
}

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

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetHeaders()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetHeaders()[k]) {
				return false
			}
		}

	}

	if len(m.GetQueryParams()) != len(target.GetQueryParams()) {
		return false
	}
	for k, v := range m.GetQueryParams() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetQueryParams()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetQueryParams()[k]) {
				return false
			}
		}

	}

	switch m.OutgoingBody.(type) {

	case *RequestTemplate_Json:
		if _, ok := target.OutgoingBody.(*RequestTemplate_Json); !ok {
			return false
		}

		if h, ok := interface{}(m.GetJson()).(equality.Equalizer); ok {
			if !h.Equal(target.GetJson()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetJson(), target.GetJson()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.OutgoingBody != target.OutgoingBody {
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
func (m *QueryMatcher) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*QueryMatcher)
	if !ok {
		that2, ok := that.(QueryMatcher)
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

	switch m.Match.(type) {

	case *QueryMatcher_FieldMatcher_:
		if _, ok := target.Match.(*QueryMatcher_FieldMatcher_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetFieldMatcher()).(equality.Equalizer); ok {
			if !h.Equal(target.GetFieldMatcher()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetFieldMatcher(), target.GetFieldMatcher()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Match != target.Match {
			return false
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

	if h, ok := interface{}(m.GetMatcher()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMatcher()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMatcher(), target.GetMatcher()) {
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

	default:
		// m is nil but target is not nil
		if m.Resolver != target.Resolver {
			return false
		}
	}

	return true
}

// Equal function
func (m *GraphQLSchema) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GraphQLSchema)
	if !ok {
		that2, ok := that.(GraphQLSchema)
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

	if strings.Compare(m.GetSchema(), target.GetSchema()) != 0 {
		return false
	}

	if m.GetEnableIntrospection() != target.GetEnableIntrospection() {
		return false
	}

	if len(m.GetResolutions()) != len(target.GetResolutions()) {
		return false
	}
	for idx, v := range m.GetResolutions() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetResolutions()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetResolutions()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *ValueProvider_GraphQLArgExtraction) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ValueProvider_GraphQLArgExtraction)
	if !ok {
		that2, ok := that.(ValueProvider_GraphQLArgExtraction)
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

	if strings.Compare(m.GetArgName(), target.GetArgName()) != 0 {
		return false
	}

	if len(m.GetPath()) != len(target.GetPath()) {
		return false
	}
	for idx, v := range m.GetPath() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetPath()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetPath()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *ValueProvider_GraphQLParentExtraction) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ValueProvider_GraphQLParentExtraction)
	if !ok {
		that2, ok := that.(ValueProvider_GraphQLParentExtraction)
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

	if len(m.GetPath()) != len(target.GetPath()) {
		return false
	}
	for idx, v := range m.GetPath() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetPath()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetPath()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *ValueProvider_TypedValueProvider) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ValueProvider_TypedValueProvider)
	if !ok {
		that2, ok := that.(ValueProvider_TypedValueProvider)
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

	switch m.ValProvider.(type) {

	case *ValueProvider_TypedValueProvider_Header:
		if _, ok := target.ValProvider.(*ValueProvider_TypedValueProvider_Header); !ok {
			return false
		}

		if strings.Compare(m.GetHeader(), target.GetHeader()) != 0 {
			return false
		}

	case *ValueProvider_TypedValueProvider_Value:
		if _, ok := target.ValProvider.(*ValueProvider_TypedValueProvider_Value); !ok {
			return false
		}

		if strings.Compare(m.GetValue(), target.GetValue()) != 0 {
			return false
		}

	default:
		// m is nil but target is not nil
		if m.ValProvider != target.ValProvider {
			return false
		}
	}

	return true
}

// Equal function
func (m *JsonKeyValue_JsonValueList) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*JsonKeyValue_JsonValueList)
	if !ok {
		that2, ok := that.(JsonKeyValue_JsonValueList)
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

	if len(m.GetValues()) != len(target.GetValues()) {
		return false
	}
	for idx, v := range m.GetValues() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetValues()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetValues()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *JsonKeyValue_JsonValue) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*JsonKeyValue_JsonValue)
	if !ok {
		that2, ok := that.(JsonKeyValue_JsonValue)
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

	switch m.JsonVal.(type) {

	case *JsonKeyValue_JsonValue_Node:
		if _, ok := target.JsonVal.(*JsonKeyValue_JsonValue_Node); !ok {
			return false
		}

		if h, ok := interface{}(m.GetNode()).(equality.Equalizer); ok {
			if !h.Equal(target.GetNode()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetNode(), target.GetNode()) {
				return false
			}
		}

	case *JsonKeyValue_JsonValue_ValueProvider:
		if _, ok := target.JsonVal.(*JsonKeyValue_JsonValue_ValueProvider); !ok {
			return false
		}

		if h, ok := interface{}(m.GetValueProvider()).(equality.Equalizer); ok {
			if !h.Equal(target.GetValueProvider()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetValueProvider(), target.GetValueProvider()) {
				return false
			}
		}

	case *JsonKeyValue_JsonValue_List:
		if _, ok := target.JsonVal.(*JsonKeyValue_JsonValue_List); !ok {
			return false
		}

		if h, ok := interface{}(m.GetList()).(equality.Equalizer); ok {
			if !h.Equal(target.GetList()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetList(), target.GetList()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.JsonVal != target.JsonVal {
			return false
		}
	}

	return true
}

// Equal function
func (m *QueryMatcher_FieldMatcher) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*QueryMatcher_FieldMatcher)
	if !ok {
		that2, ok := that.(QueryMatcher_FieldMatcher)
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

	if strings.Compare(m.GetType(), target.GetType()) != 0 {
		return false
	}

	if strings.Compare(m.GetField(), target.GetField()) != 0 {
		return false
	}

	return true
}
