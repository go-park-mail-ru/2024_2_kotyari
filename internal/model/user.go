package model

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserApiRequest struct {
	Email          string `json:"email"`
	Username       string `json:"username,omitempty"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password,omitempty"`
}

type UsernameResponse struct {
	Username string `json:"username"`
}
