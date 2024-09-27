package handlers

type credsApiRequest struct {
	Email          string `json:"email"`
	Username       string `json:"username,omitempty"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password,omitempty"`
}
