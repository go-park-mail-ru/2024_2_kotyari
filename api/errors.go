package api

import "errors"

type HTTPErrorResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

var (
	ErrInvalidJSONFormat   = errors.New("invalid JSON format")
	ErrEmptyRequestParams  = errors.New("empty request params")
	ErrWrongUsernameFormat = errors.New("wrong username format")
	ErrWrongEmailFormat    = errors.New("wrong email format")
	ErrWrongPasswordFormat = errors.New("wrong password format")
)
