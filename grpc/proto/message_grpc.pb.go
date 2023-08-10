// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/message.proto

package proto

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

// MessangerClient is the client API for Messanger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessangerClient interface {
	Messager(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type messangerClient struct {
	cc grpc.ClientConnInterface
}

func NewMessangerClient(cc grpc.ClientConnInterface) MessangerClient {
	return &messangerClient{cc}
}

func (c *messangerClient) Messager(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/messanger/messager", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessangerServer is the server API for Messanger service.
// All implementations must embed UnimplementedMessangerServer
// for forward compatibility
type MessangerServer interface {
	Messager(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedMessangerServer()
}

// UnimplementedMessangerServer must be embedded to have forward compatible implementations.
type UnimplementedMessangerServer struct {
}

func (UnimplementedMessangerServer) Messager(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Messager not implemented")
}
func (UnimplementedMessangerServer) mustEmbedUnimplementedMessangerServer() {}

// UnsafeMessangerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessangerServer will
// result in compilation errors.
type UnsafeMessangerServer interface {
	mustEmbedUnimplementedMessangerServer()
}

func RegisterMessangerServer(s grpc.ServiceRegistrar, srv MessangerServer) {
	s.RegisterService(&Messanger_ServiceDesc, srv)
}

func _Messanger_Messager_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessangerServer).Messager(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messanger/messager",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessangerServer).Messager(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Messanger_ServiceDesc is the grpc.ServiceDesc for Messanger service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Messanger_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "messanger",
	HandlerType: (*MessangerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "messager",
			Handler:    _Messanger_Messager_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/message.proto",
}
