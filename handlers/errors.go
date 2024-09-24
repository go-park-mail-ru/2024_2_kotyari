package handlers

import "errors"

type HTTPErrorResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

var (
	ErrInvalidEmailFormat      = errors.New("Incorrect format of email")
	ErrInvalidPasswordFormat   = errors.New("Password must include uppercase, lowercase, number, and special character @$%*?&#.")
	ErrUnauthorizedCredentials = errors.New("Wrong email or password")
	ErrSessionCreationError    = errors.New("Error creating session")
	ErrSessionSaveError        = errors.New("Error when saving session")
	ErrLogoutError             = errors.New("Error when terminating session")
	ErrInternalServerError     = errors.New("Internal Server Error")
	ErrUnauthorizedMessage     = errors.New("Not authorized")
)
