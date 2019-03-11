package knative_test

import (
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/solo-io/go-utils/testutils/exec"
	"github.com/solo-io/go-utils/testutils/helper"

	"github.com/solo-io/solo-kit/pkg/utils/log"

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
			ConnectionTimeout: 5,
		}, "Hello Go Sample v1!", 1, time.Minute*2)
	})
})

func deployKnativeTestService() {
	b, err := ioutil.ReadFile(knativeTestServiceFile())
	ExpectWithOffset(1, err).NotTo(HaveOccurred())

	err = exec.RunCommandInput(string(b), testHelper.RootDir, true, "kubectl", "apply", "-f", "-")
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
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
