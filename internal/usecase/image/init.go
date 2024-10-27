package image

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/file"
	"os"
)

type filesUsecase interface {
	SaveFile(filename string, file *os.File, checkFile file.CheckFileFunc) (string, error)
}

type ImagesUsecase struct {
	files filesUsecase
}

func NewImagesUsecase(files filesUsecase) *ImagesUsecase {
	return &ImagesUsecase{
		files: files,
	}
}
