package handlers

import (
	"encoding/json"
	"net/http"

	"2024_2_kotyari/config"
	"2024_2_kotyari/db"
	"2024_2_kotyari/sessions"
)

// respondWithError отправляет JSON ответ с ошибкой и устанавливает код статуса HTTP
func respondWithError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code) // Устанавливаем HTTP статус-код
	writeJSON(w, HTTPErrorResponse{
		ErrorCode:    code,
		ErrorMessage: message,
	})
}

// validateCredentials проверяет учетные данные пользователя
func validateCredentials(w http.ResponseWriter, creds db.User) bool {
	if !isValidEmail(creds.Email) {
		respondWithError(w, http.StatusBadRequest, ErrInvalidEmailFormat.Error())
		return false
	}
	if !isValidPassword(creds.Password) {
		respondWithError(w, http.StatusBadRequest, ErrInvalidPasswordFormat.Error())
		return false
	}

	return true
}

// LoginHandler обрабатывает запросы на вход
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, ErrMethodNotAllowed.Error())
		return
	}

	var creds db.User
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, ErrInvalidJSONFormat.Error())
		return
	}

	if !validateCredentials(w, creds) {
		return
	}

	user, exists := db.GetUserByEmail(creds.Email)
	if !exists || user.Password != creds.Password {
		respondWithError(w, http.StatusUnauthorized, ErrUnauthorizedCredentials.Error())
		return
	}

	session, err := sessions.GetSession(r, config.SessionName)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, ErrSessionCreationError.Error())
		return
	}

	session.Values["user_id"] = creds.Email
	err = sessions.SaveSession(r, w, session)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, ErrSessionSaveError.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Успешный вход"))
}

// LogoutHandler очищает куки и завершает сессию
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем сессию из запроса
	session, err := sessions.GetSession(r, config.SessionName)
	if err != nil {
		http.Error(w, ErrSessionRetrievalError.Error(), http.StatusInternalServerError)
		return
	}
	// Очищаем значения сессии, создавая новую пустую мапу
	session.Values = make(map[interface{}]interface{})

	// Устанавливаем время жизни сессии в -1, что означает, что сессия будет завершена
	session.Options.MaxAge = -1

	// Сохраняем изменения сессии
	err = sessions.SaveSession(r, w, session)
	if err != nil {
		http.Error(w, ErrLogoutError.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Вы вышли из системы"))
}

// writeJSON отправляет JSON ответ
func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, ErrInternalServerError.Error(), http.StatusInternalServerError)
	}

}
