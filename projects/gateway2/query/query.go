package query

import (
	"context"
	"fmt"
	"strings"

	"github.com/solo-io/gloo/projects/gateway2/wellknown"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	apiv1 "sigs.k8s.io/gateway-api/apis/v1"
	apiv1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"
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

type From struct {
	GroupKind metav1.GroupKind
	Namespace string
}

func FromRoute(hr *apiv1.HTTPRoute) From {
	return From{
		GroupKind: metav1.GroupKind(wellknown.HTTPRouteGVK.GroupKind()),
		Namespace: hr.Namespace,
	}
}

func FromGateway(gw *apiv1.Gateway) From {
	return From{
		GroupKind: metav1.GroupKind(wellknown.GatewayGVK.GroupKind()),
		Namespace: gw.Namespace,
	}
}

type GatewayQueries interface {
	// Given a backendRef that resides in namespace obj, return the service that backs it.
	// This will error with `ErrMissingReferenceGrant` if there is no reference grant allowing the reference
	// return value depends on the group/kind in the backendRef.
	GetBackendForRef(ctx context.Context, obj From, backendRef *apiv1.BackendObjectReference) (client.Object, error)

	GetSecretForRef(ctx context.Context, obj From, secretRef apiv1.SecretObjectReference) (*corev1.Secret, error)

	GetLocalObjRef(ctx context.Context, from From, localObjRef apiv1.LocalObjectReference) (client.Object, error)

	// GetRoutesForGateway finds the top level HTTPRoutes attached to a Gateway
	GetRoutesForGateway(ctx context.Context, gw *apiv1.Gateway) (RoutesForGwResult, error)
	// GetHTTPRouteChain resolves backends and delegated routes for a HTTPRoute
	GetHTTPRouteChain(ctx context.Context, route apiv1.HTTPRoute, hostnames []string, parentRef apiv1.ParentReference) *HTTPRouteInfo
}

type RoutesForGwResult struct {
	// key is listener name
	ListenerResults map[string]*ListenerResult
	RouteErrors     []*RouteError
}

type ListenerResult struct {
	Error  error
	Routes []*HTTPRouteInfo
}

type RouteError struct {
	Route     apiv1.HTTPRoute
	ParentRef apiv1.ParentReference
	Error     Error
}

func NewData(c client.Client, scheme *runtime.Scheme) GatewayQueries {
	return &gatewayQueries{c, scheme}
}

type gatewayQueries struct {
	client client.Client
	scheme *runtime.Scheme
}

func (r *gatewayQueries) referenceAllowed(ctx context.Context, from metav1.GroupKind, fromns string, to metav1.GroupKind, tons, toname string) (bool, error) {
	var list apiv1beta1.ReferenceGrantList
	err := r.client.List(ctx, &list, client.InNamespace(tons), client.MatchingFieldsSelector{Selector: fields.OneTermEqualSelector(ReferenceGrantFromField, fromns)})
	if err != nil {
		return false, err
	}

	return ReferenceAllowed(ctx, from, fromns, to, toname, list.Items), nil
}

func parentRefMatchListener(ref apiv1.ParentReference, l *apiv1.Listener) bool {
	if ref.Port != nil && *ref.Port != l.Port {
		return false
	}
	if ref.SectionName != nil && *ref.SectionName != l.Name {
		return false
	}
	return true
}

func getParentRefsForGw(gw *apiv1.Gateway, hr *apiv1.HTTPRoute) []apiv1.ParentReference {
	var ret []apiv1.ParentReference
	for _, pRef := range hr.Spec.ParentRefs {

		if pRef.Group != nil && *pRef.Group != "gateway.networking.k8s.io" {
			continue
		}
		if pRef.Kind != nil && *pRef.Kind != "Gateway" {
			continue
		}
		ns := hr.Namespace
		if pRef.Namespace != nil {
			ns = string(*pRef.Namespace)
		}

		if ns == gw.Namespace && string(pRef.Name) == gw.Name {
			ret = append(ret, pRef)
		}
	}
	return ret
}

func hostnameIntersect(l *apiv1.Listener, hr *apiv1.HTTPRoute) (bool, []string) {
	var hostnames []string
	if l.Hostname == nil {
		for _, h := range hr.Spec.Hostnames {
			hostnames = append(hostnames, string(h))
		}
		return true, hostnames
	}
	var listenerHostname string = string(*l.Hostname)

	if strings.HasPrefix(listenerHostname, "*.") {
		if hr.Spec.Hostnames == nil {
			return true, []string{listenerHostname}
		}

		for _, hostname := range hr.Spec.Hostnames {
			hrHost := string(hostname)
			if strings.HasSuffix(hrHost, listenerHostname[1:]) {
				hostnames = append(hostnames, hrHost)
			}
		}
		return len(hostnames) > 0, hostnames
	} else {
		if len(hr.Spec.Hostnames) == 0 {
			return true, []string{listenerHostname}
		}
		for _, hostname := range hr.Spec.Hostnames {
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

func (r *gatewayQueries) GetSecretForRef(ctx context.Context, obj From, secretRef apiv1.SecretObjectReference) (*corev1.Secret, error) {
	secretGK := metav1.GroupKind{Group: corev1.GroupName, Kind: wellknown.SecretKind}
	refGk := secretGK
	if secretRef.Group != nil {
		refGk.Group = string(*secretRef.Group)
	}
	if secretRef.Kind != nil {
		refGk.Kind = string(*secretRef.Kind)
	}
	if refGk != secretGK {
		return nil, fmt.Errorf("only support core Secret references")
	}

	secretObj, err := r.getRef(ctx, obj, string(secretRef.Name), secretRef.Namespace, secretGK)
	if err != nil {
		return nil, err
	}
	secret, ok := secretObj.(*corev1.Secret)
	if !ok {
		return nil, fmt.Errorf("did not get corev1.Secret for %v", secretGK)
	}

	return secret, nil
}

func (r *gatewayQueries) GetLocalObjRef(ctx context.Context, obj From, localObjRef apiv1.LocalObjectReference) (client.Object, error) {
	refGroup := ""
	if localObjRef.Group != "" {
		refGroup = string(localObjRef.Group)
	}

	if localObjRef.Kind == "" {
		return nil, ErrLocalObjRefMissingKind
	}
	refKind := localObjRef.Kind

	localObjGK := metav1.GroupKind{Group: refGroup, Kind: string(refKind)}
	return r.getRef(ctx, obj, string(localObjRef.Name), nil, localObjGK)
}

func (r *gatewayQueries) GetBackendForRef(ctx context.Context, obj From, backend *apiv1.BackendObjectReference) (client.Object, error) {
	backendKind := "Service"
	backendGroup := ""

	if backend.Group != nil {
		backendGroup = string(*backend.Group)
	}
	if backend.Kind != nil {
		backendKind = string(*backend.Kind)
	}
	backendGK := metav1.GroupKind{Group: backendGroup, Kind: backendKind}

	return r.getRef(ctx, obj, string(backend.Name), backend.Namespace, backendGK)
}

func (r *gatewayQueries) getRef(ctx context.Context, from From, backendName string, backendNS *apiv1.Namespace, backendGK metav1.GroupKind) (client.Object, error) {
	fromNs := from.Namespace
	if fromNs == "" {
		fromNs = "default"
	}
	ns := fromNs
	if backendNS != nil {
		ns = string(*backendNS)
	}
	if ns != fromNs {
		fromgk := from.GroupKind
		// check if we're allowed to reference this namespace
		allowed, err := r.referenceAllowed(ctx, fromgk, fromNs, backendGK, ns, backendName)
		if err != nil {
			return nil, err
		}
		if !allowed {
			return nil, ErrMissingReferenceGrant
		}
	}

	gk := schema.GroupKind{Group: backendGK.Group, Kind: backendGK.Kind}

	versions := r.scheme.VersionsForGroupKind(gk)
	// versions are prioritized by order in the scheme, so we can just take the first one
	if len(versions) == 0 {
		return nil, ErrUnknownBackendKind
	}
	newObj, err := r.scheme.New(gk.WithVersion(versions[0].Version))
	if err != nil {
		return nil, err
	}
	ret, ok := newObj.(client.Object)
	if !ok {
		return nil, fmt.Errorf("new object is not a client.Object")
	}

	err = r.client.Get(ctx, types.NamespacedName{Namespace: ns, Name: backendName}, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func isHttpRouteAllowed(allowedKinds []metav1.GroupKind) bool {
	return isRouteAllowed(apiv1.GroupName, wellknown.HTTPRouteKind, allowedKinds)
}

func isRouteAllowed(group, kind string, allowedKinds []metav1.GroupKind) bool {
	for _, k := range allowedKinds {
		var allowedGroup string = k.Group
		if allowedGroup == "" {
			allowedGroup = apiv1.GroupName
		}

		if allowedGroup == group && k.Kind == kind {
			return true
		}
	}
	return false
}

func SameNamespace(ns string) func(string) bool {
	return func(s string) bool {
		return ns == s
	}
}

func AllNamespace() func(string) bool {
	return func(s string) bool {
		return true
	}
}

func (r *gatewayQueries) NamespaceSelector(sel labels.Selector) func(string) bool {
	return func(s string) bool {
		var ns corev1.Namespace
		r.client.Get(context.TODO(), types.NamespacedName{Name: s}, &ns)
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
