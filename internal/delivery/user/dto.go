package user

import "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"

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
	Username  string `json:"username"`
	City      string `json:"city"`
	AvatarUrl string `json:"avatar_url"`
}

func (ur *UsersSignUpRequest) ToModel() model.User {
	return model.User{
		Email:    ur.Email,
		Username: ur.Username,
		Password: ur.Password,
	}
}

func (ul *UsersLoginRequest) ToModel() model.User {
	return model.User{
		Email:    ul.Email,
		Password: ul.Password,
	}
}
