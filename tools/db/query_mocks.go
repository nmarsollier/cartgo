// Code generated by MockGen. DO NOT EDIT.
// Source: ./tools/db/query.go

// Package db is a generated GoMock package.
package db

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockQuery is a mock of Query interface.
type MockQuery[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockQueryMockRecorder[T]
}

// MockQueryMockRecorder is the mock recorder for MockQuery.
type MockQueryMockRecorder[T any] struct {
	mock *MockQuery[T]
}

// NewMockQuery creates a new mock instance.
func NewMockQuery[T any](ctrl *gomock.Controller) *MockQuery[T] {
	mock := &MockQuery[T]{ctrl: ctrl}
	mock.recorder = &MockQueryMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuery[T]) EXPECT() *MockQueryMockRecorder[T] {
	return m.recorder
}

// Row mocks base method.
func (m *MockQuery[T]) Row(query string, args ...interface{}) (*T, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Row", varargs...)
	ret0, _ := ret[0].(*T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Row indicates an expected call of Row.
func (mr *MockQueryMockRecorder[T]) Row(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Row", reflect.TypeOf((*MockQuery[T])(nil).Row), varargs...)
}