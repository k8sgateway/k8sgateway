package shadowing

import (
	envoycore "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	envoyroute "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/shadowing"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
)

var _ = Describe("Plugin", func() {

	It("should work on valid inputs, with uninitialized outputs", func() {
		p := NewPlugin()

		upRef := &core.ResourceRef{
			Name:      "some-upstream",
			Namespace: "default",
		}
		in := &v1.Route{
			RoutePlugins: &v1.RoutePlugins{
				Shadowing: &shadowing.RouteShadowing{
					Upstream:   upRef,
					Percentage: 100,
				},
			},
		}
		out := &envoyroute.Route{}
		err := p.ProcessRoute(plugins.RouteParams{}, in, out)
		Expect(err).NotTo(HaveOccurred())
		checkFraction(out.GetRoute().RequestMirrorPolicy.RuntimeFraction, 100)
		Expect(out.GetRoute().RequestMirrorPolicy.Cluster).To(Equal("some-upstream_default"))
	})

	It("should work on valid inputs, with initialized outputs", func() {
		p := NewPlugin()

		upRef := &core.ResourceRef{
			Name:      "some-upstream",
			Namespace: "default",
		}
		in := &v1.Route{
			RoutePlugins: &v1.RoutePlugins{
				Shadowing: &shadowing.RouteShadowing{
					Upstream:   upRef,
					Percentage: 100,
				},
			},
		}
		var out = &envoyroute.Route{
			Action: &envoyroute.Route_Route{
				Route: &envoyroute.RouteAction{
					PrefixRewrite: "/something/set/by/another/plugin",
				},
			},
		}
		err := p.ProcessRoute(plugins.RouteParams{}, in, out)
		Expect(err).NotTo(HaveOccurred())
		checkFraction(out.GetRoute().RequestMirrorPolicy.RuntimeFraction, 100)
		Expect(out.GetRoute().RequestMirrorPolicy.Cluster).To(Equal("some-upstream_default"))
		Expect(out.GetRoute().PrefixRewrite).To(Equal("/something/set/by/another/plugin"))
	})

	It("should not error on empty configs", func() {
		p := NewPlugin()
		in := &v1.Route{}
		out := &envoyroute.Route{}
		err := p.ProcessRoute(plugins.RouteParams{}, in, out)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should error when set on invalid routes", func() {
		p := NewPlugin()

		upRef := &core.ResourceRef{
			Name:      "some-upstream",
			Namespace: "default",
		}
		in := &v1.Route{
			RoutePlugins: &v1.RoutePlugins{
				Shadowing: &shadowing.RouteShadowing{
					Upstream:   upRef,
					Percentage: 100,
				},
			},
		}
		// a redirect route is not a valid target for this plugin
		out := &envoyroute.Route{
			Action: &envoyroute.Route_Redirect{
				Redirect: &envoyroute.RedirectAction{},
			},
		}
		err := p.ProcessRoute(plugins.RouteParams{}, in, out)
		Expect(err).To(HaveOccurred())
		Expect(err).To(Equal(InvalidRouteActionError))

		// a direct response route is not a valid target for this plugin
		out = &envoyroute.Route{
			Action: &envoyroute.Route_DirectResponse{
				DirectResponse: &envoyroute.DirectResponseAction{},
			},
		}
		err = p.ProcessRoute(plugins.RouteParams{}, in, out)
		Expect(err).To(HaveOccurred())
		Expect(err).To(Equal(InvalidRouteActionError))
	})

	It("should error when given invalid specs", func() {
		p := NewPlugin()

		upRef := &core.ResourceRef{
			Name:      "some-upstream",
			Namespace: "default",
		}
		in := &v1.Route{
			RoutePlugins: &v1.RoutePlugins{
				Shadowing: &shadowing.RouteShadowing{
					Upstream:   upRef,
					Percentage: 200,
				},
			},
		}
		out := &envoyroute.Route{}
		err := p.ProcessRoute(plugins.RouteParams{}, in, out)
		Expect(err).To(HaveOccurred())
		Expect(err).To(Equal(InvalidNumeratorError(200)))

		in = &v1.Route{
			RoutePlugins: &v1.RoutePlugins{
				Shadowing: &shadowing.RouteShadowing{
					Percentage: 100,
				},
			},
		}
		out = &envoyroute.Route{}
		err = p.ProcessRoute(plugins.RouteParams{}, in, out)
		Expect(err).To(HaveOccurred())
		Expect(err).To(Equal(UnspecifiedUpstreamError))
	})

})

func checkFraction(frac *envoycore.RuntimeFractionalPercent, percentage float32) {
	Expect(frac.DefaultValue.Numerator).To(Equal(uint32(percentage * 10000)))
}
