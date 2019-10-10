// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/gloo/pkg/utils/selectionutils (interfaces: VirtualServiceSelector)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	reflect "reflect"
)

// MockVirtualServiceSelector is a mock of VirtualServiceSelector interface
type MockVirtualServiceSelector struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualServiceSelectorMockRecorder
}

// MockVirtualServiceSelectorMockRecorder is the mock recorder for MockVirtualServiceSelector
type MockVirtualServiceSelectorMockRecorder struct {
	mock *MockVirtualServiceSelector
}

// NewMockVirtualServiceSelector creates a new mock instance
func NewMockVirtualServiceSelector(ctrl *gomock.Controller) *MockVirtualServiceSelector {
	mock := &MockVirtualServiceSelector{ctrl: ctrl}
	mock.recorder = &MockVirtualServiceSelectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVirtualServiceSelector) EXPECT() *MockVirtualServiceSelectorMockRecorder {
	return m.recorder
}

// SelectOrCreateVirtualService mocks base method
func (m *MockVirtualServiceSelector) SelectOrCreateVirtualService(arg0 context.Context, arg1 *core.ResourceRef) (*v1.VirtualService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectOrCreateVirtualService", arg0, arg1)
	ret0, _ := ret[0].(*v1.VirtualService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectOrCreateVirtualService indicates an expected call of SelectOrCreateVirtualService
func (mr *MockVirtualServiceSelectorMockRecorder) SelectOrCreateVirtualService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectOrCreateVirtualService", reflect.TypeOf((*MockVirtualServiceSelector)(nil).SelectOrCreateVirtualService), arg0, arg1)
}
