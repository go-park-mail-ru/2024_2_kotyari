package profile

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/profile/grpc_gen"
	"google.golang.org/grpc"
)

type ProfilesGrpcDelivery struct {
	profileManager profileManager
	grpc_gen.UnimplementedProfileServer
}

func NewProfilesDelivery(profileManager profileManager) *ProfilesGrpcDelivery {

	return &ProfilesGrpcDelivery{profileManager: profileManager}
}

func (p *ProfilesGrpcDelivery) GetProfile(ctx context.Context, req *grpc_gen.GetProfileRequest) (*grpc_gen.GetProfileResponse, error) {
	panic("implement me")
}

func Register(grpcServer *grpc.Server, profile ProfilesGrpcDelivery) {
	grpc_gen.RegisterProfileServer(grpcServer, profile)
}

func (p *ProfilesGrpcDelivery) UpdateProfileInfo(ctx context.Context, req *grpc_gen.UpdateProfileInfoRequest) (*grpc_gen.UpdateProfileInfoResponse, error) {
	panic("implement me")
}

func (p *ProfilesGrpcDelivery) UpdateProfileAvatar(ctx context.Context, req *grpc_gen.UpdateProfileInfoRequest) (*grpc_gen.UpdateProfileInfoResponse, error) {
	panic("implement me")
}
