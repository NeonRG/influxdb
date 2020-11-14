// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/influxdata/influxdb/v2/storage (interfaces: EngineSchema)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	influxdb "github.com/influxdata/influxdb/v2"
)

// MockEngineSchema is a mock of EngineSchema interface
type MockEngineSchema struct {
	ctrl     *gomock.Controller
	recorder *MockEngineSchemaMockRecorder
}

// MockEngineSchemaMockRecorder is the mock recorder for MockEngineSchema
type MockEngineSchemaMockRecorder struct {
	mock *MockEngineSchema
}

// NewMockEngineSchema creates a new mock instance
func NewMockEngineSchema(ctrl *gomock.Controller) *MockEngineSchema {
	mock := &MockEngineSchema{ctrl: ctrl}
	mock.recorder = &MockEngineSchemaMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEngineSchema) EXPECT() *MockEngineSchemaMockRecorder {
	return m.recorder
}

// CreateBucket mocks base method
func (m *MockEngineSchema) CreateBucket(arg0 context.Context, arg1 *influxdb.Bucket) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBucket", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBucket indicates an expected call of CreateBucket
func (mr *MockEngineSchemaMockRecorder) CreateBucket(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBucket", reflect.TypeOf((*MockEngineSchema)(nil).CreateBucket), arg0, arg1)
}

// DeleteBucket mocks base method
func (m *MockEngineSchema) DeleteBucket(arg0 context.Context, arg1, arg2 influxdb.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBucket", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBucket indicates an expected call of DeleteBucket
func (mr *MockEngineSchemaMockRecorder) DeleteBucket(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBucket", reflect.TypeOf((*MockEngineSchema)(nil).DeleteBucket), arg0, arg1, arg2)
}

// UpdateBucketRetentionPeriod mocks base method
func (m *MockEngineSchema) UpdateBucketRetentionPeriod(arg0 context.Context, arg1 influxdb.ID, arg2 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBucketRetentionPeriod", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBucketRetentionPeriod indicates an expected call of UpdateBucketRetentionPeriod
func (mr *MockEngineSchemaMockRecorder) UpdateBucketRetentionPeriod(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBucketRetentionPeriod", reflect.TypeOf((*MockEngineSchema)(nil).UpdateBucketRetentionPeriod), arg0, arg1, arg2)
}
