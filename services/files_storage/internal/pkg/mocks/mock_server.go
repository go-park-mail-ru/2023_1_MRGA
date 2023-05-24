// Code generated by MockGen. DO NOT EDIT.
// Source: services/files_storage/internal/app/iserver.go

// Package mocks is a generated GoMock package.
package mocks

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIServer is a mock of IServer interface.
type MockIServer struct {
	ctrl     *gomock.Controller
	recorder *MockIServerMockRecorder
}

// MockIServerMockRecorder is the mock recorder for MockIServer.
type MockIServerMockRecorder struct {
	mock *MockIServer
}

// NewMockIServer creates a new mock instance.
func NewMockIServer(ctrl *gomock.Controller) *MockIServer {
	mock := &MockIServer{ctrl: ctrl}
	mock.recorder = &MockIServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIServer) EXPECT() *MockIServerMockRecorder {
	return m.recorder
}

// GetFile mocks base method.
func (m *MockIServer) GetFile(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetFile", arg0, arg1)
}

// GetFile indicates an expected call of GetFile.
func (mr *MockIServerMockRecorder) GetFile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFile", reflect.TypeOf((*MockIServer)(nil).GetFile), arg0, arg1)
}

// GetFileByPath mocks base method.
func (m *MockIServer) GetFileByPath(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetFileByPath", arg0, arg1)
}

// GetFileByPath indicates an expected call of GetFileByPath.
func (mr *MockIServerMockRecorder) GetFileByPath(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileByPath", reflect.TypeOf((*MockIServer)(nil).GetFileByPath), arg0, arg1)
}

// UploadFile mocks base method.
func (m *MockIServer) UploadFile(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UploadFile", arg0, arg1)
}

// UploadFile indicates an expected call of UploadFile.
func (mr *MockIServerMockRecorder) UploadFile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFile", reflect.TypeOf((*MockIServer)(nil).UploadFile), arg0, arg1)
}

// UploadPhoto mocks base method.
func (m *MockIServer) UploadPhoto(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UploadPhoto", arg0, arg1)
}

// UploadPhoto indicates an expected call of UploadPhoto.
func (mr *MockIServerMockRecorder) UploadPhoto(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadPhoto", reflect.TypeOf((*MockIServer)(nil).UploadPhoto), arg0, arg1)
}
