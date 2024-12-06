// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/external/envoy/config/core/v3/socket_option.proto

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
func (m *SocketOption) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*SocketOption)
	if !ok {
		that2, ok := that.(SocketOption)
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

	if strings.Compare(m.GetDescription(), target.GetDescription()) != 0 {
		return false
	}

	if m.GetLevel() != target.GetLevel() {
		return false
	}

	if m.GetName() != target.GetName() {
		return false
	}

	if m.GetState() != target.GetState() {
		return false
	}

	switch m.Value.(type) {

	case *SocketOption_IntValue:
		if _, ok := target.Value.(*SocketOption_IntValue); !ok {
			return false
		}

		if m.GetIntValue() != target.GetIntValue() {
			return false
		}

	case *SocketOption_BufValue:
		if _, ok := target.Value.(*SocketOption_BufValue); !ok {
			return false
		}

		if bytes.Compare(m.GetBufValue(), target.GetBufValue()) != 0 {
			return false
		}

	default:
		// m is nil but target is not nil
		if m.Value != target.Value {
			return false
		}
	}

	return true
}