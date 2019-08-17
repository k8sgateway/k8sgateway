// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/gloo/projects/gloo/pkg/plugins/kubernetes (interfaces: KubePluginSharedFactory)

// Package mock_kubernetes is a generated GoMock package.
package mock_kubernetes

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/client-go/listers/core/v1"
)

// MockKubePluginSharedFactory is a mock of KubePluginSharedFactory interface
type MockKubePluginSharedFactory struct {
	ctrl     *gomock.Controller
	recorder *MockKubePluginSharedFactoryMockRecorder
}

// MockKubePluginSharedFactoryMockRecorder is the mock recorder for MockKubePluginSharedFactory
type MockKubePluginSharedFactoryMockRecorder struct {
	mock *MockKubePluginSharedFactory
}

// NewMockKubePluginSharedFactory creates a new mock instance
func NewMockKubePluginSharedFactory(ctrl *gomock.Controller) *MockKubePluginSharedFactory {
	mock := &MockKubePluginSharedFactory{ctrl: ctrl}
	mock.recorder = &MockKubePluginSharedFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKubePluginSharedFactory) EXPECT() *MockKubePluginSharedFactoryMockRecorder {
	return m.recorder
}

// EndpointsLister mocks base method
func (m *MockKubePluginSharedFactory) EndpointsLister(arg0 string) v1.EndpointsLister {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EndpointsLister", arg0)
	ret0, _ := ret[0].(v1.EndpointsLister)
	return ret0
}

// EndpointsLister indicates an expected call of EndpointsLister
func (mr *MockKubePluginSharedFactoryMockRecorder) EndpointsLister(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndpointsLister", reflect.TypeOf((*MockKubePluginSharedFactory)(nil).EndpointsLister), arg0)
}

// Subscribe mocks base method
func (m *MockKubePluginSharedFactory) Subscribe() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockKubePluginSharedFactoryMockRecorder) Subscribe() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockKubePluginSharedFactory)(nil).Subscribe))
}

// Unsubscribe mocks base method
func (m *MockKubePluginSharedFactory) Unsubscribe(arg0 <-chan struct{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Unsubscribe", arg0)
}

// Unsubscribe indicates an expected call of Unsubscribe
func (mr *MockKubePluginSharedFactoryMockRecorder) Unsubscribe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockKubePluginSharedFactory)(nil).Unsubscribe), arg0)
}
