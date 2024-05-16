package deployer

import (
	"fmt"
	"sort"
	"strings"

	"github.com/rotisserie/eris"
	"github.com/solo-io/gloo/projects/gateway2/extensions"
	"github.com/solo-io/gloo/projects/gateway2/pkg/api/gateway.gloo.solo.io/v1alpha1"
	v1alpha1kube "github.com/solo-io/gloo/projects/gateway2/pkg/api/gateway.gloo.solo.io/v1alpha1/kube"
	"github.com/solo-io/gloo/projects/gateway2/ports"
	"golang.org/x/exp/slices"
	"k8s.io/utils/ptr"
	api "sigs.k8s.io/gateway-api/apis/v1"
)

// This file contains helper functions that generate helm values in the format needed
// by the deployer.

var (
	ComponentLogLevelEmptyError = func(key string, value string) error {
		return eris.Errorf("an empty key or value was provided in componentLogLevels: key=%s, value=%s", key, value)
	}
)

// Extract the listener ports from a Gateway. These will be used to populate:
// 1. the ports exposed on the envoy container
// 2. the ports exposed on the proxy service
func getPortsValues(gw *api.Gateway) []helmPort {
	gwPorts := []helmPort{}
	for _, l := range gw.Spec.Listeners {
		listenerPort := uint16(l.Port)

		// only process this port if we haven't already processed a listener with the same port
		if slices.IndexFunc(gwPorts, func(p helmPort) bool { return *p.Port == listenerPort }) != -1 {
			continue
		}

		targetPort := ports.TranslatePort(listenerPort)
		portName := string(l.Name)
		protocol := "TCP"

		gwPorts = append(gwPorts, helmPort{
			Port:       &listenerPort,
			TargetPort: &targetPort,
			Name:       &portName,
			Protocol:   &protocol,
		})
	}
	return gwPorts
}

// Convert autoscaling values from GatewayParameters into helm values to be used by the deployer.
func getAutoscalingValues(autoscaling *v1alpha1kube.Autoscaling) *helmAutoscaling {
	hpaConfig := autoscaling.GetHorizontalPodAutoscaler()
	if hpaConfig == nil {
		return nil
	}

	trueVal := true
	autoscalingVals := &helmAutoscaling{
		Enabled: &trueVal,
	}
	if hpaConfig.GetMinReplicas() != nil {
		minReplicas := hpaConfig.GetMinReplicas().GetValue()
		autoscalingVals.MinReplicas = &minReplicas
	}
	if hpaConfig.GetMaxReplicas() != nil {
		maxReplicas := hpaConfig.GetMaxReplicas().GetValue()
		autoscalingVals.MaxReplicas = &maxReplicas
	}
	if hpaConfig.GetTargetCpuUtilizationPercentage() != nil {
		cpuPercent := hpaConfig.GetTargetCpuUtilizationPercentage().GetValue()
		autoscalingVals.TargetCPUUtilizationPercentage = &cpuPercent
	}
	if hpaConfig.GetTargetMemoryUtilizationPercentage() != nil {
		memPercent := hpaConfig.GetTargetMemoryUtilizationPercentage().GetValue()
		autoscalingVals.TargetMemoryUtilizationPercentage = &memPercent
	}

	return autoscalingVals
}

// Convert service values from GatewayParameters into helm values to be used by the deployer.
func getServiceValues(svcConfig *v1alpha1kube.Service) *helmService {
	// convert the service type enum to its string representation;
	// if type is not set, it will default to 0 ("ClusterIP")
	svcType := v1alpha1kube.Service_ServiceType_name[int32(svcConfig.GetType())]
	clusterIp := svcConfig.GetClusterIP()
	return &helmService{
		Type:             &svcType,
		ClusterIP:        &clusterIp,
		ExtraAnnotations: svcConfig.GetExtraAnnotations(),
		ExtraLabels:      svcConfig.GetExtraLabels(),
	}
}

func getDefaultSdsValues(defaultSds extensions.Image) *helmSds {
	defaultSdsImage := helmImage{
		Repository: ptr.To(defaultSds.Repository),
		Tag:        ptr.To(defaultSds.Tag),
	}

	return &helmSds{
		Image: ptr.To(defaultSdsImage),
		SdsBootstrap: &sdsBootstrap{
			LogLevel: ptr.To("info"),
		},
	}
}

// Convert sds values from GatewayParameters into helm values to be used by the deployer.
func getSdsValues(sdsConfig *v1alpha1.SdsIntegration, defaultSds extensions.Image) *helmSds {
	// if sdsConfig is nil, sds is disabled
	if sdsConfig == nil {
		return nil
	}

	var sds *helmSds
	// if sdsConfig is not nil, but unset use defaults for sds config
	if sdsConfig.GetSdsContainer() == nil {
		sds = getDefaultSdsValues(defaultSds)
		sds.Istio = getIstioValues(sdsConfig.GetIstioIntegration())
	} else {
		// Use GatewayParameter overrides if provided
		var bootstrap *sdsBootstrap
		if sdsConfig.GetSdsContainer() != nil {
			bootstrap = &sdsBootstrap{
				LogLevel: ptr.To(sdsConfig.GetSdsContainer().GetBootstrap().GetLogLevel()),
			}
		}

		// convert the service type enum to its string representation;
		// if type is not set, it will default to 0 ("ClusterIP")
		sds = &helmSds{
			Image:           getMergedSdsImageValues(defaultSds, sdsConfig.GetSdsContainer().GetImage()),
			Resources:       sdsConfig.GetSdsContainer().GetResources(),
			SecurityContext: sdsConfig.GetSdsContainer().GetSecurityContext(),
			SdsBootstrap:    bootstrap,
			Istio:           getIstioValues(sdsConfig.GetIstioIntegration()),
		}
	}
	return sds
}

// Convert istio values from GatewayParameters into helm values to be used by the deployer.
func getIstioValues(istioConfig *v1alpha1.IstioIntegration) *helmIstioSds {
	var istioVals *helmIstioSds

	defaultIstioImage := helmImage{
		Registry:   ptr.To("docker.io/istio"),
		Repository: ptr.To("proxyv2"),
		Tag:        ptr.To("1.19.0"),
	}

	// if istioConfig is nil, istio sds is disabled and values can be ignored
	if istioConfig != nil && istioConfig.GetEnabled() != nil && !istioConfig.GetEnabled().GetValue() {
		// If istioConfig is not nil, but unset use defaults for istio config
		if istioConfig.GetIstioContainer() == nil {
			istioVals = &helmIstioSds{
				Enabled:  ptr.To(istioConfig.GetEnabled().GetValue()),
				Image:    ptr.To(defaultIstioImage),
				LogLevel: ptr.To("warning"),
			}
		} else {
			// Use GatewayParameter overrides if provided
			istioVals = &helmIstioSds{
				Enabled:         ptr.To(istioConfig.GetEnabled().GetValue()),
				Image:           getImage(istioConfig.GetIstioContainer().GetImage()),
				LogLevel:        ptr.To(istioConfig.GetIstioContainer().GetLogLevel()),
				Resources:       istioConfig.GetIstioContainer().GetResources(),
				SecurityContext: istioConfig.GetIstioContainer().GetSecurityContext(),
			}
		}
	}
	return istioVals
}

// Convert an image from GatewayParameters into a helm image to be used by the deployer.
func getImage(overrideImage *v1alpha1kube.Image) *helmImage {
	// if no overrides are provided, use the default values
	if overrideImage == nil {
		return nil
	}

	image := &helmImage{}

	if overrideImage.GetRepository() != "" {
		image.Repository = ptr.To(overrideImage.GetRepository())
	}
	if overrideImage.GetTag() != "" {
		image.Tag = ptr.To(overrideImage.GetTag())
	}
	if overrideImage.GetRegistry() != "" {
		image.Registry = ptr.To(overrideImage.GetRegistry())
	}
	if overrideImage.GetDigest() != "" {
		image.Digest = ptr.To(overrideImage.GetDigest())
	}

	// get the string representation of pull policy, unless it's unspecified, in which case we
	// leave it empty to fall back to the default value
	pullPolicy := ""
	if overrideImage.GetPullPolicy() != v1alpha1kube.Image_Unspecified {
		pullPolicy = v1alpha1kube.Image_PullPolicy_name[int32(overrideImage.GetPullPolicy())]
	}
	image.PullPolicy = ptr.To(pullPolicy)

	return image
}

// Get the default image values for the envoy container in the proxy deployment.
// Typically this is a gloo envoy wrapper image.
func getDefaultEnvoyImageValues(image extensions.Image) *helmImage {
	// Get the envoy repo and tag from the k8s gw extensions.
	// The other default values (registry and pullPolicy) are statically defined in the deployer
	// helm chart.
	return &helmImage{
		Repository: &image.Repository,
		Tag:        &image.Tag,
	}
}

// Get the image values for the sds container in the proxy deployment. This is done by:
// 1. getting the image values from a GatewayParameter
// 2. for values not provided, fall back to the defaults (if any) from the k8s gw extensions
func getMergedSdsImageValues(defaultImage extensions.Image, overrideImage *v1alpha1kube.Image) *helmImage {
	// if no overrides are provided, use the default values
	if overrideImage == nil {
		return &helmImage{
			Repository: ptr.To(defaultImage.Repository),
			Tag:        ptr.To(defaultImage.Tag),
		}
	}

	// for repo and tag, fall back to defaults if not provided
	repository := overrideImage.GetRepository()
	if repository == "" {
		repository = defaultImage.Repository
	}
	tag := overrideImage.GetTag()
	if tag == "" {
		tag = defaultImage.Tag
	}

	registry := overrideImage.GetRegistry()
	digest := overrideImage.GetDigest()

	// get the string representation of pull policy, unless it's unspecified, in which case we
	// leave it empty to fall back to the default value
	pullPolicy := ""
	if overrideImage.GetPullPolicy() != v1alpha1kube.Image_Unspecified {
		pullPolicy = v1alpha1kube.Image_PullPolicy_name[int32(overrideImage.GetPullPolicy())]
	}

	return &helmImage{
		Registry:   &registry,
		Repository: &repository,
		Tag:        &tag,
		Digest:     &digest,
		PullPolicy: &pullPolicy,
	}
}

// Get the image values for the envoy container in the proxy deployment. This is done by:
// 1. getting the image values from a GatewayParameter
// 2. for values not provided, fall back to the defaults (if any) from the k8s gw extensions
func getMergedEnvoyImageValues(defaultImage extensions.Image, overrideImage *v1alpha1kube.Image) *helmImage {
	// if no overrides are provided, use the default values
	if overrideImage == nil {
		return getDefaultEnvoyImageValues(defaultImage)
	}

	// for repo and tag, fall back to defaults if not provided
	repository := overrideImage.GetRepository()
	if repository == "" {
		repository = defaultImage.Repository
	}
	tag := overrideImage.GetTag()
	if tag == "" {
		tag = defaultImage.Tag
	}

	registry := overrideImage.GetRegistry()
	digest := overrideImage.GetDigest()

	// get the string representation of pull policy, unless it's unspecified, in which case we
	// leave it empty to fall back to the default value
	pullPolicy := ""
	if overrideImage.GetPullPolicy() != v1alpha1kube.Image_Unspecified {
		pullPolicy = v1alpha1kube.Image_PullPolicy_name[int32(overrideImage.GetPullPolicy())]
	}

	return &helmImage{
		Registry:   &registry,
		Repository: &repository,
		Tag:        &tag,
		Digest:     &digest,
		PullPolicy: &pullPolicy,
	}
}

// ComponentLogLevelsToString converts the key-value pairs in the map into a string of the
// format: key1:value1,key2:value2,key3:value3, where the keys are sorted alphabetically.
// If an empty map is passed in, then an empty string is returned.
// Map keys and values may not be empty.
// No other validation is currently done on the keys/values.
func ComponentLogLevelsToString(vals map[string]string) (string, error) {
	if len(vals) == 0 {
		return "", nil
	}

	parts := make([]string, 0, len(vals))
	for k, v := range vals {
		if k == "" || v == "" {
			return "", ComponentLogLevelEmptyError(k, v)
		}
		parts = append(parts, fmt.Sprintf("%s:%s", k, v))
	}
	sort.Strings(parts)
	return strings.Join(parts, ","), nil
}
