package orders

import (
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (h *OrdersHandler) GetNearestDeliveryDate(w http.ResponseWriter, r *http.Request) {
	requestID, err := utils.GetContextRequestID(r.Context())
	if err != nil {
		h.logger.Error("[OrdersHandler.GetNearestDeliveryDate] No request ID")
		utils.WriteErrorJSONByError(w, err, h.errResolver)

		return
	}

	h.logger.Info("[OrdersHandler.GetNearestDeliveryDate] Started executing", slog.Any("request-id", requestID))

	userID, ok := utils.GetContextSessionUserID(r.Context())
	if !ok {
		utils.WriteErrorJSON(w, http.StatusUnauthorized, errs.UserNotAuthorized)
	}

	deliveryDate, err := h.ordersManager.GetNearestDeliveryDate(r.Context(), userID)
	if err != nil {
		h.logger.Error("[delivery.GetNearestDeliveryDate] Failed to fetch nearest delivery date", slog.String("error", err.Error()))
		utils.WriteErrorJSONByError(w, errs.ErrGetNearestDeliveryDate, h.errResolver)
		return
	}

	h.logger.Info("[delivery.GetNearestDeliveryDate] Fetched nearest delivery date successfully")
	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"delivery_date": deliveryDate,
	})
}
