// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/gloo/projects/gateway2/query (interfaces: GatewayQueries)
//
// Generated by this command:
//
//	mockgen -destination mocks/mock_queries.go -package mocks github.com/solo-io/gloo/projects/gateway2/query GatewayQueries
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	query "github.com/solo-io/gloo/projects/gateway2/query"
	gomock "github.com/golang/mock/gomock"
	client "sigs.k8s.io/controller-runtime/pkg/client"
	v1 "sigs.k8s.io/gateway-api/apis/v1"
)

// MockGatewayQueries is a mock of GatewayQueries interface.
type MockGatewayQueries struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayQueriesMockRecorder
}

// MockGatewayQueriesMockRecorder is the mock recorder for MockGatewayQueries.
type MockGatewayQueriesMockRecorder struct {
	mock *MockGatewayQueries
}

// NewMockGatewayQueries creates a new mock instance.
func NewMockGatewayQueries(ctrl *gomock.Controller) *MockGatewayQueries {
	mock := &MockGatewayQueries{ctrl: ctrl}
	mock.recorder = &MockGatewayQueriesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGatewayQueries) EXPECT() *MockGatewayQueriesMockRecorder {
	return m.recorder
}

// GetBackendForRef mocks base method.
func (m *MockGatewayQueries) GetBackendForRef(arg0 context.Context, arg1 query.From, arg2 *v1.BackendObjectReference) (client.Object, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBackendForRef", arg0, arg1, arg2)
	ret0, _ := ret[0].(client.Object)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBackendForRef indicates an expected call of GetBackendForRef.
func (mr *MockGatewayQueriesMockRecorder) GetBackendForRef(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackendForRef", reflect.TypeOf((*MockGatewayQueries)(nil).GetBackendForRef), arg0, arg1, arg2)
}

// GetLocalObjRef mocks base method.
func (m *MockGatewayQueries) GetLocalObjRef(arg0 context.Context, arg1 query.From, arg2 v1.LocalObjectReference) (client.Object, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLocalObjRef", arg0, arg1, arg2)
	ret0, _ := ret[0].(client.Object)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLocalObjRef indicates an expected call of GetLocalObjRef.
func (mr *MockGatewayQueriesMockRecorder) GetLocalObjRef(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLocalObjRef", reflect.TypeOf((*MockGatewayQueries)(nil).GetLocalObjRef), arg0, arg1, arg2)
}

// GetRoutesForGw mocks base method.
func (m *MockGatewayQueries) GetRoutesForGw(arg0 context.Context, arg1 *v1.Gateway) (query.RoutesForGwResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoutesForGw", arg0, arg1)
	ret0, _ := ret[0].(query.RoutesForGwResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoutesForGw indicates an expected call of GetRoutesForGw.
func (mr *MockGatewayQueriesMockRecorder) GetRoutesForGw(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoutesForGw", reflect.TypeOf((*MockGatewayQueries)(nil).GetRoutesForGw), arg0, arg1)
}

// GetSecretForRef mocks base method.
func (m *MockGatewayQueries) GetSecretForRef(arg0 context.Context, arg1 query.From, arg2 v1.SecretObjectReference) (client.Object, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretForRef", arg0, arg1, arg2)
	ret0, _ := ret[0].(client.Object)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretForRef indicates an expected call of GetSecretForRef.
func (mr *MockGatewayQueriesMockRecorder) GetSecretForRef(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretForRef", reflect.TypeOf((*MockGatewayQueries)(nil).GetSecretForRef), arg0, arg1, arg2)
}

// ObjToFrom mocks base method.
func (m *MockGatewayQueries) ObjToFrom(arg0 client.Object) query.From {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObjToFrom", arg0)
	ret0, _ := ret[0].(query.From)
	return ret0
}

// ObjToFrom indicates an expected call of ObjToFrom.
func (mr *MockGatewayQueriesMockRecorder) ObjToFrom(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjToFrom", reflect.TypeOf((*MockGatewayQueries)(nil).ObjToFrom), arg0)
}
