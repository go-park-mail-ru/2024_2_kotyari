// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: init.proto

package rating_updater

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RatingUpdater_UpdateRating_FullMethodName = "/rating_updater.RatingUpdater/UpdateRating"
)

// RatingUpdaterClient is the client API for RatingUpdater service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RatingUpdaterClient interface {
	UpdateRating(ctx context.Context, in *UpdateRatingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type ratingUpdaterClient struct {
	cc grpc.ClientConnInterface
}

func NewRatingUpdaterClient(cc grpc.ClientConnInterface) RatingUpdaterClient {
	return &ratingUpdaterClient{cc}
}

func (c *ratingUpdaterClient) UpdateRating(ctx context.Context, in *UpdateRatingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RatingUpdater_UpdateRating_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RatingUpdaterServer is the server API for RatingUpdater service.
// All implementations must embed UnimplementedRatingUpdaterServer
// for forward compatibility.
type RatingUpdaterServer interface {
	UpdateRating(context.Context, *UpdateRatingRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedRatingUpdaterServer()
}

// UnimplementedRatingUpdaterServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRatingUpdaterServer struct{}

func (UnimplementedRatingUpdaterServer) UpdateRating(context.Context, *UpdateRatingRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRating not implemented")
}
func (UnimplementedRatingUpdaterServer) mustEmbedUnimplementedRatingUpdaterServer() {}
func (UnimplementedRatingUpdaterServer) testEmbeddedByValue()                       {}

// UnsafeRatingUpdaterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RatingUpdaterServer will
// result in compilation errors.
type UnsafeRatingUpdaterServer interface {
	mustEmbedUnimplementedRatingUpdaterServer()
}

func RegisterRatingUpdaterServer(s grpc.ServiceRegistrar, srv RatingUpdaterServer) {
	// If the following call pancis, it indicates UnimplementedRatingUpdaterServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RatingUpdater_ServiceDesc, srv)
}

func _RatingUpdater_UpdateRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingUpdaterServer).UpdateRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingUpdater_UpdateRating_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingUpdaterServer).UpdateRating(ctx, req.(*UpdateRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RatingUpdater_ServiceDesc is the grpc.ServiceDesc for RatingUpdater service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RatingUpdater_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rating_updater.RatingUpdater",
	HandlerType: (*RatingUpdaterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateRating",
			Handler:    _RatingUpdater_UpdateRating_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "init.proto",
}
