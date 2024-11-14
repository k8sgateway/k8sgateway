// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/gloo/projects/controller/pkg/api/v1 (interfaces: SettingsClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/gloo/projects/controller/pkg/api/v1"
	clients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

// MockSettingsClient is a mock of SettingsClient interface
type MockSettingsClient struct {
	ctrl     *gomock.Controller
	recorder *MockSettingsClientMockRecorder
}

// MockSettingsClientMockRecorder is the mock recorder for MockSettingsClient
type MockSettingsClientMockRecorder struct {
	mock *MockSettingsClient
}

// NewMockSettingsClient creates a new mock instance
func NewMockSettingsClient(ctrl *gomock.Controller) *MockSettingsClient {
	mock := &MockSettingsClient{ctrl: ctrl}
	mock.recorder = &MockSettingsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSettingsClient) EXPECT() *MockSettingsClientMockRecorder {
	return m.recorder
}

// BaseClient mocks base method
func (m *MockSettingsClient) BaseClient() clients.ResourceClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseClient")
	ret0, _ := ret[0].(clients.ResourceClient)
	return ret0
}

// BaseClient indicates an expected call of BaseClient
func (mr *MockSettingsClientMockRecorder) BaseClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseClient", reflect.TypeOf((*MockSettingsClient)(nil).BaseClient))
}

// Delete mocks base method
func (m *MockSettingsClient) Delete(arg0, arg1 string, arg2 clients.DeleteOpts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockSettingsClientMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSettingsClient)(nil).Delete), arg0, arg1, arg2)
}

// List mocks base method
func (m *MockSettingsClient) List(arg0 string, arg1 clients.ListOpts) (v1.SettingsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(v1.SettingsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockSettingsClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSettingsClient)(nil).List), arg0, arg1)
}

// Read mocks base method
func (m *MockSettingsClient) Read(arg0, arg1 string, arg2 clients.ReadOpts) (*v1.Settings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Settings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockSettingsClientMockRecorder) Read(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockSettingsClient)(nil).Read), arg0, arg1, arg2)
}

// Register mocks base method
func (m *MockSettingsClient) Register() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register")
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockSettingsClientMockRecorder) Register() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockSettingsClient)(nil).Register))
}

// Watch mocks base method
func (m *MockSettingsClient) Watch(arg0 string, arg1 clients.WatchOpts) (<-chan v1.SettingsList, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(<-chan v1.SettingsList)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Watch indicates an expected call of Watch
func (mr *MockSettingsClientMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockSettingsClient)(nil).Watch), arg0, arg1)
}

// Write mocks base method
func (m *MockSettingsClient) Write(arg0 *v1.Settings, arg1 clients.WriteOpts) (*v1.Settings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0, arg1)
	ret0, _ := ret[0].(*v1.Settings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockSettingsClientMockRecorder) Write(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockSettingsClient)(nil).Write), arg0, arg1)
}
