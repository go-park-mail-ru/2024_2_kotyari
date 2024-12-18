package file

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestGetFile(t *testing.T) {
	baseDir := t.TempDir()
	repo := &FilesRepo{
		baseUrl: baseDir,
	}

	tests := []struct {
		name        string
		filename    string
		setup       func() string
		expectedErr error
	}{
		{
			name:     "Success",
			filename: "test-file.txt",
			setup: func() string {
				filePath := filepath.Join(baseDir, "test-file.txt")
				err := os.WriteFile(filePath, []byte("content"), 0644)
				if err != nil {
					t.Fatalf("Ошибка при создании файла для теста: %v", err)
				}
				return filePath
			},
			expectedErr: nil,
		},
		{
			name:        "File Not Found",
			filename:    "non-existent.txt",
			setup:       func() string { return "" },
			expectedErr: ErrFileDoesNotExist,
		},
		{
			name:        "Access Denied",
			filename:    "../unauthorized.txt",
			setup:       func() string { return "" },
			expectedErr: ErrAccessDenied,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			file, err := repo.GetFile(tt.filename)

			if tt.expectedErr != nil {
				assert.Nil(t, file)
				assert.ErrorIs(t, err, tt.expectedErr)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, file)
				_ = file.Close()
			}
		})
	}
}

func TestSaveFile(t *testing.T) {
	baseDir := t.TempDir()
	repo := &FilesRepo{
		baseUrl: baseDir,
	}

	tests := []struct {
		name        string
		filename    string
		setup       func() (*os.File, string)
		expectedErr error
	}{
		{
			name:     "Success",
			filename: "new-file.txt",
			setup: func() (*os.File, string) {
				tempFile, _ := os.CreateTemp("", "temp-file")
				_, _ = tempFile.WriteString("test content")
				return tempFile, filepath.Join(baseDir, "new-file.txt")
			},
			expectedErr: nil,
		},
		{
			name:     "File Already Exists",
			filename: "existing-file.txt",
			setup: func() (*os.File, string) {
				filePath := filepath.Join(baseDir, "existing-file.txt")
				err := os.WriteFile(filePath, []byte("existing content"), 0644)
				assert.NoError(t, err)
				return nil, filePath
			},
			expectedErr: nil,
		},
		{
			name:     "Error Creating File",
			filename: "invalid-path/invalid-file.txt",
			setup: func() (*os.File, string) {
				tempFile, _ := os.CreateTemp("", "temp-file")
				_, _ = tempFile.WriteString("test content")
				return tempFile, ""
			},
			expectedErr: fmt.Errorf("ошибка при создании файла"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputFile, expectedPath := tt.setup()
			defer func() {
				if inputFile != nil {
					_ = inputFile.Close()
				}
			}()

			result, err := repo.SaveFile(tt.filename, inputFile)

			if tt.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, expectedPath, result)
				_, statErr := os.Stat(result)
				assert.NoError(t, statErr)
			}
		})
	}
}
