package user

import (
	"errors"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"net/http"
)

func (d *UsersHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(utils.SessionName)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			utils.WriteJSON(w, http.StatusUnauthorized, errs.HTTPErrorResponse{
				ErrorMessage: errs.UserNotAuthorized.Error(),
			})

			return
		}

		utils.WriteJSON(w, http.StatusInternalServerError, errs.HTTPErrorResponse{
			ErrorMessage: errs.InternalServerError.Error(),
		})

		return
	}
	d.userClientGrpc.GetUserById()
	user, err := d.userClientGrpc.GetUserBySessionID(r.Context(), cookie.Value)
	if err != nil {
		err, code := d.errResolver.Get(err)
		utils.WriteJSON(w, code, errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})

		return
	}

	utils.WriteJSON(w, http.StatusOK, UsersDefaultResponse{
		Username:  user.Username,
		City:      user.City,
		AvatarUrl: user.AvatarUrl,
	})
}
