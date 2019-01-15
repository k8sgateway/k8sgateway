package kube2e_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/test/helpers"
	"github.com/solo-io/go-utils/kubeutils"
	"github.com/solo-io/solo-kit/test/setup"
	"io/ioutil"
	"k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"path/filepath"
)

var _ = FDescribe("Kube2e: Knative-Ingress", func() {
	BeforeEach(func() {
		deployKnative()
	})
	AfterEach(func() {
		deleteKnative()
	})
	It("works", func() {
		cfg, err := kubeutils.GetConfig("", "")
		Expect(err).NotTo(HaveOccurred())

		kube, err := kubernetes.NewForConfig(cfg)
		Expect(err).NotTo(HaveOccurred())
		kubeIngressClient := kube.ExtensionsV1beta1().Ingresses(namespace)
		backend := &v1beta1.IngressBackend{
			ServiceName: "testrunner",
			ServicePort: intstr.IntOrString{
				IntVal: testRunnerPort,
			},
		}
		kubeIng, err := kubeIngressClient.Create(&v1beta1.Ingress{
			ObjectMeta: metav1.ObjectMeta{
				Name:        "simple-ingress-route",
				Namespace:   namespace,
				Annotations: map[string]string{"kubernetes.io/ingress.class": "gloo"},
			},
			Spec: v1beta1.IngressSpec{
				Backend: backend,
				//TLS: []v1beta1.IngressTLS{
				//	{
				//		Hosts:      []string{"some.host"},
				//		SecretName: "doesntexistanyway",
				//	},
				//},
				Rules: []v1beta1.IngressRule{
					{
						//Host: "some.host",
						IngressRuleValue: v1beta1.IngressRuleValue{
							HTTP: &v1beta1.HTTPIngressRuleValue{
								Paths: []v1beta1.HTTPIngressPath{
									{
										Backend: *backend,
									},
								},
							},
						},
					},
				},
			},
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(kubeIng).NotTo(BeNil())

		ingressProxy := "ingress-proxy"
		ingressPort := 80
		setup.CurlEventuallyShouldRespond(setup.CurlOpts{
			Protocol: "http",
			Path:     "/",
			Method:   "GET",
			Host:     ingressProxy,
			Service:  ingressProxy,
			Port:     ingressPort,
		}, helpers.SimpleHttpResponse)
	})
})

func deployKnative() {
	b, err := ioutil.ReadFile(KnativeManifest())
	Expect(err).NotTo(HaveOccurred())

	err = helpers.RunCommandInput(string(b), true, "kubectl", "apply", "-f", "-")
	Expect(err).NotTo(HaveOccurred())
}

func deleteKnative() {
	b, err := ioutil.ReadFile(KnativeManifest())
	Expect(err).NotTo(HaveOccurred())

	err = helpers.RunCommandInput(string(b), true, "kubectl", "apply", "-f", "-")
	Expect(err).NotTo(HaveOccurred())
}

func KnativeManifest() string {
	return filepath.Join(helpers.GlooDir(), "test", "kube2e", "artifacts", "knative-no-istio.yaml")
}