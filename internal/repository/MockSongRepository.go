// Code generated by MockGen. DO NOT EDIT.
// Source: jukebox-app/internal/core/repository (interfaces: SongRepository)

// Package repository is a generated GoMock package.
package repository

import (
	"context"
	"reflect"

	"github.com/golang/mock/gomock"

	"jukebox-app/internal/model"
)

// MockSongRepository is a mock of SongRepository interface.
type MockSongRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSongRepositoryMockRecorder
}

// MockSongRepositoryMockRecorder is the mock recorder for MockSongRepository.
type MockSongRepositoryMockRecorder struct {
	mock *MockSongRepository
}

// NewMockSongRepository creates a new mock instance.
func NewMockSongRepository(ctrl *gomock.Controller) *MockSongRepository {
	mock := &MockSongRepository{ctrl: ctrl}
	mock.recorder = &MockSongRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSongRepository) EXPECT() *MockSongRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockSongRepository) Create(arg0 context.Context, arg1 *model.Song) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockSongRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSongRepository)(nil).Create), arg0, arg1)
}

// DeleteById mocks base method.
func (m *MockSongRepository) DeleteById(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteById", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteById indicates an expected call of DeleteById.
func (mr *MockSongRepositoryMockRecorder) DeleteById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteById", reflect.TypeOf((*MockSongRepository)(nil).DeleteById), arg0, arg1)
}

// FindAll mocks base method.
func (m *MockSongRepository) FindAll(arg0 context.Context) (*[]model.Song, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", arg0)
	ret0, _ := ret[0].(*[]model.Song)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockSongRepositoryMockRecorder) FindAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockSongRepository)(nil).FindAll), arg0)
}

// FindByArtistId mocks base method.
func (m *MockSongRepository) FindByArtistId(arg0 context.Context, arg1 int64) (*[]model.Song, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByArtistId", arg0, arg1)
	ret0, _ := ret[0].(*[]model.Song)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByArtistId indicates an expected call of FindByArtistId.
func (mr *MockSongRepositoryMockRecorder) FindByArtistId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByArtistId", reflect.TypeOf((*MockSongRepository)(nil).FindByArtistId), arg0, arg1)
}

// FindByCode mocks base method.
func (m *MockSongRepository) FindByCode(arg0 context.Context, arg1 int64) (*model.Song, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByCode", arg0, arg1)
	ret0, _ := ret[0].(*model.Song)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByCode indicates an expected call of FindByCode.
func (mr *MockSongRepositoryMockRecorder) FindByCode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByCode", reflect.TypeOf((*MockSongRepository)(nil).FindByCode), arg0, arg1)
}

// FindById mocks base method.
func (m *MockSongRepository) FindById(arg0 context.Context, arg1 int64) (*model.Song, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", arg0, arg1)
	ret0, _ := ret[0].(*model.Song)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockSongRepositoryMockRecorder) FindById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockSongRepository)(nil).FindById), arg0, arg1)
}

// FindByName mocks base method.
func (m *MockSongRepository) FindByName(arg0 context.Context, arg1 string) (*model.Song, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", arg0, arg1)
	ret0, _ := ret[0].(*model.Song)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByName indicates an expected call of FindByName.
func (mr *MockSongRepositoryMockRecorder) FindByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockSongRepository)(nil).FindByName), arg0, arg1)
}

// Update mocks base method.
func (m *MockSongRepository) Update(arg0 context.Context, arg1 *model.Song) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockSongRepositoryMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSongRepository)(nil).Update), arg0, arg1)
}
