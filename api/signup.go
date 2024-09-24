package api

import (
	"encoding/json"
	"net/http"
	"net/mail"

	"2024_2_kotyari/db"
	"2024_2_kotyari/errs"
)

// SignupHandler handles user signup requests
// @Summary      Signup a new user
// @Description  This endpoint creates a new user in the system
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body  signupApiRequest  true  "Signup Request Body"
// @Success      200   {string}  string  "OK"
// @Failure      400   {object}  errs.HTTPErrorResponse "Invalid request"
// @Failure      409   {object}  errs.HTTPErrorResponse "User already exists"
// @Failure      500   {object}  errs.HTTPErrorResponse "Internal server error"
// @Router       /signup [post]
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var signupRequest signupApiRequest

	err := json.NewDecoder(r.Body).Decode(&signupRequest)
	if err != nil {
		writeJSON(w, errs.HTTPErrorResponse{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: errs.InvalidJSONFormat.Error(),
		})
		return
	}

	err = validateParams(signupRequest)
	if err != nil {
		writeJSON(w, errs.HTTPErrorResponse{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: err.Error(),
		})
		return
	}

	err = db.CreateUser(signupRequest.Email, signupRequest.User)
	if err != nil {
		writeJSON(w, errs.HTTPErrorResponse{
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
		return errs.EmptyRequestParams
	}

	if len(signupRequest.Username) < 3 || len(signupRequest.Username) > 20 {
		return errs.WrongUsernameFormat
	}

	_, err := mail.ParseAddress(signupRequest.Email)
	if err != nil {
		return errs.WrongEmailFormat
	}

	if len(signupRequest.Password) < 5 {
		return errs.WrongPasswordFormat
	}

	return nil
}
