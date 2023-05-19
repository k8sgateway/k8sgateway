// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/waf/waf.proto

package waf

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
func (m *Settings) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Settings)
	if !ok {
		that2, ok := that.(Settings)
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

	if m.GetDisabled() != target.GetDisabled() {
		return false
	}

	if strings.Compare(m.GetCustomInterventionMessage(), target.GetCustomInterventionMessage()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetCoreRuleSet()).(equality.Equalizer); ok {
		if !h.Equal(target.GetCoreRuleSet()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetCoreRuleSet(), target.GetCoreRuleSet()) {
			return false
		}
	}

	if len(m.GetRuleSets()) != len(target.GetRuleSets()) {
		return false
	}
	for idx, v := range m.GetRuleSets() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetRuleSets()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetRuleSets()[idx]) {
				return false
			}
		}

	}

	if len(m.GetConfigMapRuleSets()) != len(target.GetConfigMapRuleSets()) {
		return false
	}
	for idx, v := range m.GetConfigMapRuleSets() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetConfigMapRuleSets()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetConfigMapRuleSets()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetAuditLogging()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAuditLogging()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAuditLogging(), target.GetAuditLogging()) {
			return false
		}
	}

	if m.GetRequestHeadersOnly() != target.GetRequestHeadersOnly() {
		return false
	}

	if m.GetResponseHeadersOnly() != target.GetResponseHeadersOnly() {
		return false
	}

	return true
}

// Equal function
func (m *RuleSetFromConfigMap) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RuleSetFromConfigMap)
	if !ok {
		that2, ok := that.(RuleSetFromConfigMap)
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

	if h, ok := interface{}(m.GetConfigMapRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetConfigMapRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetConfigMapRef(), target.GetConfigMapRef()) {
			return false
		}
	}

	if len(m.GetDataMapKeys()) != len(target.GetDataMapKeys()) {
		return false
	}
	for idx, v := range m.GetDataMapKeys() {

		if strings.Compare(v, target.GetDataMapKeys()[idx]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *CoreRuleSet) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*CoreRuleSet)
	if !ok {
		that2, ok := that.(CoreRuleSet)
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

	switch m.CustomSettingsType.(type) {

	case *CoreRuleSet_CustomSettingsString:
		if _, ok := target.CustomSettingsType.(*CoreRuleSet_CustomSettingsString); !ok {
			return false
		}

		if strings.Compare(m.GetCustomSettingsString(), target.GetCustomSettingsString()) != 0 {
			return false
		}

	case *CoreRuleSet_CustomSettingsFile:
		if _, ok := target.CustomSettingsType.(*CoreRuleSet_CustomSettingsFile); !ok {
			return false
		}

		if strings.Compare(m.GetCustomSettingsFile(), target.GetCustomSettingsFile()) != 0 {
			return false
		}

	default:
		// m is nil but target is not nil
		if m.CustomSettingsType != target.CustomSettingsType {
			return false
		}
	}

	return true
}
