// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: map.proto

package world

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

const (
	Map_Get_FullMethodName = "/box.svc.world.Map/Get"
)

// MapClient is the client API for Map service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MapClient interface {
	Get(ctx context.Context, in *GetMapRequest, opts ...grpc.CallOption) (*GetMapResponse, error)
}

type mapClient struct {
	cc grpc.ClientConnInterface
}

func NewMapClient(cc grpc.ClientConnInterface) MapClient {
	return &mapClient{cc}
}

func (c *mapClient) Get(ctx context.Context, in *GetMapRequest, opts ...grpc.CallOption) (*GetMapResponse, error) {
	out := new(GetMapResponse)
	err := c.cc.Invoke(ctx, Map_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MapServer is the server API for Map service.
// All implementations must embed UnimplementedMapServer
// for forward compatibility
type MapServer interface {
	Get(context.Context, *GetMapRequest) (*GetMapResponse, error)
	mustEmbedUnimplementedMapServer()
}

// UnimplementedMapServer must be embedded to have forward compatible implementations.
type UnimplementedMapServer struct {
}

func (UnimplementedMapServer) Get(context.Context, *GetMapRequest) (*GetMapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedMapServer) mustEmbedUnimplementedMapServer() {}

// UnsafeMapServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MapServer will
// result in compilation errors.
type UnsafeMapServer interface {
	mustEmbedUnimplementedMapServer()
}

func RegisterMapServer(s grpc.ServiceRegistrar, srv MapServer) {
	s.RegisterService(&Map_ServiceDesc, srv)
}

func _Map_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Map_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapServer).Get(ctx, req.(*GetMapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Map_ServiceDesc is the grpc.ServiceDesc for Map service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Map_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "box.svc.world.Map",
	HandlerType: (*MapServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Map_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "map.proto",
}