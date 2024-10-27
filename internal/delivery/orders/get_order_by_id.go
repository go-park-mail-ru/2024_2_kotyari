package orders

import (
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (h *OrdersHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := uuid.Parse(idStr)
	if err != nil {
		h.logger.Error("Invalid order ID format", slog.String("orderID", idStr))
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid order ID format"})
		return
	}

	ord, err := h.ordersManager.GetOrderByID(r.Context(), id)
	if err != nil {
		h.logger.Error("Order not found", slog.String("orderID", id.String()))
		utils.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "Order not found"})
		return
	}

	h.logger.Info("Order retrieved successfully", slog.String("orderID", id.String()))
	utils.WriteJSON(w, http.StatusOK, getOrderByIDResponse{
		Status: http.StatusOK,
		Body:   ToOrderDTO(*ord),
	})
}
