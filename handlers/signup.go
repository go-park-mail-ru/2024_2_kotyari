package handlers

import (
	"encoding/json"
	"net/http"

	"2024_2_kotyari/config"
	"2024_2_kotyari/db"
	"2024_2_kotyari/errs"
)

// Signup handles user signup requests
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
func (s *Server) Signup(w http.ResponseWriter, r *http.Request) {
	var signupRequest credsApiRequest

	err := json.NewDecoder(r.Body).Decode(&signupRequest)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, errs.HTTPErrorResponse{
			ErrorMessage: errs.InvalidJSONFormat.Error(),
		})
		return
	}
	if !validateRegistration(&w, signupRequest) {
		return
	}

	_, exists := db.GetUserByEmail(signupRequest.Email)
	if exists {
		writeJSON(w, http.StatusConflict, errs.HTTPErrorResponse{
			ErrorMessage: errs.UserAlreadyExists.Error(),
		})
		return
	}

	// Генерация соли и хеширование пароля
	salt, err := generateSalt()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, errs.HTTPErrorResponse{
			ErrorMessage: errs.InternalServerError.Error(),
		})
		return
	}
	hashedPassword := hashPassword(signupRequest.Password, salt)

	// Сохраняем нового пользователя
	user := db.User{
		Username:     signupRequest.Username,
		PasswordHash: hashedPassword,
	}
	db.CreateUser(signupRequest.Email, user)

	session, err := s.sessions.Get(r, config.GetSessionName())
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, errs.HTTPErrorResponse{
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
			ErrorMessage: errs.SessionSaveError.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, nil)
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
