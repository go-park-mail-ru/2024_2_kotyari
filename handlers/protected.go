package handlers

import (
	"fmt"
	"net/http"

	"2024_2_kotyari/config"
	"2024_2_kotyari/sessions"
)

// ProtectedHandler — пример защищенного маршрута
func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	session, err := sessions.GetSession(r, config.SessionName)
	if err != nil {
		writeJSON(w, HTTPErrorResponse{
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: ErrSessionRetrievalError.Error(),
		})
		return
	}

	userID, ok := session.Values["user_id"].(string)
	if !ok || userID == "" {
		writeJSON(w, HTTPErrorResponse{
			ErrorCode:    http.StatusUnauthorized,
			ErrorMessage: ErrUnauthorizedMessage,
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Добро пожаловать, %s!", userID)))
}
