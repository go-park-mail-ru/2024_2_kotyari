package file

import "os"

type filesRepo interface {
	GetFile(filename string) (*os.File, error)
}

type FilesDelivery struct {
	repo filesRepo
}

func NewFilesDelivery(repo filesRepo) *FilesDelivery {
	return &FilesDelivery{
		repo: repo,
	}
}
