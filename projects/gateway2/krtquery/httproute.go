package krtquery

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	k8sptr "k8s.io/utils/ptr"
	"sigs.k8s.io/controller-runtime/pkg/client"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
	v1 "sigs.k8s.io/gateway-api/apis/v1"

	"istio.io/istio/pkg/kube/krt"
	"istio.io/istio/pkg/ptr"
	"istio.io/istio/pkg/slices"

	"github.com/solo-io/gloo/projects/gateway2/query"
	"github.com/solo-io/gloo/projects/gateway2/translator/backendref"
	"github.com/solo-io/gloo/projects/gateway2/wellknown"
)

type (
	GatewayHTTPRoute = Attachment[*query.HTTPRouteInfo, *gwv1.Gateway]
	ServiceHTTPRoute = Attachment[*query.HTTPRouteInfo, *corev1.Service]
)

var _ krt.ResourceNamer = &GatewayHTTPRoute{}

type krtKey struct{}

var NonKRTError = fmt.Errorf("Called from outside a KRT Context")

// WithKRT allows wrapping the krt HandlerContext in a Go context.
func WithKRT(ctx context.Context, krtctx krt.HandlerContext) context.Context {
	return context.WithValue(ctx, krtKey{}, krtctx)
}

func GetKRT(ctx context.Context) (krt.HandlerContext, error) {
	krtctx, ok := ctx.Value(krtKey{}).(krt.HandlerContext)
	if !ok {
		return nil, NonKRTError
	}
	return krtctx, nil
}

func (q *Queries) HTTPRoutesForGateway(ctx krt.HandlerContext, gw *gwv1.Gateway) (query.RoutesForGwResult, error) {
	result := query.RoutesForGwResult{
		ListenerResults: map[string]*query.ListenerResult{},
	}

	// perform per-listener validation (only error is bad selectors for allowedRoutes)
	for _, l := range gw.Spec.Listeners {
		if l.AllowedRoutes.Namespaces.From != nil && *l.AllowedRoutes.Namespaces.From == gwv1.NamespacesFromSelector && l.AllowedRoutes.Namespaces.Selector == nil {
			lr := result.ListenerResults[string(l.Name)]
			if lr == nil {
				lr = &query.ListenerResult{}
				result.ListenerResults[string(l.Name)] = lr
			}
			lr.Error = fmt.Errorf("selector must be set")
		}
	}

	// aggregate successfully attached routes per-listener
	routes := q.GatewayHTTPRoutes.FetchAttachmentsTo(ctx, gw, "")
	for _, a := range routes {
		if a.Section == "" {
			continue
		}
		lr := result.ListenerResults[a.Section]
		if lr == nil {
			lr = &query.ListenerResult{}
			result.ListenerResults[string(a.Section)] = lr
		}
		lr.Routes = append(lr.Routes, a.Resource)
	}

	// aggregate errored routes
	errors := q.GatewayHTTPRoutes.FetchAttachmentsTo(ctx, gw, "")
	for _, a := range errors {
		result.RouteErrors = append(result.RouteErrors, &query.RouteError{
			Route:     a.Resource.HTTPRoute,
			ParentRef: a.Resource.ParentRef,
			Error:     toQueryError(a.Error),
		})
	}

	return result, nil
}

// this is ugly
func toQueryError(err error) query.Error {
	var reason gwv1.RouteConditionReason
	switch err {
	case query.ErrNotAllowedByListeners:
		reason = gwv1.RouteReasonNotAllowedByListeners
	case query.ErrNoMatchingParent:
		reason = gwv1.RouteReasonNoMatchingParent
	case query.ErrNoMatchingListenerHostname:
		reason = gwv1.RouteReasonNoMatchingListenerHostname
	}
	return query.Error{
		E:      err,
		Reason: reason,
	}
}

// routesForGateway are root routes that attach directly to the gateway.
func (q *Queries) routesForGateway(
	ctx context.Context,
) IndexedAttachments[*query.HTTPRouteInfo, *gwv1.Gateway] {
	// One HTTPRoute to many attached listeners.
	gwRoutes := krt.NewManyCollection(
		q.HTTPRoutes,
		func(ctx krt.HandlerContext, hr *gwv1.HTTPRoute) []GatewayHTTPRoute {
			var out []GatewayHTTPRoute
			for _, ref := range getParentRefsForGw(hr) {

				// first make sure the referenced gateway exists
				gw := ptr.Flatten(krt.FetchOne(ctx, q.Gateways, krt.FilterKey(parentRef(hr, ref).targetResourceKey())))
				if gw == nil {
					// when it's missing the errored attachment still needed for reports
					out = append(out, GatewayHTTPRoute{
						Resource: &query.HTTPRouteInfo{HTTPRoute: *hr, ParentRef: ref},
						Section:  string(ptr.OrDefault(ref.SectionName, "")),
						Error:    AttachmentErrTargetNotFound,
					})
					continue
				}

				// for HTTPRoute -> Gateway, unspecified Section means all listeners.
				var sections []gwv1.SectionName
				if ref.SectionName != nil {
					sections = []gwv1.SectionName{*ref.SectionName}
				} else {
					sections = slices.Map(gw.Spec.Listeners, func(l gwv1.Listener) gwv1.SectionName {
						return l.Name
					})
				}

				for _, section := range sections {
					info := &query.HTTPRouteInfo{HTTPRoute: *hr, ParentRef: ref}
					attachment := GatewayHTTPRoute{
						Resource: info,
						Target:   gw,
						Section:  string(section),
					}
					l := slices.FindFunc(gw.Spec.Listeners, func(l gwv1.Listener) bool {
						return l.Name == section
					})
					if l == nil {
						attachment.Error = AttachmentErrSectionNotFound
						out = append(out, attachment)
						continue
					}
					if !allowedRouteKind(l, metav1.GroupKind{Group: gwv1.GroupName, Kind: wellknown.HTTPRouteKind}) {
						continue
					}
					if !allowedRouteNamespace(gw, l, hr.GetNamespace()) {
						continue
					}

					// errors on referenced delgated routes are reported as part of HTTPRoute plugin
					// the errors we find here will live in the Backends map
					info.Backends = q.resolveRouteBackends(ctx, hr)
					info.Children = q.delegatedChildren(ctx, hr, nil)

					out = append(out, attachment)
				}
			}
			return out
		},
	)

	return CreateAttachmentIndexes(gwRoutes)
}

// HTTPRouteChain recursively resolves all backends of the given HTTPRoute.
// While this includes delegated HTTPRoutes, validation of matchers is not
// applied here. Instead, matcher processing is handled during translation.
// Errors for unresolved or cyclic backend references will be surfaced on the
// HTTPRouteInfo.
func (q *Queries) HTTPRouteChain(
	ctx krt.HandlerContext,
	route gwv1.HTTPRoute,
	hostnames []string,
	parentRef gwv1.ParentReference,
) *query.HTTPRouteInfo {
	return &query.HTTPRouteInfo{
		HTTPRoute:         route,
		HostnameOverrides: hostnames,
		ParentRef:         parentRef,
		Backends:          q.resolveRouteBackends(ctx, &route),
		Children:          q.delegatedChildren(ctx, &route, nil),
	}
}

// krt equiv of query.getDelegatedChildren
func (q *Queries) delegatedChildren(
	ctx krt.HandlerContext,
	parent *gwv1.HTTPRoute,
	visited sets.Set[types.NamespacedName],
) query.BackendMap[[]*query.HTTPRouteInfo] {
	if visited == nil {
		visited = sets.New[types.NamespacedName]()
	}
	parentRef := namespacedName(parent)
	visited.Insert(parentRef)

	children := query.NewBackendMap[[]*query.HTTPRouteInfo]()
	for _, parentRule := range parent.Spec.Rules {
		for _, backendRef := range parentRule.BackendRefs {
			if !backendref.RefIsHTTPRoute(backendRef.BackendObjectReference) {
				continue
			}
			referencedRoutes, err := q.fetchChildRoutes(
				ctx,
				parent.Namespace,
				backendRef,
			)
			if err != nil {
				children.AddError(backendRef.BackendObjectReference, err)
				continue
			}

			var backendRefChildren []*query.HTTPRouteInfo
			for _, childRoute := range referencedRoutes {
				childRef := namespacedName(childRoute)
				if visited.Has(childRef) {
					err := fmt.Errorf("ignoring cyclical child route %s for parent %s: %w", parentRef, childRef, query.ErrCyclicReference)
					children.AddError(backendRef.BackendObjectReference, err)
					// don't resolve child routes; the entire backendRef is invalid?
					break
				}
				routeInfo := &query.HTTPRouteInfo{
					HTTPRoute: *childRoute,
					ParentRef: gwv1.ParentReference{
						Group:     k8sptr.To(gwv1.Group(wellknown.GatewayGroup)),
						Kind:      k8sptr.To(gwv1.Kind(wellknown.HTTPRouteKind)),
						Namespace: k8sptr.To(gwv1.Namespace(parent.Namespace)),
						Name:      gwv1.ObjectName(parent.Name),
					},
					Backends: q.resolveRouteBackends(ctx, childRoute),
					Children: q.delegatedChildren(ctx, childRoute, visited),
				}
				backendRefChildren = append(backendRefChildren, routeInfo)
			}
			children.Add(backendRef.BackendObjectReference, backendRefChildren)
		}
	}
	return children
}

// TODO check reference grants
func (q *Queries) resolveRouteBackends(
	ctx krt.HandlerContext,
	hr *gwv1.HTTPRoute,
) query.BackendMap[client.Object] {
	out := query.NewBackendMap[client.Object]()
	for _, rule := range hr.Spec.Rules {
		for _, backendRef := range rule.BackendRefs {
			// TODO(stevenctl) we should probably report invalid kinds of backendRef.
			// We don't do it yet today, and we need some extensibilty here so I'm
			// not hanlding it in the move to krt; will do it with ServiceEntry work.
			resolvedBackend, err := q.GetGrantedRef(
				ctx,
				HTTPRouteGK,
				hr.GetNamespace(),
				backendref.GroupKind(backendRef.BackendObjectReference),
				string(backendRef.BackendObjectReference.Name),
				nsPtrToString(backendRef.BackendObjectReference.Namespace),
			)
			if err != nil {
				out.AddError(backendRef.BackendObjectReference, err)
				continue
			} else if resolvedBackend == nil {
				// should never happen if err is nil
				out.AddError(backendRef.BackendObjectReference, query.ErrUnresolvedReference)
				continue
			}

			out.Add(backendRef.BackendObjectReference, resolvedBackend)

		}
	}
	return out
}

func (q *Queries) fetchChildRoutes(
	ctx krt.HandlerContext,
	parentNamespace string,
	backendRef gwv1.HTTPBackendRef,
) ([]*gwv1.HTTPRoute, error) {
	delegatedNs := parentNamespace
	if backendRef.Namespace != nil {
		delegatedNs = string(*backendRef.Namespace)
	}

	// if the name is missing or '*' that means all routes in the namespace.
	filter := []krt.FetchOption{krt.FilterIndex(q.httpRoutesByNamespace, delegatedNs)}
	// if the name is specified, fetch only that one
	if backendRef.Name != "" && backendRef.Name != "*" {
		filter = append(filter, krt.FilterObjectName(types.NamespacedName{
			Name:      string(backendRef.Name),
			Namespace: delegatedNs,
		}))
	}

	var refChildren []*gwv1.HTTPRoute
	fetched := krt.Fetch(ctx, q.HTTPRoutes, filter...)
	refChildren = append(refChildren, fetched...)

	if len(refChildren) == 0 {
		return nil, query.ErrUnresolvedReference
	}
	return refChildren, nil
}

// getParentRefsForGw filters refs to only include those of kind Gateway
// If there is no Group/Kind, assume it's pointing to a Gateway.
func getParentRefsForGw(hr *gwv1.HTTPRoute) []gwv1.ParentReference {
	var ret []gwv1.ParentReference
	for _, pRef := range hr.Spec.ParentRefs {
		if pRef.Group != nil && *pRef.Group != gwv1.GroupName {
			continue
		}
		if pRef.Kind != nil && *pRef.Kind != wellknown.GatewayKind {
			continue
		}
		ret = append(ret, pRef)
	}
	return ret
}

func allowedRouteKind(l *gwv1.Listener, gk metav1.GroupKind) bool {
	// default to protocol based
	var allowedKinds []metav1.GroupKind
	switch l.Protocol {
	case gwv1.HTTPSProtocolType:
		fallthrough
	case gwv1.HTTPProtocolType:
		allowedKinds = []metav1.GroupKind{{Kind: wellknown.HTTPRouteKind, Group: gwv1.GroupName}}
		// TLS and TCP unsupported
	case gwv1.TLSProtocolType:
		fallthrough
	case gwv1.TCPProtocolType:
		allowedKinds = []metav1.GroupKind{{}}
	case gwv1.UDPProtocolType:
		allowedKinds = []metav1.GroupKind{{}}
	// custom protocols can have httproutes
	default:
		// allow custom protocols to work (i.e. istio.io/PROXY)
		allowedKinds = []metav1.GroupKind{{Kind: wellknown.HTTPRouteKind, Group: gwv1.GroupName}}
	}

	// listener specified
	if l != nil && l.AllowedRoutes != nil && l.AllowedRoutes.Kinds != nil {
		allowedKinds = slices.Map(l.AllowedRoutes.Kinds, func(k gwv1.RouteGroupKind) metav1.GroupKind {
			gk := metav1.GroupKind{Kind: string(k.Kind), Group: gwv1.GroupName}
			if k.Group != nil {
				gk.Group = string(*k.Group)
			}
			return gk
		})
	}

	return sets.New(allowedKinds...).Has(gk)
}

func allowedRouteNamespace(
	gw *gwv1.Gateway,
	l *gwv1.Listener,
	ns string,
) bool {
	var allowedNamespaces *v1.RouteNamespaces
	if l.AllowedRoutes != nil && l.AllowedRoutes.Namespaces != nil && l.AllowedRoutes.Namespaces.From != nil {
		allowedNamespaces = l.AllowedRoutes.Namespaces
	}

	switch *allowedNamespaces.From {
	case v1.NamespacesFromSelector:
		sel := l.AllowedRoutes.Namespaces.Selector
		if sel == nil {
			// we do validation separately that reports this as an error
			return false
		}

	case v1.NamespacesFromAll:
		return true
	}

	return ns == gw.Namespace
}
