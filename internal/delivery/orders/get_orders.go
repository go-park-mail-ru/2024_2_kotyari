package orders

import (
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (h *OrdersHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
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
