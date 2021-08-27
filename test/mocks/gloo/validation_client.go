// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/gloo/projects/gloo/pkg/api/grpc/validation (interfaces: GlooValidationServiceClient)

// Package mock_validation is a generated GoMock package.
package mock_validation

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	validation "github.com/solo-io/gloo/projects/gloo/pkg/api/grpc/validation"
	grpc "google.golang.org/grpc"
)

// MockGlooValidationServiceClient is a mock of GlooValidationServiceClient interface
type MockGlooValidationServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockGlooValidationServiceClientMockRecorder
}

// MockGlooValidationServiceClientMockRecorder is the mock recorder for MockGlooValidationServiceClient
type MockGlooValidationServiceClientMockRecorder struct {
	mock *MockGlooValidationServiceClient
}

// NewMockGlooValidationServiceClient creates a new mock instance
func NewMockGlooValidationServiceClient(ctrl *gomock.Controller) *MockGlooValidationServiceClient {
	mock := &MockGlooValidationServiceClient{ctrl: ctrl}
	mock.recorder = &MockGlooValidationServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGlooValidationServiceClient) EXPECT() *MockGlooValidationServiceClientMockRecorder {
	return m.recorder
}

// NotifyOnResync mocks base method
func (m *MockGlooValidationServiceClient) NotifyOnResync(arg0 context.Context, arg1 *validation.NotifyOnResyncRequest, arg2 ...grpc.CallOption) (validation.GlooValidationService_NotifyOnResyncClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NotifyOnResync", varargs...)
	ret0, _ := ret[0].(validation.GlooValidationService_NotifyOnResyncClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NotifyOnResync indicates an expected call of NotifyOnResync
func (mr *MockGlooValidationServiceClientMockRecorder) NotifyOnResync(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyOnResync", reflect.TypeOf((*MockGlooValidationServiceClient)(nil).NotifyOnResync), varargs...)
}

// Validate mocks base method
func (m *MockGlooValidationServiceClient) Validate(arg0 context.Context, arg1 *validation.GlooValidationServiceRequest, arg2 ...grpc.CallOption) (*validation.GlooValidationServiceResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Validate", varargs...)
	ret0, _ := ret[0].(*validation.GlooValidationServiceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Validate indicates an expected call of Validate
func (mr *MockGlooValidationServiceClientMockRecorder) Validate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockGlooValidationServiceClient)(nil).Validate), varargs...)
}
