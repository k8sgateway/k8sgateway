// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for proto-based Spec and Status fields

package v1alpha1

import (
	proto "github.com/golang/protobuf/proto"
	"github.com/solo-io/protoc-gen-ext/pkg/clone"
)

// DeepCopyInto for the GatewayConfig.Spec
func (in *GatewayConfigSpec) DeepCopyInto(out *GatewayConfigSpec) {
	var p *GatewayConfigSpec
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*GatewayConfigSpec)
	} else {
		p = proto.Clone(in).(*GatewayConfigSpec)
	}
	*out = *p
}

// DeepCopyInto for the GatewayConfig.Status
func (in *GatewayConfigStatus) DeepCopyInto(out *GatewayConfigStatus) {
	var p *GatewayConfigStatus
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*GatewayConfigStatus)
	} else {
		p = proto.Clone(in).(*GatewayConfigStatus)
	}
	*out = *p
}
