package krtcollections

import (
	"errors"

	"github.com/solo-io/gloo/projects/gateway2/model"
	"github.com/solo-io/gloo/projects/gateway2/translator/backendref"
	"istio.io/istio/pkg/kube/krt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
	gwv1a2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
	gwv1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

type UpstreamIndex struct {
	upstream krt.Collection[model.Upstream]
}

func NewUpstreamIndex(upstreams krt.Collection[model.Upstream]) *UpstreamIndex {
	return &UpstreamIndex{upstream: upstreams}
}

func (i *UpstreamIndex) GetUpstream(kctx krt.HandlerContext, gk schema.GroupKind, n types.NamespacedName) *model.Upstream {
	key := model.ObjectSource{
		Group:     gk.Group,
		Kind:      gk.Kind,
		Namespace: n.Namespace,
		Name:      n.Name,
	}
	return krt.FetchOne(kctx, i.upstream, krt.FilterKey(key.ResourceName()))
}

func (i *UpstreamIndex) GetUpstreamFromRef(kctx krt.HandlerContext, localns string, ref gwv1.BackendObjectReference) *model.Upstream {
	group := ""
	if ref.Group != nil {
		group = string(*ref.Group)
	}
	kind := "Service"
	if ref.Kind != nil {
		kind = string(*ref.Kind)
	}
	ns := localns
	if ref.Namespace != nil {
		ns = string(*ref.Namespace)
	}
	gk := schema.GroupKind{
		Group: group,
		Kind:  kind,
	}
	return i.GetUpstream(kctx, gk, types.NamespacedName{Namespace: ns, Name: string(ref.Name)})
}

type targetRefIndexKey struct {
	model.PolicyTargetRef
	Namespace string
}

type PolicyIndex struct {
	policies       krt.Collection[model.PolicyWrapper]
	targetRefIndex krt.Index[targetRefIndexKey, model.PolicyWrapper]
}

func NewPolicyIndex(policies krt.Collection[model.PolicyWrapper]) *PolicyIndex {
	targetRefIndex := krt.NewIndex(policies, func(p model.PolicyWrapper) []targetRefIndexKey {
		ret := make([]targetRefIndexKey, len(p.TargetRefs))
		for i, tr := range p.TargetRefs {
			ret[i] = targetRefIndexKey{
				PolicyTargetRef: tr,
				Namespace:       p.Namespace,
			}
		}
		return ret
	})
	return &PolicyIndex{policies: policies, targetRefIndex: targetRefIndex}
}

func (p *PolicyIndex) GetTargetingPolicies(kctx krt.HandlerContext, ref model.ObjectSource, sectionName string) []model.PolicyWrapper {
	// no need for ref grants here as target refs are namespace local
	targetRefIndexKey := targetRefIndexKey{
		PolicyTargetRef: model.PolicyTargetRef{
			Group: ref.Group,
			Kind:  ref.Kind,
			Name:  ref.Name,
		},
		Namespace: ref.Namespace,
	}
	return krt.Fetch(kctx, p.policies, krt.FilterIndex(p.targetRefIndex, targetRefIndexKey))
}

func (p *PolicyIndex) FetchPolicy(kctx krt.HandlerContext, ref model.ObjectSource) *model.PolicyWrapper {
	return krt.FetchOne(kctx, p.policies, krt.FilterKey(ref.ResourceName()))
}

type refGrantIndexKey struct {
	RefGrantNs string
	ToGK       metav1.GroupKind
	ToName     string
	FromGK     metav1.GroupKind
	FromNs     string
}
type RefGrantIndex struct {
	refgrants     krt.Collection[*gwv1beta1.ReferenceGrant]
	refGrantIndex krt.Index[refGrantIndexKey, *gwv1beta1.ReferenceGrant]
}

func NewRefGrantIndex(refgrants krt.Collection[*gwv1beta1.ReferenceGrant]) *RefGrantIndex {
	refGrantIndex := krt.NewIndex(refgrants, func(p *gwv1beta1.ReferenceGrant) []refGrantIndexKey {
		ret := make([]refGrantIndexKey, 0, len(p.Spec.To)*len(p.Spec.From))
		for _, from := range p.Spec.From {
			for _, to := range p.Spec.To {

				ret = append(ret, refGrantIndexKey{
					RefGrantNs: p.Namespace,
					ToGK:       metav1.GroupKind{Group: string(to.Group), Kind: string(to.Kind)},
					ToName:     strOr(to.Name, ""),
					FromGK:     metav1.GroupKind{Group: string(from.Group), Kind: string(from.Kind)},
					FromNs:     string(from.Namespace),
				})
			}
		}
		return ret
	})
	return &RefGrantIndex{refgrants: refgrants, refGrantIndex: refGrantIndex}
}

func (r *RefGrantIndex) ReferenceAllowed(kctx krt.HandlerContext, fromgk metav1.GroupKind, fromns string, to model.ObjectSource) bool {
	key := refGrantIndexKey{
		RefGrantNs: to.Namespace,
		ToGK:       metav1.GroupKind{Group: to.Group, Kind: to.Kind},
		FromGK:     fromgk,
		FromNs:     fromns,
	}
	if krt.Fetch(kctx, r.refgrants, krt.FilterIndex(r.refGrantIndex, key)) != nil {
		return true
	}
	// try with name:
	key.ToName = to.Name
	if krt.Fetch(kctx, r.refgrants, krt.FilterIndex(r.refGrantIndex, key)) != nil {
		return true
	}
	return false
}

type RouteWrapper struct {
	Route model.Route
}

func (c RouteWrapper) ResourceName() string {
	os := model.ObjectSource{
		Group:     c.Route.GetGroupKind().Group,
		Kind:      c.Route.GetGroupKind().Kind,
		Namespace: c.Route.GetNamespace(),
		Name:      c.Route.GetName(),
	}
	return os.ResourceName()
}

func (c RouteWrapper) Equals(in RouteWrapper) bool {
	return c.ResourceName() == in.ResourceName() && versionEquals(c.Route.GetSourceObject(), in.Route.GetSourceObject())
}
func versionEquals(a, b metav1.Object) bool {
	var versionEquals bool
	if a.GetGeneration() != 0 && b.GetGeneration() != 0 {
		versionEquals = a.GetGeneration() == b.GetGeneration()
	} else {
		versionEquals = a.GetResourceVersion() == b.GetResourceVersion()
	}
	return versionEquals && a.GetUID() == b.GetUID()
}

type RoutesIndex struct {
	routes          krt.Collection[RouteWrapper]
	httpRoutes      krt.Collection[model.HttpRouteIR]
	httpByNamespace krt.Index[string, model.HttpRouteIR]
	byTargetRef     krt.Index[types.NamespacedName, RouteWrapper]

	policies  *PolicyIndex
	refgrants *RefGrantIndex
	upstreams *UpstreamIndex
}

func NewRoutes(httproutes krt.Collection[*gwv1.HTTPRoute], tcproutes krt.Collection[*gwv1a2.TCPRoute], policies *PolicyIndex, upstreams *UpstreamIndex, refgrants *RefGrantIndex) *RoutesIndex {

	h := &RoutesIndex{policies: policies, refgrants: refgrants, upstreams: upstreams}
	h.httpRoutes = krt.NewCollection(httproutes, h.transformHttpRoute)
	hr := krt.NewCollection(h.httpRoutes, func(kctx krt.HandlerContext, i model.HttpRouteIR) *RouteWrapper {
		return &RouteWrapper{Route: &i}
	})
	h.routes = krt.JoinCollection([]krt.Collection[RouteWrapper]{hr})

	httpByNamespace := krt.NewIndex(h.httpRoutes, func(i model.HttpRouteIR) []string {
		return []string{i.GetNamespace()}
	})
	byTargetRef := krt.NewIndex(h.routes, func(in RouteWrapper) []types.NamespacedName {
		parentRefs := in.Route.GetParentRefs()
		ret := make([]types.NamespacedName, len(parentRefs))
		for i, pRef := range parentRefs {
			ns := strOr(pRef.Namespace, "")
			if ns == "" {
				ns = in.Route.GetNamespace()
			}
			ret[i] = types.NamespacedName{Namespace: ns, Name: string(pRef.Name)}
		}
		return ret
	})
	h.httpByNamespace = httpByNamespace
	h.byTargetRef = byTargetRef
	panic("TODO: implement tcp routes")
	return h
}

func (h *RoutesIndex) ListHttp(kctx krt.HandlerContext, ns string) []model.HttpRouteIR {
	return krt.Fetch(kctx, h.httpRoutes, krt.FilterIndex(h.httpByNamespace, ns))
}

func (h *RoutesIndex) RoutesForGateway(kctx krt.HandlerContext, nns types.NamespacedName) []model.Route {
	rts := krt.Fetch(kctx, h.routes, krt.FilterIndex(h.byTargetRef, nns))
	ret := make([]model.Route, len(rts))
	for i, r := range rts {
		ret[i] = r.Route
	}
	return ret
}

func (h *RoutesIndex) FetchHttp(kctx krt.HandlerContext, n, ns string) *model.HttpRouteIR {
	// TODO: maybe the key shouldnt include g and k?
	src := model.ObjectSource{
		Group:     gwv1.SchemeGroupVersion.Group,
		Kind:      "HTTPRoute",
		Namespace: ns,
		Name:      n,
	}
	return krt.FetchOne(kctx, h.httpRoutes, krt.FilterKey(src.ResourceName()))
}

func (h *RoutesIndex) Fetch(kctx krt.HandlerContext, gk schema.GroupKind, n, ns string) *RouteWrapper {
	// TODO: maybe the key shouldnt include g and k?
	src := model.ObjectSource{
		Group:     gk.Group,
		Kind:      gk.Kind,
		Namespace: ns,
		Name:      n,
	}
	return krt.FetchOne(kctx, h.routes, krt.FilterKey(src.ResourceName()))
}

func (h *RoutesIndex) transformHttpRoute(kctx krt.HandlerContext, i *gwv1.HTTPRoute) *model.HttpRouteIR {
	src := model.ObjectSource{
		Group:     gwv1.SchemeGroupVersion.Group,
		Kind:      "HTTPRoute",
		Namespace: i.Namespace,
		Name:      i.Name,
	}

	return &model.HttpRouteIR{
		ObjectSource:     src,
		SourceObject:     i,
		ParentRefs:       i.Spec.ParentRefs,
		Hostnames:        tostr(i.Spec.Hostnames),
		Rules:            h.transformRules(kctx, src, i.Spec.Rules),
		AttachedPolicies: toAttachedPolicies(h.policies.GetTargetingPolicies(kctx, src, "")),
	}
}
func (h *RoutesIndex) transformRules(kctx krt.HandlerContext, src model.ObjectSource, i []gwv1.HTTPRouteRule) []model.HttpRouteRuleIR {
	rules := make([]model.HttpRouteRuleIR, 0, len(i))
	for _, r := range i {

		extensionRefs := h.getExtensionRefs(kctx, src.Namespace, r)
		var policies model.AttachedPolicies
		if r.Name != nil {
			policies = toAttachedPolicies(h.policies.GetTargetingPolicies(kctx, src, string(*r.Name)))
		}

		rules = append(rules, model.HttpRouteRuleIR{
			HttpRouteRuleCommonIR: model.HttpRouteRuleCommonIR{
				SourceRule:       &r,
				ExtensionRefs:    extensionRefs,
				AttachedPolicies: policies,
			},
			Backends: h.getBackends(kctx, src, r.BackendRefs),
			Matches:  r.Matches,
			Name:     emptyIfNil(r.Name),
		})
	}
	return rules

}
func (h *RoutesIndex) getExtensionRefs(kctx krt.HandlerContext, ns string, r gwv1.HTTPRouteRule) model.AttachedPolicies {
	ret := model.AttachedPolicies{
		Policies: map[schema.GroupKind][]model.PolicyAtt{},
	}
	for _, ext := range r.Filters {
		if ext.ExtensionRef == nil {
			continue
		}
		ref := *ext.ExtensionRef
		gk := schema.GroupKind{
			Group: string(ref.Group),
			Kind:  string(ref.Kind),
		}
		key := model.ObjectSource{
			Group:     string(ref.Group),
			Kind:      string(ref.Kind),
			Namespace: ns,
			Name:      string(ref.Name),
		}
		policy := h.policies.FetchPolicy(kctx, key)
		if policy != nil {
			ret.Policies[gk] = append(ret.Policies[gk], model.PolicyAtt{PolicyIr: policy /*direct attachment - no target ref*/})
		}

	}
	return ret
}
func (h *RoutesIndex) getExtensionRefs2(kctx krt.HandlerContext, ns string, r []gwv1.HTTPRouteFilter) model.AttachedPolicies {
	ret := model.AttachedPolicies{
		Policies: map[schema.GroupKind][]model.PolicyAtt{},
	}
	for _, ext := range r {
		if ext.ExtensionRef == nil {
			panic("TODO: handle built in extensions")
			continue
		}
		ref := *ext.ExtensionRef
		gk := schema.GroupKind{
			Group: string(ref.Group),
			Kind:  string(ref.Kind),
		}
		key := model.ObjectSource{
			Group:     string(ref.Group),
			Kind:      string(ref.Kind),
			Namespace: ns,
			Name:      string(ref.Name),
		}
		policy := h.policies.FetchPolicy(kctx, key)
		if policy != nil {
			ret.Policies[gk] = append(ret.Policies[gk], model.PolicyAtt{PolicyIr: policy /*direct attachment - no target ref*/})
		}

	}
	return ret
}

func (h *RoutesIndex) getBackends(kctx krt.HandlerContext, src model.ObjectSource, i []gwv1.HTTPBackendRef) []model.HttpBackendOrDelegate {
	backends := make([]model.HttpBackendOrDelegate, 0, len(i))
	for _, ref := range i {
		extensionRefs := h.getExtensionRefs2(kctx, src.Namespace, ref.Filters)
		fromns := src.Namespace

		to := model.ObjectSource{
			Group:     strOr(ref.BackendRef.Group, ""),
			Kind:      strOr(ref.BackendRef.Kind, "Service"),
			Namespace: strOr(ref.BackendRef.Namespace, fromns),
			Name:      string(ref.BackendRef.Name),
		}
		if backendref.RefIsHTTPRoute(ref.BackendRef.BackendObjectReference) {
			backends = append(backends, model.HttpBackendOrDelegate{
				Delegate:         &to,
				AttachedPolicies: extensionRefs,
			})
			continue
		}

		var upstream *model.Upstream
		fromgk := metav1.GroupKind{
			Group: src.Group,
			Kind:  src.Kind,
		}
		var err error
		if h.refgrants.ReferenceAllowed(kctx, fromgk, fromns, to) {
			upstream = h.upstreams.GetUpstreamFromRef(kctx, src.Namespace, ref.BackendRef.BackendObjectReference)
		} else {
			err = errors.New("ErrMissingReferenceGrant")
			panic("TODO: use a common error type")
		}
		clusterName := "blackhole-cluster"
		if upstream != nil {
			panic("TODO: figure out cluster name")
			//			clusterName = model.UpstreamToClusterName(upstream)
		}
		backends = append(backends, model.HttpBackendOrDelegate{
			Backend: &model.Backend{
				Upstream:    upstream,
				ClusterName: clusterName,
				Weight:      weight(ref.Weight),
				Err:         err,
			},
			AttachedPolicies: extensionRefs,
		})
	}
	return backends
}

type GwIndex struct {
	routes    krt.Collection[model.GatewayIR]
	policies  *PolicyIndex
	refgrants *RefGrantIndex
}

//func NewGwIndex(gws krt.Collection[*gwv1.Gateway], policies *PolicyIndex, refgrants *RefGrantIndex) *HttpRoutesIndex {
//	h := &HttpRoutesIndex{policies: policies, refgrants: refgrants}
//	h.routes = krt.NewCollection(gws, h.transformGw)
//	return h
//}
//
//func (h *HttpRoutesIndex) transformGw(kctx krt.HandlerContext, i *gwv1.Gateway) *model.GatewayWithPoliciesIR {
//	src := model.ObjectSource{
//		Group:     gwv1.SchemeGroupVersion.Group,
//		Kind:      "Gateway",
//		Namespace: i.Namespace,
//		Name:      i.Name,
//	}
//
//	return &model.HttpRouteIR{
//		ObjectSource:     src,
//		SourceObject:     i,
//		ParentRefs:       i.Spec.ParentRefs,
//		Hostnames:        tostr(i.Spec.Hostnames),
//		Rules:            h.transformRules(kctx, src, i.Spec.Rules),
//		AttachedPolicies: toAttachedPolicies(h.policies.GetTargetingPolicies(kctx, src, "")),
//	}
//}

func strOr[T ~string](s *T, def string) string {
	if s == nil {
		return def
	}
	return string(*s)
}

func weight(w *int32) uint32 {
	if w == nil {
		return 1
	}
	return uint32(*w)
}

func toAttachedPolicies(policies []model.PolicyWrapper) model.AttachedPolicies {
	ret := model.AttachedPolicies{
		Policies: map[schema.GroupKind][]model.PolicyAtt{},
	}
	for _, p := range policies {
		gk := schema.GroupKind{
			Group: p.Group,
			Kind:  p.Kind,
		}
		ret.Policies[gk] = append(ret.Policies[gk], model.PolicyAtt{PolicyIr: p.PolicyIR})
	}
	return ret
}

func emptyIfNil(s *gwv1.SectionName) string {
	if s == nil {
		return ""
	}
	return string(*s)
}

func tostr(in []gwv1.Hostname) []string {
	out := make([]string, len(in))
	for i, h := range in {
		out[i] = string(h)
	}
	return out
}
