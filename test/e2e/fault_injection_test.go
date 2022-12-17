package e2e_test

import (
	"fmt"
	"net/http"
	"time"

	"github.com/solo-io/gloo/test/e2e"

	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	"github.com/solo-io/gloo/test/helpers"
	matchers2 "github.com/solo-io/gloo/test/matchers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	fault "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/faultinjection"
	"github.com/solo-io/solo-kit/pkg/utils/prototime"
)

var _ = Describe("Fault Injection", func() {

	var (
		testContext *e2e.TestContext
	)

	BeforeEach(func() {
		testContext = testContextFactory.NewTestContext()
		testContext.BeforeEach()
	})

	AfterEach(func() {
		testContext.AfterEach()
	})

	JustBeforeEach(func() {
		testContext.JustBeforeEach()
	})

	JustAfterEach(func() {
		testContext.JustAfterEach()
	})

	Context("Envoy Abort Fault", func() {

		BeforeEach(func() {
			vs := helpers.NewVirtualServiceBuilder().
				WithName("vs-test").
				WithNamespace(writeNamespace).
				WithDomain("test.com").
				WithRoutePrefixMatcher("test", "/").
				WithRouteOptions("test", &gloov1.RouteOptions{
					Faults: &fault.RouteFaults{
						Abort: &fault.RouteAbort{
							HttpStatus: uint32(503),
							Percentage: float32(100),
						},
					},
				}).
				WithRouteActionToUpstream("test", testContext.TestUpstream().Upstream).
				Build()

			testContext.ResourcesToCreate().VirtualServices = v1.VirtualServiceList{
				vs,
			}
		})

		It("works", func() {
			req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/", "localhost", defaults.HttpPort), nil)
			Expect(err).NotTo(HaveOccurred())
			req.Host = "test.com" // to match the vs-test

			Eventually(func(g Gomega) (*http.Response, error) {
				return http.DefaultClient.Do(req)
			}, "5s", ".5s").Should(matchers2.MatchHttpResponse(&http.Response{
				StatusCode: http.StatusServiceUnavailable,
			}))

		})
	})

	Context("Envoy Delay Fault", func() {

		BeforeEach(func() {
			vs := helpers.NewVirtualServiceBuilder().
				WithName("vs-test").
				WithNamespace(writeNamespace).
				WithDomain("test.com").
				WithRoutePrefixMatcher("test", "/").
				WithRouteOptions("test", &gloov1.RouteOptions{
					Faults: &fault.RouteFaults{
						Delay: &fault.RouteDelay{
							FixedDelay: prototime.DurationToProto(time.Second * 3),
							Percentage: float32(100),
						},
					},
				}).
				WithRouteActionToUpstream("test", testContext.TestUpstream().Upstream).
				Build()

			testContext.ResourcesToCreate().VirtualServices = v1.VirtualServiceList{
				vs,
			}
		})

		It("works", func() {
			req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/", "localhost", defaults.HttpPort), nil)
			Expect(err).NotTo(HaveOccurred())
			req.Host = "test.com" // to match the vs-test

			Eventually(func(g Gomega) *http.Response {
				start := time.Now()
				response, err := http.DefaultClient.Do(req)
				g.Expect(err).NotTo(HaveOccurred())

				elapsed := time.Now().Sub(start)
				// This test regularly flakes, and the error is usually of the form:
				// "Elapsed time 2.998280684s not longer than delay 3s"
				// There's a small precision issue when communicating with Envoy, so including a small
				// margin of error to eliminate the test flake.
				marginOfError := 100 * time.Millisecond
				g.Expect(elapsed + marginOfError).To(BeNumerically(">", 3*time.Second))

				return response
			}, "20s", ".1s").Should(matchers2.MatchHttpResponse(&http.Response{
				StatusCode: http.StatusOK,
			}))

		})
	})
})
