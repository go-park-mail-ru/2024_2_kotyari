package profile

import (
	"errors"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
)

// uploadAvatarFromRequest загружает аватар из запроса.
//
// Возвращает путь к сохраненному файлу, сообщение об ошибке, которое будет отправлено в ответе,
// и саму ошибку.
func (pd *ProfilesDelivery) uploadAvatarFromRequest(r *http.Request) (string, error, error) {
	file, header, err := r.FormFile("avatar")
	if err != nil {
		return "", errs.AvatarFileReadError, err
	}

	defer func(f multipart.File) {
		closeErr := f.Close()
		if closeErr != nil {
			pd.log.Error("Ошибка закрытия multipart file", slog.String("error", closeErr.Error()))
		}
	}(file)

	if header.Size > maxUploadSize {
		return "", errs.AvatarFileSizeExceedsLimit, errors.New("[ ProfilesDelivery.UpdateProfileAvatar ] Размер файла превышает 10 МБ")
	}

	tempFile, err := os.CreateTemp("", "avatar-*")
	if err != nil {
		return "", errs.InternalServerError, err
	}

	defer func(name string) {
		removeErr := os.Remove(name)
		if removeErr != nil {
			pd.log.Error("[ ProfilesDelivery.UpdateProfileAvatar ] ошибка удаления temp файла",
				slog.String("error", removeErr.Error()),
			)
		}
	}(tempFile.Name())

	if _, err = io.Copy(tempFile, file); err != nil {
		return "", errs.AvatarUploadError, err
	}

	if _, err = tempFile.Seek(0, 0); err != nil {
		return "", errs.InternalServerError, err
	}

	path, err := pd.imageSaver.SaveImage(tempFile.Name(), tempFile)
	if err != nil {
		return "", errs.AvatarImageSaveError, err
	}

	return path, nil, nil
}
