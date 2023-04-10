// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: creation.proto

package proto

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

// CreationClient is the client API for Creation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CreationClient interface {
	GetTemplate(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTemplateResponse, error)
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
}

type creationClient struct {
	cc grpc.ClientConnInterface
}

func NewCreationClient(cc grpc.ClientConnInterface) CreationClient {
	return &creationClient{cc}
}

func (c *creationClient) GetTemplate(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTemplateResponse, error) {
	out := new(GetTemplateResponse)
	err := c.cc.Invoke(ctx, "/Creation/GetTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/Creation/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreationServer is the server API for Creation service.
// All implementations must embed UnimplementedCreationServer
// for forward compatibility
type CreationServer interface {
	GetTemplate(context.Context, *emptypb.Empty) (*GetTemplateResponse, error)
	Add(context.Context, *AddRequest) (*AddResponse, error)
	mustEmbedUnimplementedCreationServer()
}

// UnimplementedCreationServer must be embedded to have forward compatible implementations.
type UnimplementedCreationServer struct {
}

func (UnimplementedCreationServer) GetTemplate(context.Context, *emptypb.Empty) (*GetTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTemplate not implemented")
}
func (UnimplementedCreationServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedCreationServer) mustEmbedUnimplementedCreationServer() {}

// UnsafeCreationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CreationServer will
// result in compilation errors.
type UnsafeCreationServer interface {
	mustEmbedUnimplementedCreationServer()
}

func RegisterCreationServer(s grpc.ServiceRegistrar, srv CreationServer) {
	s.RegisterService(&Creation_ServiceDesc, srv)
}

func _Creation_GetTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServer).GetTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Creation/GetTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServer).GetTemplate(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Creation_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Creation/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Creation_ServiceDesc is the grpc.ServiceDesc for Creation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Creation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Creation",
	HandlerType: (*CreationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTemplate",
			Handler:    _Creation_GetTemplate_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Creation_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "creation.proto",
}
