package profile

import (
	"context"
	profilegrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/profile/gen"
)

func (p *ProfilesGrpc) GetProfile(ctx context.Context, in *profilegrpc.GetProfileRequest) (*profilegrpc.GetProfileResponse, error) {
	profile, err := p.manager.GetProfile(ctx, in.GetUserId())
	if err != nil {
		return nil, err
	}

	return &profilegrpc.GetProfileResponse{
		UserId:    profile.ID,
		Email:     profile.Email,
		Username:  profile.Username,
		Gender:    profile.Gender,
		Age:       profile.Age,
		AvatarUrl: profile.AvatarURL,
	}, nil
}
