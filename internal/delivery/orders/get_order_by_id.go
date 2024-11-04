package orders

import (
	"errors"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (h *OrdersHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	userID := utils.GetContextSessionUserID(r.Context())

	vars := mux.Vars(r)
	idStr := vars["id"]
	deliveryDateStr := vars["delivery_date"]

	id, err := uuid.Parse(idStr)
	if err != nil {
		h.logger.Error("[delivery.GetOrderById] Invalid order ID format", slog.String("orderID", idStr))
		utils.WriteErrorJSONByError(w, errs.ErrInvalidOrderIDFormat, h.errResolver)
		return
	}

	var deliveryDate time.Time
	if deliveryDate, err = time.Parse("2006-01-02T15:04:05.000Z", deliveryDateStr); err != nil {
		if deliveryDate, err = time.Parse("2006-01-02T15:04:05.000000Z", deliveryDateStr); err != nil {
			h.logger.Error("[delivery.GetOrderById] Invalid delivery date format", slog.String("deliveryDate", deliveryDateStr))
			utils.WriteErrorJSONByError(w, errs.ErrInvalidDeliveryDateFormat, h.errResolver)
			return
		}
	}

	order, err := h.ordersManager.GetOrderById(r.Context(), id, deliveryDate, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			h.logger.Warn("[delivery.GetOrderById] Order not found", slog.String("orderID", id.String()))
			utils.WriteErrorJSONByError(w, errs.ErrOrderNotFound, h.errResolver)
		} else {
			h.logger.Error("[delivery.GetOrderById] Failed to retrieve order", slog.String("orderID", id.String()), slog.String("error", err.Error()))
			utils.WriteErrorJSONByError(w, errs.ErrRetrieveOrder, h.errResolver)
		}
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
