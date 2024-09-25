package handlers

import (
	"encoding/json"
	"github.com/gorilla/sessions"
	"net/http"

	"2024_2_kotyari/config"
	"2024_2_kotyari/db"
	"2024_2_kotyari/errs"
)

// validateCredentials проверяет учетные данные пользователя
func validateCredentials(w *http.ResponseWriter, creds credsApiRequest, requireUsername bool) bool {
	switch {
	case !isValidEmail(creds.Email):
		writeJSON(*w, http.StatusBadRequest, errs.HTTPErrorResponse{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: errs.InvalidEmailFormat.Error(),
		})
		return false
	case !isValidPassword(creds.Password):
		writeJSON(*w, http.StatusBadRequest, errs.HTTPErrorResponse{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: errs.InvalidPasswordFormat.Error(),
		})
		return false
	}

	if requireUsername {
		if !isValidUsername(creds.Username) {
			writeJSON(*w, http.StatusBadRequest, errs.HTTPErrorResponse{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: errs.InvalidUsernameFormat.Error(),
			})
			return false
		}
	}

	return true
}

type Server struct {
	sessions sessions.Store
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		sessions: sessions.NewCookieStore([]byte(cfg.SessionKey)),
	}
}

// LoginHandler обрабатывает вход пользователя и создает сессию
// @Summary Логин пользователя
// @Description Проверяет учетные данные пользователя и создает сессию при успешной аутентификации
// @Tags auth
// @Accept json
// @Produce json
// @Param user body db.User true "Данные пользователя"
// @Success 200 {string} string "Успешный вход"
// @Failure 400 {string} string "Неправильный запрос"
// @Failure 401 {string} string "Неправильные учетные данные"
// @Failure 500 {string} string "Ошибка при создании сессии"
// @Router /login [post]
func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds credsApiRequest
	session, err := s.sessions.Get(r, config.GetSessionName())
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, errs.HTTPErrorResponse{
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: errs.SessionCreationError.Error(),
		})
		return
	}
	if email, isAuthenticated := session.Values["user_id"].(string); isAuthenticated {
		user, _ := db.GetUserByEmail(email)
		writeJSON(w, http.StatusOK, struct {
			Code     int
			Username string
		}{
			Code:     http.StatusOK,
			Username: user.Username,
		})
		return
	}

	err = json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, errs.HTTPErrorResponse{
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: errs.InternalServerError.Error(),
		})
		return
	}

	if !validateCredentials(&w, creds, false) {
		return
	}

	user, exists := db.GetUserByEmail(creds.Email)
	if !exists || user.Password != creds.Password {
		writeJSON(w, http.StatusUnauthorized, errs.HTTPErrorResponse{
			ErrorCode:    http.StatusUnauthorized,
			ErrorMessage: errs.UnauthorizedCredentials.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, struct {
		Code     int
		Username string
	}{
		Code:     http.StatusOK,
		Username: user.Username,
	})
}

// LogoutHandler очищает куки и завершает сессию
// @Summary Логаут пользователя
// @Description Завершает сессию пользователя, очищая куки и удаляя все значения из сессии
// @Tags auth
// @Produce json
// @Success 200 {string} string "Вы успешно вышли"
// @Failure 401 {string} string "Пользователь не авторизован"
// @Failure 500 {string} string "Ошибка при завершении сессии"
// @Router /logout [post]
func (s *Server) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем сессию из запроса
	session, err := s.sessions.Get(r, config.GetSessionName())
	if err != nil {
		http.Error(w, errs.UnauthorizedMessage.Error(), http.StatusUnauthorized)
		return
	}
	// Очищаем значения сессии, создавая новую пустую мапу
	session.Values = make(map[interface{}]interface{})

	// Устанавливаем время жизни сессии в -1, что означает, что сессия будет завершена
	session.Options.MaxAge = -1

	// Сохраняем изменения сессии
	err = s.sessions.Save(r, w, session)
	if err != nil {
		http.Error(w, errs.LogoutError.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusNoContent, nil)
}
