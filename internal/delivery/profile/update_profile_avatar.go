package profile

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"io"
	"log/slog"
	"net/http"
	"os"
)

const maxUploadSize = 1024 * 1024 * 10

func (h *ProfilesDelivery) UpdateProfileAvatar(writer http.ResponseWriter, request *http.Request) {

	id := utils.GetContextSessionUserID(request.Context())

	file, header, err := request.FormFile("file")

	if err != nil {
		http.Error(writer, "Не удалось прочитать файл", http.StatusBadRequest)
		return
	}

	defer file.Close()

	if header.Size > maxUploadSize {
		http.Error(writer, "Размер файла превышает 10 МБ", http.StatusBadRequest)
		return
	}

	tempFile, err := os.CreateTemp("", "avatar-*")
	if err != nil {
		http.Error(writer, "Не удалось создать временный файл", http.StatusInternalServerError)
		return
	}
	defer os.Remove(tempFile.Name())

	_, err = io.Copy(tempFile, file)
	if err != nil {
		http.Error(writer, "Не удалось сохранить файл", http.StatusInternalServerError)
		return
	}

	if _, err = tempFile.Seek(0, 0); err != nil {
		http.Error(writer, "Не удалось сбросить указатель файла", http.StatusInternalServerError)
		return
	}

	err = h.profileManager.UpdateProfileAvatar(request.Context(), id, tempFile)

	if err != nil {
		h.log.Error("[ ProfilesDelivery.UpdateProfileAvatar ]", slog.String("error", err.Error()))

		return
	}

}
