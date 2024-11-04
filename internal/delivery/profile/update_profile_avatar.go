package profile

import (
	"errors"
	"fmt"
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
		utils.WriteErrorJSON(writer, http.StatusBadRequest,
			errors.New("не удалось прочитать файл, попробуйте позже"))
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
		utils.WriteErrorJSON(writer, http.StatusInternalServerError, errors.New("внутренняя ошибка сервера, попробуйте позже"))

		return
	}
	defer func(name string) {
		err = os.Remove(name)
		if err != nil {
			fmt.Printf("ошибка удаления temp, %s", err.Error())
		}
	}(tempFile.Name())

	fmt.Println("tempFile", tempFile.Name())

	_, err = io.Copy(tempFile, file)
	if err != nil {
		h.log.Error("[ ProfilesDelivery.UpdateProfileAvatar ]Не удалось сохранить файл", slog.String("error", err.Error()))
		utils.WriteErrorJSON(writer, http.StatusInternalServerError, errors.New("внутренняя ошибка сервера, попробуйте позже"))

		return
	}

	if _, err = tempFile.Seek(0, 0); err != nil {
		h.log.Error("[ ProfilesDelivery.UpdateProfileAvatar ] Не удалось сбросить указатель файла", slog.String("error", err.Error()))
		utils.WriteErrorJSON(writer, http.StatusInternalServerError, errors.New("внутренняя ошибка сервера, попробуйте позже"))

		return
	}

	filepath, err := h.profileManager.UpdateProfileAvatar(request.Context(), userID, tempFile)
	if err != nil {
		h.log.Error("[ ProfilesDelivery.UpdateProfileAvatar ]", slog.String("error", err.Error()))
		if errors.Is(err, errs.ErrFileTypeNotAllowed) {
			utils.WriteErrorJSON(writer, http.StatusBadRequest, errs.ErrFileTypeNotAllowed)

			return
		}

		utils.WriteErrorJSON(writer, http.StatusInternalServerError, errors.New("не удалось обновить аватар профиля, попробуйте позже"))
		return
	}

	avatarResponse := AvatarResponse{
		AvatarUrl: filepath,
	}

	utils.WriteJSON(writer, http.StatusOK, avatarResponse)
}
