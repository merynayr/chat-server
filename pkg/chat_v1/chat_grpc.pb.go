// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: chat.proto

package chat_v1

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ChatV1_CreateChat_FullMethodName  = "/chat_v1.ChatV1/CreateChat"
	ChatV1_DeleteChat_FullMethodName  = "/chat_v1.ChatV1/DeleteChat"
	ChatV1_SendMessage_FullMethodName = "/chat_v1.ChatV1/SendMessage"
)

// ChatV1Client is the client API for ChatV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatV1Client interface {
	// Создает новый чат
	CreateChat(ctx context.Context, in *CreateChatRequest, opts ...grpc.CallOption) (*CreateChatResponse, error)
	// Удаляет чат по id
	DeleteChat(ctx context.Context, in *DeleteChatRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Отправляет сообщение в чат
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type chatV1Client struct {
	cc grpc.ClientConnInterface
}

func NewChatV1Client(cc grpc.ClientConnInterface) ChatV1Client {
	return &chatV1Client{cc}
}

func (c *chatV1Client) CreateChat(ctx context.Context, in *CreateChatRequest, opts ...grpc.CallOption) (*CreateChatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateChatResponse)
	err := c.cc.Invoke(ctx, ChatV1_CreateChat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatV1Client) DeleteChat(ctx context.Context, in *DeleteChatRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, ChatV1_DeleteChat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatV1Client) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, ChatV1_SendMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatV1Server is the server API for ChatV1 service.
// All implementations must embed UnimplementedChatV1Server
// for forward compatibility
type ChatV1Server interface {
	// Создает новый чат
	CreateChat(context.Context, *CreateChatRequest) (*CreateChatResponse, error)
	// Удаляет чат по id
	DeleteChat(context.Context, *DeleteChatRequest) (*empty.Empty, error)
	// Отправляет сообщение в чат
	SendMessage(context.Context, *SendMessageRequest) (*empty.Empty, error)
	mustEmbedUnimplementedChatV1Server()
}

// UnimplementedChatV1Server must be embedded to have forward compatible implementations.
type UnimplementedChatV1Server struct {
}

func (UnimplementedChatV1Server) CreateChat(context.Context, *CreateChatRequest) (*CreateChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChat not implemented")
}
func (UnimplementedChatV1Server) DeleteChat(context.Context, *DeleteChatRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChat not implemented")
}
func (UnimplementedChatV1Server) SendMessage(context.Context, *SendMessageRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChatV1Server) mustEmbedUnimplementedChatV1Server() {}

// UnsafeChatV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatV1Server will
// result in compilation errors.
type UnsafeChatV1Server interface {
	mustEmbedUnimplementedChatV1Server()
}

func RegisterChatV1Server(s grpc.ServiceRegistrar, srv ChatV1Server) {
	s.RegisterService(&ChatV1_ServiceDesc, srv)
}

func _ChatV1_CreateChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatV1Server).CreateChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatV1_CreateChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatV1Server).CreateChat(ctx, req.(*CreateChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatV1_DeleteChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatV1Server).DeleteChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatV1_DeleteChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatV1Server).DeleteChat(ctx, req.(*DeleteChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatV1_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatV1Server).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatV1_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatV1Server).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatV1_ServiceDesc is the grpc.ServiceDesc for ChatV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat_v1.ChatV1",
	HandlerType: (*ChatV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChat",
			Handler:    _ChatV1_CreateChat_Handler,
		},
		{
			MethodName: "DeleteChat",
			Handler:    _ChatV1_DeleteChat_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _ChatV1_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}
