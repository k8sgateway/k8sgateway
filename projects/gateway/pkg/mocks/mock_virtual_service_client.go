// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/gloo/projects/gateway/pkg/api/v1 (interfaces: VirtualServiceClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	clients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

// MockVirtualServiceClient is a mock of VirtualServiceClient interface
type MockVirtualServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualServiceClientMockRecorder
}

// MockVirtualServiceClientMockRecorder is the mock recorder for MockVirtualServiceClient
type MockVirtualServiceClientMockRecorder struct {
	mock *MockVirtualServiceClient
}

// NewMockVirtualServiceClient creates a new mock instance
func NewMockVirtualServiceClient(ctrl *gomock.Controller) *MockVirtualServiceClient {
	mock := &MockVirtualServiceClient{ctrl: ctrl}
	mock.recorder = &MockVirtualServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVirtualServiceClient) EXPECT() *MockVirtualServiceClientMockRecorder {
	return m.recorder
}

// BaseClient mocks base method
func (m *MockVirtualServiceClient) BaseClient() clients.ResourceClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseClient")
	ret0, _ := ret[0].(clients.ResourceClient)
	return ret0
}

// BaseClient indicates an expected call of BaseClient
func (mr *MockVirtualServiceClientMockRecorder) BaseClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseClient", reflect.TypeOf((*MockVirtualServiceClient)(nil).BaseClient))
}

// Delete mocks base method
func (m *MockVirtualServiceClient) Delete(arg0, arg1 string, arg2 clients.DeleteOpts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockVirtualServiceClientMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVirtualServiceClient)(nil).Delete), arg0, arg1, arg2)
}

// List mocks base method
func (m *MockVirtualServiceClient) List(arg0 string, arg1 clients.ListOpts) (v1.VirtualServiceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(v1.VirtualServiceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockVirtualServiceClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVirtualServiceClient)(nil).List), arg0, arg1)
}

// Read mocks base method
func (m *MockVirtualServiceClient) Read(arg0, arg1 string, arg2 clients.ReadOpts) (*v1.VirtualService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.VirtualService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockVirtualServiceClientMockRecorder) Read(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockVirtualServiceClient)(nil).Read), arg0, arg1, arg2)
}

// Register mocks base method
func (m *MockVirtualServiceClient) Register() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register")
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockVirtualServiceClientMockRecorder) Register() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockVirtualServiceClient)(nil).Register))
}

// Watch mocks base method
func (m *MockVirtualServiceClient) Watch(arg0 string, arg1 clients.WatchOpts) (<-chan v1.VirtualServiceList, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(<-chan v1.VirtualServiceList)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Watch indicates an expected call of Watch
func (mr *MockVirtualServiceClientMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockVirtualServiceClient)(nil).Watch), arg0, arg1)
}

// Write mocks base method
func (m *MockVirtualServiceClient) Write(arg0 *v1.VirtualService, arg1 clients.WriteOpts) (*v1.VirtualService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0, arg1)
	ret0, _ := ret[0].(*v1.VirtualService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockVirtualServiceClientMockRecorder) Write(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockVirtualServiceClient)(nil).Write), arg0, arg1)
}
