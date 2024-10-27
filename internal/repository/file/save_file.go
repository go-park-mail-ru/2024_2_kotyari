package file

import (
	"fmt"
	"os"
	"path/filepath"
)

func (repo *FilesRepo) SaveFile(filename string, file *os.File) (string, error) {
	fullPath := filepath.Join(repo.baseUrl, filename)

	if _, err := os.Stat(fullPath); err == nil {
		return fullPath, nil
	}

	destFile, err := os.Create(fullPath)
	if err != nil {
		return "", fmt.Errorf("ошибка при создании файла: %w", err)
	}

	defer func(destFile *os.File) {
		err = destFile.Close()
		if err != nil {
			repo.log.Error("ошибка закрытия файла")
		}
	}(destFile)

	_, err = file.Seek(0, 0)
	if err != nil {
		return "", err
	}

	if _, err = destFile.ReadFrom(file); err != nil {
		return "", fmt.Errorf("ошибка при копировании содержимого файла: %w", err)
	}

	return fullPath, nil
}
