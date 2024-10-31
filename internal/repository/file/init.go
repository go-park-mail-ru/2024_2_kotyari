package file

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
)

const baseUrl = "./files"

var ErrFileDoesNotExist = errors.New("file does not exist")
var ErrAccessDenied = errors.New("access denied to file outside allowed directory")

type FilesRepo struct {
	log     *slog.Logger
	baseUrl string
}

func NewFilesRepo(log *slog.Logger) (*FilesRepo, error) {
	projectRoot, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении корня проекта: %w", err)
	}

	fullPath := filepath.Join(projectRoot, baseUrl)

	if _, err = os.Stat(fullPath); os.IsNotExist(err) {
		if err = os.MkdirAll(fullPath, os.ModePerm); err != nil {
			return nil, fmt.Errorf("ошибка при создании директории %s: %w", fullPath, err)
		}
		log.Debug("Директория успешно создана в корне проекта", slog.String("path", fullPath))
	}
	return &FilesRepo{
		log:     log,
		baseUrl: baseUrl,
	}, nil
}
