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

func (us *UsersSignUpRequest) ToModel() model.User {
	return model.User{
		Email:    us.Email,
		Username: us.Username,
		Password: us.Password,
	}
}

func (ul *UsersLoginRequest) ToModel() model.User {
	return model.User{
		Email:    ul.Email,
		Password: ul.Password,
	}
}

func (us *UsersSignUpRequest) ToGrpcSignupRequest() *grpc_gen.UsersSignUpRequest {
	return &grpc_gen.UsersSignUpRequest{
		Username: us.Username,
		Email:    us.Email,
		Password: us.Password,
	}
}

func (ul *UsersLoginRequest) ToGrpcLoginRequest() *grpc_gen.UsersLoginRequest {
	return &grpc_gen.UsersLoginRequest{
		Email:    ul.Email,
		Password: ul.Password,
	}
}
