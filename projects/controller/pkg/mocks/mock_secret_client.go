// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/gloo/projects/controller/pkg/api/v1 (interfaces: SecretClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/gloo/projects/controller/pkg/api/v1"
	clients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

// MockSecretClient is a mock of SecretClient interface
type MockSecretClient struct {
	ctrl     *gomock.Controller
	recorder *MockSecretClientMockRecorder
}

// MockSecretClientMockRecorder is the mock recorder for MockSecretClient
type MockSecretClientMockRecorder struct {
	mock *MockSecretClient
}

// NewMockSecretClient creates a new mock instance
func NewMockSecretClient(ctrl *gomock.Controller) *MockSecretClient {
	mock := &MockSecretClient{ctrl: ctrl}
	mock.recorder = &MockSecretClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSecretClient) EXPECT() *MockSecretClientMockRecorder {
	return m.recorder
}

// BaseClient mocks base method
func (m *MockSecretClient) BaseClient() clients.ResourceClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseClient")
	ret0, _ := ret[0].(clients.ResourceClient)
	return ret0
}

// BaseClient indicates an expected call of BaseClient
func (mr *MockSecretClientMockRecorder) BaseClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseClient", reflect.TypeOf((*MockSecretClient)(nil).BaseClient))
}

// Delete mocks base method
func (m *MockSecretClient) Delete(arg0, arg1 string, arg2 clients.DeleteOpts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockSecretClientMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSecretClient)(nil).Delete), arg0, arg1, arg2)
}

// List mocks base method
func (m *MockSecretClient) List(arg0 string, arg1 clients.ListOpts) (v1.SecretList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(v1.SecretList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockSecretClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSecretClient)(nil).List), arg0, arg1)
}

// Read mocks base method
func (m *MockSecretClient) Read(arg0, arg1 string, arg2 clients.ReadOpts) (*v1.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockSecretClientMockRecorder) Read(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockSecretClient)(nil).Read), arg0, arg1, arg2)
}

// Register mocks base method
func (m *MockSecretClient) Register() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register")
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockSecretClientMockRecorder) Register() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockSecretClient)(nil).Register))
}

// Watch mocks base method
func (m *MockSecretClient) Watch(arg0 string, arg1 clients.WatchOpts) (<-chan v1.SecretList, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(<-chan v1.SecretList)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Watch indicates an expected call of Watch
func (mr *MockSecretClientMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockSecretClient)(nil).Watch), arg0, arg1)
}

// Write mocks base method
func (m *MockSecretClient) Write(arg0 *v1.Secret, arg1 clients.WriteOpts) (*v1.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0, arg1)
	ret0, _ := ret[0].(*v1.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockSecretClientMockRecorder) Write(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockSecretClient)(nil).Write), arg0, arg1)
}
