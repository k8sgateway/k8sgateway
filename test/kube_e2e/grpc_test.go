package kube_e2e

import (
	"time"

	. "github.com/onsi/ginkgo"
	"github.com/solo-io/gloo/pkg/api/types/v1"
	. "github.com/solo-io/gloo/test/helpers"
)

var _ = Describe("GRPC Function Discovery", func() {
	upstreamName := namespace + "-grpc-test-service-8080"
	Context("creating a vService with a route to a GRPC function generated by"+
		" gloo-function-discovery", func() {
		listPath := "/list"
		vServiceName := "grpc-test"
		BeforeEach(func() {
			_, err := gloo.V1().VirtualServices().Create(&v1.VirtualService{
				Name: vServiceName,
				Routes: []*v1.Route{
					{
						Matcher: &v1.Route_RequestMatcher{
							RequestMatcher: &v1.RequestMatcher{
								Path: &v1.RequestMatcher_PathExact{
									PathExact: listPath,
								},
							},
						},
						SingleDestination: &v1.Destination{
							DestinationType: &v1.Destination_Function{
								Function: &v1.FunctionDestination{
									FunctionName: "ListShelves",
									UpstreamName: upstreamName,
								},
							},
						},
					},
				},
			})
			Must(err)
		})
		AfterEach(func() {
			gloo.V1().VirtualServices().Delete(vServiceName)
		})
		It("should route to the grpc function", func() {
			curlEventuallyShouldRespond(curlOpts{
				path: listPath,
			}, "< HTTP/1.1 200", time.Minute*5)
			curlEventuallyShouldRespond(curlOpts{
				path: listPath,
			}, "{}", time.Minute*5)
		})
	})
})
