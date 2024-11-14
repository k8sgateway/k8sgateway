// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/gloo/projects/controller/pkg/api/v1 (interfaces: EndpointClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/gloo/projects/controller/pkg/api/v1"
	clients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

// MockEndpointClient is a mock of EndpointClient interface
type MockEndpointClient struct {
	ctrl     *gomock.Controller
	recorder *MockEndpointClientMockRecorder
}

// MockEndpointClientMockRecorder is the mock recorder for MockEndpointClient
type MockEndpointClientMockRecorder struct {
	mock *MockEndpointClient
}

// NewMockEndpointClient creates a new mock instance
func NewMockEndpointClient(ctrl *gomock.Controller) *MockEndpointClient {
	mock := &MockEndpointClient{ctrl: ctrl}
	mock.recorder = &MockEndpointClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEndpointClient) EXPECT() *MockEndpointClientMockRecorder {
	return m.recorder
}

// BaseClient mocks base method
func (m *MockEndpointClient) BaseClient() clients.ResourceClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseClient")
	ret0, _ := ret[0].(clients.ResourceClient)
	return ret0
}

// BaseClient indicates an expected call of BaseClient
func (mr *MockEndpointClientMockRecorder) BaseClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseClient", reflect.TypeOf((*MockEndpointClient)(nil).BaseClient))
}

// Delete mocks base method
func (m *MockEndpointClient) Delete(arg0, arg1 string, arg2 clients.DeleteOpts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockEndpointClientMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockEndpointClient)(nil).Delete), arg0, arg1, arg2)
}

// List mocks base method
func (m *MockEndpointClient) List(arg0 string, arg1 clients.ListOpts) (v1.EndpointList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(v1.EndpointList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockEndpointClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockEndpointClient)(nil).List), arg0, arg1)
}

// Read mocks base method
func (m *MockEndpointClient) Read(arg0, arg1 string, arg2 clients.ReadOpts) (*v1.Endpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Endpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockEndpointClientMockRecorder) Read(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockEndpointClient)(nil).Read), arg0, arg1, arg2)
}

// Register mocks base method
func (m *MockEndpointClient) Register() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register")
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockEndpointClientMockRecorder) Register() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockEndpointClient)(nil).Register))
}

// Watch mocks base method
func (m *MockEndpointClient) Watch(arg0 string, arg1 clients.WatchOpts) (<-chan v1.EndpointList, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(<-chan v1.EndpointList)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Watch indicates an expected call of Watch
func (mr *MockEndpointClientMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockEndpointClient)(nil).Watch), arg0, arg1)
}

// Write mocks base method
func (m *MockEndpointClient) Write(arg0 *v1.Endpoint, arg1 clients.WriteOpts) (*v1.Endpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0, arg1)
	ret0, _ := ret[0].(*v1.Endpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockEndpointClientMockRecorder) Write(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockEndpointClient)(nil).Write), arg0, arg1)
}
