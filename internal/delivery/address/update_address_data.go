package address

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (h *AddressDelivery) UpdateAddressData(writer http.ResponseWriter, request *http.Request) {
	userID, ok := utils.GetContextSessionUserID(request.Context())
	if !ok {
		utils.WriteErrorJSON(writer, http.StatusUnauthorized, errs.UserNotAuthorized)
	}

	var req UpdateAddressRequest

	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		h.log.Error("[ AddressDelivery.UpdateAddressData ] Ошибка десериализации запроса",
			slog.String("error", err.Error()),
		)

		utils.WriteErrorJSON(writer, http.StatusBadRequest, errs.InvalidJSONFormat)
		return
	}

	newAddressData := req.ToModel()

	if err := h.addressManager.UpdateUsersAddress(request.Context(), userID, newAddressData); err != nil {
		h.log.Warn("[ AddressDelivery.UpdateAddressData ] Не удалось обновить данные профиля", slog.String("error", err.Error()))

		switch {
		case errors.Is(err, errs.InvalidEmailFormat):
			utils.WriteErrorJSON(writer, http.StatusBadRequest, err)
		case errors.Is(err, errs.InvalidUsernameFormat):
			utils.WriteErrorJSON(writer, http.StatusBadRequest, err)
		default:
			utils.WriteErrorJSON(writer, http.StatusInternalServerError, errs.InternalServerError)
		}

		return
	}

	utils.WriteJSON(writer, http.StatusOK, req)
}
