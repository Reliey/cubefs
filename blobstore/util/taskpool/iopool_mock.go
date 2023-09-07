// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cubefs/cubefs/blobstore/util/taskpool (interfaces: IoPool)

// Package taskpool is a generated GoMock package.
package taskpool

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIoPool is a mock of IoPool interface.
type MockIoPool struct {
	ctrl     *gomock.Controller
	recorder *MockIoPoolMockRecorder
}

// MockIoPoolMockRecorder is the mock recorder for MockIoPool.
type MockIoPoolMockRecorder struct {
	mock *MockIoPool
}

// NewMockIoPool creates a new mock instance.
func NewMockIoPool(ctrl *gomock.Controller) *MockIoPool {
	mock := &MockIoPool{ctrl: ctrl}
	mock.recorder = &MockIoPoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIoPool) EXPECT() *MockIoPoolMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockIoPool) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockIoPoolMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIoPool)(nil).Close))
}

// Submit mocks base method.
func (m *MockIoPool) Submit(arg0 uint64, arg1 func()) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Submit", arg0, arg1)
}

// Submit indicates an expected call of Submit.
func (mr *MockIoPoolMockRecorder) Submit(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Submit", reflect.TypeOf((*MockIoPool)(nil).Submit), arg0, arg1)
}
