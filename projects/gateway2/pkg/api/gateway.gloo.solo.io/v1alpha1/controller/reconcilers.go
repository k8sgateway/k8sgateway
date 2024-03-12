// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./reconcilers.go -destination mocks/reconcilers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	gateway_gloo_solo_io_v1alpha1 "github.com/solo-io/gloo/projects/gateway2/pkg/api/gateway.gloo.solo.io/v1alpha1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the GatewayConfig Resource.
// implemented by the user
type GatewayConfigReconciler interface {
	ReconcileGatewayConfig(obj *gateway_gloo_solo_io_v1alpha1.GatewayConfig) (reconcile.Result, error)
}

// Reconcile deletion events for the GatewayConfig Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type GatewayConfigDeletionReconciler interface {
	ReconcileGatewayConfigDeletion(req reconcile.Request) error
}

type GatewayConfigReconcilerFuncs struct {
	OnReconcileGatewayConfig         func(obj *gateway_gloo_solo_io_v1alpha1.GatewayConfig) (reconcile.Result, error)
	OnReconcileGatewayConfigDeletion func(req reconcile.Request) error
}

func (f *GatewayConfigReconcilerFuncs) ReconcileGatewayConfig(obj *gateway_gloo_solo_io_v1alpha1.GatewayConfig) (reconcile.Result, error) {
	if f.OnReconcileGatewayConfig == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileGatewayConfig(obj)
}

func (f *GatewayConfigReconcilerFuncs) ReconcileGatewayConfigDeletion(req reconcile.Request) error {
	if f.OnReconcileGatewayConfigDeletion == nil {
		return nil
	}
	return f.OnReconcileGatewayConfigDeletion(req)
}

// Reconcile and finalize the GatewayConfig Resource
// implemented by the user
type GatewayConfigFinalizer interface {
	GatewayConfigReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	GatewayConfigFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeGatewayConfig(obj *gateway_gloo_solo_io_v1alpha1.GatewayConfig) error
}

type GatewayConfigReconcileLoop interface {
	RunGatewayConfigReconciler(ctx context.Context, rec GatewayConfigReconciler, predicates ...predicate.Predicate) error
}

type gatewayConfigReconcileLoop struct {
	loop reconcile.Loop
}

func NewGatewayConfigReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) GatewayConfigReconcileLoop {
	return &gatewayConfigReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &gateway_gloo_solo_io_v1alpha1.GatewayConfig{}, options),
	}
}

func (c *gatewayConfigReconcileLoop) RunGatewayConfigReconciler(ctx context.Context, reconciler GatewayConfigReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericGatewayConfigReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(GatewayConfigFinalizer); ok {
		reconcilerWrapper = genericGatewayConfigFinalizer{
			genericGatewayConfigReconciler: genericReconciler,
			finalizingReconciler:           finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericGatewayConfigHandler implements a generic reconcile.Reconciler
type genericGatewayConfigReconciler struct {
	reconciler GatewayConfigReconciler
}

func (r genericGatewayConfigReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*gateway_gloo_solo_io_v1alpha1.GatewayConfig)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: GatewayConfig handler received event for %T", object)
	}
	return r.reconciler.ReconcileGatewayConfig(obj)
}

func (r genericGatewayConfigReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(GatewayConfigDeletionReconciler); ok {
		return deletionReconciler.ReconcileGatewayConfigDeletion(request)
	}
	return nil
}

// genericGatewayConfigFinalizer implements a generic reconcile.FinalizingReconciler
type genericGatewayConfigFinalizer struct {
	genericGatewayConfigReconciler
	finalizingReconciler GatewayConfigFinalizer
}

func (r genericGatewayConfigFinalizer) FinalizerName() string {
	return r.finalizingReconciler.GatewayConfigFinalizerName()
}

func (r genericGatewayConfigFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*gateway_gloo_solo_io_v1alpha1.GatewayConfig)
	if !ok {
		return errors.Errorf("internal error: GatewayConfig handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeGatewayConfig(obj)
}
