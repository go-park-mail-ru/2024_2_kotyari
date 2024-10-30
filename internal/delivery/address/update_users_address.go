package address

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"log/slog"
	"net/http"
)

func (h *AddressDelivery) UpdateAddressData(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	h.log.Info("Обработка запроса на обновление адреса профиля", "request", request)

	id := utils.GetContextSessionUserID(request.Context())

	var req AddressRequest

	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		h.log.Error("Ошибка десериализации запроса", slog.String("error", err.Error()))
		utils.WriteJSON(writer, http.StatusBadRequest, &errs.HTTPErrorResponse{ErrorMessage: errs.InvalidJSONFormat.Error()})
		return
	}

	defer request.Body.Close()
	h.log.Debug("Получено тело запроса", "req", req)

	oldAddressData, err := h.addressManager.GetAddressByProfileID(ctx, uint32(id))
	if err != nil {
		h.log.Warn("Не удалось получить старый адрес профиля", slog.String("error", err.Error()))
		utils.WriteJSON(writer, http.StatusNotFound, &errs.HTTPErrorResponse{ErrorMessage: errs.UserDoesNotExist.Error()})
		return
	}

	newAddressData := model.Address{
		City:   req.City,
		Street: req.Street,
		House:  req.House,
		Flat:   req.Flat,
	}

	if err := h.addressManager.UpdateUsersAddress(oldAddressData.Id, newAddressData); err != nil {
		h.log.Warn("Не удалось обновить данные профиля", slog.String("error", err.Error()))

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

	h.log.Info("Данные адреса обновлены", "id", id)
	utils.WriteJSON(writer, http.StatusOK, req)
}
