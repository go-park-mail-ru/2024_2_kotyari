package file

import (
	"path/filepath"
	"strings"
)

func (repo *FilesRepo) buildPath(filename string) (string, error) {
	fullPath := filepath.Join(repo.baseUrl, filename)

	if !strings.HasPrefix(fullPath, repo.baseUrl) {
		return "", ErrAccessDenied
	}

	return fullPath, nil
}
