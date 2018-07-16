package rest

import (
	"github.com/gogo/protobuf/types"
	"github.com/solo-io/gloo/pkg/protoutil"
)

func DecodeRouteExtension(generic *types.Struct) (RouteExtension, error) {
	var s RouteExtension
	err := protoutil.UnmarshalStruct(generic, &s)
	return s, err
}

func EncodeRouteExtension(spec RouteExtension) *types.Struct {
	v1Spec, err := protoutil.MarshalStruct(spec)
	if err != nil {
		panic(err)
	}
	return v1Spec
}

func DecodeFunctionSpec(generic *types.Struct) (TransformationSpec, error) {
	var s TransformationSpec
	err := protoutil.UnmarshalStruct(generic, &s)
	return s, err
}

func EncodeFunctionSpec(spec TransformationSpec) *types.Struct {
	v1Spec, err := protoutil.MarshalStruct(spec)
	if err != nil {
		panic(err)
	}
	return v1Spec
}
