package handlers

import (
	"encoding/json"
	"github.com/gorilla/sessions"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/db"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
)

type AuthApp struct {
	db       *db.Users
	sessions *sessionManager
}

func NewApp() *AuthApp {
	return &AuthApp{
		sessions: newSessions(),
		db:       db.InitUsersWithData(),
	}
}

type Server struct {
	sessions sessions.Store
}

func newAppForTests() *AuthApp {
	return &AuthApp{
		sessions: newTestSessions(),
		db:       db.InitUsersWithData(),
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
	var creds credsApiRequest

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

	user, exists := a.db.GetUserByEmail(creds.Email)
	if !exists || !verifyPassword(user.Password, creds.Password) {
		writeJSON(w, http.StatusUnauthorized, errs.HTTPErrorResponse{
			ErrorMessage: errs.UnauthorizedCredentials.Error(),
		})
		return
	}

	session.Values["user_id"] = creds.Email
	session.Options.MaxAge = 3600 * 10
	session.Options.HttpOnly = true
	err = a.sessions.Save(w, r, session)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, errs.HTTPErrorResponse{
			ErrorMessage: errs.SessionSaveError.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, UsernameResponse{Username: user.Username})
}

// Logout очищает куки и завершает сессию
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

	// Устанавливаем время жизни сессии в -1, что означает, что сессия будет завершена
	session.Options.MaxAge = -1
	err = a.sessions.Save(w, r, session)
	if err != nil {
		http.Error(w, errs.LogoutError.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusNoContent, nil)
}

// IsLogin проверяет, авторизован ли пользователь, и возвращает его имя пользователя
// @Summary Проверка авторизации пользователя
// @Description Проверяет, авторизован ли пользователь, и возвращает его имя пользователя
// @Tags auth
// @Produce json
// @Success 200 {object} UsernameResponse "Информация о пользователе"
// @Failure 401 {string} string "Пользователь не авторизован"
// @Router /islogin [get]
func (a *AuthApp) IsLogin(w http.ResponseWriter, r *http.Request) {
	session, err := a.sessions.Get(r)
	if err != nil {
		http.Error(w, errs.UnauthorizedMessage.Error(), http.StatusUnauthorized)
		return
	}

	if email, isAuthenticated := session.Values["user_id"].(string); isAuthenticated {
		user, _ := a.db.GetUserByEmail(email)
		writeJSON(w, http.StatusOK, UsernameResponse{Username: user.Username})
		return
	}

	http.Error(w, errs.UnauthorizedMessage.Error(), http.StatusUnauthorized)
}
