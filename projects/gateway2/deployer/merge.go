package deployer

import (
	kubecorev1 "github.com/solo-io/gloo/projects/gateway2/pkg/api/external/kubernetes/api/core/v1"
	"github.com/solo-io/gloo/projects/gateway2/pkg/api/gateway.gloo.solo.io/v1alpha1"
	"github.com/solo-io/gloo/projects/gateway2/pkg/api/gateway.gloo.solo.io/v1alpha1/kube"
)

// TODO(jbohanon) most of this code should probably generated, and almost none of it should
// live here

// mergePointers will decide whether to use dst or src without dereferencing or recursing
func mergePointers[T any](dst, src *T) *T {
	// nil src override means just use dst
	if src == nil {
		return dst
	}

	// given non-nil src override, use that instead
	return src
}

// deepMergeMaps will use dst if src is nil, src if dest is nil, or add all entries from src into dst
// if neither are nil
func deepMergeMaps[keyT comparable, valT any](dst, src map[keyT]valT) map[keyT]valT {
	// nil src override means just use dst
	if src == nil {
		return dst
	}

	if dst == nil || len(src) == 0 {
		return src
	}

	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func deepMergeSlices[T any](dst, src []T) []T {
	// nil src override means just use dst
	if src == nil {
		return dst
	}

	if dst == nil || len(src) == 0 {
		return src
	}

	dst = append(dst, src...)

	return dst
}

func deepMergeGatewayParameters(dst, src *v1alpha1.GatewayParameters) *v1alpha1.GatewayParameters {
	// nil src override means just use dst
	if src == nil || src.Spec.GetKube() == nil {
		return dst
	}

	if dst == nil || dst.Spec.GetKube() == nil {
		return src
	}

	dstKube := dst.Spec.GetKube()
	srcKube := src.Spec.GetKube()

	dstKube.EnvoyContainer = deepMergeEnvoyContainer(dstKube.GetEnvoyContainer(), srcKube.GetEnvoyContainer())

	dstKube.PodTemplate = deepMergePodTemplate(dstKube.GetPodTemplate(), srcKube.GetPodTemplate())

	dstKube.Service = deepMergeService(dstKube.GetService(), srcKube.GetService())

	dstKube.Autoscaling = deepMergeAutoscaling(dstKube.GetAutoscaling(), srcKube.GetAutoscaling())

	dstKube.Sds = deepMergeSdsIntegration(dstKube.GetSds(), srcKube.GetSds())

	if srcKube.GetWorkloadType() == nil {
		return dst
	}

	switch dstWorkload := dstKube.GetWorkloadType().(type) {
	case *v1alpha1.KubernetesProxyConfig_Deployment:
		srcWorkload, ok := srcKube.GetWorkloadType().(*v1alpha1.KubernetesProxyConfig_Deployment)
		if !ok {
			dstWorkload = srcWorkload
			break
		}
		dstWorkload = deepMergeDeploymentWorkloadType(dstWorkload, srcWorkload)
	default:
		// TODO(jbohanon) log or something? Shouldn't happen unless a new type is added
		break
	}

	return dst

}

func deepMergePodTemplate(dst, src *kube.Pod) *kube.Pod {
	// nil src override means just use dst
	if src == nil {
		return dst
	}

	if dst == nil {
		return src
	}

	dst.ExtraLabels = deepMergeMaps(dst.GetExtraLabels(), src.GetExtraLabels())

	dst.ExtraAnnotations = deepMergeMaps(dst.GetExtraAnnotations(), src.GetExtraAnnotations())

	dst.SecurityContext = deepMergePodSecurityContext(dst.GetSecurityContext(), src.GetSecurityContext())

	dst.ImagePullSecrets = deepMergeSlices(dst.GetImagePullSecrets(), src.GetImagePullSecrets())

	dst.NodeSelector = deepMergeMaps(dst.GetNodeSelector(), src.GetNodeSelector())

	dst.Affinity = deepMergeAffinity(dst.GetAffinity(), src.GetAffinity())

	dst.Tolerations = deepMergeSlices(dst.GetTolerations(), src.GetTolerations())

	return dst
}

func deepMergePodSecurityContext(dst, src *kubecorev1.PodSecurityContext) *kubecorev1.PodSecurityContext {
	// nil src override means just use dst
	if src == nil {
		return dst
	}

	if dst == nil {
		dst = src
		return src
	}

	dst.SeLinuxOptions = deepMergeSELinuxOptions(dst.GetSeLinuxOptions(), src.GetSeLinuxOptions())

	dst.WindowsOptions = deepMergeWindowsSecurityContextOptions(dst.GetWindowsOptions(), src.GetWindowsOptions())

	// We don't use getter here because getter returns zero value for nil, but we need
	// to know if it was nil
	dst.RunAsUser = mergePointers(dst.RunAsUser, src.RunAsUser)

	dst.RunAsGroup = mergePointers(dst.RunAsGroup, src.RunAsGroup)

	dst.RunAsNonRoot = mergePointers(dst.RunAsNonRoot, src.RunAsNonRoot)

	dst.SupplementalGroups = deepMergeSlices(dst.GetSupplementalGroups(), src.GetSupplementalGroups())

	dst.FsGroup = mergePointers(dst.FsGroup, src.FsGroup)

	dst.Sysctls = deepMergeSlices(dst.GetSysctls(), src.GetSysctls())

	dst.FsGroupChangePolicy = mergePointers(dst.FsGroupChangePolicy, src.FsGroupChangePolicy)

	dst.SeccompProfile = deepMergeSeccompProfile(dst.SeccompProfile, src.SeccompProfile)

	return dst
}

func deepMergeSELinuxOptions(dst, src *kubecorev1.SELinuxOptions) *kubecorev1.SELinuxOptions {
	// nil src override means just use dst
	if src == nil {
		return dst
	}

	if dst == nil {
		return src
	}

	dst.User = mergePointers(dst.User, src.User)
	dst.Role = mergePointers(dst.Role, src.Role)
	dst.Type = mergePointers(dst.Type, src.Type)
	dst.Level = mergePointers(dst.Level, src.Level)

	return dst
}
func deepMergeWindowsSecurityContextOptions(dst, src *kubecorev1.WindowsSecurityContextOptions) *kubecorev1.WindowsSecurityContextOptions {
	// nil src override means just use dst
	if src == nil {
		return dst
	}

	if dst == nil {
		dst = src
		return src
	}

	dst.GmsaCredentialSpecName = mergePointers(dst.GmsaCredentialSpecName, src.GmsaCredentialSpecName)
	dst.GmsaCredentialSpec = mergePointers(dst.GmsaCredentialSpec, src.GmsaCredentialSpec)
	dst.RunAsUserName = mergePointers(dst.RunAsUserName, src.RunAsUserName)
	dst.HostProcess = mergePointers(dst.HostProcess, src.HostProcess)

	return dst
}
func deepMergeSeccompProfile(dst, src *kubecorev1.SeccompProfile) *kubecorev1.SeccompProfile {
	// nil src override means just use dst
	if src == nil {
		return dst
	}

	if dst == nil {
		return src
	}

	dst.Type = mergePointers(dst.Type, src.Type)
	dst.LocalhostProfile = mergePointers(dst.LocalhostProfile, src.LocalhostProfile)

	return dst
}
func deepMergeAffinity(dst, src *kubecorev1.Affinity) *kubecorev1.Affinity {
	// nil src override means just use dst
	if src == nil {
		return dst
	}

	if dst == nil {
		return src
	}

	dst.NodeAffinity = deepMergeNodeAffinity(dst.GetNodeAffinity(), src.GetNodeAffinity())

	dst.PodAffinity = deepMergePodAffinity(dst.GetPodAffinity(), src.GetPodAffinity())

	dst.PodAntiAffinity = deepMergePodAntiAffinity(dst.GetPodAntiAffinity(), src.GetPodAntiAffinity())

	return dst
}

func deepMergeNodeAffinity(dst, src *kubecorev1.NodeAffinity) *kubecorev1.NodeAffinity {
	// nil src override means just use dst
	if src == nil {
		return dst
	}

	if dst == nil {
		return src
	}

	dst.RequiredDuringSchedulingIgnoredDuringExecution = deepMergeNodeSelector(dst.GetRequiredDuringSchedulingIgnoredDuringExecution(), src.GetRequiredDuringSchedulingIgnoredDuringExecution())

	dst.PreferredDuringSchedulingIgnoredDuringExecution = deepMergeSlices(dst.GetPreferredDuringSchedulingIgnoredDuringExecution(), src.GetPreferredDuringSchedulingIgnoredDuringExecution())

	return dst
}
func deepMergeNodeSelector(dst, src *kubecorev1.NodeSelector) *kubecorev1.NodeSelector {
	// nil src override means just use dst
	if src == nil {
		return dst
	}

	if dst == nil {
		return src
	}

	dst.NodeSelectorTerms = deepMergeSlices(dst.GetNodeSelectorTerms(), src.GetNodeSelectorTerms())

	return dst
}
func deepMergePodAffinity(dst, src *kubecorev1.PodAffinity) *kubecorev1.PodAffinity {
	// nil src override means just use dst
	if src == nil {
		return dst
	}

	if dst == nil {
		return src
	}

	dst.RequiredDuringSchedulingIgnoredDuringExecution = deepMergeSlices(dst.GetRequiredDuringSchedulingIgnoredDuringExecution(), src.GetRequiredDuringSchedulingIgnoredDuringExecution())

	dst.PreferredDuringSchedulingIgnoredDuringExecution = deepMergeSlices(dst.GetPreferredDuringSchedulingIgnoredDuringExecution(), src.GetPreferredDuringSchedulingIgnoredDuringExecution())

	return dst
}
func deepMergePodAntiAffinity(dst, src *kubecorev1.PodAntiAffinity) *kubecorev1.PodAntiAffinity {
	// nil src override means just use dst
	if src == nil {
		return dst
	}

	if dst == nil {
		return src
	}

	dst.RequiredDuringSchedulingIgnoredDuringExecution = deepMergeSlices(dst.GetRequiredDuringSchedulingIgnoredDuringExecution(), src.GetRequiredDuringSchedulingIgnoredDuringExecution())

	dst.PreferredDuringSchedulingIgnoredDuringExecution = deepMergeSlices(dst.GetPreferredDuringSchedulingIgnoredDuringExecution(), src.GetPreferredDuringSchedulingIgnoredDuringExecution())

	return dst
}
func deepMergeService(dst, src *kube.Service) *kube.Service {
	// nil src override means just use dst
	if src == nil {
		return dst
	}

	if dst == nil {
		return src
	}

	// This is not nullable and as such is a required field for overrides
	// TODO(jbohanon) note this for documentation
	dst.Type = src.GetType()

	// This is not nullable and as such is a required field for overrides
	// TODO(jbohanon) note this for documentation
	dst.ClusterIP = src.GetClusterIP()

	dst.ExtraLabels = deepMergeMaps(dst.GetExtraLabels(), src.GetExtraLabels())

	dst.ExtraAnnotations = deepMergeMaps(dst.GetExtraAnnotations(), src.GetExtraAnnotations())

	return dst
}
func deepMergeAutoscaling(dst, src *kube.Autoscaling) *kube.Autoscaling {
	// nil src override means just use dst
	if src == nil {
		return dst
	}

	if dst == nil {
		return src
	}

	dst.HorizontalPodAutoscaler = deepMergeHorizontalPodAutoscaler(dst.GetHorizontalPodAutoscaler(), src.GetHorizontalPodAutoscaler())

	return dst
}
func deepMergeHorizontalPodAutoscaler(dst, src *kube.HorizontalPodAutoscaler) *kube.HorizontalPodAutoscaler {
	// nil src override means just use dst
	if src == nil {
		return dst
	}

	if dst == nil {
		return src
	}

	dst.MinReplicas = mergePointers(dst.GetMinReplicas(), src.GetMinReplicas())
	dst.MaxReplicas = mergePointers(dst.GetMaxReplicas(), src.GetMaxReplicas())
	dst.TargetCpuUtilizationPercentage = mergePointers(dst.GetTargetCpuUtilizationPercentage(), src.GetTargetCpuUtilizationPercentage())
	dst.TargetMemoryUtilizationPercentage = mergePointers(dst.GetTargetMemoryUtilizationPercentage(), src.GetTargetMemoryUtilizationPercentage())

	return dst
}
func deepMergeSdsIntegration(dst, src *v1alpha1.SdsIntegration) *v1alpha1.SdsIntegration {
	// nil src override means just use dst
	if src == nil {
		return dst
	}

	if dst == nil {
		return src
	}

	// SdsContainer, IstioIntegration
	dst.SdsContainer = deepMergeSdsContainer(dst.GetSdsContainer(), src.GetSdsContainer())
	dst.IstioIntegration = deepMergeIstioIntegration(dst.GetIstioIntegration(), src.GetIstioIntegration())

	return dst
}
func deepMergeSdsContainer(dst, src *v1alpha1.SdsContainer) *v1alpha1.SdsContainer {
	// nil src override means just use dst
	if src == nil {
		return dst
	}

	if dst == nil {
		return src
	}

	dst.Image = deepMergeImage(dst.GetImage(), src.GetImage())
	dst.SecurityContext = deepMergeSecurityContext(dst.GetSecurityContext(), src.GetSecurityContext())
	dst.Resources = deepMergeResourceRequirements(dst.GetResources(), src.GetResources())
	dst.Bootstrap = deepMergeSdsBootstrap(dst.GetBootstrap(), src.GetBootstrap())

	return dst
}
func deepMergeSdsBootstrap(dst, src *v1alpha1.SdsBootstrap) *v1alpha1.SdsBootstrap {
	// nil src override means just use dst
	if src == nil {
		return dst
	}

	if dst == nil {
		return src
	}

	if src.GetLogLevel() != "" {
		dst.LogLevel = src.GetLogLevel()
	}

	return dst
}
func deepMergeIstioIntegration(dst, src *v1alpha1.IstioIntegration) *v1alpha1.IstioIntegration {
	// nil src override means just use dst
	if src == nil {
		return dst
	}

	if dst == nil {
		return src
	}

	dst.IstioContainer = deepMergeIstioContainer(dst.GetIstioContainer(), src.GetIstioContainer())

	if istioDiscoveryAddress := src.GetIstioDiscoveryAddress(); istioDiscoveryAddress != "" {
		dst.IstioDiscoveryAddress = istioDiscoveryAddress
	}
	if istioMetaMeshId := src.GetIstioMetaMeshId(); istioMetaMeshId != "" {
		dst.IstioMetaMeshId = istioMetaMeshId
	}
	if istioMetaClusterId := src.GetIstioMetaClusterId(); istioMetaClusterId != "" {
		dst.IstioMetaClusterId = istioMetaClusterId
	}

	return dst
}
func deepMergeIstioContainer(dst, src *v1alpha1.IstioContainer) *v1alpha1.IstioContainer {
	// nil src override means just use dst
	if src == nil {
		return dst
	}

	if dst == nil {
		return src
	}
	/*
		Image *kube.Image `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
		SecurityContext *v1.SecurityContext `protobuf:"bytes,2,opt,name=security_context,json=securityContext,proto3" json:"security_context,omitempty"`
		Resources *kube.ResourceRequirements `protobuf:"bytes,3,opt,name=resources,proto3" json:"resources,omitempty"`
		LogLevel string `protobuf:"bytes,4,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
	*/

	dst.Image = deepMergeImage(dst.GetImage(), src.GetImage())
	dst.SecurityContext = deepMergeSecurityContext(dst.GetSecurityContext(), src.GetSecurityContext())
	dst.Resources = deepMergeResourceRequirements(dst.GetResources(), src.GetResources())

	if logLevel := src.GetLogLevel(); logLevel != "" {
		dst.LogLevel = logLevel
	}

	return dst
}
func deepMergeEnvoyContainer(dst, src *v1alpha1.EnvoyContainer) *v1alpha1.EnvoyContainer {
	// nil src override means just use dst
	if src == nil {
		return dst
	}

	if dst == nil {
		return src
	}

	dst.Image = deepMergeImage(dst.GetImage(), src.GetImage())

	dst.Bootstrap = deepMergeEnvoyBootstrap(dst.GetBootstrap(), src.GetBootstrap())

	dst.Resources = deepMergeResourceRequirements(dst.GetResources(), src.GetResources())

	dst.SecurityContext = deepMergeSecurityContext(dst.GetSecurityContext(), src.GetSecurityContext())

	return dst
}
func deepMergeImage(dst, src *kube.Image) *kube.Image {
	// nil src override means just use dst
	if src == nil {
		return dst
	}

	if dst == nil {
		return src
	}

	// because all fields are not nullable, we treat empty strings as empty values
	// and do not override with them

	if src.GetRegistry() != "" {
		dst.Registry = src.GetRegistry()
	}

	if src.GetRepository() != "" {
		dst.Repository = src.GetRepository()
	}

	if src.GetTag() != "" {
		dst.Tag = src.GetTag()
	}

	if src.GetDigest() != "" {
		dst.Digest = src.GetDigest()
	}

	dst.PullPolicy = src.GetPullPolicy()

	return dst
}

func deepMergeEnvoyBootstrap(dst, src *v1alpha1.EnvoyBootstrap) *v1alpha1.EnvoyBootstrap {
	// nil src override means just use dst
	if src == nil {
		return dst
	}

	if dst == nil {
		return src
	}
	if src.GetLogLevel() != "" {
		dst.LogLevel = src.GetLogLevel()
	}

	dst.ComponentLogLevels = deepMergeMaps(dst.GetComponentLogLevels(), src.GetComponentLogLevels())

	return dst
}
func deepMergeResourceRequirements(dst, src *kube.ResourceRequirements) *kube.ResourceRequirements {
	// nil src override means just use dst
	if src == nil {
		return dst
	}

	if dst == nil {
		return src
	}

	dst.Limits = deepMergeMaps(dst.GetLimits(), src.GetLimits())

	dst.Requests = deepMergeMaps(dst.GetRequests(), src.GetRequests())

	return dst
}
func deepMergeSecurityContext(dst, src *kubecorev1.SecurityContext) *kubecorev1.SecurityContext {
	// nil src override means just use dst
	if src == nil {
		return dst
	}

	if dst == nil {
		return src
	}

	dst.Capabilities = deepMergeCapabilities(dst.GetCapabilities(), src.GetCapabilities())

	dst.SeLinuxOptions = deepMergeSELinuxOptions(dst.GetSeLinuxOptions(), src.GetSeLinuxOptions())

	dst.WindowsOptions = deepMergeWindowsSecurityContextOptions(dst.GetWindowsOptions(), src.GetWindowsOptions())

	// We don't use getter here because getter returns zero value for nil, but we need
	// to know if it was nil
	dst.RunAsUser = mergePointers(dst.RunAsUser, src.RunAsUser)

	dst.RunAsGroup = mergePointers(dst.RunAsGroup, src.RunAsGroup)

	dst.RunAsNonRoot = mergePointers(dst.RunAsNonRoot, src.RunAsNonRoot)

	dst.Privileged = mergePointers(dst.Privileged, src.Privileged)

	dst.ReadOnlyRootFilesystem = mergePointers(dst.ReadOnlyRootFilesystem, src.ReadOnlyRootFilesystem)

	dst.AllowPrivilegeEscalation = mergePointers(dst.AllowPrivilegeEscalation, src.AllowPrivilegeEscalation)

	dst.ProcMount = mergePointers(dst.ProcMount, src.ProcMount)

	dst.SeccompProfile = deepMergeSeccompProfile(dst.GetSeccompProfile(), src.GetSeccompProfile())

	return dst
}
func deepMergeCapabilities(dst, src *kubecorev1.Capabilities) *kubecorev1.Capabilities {
	// nil src override means just use dst
	if src == nil {
		return dst
	}

	if dst == nil {
		return src
	}

	dst.Add = deepMergeSlices(dst.GetAdd(), src.GetAdd())
	dst.Drop = deepMergeSlices(dst.GetDrop(), src.GetDrop())

	return dst
}
func deepMergeDeploymentWorkloadType(dst, src *v1alpha1.KubernetesProxyConfig_Deployment) *v1alpha1.KubernetesProxyConfig_Deployment {
	// nil src override means just use dst
	if src == nil {
		return dst
	}

	if dst == nil {
		return src
	}

	dstDeployment := dst.Deployment
	srcDeployment := src.Deployment

	if srcDeployment == nil {
		return dst
	}
	if dstDeployment == nil {
		return src
	}

	// we can use the getter here since the value is a pb wrapper
	dstDeployment.Replicas = mergePointers(dst.Deployment.GetReplicas(), src.Deployment.GetReplicas())

	return dst
}

// The following exists only to exclude this file from the gettercheck.
// This is a hacky workaround due to needing to pass pointers into `mergePointers` by field
// access instead of by getter. We should add a way to exclude lines from the gettercheck.

// Code generated DO NOT EDIT.
