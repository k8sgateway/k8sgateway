package version

import (
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/grpc/version"
	"github.com/solo-io/go-utils/kubeutils"
	"github.com/solo-io/go-utils/stringutils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ServerVersion interface {
	Get(opts *options.Options) (*version.ServerVersion, error)
}

type kube struct{}

var (
	KnativeUniqueContainers = []string{"knative-external-proxy", "knative-internal-proxy"}
	IngressUniqueContainers = []string{"ingress"}
	GlooEUniqueContainers = []string{"gloo-ee"}
)

func NewKube() *kube {
	return &kube{}
}

func (k *kube) Get(opts *options.Options) (*version.ServerVersion, error) {
	cfg, err := kubeutils.GetConfig("", "")
	if err != nil {
		// kubecfg is missing, therefore no cluster is present, only print client version
		return nil, nil
	}
	client, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}

	deployments, err := client.AppsV1().Deployments(opts.Metadata.Namespace).List(metav1.ListOptions{
		// search only for gloo deployments based on labels
		LabelSelector: "app=gloo",
	})
	if err != nil {
		return nil, err
	}

	var kubeContainerList []*version.Kubernetes_Container
	var foundGlooE, foundIngress, foundKnative bool
	for _, v := range deployments.Items {
		for _, container := range v.Spec.Template.Spec.Containers {
			containerInfo := parseContainerString(container)
			kubeContainerList = append(kubeContainerList, &version.Kubernetes_Container{
				Tag:      containerInfo.Tag,
				Name:     containerInfo.Repository,
				Registry: containerInfo.Registry,
			})
			switch {
			case stringutils.ContainsString(containerInfo.Repository, KnativeUniqueContainers):
				foundKnative = true
			case stringutils.ContainsString(containerInfo.Repository, IngressUniqueContainers):
				foundIngress = true
			case stringutils.ContainsString(containerInfo.Repository, GlooEUniqueContainers):
				foundGlooE = true
			}
		}
	}

	var deploymentType version.GlooType
	switch {
	case foundGlooE:
		deploymentType = version.GlooType_Enterprise
	case foundKnative:
		deploymentType = version.GlooType_Knative
	case foundIngress:
		deploymentType = version.GlooType_Knative
	default:
		deploymentType = version.GlooType_Gateway
	}

	if len(kubeContainerList) == 0 {
		return nil, nil
	}
	serverVersion := &version.ServerVersion{
		VersionType: &version.ServerVersion_Kubernetes{
			Kubernetes: &version.Kubernetes{
				Containers: kubeContainerList,
				Namespace:  opts.Metadata.Namespace,
				Type:       deploymentType,
			},
		},
	}
	return serverVersion, nil
}
