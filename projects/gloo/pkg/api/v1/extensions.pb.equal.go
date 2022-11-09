// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/extensions.proto

package v1

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
func (m *Extensions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Extensions)
	if !ok {
		that2, ok := that.(Extensions)
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

	if len(m.GetConfigs()) != len(target.GetConfigs()) {
		return false
	}
	for k, v := range m.GetConfigs() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetConfigs()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetConfigs()[k]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *Extension) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Extension)
	if !ok {
		that2, ok := that.(Extension)
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

	if h, ok := interface{}(m.GetConfig()).(equality.Equalizer); ok {
		if !h.Equal(target.GetConfig()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetConfig(), target.GetConfig()) {
			return false
		}
	}

	return true
}
