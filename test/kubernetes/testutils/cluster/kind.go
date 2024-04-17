package cluster

import (
	"fmt"
	"io"
	"testing"

	"github.com/onsi/gomega"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/solo-io/gloo/pkg/utils/kubeutils"
	"github.com/solo-io/gloo/pkg/utils/kubeutils/kubectl"
	kubetestclients "github.com/solo-io/gloo/test/kubernetes/testutils/clients"
)

// MustKindContext returns the Context for a KinD cluster with the given name
func MustKindContext(testing testing.TB, testingWriter io.Writer, clusterName string) *Context {
	testing.Helper()

	kubeCtx := fmt.Sprintf("kind-%s", clusterName)

	restCfg, err := kubeutils.GetRestConfigWithKubeContext(kubeCtx)
	gomega.Expect(err).NotTo(gomega.HaveOccurred())

	clientset, err := kubernetes.NewForConfig(restCfg)
	gomega.Expect(err).NotTo(gomega.HaveOccurred())

	clt, err := client.New(restCfg, client.Options{
		Scheme: kubetestclients.MustClientScheme(testing),
	})
	gomega.Expect(err).NotTo(gomega.HaveOccurred())

	return &Context{
		Name:        clusterName,
		KubeContext: kubeCtx,
		RestConfig:  restCfg,
		Cli:         kubectl.NewCli().WithKubeContext(kubeCtx).WithReceiver(testingWriter),
		Client:      clt,
		Clientset:   clientset,
	}
}
