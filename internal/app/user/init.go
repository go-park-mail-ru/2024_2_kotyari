package user

import (
	"context"
	proto "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/user/gen"
	"google.golang.org/grpc"
	"log/slog"
)

type usersDelivery interface {
	Register(grpcServer *grpc.Server)
	LoginUser(ctx context.Context, in *proto.UsersLoginRequest) (*proto.UsersDefaultResponse, error)
	GetUserById(ctx context.Context, in *proto.GetUserByIdRequest) (*proto.UsersDefaultResponse, error)
	CreateUser(ctx context.Context, in *proto.UsersSignUpRequest) (*proto.UsersDefaultResponse, error)
}

type UsersApp struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	delivery   usersDelivery
}

func NewUsersApp(log *slog.Logger, delivery usersDelivery) *UsersApp {
	grpcServer := grpc.NewServer()

	return &UsersApp{
		log:        log,
		gRPCServer: grpcServer,
		delivery:   delivery,
	}
}
