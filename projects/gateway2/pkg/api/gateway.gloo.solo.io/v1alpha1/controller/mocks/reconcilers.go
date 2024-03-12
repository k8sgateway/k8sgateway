// Code generated by MockGen. DO NOT EDIT.
// Source: ./reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/gloo/projects/gateway2/pkg/api/gateway.gloo.solo.io/v1alpha1"
	controller "github.com/solo-io/gloo/projects/gateway2/pkg/api/gateway.gloo.solo.io/v1alpha1/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockGatewayConfigReconciler is a mock of GatewayConfigReconciler interface.
type MockGatewayConfigReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayConfigReconcilerMockRecorder
}

// MockGatewayConfigReconcilerMockRecorder is the mock recorder for MockGatewayConfigReconciler.
type MockGatewayConfigReconcilerMockRecorder struct {
	mock *MockGatewayConfigReconciler
}

// NewMockGatewayConfigReconciler creates a new mock instance.
func NewMockGatewayConfigReconciler(ctrl *gomock.Controller) *MockGatewayConfigReconciler {
	mock := &MockGatewayConfigReconciler{ctrl: ctrl}
	mock.recorder = &MockGatewayConfigReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGatewayConfigReconciler) EXPECT() *MockGatewayConfigReconcilerMockRecorder {
	return m.recorder
}

// ReconcileGatewayConfig mocks base method.
func (m *MockGatewayConfigReconciler) ReconcileGatewayConfig(obj *v1alpha1.GatewayConfig) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileGatewayConfig", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileGatewayConfig indicates an expected call of ReconcileGatewayConfig.
func (mr *MockGatewayConfigReconcilerMockRecorder) ReconcileGatewayConfig(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileGatewayConfig", reflect.TypeOf((*MockGatewayConfigReconciler)(nil).ReconcileGatewayConfig), obj)
}

// MockGatewayConfigDeletionReconciler is a mock of GatewayConfigDeletionReconciler interface.
type MockGatewayConfigDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayConfigDeletionReconcilerMockRecorder
}

// MockGatewayConfigDeletionReconcilerMockRecorder is the mock recorder for MockGatewayConfigDeletionReconciler.
type MockGatewayConfigDeletionReconcilerMockRecorder struct {
	mock *MockGatewayConfigDeletionReconciler
}

// NewMockGatewayConfigDeletionReconciler creates a new mock instance.
func NewMockGatewayConfigDeletionReconciler(ctrl *gomock.Controller) *MockGatewayConfigDeletionReconciler {
	mock := &MockGatewayConfigDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockGatewayConfigDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGatewayConfigDeletionReconciler) EXPECT() *MockGatewayConfigDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileGatewayConfigDeletion mocks base method.
func (m *MockGatewayConfigDeletionReconciler) ReconcileGatewayConfigDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileGatewayConfigDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileGatewayConfigDeletion indicates an expected call of ReconcileGatewayConfigDeletion.
func (mr *MockGatewayConfigDeletionReconcilerMockRecorder) ReconcileGatewayConfigDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileGatewayConfigDeletion", reflect.TypeOf((*MockGatewayConfigDeletionReconciler)(nil).ReconcileGatewayConfigDeletion), req)
}

// MockGatewayConfigFinalizer is a mock of GatewayConfigFinalizer interface.
type MockGatewayConfigFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayConfigFinalizerMockRecorder
}

// MockGatewayConfigFinalizerMockRecorder is the mock recorder for MockGatewayConfigFinalizer.
type MockGatewayConfigFinalizerMockRecorder struct {
	mock *MockGatewayConfigFinalizer
}

// NewMockGatewayConfigFinalizer creates a new mock instance.
func NewMockGatewayConfigFinalizer(ctrl *gomock.Controller) *MockGatewayConfigFinalizer {
	mock := &MockGatewayConfigFinalizer{ctrl: ctrl}
	mock.recorder = &MockGatewayConfigFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGatewayConfigFinalizer) EXPECT() *MockGatewayConfigFinalizerMockRecorder {
	return m.recorder
}

// FinalizeGatewayConfig mocks base method.
func (m *MockGatewayConfigFinalizer) FinalizeGatewayConfig(obj *v1alpha1.GatewayConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeGatewayConfig", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeGatewayConfig indicates an expected call of FinalizeGatewayConfig.
func (mr *MockGatewayConfigFinalizerMockRecorder) FinalizeGatewayConfig(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeGatewayConfig", reflect.TypeOf((*MockGatewayConfigFinalizer)(nil).FinalizeGatewayConfig), obj)
}

// GatewayConfigFinalizerName mocks base method.
func (m *MockGatewayConfigFinalizer) GatewayConfigFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GatewayConfigFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GatewayConfigFinalizerName indicates an expected call of GatewayConfigFinalizerName.
func (mr *MockGatewayConfigFinalizerMockRecorder) GatewayConfigFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GatewayConfigFinalizerName", reflect.TypeOf((*MockGatewayConfigFinalizer)(nil).GatewayConfigFinalizerName))
}

// ReconcileGatewayConfig mocks base method.
func (m *MockGatewayConfigFinalizer) ReconcileGatewayConfig(obj *v1alpha1.GatewayConfig) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileGatewayConfig", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileGatewayConfig indicates an expected call of ReconcileGatewayConfig.
func (mr *MockGatewayConfigFinalizerMockRecorder) ReconcileGatewayConfig(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileGatewayConfig", reflect.TypeOf((*MockGatewayConfigFinalizer)(nil).ReconcileGatewayConfig), obj)
}

// MockGatewayConfigReconcileLoop is a mock of GatewayConfigReconcileLoop interface.
type MockGatewayConfigReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayConfigReconcileLoopMockRecorder
}

// MockGatewayConfigReconcileLoopMockRecorder is the mock recorder for MockGatewayConfigReconcileLoop.
type MockGatewayConfigReconcileLoopMockRecorder struct {
	mock *MockGatewayConfigReconcileLoop
}

// NewMockGatewayConfigReconcileLoop creates a new mock instance.
func NewMockGatewayConfigReconcileLoop(ctrl *gomock.Controller) *MockGatewayConfigReconcileLoop {
	mock := &MockGatewayConfigReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockGatewayConfigReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGatewayConfigReconcileLoop) EXPECT() *MockGatewayConfigReconcileLoopMockRecorder {
	return m.recorder
}

// RunGatewayConfigReconciler mocks base method.
func (m *MockGatewayConfigReconcileLoop) RunGatewayConfigReconciler(ctx context.Context, rec controller.GatewayConfigReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunGatewayConfigReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunGatewayConfigReconciler indicates an expected call of RunGatewayConfigReconciler.
func (mr *MockGatewayConfigReconcileLoopMockRecorder) RunGatewayConfigReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunGatewayConfigReconciler", reflect.TypeOf((*MockGatewayConfigReconcileLoop)(nil).RunGatewayConfigReconciler), varargs...)
}
