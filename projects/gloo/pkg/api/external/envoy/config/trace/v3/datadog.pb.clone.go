// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/config/trace/v3/datadog.proto

package v3

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_wrappers "github.com/golang/protobuf/ptypes/wrappers"

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
func (m *DatadogConfig) Clone() proto.Message {
	var target *DatadogConfig
	if m == nil {
		return target
	}
	target = &DatadogConfig{}

	if h, ok := interface{}(m.GetServiceName()).(clone.Cloner); ok {
		target.ServiceName = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.StringValue)
	} else {
		target.ServiceName = proto.Clone(m.GetServiceName()).(*github_com_golang_protobuf_ptypes_wrappers.StringValue)
	}

	switch m.CollectorCluster.(type) {

	case *DatadogConfig_CollectorUpstreamRef:

		if h, ok := interface{}(m.GetCollectorUpstreamRef()).(clone.Cloner); ok {
			target.CollectorCluster = &DatadogConfig_CollectorUpstreamRef{
				CollectorUpstreamRef: h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef),
			}
		} else {
			target.CollectorCluster = &DatadogConfig_CollectorUpstreamRef{
				CollectorUpstreamRef: proto.Clone(m.GetCollectorUpstreamRef()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef),
			}
		}

	case *DatadogConfig_ClusterName:

		target.CollectorCluster = &DatadogConfig_ClusterName{
			ClusterName: m.GetClusterName(),
		}

	}

	return target
}
