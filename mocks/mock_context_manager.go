// Code generated by MockGen. DO NOT EDIT.
// Source: fiap-fast-food-ms-producao/adapter/context_manager (interfaces: ContextManager)

// Package mocks is a generated GoMock package.
package mocks

import (
	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockContextManager is a mock of ContextManager interface
type MockContextManager struct {
	ctrl     *gomock.Controller
	recorder *MockContextManagerMockRecorder
}

// MockContextManagerMockRecorder is the mock recorder for MockContextManager
type MockContextManagerMockRecorder struct {
	mock *MockContextManager
}

// NewMockContextManager creates a new mock instance
func NewMockContextManager(ctrl *gomock.Controller) *MockContextManager {
	mock := &MockContextManager{ctrl: ctrl}
	mock.recorder = &MockContextManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContextManager) EXPECT() *MockContextManagerMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockContextManager) Get(arg0 string) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockContextManagerMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockContextManager)(nil).Get), arg0)
}

// PassContext mocks base method
func (m *MockContextManager) PassContext(arg0 *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PassContext", arg0)
}

// PassContext indicates an expected call of PassContext
func (mr *MockContextManagerMockRecorder) PassContext(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PassContext", reflect.TypeOf((*MockContextManager)(nil).PassContext), arg0)
}

// Set mocks base method
func (m *MockContextManager) Set(arg0 string, arg1 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", arg0, arg1)
}

// Set indicates an expected call of Set
func (mr *MockContextManagerMockRecorder) Set(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockContextManager)(nil).Set), arg0, arg1)
}
