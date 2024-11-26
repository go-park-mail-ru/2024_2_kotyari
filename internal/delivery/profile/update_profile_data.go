package profile

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (pd *ProfilesDelivery) UpdateProfileData(writer http.ResponseWriter, request *http.Request) {
	userID, ok := utils.GetContextSessionUserID(request.Context())
	if !ok {
		utils.WriteErrorJSON(writer, http.StatusUnauthorized, errs.UserNotAuthorized)
	}

	var req UpdateProfileRequest

	req.Email = pd.stringSanitizer.SanitizeString(req.Email)
	req.Username = pd.stringSanitizer.SanitizeString(req.Username)
	req.Gender = pd.stringSanitizer.SanitizeString(req.Gender)

	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		pd.log.Error("[ ProfilesDelivery.UpdateProfileData ] Ошибка десериализации запроса", slog.String("error", err.Error()))
		utils.WriteJSON(writer, http.StatusBadRequest, &errs.HTTPErrorResponse{ErrorMessage: errs.InvalidJSONFormat.Error()})
		return
	}

	oldProfileData, err := pd.profileManager.GetProfile(request.Context(), uint32(userID))
	if err != nil {
		pd.log.Warn("[ ProfilesDelivery.UpdateProfileData ] Не удалось получить старые данные профиля", slog.String("error", err.Error()))
		utils.WriteJSON(writer, http.StatusNotFound, &errs.HTTPErrorResponse{ErrorMessage: errs.UserDoesNotExist.Error()})
		return
	}

	newProfileData := model.Profile{
		Email:    req.Email,
		Username: req.Username,
		Gender:   req.Gender,
	}

	if err := pd.profileManager.UpdateProfile(request.Context(), oldProfileData, newProfileData); err != nil {
		pd.log.Warn("[ ProfilesDelivery.UpdateProfileData ] Не удалось обновить данные профиля", slog.String("error", err.Error()))

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
