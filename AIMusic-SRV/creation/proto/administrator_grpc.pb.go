// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: administrator.proto

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

// AdministratorClient is the client API for Administrator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdministratorClient interface {
	GetUsage(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Usage, error)
	ExecuteUser(ctx context.Context, in *ExecutionUserRequest, opts ...grpc.CallOption) (*ExecutionResponse, error)
	ExecuteCreation(ctx context.Context, in *ExecutionCreationRequest, opts ...grpc.CallOption) (*ExecutionResponse, error)
	EditBulletin(ctx context.Context, in *EditBulletinMsg, opts ...grpc.CallOption) (*EditBulletinMsg, error)
	UserAudit(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserAuditResponse, error)
	CreationAudit(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CreationAuditResponse, error)
	GetUserFeedback(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserFeedbackResponse, error)
}

type administratorClient struct {
	cc grpc.ClientConnInterface
}

func NewAdministratorClient(cc grpc.ClientConnInterface) AdministratorClient {
	return &administratorClient{cc}
}

func (c *administratorClient) GetUsage(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Usage, error) {
	out := new(Usage)
	err := c.cc.Invoke(ctx, "/Administrator/GetUsage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *administratorClient) ExecuteUser(ctx context.Context, in *ExecutionUserRequest, opts ...grpc.CallOption) (*ExecutionResponse, error) {
	out := new(ExecutionResponse)
	err := c.cc.Invoke(ctx, "/Administrator/ExecuteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *administratorClient) ExecuteCreation(ctx context.Context, in *ExecutionCreationRequest, opts ...grpc.CallOption) (*ExecutionResponse, error) {
	out := new(ExecutionResponse)
	err := c.cc.Invoke(ctx, "/Administrator/ExecuteCreation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *administratorClient) EditBulletin(ctx context.Context, in *EditBulletinMsg, opts ...grpc.CallOption) (*EditBulletinMsg, error) {
	out := new(EditBulletinMsg)
	err := c.cc.Invoke(ctx, "/Administrator/EditBulletin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *administratorClient) UserAudit(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserAuditResponse, error) {
	out := new(UserAuditResponse)
	err := c.cc.Invoke(ctx, "/Administrator/UserAudit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *administratorClient) CreationAudit(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CreationAuditResponse, error) {
	out := new(CreationAuditResponse)
	err := c.cc.Invoke(ctx, "/Administrator/CreationAudit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *administratorClient) GetUserFeedback(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserFeedbackResponse, error) {
	out := new(UserFeedbackResponse)
	err := c.cc.Invoke(ctx, "/Administrator/GetUserFeedback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdministratorServer is the server API for Administrator service.
// All implementations must embed UnimplementedAdministratorServer
// for forward compatibility
type AdministratorServer interface {
	GetUsage(context.Context, *emptypb.Empty) (*Usage, error)
	ExecuteUser(context.Context, *ExecutionUserRequest) (*ExecutionResponse, error)
	ExecuteCreation(context.Context, *ExecutionCreationRequest) (*ExecutionResponse, error)
	EditBulletin(context.Context, *EditBulletinMsg) (*EditBulletinMsg, error)
	UserAudit(context.Context, *emptypb.Empty) (*UserAuditResponse, error)
	CreationAudit(context.Context, *emptypb.Empty) (*CreationAuditResponse, error)
	GetUserFeedback(context.Context, *emptypb.Empty) (*UserFeedbackResponse, error)
	mustEmbedUnimplementedAdministratorServer()
}

// UnimplementedAdministratorServer must be embedded to have forward compatible implementations.
type UnimplementedAdministratorServer struct {
}

func (UnimplementedAdministratorServer) GetUsage(context.Context, *emptypb.Empty) (*Usage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsage not implemented")
}
func (UnimplementedAdministratorServer) ExecuteUser(context.Context, *ExecutionUserRequest) (*ExecutionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteUser not implemented")
}
func (UnimplementedAdministratorServer) ExecuteCreation(context.Context, *ExecutionCreationRequest) (*ExecutionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteCreation not implemented")
}
func (UnimplementedAdministratorServer) EditBulletin(context.Context, *EditBulletinMsg) (*EditBulletinMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditBulletin not implemented")
}
func (UnimplementedAdministratorServer) UserAudit(context.Context, *emptypb.Empty) (*UserAuditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAudit not implemented")
}
func (UnimplementedAdministratorServer) CreationAudit(context.Context, *emptypb.Empty) (*CreationAuditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreationAudit not implemented")
}
func (UnimplementedAdministratorServer) GetUserFeedback(context.Context, *emptypb.Empty) (*UserFeedbackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFeedback not implemented")
}
func (UnimplementedAdministratorServer) mustEmbedUnimplementedAdministratorServer() {}

// UnsafeAdministratorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdministratorServer will
// result in compilation errors.
type UnsafeAdministratorServer interface {
	mustEmbedUnimplementedAdministratorServer()
}

func RegisterAdministratorServer(s grpc.ServiceRegistrar, srv AdministratorServer) {
	s.RegisterService(&Administrator_ServiceDesc, srv)
}

func _Administrator_GetUsage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdministratorServer).GetUsage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Administrator/GetUsage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdministratorServer).GetUsage(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Administrator_ExecuteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecutionUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdministratorServer).ExecuteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Administrator/ExecuteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdministratorServer).ExecuteUser(ctx, req.(*ExecutionUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Administrator_ExecuteCreation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecutionCreationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdministratorServer).ExecuteCreation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Administrator/ExecuteCreation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdministratorServer).ExecuteCreation(ctx, req.(*ExecutionCreationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Administrator_EditBulletin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditBulletinMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdministratorServer).EditBulletin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Administrator/EditBulletin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdministratorServer).EditBulletin(ctx, req.(*EditBulletinMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Administrator_UserAudit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdministratorServer).UserAudit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Administrator/UserAudit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdministratorServer).UserAudit(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Administrator_CreationAudit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdministratorServer).CreationAudit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Administrator/CreationAudit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdministratorServer).CreationAudit(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Administrator_GetUserFeedback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdministratorServer).GetUserFeedback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Administrator/GetUserFeedback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdministratorServer).GetUserFeedback(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Administrator_ServiceDesc is the grpc.ServiceDesc for Administrator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Administrator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Administrator",
	HandlerType: (*AdministratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUsage",
			Handler:    _Administrator_GetUsage_Handler,
		},
		{
			MethodName: "ExecuteUser",
			Handler:    _Administrator_ExecuteUser_Handler,
		},
		{
			MethodName: "ExecuteCreation",
			Handler:    _Administrator_ExecuteCreation_Handler,
		},
		{
			MethodName: "EditBulletin",
			Handler:    _Administrator_EditBulletin_Handler,
		},
		{
			MethodName: "UserAudit",
			Handler:    _Administrator_UserAudit_Handler,
		},
		{
			MethodName: "CreationAudit",
			Handler:    _Administrator_CreationAudit_Handler,
		},
		{
			MethodName: "GetUserFeedback",
			Handler:    _Administrator_GetUserFeedback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "administrator.proto",
}
