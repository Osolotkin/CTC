// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: gogo.proto

package gogo

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

// GoGoServiceClient is the client API for GoGoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoGoServiceClient interface {
	Ping(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	Get(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	Post(ctx context.Context, in *KeyValuePair, opts ...grpc.CallOption) (*Message, error)
	List(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	Delete(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type goGoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGoGoServiceClient(cc grpc.ClientConnInterface) GoGoServiceClient {
	return &goGoServiceClient{cc}
}

func (c *goGoServiceClient) Ping(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/gogo.GoGoService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goGoServiceClient) Get(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/gogo.GoGoService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goGoServiceClient) Post(ctx context.Context, in *KeyValuePair, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/gogo.GoGoService/Post", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goGoServiceClient) List(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/gogo.GoGoService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goGoServiceClient) Delete(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/gogo.GoGoService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoGoServiceServer is the server API for GoGoService service.
// All implementations must embed UnimplementedGoGoServiceServer
// for forward compatibility
type GoGoServiceServer interface {
	Ping(context.Context, *Message) (*Message, error)
	Get(context.Context, *Message) (*Message, error)
	Post(context.Context, *KeyValuePair) (*Message, error)
	List(context.Context, *Message) (*Message, error)
	Delete(context.Context, *Message) (*Message, error)
	mustEmbedUnimplementedGoGoServiceServer()
}

// UnimplementedGoGoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGoGoServiceServer struct {
}

func (UnimplementedGoGoServiceServer) Ping(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedGoGoServiceServer) Get(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedGoGoServiceServer) Post(context.Context, *KeyValuePair) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
}
func (UnimplementedGoGoServiceServer) List(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedGoGoServiceServer) Delete(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedGoGoServiceServer) mustEmbedUnimplementedGoGoServiceServer() {}

// UnsafeGoGoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoGoServiceServer will
// result in compilation errors.
type UnsafeGoGoServiceServer interface {
	mustEmbedUnimplementedGoGoServiceServer()
}

func RegisterGoGoServiceServer(s grpc.ServiceRegistrar, srv GoGoServiceServer) {
	s.RegisterService(&GoGoService_ServiceDesc, srv)
}

func _GoGoService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoGoServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gogo.GoGoService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoGoServiceServer).Ping(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoGoService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoGoServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gogo.GoGoService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoGoServiceServer).Get(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoGoService_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValuePair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoGoServiceServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gogo.GoGoService/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoGoServiceServer).Post(ctx, req.(*KeyValuePair))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoGoService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoGoServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gogo.GoGoService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoGoServiceServer).List(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoGoService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoGoServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gogo.GoGoService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoGoServiceServer).Delete(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

// GoGoService_ServiceDesc is the grpc.ServiceDesc for GoGoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoGoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gogo.GoGoService",
	HandlerType: (*GoGoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _GoGoService_Ping_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _GoGoService_Get_Handler,
		},
		{
			MethodName: "Post",
			Handler:    _GoGoService_Post_Handler,
		},
		{
			MethodName: "List",
			Handler:    _GoGoService_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _GoGoService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gogo.proto",
}