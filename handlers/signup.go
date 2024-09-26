package handlers

import (
	"encoding/json"
	"net/http"

	"2024_2_kotyari/config"
	"2024_2_kotyari/db"
	"2024_2_kotyari/errs"
)

// SignupHandler handles user signup requests
// @Summary      Signup a new user
// @Description  This endpoint creates a new user in the system
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body  credsApiRequest  true  "Signup Request Body"
// @Success      200   {string}  string  "OK"
// @Failure      400   {object}  errs.HTTPErrorResponse "Invalid request"
// @Failure      409   {object}  errs.HTTPErrorResponse "User already exists"
// @Failure      500   {object}  errs.HTTPErrorResponse "Internal server error"
// @Router       /signup [post]
func (s *Server) SignupHandler(w http.ResponseWriter, r *http.Request) {
	var signupRequest credsApiRequest

	err := json.NewDecoder(r.Body).Decode(&signupRequest)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, errs.HTTPErrorResponse{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: errs.InvalidJSONFormat.Error(),
		})
		return
	}
	if !validateCredentials(&w, signupRequest, true) {
		return
	}

	err = db.CreateUser(signupRequest.Email, signupRequest.User)
	if err != nil {
		writeJSON(w, http.StatusConflict, errs.HTTPErrorResponse{
			ErrorCode:    http.StatusConflict,
			ErrorMessage: err.Error(),
		})
		return
	}

	session, err := s.sessions.Get(r, config.GetSessionName())
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, errs.HTTPErrorResponse{
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: errs.SessionCreationError.Error(),
		})
		return
	}
	session.Values["user_id"] = signupRequest.Email
	session.Options.MaxAge = 3600 * 10
	session.Options.HttpOnly = true

	err = s.sessions.Save(r, w, session)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, errs.HTTPErrorResponse{
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: errs.SessionSaveError.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, nil)
}

func writeJSON(w http.ResponseWriter, headerStatusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(headerStatusCode)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			http.Error(w, errs.InternalServerError.Error(), http.StatusInternalServerError)
		}
	}
}
