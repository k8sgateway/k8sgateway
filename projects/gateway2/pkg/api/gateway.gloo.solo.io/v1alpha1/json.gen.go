// Code generated by skv2. DO NOT EDIT.

// Generated json marshal and unmarshal functions

package v1alpha1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	jsonpb "github.com/golang/protobuf/jsonpb"
	proto "github.com/golang/protobuf/proto"
	skv2jsonpb "github.com/solo-io/skv2/pkg/kube_jsonpb"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var (
	marshaller   = &skv2jsonpb.Marshaler{}
	unmarshaller = &jsonpb.Unmarshaler{
		AllowUnknownFields: true,
	}
	strictUnmarshaller = &jsonpb.Unmarshaler{}
)

// MarshalJSON is a custom marshaler for DataPlaneConfigSpec
func (this *DataPlaneConfigSpec) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for DataPlaneConfigSpec
func (this *DataPlaneConfigSpec) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for DataPlaneConfigStatus
func (this *DataPlaneConfigStatus) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for DataPlaneConfigStatus
func (this *DataPlaneConfigStatus) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}
