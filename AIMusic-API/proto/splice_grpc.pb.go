// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: splice.proto

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

// AIMusicSpliceClient is the client API for AIMusicSplice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AIMusicSpliceClient interface {
	MidiSplice(ctx context.Context, in *InputFile, opts ...grpc.CallOption) (*OutputFile, error)
}

type aIMusicSpliceClient struct {
	cc grpc.ClientConnInterface
}

func NewAIMusicSpliceClient(cc grpc.ClientConnInterface) AIMusicSpliceClient {
	return &aIMusicSpliceClient{cc}
}

func (c *aIMusicSpliceClient) MidiSplice(ctx context.Context, in *InputFile, opts ...grpc.CallOption) (*OutputFile, error) {
	out := new(OutputFile)
	err := c.cc.Invoke(ctx, "/AIMusicSplice/MidiSplice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AIMusicSpliceServer is the server API for AIMusicSplice service.
// All implementations must embed UnimplementedAIMusicSpliceServer
// for forward compatibility
type AIMusicSpliceServer interface {
	MidiSplice(context.Context, *InputFile) (*OutputFile, error)
	mustEmbedUnimplementedAIMusicSpliceServer()
}

// UnimplementedAIMusicSpliceServer must be embedded to have forward compatible implementations.
type UnimplementedAIMusicSpliceServer struct {
}

func (UnimplementedAIMusicSpliceServer) MidiSplice(context.Context, *InputFile) (*OutputFile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MidiSplice not implemented")
}
func (UnimplementedAIMusicSpliceServer) mustEmbedUnimplementedAIMusicSpliceServer() {}

// UnsafeAIMusicSpliceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AIMusicSpliceServer will
// result in compilation errors.
type UnsafeAIMusicSpliceServer interface {
	mustEmbedUnimplementedAIMusicSpliceServer()
}

func RegisterAIMusicSpliceServer(s grpc.ServiceRegistrar, srv AIMusicSpliceServer) {
	s.RegisterService(&AIMusicSplice_ServiceDesc, srv)
}

func _AIMusicSplice_MidiSplice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InputFile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AIMusicSpliceServer).MidiSplice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AIMusicSplice/MidiSplice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AIMusicSpliceServer).MidiSplice(ctx, req.(*InputFile))
	}
	return interceptor(ctx, in, info, handler)
}

// AIMusicSplice_ServiceDesc is the grpc.ServiceDesc for AIMusicSplice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AIMusicSplice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AIMusicSplice",
	HandlerType: (*AIMusicSpliceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MidiSplice",
			Handler:    _AIMusicSplice_MidiSplice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "splice.proto",
}