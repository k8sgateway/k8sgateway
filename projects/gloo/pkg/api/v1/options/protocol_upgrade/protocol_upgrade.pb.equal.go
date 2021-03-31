// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/protocol_upgrade/protocol_upgrade.proto

package protocol_upgrade

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
func (m *ProtocolUpgradeConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ProtocolUpgradeConfig)
	if !ok {
		that2, ok := that.(ProtocolUpgradeConfig)
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

	switch m.UpgradeType.(type) {

	case *ProtocolUpgradeConfig_Websocket:
		if _, ok := target.UpgradeType.(*ProtocolUpgradeConfig_Websocket); !ok {
			return false
		}

		if h, ok := interface{}(m.GetWebsocket()).(equality.Equalizer); ok {
			if !h.Equal(target.GetWebsocket()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetWebsocket(), target.GetWebsocket()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.UpgradeType != target.UpgradeType {
			return false
		}
	}

	return true
}

// Equal function
func (m *ProtocolUpgradeConfig_ProtocolUpgradeSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ProtocolUpgradeConfig_ProtocolUpgradeSpec)
	if !ok {
		that2, ok := that.(ProtocolUpgradeConfig_ProtocolUpgradeSpec)
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

	if h, ok := interface{}(m.GetEnabled()).(equality.Equalizer); ok {
		if !h.Equal(target.GetEnabled()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetEnabled(), target.GetEnabled()) {
			return false
		}
	}

	return true
}
