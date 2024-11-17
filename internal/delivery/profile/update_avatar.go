package profile

import (
	"errors"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"os"
)

// TODO : вынести
func (pd *ProfilesDelivery) updateProfileAvatar(r *http.Request) (string, error, error) {
	file, header, err := r.FormFile("avatar")
	if err != nil {
		return "", errors.New("не удалось прочитать файл, попробуйте позже"), err
	}

	defer func(f multipart.File) {
		closeErr := f.Close()
		if closeErr != nil {
			pd.log.Error("Ошибка закрытия multipart file", slog.String("error", closeErr.Error()))
		}
	}(file)

	if header.Size > maxUploadSize {
		return "", errors.New("[ ProfilesDelivery.UpdateProfileAvatar ] Размер файла превышает 10 МБ"),
			errors.New("[ ProfilesDelivery.UpdateProfileAvatar ] Размер файла превышает 10 МБ")
	}

	tempFile, err := os.CreateTemp("", "avatar-*")
	if err != nil {
		return "", errors.New("внутренняя ошибка сервера, попробуйте позже"), err
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
		return "", errors.New("внутренняя ошибка сервера, попробуйте позже"), err
	}

	if _, err = tempFile.Seek(0, 0); err != nil {
		return "", errors.New("внутренняя ошибка сервера, попробуйте позже"), err
	}

	path, err := pd.imageSaver.SaveImage(tempFile.Name(), tempFile)
	if err != nil {
		return "", errors.New("не удалось загрузить фотографию"), err
	}

	return path, nil, nil
}
