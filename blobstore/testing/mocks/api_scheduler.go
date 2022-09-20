// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cubefs/cubefs/blobstore/api/scheduler (interfaces: IScheduler)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	scheduler "github.com/cubefs/cubefs/blobstore/api/scheduler"
	proto "github.com/cubefs/cubefs/blobstore/common/proto"
	gomock "github.com/golang/mock/gomock"
)

// MockIScheduler is a mock of IScheduler interface.
type MockIScheduler struct {
	ctrl     *gomock.Controller
	recorder *MockISchedulerMockRecorder
}

// MockISchedulerMockRecorder is the mock recorder for MockIScheduler.
type MockISchedulerMockRecorder struct {
	mock *MockIScheduler
}

// NewMockIScheduler creates a new mock instance.
func NewMockIScheduler(ctrl *gomock.Controller) *MockIScheduler {
	mock := &MockIScheduler{ctrl: ctrl}
	mock.recorder = &MockISchedulerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIScheduler) EXPECT() *MockISchedulerMockRecorder {
	return m.recorder
}

// AcquireInspectTask mocks base method.
func (m *MockIScheduler) AcquireInspectTask(arg0 context.Context) (*proto.VolumeInspectTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcquireInspectTask", arg0)
	ret0, _ := ret[0].(*proto.VolumeInspectTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcquireInspectTask indicates an expected call of AcquireInspectTask.
func (mr *MockISchedulerMockRecorder) AcquireInspectTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcquireInspectTask", reflect.TypeOf((*MockIScheduler)(nil).AcquireInspectTask), arg0)
}

// AcquireTask mocks base method.
func (m *MockIScheduler) AcquireTask(arg0 context.Context, arg1 *scheduler.AcquireArgs) (*proto.MigrateTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcquireTask", arg0, arg1)
	ret0, _ := ret[0].(*proto.MigrateTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcquireTask indicates an expected call of AcquireTask.
func (mr *MockISchedulerMockRecorder) AcquireTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcquireTask", reflect.TypeOf((*MockIScheduler)(nil).AcquireTask), arg0, arg1)
}

// AddManualMigrateTask mocks base method.
func (m *MockIScheduler) AddManualMigrateTask(arg0 context.Context, arg1 *scheduler.AddManualMigrateArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddManualMigrateTask", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddManualMigrateTask indicates an expected call of AddManualMigrateTask.
func (mr *MockISchedulerMockRecorder) AddManualMigrateTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddManualMigrateTask", reflect.TypeOf((*MockIScheduler)(nil).AddManualMigrateTask), arg0, arg1)
}

// CancelTask mocks base method.
func (m *MockIScheduler) CancelTask(arg0 context.Context, arg1 *scheduler.OperateTaskArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelTask", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelTask indicates an expected call of CancelTask.
func (mr *MockISchedulerMockRecorder) CancelTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelTask", reflect.TypeOf((*MockIScheduler)(nil).CancelTask), arg0, arg1)
}

// CompleteInspectTask mocks base method.
func (m *MockIScheduler) CompleteInspectTask(arg0 context.Context, arg1 *proto.VolumeInspectRet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompleteInspectTask", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CompleteInspectTask indicates an expected call of CompleteInspectTask.
func (mr *MockISchedulerMockRecorder) CompleteInspectTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteInspectTask", reflect.TypeOf((*MockIScheduler)(nil).CompleteInspectTask), arg0, arg1)
}

// CompleteTask mocks base method.
func (m *MockIScheduler) CompleteTask(arg0 context.Context, arg1 *scheduler.OperateTaskArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompleteTask", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CompleteTask indicates an expected call of CompleteTask.
func (mr *MockISchedulerMockRecorder) CompleteTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteTask", reflect.TypeOf((*MockIScheduler)(nil).CompleteTask), arg0, arg1)
}

// DetailMigrateTask mocks base method.
func (m *MockIScheduler) DetailMigrateTask(arg0 context.Context, arg1 *scheduler.MigrateTaskDetailArgs) (scheduler.MigrateTaskDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetailMigrateTask", arg0, arg1)
	ret0, _ := ret[0].(scheduler.MigrateTaskDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetailMigrateTask indicates an expected call of DetailMigrateTask.
func (mr *MockISchedulerMockRecorder) DetailMigrateTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetailMigrateTask", reflect.TypeOf((*MockIScheduler)(nil).DetailMigrateTask), arg0, arg1)
}

// DiskMigratingStats mocks base method.
func (m *MockIScheduler) DiskMigratingStats(arg0 context.Context, arg1 *scheduler.DiskMigratingStatsArgs) (*scheduler.DiskMigratingStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiskMigratingStats", arg0, arg1)
	ret0, _ := ret[0].(*scheduler.DiskMigratingStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DiskMigratingStats indicates an expected call of DiskMigratingStats.
func (mr *MockISchedulerMockRecorder) DiskMigratingStats(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiskMigratingStats", reflect.TypeOf((*MockIScheduler)(nil).DiskMigratingStats), arg0, arg1)
}

// LeaderStats mocks base method.
func (m *MockIScheduler) LeaderStats(arg0 context.Context) (scheduler.TasksStat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LeaderStats", arg0)
	ret0, _ := ret[0].(scheduler.TasksStat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LeaderStats indicates an expected call of LeaderStats.
func (mr *MockISchedulerMockRecorder) LeaderStats(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LeaderStats", reflect.TypeOf((*MockIScheduler)(nil).LeaderStats), arg0)
}

// ReclaimTask mocks base method.
func (m *MockIScheduler) ReclaimTask(arg0 context.Context, arg1 *scheduler.OperateTaskArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReclaimTask", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReclaimTask indicates an expected call of ReclaimTask.
func (mr *MockISchedulerMockRecorder) ReclaimTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReclaimTask", reflect.TypeOf((*MockIScheduler)(nil).ReclaimTask), arg0, arg1)
}

// RenewalTask mocks base method.
func (m *MockIScheduler) RenewalTask(arg0 context.Context, arg1 *scheduler.TaskRenewalArgs) (*scheduler.TaskRenewalRet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenewalTask", arg0, arg1)
	ret0, _ := ret[0].(*scheduler.TaskRenewalRet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RenewalTask indicates an expected call of RenewalTask.
func (mr *MockISchedulerMockRecorder) RenewalTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenewalTask", reflect.TypeOf((*MockIScheduler)(nil).RenewalTask), arg0, arg1)
}

// ReportTask mocks base method.
func (m *MockIScheduler) ReportTask(arg0 context.Context, arg1 *scheduler.TaskReportArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReportTask", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReportTask indicates an expected call of ReportTask.
func (mr *MockISchedulerMockRecorder) ReportTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportTask", reflect.TypeOf((*MockIScheduler)(nil).ReportTask), arg0, arg1)
}

// Stats mocks base method.
func (m *MockIScheduler) Stats(arg0 context.Context, arg1 string) (scheduler.TasksStat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stats", arg0, arg1)
	ret0, _ := ret[0].(scheduler.TasksStat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stats indicates an expected call of Stats.
func (mr *MockISchedulerMockRecorder) Stats(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stats", reflect.TypeOf((*MockIScheduler)(nil).Stats), arg0, arg1)
}

// UpdateVolume mocks base method.
func (m *MockIScheduler) UpdateVolume(arg0 context.Context, arg1 string, arg2 proto.Vid) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVolume", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateVolume indicates an expected call of UpdateVolume.
func (mr *MockISchedulerMockRecorder) UpdateVolume(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVolume", reflect.TypeOf((*MockIScheduler)(nil).UpdateVolume), arg0, arg1, arg2)
}
