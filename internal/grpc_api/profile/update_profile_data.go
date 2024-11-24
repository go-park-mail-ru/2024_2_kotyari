package profile

import (
	"context"
	profilegrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/profile/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/golang/protobuf/ptypes/empty"
	"log/slog"
)

func (p *ProfilesGrpc) UpdateProfileData(ctx context.Context, in *profilegrpc.UpdateProfileDataRequest) (*empty.Empty, error) {
	oldProfileData, err := p.manager.GetProfile(ctx, in.UserId)
	if err != nil {
		p.log.Warn("[ ProfilesDelivery.UpdateProfileData ] Не удалось получить старые данные профиля", slog.String("error", err.Error()))

		return nil, err
	}

	newProfileData := model.Profile{
		Email:    in.Email,
		Username: in.Username,
		Gender:   in.Gender,
	}

	if err = p.manager.UpdateProfile(ctx, oldProfileData, newProfileData); err != nil {
		p.log.Warn("[ ProfilesDelivery.UpdateProfileData ] Не удалось обновить данные профиля", slog.String("error", err.Error()))

		return nil, err
	}

	return &empty.Empty{}, nil
}
