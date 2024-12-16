// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	internal "github.com/solo-io/gloo/projects/gateway2/api/applyconfiguration/internal"
	apiv1alpha1 "github.com/solo-io/gloo/projects/gateway2/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// UpstreamApplyConfiguration represents a declarative configuration of the Upstream type for use
// with apply.
type UpstreamApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Spec                             *UpstreamSpecApplyConfiguration   `json:"spec,omitempty"`
	Status                           *UpstreamStatusApplyConfiguration `json:"status,omitempty"`
}

// Upstream constructs a declarative configuration of the Upstream type for use with
// apply.
func Upstream(name, namespace string) *UpstreamApplyConfiguration {
	b := &UpstreamApplyConfiguration{}
	b.WithName(name)
	b.WithNamespace(namespace)
	b.WithKind("Upstream")
	b.WithAPIVersion("gateway.gloo.solo.io/v1alpha1")
	return b
}

// ExtractUpstream extracts the applied configuration owned by fieldManager from
// upstream. If no managedFields are found in upstream for fieldManager, a
// UpstreamApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// upstream must be a unmodified Upstream API object that was retrieved from the Kubernetes API.
// ExtractUpstream provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractUpstream(upstream *apiv1alpha1.Upstream, fieldManager string) (*UpstreamApplyConfiguration, error) {
	return extractUpstream(upstream, fieldManager, "")
}

// ExtractUpstreamStatus is the same as ExtractUpstream except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractUpstreamStatus(upstream *apiv1alpha1.Upstream, fieldManager string) (*UpstreamApplyConfiguration, error) {
	return extractUpstream(upstream, fieldManager, "status")
}

func extractUpstream(upstream *apiv1alpha1.Upstream, fieldManager string, subresource string) (*UpstreamApplyConfiguration, error) {
	b := &UpstreamApplyConfiguration{}
	err := managedfields.ExtractInto(upstream, internal.Parser().Type("com.github.solo-io.gloo.projects.gateway2.api.v1alpha1.Upstream"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(upstream.Name)
	b.WithNamespace(upstream.Namespace)

	b.WithKind("Upstream")
	b.WithAPIVersion("gateway.gloo.solo.io/v1alpha1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *UpstreamApplyConfiguration) WithKind(value string) *UpstreamApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *UpstreamApplyConfiguration) WithAPIVersion(value string) *UpstreamApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *UpstreamApplyConfiguration) WithName(value string) *UpstreamApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *UpstreamApplyConfiguration) WithGenerateName(value string) *UpstreamApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *UpstreamApplyConfiguration) WithNamespace(value string) *UpstreamApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *UpstreamApplyConfiguration) WithUID(value types.UID) *UpstreamApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *UpstreamApplyConfiguration) WithResourceVersion(value string) *UpstreamApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *UpstreamApplyConfiguration) WithGeneration(value int64) *UpstreamApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *UpstreamApplyConfiguration) WithCreationTimestamp(value metav1.Time) *UpstreamApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *UpstreamApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *UpstreamApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *UpstreamApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *UpstreamApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *UpstreamApplyConfiguration) WithLabels(entries map[string]string) *UpstreamApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *UpstreamApplyConfiguration) WithAnnotations(entries map[string]string) *UpstreamApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *UpstreamApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *UpstreamApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.OwnerReferences = append(b.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *UpstreamApplyConfiguration) WithFinalizers(values ...string) *UpstreamApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

func (b *UpstreamApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithSpec sets the Spec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Spec field is set to the value of the last call.
func (b *UpstreamApplyConfiguration) WithSpec(value *UpstreamSpecApplyConfiguration) *UpstreamApplyConfiguration {
	b.Spec = value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *UpstreamApplyConfiguration) WithStatus(value *UpstreamStatusApplyConfiguration) *UpstreamApplyConfiguration {
	b.Status = value
	return b
}

// GetName retrieves the value of the Name field in the declarative configuration.
func (b *UpstreamApplyConfiguration) GetName() *string {
	b.ensureObjectMetaApplyConfigurationExists()
	return b.Name
}
