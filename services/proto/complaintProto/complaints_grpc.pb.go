// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: complaintProto.proto

package complaintProto

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ComplaintsClient is the client API for Complaints service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
//go:generate mockgen -source=complaints_grpc.pb.go -destination=mocks/comp_mock.go -package=mock
type ComplaintsClient interface {
	Complain(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*Response, error)
	CheckBanned(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*Response, error)
}

type complaintsClient struct {
	cc grpc.ClientConnInterface
}

func NewComplaintsClient(cc grpc.ClientConnInterface) ComplaintsClient {
	return &complaintsClient{cc}
}

func (c *complaintsClient) Complain(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/complaintProto.Complaints/Complain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *complaintsClient) CheckBanned(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/complaintProto.Complaints/CheckBanned", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComplaintsServer is the server API for Complaints service.
// All implementations must embed UnimplementedComplaintsServer
// for forward compatibility
type ComplaintsServer interface {
	Complain(context.Context, *UserId) (*Response, error)
	CheckBanned(context.Context, *UserId) (*Response, error)
	//mustEmbedUnimplementedComplaintsServer()
}

// UnimplementedComplaintsServer must be embedded to have forward compatible implementations.
type UnimplementedComplaintsServer struct {
}

func (UnimplementedComplaintsServer) Complain(context.Context, *UserId) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Complain not implemented")
}
func (UnimplementedComplaintsServer) CheckBanned(context.Context, *UserId) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckBanned not implemented")
}
func (UnimplementedComplaintsServer) mustEmbedUnimplementedComplaintsServer() {}

// UnsafeComplaintsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComplaintsServer will
// result in compilation errors.
type UnsafeComplaintsServer interface {
	mustEmbedUnimplementedComplaintsServer()
}

func RegisterComplaintsServer(s grpc.ServiceRegistrar, srv ComplaintsServer) {
	s.RegisterService(&Complaints_ServiceDesc, srv)
}

func _Complaints_Complain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComplaintsServer).Complain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/complaintProto.Complaints/Complain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComplaintsServer).Complain(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Complaints_CheckBanned_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComplaintsServer).CheckBanned(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/complaintProto.Complaints/CheckBanned",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComplaintsServer).CheckBanned(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

// Complaints_ServiceDesc is the grpc.ServiceDesc for Complaints service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Complaints_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "complaintProto.Complaints",
	HandlerType: (*ComplaintsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Complain",
			Handler:    _Complaints_Complain_Handler,
		},
		{
			MethodName: "CheckBanned",
			Handler:    _Complaints_CheckBanned_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "complaintProto.proto",
}
