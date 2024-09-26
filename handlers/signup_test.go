package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"2024_2_kotyari/config"
	"2024_2_kotyari/errs"
	"github.com/joho/godotenv"
)

func TestSignupHandler(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
		return
	}

	config.Init()
	server := NewServer(&config.Cfg)

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
			requestBody:      `{"username":"t","email":"test@test.com", "password":"test@Password1"}`,
			wantStatus:       http.StatusBadRequest,
			wantErrorMessage: errs.InvalidUsernameFormat.Error(),
		},
		{
			name:             "Invalid Email Format",
			requestBody:      `{"username":"testing","email":"test@", "password":"testPassword"}`,
			wantStatus:       http.StatusBadRequest,
			wantErrorMessage: errs.InvalidEmailFormat.Error(),
		},
		{
			name:             "Invalid Password Format",
			requestBody:      `{"username":"testing","email":"test@test.com", "password":"te"}`,
			wantStatus:       http.StatusBadRequest,
			wantErrorMessage: errs.InvalidPasswordFormat.Error(),
		},
		{
			name:        "Valid Signup",
			requestBody: `{"username":"PROSADIK","email":"testwewew@test.com", "password":"abcdefG@23"}`,
			wantStatus:  http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/signup", strings.NewReader(tt.requestBody))
			rr := httptest.NewRecorder()
			server.SignupHandler(rr, req)

			res := rr.Result()
			defer res.Body.Close()

			if res.StatusCode != tt.wantStatus {
				t.Errorf("Expected status code: %v, got: %v", tt.wantStatus, res.StatusCode)
			}

			if tt.wantStatus >= 400 {
				var httpError errs.HTTPErrorResponse
				err = json.NewDecoder(res.Body).Decode(&httpError)
				if err != nil {
					t.Fatal(err)
				}
				if httpError.ErrorCode != tt.wantStatus || httpError.ErrorMessage != tt.wantErrorMessage {
					t.Errorf("Expected error code %v, error message: %s, got error code: %v, error message: %s",
						tt.wantStatus, tt.wantErrorMessage, httpError.ErrorCode, httpError.ErrorMessage)
				}
			}
		})
	}

	t.Run("Concurrent Signups", func(t *testing.T) {
		var wg sync.WaitGroup
		requestStrings := []string{
			`{"username":"abcdefghij","email":"test@test.com", "password":"abcdefG@23"}`,
			`{"username":"abcdefghij","email":"test1@test.com", "password":"abcdefG@23"}`,
			`{"username":"abcdefghij","email":"test2@test.com", "password":"abcdefG@23"}`,
		}
		for _, requestString := range requestStrings {
			wg.Add(1)
			go func(requestString string) {
				defer wg.Done()
				req := httptest.NewRequest("POST", "/signup", strings.NewReader(requestString))
				rr := httptest.NewRecorder()

				server.SignupHandler(rr, req)

				res := rr.Result()
				defer res.Body.Close()

				if res.StatusCode != http.StatusOK {
					t.Errorf("Expected status code: %v, got: %v", http.StatusOK, res.StatusCode)
				}
			}(requestString)
		}
		wg.Wait()
	})
}
