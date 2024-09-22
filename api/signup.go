package api

import (
	"encoding/json"
	"net/http"
	"net/mail"

	"2024_2_kotyari/db"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var signupRequest signupApiRequest

	err := json.NewDecoder(r.Body).Decode(&signupRequest)
	if err != nil {
		writeJSON(w, HTTPErrorResponse{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: ErrInvalidJSONFormat.Error(),
		})
		return
	}

	err = validateParams(signupRequest)
	if err != nil {
		writeJSON(w, HTTPErrorResponse{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: err.Error(),
		})
		return
	}

	err = db.CreateUser(signupRequest.Email, signupRequest.User)
	if err != nil {
		writeJSON(w, HTTPErrorResponse{
			ErrorCode:    http.StatusConflict,
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

func validateParams(signupRequest signupApiRequest) error {
	if len(signupRequest.Username) == 0 || len(signupRequest.Email) == 0 || len(signupRequest.Password) == 0 {
		return ErrEmptyRequestParams
	}

	if len(signupRequest.Username) < 3 || len(signupRequest.Username) > 20 {
		return ErrWrongUsernameFormat
	}

	_, err := mail.ParseAddress(signupRequest.Email)
	if err != nil {
		return ErrWrongEmailFormat
	}

	if len(signupRequest.Password) < 5 {
		return ErrWrongPasswordFormat
	}

	return nil
}
