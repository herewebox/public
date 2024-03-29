// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: tenant.proto

package tenant

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
	Tenant_Get_FullMethodName = "/box.tenant.Tenant/Get"
)

// TenantClient is the client API for Tenant service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TenantClient interface {
	Get(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*GetTenantResponse, error)
}

type tenantClient struct {
	cc grpc.ClientConnInterface
}

func NewTenantClient(cc grpc.ClientConnInterface) TenantClient {
	return &tenantClient{cc}
}

func (c *tenantClient) Get(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*GetTenantResponse, error) {
	out := new(GetTenantResponse)
	err := c.cc.Invoke(ctx, Tenant_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TenantServer is the server API for Tenant service.
// All implementations must embed UnimplementedTenantServer
// for forward compatibility
type TenantServer interface {
	Get(context.Context, *GetTenantRequest) (*GetTenantResponse, error)
	mustEmbedUnimplementedTenantServer()
}

// UnimplementedTenantServer must be embedded to have forward compatible implementations.
type UnimplementedTenantServer struct {
}

func (UnimplementedTenantServer) Get(context.Context, *GetTenantRequest) (*GetTenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTenantServer) mustEmbedUnimplementedTenantServer() {}

// UnsafeTenantServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TenantServer will
// result in compilation errors.
type UnsafeTenantServer interface {
	mustEmbedUnimplementedTenantServer()
}

func RegisterTenantServer(s grpc.ServiceRegistrar, srv TenantServer) {
	s.RegisterService(&Tenant_ServiceDesc, srv)
}

func _Tenant_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tenant_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServer).Get(ctx, req.(*GetTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Tenant_ServiceDesc is the grpc.ServiceDesc for Tenant service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tenant_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "box.tenant.Tenant",
	HandlerType: (*TenantServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Tenant_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tenant.proto",
}
