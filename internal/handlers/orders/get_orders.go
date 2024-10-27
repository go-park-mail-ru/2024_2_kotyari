package orders

import (
	"encoding/json"
	"net/http"
)

func (h *OrdersHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.ordersManager.GetOrders(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch orders", http.StatusInternalServerError)
		return
	}

	var orderDTOs []OrderDTO
	for _, ord := range orders {
		orderDTOs = append(orderDTOs, ToOrderDTO(ord))
	}

	resp := GetOrdersResponse{
		Status: http.StatusOK,
		Body:   orderDTOs,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
