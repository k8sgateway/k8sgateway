// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/filters/http/ext_proc/v3/processing_mode.proto

package v3

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
func (m *ProcessingMode) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ProcessingMode)
	if !ok {
		that2, ok := that.(ProcessingMode)
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

	if m.GetRequestHeaderMode() != target.GetRequestHeaderMode() {
		return false
	}

	if m.GetResponseHeaderMode() != target.GetResponseHeaderMode() {
		return false
	}

	if m.GetRequestBodyMode() != target.GetRequestBodyMode() {
		return false
	}

	if m.GetResponseBodyMode() != target.GetResponseBodyMode() {
		return false
	}

	if m.GetRequestTrailerMode() != target.GetRequestTrailerMode() {
		return false
	}

	if m.GetResponseTrailerMode() != target.GetResponseTrailerMode() {
		return false
	}

	return true
}
