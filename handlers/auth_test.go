package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"2024_2_kotyari/db"
)

var testUser = db.User{
	Email:    "user@example.com",
	Password: "Password123",
}

var testUserIncorrectEmail = db.User{
	Email:    "invalid-email",
	Password: "Password123",
}

func TestLoginHandler(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		body       db.User
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
			body:       db.User{Email: "test@example.com", Password: "short"},
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "Пользователя не существует",
			method:     http.MethodPost,
			body:       db.User{Email: "notfound@example.com", Password: "Password123"},
			wantStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.body)
			req := httptest.NewRequest(tt.method, "/login", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			LoginHandler(w, req)

			res := w.Result()
			defer res.Body.Close()

			// для дальнейшей отладки
			bodyBytes, _ := io.ReadAll(res.Body)
			bodyString := string(bodyBytes)

			// Проверка кода ответа
			if res.StatusCode != tt.wantStatus {
				t.Errorf("Ожидалось %v, получено %v. Ответ сервера: %v", tt.wantStatus, res.StatusCode, bodyString)
			}
		})
	}
}

// Тест для LogoutHandler
func TestLogoutHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/logout", nil)
	req.AddCookie(&http.Cookie{Name: "session_id", Value: "test-session-id"}) // Имитация сессии
	w := httptest.NewRecorder()

	LogoutHandler(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("Ожидалось 200, имеем %v", res.StatusCode)
	}
}

// Тест для ProtectedHandler
func TestProtectedHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.AddCookie(&http.Cookie{Name: "session_id", Value: "test-session-id"}) // Имитация сессии
	w := httptest.NewRecorder()

	ProtectedHandler(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("Ожидалось 200, имеем %v", res.StatusCode)
	}
}
