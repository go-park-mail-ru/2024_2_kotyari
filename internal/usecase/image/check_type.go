package image

import (
	"image/gif"
	"os"

	"github.com/disintegration/imaging"
)

func (iu *ImagesUsecase) isImageFile(file *os.File) bool {
	img, err := imaging.Decode(file)
	if err != nil {
		return false
	}

	return img != nil
}

func (iu *ImagesUsecase) isGIFFile(file *os.File) bool {
	gifImage, err := gif.Decode(file)
	if err != nil {
		return false
	}

	return gifImage != nil
}
