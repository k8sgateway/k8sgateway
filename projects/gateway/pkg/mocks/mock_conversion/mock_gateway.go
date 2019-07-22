// Code generated by MockGen. DO NOT EDIT.
// Source: ./projects/gateway/pkg/conversion/gateway.go

// Package mock_conversion is a generated GoMock package.
package mock_conversion

import (
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	v2alpha1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v2alpha1"
	reflect "reflect"
)

// MockGatewayConverter is a mock of GatewayConverter interface
type MockGatewayConverter struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayConverterMockRecorder
}

// MockGatewayConverterMockRecorder is the mock recorder for MockGatewayConverter
type MockGatewayConverterMockRecorder struct {
	mock *MockGatewayConverter
}

// NewMockGatewayConverter creates a new mock instance
func NewMockGatewayConverter(ctrl *gomock.Controller) *MockGatewayConverter {
	mock := &MockGatewayConverter{ctrl: ctrl}
	mock.recorder = &MockGatewayConverterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGatewayConverter) EXPECT() *MockGatewayConverterMockRecorder {
	return m.recorder
}

// FromV1ToV2alpha1 mocks base method
func (m *MockGatewayConverter) FromV1ToV2alpha1(src *v1.Gateway) *v2alpha1.Gateway {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FromV1ToV2alpha1", src)
	ret0, _ := ret[0].(*v2alpha1.Gateway)
	return ret0
}

// FromV1ToV2alpha1 indicates an expected call of FromV1ToV2alpha1
func (mr *MockGatewayConverterMockRecorder) FromV1ToV2alpha1(src interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FromV1ToV2alpha1", reflect.TypeOf((*MockGatewayConverter)(nil).FromV1ToV2alpha1), src)
}
