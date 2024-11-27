package cart

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (ch *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	requestID, err := utils.GetContextRequestID(r.Context())
	if err != nil {
		ch.log.Error("[CartHandler.GetCart] No request ID")
		utils.WriteErrorJSONByError(w, err, ch.errResolver)

		return
	}

	ch.log.Info("[CartHandler.GetCart] Started executing", slog.Any("request-id", requestID))

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
