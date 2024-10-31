package file

import (
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"os"
	"path/filepath"
)

type CheckFileFunc func(*os.File) bool

var ErrFileTypeNotAllowed = errors.New("file type not allowed")

func (fu *FilesUsecase) SaveFile(filename string, file *os.File, checkFile CheckFileFunc) (string, error) {
	if !checkFile(file) {
		return "", ErrFileTypeNotAllowed
	}

	hash, err := utils.CalculateFileHash(file)
	if err != nil {
		return "", fmt.Errorf("ошибка при вычислении хэша файла: %w", err)
	}

	hashedFilename := fmt.Sprintf("%s%s", hash, filepath.Ext(filename))

	filePath, err := fu.repo.SaveFile(hashedFilename, file)
	if err != nil {
		return "", fmt.Errorf("ошибка при сохранении файла: %w", err)
	}

	return filePath, nil
}
