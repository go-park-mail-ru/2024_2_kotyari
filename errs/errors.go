package errs

import "errors"

type HTTPErrorResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

var (
	InvalidJSONFormat       = errors.New("invalid JSON format")
	InvalidUsernameFormat   = errors.New("wrong username format")
	InvalidEmailFormat      = errors.New("wrong email format")
	InvalidPasswordFormat   = errors.New("wrong password format")
	UserAlreadyExists       = errors.New("user already exists")
	InternalServerError     = errors.New("internal server error")
	SessionCreationError    = errors.New("error creating session")
	UnauthorizedCredentials = errors.New("wrong email or password")
	SessionSaveError        = errors.New("error when saving session")
	LogoutError             = errors.New("error when terminating session")
	UnauthorizedMessage     = errors.New("not authorized")
)
