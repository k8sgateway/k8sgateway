// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/gloo/projects/gateway/pkg/api/v1 (interfaces: RouteTableClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	clients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

// MockRouteTableClient is a mock of RouteTableClient interface
type MockRouteTableClient struct {
	ctrl     *gomock.Controller
	recorder *MockRouteTableClientMockRecorder
}

// MockRouteTableClientMockRecorder is the mock recorder for MockRouteTableClient
type MockRouteTableClientMockRecorder struct {
	mock *MockRouteTableClient
}

// NewMockRouteTableClient creates a new mock instance
func NewMockRouteTableClient(ctrl *gomock.Controller) *MockRouteTableClient {
	mock := &MockRouteTableClient{ctrl: ctrl}
	mock.recorder = &MockRouteTableClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRouteTableClient) EXPECT() *MockRouteTableClientMockRecorder {
	return m.recorder
}

// BaseClient mocks base method
func (m *MockRouteTableClient) BaseClient() clients.ResourceClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseClient")
	ret0, _ := ret[0].(clients.ResourceClient)
	return ret0
}

// BaseClient indicates an expected call of BaseClient
func (mr *MockRouteTableClientMockRecorder) BaseClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseClient", reflect.TypeOf((*MockRouteTableClient)(nil).BaseClient))
}

// Delete mocks base method
func (m *MockRouteTableClient) Delete(arg0, arg1 string, arg2 clients.DeleteOpts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockRouteTableClientMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRouteTableClient)(nil).Delete), arg0, arg1, arg2)
}

// List mocks base method
func (m *MockRouteTableClient) List(arg0 string, arg1 clients.ListOpts) (v1.RouteTableList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(v1.RouteTableList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockRouteTableClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRouteTableClient)(nil).List), arg0, arg1)
}

// Read mocks base method
func (m *MockRouteTableClient) Read(arg0, arg1 string, arg2 clients.ReadOpts) (*v1.RouteTable, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.RouteTable)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockRouteTableClientMockRecorder) Read(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockRouteTableClient)(nil).Read), arg0, arg1, arg2)
}

// Register mocks base method
func (m *MockRouteTableClient) Register() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register")
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockRouteTableClientMockRecorder) Register() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockRouteTableClient)(nil).Register))
}

// Watch mocks base method
func (m *MockRouteTableClient) Watch(arg0 string, arg1 clients.WatchOpts) (<-chan v1.RouteTableList, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(<-chan v1.RouteTableList)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Watch indicates an expected call of Watch
func (mr *MockRouteTableClientMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockRouteTableClient)(nil).Watch), arg0, arg1)
}

// Write mocks base method
func (m *MockRouteTableClient) Write(arg0 *v1.RouteTable, arg1 clients.WriteOpts) (*v1.RouteTable, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0, arg1)
	ret0, _ := ret[0].(*v1.RouteTable)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockRouteTableClientMockRecorder) Write(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockRouteTableClient)(nil).Write), arg0, arg1)
}
