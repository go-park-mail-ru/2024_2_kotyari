package image

import (
	"github.com/disintegration/imaging"
	"image/gif"
	"os"
)

func (iu *ImagesUsecase) isImageFile(file *os.File) bool {
	img, err := imaging.Decode(file)
	if err != nil {
		return false
	}

	return img != nil
}

func (iu *ImagesUsecase) isGIFFile(file *os.File) bool {
	_, err := gif.Decode(file)
	if err != nil {
		return false
	}
	return true
}
