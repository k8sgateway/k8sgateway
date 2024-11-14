// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"encoding/json"

	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/utils/protoutils"

	api "github.com/solo-io/gloo/projects/controller/pkg/api/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type metaOnly struct {
	v1.TypeMeta   `json:",inline"`
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resourceName=artifacts
// +genclient
// +genclient:noStatus
type Artifact struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec api.Artifact `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

func (o *Artifact) MarshalJSON() ([]byte, error) {
	spec, err := protoutils.MarshalMap(&o.Spec)
	if err != nil {
		return nil, err
	}
	delete(spec, "metadata")
	asMap := map[string]interface{}{
		"metadata":   o.ObjectMeta,
		"apiVersion": o.TypeMeta.APIVersion,
		"kind":       o.TypeMeta.Kind,
		"spec":       spec,
	}
	return json.Marshal(asMap)
}

func (o *Artifact) UnmarshalJSON(data []byte) error {
	var metaOnly metaOnly
	if err := json.Unmarshal(data, &metaOnly); err != nil {
		return err
	}
	var spec api.Artifact
	if err := protoutils.UnmarshalResource(data, &spec); err != nil {
		return err
	}
	spec.Metadata = nil
	*o = Artifact{
		ObjectMeta: metaOnly.ObjectMeta,
		TypeMeta:   metaOnly.TypeMeta,
		Spec:       spec,
	}

	return nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ArtifactList is a collection of Artifacts.
type ArtifactList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []Artifact `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resourceName=endpoints
// +genclient
// +genclient:noStatus
type Endpoint struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec api.Endpoint `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

func (o *Endpoint) MarshalJSON() ([]byte, error) {
	spec, err := protoutils.MarshalMap(&o.Spec)
	if err != nil {
		return nil, err
	}
	delete(spec, "metadata")
	asMap := map[string]interface{}{
		"metadata":   o.ObjectMeta,
		"apiVersion": o.TypeMeta.APIVersion,
		"kind":       o.TypeMeta.Kind,
		"spec":       spec,
	}
	return json.Marshal(asMap)
}

func (o *Endpoint) UnmarshalJSON(data []byte) error {
	var metaOnly metaOnly
	if err := json.Unmarshal(data, &metaOnly); err != nil {
		return err
	}
	var spec api.Endpoint
	if err := protoutils.UnmarshalResource(data, &spec); err != nil {
		return err
	}
	spec.Metadata = nil
	*o = Endpoint{
		ObjectMeta: metaOnly.ObjectMeta,
		TypeMeta:   metaOnly.TypeMeta,
		Spec:       spec,
	}

	return nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// EndpointList is a collection of Endpoints.
type EndpointList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []Endpoint `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resourceName=proxies
// +genclient
type Proxy struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec   api.Proxy               `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status core.NamespacedStatuses `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

func (o *Proxy) MarshalJSON() ([]byte, error) {
	spec, err := protoutils.MarshalMap(&o.Spec)
	if err != nil {
		return nil, err
	}
	delete(spec, "metadata")
	delete(spec, "namespacedStatuses")
	asMap := map[string]interface{}{
		"metadata":   o.ObjectMeta,
		"apiVersion": o.TypeMeta.APIVersion,
		"kind":       o.TypeMeta.Kind,
		"status":     o.Status,
		"spec":       spec,
	}
	return json.Marshal(asMap)
}

func (o *Proxy) UnmarshalJSON(data []byte) error {
	var metaOnly metaOnly
	if err := json.Unmarshal(data, &metaOnly); err != nil {
		return err
	}
	var spec api.Proxy
	if err := protoutils.UnmarshalResource(data, &spec); err != nil {
		return err
	}
	spec.Metadata = nil
	*o = Proxy{
		ObjectMeta: metaOnly.ObjectMeta,
		TypeMeta:   metaOnly.TypeMeta,
		Spec:       spec,
	}
	if spec.GetNamespacedStatuses() != nil {
		o.Status = *spec.NamespacedStatuses
		o.Spec.NamespacedStatuses = nil
	}

	return nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ProxyList is a collection of Proxys.
type ProxyList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []Proxy `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resourceName=secrets
// +genclient
// +genclient:noStatus
type Secret struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec api.Secret `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

func (o *Secret) MarshalJSON() ([]byte, error) {
	spec, err := protoutils.MarshalMap(&o.Spec)
	if err != nil {
		return nil, err
	}
	delete(spec, "metadata")
	asMap := map[string]interface{}{
		"metadata":   o.ObjectMeta,
		"apiVersion": o.TypeMeta.APIVersion,
		"kind":       o.TypeMeta.Kind,
		"spec":       spec,
	}
	return json.Marshal(asMap)
}

func (o *Secret) UnmarshalJSON(data []byte) error {
	var metaOnly metaOnly
	if err := json.Unmarshal(data, &metaOnly); err != nil {
		return err
	}
	var spec api.Secret
	if err := protoutils.UnmarshalResource(data, &spec); err != nil {
		return err
	}
	spec.Metadata = nil
	*o = Secret{
		ObjectMeta: metaOnly.ObjectMeta,
		TypeMeta:   metaOnly.TypeMeta,
		Spec:       spec,
	}

	return nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SecretList is a collection of Secrets.
type SecretList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []Secret `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resourceName=settings
// +genclient
type Settings struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec   api.Settings            `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status core.NamespacedStatuses `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

func (o *Settings) MarshalJSON() ([]byte, error) {
	spec, err := protoutils.MarshalMap(&o.Spec)
	if err != nil {
		return nil, err
	}
	delete(spec, "metadata")
	delete(spec, "namespacedStatuses")
	asMap := map[string]interface{}{
		"metadata":   o.ObjectMeta,
		"apiVersion": o.TypeMeta.APIVersion,
		"kind":       o.TypeMeta.Kind,
		"status":     o.Status,
		"spec":       spec,
	}
	return json.Marshal(asMap)
}

func (o *Settings) UnmarshalJSON(data []byte) error {
	var metaOnly metaOnly
	if err := json.Unmarshal(data, &metaOnly); err != nil {
		return err
	}
	var spec api.Settings
	if err := protoutils.UnmarshalResource(data, &spec); err != nil {
		return err
	}
	spec.Metadata = nil
	*o = Settings{
		ObjectMeta: metaOnly.ObjectMeta,
		TypeMeta:   metaOnly.TypeMeta,
		Spec:       spec,
	}
	if spec.GetNamespacedStatuses() != nil {
		o.Status = *spec.NamespacedStatuses
		o.Spec.NamespacedStatuses = nil
	}

	return nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SettingsList is a collection of Settingss.
type SettingsList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []Settings `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resourceName=upstreams
// +genclient
type Upstream struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec   api.Upstream            `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status core.NamespacedStatuses `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

func (o *Upstream) MarshalJSON() ([]byte, error) {
	spec, err := protoutils.MarshalMap(&o.Spec)
	if err != nil {
		return nil, err
	}
	delete(spec, "metadata")
	delete(spec, "namespacedStatuses")
	asMap := map[string]interface{}{
		"metadata":   o.ObjectMeta,
		"apiVersion": o.TypeMeta.APIVersion,
		"kind":       o.TypeMeta.Kind,
		"status":     o.Status,
		"spec":       spec,
	}
	return json.Marshal(asMap)
}

func (o *Upstream) UnmarshalJSON(data []byte) error {
	var metaOnly metaOnly
	if err := json.Unmarshal(data, &metaOnly); err != nil {
		return err
	}
	var spec api.Upstream
	if err := protoutils.UnmarshalResource(data, &spec); err != nil {
		return err
	}
	spec.Metadata = nil
	*o = Upstream{
		ObjectMeta: metaOnly.ObjectMeta,
		TypeMeta:   metaOnly.TypeMeta,
		Spec:       spec,
	}
	if spec.GetNamespacedStatuses() != nil {
		o.Status = *spec.NamespacedStatuses
		o.Spec.NamespacedStatuses = nil
	}

	return nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// UpstreamList is a collection of Upstreams.
type UpstreamList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []Upstream `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resourceName=upstreamgroups
// +genclient
type UpstreamGroup struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec   api.UpstreamGroup       `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status core.NamespacedStatuses `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

func (o *UpstreamGroup) MarshalJSON() ([]byte, error) {
	spec, err := protoutils.MarshalMap(&o.Spec)
	if err != nil {
		return nil, err
	}
	delete(spec, "metadata")
	delete(spec, "namespacedStatuses")
	asMap := map[string]interface{}{
		"metadata":   o.ObjectMeta,
		"apiVersion": o.TypeMeta.APIVersion,
		"kind":       o.TypeMeta.Kind,
		"status":     o.Status,
		"spec":       spec,
	}
	return json.Marshal(asMap)
}

func (o *UpstreamGroup) UnmarshalJSON(data []byte) error {
	var metaOnly metaOnly
	if err := json.Unmarshal(data, &metaOnly); err != nil {
		return err
	}
	var spec api.UpstreamGroup
	if err := protoutils.UnmarshalResource(data, &spec); err != nil {
		return err
	}
	spec.Metadata = nil
	*o = UpstreamGroup{
		ObjectMeta: metaOnly.ObjectMeta,
		TypeMeta:   metaOnly.TypeMeta,
		Spec:       spec,
	}
	if spec.GetNamespacedStatuses() != nil {
		o.Status = *spec.NamespacedStatuses
		o.Spec.NamespacedStatuses = nil
	}

	return nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// UpstreamGroupList is a collection of UpstreamGroups.
type UpstreamGroupList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []UpstreamGroup `json:"items" protobuf:"bytes,2,rep,name=items"`
}
