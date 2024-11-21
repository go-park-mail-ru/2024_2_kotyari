package user

import (
	proto "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/user/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func toModel(us *proto.UsersSignUpRequest) model.User {
	return model.User{
		Email:    us.GetEmail(),
		Username: us.GetUsername(),
		Password: us.GetPassword(),
	}
}

func toUserModel(us *proto.UsersLoginRequest) model.User {
	return model.User{
		Email:    us.GetEmail(),
		Password: us.GetPassword(),
	}
}
