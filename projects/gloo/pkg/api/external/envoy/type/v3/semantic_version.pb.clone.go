// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/type/v3/semantic_version.proto

package v3

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"
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
func (m *SemanticVersion) Clone() proto.Message {
	var target *SemanticVersion
	if m == nil {
		return target
	}
	target = &SemanticVersion{}

	target.MajorNumber = m.GetMajorNumber()

	target.MinorNumber = m.GetMinorNumber()

	target.Patch = m.GetPatch()

	return target
}
