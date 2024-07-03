package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=gateway-gloo,shortName=gwp
// +kubebuilder:subresource:status
type GatewayParameters struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GatewayParametersSpec   `json:"spec,omitempty"`
	Status GatewayParametersStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type GatewayParametersList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GatewayParameters `json:"items"`
}

type GatewayParametersSpec struct {
	SelfManaged *SelfManagedGateway    `json:"selfManaged,omitempty"`
	Kube        *KubernetesProxyConfig `json:"kube,omitempty"`
}

type GatewayParametersStatus struct {
}

type SelfManagedGateway struct {
}

// Configuration for the set of Kubernetes resources that will be provisioned
// for a given Gateway.
type KubernetesProxyConfig struct {
	// +kubebuilder:validation:Optional
	Deployment *ProxyDeployment `json:"deployment,omitempty"`

	// Configuration for the container running Envoy.
	// +kubebuilder:validation:Optional
	EnvoyContainer *EnvoyContainer `json:"envoyContainer,omitempty"`
	// Configuration for the container running the Secret Discovery Service (SDS).
	// +kubebuilder:validation:Optional
	SdsContainer *SdsContainer `json:"sdsContainer,omitempty"`
	// Configuration for the pods that will be created.
	// +kubebuilder:validation:Optional
	PodTemplate *Pod `json:"podTemplate,omitempty"`
	// Configuration for the Kubernetes Service that exposes the Envoy proxy over
	// the network.
	// +kubebuilder:validation:Optional
	Service *Service `json:"service,omitempty"`
	// Autoscaling configuration.
	// Autoscaling Autoscaling `json:"autoscaling,omitempty"`
	// Configuration for the Istio integration.
	// +kubebuilder:validation:Optional
	Istio *IstioIntegration `json:"istio,omitempty"`
	// Configuration for the stats server.
	// +kubebuilder:validation:Optional
	Stats *StatsConfig `json:"stats,omitempty"`
	// Configuration for the AI extension.
	// +kubebuilder:validation:Optional
	AiExtension *AiExtension `json:"aiExtension,omitempty"`
}

func (in *KubernetesProxyConfig) GetDeployment() *ProxyDeployment {
	if in == nil {
		return nil
	}
	return in.Deployment
}

func (in *KubernetesProxyConfig) GetEnvoyContainer() *EnvoyContainer {
	if in == nil {
		return nil
	}
	return in.EnvoyContainer
}

func (in *KubernetesProxyConfig) GetSdsContainer() *SdsContainer {
	if in == nil {
		return nil
	}
	return in.SdsContainer
}

func (in *KubernetesProxyConfig) GetPodTemplate() *Pod {
	if in == nil {
		return nil
	}
	return in.PodTemplate
}

func (in *KubernetesProxyConfig) GetService() *Service {
	if in == nil {
		return nil
	}
	return in.Service
}

func (in *KubernetesProxyConfig) GetIstio() *IstioIntegration {
	if in == nil {
		return nil
	}
	return in.Istio
}

func (in *KubernetesProxyConfig) GetStats() *StatsConfig {
	if in == nil {
		return nil
	}
	return in.Stats
}

func (in *KubernetesProxyConfig) GetAiExtension() *AiExtension {
	if in == nil {
		return nil
	}
	return in.AiExtension
}

type ProxyDeployment struct {
	// The number of desired pods. Defaults to 1.
	// +kubebuilder:validation:Optional
	Replicas *uint32 `json:"replicas,omitempty"`
}

func (in *ProxyDeployment) GetReplicas() *uint32 {
	if in == nil {
		return nil
	}
	return in.Replicas
}

type EnvoyContainer struct {

	// Initial envoy configuration.
	// +kubebuilder:validation:Optional
	Bootstrap *EnvoyBootstrap `json:"bootstrap,omitempty"`
	// The envoy container image. See
	// https://kubernetes.io/docs/concepts/containers/images
	// for details.
	//
	// Default values, which may be overridden individually:
	//
	//	registry: quay.io/solo-io
	//	repository: gloo-envoy-wrapper (OSS) / gloo-ee-envoy-wrapper (EE)
	//	tag: <gloo version> (OSS) / <gloo-ee version> (EE)
	//	pullPolicy: IfNotPresent
	// +kubebuilder:validation:Optional
	Image *Image `json:"image,omitempty"`
	// The security context for this container. See
	// https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.26/#securitycontext-v1-core
	// for details.
	// +kubebuilder:validation:Optional
	SecurityContext *corev1.SecurityContext `json:"securityContext,omitempty"`
	// The compute resources required by this container. See
	// https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	// for details.
	// +kubebuilder:validation:Optional
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
}

func (in *EnvoyContainer) GetBootstrap() *EnvoyBootstrap {
	if in == nil {
		return nil
	}
	return in.Bootstrap
}

func (in *EnvoyContainer) GetImage() *Image {
	if in == nil {
		return nil
	}
	return in.Image
}

func (in *EnvoyContainer) GetSecurityContext() *corev1.SecurityContext {
	if in == nil {
		return nil
	}
	return in.SecurityContext
}

func (in *EnvoyContainer) GetResources() *corev1.ResourceRequirements {
	if in == nil {
		return nil
	}
	return in.Resources
}

type EnvoyBootstrap struct {

	// Envoy log level. Options include "trace", "debug", "info", "warn", "error",
	// "critical" and "off". Defaults to "info". See
	// https://www.envoyproxy.io/docs/envoy/latest/start/quick-start/run-envoy#debugging-envoy
	// for more information.
	// +kubebuilder:validation:Optional
	LogLevel string `json:"logLevel,omitempty"`
	// Envoy log levels for specific components. The keys are component names and
	// the values are one of "trace", "debug", "info", "warn", "error",
	// "critical", or "off", e.g.
	//
	//	```yaml
	//	componentLogLevels:
	//	  upstream: debug
	//	  connection: trace
	//	```
	//
	// These will be converted to the `--component-log-level` Envoy argument
	// value. See
	// https://www.envoyproxy.io/docs/envoy/latest/start/quick-start/run-envoy#debugging-envoy
	// for more information.
	//
	// Note: the keys and values cannot be empty, but they are not otherwise validated.
	// +kubebuilder:validation:Optional
	ComponentLogLevels map[string]string `json:"componentLogLevels,omitempty"`
}

func (in *EnvoyBootstrap) GetLogLevel() string {
	if in == nil {
		return ""
	}
	return in.LogLevel
}

func (in *EnvoyBootstrap) GetComponentLogLevels() map[string]string {
	if in == nil {
		return nil
	}
	return in.ComponentLogLevels
}

type SdsContainer struct {
	Image           *Image                       `json:"image,omitempty"`
	SecurityContext *corev1.SecurityContext      `json:"securityContext,omitempty"`
	Resources       *corev1.ResourceRequirements `json:"resources,omitempty"`
	Bootstrap       *SdsBootstrap                `json:"bootstrap,omitempty"`
}

func (in *SdsContainer) GetImage() *Image {
	if in == nil {
		return nil
	}
	return in.Image
}

func (in *SdsContainer) GetSecurityContext() *corev1.SecurityContext {
	if in == nil {
		return nil
	}
	return in.SecurityContext
}

func (in *SdsContainer) GetResources() *corev1.ResourceRequirements {
	if in == nil {
		return nil
	}
	return in.Resources
}

func (in *SdsContainer) GetBootstrap() *SdsBootstrap {
	if in == nil {
		return nil
	}
	return in.Bootstrap
}

type SdsBootstrap struct {
	LogLevel string `json:"logLevel,omitempty"`
}

func (in *SdsBootstrap) GetLogLevel() string {
	if in == nil {
		return ""
	}
	return in.LogLevel
}

type IstioIntegration struct {
	// +kubebuilder:validation:Optional
	IstioProxyContainer *IstioContainer `json:"istioProxyContainer,omitempty"`
	// +kubebuilder:validation:Optional
	CustomSidecars []*corev1.Container `json:"customSidecars,omitempty"`
}

func (in *IstioIntegration) GetIstioProxyContainer() *IstioContainer {
	if in == nil {
		return nil
	}
	return in.IstioProxyContainer
}

func (in *IstioIntegration) GetCustomSidecars() []*corev1.Container {
	if in == nil {
		return nil
	}
	return in.CustomSidecars
}

type IstioContainer struct {
	// The envoy container image. See
	// https://kubernetes.io/docs/concepts/containers/images
	// for details.
	//
	// Default values, which may be overridden individually:
	//
	//	registry: quay.io/solo-io
	//	repository: gloo-envoy-wrapper (OSS) / gloo-ee-envoy-wrapper (EE)
	//	tag: <gloo version> (OSS) / <gloo-ee version> (EE)
	//	pullPolicy: IfNotPresent
	// +kubebuilder:validation:Optional
	Image *Image `json:"image,omitempty"`
	// The security context for this container. See
	// https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.26/#securitycontext-v1-core
	// for details.
	// +kubebuilder:validation:Optional
	SecurityContext *corev1.SecurityContext `json:"securityContext,omitempty"`
	// The compute resources required by this container. See
	// https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	// for details.
	// +kubebuilder:validation:Optional
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`

	// +kubebuilder:validation:Optional
	LogLevel string `json:"logLevel,omitempty"`

	// +kubebuilder:validation:Optional
	IstioDiscoveryAddress string `json:"istioDiscoveryAddress,omitempty"`

	// +kubebuilder:validation:Optional
	IstioMetaMeshId string `json:"istioMetaMeshId,omitempty"`

	// +kubebuilder:validation:Optional
	IstioMetaClusterId string `json:"istioMetaClusterId,omitempty"`
}

func (in *IstioContainer) GetImage() *Image {
	if in == nil {
		return nil
	}
	return in.Image
}

func (in *IstioContainer) GetSecurityContext() *corev1.SecurityContext {
	if in == nil {
		return nil
	}
	return in.SecurityContext
}

func (in *IstioContainer) GetResources() *corev1.ResourceRequirements {
	if in == nil {
		return nil
	}
	return in.Resources
}

func (in *IstioContainer) GetLogLevel() string {
	if in == nil {
		return ""
	}
	return in.LogLevel
}

func (in *IstioContainer) GetIstioDiscoveryAddress() string {
	if in == nil {
		return ""
	}
	return in.IstioDiscoveryAddress
}

func (in *IstioContainer) GetIstioMetaMeshId() string {
	if in == nil {
		return ""
	}
	return in.IstioMetaMeshId
}

func (in *IstioContainer) GetIstioMetaClusterId() string {
	if in == nil {
		return ""
	}
	return in.IstioMetaClusterId
}

type StatsConfig struct {
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	RoutePrefixRewrite string `json:"routePrefixRewrite,omitempty"`

	// +kubebuilder:validation:Optional
	EnableStatsRoute *bool `json:"enableStatsRoute,omitempty"`

	// +kubebuilder:validation:Optional
	StatsRoutePrefixRewrite string `json:"statsRoutePrefixRewrite,omitempty"`
}

func (in *StatsConfig) GetEnabled() *bool {
	if in == nil {
		return nil
	}
	return in.Enabled
}

func (in *StatsConfig) GetRoutePrefixRewrite() string {
	if in == nil {
		return ""
	}
	return in.RoutePrefixRewrite
}

func (in *StatsConfig) GetEnableStatsRoute() *bool {
	if in == nil {
		return nil
	}
	return in.EnableStatsRoute
}

func (in *StatsConfig) GetStatsRoutePrefixRewrite() string {
	if in == nil {
		return ""
	}
	return in.StatsRoutePrefixRewrite
}

// Configuration for the AI extension.
type AiExtension struct {
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty"`
	// The envoy container image. See
	// https://kubernetes.io/docs/concepts/containers/images
	// for details.
	//
	// Default values, which may be overridden individually:
	//
	//	registry: quay.io/solo-io
	//	repository: gloo-envoy-wrapper (OSS) / gloo-ee-envoy-wrapper (EE)
	//	tag: <gloo version> (OSS) / <gloo-ee version> (EE)
	//	pullPolicy: IfNotPresent
	// +kubebuilder:validation:Optional
	Image *Image `json:"image,omitempty"`
	// The security context for this container. See
	// https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.26/#securitycontext-v1-core
	// for details.
	// +kubebuilder:validation:Optional
	SecurityContext *corev1.SecurityContext `json:"securityContext,omitempty"`
	// The compute resources required by this container. See
	// https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	// for details.
	// +kubebuilder:validation:Optional
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`

	// +kubebuilder:validation:Optional
	Env []*corev1.EnvVar `json:"env,omitempty"`

	// +kubebuilder:validation:Optional
	Ports []*corev1.ContainerPort `json:"ports,omitempty"`
}

func (in *AiExtension) GetEnabled() *bool {
	if in == nil {
		return nil
	}
	return in.Enabled
}

func (in *AiExtension) GetImage() *Image {
	if in == nil {
		return nil
	}
	return in.Image
}

func (in *AiExtension) GetSecurityContext() *corev1.SecurityContext {
	if in == nil {
		return nil
	}
	return in.SecurityContext
}

func (in *AiExtension) GetResources() *corev1.ResourceRequirements {
	if in == nil {
		return nil
	}
	return in.Resources
}

func (in *AiExtension) GetEnv() []*corev1.EnvVar {
	if in == nil {
		return nil
	}
	return in.Env
}

func (in *AiExtension) GetPorts() []*corev1.ContainerPort {
	if in == nil {
		return nil
	}
	return in.Ports
}

func init() {
	SchemeBuilder.Register(&GatewayParameters{}, &GatewayParametersList{})
}
