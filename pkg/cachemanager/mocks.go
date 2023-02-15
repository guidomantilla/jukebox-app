// Code generated by MockGen. DO NOT EDIT.
// Source: jukebox-app/pkg/cachemanager (interfaces: CacheManager)

// Package cachemanager is a generated GoMock package.
package cachemanager

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCacheManager is a mock of CacheManager interface.
type MockCacheManager struct {
	ctrl     *gomock.Controller
	recorder *MockCacheManagerMockRecorder
}

// MockCacheManagerMockRecorder is the mock recorder for MockCacheManager.
type MockCacheManagerMockRecorder struct {
	mock *MockCacheManager
}

// NewMockCacheManager creates a new mock instance.
func NewMockCacheManager(ctrl *gomock.Controller) *MockCacheManager {
	mock := &MockCacheManager{ctrl: ctrl}
	mock.recorder = &MockCacheManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCacheManager) EXPECT() *MockCacheManagerMockRecorder {
	return m.recorder
}

// Clear mocks base method.
func (m *MockCacheManager) Clear(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clear", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Clear indicates an expected call of Clear.
func (mr *MockCacheManagerMockRecorder) Clear(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clear", reflect.TypeOf((*MockCacheManager)(nil).Clear), arg0)
}

// Delete mocks base method.
func (m *MockCacheManager) Delete(arg0 context.Context, arg1 string, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCacheManagerMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCacheManager)(nil).Delete), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockCacheManager) Get(arg0 context.Context, arg1 string, arg2, arg3 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockCacheManagerMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCacheManager)(nil).Get), arg0, arg1, arg2, arg3)
}

// GetType mocks base method.
func (m *MockCacheManager) GetType() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetType")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetType indicates an expected call of GetType.
func (mr *MockCacheManagerMockRecorder) GetType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetType", reflect.TypeOf((*MockCacheManager)(nil).GetType))
}

// Invalidate mocks base method.
func (m *MockCacheManager) Invalidate(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Invalidate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Invalidate indicates an expected call of Invalidate.
func (mr *MockCacheManagerMockRecorder) Invalidate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Invalidate", reflect.TypeOf((*MockCacheManager)(nil).Invalidate), arg0, arg1)
}

// Set mocks base method.
func (m *MockCacheManager) Set(arg0 context.Context, arg1 string, arg2, arg3 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockCacheManagerMockRecorder) Set(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockCacheManager)(nil).Set), arg0, arg1, arg2, arg3)
}