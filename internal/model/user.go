package model

type User struct {
	Username string
	Email    string
	Password string
}

func (u *User) ToUserDTO() *UserDTO {
	return &UserDTO{
		Username: u.Username,
		Email:    u.Email,
	}
}
