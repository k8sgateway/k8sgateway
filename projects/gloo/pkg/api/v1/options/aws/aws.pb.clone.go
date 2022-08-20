// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/aws/aws.proto

package aws

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_solo_kit_pkg_api_v1_resources_core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
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
func (m *UpstreamSpec) Clone() proto.Message {
	var target *UpstreamSpec
	if m == nil {
		return target
	}
	target = &UpstreamSpec{}

	target.Region = m.GetRegion()

	if h, ok := interface{}(m.GetSecretRef()).(clone.Cloner); ok {
		target.SecretRef = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	} else {
		target.SecretRef = proto.Clone(m.GetSecretRef()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	}

	if m.GetLambdaFunctions() != nil {
		target.LambdaFunctions = make([]*LambdaFunctionSpec, len(m.GetLambdaFunctions()))
		for idx, v := range m.GetLambdaFunctions() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.LambdaFunctions[idx] = h.Clone().(*LambdaFunctionSpec)
			} else {
				target.LambdaFunctions[idx] = proto.Clone(v).(*LambdaFunctionSpec)
			}

		}
	}

	target.RoleArn = m.GetRoleArn()

	target.DisableRoleChaining = m.GetDisableRoleChaining()

	return target
}

// Clone function
func (m *LambdaFunctionSpec) Clone() proto.Message {
	var target *LambdaFunctionSpec
	if m == nil {
		return target
	}
	target = &LambdaFunctionSpec{}

	target.LogicalName = m.GetLogicalName()

	target.LambdaFunctionName = m.GetLambdaFunctionName()

	target.Qualifier = m.GetQualifier()

	return target
}

// Clone function
func (m *DestinationSpec) Clone() proto.Message {
	var target *DestinationSpec
	if m == nil {
		return target
	}
	target = &DestinationSpec{}

	target.LogicalName = m.GetLogicalName()

	target.InvocationStyle = m.GetInvocationStyle()

	target.RequestTransformation = m.GetRequestTransformation()

	target.ResponseTransformation = m.GetResponseTransformation()

	target.UnwrapAsAlb = m.GetUnwrapAsAlb()

	target.UnwrapAsApiGateway = m.GetUnwrapAsApiGateway()

	return target
}
