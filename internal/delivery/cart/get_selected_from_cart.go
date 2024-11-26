package cart

import (
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (ch *CartHandler) GetSelectedFromCart(w http.ResponseWriter, r *http.Request) {
	requestID, err := utils.GetContextRequestID(r.Context())
	if err != nil {
		ch.log.Error("[CartHandler.GetSelectedFromCart] No request ID")
		utils.WriteErrorJSONByError(w, err, ch.errResolver)

		return
	}

	ch.log.Info("[CartHandler.GetSelectedFromCart] Started executing", slog.Any("request-id", requestID))

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
