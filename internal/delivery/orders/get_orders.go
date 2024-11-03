package orders

import (
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (h *OrdersHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.ordersManager.GetOrders(r.Context())
	if err != nil {
		h.logger.Error("[delivery.GetOrders] Failed to fetch orders", slog.String("error", err.Error()))
		utils.WriteErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	orderDTOs := convertOrdersToDTOs(orders)
	h.logger.Info("[delivery.GetOrders] Fetched orders successfully")

	utils.WriteJSON(w, http.StatusOK, orderDTOs)
}
