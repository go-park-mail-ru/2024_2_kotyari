// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: init.proto

package grpc_gen

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CsatService_GetCsat_FullMethodName       = "/rating.CsatService/GetCsat"
	CsatService_CreateCsat_FullMethodName    = "/rating.CsatService/CreateCsat"
	CsatService_GetStatistics_FullMethodName = "/rating.CsatService/GetStatistics"
	CsatService_UpdateCsat_FullMethodName    = "/rating.CsatService/UpdateCsat"
)

// CsatServiceClient is the client API for CsatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CsatServiceClient interface {
	GetCsat(ctx context.Context, in *GetCsatRequest, opts ...grpc.CallOption) (*GetCsatResponse, error)
	CreateCsat(ctx context.Context, in *CreateCsatRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetStatistics(ctx context.Context, in *GetStatisticsRequest, opts ...grpc.CallOption) (*GetStatisticsResponse, error)
	UpdateCsat(ctx context.Context, in *UpdateCsatRequest, opts ...grpc.CallOption) (*UpdateCsatResponse, error)
}

type csatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCsatServiceClient(cc grpc.ClientConnInterface) CsatServiceClient {
	return &csatServiceClient{cc}
}

func (c *csatServiceClient) GetCsat(ctx context.Context, in *GetCsatRequest, opts ...grpc.CallOption) (*GetCsatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCsatResponse)
	err := c.cc.Invoke(ctx, CsatService_GetCsat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *csatServiceClient) CreateCsat(ctx context.Context, in *CreateCsatRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, CsatService_CreateCsat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *csatServiceClient) GetStatistics(ctx context.Context, in *GetStatisticsRequest, opts ...grpc.CallOption) (*GetStatisticsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetStatisticsResponse)
	err := c.cc.Invoke(ctx, CsatService_GetStatistics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *csatServiceClient) UpdateCsat(ctx context.Context, in *UpdateCsatRequest, opts ...grpc.CallOption) (*UpdateCsatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCsatResponse)
	err := c.cc.Invoke(ctx, CsatService_UpdateCsat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CsatServiceServer is the server API for CsatService service.
// All implementations must embed UnimplementedCsatServiceServer
// for forward compatibility.
type CsatServiceServer interface {
	GetCsat(context.Context, *GetCsatRequest) (*GetCsatResponse, error)
	CreateCsat(context.Context, *CreateCsatRequest) (*empty.Empty, error)
	GetStatistics(context.Context, *GetStatisticsRequest) (*GetStatisticsResponse, error)
	UpdateCsat(context.Context, *UpdateCsatRequest) (*UpdateCsatResponse, error)
	mustEmbedUnimplementedCsatServiceServer()
}

// UnimplementedCsatServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCsatServiceServer struct{}

func (UnimplementedCsatServiceServer) GetCsat(context.Context, *GetCsatRequest) (*GetCsatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCsat not implemented")
}
func (UnimplementedCsatServiceServer) CreateCsat(context.Context, *CreateCsatRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCsat not implemented")
}
func (UnimplementedCsatServiceServer) GetStatistics(context.Context, *GetStatisticsRequest) (*GetStatisticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatistics not implemented")
}
func (UnimplementedCsatServiceServer) UpdateCsat(context.Context, *UpdateCsatRequest) (*UpdateCsatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCsat not implemented")
}
func (UnimplementedCsatServiceServer) mustEmbedUnimplementedCsatServiceServer() {}
func (UnimplementedCsatServiceServer) testEmbeddedByValue()                     {}

// UnsafeCsatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CsatServiceServer will
// result in compilation errors.
type UnsafeCsatServiceServer interface {
	mustEmbedUnimplementedCsatServiceServer()
}

func RegisterCsatServiceServer(s grpc.ServiceRegistrar, srv CsatServiceServer) {
	// If the following call pancis, it indicates UnimplementedCsatServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CsatService_ServiceDesc, srv)
}

func _CsatService_GetCsat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCsatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CsatServiceServer).GetCsat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CsatService_GetCsat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CsatServiceServer).GetCsat(ctx, req.(*GetCsatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CsatService_CreateCsat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCsatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CsatServiceServer).CreateCsat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CsatService_CreateCsat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CsatServiceServer).CreateCsat(ctx, req.(*CreateCsatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CsatService_GetStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatisticsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CsatServiceServer).GetStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CsatService_GetStatistics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CsatServiceServer).GetStatistics(ctx, req.(*GetStatisticsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CsatService_UpdateCsat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCsatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CsatServiceServer).UpdateCsat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CsatService_UpdateCsat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CsatServiceServer).UpdateCsat(ctx, req.(*UpdateCsatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CsatService_ServiceDesc is the grpc.ServiceDesc for CsatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CsatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rating.CsatService",
	HandlerType: (*CsatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCsat",
			Handler:    _CsatService_GetCsat_Handler,
		},
		{
			MethodName: "CreateCsat",
			Handler:    _CsatService_CreateCsat_Handler,
		},
		{
			MethodName: "GetStatistics",
			Handler:    _CsatService_GetStatistics_Handler,
		},
		{
			MethodName: "UpdateCsat",
			Handler:    _CsatService_UpdateCsat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "init.proto",
}
