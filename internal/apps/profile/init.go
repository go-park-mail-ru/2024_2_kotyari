package profile

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/profile/grpc_gen"
	"google.golang.org/grpc"
	"log/slog"
)

type profileDelivery interface {
	mustEmbedUnimplementedProfileServer()

	GetProfile(ctx context.Context, req *grpc_gen.GetProfileRequest) (*grpc_gen.GetProfileResponse, error)
	UpdateProfileInfo(ctx context.Context, req *grpc_gen.UpdateProfileInfoRequest) (*grpc_gen.UpdateProfileInfoResponse, error)
	UpdateProfileAvatar(ctx context.Context, req *grpc_gen.UpdateProfileInfoRequest) (*grpc_gen.UpdateProfileInfoResponse, error)
}

type ProfilesApp struct {
	delivery   profileDelivery
	log        *slog.Logger
	gRPCServer *grpc.Server
}

func Register(grpcServer *grpc.Server, profile profileDelivery) {
	grpc_gen.RegisterProfileServer(grpcServer, profile)
}

func NewProfilesApp(log *slog.Logger, authService authgrpc.Auth) {

}
