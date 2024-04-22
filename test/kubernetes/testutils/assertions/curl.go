package assertions

import (
	"context"
	"time"

	"github.com/solo-io/gloo/pkg/utils/requestutils/curl"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/test/gomega/matchers"
	"github.com/solo-io/gloo/test/gomega/transforms"
	"github.com/solo-io/gloo/test/kube2e/helper"
)

var (
	// curlPodObjectMeta contains the ObjectMeta for the Pod that will be used to execute curl requests
	// The `curl` pod (referenced here)
	curlPodObjectMeta = metav1.ObjectMeta{
		Name:      "curl",
		Namespace: "curl",
	}
)

func (p *Provider) CurlEventuallyResponds(curlOptions []curl.Option, expectedResponse *matchers.HttpResponse, timeout ...time.Duration) ClusterAssertion {
	return func(ctx context.Context) {
		ginkgo.GinkgoHelper()
		currentTimeout, pollingInterval := helper.GetTimeouts(timeout...)

		// for some useful-ish output
		tick := time.Tick(currentTimeout / 8)

		Eventually(func(g Gomega) {
			res := p.clusterContext.Cli.CurlFromEphemeralPod(ctx, curlPodObjectMeta, curlOptions...)
			select {
			default:
				break
			case <-tick:
				ginkgo.GinkgoWriter.Printf("want %v\nhave: %s", expectedResponse, res)
			}

			expectedResponseMatcher := WithTransform(transforms.WithCurlHttpResponse, matchers.HaveHttpResponse(expectedResponse))
			g.Expect(res).To(expectedResponseMatcher)
			ginkgo.GinkgoWriter.Printf("success: %v", res)
		}).
			WithTimeout(currentTimeout).
			WithPolling(pollingInterval).
			WithContext(ctx).
			Should(Succeed())
	}
}
