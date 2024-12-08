package krtcollections

import (
	"errors"

	"github.com/solo-io/gloo/projects/gateway2/model"
	"istio.io/istio/pkg/kube/krt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
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

type HttpRoutesIndex struct {
	routes    krt.Collection[model.HttpRouteIR]
	policies  *PolicyIndex
	refgrants *RefGrantIndex
	upstreams *UpstreamIndex
}

func NewHttpRoutes(httproutes krt.Collection[*gwv1.HTTPRoute], policies *PolicyIndex, refgrants *RefGrantIndex) *HttpRoutesIndex {
	h := &HttpRoutesIndex{policies: policies, refgrants: refgrants}
	h.routes = krt.NewCollection(httproutes, h.transformRoute)
	return h
}

func (h *HttpRoutesIndex) transformRoute(kctx krt.HandlerContext, i *gwv1.HTTPRoute) *model.HttpRouteIR {
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
func (h *HttpRoutesIndex) transformRules(kctx krt.HandlerContext, src model.ObjectSource, i []gwv1.HTTPRouteRule) []model.HttpRouteRuleIR {
	rules := make([]model.HttpRouteRuleIR, 0, len(i))
	for j, r := range i {
		matches := r.Matches
		if len(matches) == 0 {
			matches = []gwv1.HTTPRouteMatch{{}}
		}

		extensionRefs := h.getExtensionRefs(kctx, src.Namespace, r)
		var policies model.AttachedPolicies[model.HttpPolicy]
		if r.Name != nil {
			policies = toAttachedPolicies(h.policies.GetTargetingPolicies(kctx, src, string(*r.Name)))
		}

		for _, m := range matches {
			rules = append(rules, model.HttpRouteRuleIR{
				Match:            m,
				MatchIndex:       j,
				SourceRule:       &r,
				Name:             emptyIfNil(r.Name),
				ExtensionRefs:    extensionRefs,
				AttachedPolicies: policies,
				Backends:         h.getBackends(kctx, src, r.BackendRefs),
			})
		}
	}
	return rules

}
func (h *HttpRoutesIndex) getExtensionRefs(kctx krt.HandlerContext, ns string, r gwv1.HTTPRouteRule) model.AttachedPolicies[model.HttpPolicy] {
	ret := model.AttachedPolicies[model.HttpPolicy]{
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
			ret.Policies[gk] = append(ret.Policies[gk], model.PolicyAtt{PolicyIr: policy.PolicyIr /*direct attachment - no target ref*/})
		}

	}
	return ret
}
func (h *HttpRoutesIndex) getExtensionRefs2(kctx krt.HandlerContext, ns string, r []gwv1.HTTPRouteFilter) model.AttachedPolicies[model.HttpBackendPolicy] {
	ret := model.AttachedPolicies[model.HttpBackendPolicy]{
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
			ret.Policies[gk] = append(ret.Policies[gk], model.PolicyAtt{PolicyIr: policy.PolicyIr /*direct attachment - no target ref*/})
		}

	}
	return ret
}

func (h *HttpRoutesIndex) getBackends(kctx krt.HandlerContext, src model.ObjectSource, i []gwv1.HTTPBackendRef) []model.HttpBackend {
	backends := make([]model.HttpBackend, 0, len(i))
	for _, ref := range i {
		var upstream *model.Upstream
		fromgk := metav1.GroupKind{
			Group: src.Group,
			Kind:  src.Kind,
		}
		fromns := src.Namespace
		to := model.ObjectSource{
			Group:     strOr(ref.BackendRef.Group, ""),
			Kind:      strOr(ref.BackendRef.Kind, "Service"),
			Namespace: strOr(ref.BackendRef.Namespace, fromns),
			Name:      string(ref.BackendRef.Name),
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
		extensionRefs := h.getExtensionRefs2(kctx, src.Namespace, ref.Filters)
		backends = append(backends, model.HttpBackend{
			Backend: model.Backend{
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

func toAttachedPolicies(policies []model.PolicyWrapper) model.AttachedPolicies[model.HttpPolicy] {
	ret := model.AttachedPolicies[model.HttpPolicy]{
		Policies: map[schema.GroupKind][]model.PolicyAtt{},
	}
	for _, p := range policies {
		gk := schema.GroupKind{
			Group: p.Group,
			Kind:  p.Kind,
		}
		ret.Policies[gk] = append(ret.Policies[gk], model.PolicyAtt{PolicyIr: p.PolicyIr})
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
