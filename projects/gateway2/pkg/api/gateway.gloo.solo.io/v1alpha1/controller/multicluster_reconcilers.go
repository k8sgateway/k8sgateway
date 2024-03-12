// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller

import (
	"context"

	gateway_gloo_solo_io_v1alpha1 "github.com/solo-io/gloo/projects/gateway2/pkg/api/gateway.gloo.solo.io/v1alpha1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the GatewayConfig Resource across clusters.
// implemented by the user
type MulticlusterGatewayConfigReconciler interface {
	ReconcileGatewayConfig(clusterName string, obj *gateway_gloo_solo_io_v1alpha1.GatewayConfig) (reconcile.Result, error)
}

// Reconcile deletion events for the GatewayConfig Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterGatewayConfigDeletionReconciler interface {
	ReconcileGatewayConfigDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterGatewayConfigReconcilerFuncs struct {
	OnReconcileGatewayConfig         func(clusterName string, obj *gateway_gloo_solo_io_v1alpha1.GatewayConfig) (reconcile.Result, error)
	OnReconcileGatewayConfigDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterGatewayConfigReconcilerFuncs) ReconcileGatewayConfig(clusterName string, obj *gateway_gloo_solo_io_v1alpha1.GatewayConfig) (reconcile.Result, error) {
	if f.OnReconcileGatewayConfig == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileGatewayConfig(clusterName, obj)
}

func (f *MulticlusterGatewayConfigReconcilerFuncs) ReconcileGatewayConfigDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileGatewayConfigDeletion == nil {
		return nil
	}
	return f.OnReconcileGatewayConfigDeletion(clusterName, req)
}

type MulticlusterGatewayConfigReconcileLoop interface {
	// AddMulticlusterGatewayConfigReconciler adds a MulticlusterGatewayConfigReconciler to the MulticlusterGatewayConfigReconcileLoop.
	AddMulticlusterGatewayConfigReconciler(ctx context.Context, rec MulticlusterGatewayConfigReconciler, predicates ...predicate.Predicate)
}

type multiclusterGatewayConfigReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterGatewayConfigReconcileLoop) AddMulticlusterGatewayConfigReconciler(ctx context.Context, rec MulticlusterGatewayConfigReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericGatewayConfigMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterGatewayConfigReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterGatewayConfigReconcileLoop {
	return &multiclusterGatewayConfigReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &gateway_gloo_solo_io_v1alpha1.GatewayConfig{}, options)}
}

type genericGatewayConfigMulticlusterReconciler struct {
	reconciler MulticlusterGatewayConfigReconciler
}

func (g genericGatewayConfigMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterGatewayConfigDeletionReconciler); ok {
		return deletionReconciler.ReconcileGatewayConfigDeletion(cluster, req)
	}
	return nil
}

func (g genericGatewayConfigMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*gateway_gloo_solo_io_v1alpha1.GatewayConfig)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: GatewayConfig handler received event for %T", object)
	}
	return g.reconciler.ReconcileGatewayConfig(cluster, obj)
}
