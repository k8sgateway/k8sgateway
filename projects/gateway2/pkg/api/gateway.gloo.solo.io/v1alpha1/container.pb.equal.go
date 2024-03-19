// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gateway2/api/v1alpha1/container.proto

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
func (m *Image) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Image)
	if !ok {
		that2, ok := that.(Image)
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

	if strings.Compare(m.GetRegistry(), target.GetRegistry()) != 0 {
		return false
	}

	if strings.Compare(m.GetRepository(), target.GetRepository()) != 0 {
		return false
	}

	if strings.Compare(m.GetTag(), target.GetTag()) != 0 {
		return false
	}

	if strings.Compare(m.GetPullPolicy(), target.GetPullPolicy()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *ResourceRequirements) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ResourceRequirements)
	if !ok {
		that2, ok := that.(ResourceRequirements)
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

	if len(m.GetLimits()) != len(target.GetLimits()) {
		return false
	}
	for k, v := range m.GetLimits() {

		if strings.Compare(v, target.GetLimits()[k]) != 0 {
			return false
		}

	}

	if len(m.GetRequests()) != len(target.GetRequests()) {
		return false
	}
	for k, v := range m.GetRequests() {

		if strings.Compare(v, target.GetRequests()[k]) != 0 {
			return false
		}

	}

	return true
}
