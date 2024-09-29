package errs

import "errors"

type HTTPErrorResponse struct {
	ErrorMessage string `json:"error_message"`
}

var (
	InvalidJSONFormat       = errors.New("неверный формат JSON")
	InvalidUsernameFormat   = errors.New("неверный формат имени пользователя")
	InvalidEmailFormat      = errors.New("неверный формат email")
	InvalidPasswordFormat   = errors.New("неверный формат пароля")
	UserAlreadyExists       = errors.New("пользователь уже существует")
	InternalServerError     = errors.New("внутренняя ошибка сервера")
	SessionCreationError    = errors.New("ошибка при создании сессии")
	UnauthorizedCredentials = errors.New("неверный email или пароль")
	SessionSaveError        = errors.New("ошибка при сохранении сессии")
	LogoutError             = errors.New("ошибка при завершении сессии")
	UnauthorizedMessage     = errors.New("пользователь не авторизован")
	PasswordsDoNotMatch     = errors.New("пароли не совпадают")
)
