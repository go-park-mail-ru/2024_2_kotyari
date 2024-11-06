package orders

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (h *OrdersHandler) CreateOrderFromCart(w http.ResponseWriter, r *http.Request) {
	userID, ok := utils.GetContextSessionUserID(r.Context())
	if !ok {
		utils.WriteErrorJSON(w, http.StatusUnauthorized, errs.UserNotAuthorized)

		return
	}

	fmt.Println(userID)

	var request CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		h.logger.Error("[delivery.CreateOrderFromCart] Invalid request body", slog.String("error", err.Error()))
		utils.WriteErrorJSONByError(w, errs.InvalidJSONFormat, h.errResolver)

		return
	}

	order, err := h.ordersManager.CreateOrderFromCart(r.Context(), request.Address, userID)
	if err != nil {
		h.logger.Error("[delivery.CreateOrderFromCart] Failed to create order from cart", slog.String("error", err.Error()))
		utils.WriteErrorJSONByError(w, errs.InternalServerError, h.errResolver)

		return
	}

	orderDTO := ToOrderResponse(order)
	h.logger.Info("[delivery.CreateOrderFromCart] Order created from cart", slog.String("orderID", orderDTO.ID.String()))

	utils.WriteJSON(w, http.StatusOK, orderDTO)
}
