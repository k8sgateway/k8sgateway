// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gateway2/api/v1alpha1/data_plane_config.proto

package v1alpha1

import (
	reflect "reflect"
	sync "sync"

	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	v1 "github.com/solo-io/gloo/projects/gateway2/pkg/api/external/kubernetes/api/core/v1"
	v11 "github.com/solo-io/gloo/projects/gateway2/pkg/api/external/kubernetes/apimachinery/pkg/apis/meta/v1"
	kube "github.com/solo-io/gloo/projects/gateway2/pkg/api/gateway.gloo.solo.io/v1alpha1/kube"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A DataPlaneConfig contains configuration that is used to dynamically
// provision Gloo Gateway's data plane (Envoy proxy instance), based on a
// Kubernetes Gateway or GatewayClass.
type DataPlaneConfigSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProxyConfig *ProxyConfig `protobuf:"bytes,1,opt,name=proxy_config,json=proxyConfig,proto3" json:"proxy_config,omitempty"`
}

func (x *DataPlaneConfigSpec) Reset() {
	*x = DataPlaneConfigSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataPlaneConfigSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataPlaneConfigSpec) ProtoMessage() {}

func (x *DataPlaneConfigSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataPlaneConfigSpec.ProtoReflect.Descriptor instead.
func (*DataPlaneConfigSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_rawDescGZIP(), []int{0}
}

func (x *DataPlaneConfigSpec) GetProxyConfig() *ProxyConfig {
	if x != nil {
		return x.ProxyConfig
	}
	return nil
}

// Configuration for the proxy instance that implements a given Gateway.
type ProxyConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of environment/platform in which the proxy will be provisioned.
	//
	// Types that are assignable to EnvironmentType:
	//
	//	*ProxyConfig_Kube
	EnvironmentType isProxyConfig_EnvironmentType `protobuf_oneof:"environment_type"`
}

func (x *ProxyConfig) Reset() {
	*x = ProxyConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyConfig) ProtoMessage() {}

func (x *ProxyConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyConfig.ProtoReflect.Descriptor instead.
func (*ProxyConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_rawDescGZIP(), []int{1}
}

func (m *ProxyConfig) GetEnvironmentType() isProxyConfig_EnvironmentType {
	if m != nil {
		return m.EnvironmentType
	}
	return nil
}

func (x *ProxyConfig) GetKube() *KubernetesProxyConfig {
	if x, ok := x.GetEnvironmentType().(*ProxyConfig_Kube); ok {
		return x.Kube
	}
	return nil
}

type isProxyConfig_EnvironmentType interface {
	isProxyConfig_EnvironmentType()
}

type ProxyConfig_Kube struct {
	// The proxy will be deployed on Kubernetes.
	Kube *KubernetesProxyConfig `protobuf:"bytes,1,opt,name=kube,proto3,oneof"`
}

func (*ProxyConfig_Kube) isProxyConfig_EnvironmentType() {}

// Configuration for the set of Kubernetes resources that will be provisioned
// for a given Gateway.
type KubernetesProxyConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The workload type of the proxy
	//
	// Types that are assignable to WorkloadType:
	//
	//	*KubernetesProxyConfig_Deployment
	WorkloadType isKubernetesProxyConfig_WorkloadType `protobuf_oneof:"workload_type"`
	// Configuration for the Kubernetes Service that exposes the Envoy proxy over
	// the network.
	Service *kube.Service `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	// Autoscaling configuration.
	Autoscaling *kube.Autoscaling `protobuf:"bytes,3,opt,name=autoscaling,proto3" json:"autoscaling,omitempty"`
}

func (x *KubernetesProxyConfig) Reset() {
	*x = KubernetesProxyConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesProxyConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesProxyConfig) ProtoMessage() {}

func (x *KubernetesProxyConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesProxyConfig.ProtoReflect.Descriptor instead.
func (*KubernetesProxyConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_rawDescGZIP(), []int{2}
}

func (m *KubernetesProxyConfig) GetWorkloadType() isKubernetesProxyConfig_WorkloadType {
	if m != nil {
		return m.WorkloadType
	}
	return nil
}

func (x *KubernetesProxyConfig) GetDeployment() *ProxyDeployment {
	if x, ok := x.GetWorkloadType().(*KubernetesProxyConfig_Deployment); ok {
		return x.Deployment
	}
	return nil
}

func (x *KubernetesProxyConfig) GetService() *kube.Service {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *KubernetesProxyConfig) GetAutoscaling() *kube.Autoscaling {
	if x != nil {
		return x.Autoscaling
	}
	return nil
}

type isKubernetesProxyConfig_WorkloadType interface {
	isKubernetesProxyConfig_WorkloadType()
}

type KubernetesProxyConfig_Deployment struct {
	// Use a Kubernetes deployment as the proxy workload type.
	Deployment *ProxyDeployment `protobuf:"bytes,1,opt,name=deployment,proto3,oneof"`
}

func (*KubernetesProxyConfig_Deployment) isKubernetesProxyConfig_WorkloadType() {}

// Configuration for the Proxy deployment in Kubernetes.
type ProxyDeployment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of desired pods. Defaults to 1.
	Replicas *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=replicas,proto3" json:"replicas,omitempty"`
	// Configuration for the pods created by the deployment.
	PodTemplate *kube.Pod `protobuf:"bytes,2,opt,name=pod_template,json=podTemplate,proto3" json:"pod_template,omitempty"`
	// Configuration for the container running Envoy.
	EnvoyContainer *EnvoyContainer `protobuf:"bytes,3,opt,name=envoy_container,json=envoyContainer,proto3" json:"envoy_container,omitempty"`
}

func (x *ProxyDeployment) Reset() {
	*x = ProxyDeployment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyDeployment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyDeployment) ProtoMessage() {}

func (x *ProxyDeployment) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyDeployment.ProtoReflect.Descriptor instead.
func (*ProxyDeployment) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_rawDescGZIP(), []int{3}
}

func (x *ProxyDeployment) GetReplicas() *wrappers.UInt32Value {
	if x != nil {
		return x.Replicas
	}
	return nil
}

func (x *ProxyDeployment) GetPodTemplate() *kube.Pod {
	if x != nil {
		return x.PodTemplate
	}
	return nil
}

func (x *ProxyDeployment) GetEnvoyContainer() *EnvoyContainer {
	if x != nil {
		return x.EnvoyContainer
	}
	return nil
}

// Configuration for the container running Envoy.
type EnvoyContainer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The envoy container image. See
	// https://kubernetes.io/docs/concepts/containers/images
	// for details.
	Image *kube.Image `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	// Envoy log level. Options include "trace", "debug", "info", "warn", "error",
	// "critical" and "off". Defaults to "info". See
	// https://www.envoyproxy.io/docs/envoy/latest/start/quick-start/run-envoy#debugging-envoy
	// for more information.
	LogLevel string `protobuf:"bytes,2,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
	// Envoy component log level. The provided string will be used as a
	// passthrough value for the `--component-log-level` argument to Envoy, e.g.
	// `upstream:debug,connection:trace`. See
	// https://www.envoyproxy.io/docs/envoy/latest/start/quick-start/run-envoy#debugging-envoy
	// for more information.
	ComponentLogLevel string `protobuf:"bytes,3,opt,name=component_log_level,json=componentLogLevel,proto3" json:"component_log_level,omitempty"`
	// The security context for this container. See
	// https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.26/#securitycontext-v1-core
	// for details.
	SecurityContext *v1.SecurityContext `protobuf:"bytes,4,opt,name=security_context,json=securityContext,proto3" json:"security_context,omitempty"`
	// The compute resources required by this container. See
	// https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	// for details.
	Resources *kube.ResourceRequirements `protobuf:"bytes,5,opt,name=resources,proto3" json:"resources,omitempty"`
}

func (x *EnvoyContainer) Reset() {
	*x = EnvoyContainer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvoyContainer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvoyContainer) ProtoMessage() {}

func (x *EnvoyContainer) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvoyContainer.ProtoReflect.Descriptor instead.
func (*EnvoyContainer) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_rawDescGZIP(), []int{4}
}

func (x *EnvoyContainer) GetImage() *kube.Image {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *EnvoyContainer) GetLogLevel() string {
	if x != nil {
		return x.LogLevel
	}
	return ""
}

func (x *EnvoyContainer) GetComponentLogLevel() string {
	if x != nil {
		return x.ComponentLogLevel
	}
	return ""
}

func (x *EnvoyContainer) GetSecurityContext() *v1.SecurityContext {
	if x != nil {
		return x.SecurityContext
	}
	return nil
}

func (x *EnvoyContainer) GetResources() *kube.ResourceRequirements {
	if x != nil {
		return x.Resources
	}
	return nil
}

type DataPlaneConfigStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The current conditions of the DataPlaneConfig. This is not currently
	// implemented.
	Conditions []*v11.Condition `protobuf:"bytes,1,rep,name=conditions,proto3" json:"conditions,omitempty"`
}

func (x *DataPlaneConfigStatus) Reset() {
	*x = DataPlaneConfigStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataPlaneConfigStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataPlaneConfigStatus) ProtoMessage() {}

func (x *DataPlaneConfigStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataPlaneConfigStatus.ProtoReflect.Descriptor instead.
func (*DataPlaneConfigStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_rawDescGZIP(), []int{5}
}

func (x *DataPlaneConfigStatus) GetConditions() []*v11.Condition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

var File_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_rawDesc = []byte{
	0x0a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x14, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6b,
	0x75, 0x62, 0x65, 0x2f, 0x70, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x49, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69,
	0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x5d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x6f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x13, 0x44, 0x61, 0x74, 0x61,
	0x50, 0x6c, 0x61, 0x6e, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x44, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x64, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x41, 0x0a, 0x04, 0x6b, 0x75, 0x62, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48,
	0x00, 0x52, 0x04, 0x6b, 0x75, 0x62, 0x65, 0x42, 0x12, 0x0a, 0x10, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0xf9, 0x01, 0x0a, 0x15,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x47, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x48, 0x00, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3c,
	0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x0b,
	0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x41, 0x75,
	0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x52, 0x0b, 0x61, 0x75, 0x74, 0x6f, 0x73,
	0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x42, 0x0f, 0x0a, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0xdd, 0x01, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x08, 0x72,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x72, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x41, 0x0a, 0x0c, 0x70, 0x6f, 0x64, 0x5f, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6b, 0x75,
	0x62, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x6f, 0x64, 0x52, 0x0b, 0x70, 0x6f, 0x64,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x4d, 0x0a, 0x0f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x45, 0x6e, 0x76, 0x6f, 0x79, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x0e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x22, 0xb4, 0x02, 0x0a, 0x0e, 0x45, 0x6e, 0x76, 0x6f,
	0x79, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6b, 0x75, 0x62, 0x65,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x2e, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67,
	0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x4e, 0x0a, 0x10, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6b, 0x38, 0x73, 0x2e,
	0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0f,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x4d, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0x68,
	0x0a, 0x15, 0x44, 0x61, 0x74, 0x61, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4f, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6b, 0x38,
	0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72,
	0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x59, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5,
	0x04, 0x01, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x32, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_rawDescData = file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_goTypes = []interface{}{
	(*DataPlaneConfigSpec)(nil),       // 0: gateway.gloo.solo.io.DataPlaneConfigSpec
	(*ProxyConfig)(nil),               // 1: gateway.gloo.solo.io.ProxyConfig
	(*KubernetesProxyConfig)(nil),     // 2: gateway.gloo.solo.io.KubernetesProxyConfig
	(*ProxyDeployment)(nil),           // 3: gateway.gloo.solo.io.ProxyDeployment
	(*EnvoyContainer)(nil),            // 4: gateway.gloo.solo.io.EnvoyContainer
	(*DataPlaneConfigStatus)(nil),     // 5: gateway.gloo.solo.io.DataPlaneConfigStatus
	(*kube.Service)(nil),              // 6: kube.gateway.gloo.solo.io.Service
	(*kube.Autoscaling)(nil),          // 7: kube.gateway.gloo.solo.io.Autoscaling
	(*wrappers.UInt32Value)(nil),      // 8: google.protobuf.UInt32Value
	(*kube.Pod)(nil),                  // 9: kube.gateway.gloo.solo.io.Pod
	(*kube.Image)(nil),                // 10: kube.gateway.gloo.solo.io.Image
	(*v1.SecurityContext)(nil),        // 11: k8s.io.api.core.v1.SecurityContext
	(*kube.ResourceRequirements)(nil), // 12: kube.gateway.gloo.solo.io.ResourceRequirements
	(*v11.Condition)(nil),             // 13: k8s.io.apimachinery.pkg.apis.meta.v1.Condition
}
var file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_depIdxs = []int32{
	1,  // 0: gateway.gloo.solo.io.DataPlaneConfigSpec.proxy_config:type_name -> gateway.gloo.solo.io.ProxyConfig
	2,  // 1: gateway.gloo.solo.io.ProxyConfig.kube:type_name -> gateway.gloo.solo.io.KubernetesProxyConfig
	3,  // 2: gateway.gloo.solo.io.KubernetesProxyConfig.deployment:type_name -> gateway.gloo.solo.io.ProxyDeployment
	6,  // 3: gateway.gloo.solo.io.KubernetesProxyConfig.service:type_name -> kube.gateway.gloo.solo.io.Service
	7,  // 4: gateway.gloo.solo.io.KubernetesProxyConfig.autoscaling:type_name -> kube.gateway.gloo.solo.io.Autoscaling
	8,  // 5: gateway.gloo.solo.io.ProxyDeployment.replicas:type_name -> google.protobuf.UInt32Value
	9,  // 6: gateway.gloo.solo.io.ProxyDeployment.pod_template:type_name -> kube.gateway.gloo.solo.io.Pod
	4,  // 7: gateway.gloo.solo.io.ProxyDeployment.envoy_container:type_name -> gateway.gloo.solo.io.EnvoyContainer
	10, // 8: gateway.gloo.solo.io.EnvoyContainer.image:type_name -> kube.gateway.gloo.solo.io.Image
	11, // 9: gateway.gloo.solo.io.EnvoyContainer.security_context:type_name -> k8s.io.api.core.v1.SecurityContext
	12, // 10: gateway.gloo.solo.io.EnvoyContainer.resources:type_name -> kube.gateway.gloo.solo.io.ResourceRequirements
	13, // 11: gateway.gloo.solo.io.DataPlaneConfigStatus.conditions:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.Condition
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_init()
}
func file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_init() {
	if File_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataPlaneConfigSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesProxyConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyDeployment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvoyContainer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataPlaneConfigStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ProxyConfig_Kube)(nil),
	}
	file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*KubernetesProxyConfig_Deployment)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto = out.File
	file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_data_plane_config_proto_depIdxs = nil
}
