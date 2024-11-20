package user

import (
	"context"
	proto "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/user/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (um *GrpcUserManager) CreateUser(ctx context.Context, in *proto.UsersSignUpRequest) (*proto.UsersDefaultResponse, error) {
	email := in.GetEmail()
	username := in.GetUsername()
	password := in.GetPassword()

	userModel, err := um.usersManager.CreateUser(ctx, model.User{
		Email:    email,
		Username: username,
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
