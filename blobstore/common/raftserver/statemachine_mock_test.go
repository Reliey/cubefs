// Code generated by MockGen. DO NOT EDIT.
// Source: statemachine.go

// Package raftserver is a generated GoMock package.
package raftserver

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSnapshot is a mock of Snapshot interface.
type MockSnapshot struct {
	ctrl     *gomock.Controller
	recorder *MockSnapshotMockRecorder
}

// MockSnapshotMockRecorder is the mock recorder for MockSnapshot.
type MockSnapshotMockRecorder struct {
	mock *MockSnapshot
}

// NewMockSnapshot creates a new mock instance.
func NewMockSnapshot(ctrl *gomock.Controller) *MockSnapshot {
	mock := &MockSnapshot{ctrl: ctrl}
	mock.recorder = &MockSnapshotMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSnapshot) EXPECT() *MockSnapshotMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockSnapshot) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockSnapshotMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSnapshot)(nil).Close))
}

// Index mocks base method.
func (m *MockSnapshot) Index() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Index")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Index indicates an expected call of Index.
func (mr *MockSnapshotMockRecorder) Index() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Index", reflect.TypeOf((*MockSnapshot)(nil).Index))
}

// Name mocks base method.
func (m *MockSnapshot) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockSnapshotMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockSnapshot)(nil).Name))
}

// Read mocks base method.
func (m *MockSnapshot) Read() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockSnapshotMockRecorder) Read() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockSnapshot)(nil).Read))
}

// MockStateMachine is a mock of StateMachine interface.
type MockStateMachine struct {
	ctrl     *gomock.Controller
	recorder *MockStateMachineMockRecorder
}

// MockStateMachineMockRecorder is the mock recorder for MockStateMachine.
type MockStateMachineMockRecorder struct {
	mock *MockStateMachine
}

// NewMockStateMachine creates a new mock instance.
func NewMockStateMachine(ctrl *gomock.Controller) *MockStateMachine {
	mock := &MockStateMachine{ctrl: ctrl}
	mock.recorder = &MockStateMachineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStateMachine) EXPECT() *MockStateMachineMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockStateMachine) Apply(data [][]byte, index uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", data, index)
	ret0, _ := ret[0].(error)
	return ret0
}

// Apply indicates an expected call of Apply.
func (mr *MockStateMachineMockRecorder) Apply(data, index interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockStateMachine)(nil).Apply), data, index)
}

// ApplyMemberChange mocks base method.
func (m *MockStateMachine) ApplyMemberChange(cc ConfChange, index uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyMemberChange", cc, index)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyMemberChange indicates an expected call of ApplyMemberChange.
func (mr *MockStateMachineMockRecorder) ApplyMemberChange(cc, index interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyMemberChange", reflect.TypeOf((*MockStateMachine)(nil).ApplyMemberChange), cc, index)
}

// ApplySnapshot mocks base method.
func (m *MockStateMachine) ApplySnapshot(meta SnapshotMeta, st Snapshot) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplySnapshot", meta, st)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplySnapshot indicates an expected call of ApplySnapshot.
func (mr *MockStateMachineMockRecorder) ApplySnapshot(meta, st interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplySnapshot", reflect.TypeOf((*MockStateMachine)(nil).ApplySnapshot), meta, st)
}

// LeaderChange mocks base method.
func (m *MockStateMachine) LeaderChange(leader uint64, host string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LeaderChange", leader, host)
}

// LeaderChange indicates an expected call of LeaderChange.
func (mr *MockStateMachineMockRecorder) LeaderChange(leader, host interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LeaderChange", reflect.TypeOf((*MockStateMachine)(nil).LeaderChange), leader, host)
}

// Snapshot mocks base method.
func (m *MockStateMachine) Snapshot() (Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshot")
	ret0, _ := ret[0].(Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Snapshot indicates an expected call of Snapshot.
func (mr *MockStateMachineMockRecorder) Snapshot() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshot", reflect.TypeOf((*MockStateMachine)(nil).Snapshot))
}
