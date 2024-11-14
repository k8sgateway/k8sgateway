// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/v1/options/subset_spec.proto

package options

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = clone.Cloner(nil)
	_ = proto.Message(nil)
)

// Clone function
func (m *SubsetSpec) Clone() proto.Message {
	var target *SubsetSpec
	if m == nil {
		return target
	}
	target = &SubsetSpec{}

	if m.GetSelectors() != nil {
		target.Selectors = make([]*Selector, len(m.GetSelectors()))
		for idx, v := range m.GetSelectors() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Selectors[idx] = h.Clone().(*Selector)
			} else {
				target.Selectors[idx] = proto.Clone(v).(*Selector)
			}

		}
	}

	target.FallbackPolicy = m.GetFallbackPolicy()

	if h, ok := interface{}(m.GetDefaultSubset()).(clone.Cloner); ok {
		target.DefaultSubset = h.Clone().(*Subset)
	} else {
		target.DefaultSubset = proto.Clone(m.GetDefaultSubset()).(*Subset)
	}

	return target
}

// Clone function
func (m *Selector) Clone() proto.Message {
	var target *Selector
	if m == nil {
		return target
	}
	target = &Selector{}

	if m.GetKeys() != nil {
		target.Keys = make([]string, len(m.GetKeys()))
		for idx, v := range m.GetKeys() {

			target.Keys[idx] = v

		}
	}

	target.SingleHostPerSubset = m.GetSingleHostPerSubset()

	return target
}

// Clone function
func (m *Subset) Clone() proto.Message {
	var target *Subset
	if m == nil {
		return target
	}
	target = &Subset{}

	if m.GetValues() != nil {
		target.Values = make(map[string]string, len(m.GetValues()))
		for k, v := range m.GetValues() {

			target.Values[k] = v

		}
	}

	return target
}
