// Code generated by MockGen. DO NOT EDIT.
// Source: clients.go

// Package mock_version is a generated GoMock package.
package mock_version

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	version "github.com/solo-io/gloo/projects/controller/pkg/api/grpc/version"
)

// MockServerVersion is a mock of ServerVersion interface.
type MockServerVersion struct {
	ctrl     *gomock.Controller
	recorder *MockServerVersionMockRecorder
}

// MockServerVersionMockRecorder is the mock recorder for MockServerVersion.
type MockServerVersionMockRecorder struct {
	mock *MockServerVersion
}

// NewMockServerVersion creates a new mock instance.
func NewMockServerVersion(ctrl *gomock.Controller) *MockServerVersion {
	mock := &MockServerVersion{ctrl: ctrl}
	mock.recorder = &MockServerVersionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServerVersion) EXPECT() *MockServerVersionMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockServerVersion) Get(ctx context.Context) ([]*version.ServerVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx)
	ret0, _ := ret[0].([]*version.ServerVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockServerVersionMockRecorder) Get(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockServerVersion)(nil).Get), ctx)
}

// GetClusterVersion mocks base method.
func (m *MockServerVersion) GetClusterVersion() (*version.KubernetesClusterVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterVersion")
	ret0, _ := ret[0].(*version.KubernetesClusterVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterVersion indicates an expected call of GetClusterVersion.
func (mr *MockServerVersionMockRecorder) GetClusterVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterVersion", reflect.TypeOf((*MockServerVersion)(nil).GetClusterVersion))
}
