// Code generated by MockGen. DO NOT EDIT.
// Source: complaints_grpc.pb.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	complaintProto "github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/complaintProto"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockComplaintsClient is a mock of ComplaintsClient interface.
type MockComplaintsClient struct {
	ctrl     *gomock.Controller
	recorder *MockComplaintsClientMockRecorder
}

// MockComplaintsClientMockRecorder is the mock recorder for MockComplaintsClient.
type MockComplaintsClientMockRecorder struct {
	mock *MockComplaintsClient
}

// NewMockComplaintsClient creates a new mock instance.
func NewMockComplaintsClient(ctrl *gomock.Controller) *MockComplaintsClient {
	mock := &MockComplaintsClient{ctrl: ctrl}
	mock.recorder = &MockComplaintsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComplaintsClient) EXPECT() *MockComplaintsClientMockRecorder {
	return m.recorder
}

// CheckBanned mocks base method.
func (m *MockComplaintsClient) CheckBanned(ctx context.Context, in *complaintProto.UserId, opts ...grpc.CallOption) (*complaintProto.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckBanned", varargs...)
	ret0, _ := ret[0].(*complaintProto.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckBanned indicates an expected call of CheckBanned.
func (mr *MockComplaintsClientMockRecorder) CheckBanned(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckBanned", reflect.TypeOf((*MockComplaintsClient)(nil).CheckBanned), varargs...)
}

// Complain mocks base method.
func (m *MockComplaintsClient) Complain(ctx context.Context, in *complaintProto.UserId, opts ...grpc.CallOption) (*complaintProto.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Complain", varargs...)
	ret0, _ := ret[0].(*complaintProto.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Complain indicates an expected call of Complain.
func (mr *MockComplaintsClientMockRecorder) Complain(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Complain", reflect.TypeOf((*MockComplaintsClient)(nil).Complain), varargs...)
}

// MockComplaintsServer is a mock of ComplaintsServer interface.
type MockComplaintsServer struct {
	ctrl     *gomock.Controller
	recorder *MockComplaintsServerMockRecorder
}

// MockComplaintsServerMockRecorder is the mock recorder for MockComplaintsServer.
type MockComplaintsServerMockRecorder struct {
	mock *MockComplaintsServer
}

// NewMockComplaintsServer creates a new mock instance.
func NewMockComplaintsServer(ctrl *gomock.Controller) *MockComplaintsServer {
	mock := &MockComplaintsServer{ctrl: ctrl}
	mock.recorder = &MockComplaintsServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComplaintsServer) EXPECT() *MockComplaintsServerMockRecorder {
	return m.recorder
}

// CheckBanned mocks base method.
func (m *MockComplaintsServer) CheckBanned(arg0 context.Context, arg1 *complaintProto.UserId) (*complaintProto.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckBanned", arg0, arg1)
	ret0, _ := ret[0].(*complaintProto.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckBanned indicates an expected call of CheckBanned.
func (mr *MockComplaintsServerMockRecorder) CheckBanned(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckBanned", reflect.TypeOf((*MockComplaintsServer)(nil).CheckBanned), arg0, arg1)
}

// Complain mocks base method.
func (m *MockComplaintsServer) Complain(arg0 context.Context, arg1 *complaintProto.UserId) (*complaintProto.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Complain", arg0, arg1)
	ret0, _ := ret[0].(*complaintProto.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Complain indicates an expected call of Complain.
func (mr *MockComplaintsServerMockRecorder) Complain(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Complain", reflect.TypeOf((*MockComplaintsServer)(nil).Complain), arg0, arg1)
}

// MockUnsafeComplaintsServer is a mock of UnsafeComplaintsServer interface.
type MockUnsafeComplaintsServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeComplaintsServerMockRecorder
}

// MockUnsafeComplaintsServerMockRecorder is the mock recorder for MockUnsafeComplaintsServer.
type MockUnsafeComplaintsServerMockRecorder struct {
	mock *MockUnsafeComplaintsServer
}

// NewMockUnsafeComplaintsServer creates a new mock instance.
func NewMockUnsafeComplaintsServer(ctrl *gomock.Controller) *MockUnsafeComplaintsServer {
	mock := &MockUnsafeComplaintsServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeComplaintsServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeComplaintsServer) EXPECT() *MockUnsafeComplaintsServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedComplaintsServer mocks base method.
func (m *MockUnsafeComplaintsServer) mustEmbedUnimplementedComplaintsServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedComplaintsServer")
}

// mustEmbedUnimplementedComplaintsServer indicates an expected call of mustEmbedUnimplementedComplaintsServer.
func (mr *MockUnsafeComplaintsServerMockRecorder) mustEmbedUnimplementedComplaintsServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedComplaintsServer", reflect.TypeOf((*MockUnsafeComplaintsServer)(nil).mustEmbedUnimplementedComplaintsServer))
}
