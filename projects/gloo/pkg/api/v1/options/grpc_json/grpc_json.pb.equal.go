// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/grpc_json/grpc_json.proto

package grpc_json

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
func (m *GrpcJsonTranscoder) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GrpcJsonTranscoder)
	if !ok {
		that2, ok := that.(GrpcJsonTranscoder)
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

	if len(m.GetServices()) != len(target.GetServices()) {
		return false
	}
	for idx, v := range m.GetServices() {

		if strings.Compare(v, target.GetServices()[idx]) != 0 {
			return false
		}

	}

	if h, ok := interface{}(m.GetPrintOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPrintOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPrintOptions(), target.GetPrintOptions()) {
			return false
		}
	}

	if m.GetMatchIncomingRequestRoute() != target.GetMatchIncomingRequestRoute() {
		return false
	}

	if len(m.GetIgnoredQueryParameters()) != len(target.GetIgnoredQueryParameters()) {
		return false
	}
	for idx, v := range m.GetIgnoredQueryParameters() {

		if strings.Compare(v, target.GetIgnoredQueryParameters()[idx]) != 0 {
			return false
		}

	}

	if m.GetAutoMapping() != target.GetAutoMapping() {
		return false
	}

	if m.GetIgnoreUnknownQueryParameters() != target.GetIgnoreUnknownQueryParameters() {
		return false
	}

	if m.GetConvertGrpcStatus() != target.GetConvertGrpcStatus() {
		return false
	}

	if len(m.GetMethodMap()) != len(target.GetMethodMap()) {
		return false
	}
	for k, v := range m.GetMethodMap() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetMethodMap()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetMethodMap()[k]) {
				return false
			}
		}

	}

	switch m.DescriptorSet.(type) {

	case *GrpcJsonTranscoder_ProtoDescriptor:
		if _, ok := target.DescriptorSet.(*GrpcJsonTranscoder_ProtoDescriptor); !ok {
			return false
		}

		if strings.Compare(m.GetProtoDescriptor(), target.GetProtoDescriptor()) != 0 {
			return false
		}

	case *GrpcJsonTranscoder_ProtoDescriptorBin:
		if _, ok := target.DescriptorSet.(*GrpcJsonTranscoder_ProtoDescriptorBin); !ok {
			return false
		}

		if bytes.Compare(m.GetProtoDescriptorBin(), target.GetProtoDescriptorBin()) != 0 {
			return false
		}

	case *GrpcJsonTranscoder_ProtoDescriptorConfigMap:
		if _, ok := target.DescriptorSet.(*GrpcJsonTranscoder_ProtoDescriptorConfigMap); !ok {
			return false
		}

		if h, ok := interface{}(m.GetProtoDescriptorConfigMap()).(equality.Equalizer); ok {
			if !h.Equal(target.GetProtoDescriptorConfigMap()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetProtoDescriptorConfigMap(), target.GetProtoDescriptorConfigMap()) {
				return false
			}
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
func (m *GrpcJsonTranscoder_PrintOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GrpcJsonTranscoder_PrintOptions)
	if !ok {
		that2, ok := that.(GrpcJsonTranscoder_PrintOptions)
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

	if m.GetAddWhitespace() != target.GetAddWhitespace() {
		return false
	}

	if m.GetAlwaysPrintPrimitiveFields() != target.GetAlwaysPrintPrimitiveFields() {
		return false
	}

	if m.GetAlwaysPrintEnumsAsInts() != target.GetAlwaysPrintEnumsAsInts() {
		return false
	}

	if m.GetPreserveProtoFieldNames() != target.GetPreserveProtoFieldNames() {
		return false
	}

	return true
}

// Equal function
func (m *GrpcJsonTranscoder_DescriptorConfigMap) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GrpcJsonTranscoder_DescriptorConfigMap)
	if !ok {
		that2, ok := that.(GrpcJsonTranscoder_DescriptorConfigMap)
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

	if h, ok := interface{}(m.GetConfigMapRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetConfigMapRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetConfigMapRef(), target.GetConfigMapRef()) {
			return false
		}
	}

	if strings.Compare(m.GetKey(), target.GetKey()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *GrpcJsonTranscoderMethodList) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GrpcJsonTranscoderMethodList)
	if !ok {
		that2, ok := that.(GrpcJsonTranscoderMethodList)
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

	if len(m.GetMethods()) != len(target.GetMethods()) {
		return false
	}
	for idx, v := range m.GetMethods() {

		if strings.Compare(v, target.GetMethods()[idx]) != 0 {
			return false
		}

	}

	return true
}
