//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AiExtension) DeepCopyInto(out *AiExtension) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(Image)
		**out = **in
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]*v1.EnvVar, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(v1.EnvVar)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]*v1.ContainerPort, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(v1.ContainerPort)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AiExtension.
func (in *AiExtension) DeepCopy() *AiExtension {
	if in == nil {
		return nil
	}
	out := new(AiExtension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvoyBootstrap) DeepCopyInto(out *EnvoyBootstrap) {
	*out = *in
	if in.ComponentLogLevels != nil {
		in, out := &in.ComponentLogLevels, &out.ComponentLogLevels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyBootstrap.
func (in *EnvoyBootstrap) DeepCopy() *EnvoyBootstrap {
	if in == nil {
		return nil
	}
	out := new(EnvoyBootstrap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvoyContainer) DeepCopyInto(out *EnvoyContainer) {
	*out = *in
	if in.Bootstrap != nil {
		in, out := &in.Bootstrap, &out.Bootstrap
		*out = new(EnvoyBootstrap)
		(*in).DeepCopyInto(*out)
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(Image)
		**out = **in
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyContainer.
func (in *EnvoyContainer) DeepCopy() *EnvoyContainer {
	if in == nil {
		return nil
	}
	out := new(EnvoyContainer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayParameters) DeepCopyInto(out *GatewayParameters) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayParameters.
func (in *GatewayParameters) DeepCopy() *GatewayParameters {
	if in == nil {
		return nil
	}
	out := new(GatewayParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GatewayParameters) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayParametersList) DeepCopyInto(out *GatewayParametersList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GatewayParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayParametersList.
func (in *GatewayParametersList) DeepCopy() *GatewayParametersList {
	if in == nil {
		return nil
	}
	out := new(GatewayParametersList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GatewayParametersList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayParametersSpec) DeepCopyInto(out *GatewayParametersSpec) {
	*out = *in
	if in.SelfManaged != nil {
		in, out := &in.SelfManaged, &out.SelfManaged
		*out = new(SelfManagedGateway)
		**out = **in
	}
	if in.Kube != nil {
		in, out := &in.Kube, &out.Kube
		*out = new(KubernetesProxyConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayParametersSpec.
func (in *GatewayParametersSpec) DeepCopy() *GatewayParametersSpec {
	if in == nil {
		return nil
	}
	out := new(GatewayParametersSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayParametersStatus) DeepCopyInto(out *GatewayParametersStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayParametersStatus.
func (in *GatewayParametersStatus) DeepCopy() *GatewayParametersStatus {
	if in == nil {
		return nil
	}
	out := new(GatewayParametersStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IstioContainer) DeepCopyInto(out *IstioContainer) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(Image)
		**out = **in
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioContainer.
func (in *IstioContainer) DeepCopy() *IstioContainer {
	if in == nil {
		return nil
	}
	out := new(IstioContainer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IstioIntegration) DeepCopyInto(out *IstioIntegration) {
	*out = *in
	if in.IstioProxyContainer != nil {
		in, out := &in.IstioProxyContainer, &out.IstioProxyContainer
		*out = new(IstioContainer)
		(*in).DeepCopyInto(*out)
	}
	if in.CustomSidecars != nil {
		in, out := &in.CustomSidecars, &out.CustomSidecars
		*out = make([]*v1.Container, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(v1.Container)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioIntegration.
func (in *IstioIntegration) DeepCopy() *IstioIntegration {
	if in == nil {
		return nil
	}
	out := new(IstioIntegration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesProxyConfig) DeepCopyInto(out *KubernetesProxyConfig) {
	*out = *in
	if in.Deployment != nil {
		in, out := &in.Deployment, &out.Deployment
		*out = new(ProxyDeployment)
		(*in).DeepCopyInto(*out)
	}
	if in.EnvoyContainer != nil {
		in, out := &in.EnvoyContainer, &out.EnvoyContainer
		*out = new(EnvoyContainer)
		(*in).DeepCopyInto(*out)
	}
	if in.SdsContainer != nil {
		in, out := &in.SdsContainer, &out.SdsContainer
		*out = new(SdsContainer)
		(*in).DeepCopyInto(*out)
	}
	if in.PodTemplate != nil {
		in, out := &in.PodTemplate, &out.PodTemplate
		*out = new(Pod)
		(*in).DeepCopyInto(*out)
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(Service)
		(*in).DeepCopyInto(*out)
	}
	if in.Istio != nil {
		in, out := &in.Istio, &out.Istio
		*out = new(IstioIntegration)
		(*in).DeepCopyInto(*out)
	}
	if in.Stats != nil {
		in, out := &in.Stats, &out.Stats
		*out = new(StatsConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.AiExtension != nil {
		in, out := &in.AiExtension, &out.AiExtension
		*out = new(AiExtension)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesProxyConfig.
func (in *KubernetesProxyConfig) DeepCopy() *KubernetesProxyConfig {
	if in == nil {
		return nil
	}
	out := new(KubernetesProxyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pod) DeepCopyInto(out *Pod) {
	*out = *in
	if in.ExtraLabels != nil {
		in, out := &in.ExtraLabels, &out.ExtraLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ExtraAnnotations != nil {
		in, out := &in.ExtraAnnotations, &out.ExtraAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]*v1.Toleration, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(v1.Toleration)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pod.
func (in *Pod) DeepCopy() *Pod {
	if in == nil {
		return nil
	}
	out := new(Pod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxyDeployment) DeepCopyInto(out *ProxyDeployment) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(uint32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyDeployment.
func (in *ProxyDeployment) DeepCopy() *ProxyDeployment {
	if in == nil {
		return nil
	}
	out := new(ProxyDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SdsBootstrap) DeepCopyInto(out *SdsBootstrap) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SdsBootstrap.
func (in *SdsBootstrap) DeepCopy() *SdsBootstrap {
	if in == nil {
		return nil
	}
	out := new(SdsBootstrap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SdsContainer) DeepCopyInto(out *SdsContainer) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(Image)
		**out = **in
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Bootstrap != nil {
		in, out := &in.Bootstrap, &out.Bootstrap
		*out = new(SdsBootstrap)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SdsContainer.
func (in *SdsContainer) DeepCopy() *SdsContainer {
	if in == nil {
		return nil
	}
	out := new(SdsContainer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelfManagedGateway) DeepCopyInto(out *SelfManagedGateway) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelfManagedGateway.
func (in *SelfManagedGateway) DeepCopy() *SelfManagedGateway {
	if in == nil {
		return nil
	}
	out := new(SelfManagedGateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Service) DeepCopyInto(out *Service) {
	*out = *in
	if in.ExtraLabels != nil {
		in, out := &in.ExtraLabels, &out.ExtraLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ExtraAnnotations != nil {
		in, out := &in.ExtraAnnotations, &out.ExtraAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Service.
func (in *Service) DeepCopy() *Service {
	if in == nil {
		return nil
	}
	out := new(Service)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatsConfig) DeepCopyInto(out *StatsConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.EnableStatsRoute != nil {
		in, out := &in.EnableStatsRoute, &out.EnableStatsRoute
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatsConfig.
func (in *StatsConfig) DeepCopy() *StatsConfig {
	if in == nil {
		return nil
	}
	out := new(StatsConfig)
	in.DeepCopyInto(out)
	return out
}
