// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/type/tracing/v3/custom_tag.proto

package v3

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_type_metadata_v3 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/type/metadata/v3"
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
func (m *CustomTag) Clone() proto.Message {
	var target *CustomTag
	if m == nil {
		return target
	}
	target = &CustomTag{}

	target.Tag = m.GetTag()

	switch m.Type.(type) {

	case *CustomTag_Literal_:

		if h, ok := interface{}(m.GetLiteral()).(clone.Cloner); ok {
			target.Type = &CustomTag_Literal_{
				Literal: h.Clone().(*CustomTag_Literal),
			}
		} else {
			target.Type = &CustomTag_Literal_{
				Literal: proto.Clone(m.GetLiteral()).(*CustomTag_Literal),
			}
		}

	case *CustomTag_Environment_:

		if h, ok := interface{}(m.GetEnvironment()).(clone.Cloner); ok {
			target.Type = &CustomTag_Environment_{
				Environment: h.Clone().(*CustomTag_Environment),
			}
		} else {
			target.Type = &CustomTag_Environment_{
				Environment: proto.Clone(m.GetEnvironment()).(*CustomTag_Environment),
			}
		}

	case *CustomTag_RequestHeader:

		if h, ok := interface{}(m.GetRequestHeader()).(clone.Cloner); ok {
			target.Type = &CustomTag_RequestHeader{
				RequestHeader: h.Clone().(*CustomTag_Header),
			}
		} else {
			target.Type = &CustomTag_RequestHeader{
				RequestHeader: proto.Clone(m.GetRequestHeader()).(*CustomTag_Header),
			}
		}

	case *CustomTag_Metadata_:

		if h, ok := interface{}(m.GetMetadata()).(clone.Cloner); ok {
			target.Type = &CustomTag_Metadata_{
				Metadata: h.Clone().(*CustomTag_Metadata),
			}
		} else {
			target.Type = &CustomTag_Metadata_{
				Metadata: proto.Clone(m.GetMetadata()).(*CustomTag_Metadata),
			}
		}

	}

	return target
}

// Clone function
func (m *CustomTag_Literal) Clone() proto.Message {
	var target *CustomTag_Literal
	if m == nil {
		return target
	}
	target = &CustomTag_Literal{}

	target.Value = m.GetValue()

	return target
}

// Clone function
func (m *CustomTag_Environment) Clone() proto.Message {
	var target *CustomTag_Environment
	if m == nil {
		return target
	}
	target = &CustomTag_Environment{}

	target.Name = m.GetName()

	target.DefaultValue = m.GetDefaultValue()

	return target
}

// Clone function
func (m *CustomTag_Header) Clone() proto.Message {
	var target *CustomTag_Header
	if m == nil {
		return target
	}
	target = &CustomTag_Header{}

	target.Name = m.GetName()

	target.DefaultValue = m.GetDefaultValue()

	return target
}

// Clone function
func (m *CustomTag_Metadata) Clone() proto.Message {
	var target *CustomTag_Metadata
	if m == nil {
		return target
	}
	target = &CustomTag_Metadata{}

	if h, ok := interface{}(m.GetKind()).(clone.Cloner); ok {
		target.Kind = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_type_metadata_v3.MetadataKind)
	} else {
		target.Kind = proto.Clone(m.GetKind()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_type_metadata_v3.MetadataKind)
	}

	if h, ok := interface{}(m.GetMetadataKey()).(clone.Cloner); ok {
		target.MetadataKey = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_type_metadata_v3.MetadataKey)
	} else {
		target.MetadataKey = proto.Clone(m.GetMetadataKey()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_type_metadata_v3.MetadataKey)
	}

	target.DefaultValue = m.GetDefaultValue()

	return target
}
