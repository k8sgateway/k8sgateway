// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/v1/options/faultinjection/fault.proto

package faultinjection

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
func (m *RouteAbort) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RouteAbort)
	if !ok {
		that2, ok := that.(RouteAbort)
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

	if m.GetPercentage() != target.GetPercentage() {
		return false
	}

	if m.GetHttpStatus() != target.GetHttpStatus() {
		return false
	}

	return true
}

// Equal function
func (m *RouteDelay) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RouteDelay)
	if !ok {
		that2, ok := that.(RouteDelay)
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

	if m.GetPercentage() != target.GetPercentage() {
		return false
	}

	if h, ok := interface{}(m.GetFixedDelay()).(equality.Equalizer); ok {
		if !h.Equal(target.GetFixedDelay()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetFixedDelay(), target.GetFixedDelay()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *RouteFaults) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RouteFaults)
	if !ok {
		that2, ok := that.(RouteFaults)
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

	if h, ok := interface{}(m.GetAbort()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAbort()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAbort(), target.GetAbort()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetDelay()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDelay()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDelay(), target.GetDelay()) {
			return false
		}
	}

	return true
}
