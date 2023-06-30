// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/tcp/tcp.proto

package tcp

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
func (m *TcpProxySettings) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TcpProxySettings)
	if !ok {
		that2, ok := that.(TcpProxySettings)
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

	if h, ok := interface{}(m.GetMaxConnectAttempts()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMaxConnectAttempts()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMaxConnectAttempts(), target.GetMaxConnectAttempts()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetIdleTimeout()).(equality.Equalizer); ok {
		if !h.Equal(target.GetIdleTimeout()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetIdleTimeout(), target.GetIdleTimeout()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetTunnelingConfig()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTunnelingConfig()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTunnelingConfig(), target.GetTunnelingConfig()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetAccessLogFlushInterval()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAccessLogFlushInterval()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAccessLogFlushInterval(), target.GetAccessLogFlushInterval()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *HeaderValueOption) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HeaderValueOption)
	if !ok {
		that2, ok := that.(HeaderValueOption)
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

	if h, ok := interface{}(m.GetHeader()).(equality.Equalizer); ok {
		if !h.Equal(target.GetHeader()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetHeader(), target.GetHeader()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetAppend()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAppend()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAppend(), target.GetAppend()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *HeaderValue) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HeaderValue)
	if !ok {
		that2, ok := that.(HeaderValue)
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

// Equal function
func (m *TcpProxySettings_TunnelingConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TcpProxySettings_TunnelingConfig)
	if !ok {
		that2, ok := that.(TcpProxySettings_TunnelingConfig)
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

	if strings.Compare(m.GetHostname(), target.GetHostname()) != 0 {
		return false
	}

	if len(m.GetHeadersToAdd()) != len(target.GetHeadersToAdd()) {
		return false
	}
	for idx, v := range m.GetHeadersToAdd() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetHeadersToAdd()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetHeadersToAdd()[idx]) {
				return false
			}
		}

	}

	return true
}
