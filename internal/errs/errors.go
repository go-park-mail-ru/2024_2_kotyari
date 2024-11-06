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
	InvalidJSONFormat             = errors.New("неверный формат JSON")
	InvalidUsernameFormat         = errors.New("неверный формат имени пользователя")
	InvalidEmailFormat            = errors.New("неверный формат email")
	InvalidPasswordFormat         = errors.New("неверный формат пароля")
	PasswordsDoNotMatch           = errors.New("пароли не совпадают")
	UserAlreadyExists             = errors.New("пользователь уже существует")
	InternalServerError           = errors.New("внутренняя ошибка сервера")
	SessionCreationError          = errors.New("ошибка при создании сессии")
	SessionSaveError              = errors.New("ошибка при сохранении сессии")
	SessionNotFound               = errors.New("сессия не найдена")
	WrongCredentials              = errors.New("неверный email или пароль")
	UserNotAuthorized             = errors.New("пользователь не авторизован")
	LogoutError                   = errors.New("ошибка при завершении сессии")
	UserDoesNotExist              = errors.New("пользователь не существует")
	CartDoesNotExist              = errors.New("корзины не существует")
	ProductsDoesNotExists         = errors.New("продукты не добавили =(")
	ProductToModifyNotFoundInCart = errors.New("продукт не найден в корзине")
	ProductCountTooLow            = errors.New("товар закончился")
	ProductNotFound               = errors.New("продукта не существует")
	EmptyCart                     = errors.New("корзина пуста")
	ParsingURLArg                 = errors.New("ошибка парсинга аргумента URL")
	BadRequest                    = errors.New("неправильный запрос")
	ProductNotInCart              = errors.New("этого продукта нет в корзине")
	ProductAlreadyInCart          = errors.New("этот продукт уже находится в корзине")
	ProductAlreadyRemoved         = errors.New("продукт уже удален")
	CategoriesDoesNotExits        = errors.New("нет категорий")
	OptionsDoesNotExists          = errors.New("не найдены опции")
	ImagesDoesNotExists           = errors.New("нет картинок")
	AddressNotFound               = errors.New("адрес не найден")
	ErrInvalidOrderIDFormat       = errors.New("неверный формат id")
	ErrInvalidDeliveryDateFormat  = errors.New("неферный формат даты доставки")
	ErrOrderNotFound              = errors.New("заказ не найден")
	ErrRetrieveOrder              = errors.New("не удалось получить заказ")
	ErrGetNearestDeliveryDate     = errors.New("нет заказов в процесе доставки")
	ErrFileTypeNotAllowed         = errors.New("тип файла не допустим")
	RequestIDNotFound             = errors.New("не удалось получить request-id")
	NoSelectedProducts            = errors.New("не выбрано ни одного продукта")
	NoReviewsForProduct           = errors.New("для этого продукта нет отзывов")
	ReviewNotFound                = errors.New("отзыв не найден")
	ReviewAlreadyExists           = errors.New("отзыв уже существует")
	NoTitlesToSuggest             = errors.New("отсутствуют продукты для саджестов")
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
			InvalidJSONFormat:             http.StatusBadRequest,
			InvalidUsernameFormat:         http.StatusBadRequest,
			InvalidEmailFormat:            http.StatusBadRequest,
			InvalidPasswordFormat:         http.StatusBadRequest,
			UserAlreadyExists:             http.StatusConflict,
			InternalServerError:           http.StatusInternalServerError,
			SessionCreationError:          http.StatusInternalServerError,
			SessionSaveError:              http.StatusInternalServerError,
			SessionNotFound:               http.StatusUnauthorized,
			WrongCredentials:              http.StatusUnauthorized,
			UserNotAuthorized:             http.StatusUnauthorized,
			LogoutError:                   http.StatusInternalServerError,
			PasswordsDoNotMatch:           http.StatusBadRequest,
			UserDoesNotExist:              http.StatusNotFound,
			ProductsDoesNotExists:         http.StatusNotFound,
			CategoriesDoesNotExits:        http.StatusNotFound,
			OptionsDoesNotExists:          http.StatusNotFound,
			ImagesDoesNotExists:           http.StatusNotFound,
			ProductToModifyNotFoundInCart: http.StatusNotFound,
			ProductCountTooLow:            http.StatusConflict,
			ProductNotFound:               http.StatusNotFound,
			EmptyCart:                     http.StatusNotFound,
			ParsingURLArg:                 http.StatusBadRequest,
			BadRequest:                    http.StatusBadRequest,
			ProductNotInCart:              http.StatusNotFound,
			ProductAlreadyInCart:          http.StatusConflict,
			ErrFileTypeNotAllowed:         http.StatusBadRequest,
			ErrInvalidOrderIDFormat:       http.StatusBadRequest,
			ErrInvalidDeliveryDateFormat:  http.StatusBadRequest,
			ErrOrderNotFound:              http.StatusNotFound,
			ErrRetrieveOrder:              http.StatusNotFound,
			ErrGetNearestDeliveryDate:     http.StatusNotFound,
			RequestIDNotFound:             http.StatusBadRequest,
			NoSelectedProducts:            http.StatusBadRequest,
			NoReviewsForProduct:           http.StatusNotFound,
			ReviewNotFound:                http.StatusNotFound,
			ReviewAlreadyExists:           http.StatusConflict,
			NoTitlesToSuggest:             http.StatusNotFound,
			ProductAlreadyRemoved:         http.StatusBadRequest,
		},
	}
}

type HTTPErrorResponse struct {
	ErrorMessage string `json:"error_message"`
}
