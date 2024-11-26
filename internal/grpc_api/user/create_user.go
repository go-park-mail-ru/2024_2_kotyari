package user

import (
	"context"
	proto "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/user/gen"
	"log/slog"
)

func (um *UsersGrpc) CreateUser(ctx context.Context, in *proto.UsersSignUpRequest) (*proto.UsersDefaultResponse, error) {
	userModel, err := um.usersManager.CreateUser(ctx, toModel(in))
	if err != nil {
		um.log.Error("[ UsersGrpc.CreateUser ] Ошибка при отдаче на уровень usecase", slog.String("error", err.Error()))
		return nil, err
	}

	um.log.Info("[ UsersGrpc.CreateUser ]", slog.Any("user", userModel))

	return &proto.UsersDefaultResponse{
		UserId:   userModel.ID,
		Username: userModel.Username,
		City:     userModel.City,
	}, nil
}
