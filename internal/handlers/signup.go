package handlers

import (
	"encoding/json"
	"net/http"

	userD "github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/user"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	userR "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/user"
	userU "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/user"
)

func (a *AuthApp) SignUp(w http.ResponseWriter, r *http.Request) {
	var signupRequest model.UserApiRequest
	userHandler := userD.NewUserDelivery(userU.NewUserUsecase(userR.NewUserMapRepository()))

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

	err = userHandler.CreateUser(signupRequest)
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

	writeJSON(w, http.StatusOK, model.UsernameResponse{
		Username: signupRequest.Username,
	})
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
