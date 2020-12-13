// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"encoding/json"

	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/utils/protoutils"

	api "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/extauth/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type metaOnly struct {
	v1.TypeMeta   `json:",inline"`
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resourceName=authconfigs
// +genclient
type AuthConfig struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec   api.AuthConfig `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status core.Status    `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

func (o *AuthConfig) MarshalJSON() ([]byte, error) {
	spec, err := protoutils.MarshalMap(&o.Spec)
	if err != nil {
		return nil, err
	}
	delete(spec, "metadata")
	delete(spec, "status")
	asMap := map[string]interface{}{
		"metadata":   o.ObjectMeta,
		"apiVersion": o.TypeMeta.APIVersion,
		"kind":       o.TypeMeta.Kind,
		"status":     o.Status,
		"spec":       spec,
	}
	return json.Marshal(asMap)
}

func (o *AuthConfig) UnmarshalJSON(data []byte) error {
	var metaOnly metaOnly
	if err := json.Unmarshal(data, &metaOnly); err != nil {
		return err
	}
	var spec api.AuthConfig
	if err := protoutils.UnmarshalResource(data, &spec); err != nil {
		return err
	}
	spec.Metadata = nil
	*o = AuthConfig{
		ObjectMeta: metaOnly.ObjectMeta,
		TypeMeta:   metaOnly.TypeMeta,
		Spec:       spec,
	}
	if spec.Status != nil {
		o.Status = *spec.Status
		o.Spec.Status = nil
	}

	return nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AuthConfigList is a collection of AuthConfigs.
type AuthConfigList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []AuthConfig `json:"items" protobuf:"bytes,2,rep,name=items"`
}
