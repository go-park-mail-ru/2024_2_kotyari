package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

var testUser = model.UserLoginRequestDTO{
	Email:    "user@example.com",
	Password: "Password123@",
}

var testUserIncorrectEmail = model.UserLoginRequestDTO{
	Email:    "user1example.ru",
	Password: "Password123@",
}

var testUserIncorrectPass = model.UserLoginRequestDTO{
	Email:    "user1@example.ru",
	Password: "password123",
}

var testUserNotFound = model.UserLoginRequestDTO{
	Email:    "notfound@example.com",
	Password: "Password123@",
}

func TestLoginHandler(t *testing.T) {
	testsAuthManager := newTestsAuthManager()

	tests := []struct {
		name       string
		method     string
		body       model.UserLoginRequestDTO
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

			testsAuthManager.Login(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("Expected status %v, got %v", tt.wantStatus, w.Code)
			}
		})
	}
}

// Тест для LogoutHandler
func TestLogoutHandler(t *testing.T) {
	testsAuthManager := newTestsAuthManager()

	req := httptest.NewRequest(http.MethodPost, "/logout", nil)
	req.AddCookie(&http.Cookie{Name: "session_id", Value: "test-session-id"}) // Имитация сессии
	w := httptest.NewRecorder()

	testsAuthManager.Logout(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("Expected 204 No Content, got %v", w.Code)
	}
}

func TestSignupHandler(t *testing.T) {
	tests := []struct {
		name             string
		requestBody      string
		wantStatus       int
		wantErrorMessage string
	}{
		{
			name:             "Empty Request",
			requestBody:      "",
			wantStatus:       http.StatusBadRequest,
			wantErrorMessage: errs.InvalidJSONFormat.Error(),
		},
		{
			name:             "Invalid Username Format",
			requestBody:      `{"username":"t","email":"test@test.com", "password":"test@Password1", "repeat_password":"test@Password1"}`,
			wantStatus:       http.StatusBadRequest,
			wantErrorMessage: errs.InvalidUsernameFormat.Error(),
		},
		{
			name:             "Invalid Email Format",
			requestBody:      `{"username":"testing","email":"test@", "password":"testPassword", "repeat_password":"testPassword"}`,
			wantStatus:       http.StatusBadRequest,
			wantErrorMessage: errs.InvalidEmailFormat.Error(),
		},
		{
			name:             "Invalid Password Format",
			requestBody:      `{"username":"testing","email":"test@test.com", "password":"te", "repeat_password":"te"}`,
			wantStatus:       http.StatusBadRequest,
			wantErrorMessage: errs.InvalidPasswordFormat.Error(),
		},
		{
			name:             "Passwords Do Not Match",
			requestBody:      `{"username":"testing","email":"test@test.com", "password":"Password1@","repeat_password":"Password2@"}`,
			wantStatus:       http.StatusBadRequest,
			wantErrorMessage: errs.PasswordsDoNotMatch.Error(),
		},
		{
			name:        "Valid SignUp",
			requestBody: `{"username":"PROSADIK","email":"testwewew@test.com", "password":"abcdefG@23", "repeat_password":"abcdefG@23"}`,
			wantStatus:  http.StatusOK,
		},
	}

	for _, tt := range tests {
		testsAuthManager := newTestsAuthManager()

		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/signup", strings.NewReader(tt.requestBody))
			rr := httptest.NewRecorder()
			testsAuthManager.SignUp(rr, req)

			if rr.Code != tt.wantStatus {
				t.Errorf("Expected status code: %v, got: %v", tt.wantStatus, rr.Code)
			}

			if tt.wantStatus >= 400 {
				var httpError errs.HTTPErrorResponse
				err := json.NewDecoder(rr.Body).Decode(&httpError)
				if err != nil {
					t.Fatal(err)
				}

				if httpError.ErrorMessage != tt.wantErrorMessage {
					t.Errorf("Expected error message: %s, got: %s", tt.wantErrorMessage, httpError.ErrorMessage)
				}
			}
		})
	}

	t.Run("Concurrent Signups", func(t *testing.T) {
		var wg sync.WaitGroup

		testsAuthManager := newTestsAuthManager()

		requestStrings := []string{
			`{"username":"abcdefghij","email":"test@test.com", "password":"abcdefG@23", "repeat_password":"abcdefG@23"}`,
			`{"username":"abcdefghij","email":"test1@test.com", "password":"abcdefG@23", "repeat_password":"abcdefG@23"}`,
			`{"username":"abcdefghij","email":"test2@test.com", "password":"abcdefG@23", "repeat_password":"abcdefG@23"}`,
		}
		for _, requestString := range requestStrings {
			wg.Add(1)

			go func(requestString string) {
				defer wg.Done()
				req := httptest.NewRequest("POST", "/signup", strings.NewReader(requestString))
				rr := httptest.NewRecorder()

				testsAuthManager.SignUp(rr, req)

				if rr.Code != http.StatusOK {
					t.Errorf("Expected status code: %v, got: %v", http.StatusOK, rr.Code)
				}
			}(requestString)
		}
		wg.Wait()
	})
}
