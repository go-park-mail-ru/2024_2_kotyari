package orders

import (
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (h *OrdersHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.ordersManager.GetOrders(r.Context())
	if err != nil {
		h.logger.Error("Failed to fetch orders", slog.String("error", err.Error()))
		utils.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to fetch orders"})
		return
	}

	orderDTOs := convertOrdersToDTOs(orders)
	h.logger.Info("Fetched orders successfully", slog.Int("orderCount", len(orderDTOs)))

	utils.WriteJSON(w, http.StatusOK, getOrdersResponse{
		Status: http.StatusOK,
		Body:   orderDTOs,
	})
}
