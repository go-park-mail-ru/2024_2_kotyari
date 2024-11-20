package user

import (
	grpc_gen "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/user/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type UsersSignUpRequest struct {
	Email          string `json:"email"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

type UsersLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UsersDefaultResponse struct {
	UserID   uint32 `json:"user_id"`
	Username string `json:"username"`
	City     string `json:"city"`
}

func (ur *UsersSignUpRequest) ToGrpc() grpc_gen.UsersSignUpRequest {
	return grpc_gen.UsersSignUpRequest{
		Email:          ur.Email,
		Username:       ur.Username,
		HashedPassword: ur.Password,
	}
}

func (ur *UsersSignUpRequest) ToModel() model.User {
	return model.User{
		Email:    ur.Email,
		Username: ur.Username,
		Password: ur.Password,
	}
}

func (ur) FromModelToGrpcRequest() grpc_gen.UsersSignUpRequest {
	return grpc_gen.UsersSignUpRequest{
		Email:          ur.Email,
		Username:       ur.Username,
		HashedPassword: ur.Password,
	}
}

func (ul *UsersLoginRequest) ToModel() model.User {
	return model.User{
		Email:    ul.Email,
		Password: ul.Password,
	}
}
