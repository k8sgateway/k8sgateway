// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/transformation/transformation.proto

package transformation

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
func (m *ResponseMatch) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ResponseMatch)
	if !ok {
		that2, ok := that.(ResponseMatch)
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

	if len(m.GetMatchers()) != len(target.GetMatchers()) {
		return false
	}
	for idx, v := range m.GetMatchers() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetMatchers()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetMatchers()[idx]) {
				return false
			}
		}

	}

	if strings.Compare(m.GetResponseCodeDetails(), target.GetResponseCodeDetails()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetResponseTransformation()).(equality.Equalizer); ok {
		if !h.Equal(target.GetResponseTransformation()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetResponseTransformation(), target.GetResponseTransformation()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *RequestMatch) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RequestMatch)
	if !ok {
		that2, ok := that.(RequestMatch)
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

	if m.GetClearRouteCache() != target.GetClearRouteCache() {
		return false
	}

	if h, ok := interface{}(m.GetRequestTransformation()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRequestTransformation()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRequestTransformation(), target.GetRequestTransformation()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetResponseTransformation()).(equality.Equalizer); ok {
		if !h.Equal(target.GetResponseTransformation()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetResponseTransformation(), target.GetResponseTransformation()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *Transformations) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Transformations)
	if !ok {
		that2, ok := that.(Transformations)
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

	if h, ok := interface{}(m.GetRequestTransformation()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRequestTransformation()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRequestTransformation(), target.GetRequestTransformation()) {
			return false
		}
	}

	if m.GetClearRouteCache() != target.GetClearRouteCache() {
		return false
	}

	if h, ok := interface{}(m.GetResponseTransformation()).(equality.Equalizer); ok {
		if !h.Equal(target.GetResponseTransformation()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetResponseTransformation(), target.GetResponseTransformation()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *RequestResponseTransformations) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RequestResponseTransformations)
	if !ok {
		that2, ok := that.(RequestResponseTransformations)
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

	if len(m.GetRequestTransforms()) != len(target.GetRequestTransforms()) {
		return false
	}
	for idx, v := range m.GetRequestTransforms() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetRequestTransforms()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetRequestTransforms()[idx]) {
				return false
			}
		}

	}

	if len(m.GetResponseTransforms()) != len(target.GetResponseTransforms()) {
		return false
	}
	for idx, v := range m.GetResponseTransforms() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetResponseTransforms()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetResponseTransforms()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *TransformationStages) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TransformationStages)
	if !ok {
		that2, ok := that.(TransformationStages)
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

	if h, ok := interface{}(m.GetEarly()).(equality.Equalizer); ok {
		if !h.Equal(target.GetEarly()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetEarly(), target.GetEarly()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetRegular()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRegular()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRegular(), target.GetRegular()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetPostRouting()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPostRouting()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPostRouting(), target.GetPostRouting()) {
			return false
		}
	}

	if m.GetInheritTransformation() != target.GetInheritTransformation() {
		return false
	}

	if h, ok := interface{}(m.GetLogRequestResponseInfo()).(equality.Equalizer); ok {
		if !h.Equal(target.GetLogRequestResponseInfo()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetLogRequestResponseInfo(), target.GetLogRequestResponseInfo()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetEscapeCharacters()).(equality.Equalizer); ok {
		if !h.Equal(target.GetEscapeCharacters()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetEscapeCharacters(), target.GetEscapeCharacters()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *Transformation) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Transformation)
	if !ok {
		that2, ok := that.(Transformation)
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

	if m.GetLogRequestResponseInfo() != target.GetLogRequestResponseInfo() {
		return false
	}

	switch m.TransformationType.(type) {

	case *Transformation_TransformationTemplate:
		if _, ok := target.TransformationType.(*Transformation_TransformationTemplate); !ok {
			return false
		}

		if h, ok := interface{}(m.GetTransformationTemplate()).(equality.Equalizer); ok {
			if !h.Equal(target.GetTransformationTemplate()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetTransformationTemplate(), target.GetTransformationTemplate()) {
				return false
			}
		}

	case *Transformation_HeaderBodyTransform:
		if _, ok := target.TransformationType.(*Transformation_HeaderBodyTransform); !ok {
			return false
		}

		if h, ok := interface{}(m.GetHeaderBodyTransform()).(equality.Equalizer); ok {
			if !h.Equal(target.GetHeaderBodyTransform()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetHeaderBodyTransform(), target.GetHeaderBodyTransform()) {
				return false
			}
		}

	case *Transformation_XsltTransformation:
		if _, ok := target.TransformationType.(*Transformation_XsltTransformation); !ok {
			return false
		}

		if h, ok := interface{}(m.GetXsltTransformation()).(equality.Equalizer); ok {
			if !h.Equal(target.GetXsltTransformation()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetXsltTransformation(), target.GetXsltTransformation()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.TransformationType != target.TransformationType {
			return false
		}
	}

	return true
}

// Equal function
func (m *Extraction) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Extraction)
	if !ok {
		that2, ok := that.(Extraction)
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

	if strings.Compare(m.GetRegex(), target.GetRegex()) != 0 {
		return false
	}

	if m.GetSubgroup() != target.GetSubgroup() {
		return false
	}

	if h, ok := interface{}(m.GetReplacementText()).(equality.Equalizer); ok {
		if !h.Equal(target.GetReplacementText()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetReplacementText(), target.GetReplacementText()) {
			return false
		}
	}

	if m.GetMode() != target.GetMode() {
		return false
	}

	switch m.Source.(type) {

	case *Extraction_Header:
		if _, ok := target.Source.(*Extraction_Header); !ok {
			return false
		}

		if strings.Compare(m.GetHeader(), target.GetHeader()) != 0 {
			return false
		}

	case *Extraction_Body:
		if _, ok := target.Source.(*Extraction_Body); !ok {
			return false
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

	default:
		// m is nil but target is not nil
		if m.Source != target.Source {
			return false
		}
	}

	return true
}

// Equal function
func (m *TransformationTemplate) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TransformationTemplate)
	if !ok {
		that2, ok := that.(TransformationTemplate)
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

	if m.GetAdvancedTemplates() != target.GetAdvancedTemplates() {
		return false
	}

	if len(m.GetExtractors()) != len(target.GetExtractors()) {
		return false
	}
	for k, v := range m.GetExtractors() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetExtractors()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetExtractors()[k]) {
				return false
			}
		}

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

	if len(m.GetHeadersToAppend()) != len(target.GetHeadersToAppend()) {
		return false
	}
	for idx, v := range m.GetHeadersToAppend() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetHeadersToAppend()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetHeadersToAppend()[idx]) {
				return false
			}
		}

	}

	if len(m.GetHeadersToRemove()) != len(target.GetHeadersToRemove()) {
		return false
	}
	for idx, v := range m.GetHeadersToRemove() {

		if strings.Compare(v, target.GetHeadersToRemove()[idx]) != 0 {
			return false
		}

	}

	if m.GetParseBodyBehavior() != target.GetParseBodyBehavior() {
		return false
	}

	if m.GetIgnoreErrorOnParse() != target.GetIgnoreErrorOnParse() {
		return false
	}

	if len(m.GetDynamicMetadataValues()) != len(target.GetDynamicMetadataValues()) {
		return false
	}
	for idx, v := range m.GetDynamicMetadataValues() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetDynamicMetadataValues()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetDynamicMetadataValues()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetEscapeCharacters()).(equality.Equalizer); ok {
		if !h.Equal(target.GetEscapeCharacters()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetEscapeCharacters(), target.GetEscapeCharacters()) {
			return false
		}
	}

	switch m.BodyTransformation.(type) {

	case *TransformationTemplate_Body:
		if _, ok := target.BodyTransformation.(*TransformationTemplate_Body); !ok {
			return false
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

	case *TransformationTemplate_Passthrough:
		if _, ok := target.BodyTransformation.(*TransformationTemplate_Passthrough); !ok {
			return false
		}

		if h, ok := interface{}(m.GetPassthrough()).(equality.Equalizer); ok {
			if !h.Equal(target.GetPassthrough()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetPassthrough(), target.GetPassthrough()) {
				return false
			}
		}

	case *TransformationTemplate_MergeExtractorsToBody:
		if _, ok := target.BodyTransformation.(*TransformationTemplate_MergeExtractorsToBody); !ok {
			return false
		}

		if h, ok := interface{}(m.GetMergeExtractorsToBody()).(equality.Equalizer); ok {
			if !h.Equal(target.GetMergeExtractorsToBody()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetMergeExtractorsToBody(), target.GetMergeExtractorsToBody()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.BodyTransformation != target.BodyTransformation {
			return false
		}
	}

	return true
}

// Equal function
func (m *InjaTemplate) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*InjaTemplate)
	if !ok {
		that2, ok := that.(InjaTemplate)
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

	if strings.Compare(m.GetText(), target.GetText()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *Passthrough) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Passthrough)
	if !ok {
		that2, ok := that.(Passthrough)
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
func (m *MergeExtractorsToBody) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*MergeExtractorsToBody)
	if !ok {
		that2, ok := that.(MergeExtractorsToBody)
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
func (m *HeaderBodyTransform) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HeaderBodyTransform)
	if !ok {
		that2, ok := that.(HeaderBodyTransform)
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

	if m.GetAddRequestMetadata() != target.GetAddRequestMetadata() {
		return false
	}

	return true
}

// Equal function
func (m *TransformationTemplate_HeaderToAppend) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TransformationTemplate_HeaderToAppend)
	if !ok {
		that2, ok := that.(TransformationTemplate_HeaderToAppend)
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
func (m *TransformationTemplate_DynamicMetadataValue) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TransformationTemplate_DynamicMetadataValue)
	if !ok {
		that2, ok := that.(TransformationTemplate_DynamicMetadataValue)
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

	if strings.Compare(m.GetMetadataNamespace(), target.GetMetadataNamespace()) != 0 {
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
