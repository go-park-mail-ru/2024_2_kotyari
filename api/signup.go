package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/mail"

	"2024_2_kotyari/db"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, HTTPErrorResponse{
			ErrorCode:    405,
			ErrorMessage: "method not allowed",
		})
		return
	}

	var signupRequest signupApiRequest
	err := json.NewDecoder(r.Body).Decode(&signupRequest)

	if err != nil {
		writeJSON(w, HTTPErrorResponse{
			ErrorCode:    400,
			ErrorMessage: "Invalid JSON format",
		})
		return
	}

	err = sanitizeParams(signupRequest)
	if err != nil {
		writeJSON(w, HTTPErrorResponse{
			ErrorCode:    400,
			ErrorMessage: err.Error(),
		})
		return
	}

	err = db.CreateUser(signupRequest.Email, signupRequest.User)
	if err != nil {
		writeJSON(w, HTTPErrorResponse{
			ErrorCode:    409,
			ErrorMessage: err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK)
}

func writeJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "Internal server error", 500)
	}
}

func sanitizeParams(signupRequest signupApiRequest) error {
	if len(signupRequest.Username) == 0 || len(signupRequest.Email) == 0 || len(signupRequest.Password) == 0 {
		return errors.New("empty params")
	}

	if len(signupRequest.Username) < 3 || len(signupRequest.Username) > 20 {
		return errors.New("wrong username format")
	}

	_, err := mail.ParseAddress(signupRequest.Email)
	if err != nil {
		return errors.New("wrong email format")
	}

	if len(signupRequest.Password) < 5 {
		return errors.New("wrong password format")
	}

	return nil
}
