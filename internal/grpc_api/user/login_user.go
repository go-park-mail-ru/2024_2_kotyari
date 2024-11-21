package user

import (
	"context"
	proto "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/user/gen"
	"log/slog"
)

func (um *UsersGrpc) LoginUser(ctx context.Context, in *proto.UsersLoginRequest) (*proto.UsersDefaultResponse, error) {
	userModel, err := um.usersManager.LoginUser(ctx, toUserModel(in))
	if err != nil {
		um.log.Error("[ UsersGrpc.LoginUser ] Ошибка при отдаче на уровень usecase", slog.String("error", err.Error()))
		return &proto.UsersDefaultResponse{}, err
	}
	return &proto.UsersDefaultResponse{
		UserId:   userModel.ID,
		Username: userModel.Username,
		City:     userModel.City,
	}, nil
}
