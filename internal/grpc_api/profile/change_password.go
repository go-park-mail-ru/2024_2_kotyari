package profile

import (
	"context"

	profilegrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/profile/gen"
	"github.com/golang/protobuf/ptypes/empty"
)

func (p *ProfilesGrpc) ChangePassword(ctx context.Context, in *profilegrpc.ChangePasswordRequest) (*empty.Empty, error) {
	err := p.manager.ChangePassword(ctx,
		in.GetUserId(),
		in.GetOldPassword(),
		in.GetNewPassword(),
		in.GetRepeatPassword(),
	)
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
