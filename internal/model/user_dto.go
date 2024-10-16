package model

type UserSignupRequestDTO struct {
	Email          string `json:"email"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

type UserLoginRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UsernameResponse struct {
	Username string `json:"username"`
}

func (sr *UserSignupRequestDTO) ToUserModel() *User {
	return &User{
		Username: sr.Username,
		Email:    sr.Email,
		Password: sr.Password,
	}
}
