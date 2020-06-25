// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/gloo/projects/gloo/pkg/plugins/consul (interfaces: DnsResolver)

// Package mock_consul is a generated GoMock package.
package mock_consul

import (
	context "context"
	net "net"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDnsResolver is a mock of DnsResolver interface
type MockDnsResolver struct {
	ctrl     *gomock.Controller
	recorder *MockDnsResolverMockRecorder
}

// MockDnsResolverMockRecorder is the mock recorder for MockDnsResolver
type MockDnsResolverMockRecorder struct {
	mock *MockDnsResolver
}

// NewMockDnsResolver creates a new mock instance
func NewMockDnsResolver(ctrl *gomock.Controller) *MockDnsResolver {
	mock := &MockDnsResolver{ctrl: ctrl}
	mock.recorder = &MockDnsResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDnsResolver) EXPECT() *MockDnsResolverMockRecorder {
	return m.recorder
}

// Resolve mocks base method
func (m *MockDnsResolver) Resolve(arg0 context.Context, arg1 string) ([]net.IPAddr, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resolve", arg0, arg1)
	ret0, _ := ret[0].([]net.IPAddr)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Resolve indicates an expected call of Resolve
func (mr *MockDnsResolverMockRecorder) Resolve(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resolve", reflect.TypeOf((*MockDnsResolver)(nil).Resolve), arg0, arg1)
}
