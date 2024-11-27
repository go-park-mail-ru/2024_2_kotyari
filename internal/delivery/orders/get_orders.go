package orders

import (
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (h *OrdersHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	requestID, err := utils.GetContextRequestID(r.Context())
	if err != nil {
		h.logger.Error("[OrdersHandler.GetOrders] No request ID")
		utils.WriteErrorJSONByError(w, err, h.errResolver)

		return
	}

	h.logger.Info("[OrdersHandler.GetOrders] Started executing", slog.Any("request-id", requestID))

	userID, ok := utils.GetContextSessionUserID(r.Context())
	if !ok {
		utils.WriteErrorJSON(w, http.StatusUnauthorized, errs.UserNotAuthorized)
	}

	orders, err := h.ordersManager.GetOrders(r.Context(), userID)
	if err != nil {
		h.logger.Error("[delivery.GetOrders] Failed to fetch orders", slog.String("error", err.Error()))
		utils.WriteErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	orderDTOs := ConvertOrdersToResponse(orders)
	h.logger.Info("[delivery.GetOrders] Fetched orders successfully")

	utils.WriteJSON(w, http.StatusOK, orderDTOs)
}
