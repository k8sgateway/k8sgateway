package krtquery

import (
	"context"

	"github.com/solo-io/gloo/projects/gateway2/krtcollection"
	"github.com/solo-io/gloo/projects/gateway2/query"
	"github.com/solo-io/gloo/projects/gateway2/wellknown"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
	gwv1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"

	glookubev1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/kube/apis/gloo.solo.io/v1"

	"istio.io/istio/pkg/config/schema/gvr"
	"istio.io/istio/pkg/kube"
	"istio.io/istio/pkg/kube/kclient"
	"istio.io/istio/pkg/kube/krt"
	"istio.io/istio/pkg/ptr"
)

type Queries struct {
	// k8s top level
	Services   krt.Collection[*v1.Service]
	Secrets    krt.Collection[*v1.Secret]
	Namespaces krt.Collection[*v1.Namespace]

	// gwapi top level
	HTTPRoutes                 krt.Collection[*gwv1.HTTPRoute]
	httpRoutesByNamespace      krt.Index[string, *gwv1.HTTPRoute]
	Gateways                   krt.Collection[*gwv1.Gateway]
	GatewayClasses             krt.Collection[*gwv1.GatewayClass]
	ReferenceGrants            krt.Collection[*gwv1beta1.ReferenceGrant]
	referenceGrantsByNamespace krt.Index[string, *gwv1beta1.ReferenceGrant]

	// gloo top level (for now we use kube types for reference grant purposes)
	Upstreams krt.Collection[*glookubev1.Upstream]

	// derived collections
	GatewayHTTPRoutes IndexedAttachments[*query.HTTPRouteInfo, *gwv1.Gateway]
}

func New(ctx context.Context, client kube.Client) *Queries {
	filter := kclient.Filter{}

	Services := krt.WrapClient(kclient.NewFiltered[*v1.Service](client, filter), krt.WithName("Services"))
	Secrets := krt.WrapClient(kclient.NewFiltered[*v1.Secret](client, filter), krt.WithName("Secrets"))
	Namespaces := krt.WrapClient(kclient.NewFiltered[*v1.Namespace](client, filter), krt.WithName("Namespaces"))

	HTTPRoutes := krtcollection.SetupCollectionDynamic[gwv1.HTTPRoute](ctx, client, gvr.HTTPRoute_v1, krt.WithName("HTTPRoutes"))
	httpRoutesByNamespace := krt.NewNamespaceIndex(HTTPRoutes)

	Gateways := krtcollection.SetupCollectionDynamic[gwv1.Gateway](ctx, client, gvr.KubernetesGateway_v1, krt.WithName("Gateways"))
	GatewayClasses := krtcollection.SetupCollectionDynamic[gwv1.GatewayClass](ctx, client, gvr.KubernetesGateway_v1, krt.WithName("GatewayClasses"))

	ReferenceGrants := krtcollection.SetupCollectionDynamic[gwv1beta1.ReferenceGrant](ctx, client, gvr.ReferenceGrant, krt.WithName("ReferenceGrants"))
	referenceGrantsByNamespace := krt.NewNamespaceIndex(ReferenceGrants)

	Upstreams := krtcollection.SetupCollectionDynamic[glookubev1.Upstream](
		ctx,
		client,
		glookubev1.SchemeGroupVersion.WithResource("upstreams"),
		krt.WithName("RawUpstreams"),
	)

	q := Queries{
		Services:   Services,
		Secrets:    Secrets,
		Namespaces: Namespaces,

		// TODO GatewayAPI types shouldn't need dynamic client
		HTTPRoutes:                 HTTPRoutes,
		httpRoutesByNamespace:      httpRoutesByNamespace,
		Gateways:                   Gateways,
		GatewayClasses:             GatewayClasses,
		ReferenceGrants:            ReferenceGrants,
		referenceGrantsByNamespace: referenceGrantsByNamespace,

		Upstreams: Upstreams,
	}

	// build derived collections in sub-constructors
	// using the root collections from the struct
	q.GatewayHTTPRoutes = q.routesForGateway(ctx)

	return &q
}

// supported GroupKind for GetRef.
// There is no way to make this generic over the entire Scheme with krt.
var (
	ServiceGK   = metav1.GroupKind{Group: v1.GroupName, Kind: wellknown.ServiceKind}
	SercretGK   = metav1.GroupKind{Group: v1.GroupName, Kind: wellknown.SecretKind}
	HTTPRouteGK = metav1.GroupKind(wellknown.HTTPRouteGVK.GroupKind())
	UpstreamGK  = metav1.GroupKind(wellknown.HTTPRouteGVK.GroupKind())
)

// GetGrantedRef allows resolving arbitrary references to objects.
// This method will NOT enforce ReferenceGrants.
func (q *Queries) GetUngrantedRef(
	ctx krt.HandlerContext,
	targetGk metav1.GroupKind,
	targetName string,
	targetNs string,
) (client.Object, error) {
	return q.getRef(ctx, targetGk, targetName, targetNs)
}

// GetGrantedRef allows resolving arbitrary references to objects.
// This method will enforce ReferenceGrants.
func (q *Queries) GetGrantedRef(
	ctx krt.HandlerContext,
	fromGk metav1.GroupKind, fromNamespace string,
	targetGk metav1.GroupKind, targetName string, targetNs *string,
) (client.Object, error) {
	namespace := ptr.OrDefault(targetNs, fromNamespace)
	if !q.referenceAllowed(
		ctx,
		fromGk,
		fromNamespace,
		targetGk,
		targetName,
		namespace,
	) {
		return nil, query.ErrMissingReferenceGrant
	}
	return q.getRef(ctx, targetGk, targetName, namespace)
}

func (q *Queries) getRef(
	ctx krt.HandlerContext,
	gk metav1.GroupKind,
	name,
	ns string,
) (client.Object, error) {
	key := types.NamespacedName{
		Name:      name,
		Namespace: ns,
	}.String()

	var resolvedObj client.Object
	switch gk {
	case ServiceGK:
		resolvedObj = ptr.Flatten(krt.FetchOne(ctx, q.Services, krt.FilterKey(key)))
	case HTTPRouteGK:
		resolvedObj = ptr.Flatten(krt.FetchOne(ctx, q.Services, krt.FilterKey(key)))
	case UpstreamGK:
		resolvedObj = ptr.Flatten(krt.FetchOne(ctx, q.Services, krt.FilterKey(key)))
	}
	if resolvedObj == nil {
		return nil, query.ErrUnresolvedReference
	}
	return resolvedObj, nil
}

func (q *Queries) referenceAllowed(
	ctx krt.HandlerContext,
	fromGk metav1.GroupKind,
	fromNamespace string,
	targetGk metav1.GroupKind,
	targetName, targetNs string,
) bool {
	if targetNs == fromNamespace {
		return true
	}

	// fetch grants in the targeted namespace
	refGrants := krt.Fetch(ctx, q.ReferenceGrants, krt.FilterIndex(q.referenceGrantsByNamespace, targetNs))
	for _, refGrant := range refGrants {
		for _, grantFrom := range refGrant.Spec.From {
			if string(grantFrom.Namespace) != fromNamespace {
				continue
			}
			if !gkMatch(fromGk, grantFrom.Group, grantFrom.Kind) {
				continue
			}

			for _, grantTo := range refGrant.Spec.To {
				if gkMatch(targetGk, grantTo.Group, grantTo.Kind) {
					if grantTo.Name == nil || string(*grantTo.Name) == targetName {
						return true
					}
				}
			}

		}
	}
	return false
}

func gkMatch(gk metav1.GroupKind, grantGroup gwv1.Group, grantKind gwv1.Kind) bool {
	aGroup := coreIfEmpty(gk.Group)
	bGroup := coreIfEmpty(string(grantGroup))
	return aGroup == bGroup && gk.Kind == string(grantKind)
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
