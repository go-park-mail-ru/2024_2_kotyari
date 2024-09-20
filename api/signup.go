package api

import (
	"2024_2_kotyari/custom_errors"
	"2024_2_kotyari/db"
	"encoding/json"
	"errors"
	"net/http"
	"net/mail"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, custom_errors.ErrorResponse{
			ErrorCode:    405,
			ErrorMessage: "method not allowed",
		})
		return
	}

	var user db.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		writeJSON(w, custom_errors.ErrorResponse{
			ErrorCode:    400,
			ErrorMessage: "Invalid JSON format",
		})
		return
	}

	err = sanitizeParams(user)
	if err != nil {
		writeJSON(w, custom_errors.ErrorResponse{
			ErrorCode:    400,
			ErrorMessage: err.Error(),
		})
		return
	}

	err = db.CreateUser(user)
	if err != nil {
		writeJSON(w, custom_errors.ErrorResponse{
			ErrorCode:    500,
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

func sanitizeParams(user db.User) error {
	if len(user.Username) == 0 || len(user.Email) == 0 || len(user.Password) == 0 {
		return errors.New("empty params")
	}

	if len(user.Username) < 3 || len(user.Username) > 20 {
		return errors.New("wrong username format")
	}

	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return errors.New("wrong email format")
	}

	if len(user.Password) < 5 {
		return errors.New("wrong password format")
	}

	return nil
}
