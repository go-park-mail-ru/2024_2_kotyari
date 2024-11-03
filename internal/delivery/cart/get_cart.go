package cart

import (
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (ch *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	cart, err := ch.cartManip.GetCart(r.Context(), time.Now())
	if err != nil {
		err, code := ch.errResolver.Get(err)
		utils.WriteErrorJSON(w, code, err)

		return
	}

	var cartResp GetCartResponse
	productResp := make([]GetProductResponse, 0, len(cart.Products))

	for _, cartProduct := range cart.Products {
		productResp = append(productResp, productResponseFromModel(cartProduct))
	}

	cartResp = cartResponseFromModel(cart, productResp)
	utils.WriteJSON(w, http.StatusOK, cartResp)
}
