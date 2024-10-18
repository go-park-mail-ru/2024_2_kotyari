package model

type User struct {
	ID       uint32
	Email    string
	Username string
	City     string
	Password string
}

func (u *User) ToUserDTO() *UserDTO {
	return &UserDTO{
		Username: u.Username,
		Email:    u.Email,
	}
}
