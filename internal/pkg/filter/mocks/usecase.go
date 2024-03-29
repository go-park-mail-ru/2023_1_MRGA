// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	filter "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/filter"
	gomock "github.com/golang/mock/gomock"
)

// MockUseCase is a mock of UseCase interface.
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase.
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance.
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// AddFilters mocks base method.
func (m *MockUseCase) AddFilters(userId uint, FilterInp filter.FilterInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFilters", userId, FilterInp)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFilters indicates an expected call of AddFilters.
func (mr *MockUseCaseMockRecorder) AddFilters(userId, FilterInp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFilters", reflect.TypeOf((*MockUseCase)(nil).AddFilters), userId, FilterInp)
}

// ChangeFilters mocks base method.
func (m *MockUseCase) ChangeFilters(userId uint, filterInp filter.FilterInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeFilters", userId, filterInp)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeFilters indicates an expected call of ChangeFilters.
func (mr *MockUseCaseMockRecorder) ChangeFilters(userId, filterInp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeFilters", reflect.TypeOf((*MockUseCase)(nil).ChangeFilters), userId, filterInp)
}

// GetFilters mocks base method.
func (m *MockUseCase) GetFilters(userId uint) (filter.FilterInput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilters", userId)
	ret0, _ := ret[0].(filter.FilterInput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilters indicates an expected call of GetFilters.
func (mr *MockUseCaseMockRecorder) GetFilters(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilters", reflect.TypeOf((*MockUseCase)(nil).GetFilters), userId)
}

// GetUserFilters mocks base method.
func (m *MockUseCase) GetUserFilters(userId uint) (dataStruct.UserFilter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserFilters", userId)
	ret0, _ := ret[0].(dataStruct.UserFilter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserFilters indicates an expected call of GetUserFilters.
func (mr *MockUseCaseMockRecorder) GetUserFilters(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserFilters", reflect.TypeOf((*MockUseCase)(nil).GetUserFilters), userId)
}

// GetUserReasonsId mocks base method.
func (m *MockUseCase) GetUserReasonsId(userId uint) ([]uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserReasonsId", userId)
	ret0, _ := ret[0].([]uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserReasonsId indicates an expected call of GetUserReasonsId.
func (mr *MockUseCaseMockRecorder) GetUserReasonsId(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserReasonsId", reflect.TypeOf((*MockUseCase)(nil).GetUserReasonsId), userId)
}
