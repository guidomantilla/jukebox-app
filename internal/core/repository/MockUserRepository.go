// Code generated by MockGen. DO NOT EDIT.
// Source: jukebox-app/internal/core/repository (interfaces: UserRepository)

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	model "jukebox-app/internal/core/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserRepository) Create(arg0 context.Context, arg1 *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRepository)(nil).Create), arg0, arg1)
}

// DeleteById mocks base method.
func (m *MockUserRepository) DeleteById(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteById", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteById indicates an expected call of DeleteById.
func (mr *MockUserRepositoryMockRecorder) DeleteById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteById", reflect.TypeOf((*MockUserRepository)(nil).DeleteById), arg0, arg1)
}

// FindAll mocks base method.
func (m *MockUserRepository) FindAll(arg0 context.Context) (*[]model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", arg0)
	ret0, _ := ret[0].(*[]model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockUserRepositoryMockRecorder) FindAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockUserRepository)(nil).FindAll), arg0)
}

// FindByCode mocks base method.
func (m *MockUserRepository) FindByCode(arg0 context.Context, arg1 int64) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByCode", arg0, arg1)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByCode indicates an expected call of FindByCode.
func (mr *MockUserRepositoryMockRecorder) FindByCode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByCode", reflect.TypeOf((*MockUserRepository)(nil).FindByCode), arg0, arg1)
}

// FindByEmail mocks base method.
func (m *MockUserRepository) FindByEmail(arg0 context.Context, arg1 string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", arg0, arg1)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail.
func (mr *MockUserRepositoryMockRecorder) FindByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindByEmail), arg0, arg1)
}

// FindById mocks base method.
func (m *MockUserRepository) FindById(arg0 context.Context, arg1 int64) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", arg0, arg1)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockUserRepositoryMockRecorder) FindById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockUserRepository)(nil).FindById), arg0, arg1)
}

// FindByName mocks base method.
func (m *MockUserRepository) FindByName(arg0 context.Context, arg1 string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", arg0, arg1)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByName indicates an expected call of FindByName.
func (mr *MockUserRepositoryMockRecorder) FindByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockUserRepository)(nil).FindByName), arg0, arg1)
}

// Update mocks base method.
func (m *MockUserRepository) Update(arg0 context.Context, arg1 *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserRepositoryMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserRepository)(nil).Update), arg0, arg1)
}
