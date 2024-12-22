package orders

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/mailru/easyjson"
)

func (h *OrdersHandler) CreateOrderFromCart(w http.ResponseWriter, r *http.Request) {
	requestID, err := utils.GetContextRequestID(r.Context())
	if err != nil {
		h.logger.Error("[OrdersHandler.CreateOrderFromCart] No request ID")
		utils.WriteErrorJSONByError(w, err, h.errResolver)

		return
	}

	h.logger.Info("[OrdersHandler.CreateOrderFromCart] Started executing", slog.Any("request-id", requestID))

	userID, ok := utils.GetContextSessionUserID(r.Context())
	if !ok {
		utils.WriteErrorJSON(w, http.StatusUnauthorized, errs.UserNotAuthorized)
	}

	var request CreateOrderRequest
	if err = easyjson.UnmarshalFromReader(r.Body, &request); err != nil {
		h.logger.Error("[delivery.CreateOrderFromCart] Invalid request body", slog.String("error", err.Error()))
		utils.WriteErrorJSONByError(w, errs.InvalidJSONFormat, h.errResolver)
		return
	}

	order, err := h.ordersManager.CreateOrderFromCart(r.Context(), request.Address, userID, request.PromoCode)
	if err != nil {
		if errors.Is(err, errs.NoPromoCode) {
			h.logger.Error("[delivery.CreateOrderFromCart] No promo code", slog.String("error", err.Error()))
			utils.WriteErrorJSONByError(w, errs.NoPromoCode, h.errResolver)

			return
		}

		h.logger.Error("[delivery.CreateOrderFromCart] Failed to create order from cart", slog.String("error", err.Error()))
		utils.WriteErrorJSONByError(w, errs.InternalServerError, h.errResolver)
		return
	}

	orderDTO := ToOrderResponse(order)
	h.logger.Info("[delivery.CreateOrderFromCart] Order created from cart", slog.String("orderID", orderDTO.ID.String()))

	utils.WriteJSON(w, http.StatusOK, orderDTO)
}
