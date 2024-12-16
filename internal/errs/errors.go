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
	InvalidJSONFormat             = errors.New("Неверный формат JSON")
	InvalidUsernameFormat         = errors.New("Неверный формат имени пользователя")
	InvalidEmailFormat            = errors.New("Неверный формат email")
	InvalidPasswordFormat         = errors.New("Неверный формат пароля")
	PasswordsDoNotMatch           = errors.New("Пароли не совпадают")
	UserAlreadyExists             = errors.New("Пользователь уже существует")
	InternalServerError           = errors.New("Внутренняя ошибка сервера")
	SessionCreationError          = errors.New("Ошибка при создании сессии")
	SessionSaveError              = errors.New("Ошибка при сохранении сессии")
	SessionNotFound               = errors.New("Сессия не найдена")
	WrongCredentials              = errors.New("Неверный email или пароль")
	UserNotAuthorized             = errors.New("Пользователь не авторизован")
	LogoutError                   = errors.New("Ошибка при завершении сессии")
	UserDoesNotExist              = errors.New("Пользователь не существует")
	CartDoesNotExist              = errors.New("Корзины не существует")
	ProductsDoesNotExists         = errors.New("Продукты не добавили =(")
	ProductToModifyNotFoundInCart = errors.New("Продукт не найден в корзине")
	ProductCountTooLow            = errors.New("Товар закончился")
	ProductNotFound               = errors.New("Продукта не существует")
	EmptyCart                     = errors.New("Корзина пуста")
	ParsingURLArg                 = errors.New("Ошибка парсинга аргумента URL")
	BadRequest                    = errors.New("Неправильный запрос")
	ProductNotInCart              = errors.New("Этого продукта нет в корзине")
	ProductAlreadyInCart          = errors.New("Этот продукт уже находится в корзине")
	CategoriesDoesNotExits        = errors.New("Нет категорий")
	OptionsDoesNotExists          = errors.New("Не найдены опции")
	ImagesDoesNotExists           = errors.New("Нет картинок")
	AddressNotFound               = errors.New("Адрес не найден")
	ErrInvalidOrderIDFormat       = errors.New("Неверный формат id")
	ErrInvalidDeliveryDateFormat  = errors.New("Неверный формат даты доставки")
	ErrOrderNotFound              = errors.New("Заказ не найден")
	ErrRetrieveOrder              = errors.New("Не удалось получить заказ")
	ErrGetNearestDeliveryDate     = errors.New("Нет заказов в процесе доставки")
	ErrFileTypeNotAllowed         = errors.New("Тип файла не допустим")
	RequestIDNotFound             = errors.New("Не удалось получить request-id")
	NoSelectedProducts            = errors.New("Не выбрано ни одного продукта")
	AvatarFileReadError           = errors.New("Не удалось прочитать файл, попробуйте позже")
	AvatarFileSizeExceedsLimit    = errors.New("Размер файла превышает 10 МБ")
	AvatarUploadError             = errors.New("Не удалось загрузить фотографию")
	AvatarImageSaveError          = errors.New("Не удалось сохранить изображение")
	NoReviewsForProduct           = errors.New("Для этого продукта нет отзывов")
	ReviewNotFound                = errors.New("Отзыв не найден")
	ReviewAlreadyExists           = errors.New("Отзыв уже существует")
	NoTitlesToSuggest             = errors.New("Отсутствуют продукты для саджестов")
	FailedToChangeProductRating   = errors.New("Не удалось изменить рейтинг продукта")
	NoOrdersUpdates               = errors.New("У этого пользователя нет обновлений заказов")
	NoPromoCodesForUser           = errors.New("У данного пользователя нет промокодов")
	FailedToParseConfig           = errors.New("Ошибка парсинга конфигурации")
	NoPromoCode                   = errors.New("Этого промокода нет")
	FailedToRetrievePromoCode     = errors.New("Не удалось получить промокод")
	FailedToRetrievePromoCodes    = errors.New("Не удалось получить промокоды")
	ErrNotPermitted               = errors.New("Нет прав для изменения")
	WrongPassword                 = errors.New("Неправильный пароль")
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
			FailedToChangeProductRating:   http.StatusInternalServerError,
			ErrNotPermitted:               http.StatusForbidden,
			NoOrdersUpdates:               http.StatusNotFound,
			NoPromoCodesForUser:           http.StatusNotFound,
			FailedToParseConfig:           http.StatusInternalServerError,
			NoPromoCode:                   http.StatusNotFound,
			FailedToRetrievePromoCodes:    http.StatusServiceUnavailable,
			FailedToRetrievePromoCode:     http.StatusServiceUnavailable,
			WrongPassword:                 http.StatusBadRequest,
		},
	}
}

type HTTPErrorResponse struct {
	ErrorMessage string `json:"error_message"`
}
