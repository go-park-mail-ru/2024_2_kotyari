package cart

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"net/http"
)

func (ch *CartHandler) GetSelectedFromCart(w http.ResponseWriter, r *http.Request) {
	userID := utils.GetContextSessionUserID(r.Context())

	cart, err := ch.cartManager.GetSelectedFromCart(r.Context(), userID)
	if err != nil {
		err, code := ch.errResolver.Get(err)
		utils.WriteErrorJSON(w, code, err)
		return
	}

	cartResponse := cartForOrderResponseFromModel(cart)
	utils.WriteJSON(w, http.StatusOK, cartResponse)
}
