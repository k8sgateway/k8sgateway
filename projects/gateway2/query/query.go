package query

import (
	"context"
	"fmt"
	"strings"

	"istio.io/istio/pkg/kube/krt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	apiv1 "sigs.k8s.io/gateway-api/apis/v1"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
	apiv1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"

	"github.com/solo-io/gloo/projects/gateway2/ir"
	"github.com/solo-io/gloo/projects/gateway2/krtcollections"
	"github.com/solo-io/gloo/projects/gateway2/wellknown"
)

var (
	ErrMissingReferenceGrant      = fmt.Errorf("missing reference grant")
	ErrUnknownBackendKind         = fmt.Errorf("unknown backend kind")
	ErrNoMatchingListenerHostname = fmt.Errorf("no matching listener hostname")
	ErrNoMatchingParent           = fmt.Errorf("no matching parent")
	ErrNotAllowedByListeners      = fmt.Errorf("not allowed by listeners")
	ErrLocalObjRefMissingKind     = fmt.Errorf("localObjRef provided with empty kind")
	ErrCyclicReference            = fmt.Errorf("cyclic reference detected while evaluating delegated routes")
	ErrUnresolvedReference        = fmt.Errorf("unresolved reference")
)

type Error struct {
	Reason apiv1.RouteConditionReason
	E      error
}

var _ error = &Error{}

// Error implements error.
func (e *Error) Error() string {
	return string(e.Reason)
}

func (e *Error) Unwrap() error {
	return e.E
}

type GroupKindNs struct {
	gk metav1.GroupKind
	ns string
}

func (g *GroupKindNs) GroupKind() (metav1.GroupKind, error) {
	return g.gk, nil
}

func (g *GroupKindNs) Namespace() string {
	return g.ns
}

func NewGroupKindNs(gk metav1.GroupKind, ns string) *GroupKindNs {
	return &GroupKindNs{
		gk: gk,
		ns: ns,
	}
}

type From interface {
	GroupKind() (metav1.GroupKind, error)
	Namespace() string
}

type FromObject struct {
	client.Object
	Scheme *runtime.Scheme
}

func (f FromObject) GroupKind() (metav1.GroupKind, error) {
	scheme := f.Scheme
	from := f.Object
	gvks, isUnversioned, err := scheme.ObjectKinds(from)
	var zero metav1.GroupKind
	if err != nil {
		return zero, fmt.Errorf("failed to get object kind %T", from)
	}
	if isUnversioned {
		return zero, fmt.Errorf("object of type %T is not versioned", from)
	}
	if len(gvks) != 1 {
		return zero, fmt.Errorf("ambigous gvks for %T, %v", f, gvks)
	}
	gvk := gvks[0]
	return metav1.GroupKind{Group: gvk.Group, Kind: gvk.Kind}, nil
}

func (f FromObject) Namespace() string {
	return f.GetNamespace()
}

// TODO(Law): remove this type entirely?
type FromGkNs struct {
	Gk metav1.GroupKind
	Ns string
}

func (f FromGkNs) GroupKind() (metav1.GroupKind, error) {
	return f.Gk, nil
}

func (f FromGkNs) Namespace() string {
	return f.Ns
}

type GatewayQueries interface {
	GetSecretForRef(kctx krt.HandlerContext, ctx context.Context, fromGk schema.GroupKind, fromns string, secretRef apiv1.SecretObjectReference) (*ir.Secret, error)

	// GetRoutesForGateway finds the top level xRoutes attached to the provided Gateway
	GetRoutesForGateway(kctx krt.HandlerContext, ctx context.Context, gw *gwv1.Gateway) (*RoutesForGwResult, error)
	// GetRouteChain resolves backends and delegated routes for a the provided xRoute object
	GetRouteChain(kctx krt.HandlerContext,
		ctx context.Context,
		route ir.Route,
		hostnames []string,
		parentRef gwv1.ParentReference) *RouteInfo
}

type RoutesForGwResult struct {
	// key is listener name
	ListenerResults map[string]*ListenerResult
	RouteErrors     []*RouteError
}

type ListenerResult struct {
	Error  error
	Routes []*RouteInfo
}

type RouteError struct {
	Route     ir.Route
	ParentRef apiv1.ParentReference
	Error     Error
}

type options struct {
}

type Option func(*options)

func NewData(
	routes *krtcollections.RoutesIndex,
	opts ...Option,
) GatewayQueries {
	builtOpts := &options{}
	for _, opt := range opts {
		opt(builtOpts)
	}
	return &gatewayQueries{
		routes: routes,
	}
}

// NewRoutesForGwResult creates and returns a new RoutesForGwResult with initialized fields.
func NewRoutesForGwResult() *RoutesForGwResult {
	return &RoutesForGwResult{
		ListenerResults: make(map[string]*ListenerResult),
		RouteErrors:     []*RouteError{},
	}
}

type gatewayQueries struct {
	routes     *krtcollections.RoutesIndex
	secrets    *krtcollections.SecretIndex
	namespaces krt.Collection[krtcollections.NamespaceMetadata]
	refGrants  *krtcollections.RefGrantIndex
}

func parentRefMatchListener(ref *apiv1.ParentReference, l *apiv1.Listener) bool {
	if ref != nil && ref.Port != nil && *ref.Port != l.Port {
		return false
	}
	if ref.SectionName != nil && *ref.SectionName != l.Name {
		return false
	}
	return true
}

// getParentRefsForGw extracts the ParentReferences from the provided object for the provided Gateway.
// Supported object types are:
//
//   - HTTPRoute
//   - TCPRoute
func getParentRefsForGw(gw *apiv1.Gateway, obj ir.Route) []apiv1.ParentReference {
	var ret []apiv1.ParentReference

	for _, pRef := range obj.GetParentRefs() {
		if isParentRefForGw(&pRef, gw, obj.GetNamespace()) {
			ret = append(ret, pRef)
		}
	}

	return ret
}

// isParentRefForGw checks if a ParentReference is associated with the provided Gateway.
func isParentRefForGw(pRef *apiv1.ParentReference, gw *apiv1.Gateway, defaultNs string) bool {
	if gw == nil || pRef == nil {
		return false
	}

	if pRef.Group != nil && *pRef.Group != apiv1.GroupName {
		return false
	}
	if pRef.Kind != nil && *pRef.Kind != wellknown.GatewayKind {
		return false
	}

	ns := defaultNs
	if pRef.Namespace != nil {
		ns = string(*pRef.Namespace)
	}

	return ns == gw.Namespace && string(pRef.Name) == gw.Name
}

func hostnameIntersect(l *apiv1.Listener, hr *ir.HttpRouteIR) (bool, []string) {
	var hostnames []string
	if l == nil || hr == nil {
		return false, hostnames
	}
	if l.Hostname == nil {
		for _, h := range hr.Hostnames {
			hostnames = append(hostnames, string(h))
		}
		return true, hostnames
	}
	var listenerHostname string = string(*l.Hostname)

	if strings.HasPrefix(listenerHostname, "*.") {
		if hr.Hostnames == nil {
			return true, []string{listenerHostname}
		}

		for _, hostname := range hr.Hostnames {
			hrHost := string(hostname)
			if strings.HasSuffix(hrHost, listenerHostname[1:]) {
				hostnames = append(hostnames, hrHost)
			}
		}
		return len(hostnames) > 0, hostnames
	} else {
		if len(hr.Hostnames) == 0 {
			return true, []string{listenerHostname}
		}
		for _, hostname := range hr.Hostnames {
			hrHost := string(hostname)
			if hrHost == listenerHostname {
				return true, []string{listenerHostname}
			}

			if strings.HasPrefix(hrHost, "*.") {
				if strings.HasSuffix(listenerHostname, hrHost[1:]) {
					return true, []string{listenerHostname}
				}
			}
			// also possible that listener hostname is more specific than the hr hostname
		}
	}

	return false, nil
}

func (r *gatewayQueries) GetSecretForRef(kctx krt.HandlerContext, ctx context.Context, fromGk schema.GroupKind, fromns string, secretRef apiv1.SecretObjectReference) (*ir.Secret, error) {
	f := krtcollections.From{
		GroupKind: fromGk,
		Namespace: fromns,
	}
	return r.secrets.GetSecret(kctx, f, secretRef)
}

func SameNamespace(ns string) func(krt.HandlerContext, string) bool {
	return func(_ krt.HandlerContext, s string) bool {
		return ns == s
	}
}

func AllNamespace() func(krt.HandlerContext, string) bool {
	return func(krt.HandlerContext, string) bool {
		return true
	}
}

func (r *gatewayQueries) NamespaceSelector(sel labels.Selector) func(krt.HandlerContext, string) bool {
	return func(kctx krt.HandlerContext, s string) bool {
		ns := krt.FetchOne(kctx, r.namespaces, krt.FilterKey(s))
		return sel.Matches(labels.Set(ns.Labels))
	}
}

func ReferenceAllowed(ctx context.Context, fromgk metav1.GroupKind, fromns string, togk metav1.GroupKind, toname string, grantsInToNs []apiv1beta1.ReferenceGrant) bool {
	for _, refGrant := range grantsInToNs {
		for _, from := range refGrant.Spec.From {
			if string(from.Namespace) != fromns {
				continue
			}
			if coreIfEmpty(fromgk.Group) == coreIfEmpty(string(from.Group)) && fromgk.Kind == string(from.Kind) {
				for _, to := range refGrant.Spec.To {
					if coreIfEmpty(togk.Group) == coreIfEmpty(string(to.Group)) && togk.Kind == string(to.Kind) {
						if to.Name == nil || string(*to.Name) == toname {
							return true
						}
					}
				}
			}
		}
	}
	return false
}

// Note that the spec has examples where the "core" api group is explicitly specified.
// so this helper function converts an empty string (which implies core api group) to the
// explicit "core" api group. It should only be used in places where the spec specifies that empty
// group means "core" api group (some place in the spec may default to the "gateway" api group instead.
func coreIfEmpty(s string) string {
	if s == "" {
		return "core"
	}
	return s
}
