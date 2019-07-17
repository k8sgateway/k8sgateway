package knative_test

import (
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/solo-io/go-utils/testutils/exec"
	"github.com/solo-io/go-utils/testutils/helper"

	"github.com/solo-io/go-utils/log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

var _ = Describe("Kube2e: Knative-Ingress", func() {

	BeforeEach(func() {
		deployKnativeTestService()
	})

	AfterEach(func() {
		if err := deleteKnativeTestService(); err != nil {
			log.Warnf("teardown failed %v", err)
		}
	})

	It("works", func() {
		clusterIngressProxy := "clusteringress-proxy"
		clusterIngressPort := 80
		testHelper.CurlEventuallyShouldRespond(helper.CurlOpts{
			Protocol:          "http",
			Path:              "/",
			Method:            "GET",
			Host:              "helloworld-go.default.example.com",
			Service:           clusterIngressProxy,
			Port:              clusterIngressPort,
			ConnectionTimeout: 1,
			Verbose:           true,
		}, "Hello Go Sample v1!", 1, time.Minute*2, 1*time.Second)
	})
})

func deployKnativeTestService() {
	b, err := ioutil.ReadFile(knativeTestServiceFile())
	ExpectWithOffset(1, err).NotTo(HaveOccurred())

	// The webhook may take a bit of time to initially be responsive
	// See: https://github.com/istio/istio/pull/7743/files
	EventuallyWithOffset(1, func() error {
		return exec.RunCommandInput(string(b), testHelper.RootDir, true, "kubectl", "apply", "-f", "-")
	}, "30s", "5s").Should(BeNil())
}

func deleteKnativeTestService() error {
	b, err := ioutil.ReadFile(knativeTestServiceFile())
	if err != nil {
		return err
	}
	err = exec.RunCommandInput(string(b), testHelper.RootDir, true, "kubectl", "delete", "-f", "-")
	if err != nil {
		return err
	}
	return nil
}

func knativeTestServiceFile() string {
	return filepath.Join(testHelper.RootDir, "test", "kube2e", "knative", "artifacts", "knative-hello-service.yaml")
}
