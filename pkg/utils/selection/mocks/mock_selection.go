// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/gloo/pkg/utils/selection (interfaces: VirtualServiceSelector,NamespaceLister)

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

// SelectOrCreate mocks base method
func (m *MockVirtualServiceSelector) SelectOrCreate(arg0 context.Context, arg1 *core.ResourceRef) (*v1.VirtualService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectOrCreate", arg0, arg1)
	ret0, _ := ret[0].(*v1.VirtualService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectOrCreate indicates an expected call of SelectOrCreate
func (mr *MockVirtualServiceSelectorMockRecorder) SelectOrCreate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectOrCreate", reflect.TypeOf((*MockVirtualServiceSelector)(nil).SelectOrCreate), arg0, arg1)
}

// MockNamespaceLister is a mock of NamespaceLister interface
type MockNamespaceLister struct {
	ctrl     *gomock.Controller
	recorder *MockNamespaceListerMockRecorder
}

// MockNamespaceListerMockRecorder is the mock recorder for MockNamespaceLister
type MockNamespaceListerMockRecorder struct {
	mock *MockNamespaceLister
}

// NewMockNamespaceLister creates a new mock instance
func NewMockNamespaceLister(ctrl *gomock.Controller) *MockNamespaceLister {
	mock := &MockNamespaceLister{ctrl: ctrl}
	mock.recorder = &MockNamespaceListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNamespaceLister) EXPECT() *MockNamespaceListerMockRecorder {
	return m.recorder
}

// List mocks base method
func (m *MockNamespaceLister) List() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockNamespaceListerMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockNamespaceLister)(nil).List))
}
