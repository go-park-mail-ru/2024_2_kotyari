package handlers

import "errors"

type HTTPErrorResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

var (
	ErrInvalidJSONFormat       = errors.New("неверный формат JSON")
	ErrMethodNotAllowed        = errors.New("метод не разрешен")
	ErrInvalidEmailFormat      = errors.New("неверный формат email")
	ErrInvalidPasswordFormat   = errors.New("пароль должен содержать минимум 8 символов, одну цифру и одну заглавную букву")
	ErrUnauthorizedCredentials = errors.New("неправильная почта или пароль")
	ErrSessionCreationError    = errors.New("ошибка при создании сессии")
	ErrSessionSaveError        = errors.New("ошибка при сохранении сессии")
	ErrSessionRetrievalError   = errors.New("ошибка при получении сессии")
	ErrLogoutError             = errors.New("ошибка при завершении сессии")
	ErrInternalServerError     = errors.New("внутренняя ошибка сервера")
	ErrUnauthorizedMessage     = "Не авторизован"
)
