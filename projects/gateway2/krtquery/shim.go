package krtquery

import (
	"context"
	"fmt"

	"github.com/solo-io/gloo/projects/gateway2/query"
	"github.com/solo-io/gloo/projects/gateway2/wellknown"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"

	"istio.io/istio/pkg/kube/krt"
)

var _ query.GatewayQueries = &Queries{}

func (q *Queries) GetBackendForRef(ctx context.Context, fromObj query.From, backendRef *gwv1.BackendObjectReference) (client.Object, error) {
	krtctx, ok := ctx.Value(krtKey{}).(krt.HandlerContext)
	if !ok {
		return nil, NonKRTError
	}
	// backendRef defaults to Service
	targetGk := metav1.GroupKind{
		Group: v1.GroupName,
		Kind:  wellknown.ServiceKind,
	}
	if backendRef.Kind != nil {
		targetGk.Kind = string(*backendRef.Kind)
	}
	if backendRef.Group != nil {
		targetGk.Group = string(*backendRef.Group)
	}

	return q.GetGrantedRef(
		krtctx,
		fromObj.GroupKind,
		fromObj.Namespace,
		targetGk,
		string(backendRef.Name),
		nsPtrToString(backendRef.Namespace),
	)
}

// GetHTTPRouteChain resolves backends and delegated routes for a HTTPRoute
func (q *Queries) GetHTTPRouteChain(ctx context.Context, route gwv1.HTTPRoute, hostnames []string, parentRef gwv1.ParentReference) *query.HTTPRouteInfo {
	krtctx, ok := ctx.Value(krtKey{}).(krt.HandlerContext)
	if !ok {
		// TODO, just don't have a shim...
		panic("not in krt")
	}
	return q.HTTPRouteChain(krtctx, route, hostnames, parentRef)
}

func (q *Queries) GetRoutesForGateway(ctx context.Context, gw *gwv1.Gateway) (query.RoutesForGwResult, error) {
	krtctx, ok := ctx.Value(krtKey{}).(krt.HandlerContext)
	if !ok {
		return query.RoutesForGwResult{}, NonKRTError
	}
	return q.HTTPRoutesForGateway(krtctx, gw)
}

func (r *Queries) GetSecretForRef(ctx context.Context, from query.From, secretRef gwv1.SecretObjectReference) (*v1.Secret, error) {
	krtctx, err := GetKRT(ctx)
	if err != nil {
		return nil, err
	}

	refGk := SercretGK
	if secretRef.Group != nil && *secretRef.Group != gwv1.Group(SercretGK.Group) {
		refGk.Group = string(*secretRef.Group)
	}
	if secretRef.Kind != nil {
		refGk.Kind = string(*secretRef.Kind)
	}
	if refGk != SercretGK {
		return nil, fmt.Errorf("only support core Secret references")
	}

	obj, err := r.GetGrantedRef(
		krtctx,
		from.GroupKind, from.Namespace,
		SercretGK, string(secretRef.Name), nsPtrToString(secretRef.Namespace),
	)
	if err != nil {
		return nil, err
	}
	return obj.(*v1.Secret), nil
}

func (r *Queries) GetLocalObjRef(ctx context.Context, fromObj query.From, localObjRef gwv1.LocalObjectReference) (client.Object, error) {
	krtctx, err := GetKRT(ctx)
	if err != nil {
		return nil, err
	}

	refGroup := ""
	if localObjRef.Group != "" {
		refGroup = string(localObjRef.Group)
	}

	if localObjRef.Kind == "" {
		return nil, query.ErrLocalObjRefMissingKind
	}
	refKind := localObjRef.Kind

	localObjGK := metav1.GroupKind{Group: refGroup, Kind: string(refKind)}
	return r.GetUngrantedRef(krtctx, localObjGK, string(localObjRef.Name), fromObj.Namespace)
}

func nsPtrToString(ns *gwv1.Namespace) *string {
	if ns == nil {
		return nil
	}
	converted := string(*ns)
	return &converted
}
