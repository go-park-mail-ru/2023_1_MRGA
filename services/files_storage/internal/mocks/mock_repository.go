// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/go-park-mail-ru/2023_1_MRGA.git/services/files_storage/internal (interfaces: IRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIRepository is a mock of IRepository interface.
type MockIRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIRepositoryMockRecorder
}

// MockIRepositoryMockRecorder is the mock recorder for MockIRepository.
type MockIRepositoryMockRecorder struct {
	mock *MockIRepository
}

// NewMockIRepository creates a new mock instance.
func NewMockIRepository(ctrl *gomock.Controller) *MockIRepository {
	mock := &MockIRepository{ctrl: ctrl}
	mock.recorder = &MockIRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRepository) EXPECT() *MockIRepositoryMockRecorder {
	return m.recorder
}

// GetFile mocks base method.
func (m *MockIRepository) GetFile(arg0 uint) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFile", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFile indicates an expected call of GetFile.
func (mr *MockIRepositoryMockRecorder) GetFile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFile", reflect.TypeOf((*MockIRepository)(nil).GetFile), arg0)
}

// UploadFile mocks base method.
func (m *MockIRepository) UploadFile(arg0 string) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFile", arg0)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadFile indicates an expected call of UploadFile.
func (mr *MockIRepositoryMockRecorder) UploadFile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFile", reflect.TypeOf((*MockIRepository)(nil).UploadFile), arg0)
}
