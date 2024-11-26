package cart

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (ch *CartHandler) UpdatePaymentMethod(w http.ResponseWriter, r *http.Request) {
	requestID, err := utils.GetContextRequestID(r.Context())
	if err != nil {
		ch.log.Error("[CartHandler.UpdatePaymentMethod] No request ID")
		utils.WriteErrorJSONByError(w, err, ch.errResolver)

		return
	}

	ch.log.Info("[CartHandler.UpdatePaymentMethod] Started executing", slog.Any("request-id", requestID))

	userID, ok := utils.GetContextSessionUserID(r.Context())
	if !ok {
		utils.WriteErrorJSON(w, http.StatusUnauthorized, errs.UserNotAuthorized)
	}

	var requestPaymentMethod requestPaymentMethod

	if err := json.NewDecoder(r.Body).Decode(&requestPaymentMethod); err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	err = ch.cartManip.UpdatePaymentMethod(r.Context(), userID, requestPaymentMethod.PaymentMethod)
	if err != nil {
		err, code := ch.errResolver.Get(err)
		utils.WriteErrorJSON(w, code, err)

		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "payment method updated"})
}
