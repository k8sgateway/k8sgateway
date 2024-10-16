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

	google_golang_org_protobuf_types_known_durationpb "google.golang.org/protobuf/types/known/durationpb"

	google_golang_org_protobuf_types_known_wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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
		target.TokensPerFill = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.TokensPerFill = proto.Clone(m.GetTokensPerFill()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	if h, ok := interface{}(m.GetFillInterval()).(clone.Cloner); ok {
		target.FillInterval = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.FillInterval = proto.Clone(m.GetFillInterval()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
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

	if h, ok := interface{}(m.GetDefaultLimit()).(clone.Cloner); ok {
		target.DefaultLimit = h.Clone().(*TokenBucket)
	} else {
		target.DefaultLimit = proto.Clone(m.GetDefaultLimit()).(*TokenBucket)
	}

	if h, ok := interface{}(m.GetLocalRateLimitPerDownstreamConnection()).(clone.Cloner); ok {
		target.LocalRateLimitPerDownstreamConnection = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.LocalRateLimitPerDownstreamConnection = proto.Clone(m.GetLocalRateLimitPerDownstreamConnection()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	if h, ok := interface{}(m.GetEnableXRatelimitHeaders()).(clone.Cloner); ok {
		target.EnableXRatelimitHeaders = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.EnableXRatelimitHeaders = proto.Clone(m.GetEnableXRatelimitHeaders()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	return target
}
