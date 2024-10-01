package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/db"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
)

func (a *AuthApp) SignUp(w http.ResponseWriter, r *http.Request) {
	var signupRequest credsApiRequest

	err := json.NewDecoder(r.Body).Decode(&signupRequest)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, errs.HTTPErrorResponse{
			ErrorMessage: errs.InvalidJSONFormat.Error(),
		})

		return
	}

	if !validateRegistration(w, signupRequest) {
		return
	}

	_, exists := a.userDB.GetUserByEmail(signupRequest.Email)
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
		Username: signupRequest.Username,
		Password: hashedPassword,
	}

	err = a.userDB.CreateUser(signupRequest.Email, user)
	if err != nil {
		writeJSON(w, http.StatusConflict, errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})

		return
	}

	session, err := a.sessions.Get(r)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, errs.HTTPErrorResponse{
			ErrorMessage: errs.SessionCreationError.Error(),
		})

		return
	}

	session.Values[sessionKey] = signupRequest.Email
	setSessionOptions(session, 10*hour)

	err = a.sessions.Save(w, r, session)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, errs.HTTPErrorResponse{
			ErrorMessage: errs.SessionSaveError.Error(),
		})

		return
	}

	writeJSON(w, http.StatusOK, UsernameResponse{Username: user.Username})
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			return
		}
	}
}
