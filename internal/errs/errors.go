package errs

import (
	"errors"
	"net/http"
)

type HTTPErrorResponse struct {
	ErrorMessage string `json:"error_message"`
}

var (
	InvalidJSONFormat     = errors.New("неверный формат JSON")
	InvalidUsernameFormat = errors.New("неверный формат имени пользователя")
	InvalidEmailFormat    = errors.New("неверный формат email")
	InvalidPasswordFormat = errors.New("неверный формат пароля")
	PasswordsDoNotMatch   = errors.New("пароли не совпадают")
	UserAlreadyExists     = errors.New("пользователь уже существует")
	InternalServerError   = errors.New("внутренняя ошибка сервера")
	SessionCreationError  = errors.New("ошибка при создании сессии")
	SessionSaveError      = errors.New("ошибка при сохранении сессии")
	WrongCredentials      = errors.New("неверный email или пароль")
	UserNotAuthorised     = errors.New("пользователь не авторизован")
	LogoutError           = errors.New("ошибка при завершении сессии")
)

var ErrCodesMapping = map[error]int{
	InvalidJSONFormat:     http.StatusBadRequest,
	InvalidUsernameFormat: http.StatusBadRequest,
	InvalidEmailFormat:    http.StatusBadRequest,
	InvalidPasswordFormat: http.StatusBadRequest,
	UserAlreadyExists:     http.StatusConflict,
	InternalServerError:   http.StatusInternalServerError,
	SessionCreationError:  http.StatusInternalServerError,
	SessionSaveError:      http.StatusInternalServerError,
	WrongCredentials:      http.StatusUnauthorized,
	UserNotAuthorised:     http.StatusUnauthorized,
	LogoutError:           http.StatusInternalServerError,
	PasswordsDoNotMatch:   http.StatusBadRequest,
}
