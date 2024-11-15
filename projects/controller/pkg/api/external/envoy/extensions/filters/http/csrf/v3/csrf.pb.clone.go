// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/external/envoy/extensions/filters/http/csrf/v3/csrf.proto

package v3

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_config_core_v3 "github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/config/core/v3"

	github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_type_matcher_v3 "github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/type/matcher/v3"
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
func (m *CsrfPolicy) Clone() proto.Message {
	var target *CsrfPolicy
	if m == nil {
		return target
	}
	target = &CsrfPolicy{}

	if h, ok := interface{}(m.GetFilterEnabled()).(clone.Cloner); ok {
		target.FilterEnabled = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_config_core_v3.RuntimeFractionalPercent)
	} else {
		target.FilterEnabled = proto.Clone(m.GetFilterEnabled()).(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_config_core_v3.RuntimeFractionalPercent)
	}

	if h, ok := interface{}(m.GetShadowEnabled()).(clone.Cloner); ok {
		target.ShadowEnabled = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_config_core_v3.RuntimeFractionalPercent)
	} else {
		target.ShadowEnabled = proto.Clone(m.GetShadowEnabled()).(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_config_core_v3.RuntimeFractionalPercent)
	}

	if m.GetAdditionalOrigins() != nil {
		target.AdditionalOrigins = make([]*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_type_matcher_v3.StringMatcher, len(m.GetAdditionalOrigins()))
		for idx, v := range m.GetAdditionalOrigins() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.AdditionalOrigins[idx] = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_type_matcher_v3.StringMatcher)
			} else {
				target.AdditionalOrigins[idx] = proto.Clone(v).(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_type_matcher_v3.StringMatcher)
			}

		}
	}

	return target
}
