package kubeutils

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/solo-io/skv2/pkg/multicluster/kubeconfig"
)

// Inspired by: https://github.com/solo-io/gloo-mesh-enterprise/blob/main/pkg/utils/kubeutils/pods.go

// GetPodsForDeployment gets all pods backing a deployment
func GetPodsForDeployment(
	ctx context.Context,
	kubeConfig string,
	kubeContext string,
	deploymentName string,
	deploymentNamespace string,
) ([]string, error) {
	config, err := kubeconfig.GetRestConfigWithContext(kubeConfig, kubeContext, "")
	if err != nil {
		return nil, err
	}
	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	deployment, err := kubeClient.AppsV1().Deployments(deploymentNamespace).Get(ctx, deploymentName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	matchLabels := deployment.Spec.Selector.MatchLabels
	listOptions := (&client.ListOptions{
		LabelSelector: labels.SelectorFromSet(matchLabels),
		FieldSelector: fields.Set{"status.phase": "Running"}.AsSelector(),
	}).AsListOptions()

	podList, err := kubeClient.CoreV1().Pods(deploymentNamespace).List(ctx, *listOptions)
	if err != nil {
		return nil, err
	}

	pods := make([]string, len(podList.Items))
	for i := range podList.Items {
		pods[i] = podList.Items[i].Name
	}

	return pods, nil
}

// GetPodsForService gets all pods backing a deployment
func GetPodsForService(
	ctx context.Context,
	kubeConfig string,
	kubeContext string,
	serviceName string,
	serviceNamespace string,
) ([]string, error) {
	config, err := kubeconfig.GetRestConfigWithContext(kubeConfig, kubeContext, "")
	if err != nil {
		return nil, err
	}
	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	service, err := kubeClient.CoreV1().Services(serviceNamespace).Get(ctx, serviceName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	matchLabels := service.Spec.Selector
	listOptions := (&client.ListOptions{LabelSelector: labels.SelectorFromSet(matchLabels)}).AsListOptions()

	podList, err := kubeClient.CoreV1().Pods(serviceNamespace).List(ctx, *listOptions)
	if err != nil {
		return nil, err
	}

	pods := make([]string, len(podList.Items))
	for i := range podList.Items {
		pods[i] = podList.Items[i].Name
	}

	return pods, nil
}
