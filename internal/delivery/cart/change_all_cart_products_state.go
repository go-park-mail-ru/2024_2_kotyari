package cart

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (ch *CartHandler) ChangeAllCartProductsState(w http.ResponseWriter, r *http.Request) {
	userID := utils.GetContextSessionUserID(r.Context())

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
