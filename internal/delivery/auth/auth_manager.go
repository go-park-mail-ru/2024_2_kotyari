package auth

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/user"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	userR "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/user"
	userU "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/user"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"net/http"
)

type Manager struct {
	Delivery user.Delivery
	Sessions SessionInterface
}

func NewAuthManager(sessions SessionInterface) *Manager {
	return &Manager{
		Delivery: *user.NewUserDelivery(userU.NewUserUsecase(userR.NewUserMapRepository())),
		Sessions: sessions,
	}
}

func newTestsAuthManager() *Manager {
	return &Manager{
		Delivery: *user.NewUserDelivery(userU.NewUserUsecase(userR.NewUserMapRepository())),
		Sessions: newTestSessions(),
	}
}

func (am *Manager) SignUp(w http.ResponseWriter, r *http.Request) {
	user, err := am.Delivery.CreateUser(r)
	if err != nil {
		utils.WriteJSON(w, errs.ErrCodesMapping[err], errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})
	}

	session, err := am.Sessions.Get(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, errs.HTTPErrorResponse{
			ErrorMessage: errs.SessionCreationError.Error(),
		})

		return
	}

	session.Values[SessionKey] = user.Email
	setSessionOptions(session, 10*hour)
	err = am.Sessions.Save(w, r, session)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, errs.HTTPErrorResponse{
			ErrorMessage: errs.SessionSaveError.Error(),
		})

		return
	}

	utils.WriteJSON(w, http.StatusOK, model.UsernameResponse{
		Username: user.Username,
	})
}

func (am *Manager) Login(w http.ResponseWriter, r *http.Request) {
	session, err := am.Sessions.Get(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, errs.HTTPErrorResponse{
			ErrorMessage: errs.SessionCreationError.Error(),
		})

		return
	}

	user, err := am.Delivery.LoginByEmail(r)
	if err != nil {
		utils.WriteJSON(w, errs.ErrCodesMapping[err], errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})

		return
	}

	session.Values[SessionKey] = user.Email
	setSessionOptions(session, 10*hour)

	err = am.Sessions.Save(w, r, session)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, errs.HTTPErrorResponse{
			ErrorMessage: errs.SessionSaveError.Error(),
		})

		return
	}

	utils.WriteJSON(w, http.StatusOK, model.UsernameResponse{Username: user.Username})
}

func (am *Manager) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := am.Sessions.Get(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusUnauthorized, errs.HTTPErrorResponse{
			ErrorMessage: errs.UserNotAuthorized.Error(),
		})

		return
	}

	session.Values = make(map[interface{}]interface{})
	setSessionOptions(session, nullTime)

	err = am.Sessions.Save(w, r, session)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, errs.HTTPErrorResponse{
			ErrorMessage: errs.LogoutError.Error(),
		})

		return
	}

	utils.WriteJSON(w, http.StatusNoContent, nil)
}

func (am *Manager) Soon(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)

	utils.WriteJSON(w, http.StatusOK, model.UsernameResponse{
		Username: username,
	})
}
