// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/wasm/v3/wasm.proto

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
)

// Equal function
func (m *VmConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VmConfig)
	if !ok {
		that2, ok := that.(VmConfig)
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

	if strings.Compare(m.GetVmId(), target.GetVmId()) != 0 {
		return false
	}

	if strings.Compare(m.GetRuntime(), target.GetRuntime()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetCode()).(equality.Equalizer); ok {
		if !h.Equal(target.GetCode()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetCode(), target.GetCode()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetConfiguration()).(equality.Equalizer); ok {
		if !h.Equal(target.GetConfiguration()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetConfiguration(), target.GetConfiguration()) {
			return false
		}
	}

	if m.GetAllowPrecompiled() != target.GetAllowPrecompiled() {
		return false
	}

	if m.GetNackOnCodeCacheMiss() != target.GetNackOnCodeCacheMiss() {
		return false
	}

	return true
}

// Equal function
func (m *PluginConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*PluginConfig)
	if !ok {
		that2, ok := that.(PluginConfig)
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

	if strings.Compare(m.GetRootId(), target.GetRootId()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetConfiguration()).(equality.Equalizer); ok {
		if !h.Equal(target.GetConfiguration()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetConfiguration(), target.GetConfiguration()) {
			return false
		}
	}

	if m.GetFailOpen() != target.GetFailOpen() {
		return false
	}

	switch m.Vm.(type) {

	case *PluginConfig_VmConfig:

		if h, ok := interface{}(m.GetVmConfig()).(equality.Equalizer); ok {
			if !h.Equal(target.GetVmConfig()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetVmConfig(), target.GetVmConfig()) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *WasmService) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WasmService)
	if !ok {
		that2, ok := that.(WasmService)
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

	if m.GetSingleton() != target.GetSingleton() {
		return false
	}

	return true
}
