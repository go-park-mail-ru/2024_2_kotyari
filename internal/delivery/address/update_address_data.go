package address

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (h *AddressDelivery) UpdateAddressData(writer http.ResponseWriter, request *http.Request) {

	id := utils.GetContextSessionUserID(request.Context())

	var req UpdateAddressRequest

	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		h.log.Error("[ AddressDelivery.UpdateAddressData ] Ошибка десериализации запроса", slog.String("error", err.Error()))
		utils.WriteJSON(writer, http.StatusBadRequest, &errs.HTTPErrorResponse{ErrorMessage: errs.InvalidJSONFormat.Error()})
		return
	}

	oldAddressData, err := h.addressManager.GetAddressByProfileID(request.Context(), uint32(id))
	if err != nil {
		h.log.Warn("[ AddressDelivery.UpdateAddressData ] Не удалось получить старый адрес профиля", slog.String("error", err.Error()))
		utils.WriteJSON(writer, http.StatusNotFound, &errs.HTTPErrorResponse{ErrorMessage: errs.UserDoesNotExist.Error()})
		return
	}

	newAddressData := req.ToModel()

	if err := h.addressManager.UpdateUsersAddress(request.Context(), oldAddressData.Id, newAddressData); err != nil {
		h.log.Warn("[ AddressDelivery.UpdateAddressData ] Не удалось обновить данные профиля", slog.String("error", err.Error()))

		switch err {
		case errs.InvalidEmailFormat:
			utils.WriteJSON(writer, http.StatusBadRequest, &errs.HTTPErrorResponse{ErrorMessage: err.Error()})
		case errs.InvalidUsernameFormat:
			utils.WriteJSON(writer, http.StatusBadRequest, &errs.HTTPErrorResponse{ErrorMessage: err.Error()})
		default:
			utils.WriteJSON(writer, http.StatusInternalServerError, &errs.HTTPErrorResponse{ErrorMessage: errs.InternalServerError.Error()})
		}
		return
	}

	utils.WriteJSON(writer, http.StatusOK, req)
}
