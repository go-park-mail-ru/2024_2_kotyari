package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"2024_2_kotyari/errs"
)

func TestSignupHandler(t *testing.T) {
	t.Run("Test for empty request", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/signup", nil)
		rr := httptest.NewRecorder()
		SignupHandler(rr, req)

		res := rr.Result()
		var httpError errs.HTTPErrorResponse
		err := json.NewDecoder(res.Body).Decode(&httpError)
		if err != nil {
			t.Fatal(err)
		}

		if httpError.ErrorCode != http.StatusBadRequest || httpError.ErrorMessage != errs.InvalidJSONFormat.Error() {
			t.Errorf("Expected error code %v , error message: %s, got error code: %v, error message: %s",
				http.StatusBadRequest, errs.InvalidJSONFormat.Error(), httpError.ErrorCode, httpError.ErrorMessage)
		}
	})

	t.Run("Test for wrong username format", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/signup",
			strings.NewReader(`{"username":"t","email":"test@test.com", "password":"test@Password1"}`))
		rr := httptest.NewRecorder()
		SignupHandler(rr, req)

		res := rr.Result()

		var httpError errs.HTTPErrorResponse
		err := json.NewDecoder(res.Body).Decode(&httpError)
		if err != nil {
			t.Fatal(err)
		}

		if httpError.ErrorCode != http.StatusBadRequest || httpError.ErrorMessage != errs.InvalidUsernameFormat.Error() {
			t.Errorf("Expected error code %v , error message: %s, got error code: %v, error message: %s",
				http.StatusBadRequest, errs.InvalidUsernameFormat.Error(), httpError.ErrorCode, httpError.ErrorMessage)
		}
	})

	t.Run("Test for wrong email format", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/signup",
			strings.NewReader(`{"username":"testing","email":"test@", "password":"testPassword"}`))
		rr := httptest.NewRecorder()
		SignupHandler(rr, req)

		res := rr.Result()
		var httpError errs.HTTPErrorResponse
		err := json.NewDecoder(res.Body).Decode(&httpError)
		if err != nil {
			t.Fatal(err)
		}

		if httpError.ErrorCode != http.StatusBadRequest || httpError.ErrorMessage != errs.InvalidEmailFormat.Error() {
			t.Errorf("Expected error code %v , error message: %s, got error code: %v, error message: %s",
				http.StatusBadRequest, errs.InvalidEmailFormat.Error(), httpError.ErrorCode, httpError.ErrorMessage)
		}
	})

	t.Run("Test for wrong password format", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/signup",
			strings.NewReader(`{"username":"testing","email":"test@test.com", "password":"te"}`))
		rr := httptest.NewRecorder()
		SignupHandler(rr, req)

		res := rr.Result()
		var httpError errs.HTTPErrorResponse
		err := json.NewDecoder(res.Body).Decode(&httpError)
		if err != nil {
			t.Fatal(err)
		}

		if httpError.ErrorCode != http.StatusBadRequest || httpError.ErrorMessage != errs.InvalidPasswordFormat.Error() {
			t.Errorf("Expected error code %v , error message: %s, got error code: %v, error message: %s",
				http.StatusBadRequest, errs.InvalidPasswordFormat.Error(), httpError.ErrorCode, httpError.ErrorMessage)
		}
	})

	//t.Run("Test for empty params", func(t *testing.T) {
	//	req := httptest.NewRequest("POST", "/signup",
	//		strings.NewReader(`{"username":"","email":"test@test.com", "password":"test@Test1"}`))
	//	rr := httptest.NewRecorder()
	//	SignupHandler(rr, req)
	//
	//	res := rr.Result()
	//	var httpError errs.HTTPErrorResponse
	//	err := json.NewDecoder(res.Body).Decode(&httpError)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//
	//	if httpError.ErrorCode != http.StatusBadRequest || httpError.ErrorMessage != errs.EmptyRequestParams.Error() {
	//		t.Errorf("Expected error code %v , error message: %s, got error code: %v, error message: %s",
	//			http.StatusBadRequest, errs.EmptyRequestParams.Error(), httpError.ErrorCode, httpError.ErrorMessage)
	//	}
	//})

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
