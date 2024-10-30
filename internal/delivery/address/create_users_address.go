package address

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"log/slog"
	"net/http"
)

func (h *AddressDelivery) CreateAddress(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	h.log.Info("Обработка запроса на создание адреса", "request", request)

	id := utils.GetContextSessionUserID(request.Context())

	var req AddressRequest

	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		h.log.Error("Ошибка десериализации запроса", slog.String("error", err.Error()))
		utils.WriteJSON(writer, http.StatusBadRequest, &errs.HTTPErrorResponse{ErrorMessage: errs.InvalidJSONFormat.Error()})
		return
	}

	newAddressData := model.Address{
		City:   req.City,
		Street: req.Street,
		House:  req.House,
		Flat:   req.Flat,
	}

	_, err := h.addressManager.CreateUsersAddress(ctx, uint32(id), newAddressData)

	if err != nil {
		h.log.Warn("Не удалось обновить адрес профиля", slog.String("error", err.Error()))

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

	h.log.Info("Успешное получение профиля", "id", id)

	utils.WriteJSON(writer, http.StatusOK, req)
}
