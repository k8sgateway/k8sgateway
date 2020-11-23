// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/ratelimit.proto

package enterprise

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
)

// Equal function
func (m *RateLimitConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RateLimitConfig)
	if !ok {
		that2, ok := that.(RateLimitConfig)
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

	if strings.Compare(m.GetDomain(), target.GetDomain()) != 0 {
		return false
	}

	if len(m.GetDescriptors()) != len(target.GetDescriptors()) {
		return false
	}
	for idx, v := range m.GetDescriptors() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetDescriptors()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetDescriptors()[idx]) {
				return false
			}
		}

	}

	if len(m.GetSetDescriptors()) != len(target.GetSetDescriptors()) {
		return false
	}
	for idx, v := range m.GetSetDescriptors() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSetDescriptors()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSetDescriptors()[idx]) {
				return false
			}
		}

	}

	return true
}
