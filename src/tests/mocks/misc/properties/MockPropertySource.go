// Code generated by MockGen. DO NOT EDIT.
// Source: jukebox-app/src/misc/properties (interfaces: PropertySource)

// Package properties is a generated GoMock package.
package properties

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPropertySource is a mock of PropertySource interface.
type MockPropertySource struct {
	ctrl     *gomock.Controller
	recorder *MockPropertySourceMockRecorder
}

// MockPropertySourceMockRecorder is the mock recorder for MockPropertySource.
type MockPropertySourceMockRecorder struct {
	mock *MockPropertySource
}

// NewMockPropertySource creates a new mock instance.
func NewMockPropertySource(ctrl *gomock.Controller) *MockPropertySource {
	mock := &MockPropertySource{ctrl: ctrl}
	mock.recorder = &MockPropertySourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPropertySource) EXPECT() *MockPropertySourceMockRecorder {
	return m.recorder
}

// AsMap mocks base method.
func (m *MockPropertySource) AsMap() map[string]interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AsMap")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// AsMap indicates an expected call of AsMap.
func (mr *MockPropertySourceMockRecorder) AsMap() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AsMap", reflect.TypeOf((*MockPropertySource)(nil).AsMap))
}

// Get mocks base method.
func (m *MockPropertySource) Get(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockPropertySourceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPropertySource)(nil).Get), arg0)
}
