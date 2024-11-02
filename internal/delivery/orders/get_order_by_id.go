package orders

import (
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (h *OrdersHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	deliveryDateStr := vars["delivery_date"]

	id, err := uuid.Parse(idStr)
	if err != nil {
		h.logger.Error("[delivery.GetOrderById] Invalid order ID format", slog.String("orderID", idStr))
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid order ID format"})
		return
	}

	var deliveryDate time.Time
	if deliveryDate, err = time.Parse("2006-01-02T15:04:05.000Z", deliveryDateStr); err != nil {
		if deliveryDate, err = time.Parse("2006-01-02T15:04:05.000000Z", deliveryDateStr); err != nil {
			h.logger.Error("[delivery.GetOrderById] Invalid delivery date format", slog.String("deliveryDate", deliveryDateStr))
			utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid delivery date format"})
			return
		}
	}

	order, err := h.ordersManager.GetOrderById(r.Context(), id, deliveryDate)
	if err != nil {
		h.logger.Error("[delivery.GetOrderById] Order not found", slog.String("orderID", id.String()))
		utils.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "Order not found"})
		return
	}

	response := orderDTOMax{
		ID:           order.ID,
		Recipient:    order.Recipient,
		Address:      order.Address,
		Status:       order.Status,
		OrderDate:    order.OrderDate,
		DeliveryDate: order.DeliveryDate,
		Products:     convertProductsToDTO(order.Products),
	}

	h.logger.Info("[delivery.GetOrderById] Order retrieved successfully", slog.String("orderID", id.String()))
	utils.WriteJSON(w, http.StatusOK, response)
}
