// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces/cassandra.go

package mocks

import (
	gocql "github.com/gocql/gocql"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDB is a mock of DB interface
type MockDB struct {
	ctrl     *gomock.Controller
	recorder *MockDBMockRecorder
}

// MockDBMockRecorder is the mock recorder for MockDB
type MockDBMockRecorder struct {
	mock *MockDB
}

// NewMockDB creates a new mock instance
func NewMockDB(ctrl *gomock.Controller) *MockDB {
	mock := &MockDB{ctrl: ctrl}
	mock.recorder = &MockDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockDB) EXPECT() *MockDBMockRecorder {
	return _m.recorder
}

// CreateSession mocks base method
func (_m *MockDB) CreateSession() (*gocql.Session, error) {
	ret := _m.ctrl.Call(_m, "CreateSession")
	ret0, _ := ret[0].(*gocql.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession
func (_mr *MockDBMockRecorder) CreateSession() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CreateSession", reflect.TypeOf((*MockDB)(nil).CreateSession))
}

// MockSession is a mock of Session interface
type MockSession struct {
	ctrl     *gomock.Controller
	recorder *MockSessionMockRecorder
}

// MockSessionMockRecorder is the mock recorder for MockSession
type MockSessionMockRecorder struct {
	mock *MockSession
}

// NewMockSession creates a new mock instance
func NewMockSession(ctrl *gomock.Controller) *MockSession {
	mock := &MockSession{ctrl: ctrl}
	mock.recorder = &MockSessionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockSession) EXPECT() *MockSessionMockRecorder {
	return _m.recorder
}

// Close mocks base method
func (_m *MockSession) Close() {
	_m.ctrl.Call(_m, "Close")
}

// Close indicates an expected call of Close
func (_mr *MockSessionMockRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Close", reflect.TypeOf((*MockSession)(nil).Close))
}

// Query mocks base method
func (_m *MockSession) Query(_param0 string, _param1 ...interface{}) *gocql.Query {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Query", _s...)
	ret0, _ := ret[0].(*gocql.Query)
	return ret0
}

// Query indicates an expected call of Query
func (_mr *MockSessionMockRecorder) Query(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Query", reflect.TypeOf((*MockSession)(nil).Query), _s...)
}

// ExecuteBatch mocks base method
func (_m *MockSession) ExecuteBatch(_param0 *gocql.Batch) error {
	ret := _m.ctrl.Call(_m, "ExecuteBatch", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecuteBatch indicates an expected call of ExecuteBatch
func (_mr *MockSessionMockRecorder) ExecuteBatch(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ExecuteBatch", reflect.TypeOf((*MockSession)(nil).ExecuteBatch), arg0)
}

// NewBatch mocks base method
func (_m *MockSession) NewBatch(_param0 gocql.BatchType) *gocql.Batch {
	ret := _m.ctrl.Call(_m, "NewBatch", _param0)
	ret0, _ := ret[0].(*gocql.Batch)
	return ret0
}

// NewBatch indicates an expected call of NewBatch
func (_mr *MockSessionMockRecorder) NewBatch(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "NewBatch", reflect.TypeOf((*MockSession)(nil).NewBatch), arg0)
}
