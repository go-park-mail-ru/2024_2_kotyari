// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: init_promocodes.proto

package promocodes

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PromoCodes_GetUserPromoCodes_FullMethodName = "/promocodes.PromoCodes/GetUserPromoCodes"
)

// PromoCodesClient is the client API for PromoCodes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PromoCodesClient interface {
	GetUserPromoCodes(ctx context.Context, in *GetUserPromoCodesRequest, opts ...grpc.CallOption) (*GetUserPromoCodesResponse, error)
}

type promoCodesClient struct {
	cc grpc.ClientConnInterface
}

func NewPromoCodesClient(cc grpc.ClientConnInterface) PromoCodesClient {
	return &promoCodesClient{cc}
}

func (c *promoCodesClient) GetUserPromoCodes(ctx context.Context, in *GetUserPromoCodesRequest, opts ...grpc.CallOption) (*GetUserPromoCodesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserPromoCodesResponse)
	err := c.cc.Invoke(ctx, PromoCodes_GetUserPromoCodes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PromoCodesServer is the server API for PromoCodes service.
// All implementations must embed UnimplementedPromoCodesServer
// for forward compatibility.
type PromoCodesServer interface {
	GetUserPromoCodes(context.Context, *GetUserPromoCodesRequest) (*GetUserPromoCodesResponse, error)
	mustEmbedUnimplementedPromoCodesServer()
}

// UnimplementedPromoCodesServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPromoCodesServer struct{}

func (UnimplementedPromoCodesServer) GetUserPromoCodes(context.Context, *GetUserPromoCodesRequest) (*GetUserPromoCodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPromoCodes not implemented")
}
func (UnimplementedPromoCodesServer) mustEmbedUnimplementedPromoCodesServer() {}
func (UnimplementedPromoCodesServer) testEmbeddedByValue()                    {}

// UnsafePromoCodesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PromoCodesServer will
// result in compilation errors.
type UnsafePromoCodesServer interface {
	mustEmbedUnimplementedPromoCodesServer()
}

func RegisterPromoCodesServer(s grpc.ServiceRegistrar, srv PromoCodesServer) {
	// If the following call pancis, it indicates UnimplementedPromoCodesServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PromoCodes_ServiceDesc, srv)
}

func _PromoCodes_GetUserPromoCodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserPromoCodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromoCodesServer).GetUserPromoCodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PromoCodes_GetUserPromoCodes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromoCodesServer).GetUserPromoCodes(ctx, req.(*GetUserPromoCodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PromoCodes_ServiceDesc is the grpc.ServiceDesc for PromoCodes service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PromoCodes_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "promocodes.PromoCodes",
	HandlerType: (*PromoCodesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserPromoCodes",
			Handler:    _PromoCodes_GetUserPromoCodes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "init_promocodes.proto",
}
