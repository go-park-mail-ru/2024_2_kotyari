package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

func TestSignupHandler(t *testing.T) {
	t.Run("Test for empty request", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/signup", nil)
		rr := httptest.NewRecorder()
		SignupHandler(rr, req)

		res := rr.Result()
		var httpError HTTPErrorResponse
		err := json.NewDecoder(res.Body).Decode(&httpError)
		if err != nil {
			t.Fatal(err)
		}

		if httpError.ErrorCode != http.StatusBadRequest || httpError.ErrorMessage != ErrInvalidJSONFormat.Error() {
			t.Errorf("Expected error code %v , error message: %s, got error code: %v, error message: %s",
				http.StatusBadRequest, ErrInvalidJSONFormat.Error(), httpError.ErrorCode, httpError.ErrorMessage)
		}
	})

	t.Run("Test for wrong username format", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/signup",
			strings.NewReader(`{"username":"t","email":"test@test.com", "password":"testPassword"}`))
		rr := httptest.NewRecorder()
		SignupHandler(rr, req)

		res := rr.Result()
		var httpError HTTPErrorResponse
		err := json.NewDecoder(res.Body).Decode(&httpError)
		if err != nil {
			t.Fatal(err)
		}

		if httpError.ErrorCode != http.StatusBadRequest || httpError.ErrorMessage != ErrWrongUsernameFormat.Error() {
			t.Errorf("Expected error code %v , error message: %s, got error code: %v, error message: %s",
				http.StatusBadRequest, ErrWrongUsernameFormat.Error(), httpError.ErrorCode, httpError.ErrorMessage)
		}
	})

	t.Run("Test for wrong email format", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/signup",
			strings.NewReader(`{"username":"testing","email":"test@", "password":"testPassword"}`))
		rr := httptest.NewRecorder()
		SignupHandler(rr, req)

		res := rr.Result()
		var httpError HTTPErrorResponse
		err := json.NewDecoder(res.Body).Decode(&httpError)
		if err != nil {
			t.Fatal(err)
		}

		if httpError.ErrorCode != http.StatusBadRequest || httpError.ErrorMessage != ErrWrongEmailFormat.Error() {
			t.Errorf("Expected error code %v , error message: %s, got error code: %v, error message: %s",
				http.StatusBadRequest, ErrWrongEmailFormat.Error(), httpError.ErrorCode, httpError.ErrorMessage)
		}
	})

	t.Run("Test for wrong password format", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/signup",
			strings.NewReader(`{"username":"testing","email":"test@test.com", "password":"te"}`))
		rr := httptest.NewRecorder()
		SignupHandler(rr, req)

		res := rr.Result()
		var httpError HTTPErrorResponse
		err := json.NewDecoder(res.Body).Decode(&httpError)
		if err != nil {
			t.Fatal(err)
		}

		if httpError.ErrorCode != http.StatusBadRequest || httpError.ErrorMessage != ErrWrongPasswordFormat.Error() {
			t.Errorf("Expected error code %v , error message: %s, got error code: %v, error message: %s",
				http.StatusBadRequest, ErrWrongPasswordFormat.Error(), httpError.ErrorCode, httpError.ErrorMessage)
		}
	})

	t.Run("Test for empty params", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/signup",
			strings.NewReader(`{"username":"","email":"test@test.com", "password":"testtest"}`))
		rr := httptest.NewRecorder()
		SignupHandler(rr, req)

		res := rr.Result()
		var httpError HTTPErrorResponse
		err := json.NewDecoder(res.Body).Decode(&httpError)
		if err != nil {
			t.Fatal(err)
		}

		if httpError.ErrorCode != http.StatusBadRequest || httpError.ErrorMessage != ErrEmptyRequestParams.Error() {
			t.Errorf("Expected error code %v , error message: %s, got error code: %v, error message: %s",
				http.StatusBadRequest, ErrEmptyRequestParams.Error(), httpError.ErrorCode, httpError.ErrorMessage)
		}
	})

	t.Run("Test for OK", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/signup",
			strings.NewReader(`{"username":"testing","email":"test@test.com", "password":"testtest"}`))
		rr := httptest.NewRecorder()
		SignupHandler(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expexted status code: %v, got: %v", http.StatusOK, rr.Code)
		}
	})

	t.Run("Test concurrent signups", func(t *testing.T) {
		var wg sync.WaitGroup
		requestStrings := []string{
			`{"username":"testing","email":"test@test.com", "password":"testtest"}`,
			`{"username":"testing","email":"test1@test.com", "password":"testtest"}`,
			`{"username":"testing","email":"test2@test.com", "password":"testtest"}`,
		}
		for _, requestString := range requestStrings {
			wg.Add(1)
			go func(requestString string) {
				defer wg.Done()
				req := httptest.NewRequest("POST", "/signup", strings.NewReader(requestString))
				rr := httptest.NewRecorder()
				SignupHandler(rr, req)

				if rr.Code != http.StatusOK {
					t.Errorf("Expexted status code: %v, got: %v", http.StatusOK, rr.Code)
				}
			}(requestString)
		}
		wg.Wait()
	})
}
