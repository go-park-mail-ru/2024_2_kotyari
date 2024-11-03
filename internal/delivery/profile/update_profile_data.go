package profile

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (h *ProfilesDelivery) UpdateProfileData(writer http.ResponseWriter, request *http.Request) {

	id := utils.GetContextSessionUserID(request.Context())

	var req UpdateProfileRequest

	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		h.log.Error("[ ProfilesDelivery.UpdateProfileData ] Ошибка десериализации запроса", slog.String("error", err.Error()))
		utils.WriteJSON(writer, http.StatusBadRequest, &errs.HTTPErrorResponse{ErrorMessage: errs.InvalidJSONFormat.Error()})
		return
	}

	oldProfileData, err := h.profileManager.GetProfile(request.Context(), uint32(id))
	if err != nil {
		h.log.Warn("[ ProfilesDelivery.UpdateProfileData ] Не удалось получить старые данные профиля", slog.String("error", err.Error()))
		utils.WriteJSON(writer, http.StatusNotFound, &errs.HTTPErrorResponse{ErrorMessage: errs.UserDoesNotExist.Error()})
		return
	}

	newProfileData := model.Profile{
		Email:    req.Email,
		Username: req.Username,
		Gender:   req.Gender,
	}

	if err := h.profileManager.UpdateProfile(request.Context(), oldProfileData, newProfileData); err != nil {
		h.log.Warn("[ ProfilesDelivery.UpdateProfileData ] Не удалось обновить данные профиля", slog.String("error", err.Error()))

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