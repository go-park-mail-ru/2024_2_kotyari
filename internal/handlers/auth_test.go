package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testUser = credsApiRequest{
	Email:    "user@example.com",
	Password: "Password123@",
}

var testUserIncorrectEmail = credsApiRequest{
	Email:    "user1example.ru",
	Username: "Goshanchik",
	Password: "Password123@",
}

var testUserIncorrectPass = credsApiRequest{
	Email:    "user1@example.ru",
	Username: "Goshanchik",
	Password: "password123",
}

var testUserNotFound = credsApiRequest{
	Email:    "notfound@example.com",
	Username: "Goshanchik",
	Password: "Password123@",
}

func TestLoginHandler(t *testing.T) {
	server := newAppForTests()

	tests := []struct {
		name       string
		method     string
		body       credsApiRequest
		wantStatus int
	}{
		{
			name:       "Валидный логин",
			method:     http.MethodPost,
			body:       testUser,
			wantStatus: http.StatusOK,
		},
		{
			name:       "Некорректный формат почты",
			method:     http.MethodPost,
			body:       testUserIncorrectEmail,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "Некорректный формат пароля",
			method:     http.MethodPost,
			body:       testUserIncorrectPass,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "Пользователя не существует",
			method:     http.MethodPost,
			body:       testUserNotFound,
			wantStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.body)
			req := httptest.NewRequest(tt.method, "/login", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			server.Login(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("Expected status %v, got %v", tt.wantStatus, w.Code)
			}
		})
	}
}

// Тест для LogoutHandler
func TestLogoutHandler(t *testing.T) {
	server := newAppForTests()

	req := httptest.NewRequest(http.MethodGet, "/logout", nil)
	req.AddCookie(&http.Cookie{Name: "session_id", Value: "test-session-id"}) // Имитация сессии
	w := httptest.NewRecorder()
	server.Logout(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("Ожидалось 200, имеем %v", w.Code)
	}
}
