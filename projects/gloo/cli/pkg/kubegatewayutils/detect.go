package kubegatewayutils

import (
	"context"
	"strconv"

	"github.com/rotisserie/eris"
	"github.com/solo-io/gloo/pkg/utils/kubeutils"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	cliconstants "github.com/solo-io/gloo/projects/gloo/cli/pkg/constants"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"
	"github.com/solo-io/gloo/projects/gloo/constants"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"
)

const (
	KubeGatewayNetworkingGroup = "gateway.networking.k8s.io"
)

// Returns true if Kubernetes Gateway API CRDs are on the cluster.
// Note: this doesn't check for specific CRD names; it returns true if *any* k8s Gateway CRD is detected
func DetectKubeGatewayCrds(cfg *rest.Config) (bool, error) {
	discClient, err := discovery.NewDiscoveryClientForConfig(cfg)
	if err != nil {
		return false, err
	}

	groups, err := discClient.ServerGroups()
	if err != nil {
		return false, err
	}

	// Check if gateway group exists
	for _, group := range groups.Groups {
		if group.Name == KubeGatewayNetworkingGroup {
			return true, nil
		}
	}

	return false, nil
}

// Returns true if the GG_EXPERIMENTAL_K8S_GW_CONTROLLER env var is true in the gloo deployment.
// Note: This is tied up with the GG implementation and will need to be updated if it changes
func DetectKubeGatewayEnabled(ctx context.Context, opts *options.Options) (bool, error) {
	// check if kube gateway integration is enabled by checking if the controller env variable is set in the
	// gloo deployment
	client, err := helpers.GetKubernetesClient(opts.Top.KubeContext)
	if err != nil {
		return false, eris.Wrapf(err, "could not get kubernetes client")
	}

	glooDeployment, err := client.AppsV1().Deployments(opts.Metadata.GetNamespace()).Get(ctx, kubeutils.GlooDeploymentName, metav1.GetOptions{})
	if err != nil {
		return false, eris.Wrapf(err, "could not get gloo deployment")
	}

	var glooContainer *corev1.Container
	for _, container := range glooDeployment.Spec.Template.Spec.Containers {
		if container.Name == cliconstants.GlooContainerName {
			glooContainer = &container
			break
		}
	}
	if glooContainer == nil {
		return false, eris.New("could not find gloo container in gloo deployment")
	}

	for _, envVar := range glooContainer.Env {
		if envVar.Name == constants.GlooGatewayEnableK8sGwControllerEnv {
			val, err := strconv.ParseBool(envVar.Value)
			if err != nil {
				return false, eris.Wrapf(err, "could not parse value of %s env var in gloo deployment", constants.GlooGatewayEnableK8sGwControllerEnv)
			}
			return val, nil
		}
	}
	return false, nil
}
