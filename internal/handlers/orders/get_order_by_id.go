package orders

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

func (h *OrdersHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/order/"):]

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid order ID format", http.StatusBadRequest)
		return
	}

	ord, err := h.ordersManager.GetOrderByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	resp := GetOrderByIDResponse{
		Status: http.StatusOK,
		Body:   ToOrderDTO(*ord),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
