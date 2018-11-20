package kubernetes_test

import (
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/solo-projects/projects/gloo/pkg/discovery"
	kubeplugin "github.com/solo-io/solo-projects/projects/gloo/pkg/plugins/kubernetes"
	kubev1 "k8s.io/api/core/v1"
	"time"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
	"path/filepath"

	"github.com/solo-io/solo-kit/pkg/utils/log"
	"github.com/solo-io/solo-kit/test/helpers"
	"github.com/solo-io/solo-kit/test/setup"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var _ = Describe("Kubernetes", func() {
	var (
		namespace  string
		cfg        *rest.Config
		kubeClient kubernetes.Interface

		baseLabels = map[string]string{
			"tacos": "burritos",
		}

		extendedLabels = map[string]string{
			"tacos": "burritos",
			"pizza": "frenchfries",
		}
	)

	if os.Getenv("RUN_KUBE_TESTS") != "1" {
		log.Printf("This test creates kubernetes resources and is disabled by default. To enable, set RUN_KUBE_TESTS=1 in your env.")
		return
	}
	BeforeEach(func() {
		namespace = helpers.RandString(8)
		err := setup.SetupKubeForTest(namespace)
		Expect(err).NotTo(HaveOccurred())
		kubeconfigPath := filepath.Join(os.Getenv("HOME"), ".kube", "config")
		cfg, err = clientcmd.BuildConfigFromFlags("", kubeconfigPath)
		Expect(err).NotTo(HaveOccurred())
		kubeClient, err = kubernetes.NewForConfig(cfg)
		Expect(err).NotTo(HaveOccurred())
		// create a service
		// create 2 pods for that service
		// one with extra labels, one without
		svc, err := kubeClient.CoreV1().Services(namespace).Create(&kubev1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: namespace,
				Name:      "i-hate-writing-tests",
			},
			Spec: kubev1.ServiceSpec{
				Selector: baseLabels,
				Ports: []kubev1.ServicePort{
					{
						Name: "bar",
						Port: 8080,
					},
					{
						Name: "foo",
						Port: 9090,
					},
				},
			},
		})
		Expect(err).NotTo(HaveOccurred())
		_, err = kubeClient.CoreV1().Pods(namespace).Create(&kubev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "pod-for-" + svc.Name + "-basic",
				Namespace: namespace,
				Labels:    baseLabels,
			},
			Spec: kubev1.PodSpec{
				Containers: []kubev1.Container{
					{
						Name:  "nginx",
						Image: "nginx:latest",
					},
				},
			},
		})
		Expect(err).NotTo(HaveOccurred())
		_, err = kubeClient.CoreV1().Pods(namespace).Create(&kubev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "pod-for-" + svc.Name + "-extra",
				Namespace: namespace,
				Labels:    extendedLabels,
			},
			Spec: kubev1.PodSpec{
				Containers: []kubev1.Container{
					{
						Name:  "nginx",
						Image: "nginx:latest",
					},
				},
			},
		})
		Expect(err).NotTo(HaveOccurred())
	})
	AfterEach(func() {
		setup.TeardownKube(namespace)
	})

	It("uses json keys when serializing", func() {
		plug := kubeplugin.NewPlugin(kubeClient).(discovery.DiscoveryPlugin)
		upstreams, errs, err := plug.DiscoverUpstreams([]string{namespace}, namespace, clients.WatchOpts{
			Ctx:         context.TODO(),
			RefreshRate: time.Second,
		}, discovery.Opts{})
		Expect(err).NotTo(HaveOccurred())

		for {
			select {
			case <-time.After(time.Second * 2):
				Fail("no upstreams detected after 2s")
			case upstreamList := <-upstreams:
				Expect(upstreamList).To(HaveLen(2))
			case err, ok := <-errs:
				if !ok {
					return
				}
				Expect(err).NotTo(HaveOccurred())
			}
		}
	})
})
