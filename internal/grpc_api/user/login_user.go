package user

import (
	"context"
	proto "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/user/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (um *GrpcUserManager) LoginUser(ctx context.Context, in *proto.UsersLoginRequest) (*proto.UsersDefaultResponse, error) {
	email := in.GetEmail()
	password := in.GetPassword()

	userModel, err := um.usersManager.LoginUser(ctx, model.User{
		Email:    email,
		Password: password,
	})

	if err != nil {
		return &proto.UsersDefaultResponse{}, err
	}
	return &proto.UsersDefaultResponse{
		UserId:   userModel.ID,
		Username: userModel.Username,
		City:     userModel.City,
	}, nil
}
