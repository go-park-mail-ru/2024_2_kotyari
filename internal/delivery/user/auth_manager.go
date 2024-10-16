package user

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	userR "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/user"
	userU "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/user"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"net/http"
)

type AuthManager struct {
	userDB   userU.RepoInterface
	delivery Delivery
	sessions SessionInterface
}

func NewAuthManager(sessions SessionInterface) *AuthManager {
	return &AuthManager{
		userDB:   userR.InitUsers(),
		delivery: *NewUserDelivery(userU.NewUserUsecase(userR.NewUserMapRepository())),
		sessions: sessions,
	}
}

func newTestsAuthManager() *AuthManager {
	userTestDB := userR.InitUsersWithData()
	return &AuthManager{
		userDB:   userTestDB,
		delivery: *NewUserDelivery(userU.NewUserUsecase(userR.NewUserMapRepository())),
		sessions: newTestSessions(),
	}
}

func (am *AuthManager) SignUp(w http.ResponseWriter, r *http.Request) {
	user, err := am.delivery.CreateUser(r)
	if err != nil {
		utils.WriteJSON(w, errs.ErrCodesMapping[err], errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})
	}

	session, err := am.sessions.Get(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, errs.HTTPErrorResponse{
			ErrorMessage: errs.SessionCreationError.Error(),
		})

		return
	}

	session.Values[sessionKey] = user.Email
	setSessionOptions(session, 10*hour)
	err = am.sessions.Save(w, r, session)
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

func (am *AuthManager) Login(w http.ResponseWriter, r *http.Request) {
	session, err := am.sessions.Get(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, errs.HTTPErrorResponse{
			ErrorMessage: errs.SessionCreationError.Error(),
		})

		return
	}

	user, err := am.delivery.GetUserByEmail(r)
	if err != nil {
		utils.WriteJSON(w, errs.ErrCodesMapping[err], errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})

		return
	}

	session.Values[sessionKey] = user.Email
	setSessionOptions(session, 10*hour)

	err = am.sessions.Save(w, r, session)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, errs.HTTPErrorResponse{
			ErrorMessage: errs.SessionSaveError.Error(),
		})

		return
	}

	utils.WriteJSON(w, http.StatusOK, model.UsernameResponse{Username: user.Username})
}

func (am *AuthManager) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := am.sessions.Get(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusUnauthorized, errs.HTTPErrorResponse{
			ErrorMessage: errs.UserNotAuthorised.Error(),
		})

		return
	}

	session.Values = make(map[interface{}]interface{})
	setSessionOptions(session, nullTime)

	err = am.sessions.Save(w, r, session)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, errs.HTTPErrorResponse{
			ErrorMessage: errs.LogoutError.Error(),
		})

		return
	}

	utils.WriteJSON(w, http.StatusNoContent, nil)
}
