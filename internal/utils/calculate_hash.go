package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func CalculateFileHash(file *os.File) (string, error) {
	_, err := file.Seek(0, 0)
	if err != nil {
		return "", err
	}

	hash := md5.New()
	if _, err = io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
