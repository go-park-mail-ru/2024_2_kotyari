package cart

import (
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (ch *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	cart, err := ch.cartGetter.GetCart(r.Context(), time.Now())
	if err != nil {
		/// TODO: Change code type
		utils.WriteErrorJSON(w, http.StatusBadRequest, err)

		return
	}

	if len(cart.Products) == 0 {

	}

	var cartResp CartResponse
	productResp := make([]ProductResponse, 0, len(cart.Products))

	for _, cartProduct := range cart.Products {
		productResp = append(productResp, ProductResponseFromModel(cartProduct))
	}

	cartResp = CartResponseFromModel(cart, productResp)
	utils.WriteJSON(w, http.StatusOK, cartResp)
}
