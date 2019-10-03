// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/envoyproxy/go-control-plane/envoy/service/metrics/v2 (interfaces: MetricsService_StreamMetricsServer)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v2 "github.com/envoyproxy/go-control-plane/envoy/service/metrics/v2"
	gomock "github.com/golang/mock/gomock"
	metadata "google.golang.org/grpc/metadata"
)

// MockMetricsService_StreamMetricsServer is a mock of MetricsService_StreamMetricsServer interface
type MockMetricsService_StreamMetricsServer struct {
	ctrl     *gomock.Controller
	recorder *MockMetricsService_StreamMetricsServerMockRecorder
}

// MockMetricsService_StreamMetricsServerMockRecorder is the mock recorder for MockMetricsService_StreamMetricsServer
type MockMetricsService_StreamMetricsServerMockRecorder struct {
	mock *MockMetricsService_StreamMetricsServer
}

// NewMockMetricsService_StreamMetricsServer creates a new mock instance
func NewMockMetricsService_StreamMetricsServer(ctrl *gomock.Controller) *MockMetricsService_StreamMetricsServer {
	mock := &MockMetricsService_StreamMetricsServer{ctrl: ctrl}
	mock.recorder = &MockMetricsService_StreamMetricsServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMetricsService_StreamMetricsServer) EXPECT() *MockMetricsService_StreamMetricsServerMockRecorder {
	return m.recorder
}

// Context mocks base method
func (m *MockMetricsService_StreamMetricsServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockMetricsService_StreamMetricsServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockMetricsService_StreamMetricsServer)(nil).Context))
}

// Recv mocks base method
func (m *MockMetricsService_StreamMetricsServer) Recv() (*v2.StreamMetricsMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*v2.StreamMetricsMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockMetricsService_StreamMetricsServerMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockMetricsService_StreamMetricsServer)(nil).Recv))
}

// RecvMsg mocks base method
func (m *MockMetricsService_StreamMetricsServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockMetricsService_StreamMetricsServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockMetricsService_StreamMetricsServer)(nil).RecvMsg), arg0)
}

// SendAndClose mocks base method
func (m *MockMetricsService_StreamMetricsServer) SendAndClose(arg0 *v2.StreamMetricsResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendAndClose", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendAndClose indicates an expected call of SendAndClose
func (mr *MockMetricsService_StreamMetricsServerMockRecorder) SendAndClose(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAndClose", reflect.TypeOf((*MockMetricsService_StreamMetricsServer)(nil).SendAndClose), arg0)
}

// SendHeader mocks base method
func (m *MockMetricsService_StreamMetricsServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader
func (mr *MockMetricsService_StreamMetricsServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockMetricsService_StreamMetricsServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method
func (m *MockMetricsService_StreamMetricsServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockMetricsService_StreamMetricsServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockMetricsService_StreamMetricsServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method
func (m *MockMetricsService_StreamMetricsServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader
func (mr *MockMetricsService_StreamMetricsServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockMetricsService_StreamMetricsServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method
func (m *MockMetricsService_StreamMetricsServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer
func (mr *MockMetricsService_StreamMetricsServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockMetricsService_StreamMetricsServer)(nil).SetTrailer), arg0)
}
