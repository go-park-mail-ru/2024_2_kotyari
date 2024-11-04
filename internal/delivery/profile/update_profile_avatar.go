package profile

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"io"
	"log/slog"
	"net/http"
	"os"
)

const maxUploadSize = 1024 * 1024 * 10

func (h *ProfilesDelivery) UpdateProfileAvatar(writer http.ResponseWriter, request *http.Request) {
	userID, ok := utils.GetContextSessionUserID(request.Context())
	if !ok {
		utils.WriteErrorJSON(writer, http.StatusUnauthorized, errs.UserNotAuthorized)
	}

	file, header, err := request.FormFile("avatar")

	if err != nil {
		h.log.Error("[ ProfilesDelivery.UpdateProfileAvatar ] Не удалось прочитать файл", slog.String("error", err.Error()))
		http.Error(writer, "Не удалось прочитать файл", http.StatusBadRequest)
		return
	}

	defer file.Close()

	if header.Size > maxUploadSize {
		h.log.Error("[ ProfilesDelivery.UpdateProfileAvatar ] Размер файла превышает 10 МБ", slog.String("error", err.Error()))
		http.Error(writer, "Размер файла превышает 10 МБ", http.StatusBadRequest)
		return
	}

	tempFile, err := os.CreateTemp("", "avatar-*")
	if err != nil {
		h.log.Error("[ ProfilesDelivery.UpdateProfileAvatar ] Не удалось создать временный файл", slog.String("error", err.Error()))
		http.Error(writer, "Не удалось создать временный файл", http.StatusInternalServerError)
		return
	}
	defer os.Remove(tempFile.Name())

	_, err = io.Copy(tempFile, file)
	if err != nil {
		h.log.Error("[ ProfilesDelivery.UpdateProfileAvatar ]Не удалось сохранить файл", slog.String("error", err.Error()))
		http.Error(writer, "Не удалось сохранить файл", http.StatusInternalServerError)
		return
	}

	if _, err = tempFile.Seek(0, 0); err != nil {
		h.log.Error("[ ProfilesDelivery.UpdateProfileAvatar ] Не удалось сбросить указатель файла", slog.String("error", err.Error()))
		http.Error(writer, "Не удалось сбросить указатель файла", http.StatusInternalServerError)
		return
	}

	filepath, err := h.profileManager.UpdateProfileAvatar(request.Context(), userID, tempFile)

	if err != nil {
		h.log.Error("[ ProfilesDelivery.UpdateProfileAvatar ]", slog.String("error", err.Error()))
		http.Error(writer, "Не удалось обновить аватар профиля", http.StatusInternalServerError)
		return
	}

	avatarResponse := AvatarResponse{
		AvatarUrl: filepath,
	}

	utils.WriteJSON(writer, http.StatusOK, avatarResponse)
}
