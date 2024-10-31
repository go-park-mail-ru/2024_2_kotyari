package image

import "os"

func (iu *ImagesUsecase) SaveImage(filename string, file *os.File) (string, error) {
	imagePath, err := iu.files.SaveFile(filename, file, iu.isImageFile)
	if err != nil {
		return "", err
	}

	if imagePath != "" {
		return imagePath, nil
	}

	_, err = iu.files.SaveFile(filename, file, iu.isGIFFile)
	return imagePath, err
}
