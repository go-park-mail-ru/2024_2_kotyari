package cart

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (ch *CartHandler) ChangeAllCartProductsState(w http.ResponseWriter, r *http.Request) {
	userID, ok := utils.GetContextSessionUserID(r.Context())
	if !ok {
		utils.WriteErrorJSON(w, http.StatusUnauthorized, errs.UserNotAuthorized)
	}

	var selectState bool

	switch r.Method {
	case http.MethodPatch:
		selectState = true
	case http.MethodDelete:
		selectState = false
	}

	err := ch.cartManip.ChangeAllCartProductsState(r.Context(), userID, selectState)
	if err != nil {
		err, code := ch.errResolver.Get(err)
		utils.WriteErrorJSON(w, code, err)

		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}
