// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/graphql/graphql.proto

package graphql

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
func (m *ServiceSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ServiceSpec)
	if !ok {
		that2, ok := that.(ServiceSpec)
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

	if h, ok := interface{}(m.GetEndpoint()).(equality.Equalizer); ok {
		if !h.Equal(target.GetEndpoint()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetEndpoint(), target.GetEndpoint()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ServiceSpec_Endpoint) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ServiceSpec_Endpoint)
	if !ok {
		that2, ok := that.(ServiceSpec_Endpoint)
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

	if strings.Compare(m.GetUrl(), target.GetUrl()) != 0 {
		return false
	}

	return true
}
