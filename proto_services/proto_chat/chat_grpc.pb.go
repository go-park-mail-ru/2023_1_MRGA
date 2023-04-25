// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto_services/proto_chat/chat.proto

package proto_chat

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetRecentMessages(ctx context.Context, in *GetResentMessagesRequest, opts ...grpc.CallOption) (ChatService_GetRecentMessagesClient, error)
	GetConversationMessages(ctx context.Context, in *Message, opts ...grpc.CallOption) (ChatService_GetConversationMessagesClient, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto_chat.ChatService/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetRecentMessages(ctx context.Context, in *GetResentMessagesRequest, opts ...grpc.CallOption) (ChatService_GetRecentMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[0], "/proto_chat.ChatService/GetRecentMessages", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceGetRecentMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatService_GetRecentMessagesClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type chatServiceGetRecentMessagesClient struct {
	grpc.ClientStream
}

func (x *chatServiceGetRecentMessagesClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) GetConversationMessages(ctx context.Context, in *Message, opts ...grpc.CallOption) (ChatService_GetConversationMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[1], "/proto_chat.ChatService/GetConversationMessages", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceGetConversationMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatService_GetConversationMessagesClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type chatServiceGetConversationMessagesClient struct {
	grpc.ClientStream
}

func (x *chatServiceGetConversationMessagesClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	SendMessage(context.Context, *Message) (*emptypb.Empty, error)
	GetRecentMessages(*GetResentMessagesRequest, ChatService_GetRecentMessagesServer) error
	GetConversationMessages(*Message, ChatService_GetConversationMessagesServer) error
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) SendMessage(context.Context, *Message) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChatServiceServer) GetRecentMessages(*GetResentMessagesRequest, ChatService_GetRecentMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetRecentMessages not implemented")
}
func (UnimplementedChatServiceServer) GetConversationMessages(*Message, ChatService_GetConversationMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetConversationMessages not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_chat.ChatService/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SendMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetRecentMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetResentMessagesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).GetRecentMessages(m, &chatServiceGetRecentMessagesServer{stream})
}

type ChatService_GetRecentMessagesServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type chatServiceGetRecentMessagesServer struct {
	grpc.ServerStream
}

func (x *chatServiceGetRecentMessagesServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _ChatService_GetConversationMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Message)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).GetConversationMessages(m, &chatServiceGetConversationMessagesServer{stream})
}

type ChatService_GetConversationMessagesServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type chatServiceGetConversationMessagesServer struct {
	grpc.ServerStream
}

func (x *chatServiceGetConversationMessagesServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto_chat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _ChatService_SendMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetRecentMessages",
			Handler:       _ChatService_GetRecentMessages_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetConversationMessages",
			Handler:       _ChatService_GetConversationMessages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto_services/proto_chat/chat.proto",
}
