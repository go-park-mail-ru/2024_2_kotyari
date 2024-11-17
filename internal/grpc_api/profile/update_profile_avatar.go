package profile

import (
	"context"

	profilegrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/profile/gen"
)

func (p *ProfilesGrpc) UpdateProfileAvatar(ctx context.Context, in *profilegrpc.UpdateAvatarRequest) (*profilegrpc.UpdateAvatarResponse, error) {
	err := p.avatarSaver.UpdateProfileAvatar(ctx, in.UserId, in.Filepath)
	if err != nil {
		return nil, err
	}

	return &profilegrpc.UpdateAvatarResponse{
		Filepath: in.Filepath,
	}, nil
}
