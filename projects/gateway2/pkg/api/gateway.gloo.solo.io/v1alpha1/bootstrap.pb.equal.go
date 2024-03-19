// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gateway2/api/v1alpha1/bootstrap.proto

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
func (m *EnvoyBootstrap) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*EnvoyBootstrap)
	if !ok {
		that2, ok := that.(EnvoyBootstrap)
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

	if h, ok := interface{}(m.GetXdsServer()).(equality.Equalizer); ok {
		if !h.Equal(target.GetXdsServer()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetXdsServer(), target.GetXdsServer()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *XdsServer) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*XdsServer)
	if !ok {
		that2, ok := that.(XdsServer)
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

	if h, ok := interface{}(m.GetAddress()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAddress()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAddress(), target.GetAddress()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetPortValue()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPortValue()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPortValue(), target.GetPortValue()) {
			return false
		}
	}

	return true
}
