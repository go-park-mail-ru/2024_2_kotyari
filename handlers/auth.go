package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"2024_2_kotyari/config"
	"2024_2_kotyari/db"
	"github.com/gorilla/sessions"
)

// respondWithError отправляет JSON ответ с ошибкой и устанавливает код статуса HTTP
func respondWithError(w *http.ResponseWriter, code int, message string) {
	(*w).WriteHeader(code) // Устанавливаем HTTP статус-код
	writeJSON(*w, HTTPErrorResponse{
		ErrorCode:    code,
		ErrorMessage: message,
	})
}

// validateCredentials проверяет учетные данные пользователя
func validateCredentials(w *http.ResponseWriter, creds loginApiRequest) bool {
	switch {
	case !isValidEmail(creds.Email):
		respondWithError(w, http.StatusBadRequest, ErrInvalidEmailFormat.Error())
		return false
	case !isValidPassword(creds.Password):
		respondWithError(w, http.StatusBadRequest, ErrInvalidPasswordFormat.Error())
		return false
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
	var creds loginApiRequest
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		respondWithError(&w, http.StatusInternalServerError, ErrInternalServerError.Error())
		return
	}

	if !validateCredentials(&w, creds) {
		return
	}

	user, exists := db.GetUserByEmail(creds.Email)
	if !exists || user.Password != creds.Password {
		respondWithError(&w, http.StatusUnauthorized, ErrUnauthorizedCredentials.Error())
		return
	}

	session, err := s.sessions.Get(r, config.GetSessionName())
	if err != nil {
		respondWithError(&w, http.StatusInternalServerError, ErrSessionCreationError.Error())
		return
	}

	session.Values["user_id"] = creds.Email
	err = s.sessions.Save(r, w, session)
	if err != nil {
		log.Printf("error saving session: %v", err)
		log.Println("sessions: ", session)
		respondWithError(&w, http.StatusInternalServerError, ErrSessionSaveError.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
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
		http.Error(w, ErrUnauthorizedMessage.Error(), http.StatusUnauthorized)
		return
	}
	// Очищаем значения сессии, создавая новую пустую мапу
	session.Values = make(map[interface{}]interface{})

	// Устанавливаем время жизни сессии в -1, что означает, что сессия будет завершена
	session.Options.MaxAge = -1

	// Сохраняем изменения сессии
	err = s.sessions.Save(r, w, session)
	if err != nil {
		http.Error(w, ErrLogoutError.Error(), http.StatusInternalServerError)
		return
	}
}

// writeJSON отправляет JSON ответ
func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, ErrInternalServerError.Error(), http.StatusInternalServerError)
	}

}
