// Code generated by MockGen. DO NOT EDIT.
// Source: database/repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	database "github.com/SushanthGunjal/BooleanAsAService/database"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRepositoryController is a mock of RepositoryController interface
type MockRepositoryController struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryControllerMockRecorder
}

// MockRepositoryControllerMockRecorder is the mock recorder for MockRepositoryController
type MockRepositoryControllerMockRecorder struct {
	mock *MockRepositoryController
}

// NewMockRepositoryController creates a new mock instance
func NewMockRepositoryController(ctrl *gomock.Controller) *MockRepositoryController {
	mock := &MockRepositoryController{ctrl: ctrl}
	mock.recorder = &MockRepositoryControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepositoryController) EXPECT() *MockRepositoryControllerMockRecorder {
	return m.recorder
}

// DatabaseCreate mocks base method
func (m *MockRepositoryController) DatabaseCreate(service *database.Services) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DatabaseCreate", service)
	ret0, _ := ret[0].(error)
	return ret0
}

// DatabaseCreate indicates an expected call of DatabaseCreate
func (mr *MockRepositoryControllerMockRecorder) DatabaseCreate(service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatabaseCreate", reflect.TypeOf((*MockRepositoryController)(nil).DatabaseCreate), service)
}

// DatabaseGet mocks base method
func (m *MockRepositoryController) DatabaseGet(service *database.Services) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DatabaseGet", service)
	ret0, _ := ret[0].(error)
	return ret0
}

// DatabaseGet indicates an expected call of DatabaseGet
func (mr *MockRepositoryControllerMockRecorder) DatabaseGet(service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatabaseGet", reflect.TypeOf((*MockRepositoryController)(nil).DatabaseGet), service)
}

// DatabaseSave mocks base method
func (m *MockRepositoryController) DatabaseSave(service *database.Services) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DatabaseSave", service)
	ret0, _ := ret[0].(error)
	return ret0
}

// DatabaseSave indicates an expected call of DatabaseSave
func (mr *MockRepositoryControllerMockRecorder) DatabaseSave(service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatabaseSave", reflect.TypeOf((*MockRepositoryController)(nil).DatabaseSave), service)
}

// DatabaseDelete mocks base method
func (m *MockRepositoryController) DatabaseDelete(service *database.Services) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DatabaseDelete", service)
	ret0, _ := ret[0].(error)
	return ret0
}

// DatabaseDelete indicates an expected call of DatabaseDelete
func (mr *MockRepositoryControllerMockRecorder) DatabaseDelete(service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatabaseDelete", reflect.TypeOf((*MockRepositoryController)(nil).DatabaseDelete), service)
}
