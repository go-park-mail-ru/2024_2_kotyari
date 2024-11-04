package cart

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (ch *CartHandler) GetSelectedFromCart(w http.ResponseWriter, r *http.Request) {
	userID, ok := utils.GetContextSessionUserID(r.Context())
	if !ok {
		utils.WriteErrorJSON(w, http.StatusUnauthorized, errs.UserNotAuthorized)
	}

	cart, err := ch.cartManager.GetSelectedFromCart(r.Context(), userID)
	if err != nil {
		err, code := ch.errResolver.Get(err)
		utils.WriteErrorJSON(w, code, err)
		return
	}

	cartResponse := cartForOrderResponseFromModel(cart)
	utils.WriteJSON(w, http.StatusOK, cartResponse)
}
