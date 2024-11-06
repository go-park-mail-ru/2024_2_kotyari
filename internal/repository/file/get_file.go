package file

import (
	"fmt"
	"os"
)

func (repo *FilesRepo) GetFile(filename string) (*os.File, error) {
	fullPath, err := repo.buildPath(filename)
	if err != nil {

		return nil, err
	}

	file, err := os.Open(fullPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrFileDoesNotExist
		}

		return nil, fmt.Errorf("ошибка при открытии файла: %w", err)
	}

	return file, nil
}
