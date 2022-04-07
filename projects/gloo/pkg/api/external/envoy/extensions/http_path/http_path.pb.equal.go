// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/http_path/http_path.proto

package http_path

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
func (m *HttpPath) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HttpPath)
	if !ok {
		that2, ok := that.(HttpPath)
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

	if h, ok := interface{}(m.GetHttpHealthCheck()).(equality.Equalizer); ok {
		if !h.Equal(target.GetHttpHealthCheck()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetHttpHealthCheck(), target.GetHttpHealthCheck()) {
			return false
		}
	}

	return true
}
