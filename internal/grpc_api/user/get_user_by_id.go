package user

import (
	"context"

	proto "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/user/gen"
)

func (um *GrpcUserManager) GetUserById(ctx context.Context, in *proto.GetUserByIdRequest) (*proto.UsersDefaultResponse, error) {
	userId := in.GetUserId()

	userModel, err := um.userGetter.GetUserByUserID(ctx, userId)

	if err != nil {
		return &proto.UsersDefaultResponse{}, err
	}
	return &proto.UsersDefaultResponse{
		UserId:   userModel.ID,
		Username: userModel.Username,
		City:     userModel.City,
	}, nil
}
