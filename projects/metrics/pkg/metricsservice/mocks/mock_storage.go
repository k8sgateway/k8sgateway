// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/gloo/projects/metrics/pkg/metricsservice (interfaces: StorageClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	metricsservice "github.com/solo-io/gloo/projects/metrics/pkg/metricsservice"
	reflect "reflect"
)

// MockStorageClient is a mock of StorageClient interface
type MockStorageClient struct {
	ctrl     *gomock.Controller
	recorder *MockStorageClientMockRecorder
}

// MockStorageClientMockRecorder is the mock recorder for MockStorageClient
type MockStorageClientMockRecorder struct {
	mock *MockStorageClient
}

// NewMockStorageClient creates a new mock instance
func NewMockStorageClient(ctrl *gomock.Controller) *MockStorageClient {
	mock := &MockStorageClient{ctrl: ctrl}
	mock.recorder = &MockStorageClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorageClient) EXPECT() *MockStorageClientMockRecorder {
	return m.recorder
}

// GetUsage mocks base method
func (m *MockStorageClient) GetUsage() (*metricsservice.GlobalUsage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsage")
	ret0, _ := ret[0].(*metricsservice.GlobalUsage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsage indicates an expected call of GetUsage
func (mr *MockStorageClientMockRecorder) GetUsage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsage", reflect.TypeOf((*MockStorageClient)(nil).GetUsage))
}

// RecordUsage mocks base method
func (m *MockStorageClient) RecordUsage(arg0 *metricsservice.GlobalUsage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordUsage", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecordUsage indicates an expected call of RecordUsage
func (mr *MockStorageClientMockRecorder) RecordUsage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordUsage", reflect.TypeOf((*MockStorageClient)(nil).RecordUsage), arg0)
}
