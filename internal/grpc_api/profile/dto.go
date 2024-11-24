package profile

import (
	profile_grpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/profile/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func toGrpcModel(p model.Profile) *profile_grpc.GetProfileResponse {
	return &profile_grpc.GetProfileResponse{
		Email:     p.Email,
		Username:  p.Username,
		Gender:    p.Gender,
		AvatarUrl: p.AvatarURL,
	}
}
