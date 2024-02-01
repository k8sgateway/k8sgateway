// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/gloo/projects/gloo/pkg/api/v1 (interfaces: ProxyClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	clients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
	gomock "go.uber.org/mock/gomock"
)

// MockProxyClient is a mock of ProxyClient interface
type MockProxyClient struct {
	ctrl     *gomock.Controller
	recorder *MockProxyClientMockRecorder
}

// MockProxyClientMockRecorder is the mock recorder for MockProxyClient
type MockProxyClientMockRecorder struct {
	mock *MockProxyClient
}

// NewMockProxyClient creates a new mock instance
func NewMockProxyClient(ctrl *gomock.Controller) *MockProxyClient {
	mock := &MockProxyClient{ctrl: ctrl}
	mock.recorder = &MockProxyClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProxyClient) EXPECT() *MockProxyClientMockRecorder {
	return m.recorder
}

// BaseClient mocks base method
func (m *MockProxyClient) BaseClient() clients.ResourceClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseClient")
	ret0, _ := ret[0].(clients.ResourceClient)
	return ret0
}

// BaseClient indicates an expected call of BaseClient
func (mr *MockProxyClientMockRecorder) BaseClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseClient", reflect.TypeOf((*MockProxyClient)(nil).BaseClient))
}

// Delete mocks base method
func (m *MockProxyClient) Delete(arg0, arg1 string, arg2 clients.DeleteOpts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockProxyClientMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProxyClient)(nil).Delete), arg0, arg1, arg2)
}

// List mocks base method
func (m *MockProxyClient) List(arg0 string, arg1 clients.ListOpts) (v1.ProxyList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(v1.ProxyList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockProxyClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockProxyClient)(nil).List), arg0, arg1)
}

// Read mocks base method
func (m *MockProxyClient) Read(arg0, arg1 string, arg2 clients.ReadOpts) (*v1.Proxy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Proxy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockProxyClientMockRecorder) Read(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockProxyClient)(nil).Read), arg0, arg1, arg2)
}

// Register mocks base method
func (m *MockProxyClient) Register() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register")
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockProxyClientMockRecorder) Register() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockProxyClient)(nil).Register))
}

// Watch mocks base method
func (m *MockProxyClient) Watch(arg0 string, arg1 clients.WatchOpts) (<-chan v1.ProxyList, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(<-chan v1.ProxyList)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Watch indicates an expected call of Watch
func (mr *MockProxyClientMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockProxyClient)(nil).Watch), arg0, arg1)
}

// Write mocks base method
func (m *MockProxyClient) Write(arg0 *v1.Proxy, arg1 clients.WriteOpts) (*v1.Proxy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0, arg1)
	ret0, _ := ret[0].(*v1.Proxy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockProxyClientMockRecorder) Write(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockProxyClient)(nil).Write), arg0, arg1)
}
