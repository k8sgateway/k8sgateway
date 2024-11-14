// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/v1/options/aws/ec2/aws_ec2.proto

package ec2

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
func (m *UpstreamSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*UpstreamSpec)
	if !ok {
		that2, ok := that.(UpstreamSpec)
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

	if strings.Compare(m.GetRegion(), target.GetRegion()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetSecretRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSecretRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSecretRef(), target.GetSecretRef()) {
			return false
		}
	}

	if strings.Compare(m.GetRoleArn(), target.GetRoleArn()) != 0 {
		return false
	}

	if len(m.GetFilters()) != len(target.GetFilters()) {
		return false
	}
	for idx, v := range m.GetFilters() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetFilters()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetFilters()[idx]) {
				return false
			}
		}

	}

	if m.GetPublicIp() != target.GetPublicIp() {
		return false
	}

	if m.GetPort() != target.GetPort() {
		return false
	}

	return true
}

// Equal function
func (m *TagFilter) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TagFilter)
	if !ok {
		that2, ok := that.(TagFilter)
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

	switch m.Spec.(type) {

	case *TagFilter_Key:
		if _, ok := target.Spec.(*TagFilter_Key); !ok {
			return false
		}

		if strings.Compare(m.GetKey(), target.GetKey()) != 0 {
			return false
		}

	case *TagFilter_KvPair_:
		if _, ok := target.Spec.(*TagFilter_KvPair_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetKvPair()).(equality.Equalizer); ok {
			if !h.Equal(target.GetKvPair()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetKvPair(), target.GetKvPair()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Spec != target.Spec {
			return false
		}
	}

	return true
}

// Equal function
func (m *TagFilter_KvPair) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TagFilter_KvPair)
	if !ok {
		that2, ok := that.(TagFilter_KvPair)
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

	if strings.Compare(m.GetValue(), target.GetValue()) != 0 {
		return false
	}

	return true
}
