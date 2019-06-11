package upstreams_test

import (
	"fmt"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/kubernetes"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	skkube "github.com/solo-io/solo-kit/pkg/api/v1/resources/common/kubernetes"
	corev1 "k8s.io/api/core/v1"
)

var T *testing.T

func TestClients(t *testing.T) {
	RegisterFailHandler(Fail)
	T = t
	RunSpecs(t, "Hybrid Upstreams Client Suite")
}

var getService = func(name, namespace, version string, ports []int32) *skkube.Service {
	svc := skkube.NewService("svc-ns-1", "svc-1")
	svc.ObjectMeta.ResourceVersion = version
	svc.Spec = corev1.ServiceSpec{}
	for i, port := range ports {
		svc.Spec.Ports = append(svc.Spec.Ports, corev1.ServicePort{
			Name: fmt.Sprintf("port-%d", i),
			Port: port,
		})
	}
	return svc
}

var getUpstream = func(name, namespace, svcName, svcNs string, port uint32) *v1.Upstream {
	return &v1.Upstream{
		Metadata: core.Metadata{
			Name:      name,
			Namespace: namespace,
		},
		UpstreamSpec: &v1.UpstreamSpec{
			UpstreamType: &v1.UpstreamSpec_Kube{
				Kube: &kubernetes.UpstreamSpec{
					ServiceName:      svcName,
					ServiceNamespace: svcNs,
					ServicePort:      port,
				}},
		},
	}
}
