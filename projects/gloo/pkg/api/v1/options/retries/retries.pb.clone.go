// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/retries/retries.proto

package retries

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_duration "github.com/golang/protobuf/ptypes/duration"

	github_com_golang_protobuf_ptypes_wrappers "github.com/golang/protobuf/ptypes/wrappers"
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
func (m *RetryBackOff) Clone() proto.Message {
	var target *RetryBackOff
	if m == nil {
		return target
	}
	target = &RetryBackOff{}

	if h, ok := interface{}(m.GetBaseInterval()).(clone.Cloner); ok {
		target.BaseInterval = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.BaseInterval = proto.Clone(m.GetBaseInterval()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if h, ok := interface{}(m.GetMaxInterval()).(clone.Cloner); ok {
		target.MaxInterval = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.MaxInterval = proto.Clone(m.GetMaxInterval()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	return target
}

// Clone function
func (m *RetryPolicy) Clone() proto.Message {
	var target *RetryPolicy
	if m == nil {
		return target
	}
	target = &RetryPolicy{}

	target.RetryOn = m.GetRetryOn()

	target.NumRetries = m.GetNumRetries()

	if h, ok := interface{}(m.GetPerTryTimeout()).(clone.Cloner); ok {
		target.PerTryTimeout = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.PerTryTimeout = proto.Clone(m.GetPerTryTimeout()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if h, ok := interface{}(m.GetRetryBackOff()).(clone.Cloner); ok {
		target.RetryBackOff = h.Clone().(*RetryBackOff)
	} else {
		target.RetryBackOff = proto.Clone(m.GetRetryBackOff()).(*RetryBackOff)
	}

	if m.GetRetriableStatusCodes() != nil {
		target.RetriableStatusCodes = make([]uint32, len(m.GetRetriableStatusCodes()))
		for idx, v := range m.GetRetriableStatusCodes() {

			target.RetriableStatusCodes[idx] = v

		}
	}

	switch m.PriorityPredicate.(type) {

	case *RetryPolicy_PreviousPriorities_:

		if h, ok := interface{}(m.GetPreviousPriorities()).(clone.Cloner); ok {
			target.PriorityPredicate = &RetryPolicy_PreviousPriorities_{
				PreviousPriorities: h.Clone().(*RetryPolicy_PreviousPriorities),
			}
		} else {
			target.PriorityPredicate = &RetryPolicy_PreviousPriorities_{
				PreviousPriorities: proto.Clone(m.GetPreviousPriorities()).(*RetryPolicy_PreviousPriorities),
			}
		}

	}

	return target
}

// Clone function
func (m *RetryPolicy_PreviousPriorities) Clone() proto.Message {
	var target *RetryPolicy_PreviousPriorities
	if m == nil {
		return target
	}
	target = &RetryPolicy_PreviousPriorities{}

	if h, ok := interface{}(m.GetUpdateFrequency()).(clone.Cloner); ok {
		target.UpdateFrequency = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.UpdateFrequency = proto.Clone(m.GetUpdateFrequency()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	return target
}
