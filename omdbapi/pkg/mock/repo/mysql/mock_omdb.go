// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/henprasetya/omdbapi/pkg/repo/mysql (interfaces: OmdbMysql)

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockOmdbMysql is a mock of OmdbMysql interface
type MockOmdbMysql struct {
	ctrl     *gomock.Controller
	recorder *MockOmdbMysqlMockRecorder
}

// MockOmdbMysqlMockRecorder is the mock recorder for MockOmdbMysql
type MockOmdbMysqlMockRecorder struct {
	mock *MockOmdbMysql
}

// NewMockOmdbMysql creates a new mock instance
func NewMockOmdbMysql(ctrl *gomock.Controller) *MockOmdbMysql {
	mock := &MockOmdbMysql{ctrl: ctrl}
	mock.recorder = &MockOmdbMysqlMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOmdbMysql) EXPECT() *MockOmdbMysqlMockRecorder {
	return m.recorder
}

// SelectFromDb mocks base method
func (m *MockOmdbMysql) SelectFromDb() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SelectFromDb")
}

// SelectFromDb indicates an expected call of SelectFromDb
func (mr *MockOmdbMysqlMockRecorder) SelectFromDb() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectFromDb", reflect.TypeOf((*MockOmdbMysql)(nil).SelectFromDb))
}
