package cart

import (
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (ch *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	userID, ok := utils.GetContextSessionUserID(r.Context())
	if !ok {
		utils.WriteErrorJSON(w, http.StatusUnauthorized, errs.UserNotAuthorized)

		return
	}

	cart, err := ch.cartManip.GetCart(r.Context(), userID, time.Now())
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
