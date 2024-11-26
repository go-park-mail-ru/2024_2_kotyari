package address

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (h *AddressDelivery) UpdateAddressData(w http.ResponseWriter, r *http.Request) {
	requestID, err := utils.GetContextRequestID(r.Context())
	if err != nil {
		h.log.Error("[AddressDelivery.UpdateAddressData] No r ID")
		utils.WriteErrorJSONByError(w, err, h.errResolver)

		return
	}

	h.log.Info("[AddressDelivery.UpdateAddressData] Started executing", slog.Any("r-id", requestID))

	userID, ok := utils.GetContextSessionUserID(r.Context())
	if !ok {
		utils.WriteErrorJSON(w, http.StatusUnauthorized, errs.UserNotAuthorized)
	}

	var req UpdateAddressRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error("[ AddressDelivery.UpdateAddressData ] Ошибка десериализации запроса",
			slog.String("error", err.Error()),
		)

		utils.WriteErrorJSON(w, http.StatusBadRequest, errs.InvalidJSONFormat)
		return
	}

	newAddressData := req.ToModel()

	if err := h.addressManager.UpdateUsersAddress(r.Context(), userID, newAddressData); err != nil {
		h.log.Warn("[ AddressDelivery.UpdateAddressData ] Не удалось обновить данные профиля", slog.String("error", err.Error()))

		switch {
		case errors.Is(err, errs.InvalidEmailFormat):
			utils.WriteErrorJSON(w, http.StatusBadRequest, err)
		case errors.Is(err, errs.InvalidUsernameFormat):
			utils.WriteErrorJSON(w, http.StatusBadRequest, err)
		default:
			utils.WriteErrorJSON(w, http.StatusInternalServerError, errs.InternalServerError)
		}

		return
	}

	utils.WriteJSON(w, http.StatusOK, req)
}
