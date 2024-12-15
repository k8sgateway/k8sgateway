package krtcollections

import (
	"context"
	"time"

	"istio.io/istio/pkg/kube/krt"
	"k8s.io/apimachinery/pkg/runtime/schema"

	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	envoytype "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	extensionplug "github.com/solo-io/gloo/projects/gateway2/extensions2/plugin"
	extensionsplug "github.com/solo-io/gloo/projects/gateway2/extensions2/plugin"
	"github.com/solo-io/gloo/projects/gateway2/ir"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
)

var (
	VirtualBuiltInGK = schema.GroupKind{
		Group: "builtin",
		Kind:  "builtin",
	}
)

type builtinPlugin struct {
	spec     gwv1.HTTPRouteFilter
	mutation func(outputRoute *envoy_config_route_v3.Route) error
}

func (d *builtinPlugin) CreationTime() time.Time {
	return time.Time{}
}

func (d *builtinPlugin) Equals(in any) bool {
	// we don't really need equality check here, because this policy is embedded in the httproute,
	// and we have generation based equality checks for that already.
	return true
	// d2, ok := in.(*builtinPlugin)
	//
	//	if !ok {
	//		return false
	//	}
	//
	// // TODO: implement equality check
	// return d.spec == d2.spec
}

type builtinPluginGwPass struct {
}

func NewBuiltInIr(kctx krt.HandlerContext, f gwv1.HTTPRouteFilter, fromgk schema.GroupKind, fromns string, refgrants *RefGrantIndex, ups *UpstreamIndex) ir.PolicyIR {
	return &builtinPlugin{
		spec:     f,
		mutation: convert(kctx, f, fromgk, fromns, refgrants, ups),
	}
}

func NewBuiltinPlugin(ctx context.Context) extensionplug.Plugin {

	return extensionplug.Plugin{
		ContributesPolicies: map[schema.GroupKind]extensionsplug.PolicyPlugin{
			VirtualBuiltInGK: {
				//AttachmentPoints: []ir.AttachmentPoints{ir.HttpAttachmentPoint},
				NewGatewayTranslationPass: NewGatewayTranslationPass,
			},
		},
	}
}

func convert(kctx krt.HandlerContext, f gwv1.HTTPRouteFilter, fromgk schema.GroupKind, fromns string, refgrants *RefGrantIndex, ups *UpstreamIndex) func(outputRoute *envoy_config_route_v3.Route) error {
	switch f.Type {
	case gwv1.HTTPRouteFilterRequestMirror:
		return convertMirror(kctx, f.RequestMirror, fromgk, fromns, refgrants, ups)
	case gwv1.HTTPRouteFilterRequestHeaderModifier:
		return convertHeaderModifier(kctx, f.RequestHeaderModifier)
	}
	return nil
}

func convertHeaderModifier(kctx krt.HandlerContext, f *gwv1.HTTPHeaderFilter) func(outputRoute *envoy_config_route_v3.Route) error {
	if f == nil {
		return nil
	}
	var headersToAddd []*envoy_config_core_v3.HeaderValueOption
	// TODO: add validation for header names/values with CheckForbiddenCustomHeaders
	for _, h := range f.Add {
		headersToAddd = append(headersToAddd, &envoy_config_core_v3.HeaderValueOption{
			Header: &envoy_config_core_v3.HeaderValue{
				Key:   string(h.Name),
				Value: h.Value,
			},
			AppendAction: envoy_config_core_v3.HeaderValueOption_APPEND_IF_EXISTS_OR_ADD,
		})
	}
	for _, h := range f.Set {
		headersToAddd = append(headersToAddd, &envoy_config_core_v3.HeaderValueOption{
			Header: &envoy_config_core_v3.HeaderValue{
				Key:   string(h.Name),
				Value: h.Value,
			},
			AppendAction: envoy_config_core_v3.HeaderValueOption_APPEND_IF_EXISTS_OR_ADD,
		})
	}
	toremove := f.Remove

	return func(outputRoute *envoy_config_route_v3.Route) error {
		outputRoute.RequestHeadersToAdd = append(outputRoute.RequestHeadersToAdd, headersToAddd...)
		outputRoute.RequestHeadersToRemove = append(outputRoute.RequestHeadersToRemove, toremove...)
		return nil
	}
}

func convertMirror(kctx krt.HandlerContext, f *gwv1.HTTPRequestMirrorFilter, fromgk schema.GroupKind, fromns string, refgrants *RefGrantIndex, ups *UpstreamIndex) func(outputRoute *envoy_config_route_v3.Route) error {
	if f == nil {
		return nil
	}
	to := toFromBackendRef(fromns, f.BackendRef)
	if !refgrants.ReferenceAllowed(kctx, fromgk, fromns, to) {
		// TODO: report error
		return nil
	}
	up, err := ups.getUpstreamFromRef(kctx, fromns, f.BackendRef)
	if err != nil {
		// TODO: report error
		return nil
	}
	fraction := getFactionPercent(*f)
	return func(outputRoute *envoy_config_route_v3.Route) error {

		route := outputRoute.GetRoute()
		if route == nil {
			// TODO: report error
			return nil
		}

		route.RequestMirrorPolicies = append(route.RequestMirrorPolicies, &envoy_config_route_v3.RouteAction_RequestMirrorPolicy{
			Cluster:         up.ClusterName(),
			RuntimeFraction: fraction,
		})

		return nil
	}
}

func getFactionPercent(f gwv1.HTTPRequestMirrorFilter) *envoy_config_core_v3.RuntimeFractionalPercent {
	if f.Percent != nil {
		return &envoy_config_core_v3.RuntimeFractionalPercent{
			DefaultValue: &envoytype.FractionalPercent{
				Numerator:   uint32(*f.Percent),
				Denominator: envoytype.FractionalPercent_HUNDRED,
			},
		}
	}
	if f.Fraction != nil {
		denom := 100.0
		if f.Fraction.Denominator != nil {
			denom = float64(*f.Fraction.Denominator)
		}
		ratio := float64(f.Fraction.Numerator) / denom
		return &envoy_config_core_v3.RuntimeFractionalPercent{
			DefaultValue: toEnvoyPercentage(ratio),
		}
	}

	return &envoy_config_core_v3.RuntimeFractionalPercent{
		DefaultValue: &envoytype.FractionalPercent{
			Numerator:   uint32(100),
			Denominator: envoytype.FractionalPercent_HUNDRED,
		},
	}
}

func toEnvoyPercentage(percentage float64) *envoytype.FractionalPercent {
	return &envoytype.FractionalPercent{
		Numerator:   uint32(percentage * 10000),
		Denominator: envoytype.FractionalPercent_MILLION,
	}
}

func NewGatewayTranslationPass(ctx context.Context, tctx ir.GwTranslationCtx) ir.ProxyTranslationPass {
	return &builtinPluginGwPass{}
}
func (p *builtinPlugin) Name() string {
	return "builtin"
}

// called 1 time for each listener
func (p *builtinPluginGwPass) ApplyListenerPlugin(ctx context.Context, pCtx *ir.ListenerContext, out *envoy_config_listener_v3.Listener) {
}

func (p *builtinPluginGwPass) ApplyVhostPlugin(ctx context.Context, pCtx *ir.VirtualHostContext, out *envoy_config_route_v3.VirtualHost) {
}

// called 0 or more times
func (p *builtinPluginGwPass) ApplyForRoute(ctx context.Context, pCtx *ir.RouteContext, outputRoute *envoy_config_route_v3.Route) error {

	policy, ok := pCtx.Policy.(*builtinPlugin)
	if !ok {
		return nil
	}

	if policy.mutation == nil {
		// TODO: report error
		return nil
	}

	return policy.mutation(outputRoute)
}

func (p *builtinPluginGwPass) ApplyForRouteBackend(
	ctx context.Context,
	policy ir.PolicyIR,
	pCtx *ir.RouteBackendContext,
) error {
	return nil
}

func (p *builtinPluginGwPass) HttpFilters(ctx context.Context, fcc ir.FilterChainCommon) ([]plugins.StagedHttpFilter, error) {
	return nil, nil
}

func (p *builtinPluginGwPass) UpstreamHttpFilters(ctx context.Context) ([]plugins.StagedUpstreamHttpFilter, error) {
	return nil, nil
}

func (p *builtinPluginGwPass) NetworkFilters(ctx context.Context) ([]plugins.StagedNetworkFilter, error) {
	return nil, nil
}

// called 1 time (per envoy proxy). replaces GeneratedResources
func (p *builtinPluginGwPass) ResourcesToAdd(ctx context.Context) ir.Resources {
	return ir.Resources{}
}
