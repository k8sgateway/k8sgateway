// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/dlp/dlp.proto

package dlp

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
func (m *FilterConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*FilterConfig)
	if !ok {
		that2, ok := that.(FilterConfig)
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

	if len(m.GetDlpRules()) != len(target.GetDlpRules()) {
		return false
	}
	for idx, v := range m.GetDlpRules() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetDlpRules()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetDlpRules()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *DlpRule) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*DlpRule)
	if !ok {
		that2, ok := that.(DlpRule)
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

	if h, ok := interface{}(m.GetMatcher()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMatcher()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMatcher(), target.GetMatcher()) {
			return false
		}
	}

	if len(m.GetActions()) != len(target.GetActions()) {
		return false
	}
	for idx, v := range m.GetActions() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetActions()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetActions()[idx]) {
				return false
			}
		}

	}

	if m.GetTransformAccessLogs() != target.GetTransformAccessLogs() {
		return false
	}

	return true
}

// Equal function
func (m *Config) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Config)
	if !ok {
		that2, ok := that.(Config)
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

	if len(m.GetActions()) != len(target.GetActions()) {
		return false
	}
	for idx, v := range m.GetActions() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetActions()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetActions()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *Action) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Action)
	if !ok {
		that2, ok := that.(Action)
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

	if m.GetActionType() != target.GetActionType() {
		return false
	}

	if h, ok := interface{}(m.GetCustomAction()).(equality.Equalizer); ok {
		if !h.Equal(target.GetCustomAction()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetCustomAction(), target.GetCustomAction()) {
			return false
		}
	}

	if m.GetShadow() != target.GetShadow() {
		return false
	}

	return true
}

// Equal function
func (m *CustomAction) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*CustomAction)
	if !ok {
		that2, ok := that.(CustomAction)
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

	if strings.Compare(m.GetName(), target.GetName()) != 0 {
		return false
	}

	if len(m.GetRegex()) != len(target.GetRegex()) {
		return false
	}
	for idx, v := range m.GetRegex() {

		if strings.Compare(v, target.GetRegex()[idx]) != 0 {
			return false
		}

	}

	if strings.Compare(m.GetMaskChar(), target.GetMaskChar()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetPercent()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPercent()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPercent(), target.GetPercent()) {
			return false
		}
	}

	return true
}
