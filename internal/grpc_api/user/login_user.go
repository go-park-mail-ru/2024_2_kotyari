package user

import (
	"context"
	"errors"
	proto "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/user/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (um *UsersGrpc) LoginUser(ctx context.Context, in *proto.UsersLoginRequest) (*proto.UsersDefaultResponse, error) {
	userModel, err := um.usersManager.LoginUser(ctx, toUserModel(in))
	if err != nil {
		switch {
		case errors.Is(err, errs.UserDoesNotExist):
			um.log.Error("[ UsersGrpc.LoginUser ] ", slog.String("error", err.Error()))

			return &proto.UsersDefaultResponse{}, status.Error(codes.NotFound, err.Error())
		case errors.Is(err, errs.WrongCredentials):
			um.log.Error("[ UsersGrpc.LoginUser ] ", slog.String("error", err.Error()))

			return &proto.UsersDefaultResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}

		um.log.Error("[ UsersGrpc.LoginUser ] ", slog.String("error", err.Error()))

		return &proto.UsersDefaultResponse{}, status.Error(codes.Internal, errs.InternalServerError.Error())
	}

	return &proto.UsersDefaultResponse{
		UserId:   userModel.ID,
		Username: userModel.Username,
		City:     userModel.City,
	}, nil
}
