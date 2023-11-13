// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/gloo/projects/gateway/pkg/api/v1 (interfaces: GatewayClient)

// Package mocks is a generated GoMock package.
package mock_v1

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	clients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

// MockGatewayClient is a mock of GatewayClient interface
type MockGatewayClient struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayClientMockRecorder
}

// MockGatewayClientMockRecorder is the mock recorder for MockGatewayClient
type MockGatewayClientMockRecorder struct {
	mock *MockGatewayClient
}

// NewMockGatewayClient creates a new mock instance
func NewMockGatewayClient(ctrl *gomock.Controller) *MockGatewayClient {
	mock := &MockGatewayClient{ctrl: ctrl}
	mock.recorder = &MockGatewayClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGatewayClient) EXPECT() *MockGatewayClientMockRecorder {
	return m.recorder
}

// BaseClient mocks base method
func (m *MockGatewayClient) BaseClient() clients.ResourceClient {
	ret := m.ctrl.Call(m, "BaseClient")
	ret0, _ := ret[0].(clients.ResourceClient)
	return ret0
}

// BaseClient indicates an expected call of BaseClient
func (mr *MockGatewayClientMockRecorder) BaseClient() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseClient", reflect.TypeOf((*MockGatewayClient)(nil).BaseClient))
}

// Delete mocks base method
func (m *MockGatewayClient) Delete(arg0, arg1 string, arg2 clients.DeleteOpts) error {
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockGatewayClientMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockGatewayClient)(nil).Delete), arg0, arg1, arg2)
}

// List mocks base method
func (m *MockGatewayClient) List(arg0 string, arg1 clients.ListOpts) (v1.GatewayList, error) {
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(v1.GatewayList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockGatewayClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockGatewayClient)(nil).List), arg0, arg1)
}

// Read mocks base method
func (m *MockGatewayClient) Read(arg0, arg1 string, arg2 clients.ReadOpts) (*v1.Gateway, error) {
	ret := m.ctrl.Call(m, "Read", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Gateway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockGatewayClientMockRecorder) Read(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockGatewayClient)(nil).Read), arg0, arg1, arg2)
}

// Register mocks base method
func (m *MockGatewayClient) Register() error {
	ret := m.ctrl.Call(m, "Register")
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockGatewayClientMockRecorder) Register() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockGatewayClient)(nil).Register))
}

// Watch mocks base method
func (m *MockGatewayClient) Watch(arg0 string, arg1 clients.WatchOpts) (<-chan v1.GatewayList, <-chan error, error) {
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(<-chan v1.GatewayList)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Watch indicates an expected call of Watch
func (mr *MockGatewayClientMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockGatewayClient)(nil).Watch), arg0, arg1)
}

// Write mocks base method
func (m *MockGatewayClient) Write(arg0 *v1.Gateway, arg1 clients.WriteOpts) (*v1.Gateway, error) {
	ret := m.ctrl.Call(m, "Write", arg0, arg1)
	ret0, _ := ret[0].(*v1.Gateway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockGatewayClientMockRecorder) Write(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockGatewayClient)(nil).Write), arg0, arg1)
}
