// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/type/matcher/v3/regex.proto

package v3

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
func (m *RegexMatcher) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RegexMatcher)
	if !ok {
		that2, ok := that.(RegexMatcher)
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

	if strings.Compare(m.GetRegex(), target.GetRegex()) != 0 {
		return false
	}

	switch m.EngineType.(type) {

	case *RegexMatcher_GoogleRe2:
		if _, ok := target.EngineType.(*RegexMatcher_GoogleRe2); !ok {
			return false
		}

		if h, ok := interface{}(m.GetGoogleRe2()).(equality.Equalizer); ok {
			if !h.Equal(target.GetGoogleRe2()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetGoogleRe2(), target.GetGoogleRe2()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.EngineType != target.EngineType {
			return false
		}
	}

	return true
}

// Equal function
func (m *RegexMatchAndSubstitute) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RegexMatchAndSubstitute)
	if !ok {
		that2, ok := that.(RegexMatchAndSubstitute)
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

	if h, ok := interface{}(m.GetPattern()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPattern()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPattern(), target.GetPattern()) {
			return false
		}
	}

	if strings.Compare(m.GetSubstitution(), target.GetSubstitution()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *RegexMatcher_GoogleRE2) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RegexMatcher_GoogleRE2)
	if !ok {
		that2, ok := that.(RegexMatcher_GoogleRE2)
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

	if h, ok := interface{}(m.GetMaxProgramSize()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMaxProgramSize()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMaxProgramSize(), target.GetMaxProgramSize()) {
			return false
		}
	}

	return true
}
