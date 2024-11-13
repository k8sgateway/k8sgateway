package convert

import (
	"errors"
	"fmt"
	"github.com/solo-io/gloo/pkg/schemes"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sigs.k8s.io/yaml"
	"strings"

	gloogwv1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	gatewaykube "github.com/solo-io/gloo/projects/gateway/pkg/api/v1/kube/apis/gateway.solo.io/v1"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/core/matchers"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/gcp"
	glookube "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/kube/apis/gloo.solo.io/v1"
	v3 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	v2 "github.com/solo-io/solo-apis/pkg/api/enterprise.gloo.solo.io/v1"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/wrapperspb"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/utils/ptr"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
)

const (
	RandomSuffix = 4
	RandomSeed   = 1
)

var runtimeScheme *runtime.Scheme
var codecs serializer.CodecFactory
var decoder runtime.Decoder

func RootCmd() *cobra.Command {
	opts := &Options{}
	cmd := &cobra.Command{
		Use:   "convert",
		Short: "Convert Gloo Edge APIs to Gateway API",
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(opts)
		},
	}
	opts.addToFlags(cmd.PersistentFlags())
	cmd.SilenceUsage = true
	return cmd
}

func run(opts *Options) error {

	foundFiles, err := findFiles(opts)
	if err != nil {
		return err
	}

	filesMetrics.Add(float64(len(foundFiles)))
	var inputs []*GlooEdgeInput

	for _, file := range foundFiles {
		input, err := translateFileToEdgeInput(file)
		if err != nil {
			return err
		}
		inputs = append(inputs, input)
	}

	// preprocessing
	for _, input := range inputs {
		NewPreprocessor().Preprocess(input)
	}

	// now we need to convert the easy stuff like route tables
	var outputs []*GatewayAPIOutput
	for _, input := range inputs {
		output, err := translateEdgeAPIToGatewayAPI(input, opts)
		if err != nil {
			return err
		}
		outputs = append(outputs, output)
	}

	// we need to parse through the known delegated routes and translate them
	for _, output := range outputs {
		doDelegation(output)
	}

	// write all the outputs to their files
	for _, output := range outputs {
		//only write or
		txt, err := output.ToString()
		if err != nil {
			return err
		}

		if opts.Overwrite {
			fileNameSplit := strings.Split(output.FileName, ".")
			// assuming anything before the first . is the file name
			filename := fmt.Sprintf("%s.%s", fileNameSplit[0], strings.Join(fileNameSplit[1:], "."))
			if opts.OverwriteSuffix != "" {
				fmt.Sprintf("%s-%s.%s", fileNameSplit[0], opts.OverwriteSuffix, strings.Join(fileNameSplit[1:], "."))
			}
			if output.HasItems() {
				_, _ = fmt.Fprintf(os.Stdout, "Writing File: %s\n", filename)
				file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
				if err != nil {
					log.Fatal(err)
				}
				defer file.Close()
				fmt.Fprintf(file, "%s", txt)
			} else {
				_, _ = fmt.Fprintf(os.Stdout, "Skipping File because no Edge APIs Detected: %s\n", output.FileName)
			}
		} else {
			_, _ = fmt.Fprintf(os.Stdout, "\n\n---\n# --------------------------------\n# %s\n# --------------------------------", output.FileName)
			_, _ = fmt.Fprintf(os.Stdout, "%s\n", txt)
		}
		if opts.Stats {
			totalLines.WithLabelValues("Gateway API").Add(float64(len(strings.Split(txt, "\n"))))
		}
	}
	if opts.Stats {
		//count total lines of generated yaml (probably expensive)
		for _, input := range inputs {
			txt, _ := input.ToString()
			totalLines.WithLabelValues("Gloo").Add(float64(len(strings.Split(txt, "\n"))))
		}
		printMetrics(outputs)
	}

	return nil
}

// find all the HTTPRoutes that match and add a parent ref to the correct parent
func doDelegation(output *GatewayAPIOutput) {
	for _, delegateRoute := range output.DelegationReferences {
		addParentRefsToHTTPRoutes(output, delegateRoute)
	}
}

func addParentRefsToHTTPRoutes(output *GatewayAPIOutput, delegateRoute *DelegateParentReference) {

	// for each delegate route we need to go find the RouteTable that matches and assign a parent
	for _, httpRoute := range output.HTTPRoutes {
		if len(httpRoute.Labels) > 0 && doHttpRouteLabelsMatch(delegateRoute.Labels, httpRoute.Labels) {

			parentRef := gwv1.ParentReference{
				Name:      gwv1.ObjectName(delegateRoute.ParentName),
				Namespace: (*gwv1.Namespace)(&delegateRoute.ParentNamespace),
				Kind:      (*gwv1.Kind)(ptr.To("HTTPRoute")),
				Group:     (*gwv1.Group)(ptr.To("gateway.networking.k8s.io")),
			}

			httpRoute.Spec.ParentRefs = append(httpRoute.Spec.ParentRefs, parentRef)
		}
	}
}

func doHttpRouteLabelsMatch(matches map[string]string, labels map[string]string) bool {
	for k, v := range matches {
		if labels[k] != v {
			return false
		}
	}
	return true
}

func translateEdgeAPIToGatewayAPI(input *GlooEdgeInput, opts *Options) (*GatewayAPIOutput, error) {

	output := &GatewayAPIOutput{
		FileName:           input.FileName,
		YamlObjects:        input.YamlObjects,
		VirtualHostOptions: input.VirtualHostOptions,
		RouteOptions:       input.RouteOptions,
		AuthConfigs:        input.AuthConfigs,
		// Gateways:           input.Gateways,
	}

	for _, upstream := range input.Upstreams {
		if opts.GCPRegex != "" {
			newUpstream, err := convertGCPUpstream(upstream, opts.GCPRegex)
			if err != nil {
				return nil, err
			}
			// only update it if we get something back
			if newUpstream != nil {
				upstream = newUpstream
			}
		}
		output.Upstreams = append(output.Upstreams, upstream)
	}

	for _, authConfig := range input.AuthConfigs {
		// In the past users had to place gcp auth in authConfig instead of Upstream, if true then we need to remove the AuthConfig
		if opts.RemoveGCPAUthConfig && authConfigContainsGCPAuth(authConfig) {
			newAuthConfig := removeGCPAuthFromAuthConfig(authConfig)

			// if gcp_auth was the only config then we dont need to render it
			if newAuthConfig != nil && len(newAuthConfig.Spec.Configs) > 0 {
				output.AuthConfigs = append(output.AuthConfigs, newAuthConfig)
			}
		}
	}

	for _, routeTable := range input.RouteTables {
		httpRoute, routeOptions, delegates, err := convertRouteTableToHTTPRoute(routeTable)
		if err != nil {
			return nil, err
		}
		output.HTTPRoutes = append(output.HTTPRoutes, httpRoute)
		if len(delegates) > 0 {
			output.DelegationReferences = append(output.DelegationReferences, delegates...)
		}
		output.RouteOptions = append(output.RouteOptions, routeOptions...)
	}

	for _, virtualService := range input.VirtualServices {
		httpRoute, routeOptions, virtualHostOptions, delegates, err := convertVSToHTTPRoute(virtualService)
		if err != nil {
			return nil, err
		}
		output.HTTPRoutes = append(output.HTTPRoutes, httpRoute)
		output.RouteOptions = append(output.RouteOptions, routeOptions...)
		output.VirtualHostOptions = append(output.VirtualHostOptions, virtualHostOptions...)
		if len(delegates) > 0 {
			output.DelegationReferences = append(output.DelegationReferences, delegates...)
		}
	}

	return output, nil
}

func convertVSToHTTPRoute(vs *gatewaykube.VirtualService) (
	*gwv1.HTTPRoute,
	[]*gatewaykube.RouteOption,
	[]*gatewaykube.VirtualHostOption,
	[]*DelegateParentReference,
	error,
) {
	var options []*gatewaykube.RouteOption
	var vhOptions []*gatewaykube.VirtualHostOption
	var delegates []*DelegateParentReference
	hr := &gwv1.HTTPRoute{
		TypeMeta:   vs.TypeMeta,
		ObjectMeta: vs.ObjectMeta,
		Spec: gwv1.HTTPRouteSpec{
			CommonRouteSpec: gwv1.CommonRouteSpec{
				ParentRefs: []gwv1.ParentReference{
					{
						Name:      gwv1.ObjectName("http"),
						Namespace: (*gwv1.Namespace)(ptr.To("gloo-system")),
						Kind:      (*gwv1.Kind)(ptr.To("Gateway")),
						Group:     (*gwv1.Group)(ptr.To("gateway.networking.k8s.io")),
					},
				},
			},
			Hostnames: convertDomains(vs.Spec.VirtualHost.Domains),
			Rules:     []gwv1.HTTPRouteRule{},
		},
	}
	// To avoid naming collisions we need to postpend VW HTTPRoutes with -vs
	hr.GetObjectMeta().SetName(hr.GetObjectMeta().GetName() + "-vs")

	for _, route := range vs.Spec.VirtualHost.Routes {
		rr, option, genDelegates, err := convertRouteToRule(route, vs.Name, vs.Namespace)
		if err != nil {
			return nil, nil, nil, nil, err
		}
		if genDelegates != nil {
			delegates = append(delegates, genDelegates)
		}

		hr.Spec.Rules = append(hr.Spec.Rules, rr)
		if option != nil {
			options = append(options, option)
		}
		if route.GetOptions() != nil {
			routeOptions := route.GetOptions()

			// prefix rewrite, sets it on HTTPRoute
			if routeOptions.GetPrefixRewrite() != nil {
				rr.Filters = append(rr.Filters, gwv1.HTTPRouteFilter{
					Type: gwv1.HTTPRouteFilterURLRewrite,
					URLRewrite: &gwv1.HTTPURLRewriteFilter{
						Path: &gwv1.HTTPPathModifier{
							Type:               gwv1.PrefixMatchHTTPPathModifier,
							ReplacePrefixMatch: ptr.To(routeOptions.GetPrefixRewrite().Value),
						},
					},
				})
			}
		}
		if vs.Spec.VirtualHost.GetOptions() != nil {
			vo := vs.Spec.VirtualHost.GetOptions()

			opt, filter := convertVirtualHostOptions(vo, vs.Namespace, vs.Name)
			if filter != nil {
				rr.Filters = append(rr.Filters, *filter)
			}
			vhOptions = append(vhOptions, opt)
		}
	}

	return hr, options, vhOptions, delegates, nil
}

func convertVirtualHostOptions(
	options *gloov1.VirtualHostOptions,
	namespace string,
	routeName string,
) (*gatewaykube.VirtualHostOption, *gwv1.HTTPRouteFilter) {
	var ro *gatewaykube.VirtualHostOption
	var filter *gwv1.HTTPRouteFilter
	associationID := RandStringRunes(8)
	if routeName == "" {
		routeName = "vh-association"
	}
	associationName := fmt.Sprintf("%s-%s", routeName, associationID)

	// converts options to RouteOptions but we need to this for everything except prefix rewrite
	ro = &gatewaykube.VirtualHostOption{
		TypeMeta: v1.TypeMeta{
			Kind:       "VirtualHostOption",
			APIVersion: "gateway.solo.io/v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      associationName,
			Namespace: namespace,
		},
		Spec: gloogwv1.VirtualHostOption{
			Options: options,
			// TODO we just reference a non existent gateway today
			TargetRefs: []*v3.PolicyTargetReferenceWithSectionName{
				{
					Group:     "gateway.networking.k8s.io",
					Kind:      "Gateway",
					Name:      "http",
					Namespace: wrapperspb.String("gloo-system"),
				},
			},
		},
	}

	filter = &gwv1.HTTPRouteFilter{
		Type: gwv1.HTTPRouteFilterExtensionRef,
		ExtensionRef: &gwv1.LocalObjectReference{
			Group: "gateway.solo.io",
			Kind:  "RouteOption",
			Name:  gwv1.ObjectName(associationName),
		},
	}

	return ro, filter
}

func convertRouteTableToHTTPRoute(rt *gatewaykube.RouteTable) (
	*gwv1.HTTPRoute,
	[]*gatewaykube.RouteOption,
	[]*DelegateParentReference,
	error,
) {

	var options []*gatewaykube.RouteOption
	var delegates []*DelegateParentReference

	hr := &gwv1.HTTPRoute{
		TypeMeta:   rt.TypeMeta,
		ObjectMeta: rt.ObjectMeta,
		Spec: gwv1.HTTPRouteSpec{
			// CommonRouteSpec: gwv1.CommonRouteSpec{},
			// Hostnames: [],
			Rules: []gwv1.HTTPRouteRule{},
		},
	}

	for _, route := range rt.Spec.Routes {
		rule, option, delegate, err := convertRouteToRule(route, rt.Name, rt.Namespace)
		if err != nil {
			return nil, nil, nil, err
		}
		hr.Spec.Rules = append(hr.Spec.Rules, rule)
		if option != nil {
			options = append(options, option)
		}
		if delegate != nil {
			delegates = append(delegates, delegate)
		}
	}

	return hr, options, delegates, nil
}

func convertRouteToRule(r *gloogwv1.Route, routeTableName string, routeTableNamespace string) (
	gwv1.HTTPRouteRule,
	*gatewaykube.RouteOption,
	*DelegateParentReference,
	error,
) {
	var ro *gatewaykube.RouteOption

	rr := gwv1.HTTPRouteRule{
		Matches:     []gwv1.HTTPRouteMatch{},
		Filters:     []gwv1.HTTPRouteFilter{},
		BackendRefs: []gwv1.HTTPBackendRef{},
	}

	for _, m := range r.Matchers {
		match, err := convertMatch(m, routeTableName)
		if err != nil {
			return rr, nil, nil, err
		}
		rr.Matches = append(rr.Matches, match)
	}
	if r.GetRedirectAction() != nil {

		rf, err := generateFilterForRedirectAction(r)
		if err != nil {
			return rr, nil, nil, err
		}
		rr.Filters = append(rr.Filters, rf)
	}
	if r.GetOptions() != nil {
		options := r.GetOptions()

		// prefix rewrite, sets it on HTTPRoute
		if options.GetPrefixRewrite() != nil {
			rf, err := generateFilterForURLRewrite(r)
			if err != nil {
				return rr, nil, nil, err
			}
			rr.Filters = append(rr.Filters, rf)
		}

		var filter *gwv1.HTTPRouteFilter
		ro, filter = convertRouteOptions(options, r.Name, routeTableNamespace)
		if filter != nil {
			rr.Filters = append(rr.Filters, *filter)
		}
	}
	var delegate *DelegateParentReference

	if r.GetRouteAction() != nil && r.GetRouteAction().GetSingle() != nil {
		// single static upstream
		if r.GetRouteAction().GetSingle().GetUpstream() != nil {
			backendRef := generateBackendRefForSingleUpstream(r)

			rr.BackendRefs = append(rr.BackendRefs, backendRef)
		}
	} else if r.GetDelegateAction() != nil {
		// intermediate delegation step. This is a placeholder for the next path to do delegation
		backendRef, genDelegates := generateBackendRefForDelegateAction(r, routeTableName, routeTableNamespace)

		if len(backendRef) > 0 {
			for _, b := range backendRef {
				rr.BackendRefs = append(rr.BackendRefs, *b)
			}
		}
		delegate = genDelegates
	}

	return rr, ro, delegate, nil
}

func convertMatch(m *matchers.Matcher, routeTableName string) (gwv1.HTTPRouteMatch, error) {
	hrm := gwv1.HTTPRouteMatch{
		QueryParams: []gwv1.HTTPQueryParamMatch{},
	}

	// header matching
	if len(m.Headers) > 0 {
		hrm.Headers = []gwv1.HTTPHeaderMatch{}
		for _, h := range m.Headers {
			// support invert header match https://github.com/solo-io/gloo/blob/main/projects/gateway2/translator/httproute/gateway_http_route_translator.go#L274
			if h.InvertMatch == true {
				return hrm, errors.New("invert match not currently supported")
			}
			if h.Regex {
				hrm.Headers = append(hrm.Headers, gwv1.HTTPHeaderMatch{
					Type:  ptr.To(gwv1.HeaderMatchRegularExpression),
					Value: h.Value,
					Name:  gwv1.HTTPHeaderName(h.Name),
				})
			} else {
				hrm.Headers = append(hrm.Headers, gwv1.HTTPHeaderMatch{
					Type:  ptr.To(gwv1.HeaderMatchExact),
					Value: h.Value,
					Name:  gwv1.HTTPHeaderName(h.Name),
				})
			}
		}

	}

	// method matching
	if len(m.Methods) > 0 {
		if len(m.Methods) > 1 {
			return hrm, errors.New(fmt.Sprintf("Gateway API only supports 1 method match per rule and %d were detected for RouteTable %s", len(m.Methods), routeTableName))
		}
		hrm.Method = (*gwv1.HTTPMethod)(ptr.To(m.Methods[0]))
	}

	// query param matching
	if len(m.QueryParameters) > 0 {
		for _, m := range m.QueryParameters {
			if m.Regex {
				hrm.QueryParams = append(hrm.QueryParams, gwv1.HTTPQueryParamMatch{
					Type:  ptr.To(gwv1.QueryParamMatchRegularExpression),
					Name:  (gwv1.HTTPHeaderName)(m.Name),
					Value: m.Value,
				})
			} else {
				hrm.QueryParams = append(hrm.QueryParams, gwv1.HTTPQueryParamMatch{
					Type:  ptr.To(gwv1.QueryParamMatchExact),
					Name:  (gwv1.HTTPHeaderName)(m.Name),
					Value: m.Value,
				})
			}
		}
	}

	// Path matching
	if m.GetPathSpecifier() != nil {
		if m.GetPrefix() != "" {
			hrm.Path = &gwv1.HTTPPathMatch{
				Type:  ptr.To(gwv1.PathMatchPathPrefix),
				Value: ptr.To(m.GetPrefix()),
			}
		}
		if m.GetExact() != "" {
			hrm.Path = &gwv1.HTTPPathMatch{
				Type:  ptr.To(gwv1.PathMatchExact),
				Value: ptr.To(m.GetExact()),
			}
		}
		if m.GetRegex() != "" {
			hrm.Path = &gwv1.HTTPPathMatch{
				Type:  ptr.To(gwv1.PathMatchRegularExpression),
				Value: ptr.To(m.GetRegex()),
			}
		}
	}
	return hrm, nil
}

// Converts a single upstream to a GatewayAPI backend ref
func generateBackendRefForSingleUpstream(r *gloogwv1.Route) gwv1.HTTPBackendRef {
	upstream := r.GetRouteAction().GetSingle().GetUpstream()

	// static upstream reference
	backendRef := gwv1.HTTPBackendRef{
		BackendRef: gwv1.BackendRef{
			BackendObjectReference: gwv1.BackendObjectReference{
				Name:      gwv1.ObjectName(upstream.GetName()),
				Namespace: (*gwv1.Namespace)(ptr.To(upstream.GetNamespace())),
				Kind:      (*gwv1.Kind)(ptr.To("Upstream")),
				Group:     (*gwv1.Group)(ptr.To("gloo.solo.io")),
			},
		},
	}

	// AWS lambda integration
	if r.GetRouteAction().GetSingle().GetDestinationSpec() != nil && r.GetRouteAction().GetSingle().GetDestinationSpec().GetAws() != nil {
		// we need to add a parameter for the lambda name reference
		backendRef.Filters = append(backendRef.Filters, gwv1.HTTPRouteFilter{
			Type: gwv1.HTTPRouteFilterExtensionRef,
			ExtensionRef: &gwv1.LocalObjectReference{
				Kind:  "Parameter",
				Group: "gloo.solo.io",
				Name:  (gwv1.ObjectName)(r.GetRouteAction().GetSingle().GetDestinationSpec().GetAws().LogicalName),
			},
		})
	}
	return backendRef
}

func generateFilterForURLRewrite(r *gloogwv1.Route) (gwv1.HTTPRouteFilter, error) {

	rf := gwv1.HTTPRouteFilter{
		Type: gwv1.HTTPRouteFilterURLRewrite,
		URLRewrite: &gwv1.HTTPURLRewriteFilter{
			Path: &gwv1.HTTPPathModifier{},
		},
	}
	match, err := isExactMatch(r.GetMatchers())
	if err != nil {
		return rf, errors.New(fmt.Sprintf("RouteTable %s has multiple matchers in same route", r.Name))
	}
	if match {
		rf.URLRewrite.Path.Type = gwv1.FullPathHTTPPathModifier
		rf.URLRewrite.Path.ReplaceFullPath = ptr.To(r.GetOptions().GetPrefixRewrite().Value)
		rf.URLRewrite.Path.ReplacePrefixMatch = nil
	}
	match, err = isPrefixMatch(r.GetMatchers())
	if err != nil {
		return rf, errors.New(fmt.Sprintf("RouteTable %s has multiple matchers in same route", r.Name))
	}

	if match {
		rf.URLRewrite.Path.Type = gwv1.PrefixMatchHTTPPathModifier
		rf.URLRewrite.Path.ReplacePrefixMatch = ptr.To(r.GetOptions().GetPrefixRewrite().Value)
		rf.URLRewrite.Path.ReplaceFullPath = nil
	}

	// regex rewrite, NOT SUPPORTED IN GATEWAY API
	if r.GetOptions().GetRegexRewrite() != nil {
		return rf, errors.New(fmt.Sprintf("regex rewrite not supported, need to convert to another match"))
	}
	// rr.Filters = append(rr.Filters, gwv1.HTTPRouteFilter{
	// 	Type: gwv1.HTTPRouteFilterURLRewrite,
	// 	URLRewrite: &gwv1.HTTPURLRewriteFilter{
	// 		Path: &gwv1.HTTPPathModifier{
	// 			Type:               gwv1.HTTPPathModifierType(gwv1.PathMatchRegularExpression),
	// 			ReplacePrefixMatch: ptr.To(options.GetRegexRewrite().get),
	// 		},
	// 	},
	// })
	// }

	return rf, nil
}

func generateBackendRefForDelegateAction(
	r *gloogwv1.Route,
	routeTableName string,
	routeTableNamespace string,
) ([]*gwv1.HTTPBackendRef, *DelegateParentReference) {
	var backends []*gwv1.HTTPBackendRef
	if r.GetDelegateAction().GetRef() != nil {
		delegate := r.GetDelegateAction().GetRef()
		backendRef := &gwv1.HTTPBackendRef{
			BackendRef: gwv1.BackendRef{
				BackendObjectReference: gwv1.BackendObjectReference{
					Name:      gwv1.ObjectName(delegate.GetName()),
					Namespace: (*gwv1.Namespace)(ptr.To(delegate.GetNamespace())),
					Kind:      (*gwv1.Kind)(ptr.To("HTTPRoute")),
					Group:     (*gwv1.Group)(ptr.To("gateway.networking.k8s.io")),
				},
			},
		}
		backends = append(backends, backendRef)
		return backends, nil
	} else if r.GetDelegateAction().GetSelector() != nil {

		selector := r.GetDelegateAction().GetSelector()

		delegateParentRef := &DelegateParentReference{
			Labels:          selector.Labels,
			ParentName:      routeTableName,
			ParentNamespace: routeTableNamespace,
		}
		if len(r.GetDelegateAction().GetSelector().Namespaces) > 0 {
			for _, namespace := range r.GetDelegateAction().GetSelector().Namespaces {
				backendRef := &gwv1.HTTPBackendRef{
					BackendRef: gwv1.BackendRef{
						BackendObjectReference: gwv1.BackendObjectReference{
							Name:      "*",
							Namespace: (*gwv1.Namespace)(&namespace),
							Kind:      (*gwv1.Kind)(ptr.To("HTTPRoute")),
							Group:     (*gwv1.Group)(ptr.To("gateway.networking.k8s.io")),
						},
					},
				}
				backends = append(backends, backendRef)
			}
		} else {
			// default is gloo system for namespace if none are selected
			backendRef := &gwv1.HTTPBackendRef{
				BackendRef: gwv1.BackendRef{
					BackendObjectReference: gwv1.BackendObjectReference{
						Name:      "*",
						Namespace: (*gwv1.Namespace)(ptr.To("gloo-system")),
						Kind:      (*gwv1.Kind)(ptr.To("HTTPRoute")),
						Group:     (*gwv1.Group)(ptr.To("gateway.networking.k8s.io")),
					},
				},
			}
			backends = append(backends, backendRef)
		}

		return backends, delegateParentRef
	}
	return nil, nil
}

func generateFilterForRedirectAction(r *gloogwv1.Route) (gwv1.HTTPRouteFilter, error) {
	var statusCode int

	switch r.GetRedirectAction().ResponseCode {
	case gloov1.RedirectAction_MOVED_PERMANENTLY:
		statusCode = 301
	case gloov1.RedirectAction_FOUND:
		statusCode = 302
	case gloov1.RedirectAction_SEE_OTHER:
		statusCode = 303
	case gloov1.RedirectAction_TEMPORARY_REDIRECT:
		statusCode = 307
	case gloov1.RedirectAction_PERMANENT_REDIRECT:
		statusCode = 308
	default:
		statusCode = 301
	}

	rf := gwv1.HTTPRouteFilter{
		Type: gwv1.HTTPRouteFilterRequestRedirect,
		RequestRedirect: &gwv1.HTTPRequestRedirectFilter{
			StatusCode: ptr.To(statusCode),
			Path:       &gwv1.HTTPPathModifier{},
		},
	}
	if r.GetRedirectAction().HostRedirect != "" {
		rf.RequestRedirect.Hostname = ptr.To(gwv1.PreciseHostname(r.GetRedirectAction().HostRedirect))
	}
	if r.GetRedirectAction().PortRedirect != nil {
		rf.RequestRedirect.Port = ptr.To(gwv1.PortNumber(r.GetRedirectAction().PortRedirect.Value))
	}
	if r.GetRedirectAction().HttpsRedirect == true {
		rf.RequestRedirect.Scheme = ptr.To("https")
	}

	// we dont support stripQuery https://github.com/solo-io/gloo/blob/main/projects/gateway2/translator/plugins/redirect/redirect_plugin.go#L43
	if r.GetRedirectAction().StripQuery == true {
		return rf, errors.New("strip query not supported by Gateway API")
	}

	if r.GetRedirectAction().GetPathRedirect() != "" {

		match, err := isExactMatch(r.GetMatchers())
		if err != nil {
			return rf, errors.New(fmt.Sprintf("RouteTable %s has multiple matchers in same route", r.Name))
		}
		if match {
			rf.RequestRedirect.Path.Type = gwv1.FullPathHTTPPathModifier
			rf.RequestRedirect.Path.ReplaceFullPath = ptr.To(r.GetRedirectAction().GetPathRedirect())
			rf.RequestRedirect.Path.ReplacePrefixMatch = nil
		}
		match, err = isPrefixMatch(r.GetMatchers())
		if err != nil {
			return rf, errors.New(fmt.Sprintf("RouteTable %s has multiple matchers in same route", r.Name))
		}

		if match {
			rf.RequestRedirect.Path.Type = gwv1.PrefixMatchHTTPPathModifier
			rf.RequestRedirect.Path.ReplacePrefixMatch = ptr.To(r.GetRedirectAction().GetPathRedirect())
			rf.RequestRedirect.Path.ReplaceFullPath = nil
		}
	}
	return rf, nil
}

func convertRouteOptions(
	options *gloov1.RouteOptions,
	routeName string,
	routeTableNamespace string,
) (*gatewaykube.RouteOption, *gwv1.HTTPRouteFilter) {
	var ro *gatewaykube.RouteOption
	var filter *gwv1.HTTPRouteFilter
	associationID := RandStringRunes(RandomSuffix)
	if routeName == "" {
		routeName = "route-association"
	}
	associationName := fmt.Sprintf("%s-%s", routeName, associationID)

	// converts options to RouteOptions but we need to this for everything except prefixrewrite
	if isRouteOptionsSet(options) {
		ro = &gatewaykube.RouteOption{
			TypeMeta: v1.TypeMeta{
				Kind:       "RouteOption",
				APIVersion: "gateway.solo.io/v1",
			},
			ObjectMeta: v1.ObjectMeta{
				Name:      associationName,
				Namespace: routeTableNamespace,
			},
			Spec: gloogwv1.RouteOption{
				Options: options,
			},
		}

		// Because we move rewrites to a filter we need to remove it from RouteOptions
		if options.GetPrefixRewrite() != nil {
			ro.Spec.GetOptions().PrefixRewrite = nil
		}

		filter = &gwv1.HTTPRouteFilter{
			Type: gwv1.HTTPRouteFilterExtensionRef,
			ExtensionRef: &gwv1.LocalObjectReference{
				Group: "gateway.solo.io",
				Kind:  "RouteOption",
				Name:  gwv1.ObjectName(associationName),
			},
		}
	}
	return ro, filter
}

// TODO this is a mess
func isRouteOptionsSet(options *gloov1.RouteOptions) bool {
	return options.GetExtProc() != nil || options.GetCors() != nil || options.GetRetries() != nil || options.GetTimeout() != nil ||
		options.GetStagedTransformations() != nil || options.GetAutoHostRewrite() != nil ||
		options.GetFaults() != nil || options.GetExtensions() != nil || options.GetTracing() != nil || options.GetShadowing() != nil ||
		options.GetHeaderManipulation() != nil || options.GetAppendXForwardedHost() != nil || options.GetLbHash() != nil || options.GetUpgrades() != nil ||
		options.GetRatelimit() != nil || options.GetRatelimitBasic() != nil || options.GetWaf() != nil || options.GetJwtConfig() != nil || options.GetRbac() != nil ||
		options.GetDlp() != nil || options.GetStagedTransformations() != nil || options.GetEnvoyMetadata() != nil || options.GetMaxStreamDuration() != nil ||
		options.GetIdleTimeout() != nil || options.GetRegexRewrite() != nil || options.GetExtauth() != nil
}

func removeGCPAuthFromAuthConfig(config *v2.AuthConfig) *v2.AuthConfig {
	newAuth := &v2.AuthConfig{
		TypeMeta:   config.TypeMeta,
		ObjectMeta: config.ObjectMeta,
		Spec: v2.AuthConfigSpec{
			BooleanExpr:    config.Spec.BooleanExpr,
			FailOnRedirect: config.Spec.FailOnRedirect,
		},
	}

	for _, config := range config.Spec.Configs {
		const gcpAuthPluginName = "gcp_auth"
		if config.GetPluginAuth() == nil {
			newAuth.Spec.Configs = append(newAuth.Spec.Configs, config)
		} else if config.GetPluginAuth().Name != gcpAuthPluginName {
			newAuth.Spec.Configs = append(newAuth.Spec.Configs, config)
		}
	}
	return newAuth
}

func authConfigContainsGCPAuth(config *v2.AuthConfig) bool {
	if len(config.Spec.Configs) > 0 {
		for _, config := range config.Spec.Configs {
			if config.GetPluginAuth() != nil && config.GetPluginAuth().Name == "gcp_auth" {
				return true
			}
		}
	}
	return false
}

// this function looks for a matching regex in the upstream host to determine if its a gcp function or not
func convertGCPUpstream(upstream *glookube.Upstream, regex string) (*glookube.Upstream, error) {
	if upstream.Spec.GetStatic() != nil && len(upstream.Spec.GetStatic().GetHosts()) > 0 {
		if len(upstream.Spec.GetStatic().GetHosts()) > 1 {
			return nil, errors.New("unable to convert upstream to gcp, more than one host listed: " + upstream.Name)
		}
		for _, h := range upstream.Spec.GetStatic().Hosts {
			match, _ := regexp.MatchString(regex, h.Addr)
			if match {
				// log.Printf("Found regex %s: %s File: %s", upstreamRegex, h.Addr, filePath)
				// we need to replace the Upstream with a new gcp one
				newO := &glookube.Upstream{
					TypeMeta:   upstream.TypeMeta,
					ObjectMeta: upstream.ObjectMeta,
					Spec: gloov1.Upstream{
						UpstreamType: &gloov1.Upstream_Gcp{
							Gcp: &gcp.UpstreamSpec{
								Host: h.Addr,
							},
						},
					},
				}
				return newO, nil
			}
		}
	}
	return nil, nil
}

func translateFileToEdgeInput(fileName string) (*GlooEdgeInput, error) {

	gei := &GlooEdgeInput{
		FileName: fileName,
	}

	// Read the file
	data, err := os.ReadFile(fileName)
	if err != nil {
		return gei, err
	}
	for _, resourceYAML := range strings.Split(string(data), "---") {
		if len(resourceYAML) == 0 {
			continue
		}
		// yaml to object
		obj, k, err := decoder.Decode([]byte(resourceYAML), nil, nil)
		if err != nil {
			if runtime.IsNotRegisteredError(err) {
				// we just want to add the yaml and move on
				gei.YamlObjects = append(gei.YamlObjects, resourceYAML)
				continue
			}

			// TODO if we cant decode it, don't do anything and continue
			//log.Printf("# Skipping object due to error file parsing error %s", err)
			continue
		}

		// a lot of times lists are missing the group so this object doesnt match
		if k.Kind == "List" {
			var list unstructured.UnstructuredList
			if err := yaml.Unmarshal([]byte(resourceYAML), &list); err != nil {
				return nil, err
			}

			for _, item := range list.Items {

				tmpGei, err := parseObjects(item, k)
				if err != nil {
					return nil, err
				}
				gei.RouteOptions = append(gei.RouteOptions, tmpGei.RouteOptions...)
				gei.VirtualServices = append(gei.VirtualServices, tmpGei.VirtualServices...)
				gei.YamlObjects = append(gei.YamlObjects, tmpGei.YamlObjects...)
				gei.RouteTables = append(gei.RouteTables, tmpGei.RouteTables...)
				gei.VirtualHostOptions = append(gei.VirtualHostOptions, tmpGei.VirtualHostOptions...)
				gei.Upstreams = append(gei.Upstreams, tmpGei.Upstreams...)
				gei.AuthConfigs = append(gei.AuthConfigs, tmpGei.AuthConfigs...)
			}
			continue
		}

		switch o := obj.(type) {
		case *v2.AuthConfig:
			glooConfigMetric.WithLabelValues("AuthConfig").Inc()
			gei.AuthConfigs = append(gei.AuthConfigs, o)
		case *glookube.Upstream:
			glooConfigMetric.WithLabelValues("Upstream").Inc()
			gei.Upstreams = append(gei.Upstreams, o)
		case *glookube.UpstreamList:
			for _, upstream := range o.Items {
				glooConfigMetric.WithLabelValues("Upstream").Inc()
				gei.Upstreams = append(gei.Upstreams, &upstream)
			}
		case *gatewaykube.RouteTable:
			glooConfigMetric.WithLabelValues("RouteTable").Inc()
			gei.RouteTables = append(gei.RouteTables, o)
		case *gatewaykube.RouteTableList:
			for _, routeTable := range o.Items {
				glooConfigMetric.WithLabelValues("RouteTable").Inc()
				gei.RouteTables = append(gei.RouteTables, &routeTable)
			}
		case *gatewaykube.VirtualService:
			glooConfigMetric.WithLabelValues("VirtualService").Inc()
			gei.VirtualServices = append(gei.VirtualServices, o)
		case *gatewaykube.VirtualServiceList:
			for _, vs := range o.Items {
				glooConfigMetric.WithLabelValues("VirtualService").Inc()
				gei.VirtualServices = append(gei.VirtualServices, &vs)
			}
		case *gatewaykube.RouteOption:
			glooConfigMetric.WithLabelValues("RouteOption").Inc()
			gei.RouteOptions = append(gei.RouteOptions, o)
		case *gatewaykube.RouteOptionList:
			for _, ro := range o.Items {
				glooConfigMetric.WithLabelValues("RouteOption").Inc()
				gei.RouteOptions = append(gei.RouteOptions, &ro)
			}
		case *gatewaykube.VirtualHostOption:
			glooConfigMetric.WithLabelValues("VirtualHostOption").Inc()
			gei.VirtualHostOptions = append(gei.VirtualHostOptions, o)
		case *gatewaykube.Gateway:
			glooConfigMetric.WithLabelValues("Gateway").Inc()
			gei.YamlObjects = append(gei.YamlObjects, resourceYAML)
		default:
			// if we dont know what type it is we just add it back
			// no change so just add it back
			glooConfigMetric.WithLabelValues(k.Kind).Inc()
			gei.YamlObjects = append(gei.YamlObjects, resourceYAML)
		}
	}
	return gei, nil
}

func parseObjects(item unstructured.Unstructured, k *schema.GroupVersionKind) (*GlooEdgeInput, error) {
	gei := &GlooEdgeInput{}

	resourceYaml, err := yaml.Marshal(item)
	if err != nil {
		return nil, err
	}
	gvk := item.GroupVersionKind()
	obj, err := runtimeScheme.New(gvk)
	if runtime.IsNotRegisteredError(err) {
		// we just want to add the yaml and move on
		gei.YamlObjects = append(gei.YamlObjects, string(resourceYaml))
		return gei, nil
	} else if err != nil {
		return nil, err
	}

	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(item.Object, obj); err != nil {
		return nil, errors.New(fmt.Sprintf("Error converting unstructured to typed: %v", err))
	}

	switch o := obj.(type) {
	case *v2.AuthConfig:
		glooConfigMetric.WithLabelValues("AuthConfig").Inc()
		gei.AuthConfigs = append(gei.AuthConfigs, o)
	case *glookube.Upstream:
		glooConfigMetric.WithLabelValues("Upstream").Inc()
		gei.Upstreams = append(gei.Upstreams, o)
	case *gatewaykube.RouteTable:
		glooConfigMetric.WithLabelValues("RouteTable").Inc()
		gei.RouteTables = append(gei.RouteTables, o)
	case *gatewaykube.VirtualService:
		glooConfigMetric.WithLabelValues("VirtualService").Inc()
		gei.VirtualServices = append(gei.VirtualServices, o)
	case *gatewaykube.RouteOption:
		glooConfigMetric.WithLabelValues("RouteOption").Inc()
		gei.RouteOptions = append(gei.RouteOptions, o)
	case *gatewaykube.VirtualHostOption:
		glooConfigMetric.WithLabelValues("VirtualHostOption").Inc()
		gei.VirtualHostOptions = append(gei.VirtualHostOptions, o)
	case *gatewaykube.Gateway:
		glooConfigMetric.WithLabelValues("Gateway").Inc()
		gei.YamlObjects = append(gei.YamlObjects, string(resourceYaml))
	default:
		// if we dont know what type it is we just add it back
		// no change so just add it back
		glooConfigMetric.WithLabelValues(k.Kind).Inc()
		gei.YamlObjects = append(gei.YamlObjects, string(resourceYaml))
	}
	return gei, nil
}

func init() {
	runtimeScheme = runtime.NewScheme()

	//if err := metav1.AddToSche; err != nil {
	//	log.Fatal(err)
	//}
	//if err := glookube.AddToScheme(runtimeScheme); err != nil {
	//	log.Fatal(err)
	//}
	//if err := gatewaykube.AddToScheme(runtimeScheme); err != nil {
	//	log.Fatal(err)
	//}
	//if err := v2.AddToScheme(runtimeScheme); err != nil {
	//	log.Fatal(err)
	//}
	//if err := gwv1.Install(runtimeScheme); err != nil {
	//	log.Fatal(err)
	//}

	if err := schemes.SchemeBuilder.AddToScheme(runtimeScheme); err != nil {
		log.Fatal(err)
	}

	codecs = serializer.NewCodecFactory(runtimeScheme)
	decoder = codecs.UniversalDeserializer()
}

func findFiles(opts *Options) ([]string, error) {
	var files []string
	if opts.Directory != "" {
		fs, err := findYamlFiles(opts.Directory)
		if err != nil {
			return nil, err
		}
		files = fs
	} else {
		files = append(files, opts.InputFile)
	}

	return files, nil
}

func findYamlFiles(directory string) ([]string, error) {
	var files []string
	libRegEx, e := regexp.Compile("^.+\\.(yaml|yml)$")
	if e != nil {
		return nil, e
	}

	e = filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err == nil && libRegEx.MatchString(info.Name()) {
			// println(path)
			if !strings.Contains(info.Name(), "kustomization") {
				files = append(files, path)
			}
		}
		return nil
	})
	if e != nil {
		return nil, e
	}
	return files, nil
}

// This function validates that the RouteRable matchers are the same match type prefix or exact
// The reason being is that if you are doing a rewrite you can only have one type of filter applied
func validateMatchersAreSame(matches []*matchers.Matcher) error {

	var foundExact, foundPrefix, foundRegex bool
	for _, m := range matches {
		if m.GetExact() != "" {
			if foundPrefix || foundRegex {
				return errors.New("multiple matchers found")
			}
			foundExact = true
		}
		if m.GetPrefix() != "" {
			if foundExact || foundRegex {
				return errors.New("multiple matchers found")
			}
			foundPrefix = true
		}
		if m.GetRegex() != "" {
			if foundExact || foundPrefix {
				return errors.New("multiple matchers found")
			}
			foundRegex = true
		}
	}
	return nil
}

// tests to see if all matchers are exact
func isExactMatch(matches []*matchers.Matcher) (bool, error) {
	if err := validateMatchersAreSame(matches); err != nil {
		return false, err
	}
	for _, m := range matches {
		if m.GetExact() != "" {
			return true, nil
		}
	}
	return false, nil
}

// tests to see if all matchers are exact
func isPrefixMatch(matches []*matchers.Matcher) (bool, error) {
	if err := validateMatchersAreSame(matches); err != nil {
		return false, err
	}
	for _, m := range matches {
		if m.GetPrefix() != "" {
			return true, nil
		}
	}
	return false, nil
}

// tests to see if all matchers are regex
func isRegexMatch(matches []*matchers.Matcher) (bool, error) {
	if err := validateMatchersAreSame(matches); err != nil {
		return false, err
	}
	for _, m := range matches {
		if m.GetRegex() != "" {
			return true, nil
		}
	}
	return false, nil
}
