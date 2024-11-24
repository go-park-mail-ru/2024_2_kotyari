package file

import (
	"fmt"
	"log"
	"log/slog"
	"os"
)

func (repo *FilesRepo) GetFile(filename string) (*os.File, error) {
	fullPath, err := repo.buildPath(filename)
	if err != nil {

		log.Printf("[ FilesRepo.GetFile ] ошибка %s", err.Error())
		return nil, err
	}

	file, err := os.Open(fullPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrFileDoesNotExist
		}

		return nil, fmt.Errorf("ошибка при открытии файла: %w", err)
	}

	repo.log.Debug("Файл успешно открыт для чтения", slog.String("path", fullPath))
	return file, nil
}
