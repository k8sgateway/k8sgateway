// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/consul/consul.proto

package consul

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options"
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

	target.ServiceName = m.GetServiceName()

	if m.GetServiceTags() != nil {
		target.ServiceTags = make([]string, len(m.GetServiceTags()))
		for idx, v := range m.GetServiceTags() {

			target.ServiceTags[idx] = v

		}
	}

	if m.GetSubsetTags() != nil {
		target.SubsetTags = make([]string, len(m.GetSubsetTags()))
		for idx, v := range m.GetSubsetTags() {

			target.SubsetTags[idx] = v

		}
	}

	if m.GetInstanceTags() != nil {
		target.InstanceTags = make([]string, len(m.GetInstanceTags()))
		for idx, v := range m.GetInstanceTags() {

			target.InstanceTags[idx] = v

		}
	}

	if m.GetInstanceBlacklistTags() != nil {
		target.InstanceBlacklistTags = make([]string, len(m.GetInstanceBlacklistTags()))
		for idx, v := range m.GetInstanceBlacklistTags() {

			target.InstanceBlacklistTags[idx] = v

		}
	}

	if h, ok := interface{}(m.GetServiceSpec()).(clone.Cloner); ok {
		target.ServiceSpec = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options.ServiceSpec)
	} else {
		target.ServiceSpec = proto.Clone(m.GetServiceSpec()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options.ServiceSpec)
	}

	target.ConsistencyMode = m.GetConsistencyMode()

	if h, ok := interface{}(m.GetQueryOptions()).(clone.Cloner); ok {
		target.QueryOptions = h.Clone().(*QueryOptions)
	} else {
		target.QueryOptions = proto.Clone(m.GetQueryOptions()).(*QueryOptions)
	}

	target.ConnectEnabled = m.GetConnectEnabled()

	if m.GetDataCenters() != nil {
		target.DataCenters = make([]string, len(m.GetDataCenters()))
		for idx, v := range m.GetDataCenters() {

			target.DataCenters[idx] = v

		}
	}

	return target
}
