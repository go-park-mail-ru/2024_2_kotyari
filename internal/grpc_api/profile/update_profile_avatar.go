package profile

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"

	profilegrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/profile/gen"
)

func (p *ProfilesGrpc) UpdateProfileAvatar(ctx context.Context, in *profilegrpc.UpdateAvatarRequest) (*empty.Empty, error) {
	err := p.avatarSaver.UpdateProfileAvatar(ctx, in.UserId, in.Filepath)
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
