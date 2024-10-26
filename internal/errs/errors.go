package errs

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type GetErrorCode interface {
	Get(error) (error, int)
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
	SessionNotFound       = errors.New("сессия не найдена")
	WrongCredentials      = errors.New("неверный email или пароль")
	UserNotAuthorized     = errors.New("пользователь не авторизован")
	LogoutError           = errors.New("ошибка при завершении сессии")
	UserDoesNotExist      = errors.New("пользователь не существует")
)

type ErrorStore struct {
	mux        sync.RWMutex
	errorCodes map[error]int
}

func (e *ErrorStore) Get(err error) (error, int) {
	e.mux.RLock()
	defer e.mux.RUnlock()
	errCode, present := e.errorCodes[err]
	if !present {
		log.Println(fmt.Errorf("unexpected error occured: %w", err))
		return InternalServerError, http.StatusInternalServerError
	}

	return err, errCode
}

func NewErrorStore() *ErrorStore {
	return &ErrorStore{
		mux: sync.RWMutex{},
		errorCodes: map[error]int{
			InvalidJSONFormat:     http.StatusBadRequest,
			InvalidUsernameFormat: http.StatusBadRequest,
			InvalidEmailFormat:    http.StatusBadRequest,
			InvalidPasswordFormat: http.StatusBadRequest,
			UserAlreadyExists:     http.StatusConflict,
			InternalServerError:   http.StatusInternalServerError,
			SessionCreationError:  http.StatusInternalServerError,
			SessionSaveError:      http.StatusInternalServerError,
			SessionNotFound:       http.StatusUnauthorized,
			WrongCredentials:      http.StatusUnauthorized,
			UserNotAuthorized:     http.StatusUnauthorized,
			LogoutError:           http.StatusInternalServerError,
			PasswordsDoNotMatch:   http.StatusBadRequest,
			UserDoesNotExist:      http.StatusNotFound,
		},
	}
}

type HTTPErrorResponse struct {
	ErrorMessage string `json:"error_message"`
}
