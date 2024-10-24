package auth

//
//import (
//	"net/http"
//
//	"github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/user"
//	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
//	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
//	userR "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/user"
//	userU "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/user"
//	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
//)
//
//type AuthManager interface {
//	SignUp(w http.ResponseWriter, r *http.Request)
//	Login(w http.ResponseWriter, r *http.Request)
//	Logout(w http.ResponseWriter, r *http.Request)
//	UnimplementedRoutesHandler(w http.ResponseWriter, r *http.Request)
//}
//
//type Manager struct {
//	Delivery user.UserDelivery
//	Sessions SessionInterface
//}
//
//func NewAuthManager(sessions SessionInterface) AuthManager {
//	return &Manager{
//		Delivery: user.NewUserDelivery(userU.NewUserUseCase(userR.NewUserMapRepository())),
//		Sessions: sessions,
//	}
//}
//
//func newTestsAuthManager() *Manager {
//	return &Manager{
//		Delivery: user.NewUserDelivery(userU.NewUserUseCase(userR.NewUserMapRepository())),
//		Sessions: newTestSessions(),
//	}
//}
//
//func (am *Manager) SignUp(w http.ResponseWriter, r *http.Request) {
//	userDTO, err := am.Delivery.CreateUser(r)
//	if err != nil {
//		utils.WriteJSON(w, errs.ErrCodesMapping[err], errs.HTTPErrorResponse{
//			ErrorMessage: err.Error(),
//		})
//
//		return
//	}
//
//	session, err := am.Sessions.Get(r)
//	if err != nil {
//		utils.WriteJSON(w, http.StatusInternalServerError, errs.HTTPErrorResponse{
//			ErrorMessage: errs.SessionCreationError.Error(),
//		})
//
//		return
//	}
//
//	session.Values[SessionKey] = userDTO.Email
//	setSessionOptions(session, defaultSessionSetTime)
//	err = am.Sessions.Save(w, r, session)
//	if err != nil {
//		utils.WriteJSON(w, http.StatusInternalServerError, errs.HTTPErrorResponse{
//			ErrorMessage: errs.SessionSaveError.Error(),
//		})
//
//		return
//	}
//
//	utils.WriteJSON(w, http.StatusOK, model.UsernameResponse{
//		Username: userDTO.Username,
//	})
//}
//
//func (am *Manager) Login(w http.ResponseWriter, r *http.Request) {
//	session, err := am.Sessions.Get(r)
//	if err != nil {
//		utils.WriteJSON(w, http.StatusInternalServerError, errs.HTTPErrorResponse{
//			ErrorMessage: errs.SessionCreationError.Error(),
//		})
//
//		return
//	}
//
//	userDTO, err := am.Delivery.LoginByEmail(r)
//	if err != nil {
//		utils.WriteJSON(w, errs.ErrCodesMapping[err], errs.HTTPErrorResponse{
//			ErrorMessage: err.Error(),
//		})
//
//		return
//	}
//
//	session.Values[SessionKey] = userDTO.Email
//	setSessionOptions(session, defaultSessionSetTime)
//	err = am.Sessions.Save(w, r, session)
//	if err != nil {
//		utils.WriteJSON(w, http.StatusInternalServerError, errs.HTTPErrorResponse{
//			ErrorMessage: errs.SessionSaveError.Error(),
//		})
//
//		return
//	}
//
//	utils.WriteJSON(w, http.StatusOK, model.UsernameResponse{
//		Username: userDTO.Username,
//	})
//}
//
//func (am *Manager) Logout(w http.ResponseWriter, r *http.Request) {
//	session, err := am.Sessions.Get(r)
//	if err != nil {
//		utils.WriteJSON(w, http.StatusUnauthorized, errs.HTTPErrorResponse{
//			ErrorMessage: errs.UserNotAuthorized.Error(),
//		})
//
//		return
//	}
//
//	setSessionOptions(session, deleteSession)
//	err = am.Sessions.Save(w, r, session)
//	if err != nil {
//		utils.WriteJSON(w, http.StatusInternalServerError, errs.HTTPErrorResponse{
//			ErrorMessage: errs.LogoutError.Error(),
//		})
//
//		return
//	}
//
//	utils.WriteJSON(w, http.StatusNoContent, nil)
//}
//
//func (am *Manager) UnimplementedRoutesHandler(w http.ResponseWriter, r *http.Request) {
//	email := r.Context().Value(SessionKey).(string)
//	userDTO, err := am.Delivery.GetSessionUser(email)
//	if err != nil {
//		utils.WriteJSON(w, errs.ErrCodesMapping[err], errs.HTTPErrorResponse{
//			ErrorMessage: err.Error(),
//		})
//
//		return
//	}
//
//	utils.WriteJSON(w, http.StatusOK, model.UsernameResponse{
//		Username: userDTO.Username,
//	})
//}
