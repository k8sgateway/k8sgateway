// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/v1/options/rest/rest.proto

package rest

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_transformation "github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/extensions/transformation"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_transformation "github.com/solo-io/gloo/projects/controller/pkg/api/v1/options/transformation"
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
func (m *ServiceSpec) Clone() proto.Message {
	var target *ServiceSpec
	if m == nil {
		return target
	}
	target = &ServiceSpec{}

	if m.GetTransformations() != nil {
		target.Transformations = make(map[string]*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_transformation.TransformationTemplate, len(m.GetTransformations()))
		for k, v := range m.GetTransformations() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Transformations[k] = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_transformation.TransformationTemplate)
			} else {
				target.Transformations[k] = proto.Clone(v).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_transformation.TransformationTemplate)
			}

		}
	}

	if h, ok := interface{}(m.GetSwaggerInfo()).(clone.Cloner); ok {
		target.SwaggerInfo = h.Clone().(*ServiceSpec_SwaggerInfo)
	} else {
		target.SwaggerInfo = proto.Clone(m.GetSwaggerInfo()).(*ServiceSpec_SwaggerInfo)
	}

	return target
}

// Clone function
func (m *DestinationSpec) Clone() proto.Message {
	var target *DestinationSpec
	if m == nil {
		return target
	}
	target = &DestinationSpec{}

	target.FunctionName = m.GetFunctionName()

	if h, ok := interface{}(m.GetParameters()).(clone.Cloner); ok {
		target.Parameters = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_transformation.Parameters)
	} else {
		target.Parameters = proto.Clone(m.GetParameters()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_transformation.Parameters)
	}

	if h, ok := interface{}(m.GetResponseTransformation()).(clone.Cloner); ok {
		target.ResponseTransformation = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_transformation.TransformationTemplate)
	} else {
		target.ResponseTransformation = proto.Clone(m.GetResponseTransformation()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_transformation.TransformationTemplate)
	}

	return target
}

// Clone function
func (m *ServiceSpec_SwaggerInfo) Clone() proto.Message {
	var target *ServiceSpec_SwaggerInfo
	if m == nil {
		return target
	}
	target = &ServiceSpec_SwaggerInfo{}

	switch m.SwaggerSpec.(type) {

	case *ServiceSpec_SwaggerInfo_Url:

		target.SwaggerSpec = &ServiceSpec_SwaggerInfo_Url{
			Url: m.GetUrl(),
		}

	case *ServiceSpec_SwaggerInfo_Inline:

		target.SwaggerSpec = &ServiceSpec_SwaggerInfo_Inline{
			Inline: m.GetInline(),
		}

	}

	return target
}
