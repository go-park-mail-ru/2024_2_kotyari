package orders

import (
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (h *OrdersHandler) CreateOrderFromCart(w http.ResponseWriter, r *http.Request) {
	var request createOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		h.logger.Error("Invalid request body", slog.String("error", err.Error()))
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	order, err := h.ordersManager.CreateOrderFromCart(r.Context(), request.Address)
	if err != nil {
		h.logger.Error("Failed to create order from cart", slog.String("error", err.Error()))
		utils.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to create order"})
		return
	}

	orderDTO := ToOrderDTO(*order)
	h.logger.Info("Order created from cart", slog.String("orderID", orderDTO.ID.String()))

	utils.WriteJSON(w, http.StatusOK, createOrderResponse{
		Status: http.StatusOK,
		Body:   orderDTO,
	})
}
