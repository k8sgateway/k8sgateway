// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/external/envoy/extensions/filters/http/buffer/v3/buffer.proto

package v3

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

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
func (m *Buffer) Clone() proto.Message {
	var target *Buffer
	if m == nil {
		return target
	}
	target = &Buffer{}

	if h, ok := interface{}(m.GetMaxRequestBytes()).(clone.Cloner); ok {
		target.MaxRequestBytes = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.MaxRequestBytes = proto.Clone(m.GetMaxRequestBytes()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	return target
}

// Clone function
func (m *BufferPerRoute) Clone() proto.Message {
	var target *BufferPerRoute
	if m == nil {
		return target
	}
	target = &BufferPerRoute{}

	switch m.Override.(type) {

	case *BufferPerRoute_Disabled:

		target.Override = &BufferPerRoute_Disabled{
			Disabled: m.GetDisabled(),
		}

	case *BufferPerRoute_Buffer:

		if h, ok := interface{}(m.GetBuffer()).(clone.Cloner); ok {
			target.Override = &BufferPerRoute_Buffer{
				Buffer: h.Clone().(*Buffer),
			}
		} else {
			target.Override = &BufferPerRoute_Buffer{
				Buffer: proto.Clone(m.GetBuffer()).(*Buffer),
			}
		}

	}

	return target
}
