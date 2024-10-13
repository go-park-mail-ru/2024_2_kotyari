package handlers

import (
	"encoding/json"

	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/db"
	userD "github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/user"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	userR "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/user"
	userU "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/user"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/gorilla/sessions"
)

type AuthApp struct {
	userDB   db.UserManager
	sessions SessionInterface
}

func NewApp(users db.UserManager, sessions SessionInterface) *AuthApp {
	return &AuthApp{
		sessions: sessions,
		userDB:   users,
	}
}

func newAppForTests() *AuthApp {
	userDB := db.InitUsersWithData()
	return &AuthApp{
		sessions: newTestSessions(),
		userDB:   userDB,
	}
}

// Login обрабатывает вход пользователя и создает сессию
// @Summary Логин пользователя
// @Description Проверяет учетные данные пользователя и создает сессию при успешной аутентификации
// @Tags auth
// @Accept json
// @Produce json
// @Param user body db.User true "Данные пользователя"
// @Success 200 {object} UsernameResponse "Имя пользователя"
// @Failure 400 {string} string "Неправильный запрос"
// @Failure 401 {string} string "Неправильные учетные данные"
// @Failure 500 {string} string "Ошибка при создании сессии"
// @Router /login [post]
func (a *AuthApp) Login(w http.ResponseWriter, r *http.Request) {
	var creds model.UserApiRequest
	userHandler := userD.NewUserDelivery(userU.NewUserUsecase(userR.NewUserMapRepository()))
	session, err := a.sessions.Get(r)

	if err != nil {
		writeJSON(w, http.StatusInternalServerError, errs.HTTPErrorResponse{
			ErrorMessage: errs.SessionCreationError.Error(),
		})

		return
	}

	err = json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, errs.HTTPErrorResponse{
			ErrorMessage: errs.InternalServerError.Error(),
		})
		return
	}

	if !validateLogin(w, creds) {
		return
	}

	user, exists := userHandler.GetUserByEmail(creds)
	if !exists || !utils.VerifyPassword(user.Password, creds.Password) {
		writeJSON(w, http.StatusUnauthorized, errs.HTTPErrorResponse{
			ErrorMessage: errs.UnauthorizedCredentials.Error(),
		})

		return
	}

	session.Values[sessionKey] = creds.Email
	setSessionOptions(session, 10*hour)

	err = a.sessions.Save(w, r, session)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, errs.HTTPErrorResponse{
			ErrorMessage: errs.SessionSaveError.Error(),
		})

		return
	}

	writeJSON(w, http.StatusOK, model.UsernameResponse{Username: user.Username})
}

// Logout завершает сессию пользователя
// @Summary Логаут пользователя
// @Description Завершает сессию пользователя, очищая куки и удаляя все значения из сессии
// @Tags auth
// @Produce json
// @Success 204
// @Failure 401 {string} string "Пользователь не авторизован"
// @Failure 500 {string} string "Ошибка при завершении сессии"
// @Router /logout [post]
func (a *AuthApp) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := a.sessions.Get(r)
	if err != nil {
		http.Error(w, errs.UnauthorizedMessage.Error(), http.StatusUnauthorized)
		return
	}
	// Очищаем значения сессии, создавая новую пустую мапу
	session.Values = make(map[interface{}]interface{})
	setSessionOptions(session, nullTime)

	err = a.sessions.Save(w, r, session)
	if err != nil {
		http.Error(w, errs.LogoutError.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusNoContent, nil)
}

// IsLogin проверяет, авторизован ли пользователь
// @Summary Проверка авторизации пользователя
// @Description Проверяет, авторизован ли пользователь, и возвращает его имя пользователя
// @Tags auth
// @Produce json
// @Success 200 {object} UsernameResponse "Информация о пользователе"
// @Failure 401 {string} string "Пользователь не авторизован"
// @Router /islogin [get]
func (a *AuthApp) IsLogin(w http.ResponseWriter, r *http.Request) {
	userHandler := userD.NewUserDelivery(userU.NewUserUsecase(userR.NewUserMapRepository()))
	session, err := a.sessions.Get(r)

	if err != nil {
		http.Error(w, errs.UnauthorizedMessage.Error(), http.StatusUnauthorized)

		return
	}

	if email, isAuthenticated := session.Values[sessionKey].(string); isAuthenticated {
		user, _ := userHandler.GetUserByEmail(model.UserApiRequest{Email: email})
		writeJSON(w, http.StatusOK, model.UsernameResponse{Username: user.Username})
		return
	}

	http.Error(w, errs.UnauthorizedMessage.Error(), http.StatusUnauthorized)
}
