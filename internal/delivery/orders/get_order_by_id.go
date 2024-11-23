package orders

import (
	"errors"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"log/slog"
	"net/http"
)

func (h *OrdersHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	userID, ok := utils.GetContextSessionUserID(r.Context())
	if !ok {
		utils.WriteErrorJSON(w, http.StatusUnauthorized, errs.UserNotAuthorized)
	}

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := uuid.Parse(idStr)
	if err != nil {
		h.logger.Error("[delivery.GetOrderById] Invalid order ID format", slog.String("orderID", idStr))
		utils.WriteErrorJSONByError(w, errs.ErrInvalidOrderIDFormat, h.errResolver)
		return
	}

	order, err := h.ordersManager.GetOrderById(r.Context(), id, userID)
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

	response := OrderMaxResponse{
		ID:           order.ID,
		Recipient:    order.Recipient,
		Address:      order.Address,
		Status:       order.Status,
		OrderDate:    order.OrderDate,
		DeliveryDate: order.DeliveryDate,
		Products:     ConvertProductsToDTO(order.Products),
	}

	h.logger.Info("[delivery.GetOrderById] Order retrieved successfully", slog.String("orderID", id.String()))
	utils.WriteJSON(w, http.StatusOK, response)
}
