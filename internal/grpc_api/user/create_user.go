package user

import (
	"context"
	"errors"
	"log/slog"

	proto "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/user/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (um *UsersGrpc) CreateUser(ctx context.Context, in *proto.UsersSignUpRequest) (*proto.UsersDefaultResponse, error) {
	userModel, err := um.usersManager.CreateUser(ctx, toModel(in))
	if err != nil {
		if errors.Is(err, errs.UserAlreadyExists) {
			um.log.Error("[ UsersGrpc.CreateUser ] Пользователь уже существует",
				slog.String("error", err.Error()))

			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

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
