// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/local_ratelimit/local_ratelimit.proto

package local_ratelimit

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
func (m *TokenBucket) Clone() proto.Message {
	var target *TokenBucket
	if m == nil {
		return target
	}
	target = &TokenBucket{}

	target.MaxTokens = m.GetMaxTokens()

	if h, ok := interface{}(m.GetTokensPerFill()).(clone.Cloner); ok {
		target.TokensPerFill = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.TokensPerFill = proto.Clone(m.GetTokensPerFill()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	if h, ok := interface{}(m.GetFillInterval()).(clone.Cloner); ok {
		target.FillInterval = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.FillInterval = proto.Clone(m.GetFillInterval()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	return target
}

// Clone function
func (m *Settings) Clone() proto.Message {
	var target *Settings
	if m == nil {
		return target
	}
	target = &Settings{}

	if h, ok := interface{}(m.GetDefaults()).(clone.Cloner); ok {
		target.Defaults = h.Clone().(*TokenBucket)
	} else {
		target.Defaults = proto.Clone(m.GetDefaults()).(*TokenBucket)
	}

	target.LocalRateLimitPerDownstreamConnection = m.GetLocalRateLimitPerDownstreamConnection()

	target.EnableXRatelimitHeaders = m.GetEnableXRatelimitHeaders()

	return target
}
