package errs

import "errors"

type HTTPErrorResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

var (
	InvalidJSONFormat   = errors.New("invalid JSON format")
	EmptyRequestParams  = errors.New("empty request params")
	WrongUsernameFormat = errors.New("wrong username format")
	WrongEmailFormat    = errors.New("wrong email format")
	WrongPasswordFormat = errors.New("wrong password format")
	UserAlreadyExists   = errors.New("user already exists")
	InternalServerError = errors.New("internal server error")
)
