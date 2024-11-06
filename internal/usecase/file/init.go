package file

import (
	"log/slog"
	"os"
)

type filesRepo interface {
	GetFile(filename string) (*os.File, error)
	SaveFile(filename string, file *os.File) (string, error)
}

type FilesUsecase struct {
	repo filesRepo
	log  *slog.Logger
}

func NewFilesUsecase(filesRepo filesRepo, log *slog.Logger) *FilesUsecase {
	return &FilesUsecase{
		repo: filesRepo,
		log:  log,
	}
}
