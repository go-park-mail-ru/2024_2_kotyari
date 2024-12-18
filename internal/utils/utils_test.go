package utils

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCalculateFileHash(t *testing.T) {
	tempFile, err := os.CreateTemp("", "testfile")
	assert.NoError(t, err)
	defer os.Remove(tempFile.Name())

	fileContent := "test content"
	_, err = tempFile.WriteString(fileContent)
	assert.NoError(t, err)

	_, err = tempFile.Seek(0, 0)
	assert.NoError(t, err)

	expectedHash := md5.Sum([]byte(fileContent))

	hash, err := CalculateFileHash(tempFile)
	assert.NoError(t, err)
	assert.Equal(t, hex.EncodeToString(expectedHash[:]), hash)
}

func TestStrToUint32(t *testing.T) {
	result, err := StrToUint32("123")
	assert.NoError(t, err)
	assert.Equal(t, uint32(123), result)

	_, err = StrToUint32("not-a-number")
	assert.Error(t, err)
}

func TestInputValidator_SanitizeString(t *testing.T) {
	validator := NewInputValidator()
	input := "<script>alert(1)</script>"
	expected := ""
	assert.Equal(t, expected, validator.SanitizeString(input))
}

func TestIsExpired(t *testing.T) {
	assert.True(t, IsExpired(time.Now().Add(-time.Hour).Unix()))
	assert.False(t, IsExpired(time.Now().Add(time.Hour).Unix()))
}

func TestWriteJSON(t *testing.T) {
	recorder := httptest.NewRecorder()
	WriteJSON(recorder, http.StatusOK, map[string]string{"key": "value"})

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Header().Get("Content-Type"), "application/json")
	assert.Contains(t, recorder.Body.String(), "key")
}

func TestWriteErrorJSON(t *testing.T) {
	recorder := httptest.NewRecorder()
	WriteErrorJSON(recorder, http.StatusBadRequest, assert.AnError)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "error")
}

func TestGenerateSalt(t *testing.T) {
	salt, err := GenerateSalt()
	assert.NoError(t, err)
	assert.Equal(t, saltLength, len(salt))
}

func TestHashPasswordAndVerifyPassword(t *testing.T) {
	password := "StrongPassword123"
	salt, _ := GenerateSalt()
	hash := HashPassword(password, salt)

	assert.True(t, VerifyPassword(hash, password))
	assert.False(t, VerifyPassword(hash, "WrongPassword"))
}

func TestValidateRegistration(t *testing.T) {
	err := ValidateRegistration("test@example.com", "username", "Password1", "Password1")
	assert.NoError(t, err)

	err = ValidateRegistration("invalid-email", "username", "Password1", "Password1")
	assert.Error(t, err)

	err = ValidateRegistration("test@example.com", "username", "pass", "pass")
	assert.Error(t, err)

	err = ValidateRegistration("test@example.com", "username", "Password1", "DifferentPassword")
	assert.Error(t, err)
}

func TestIsValidEmail(t *testing.T) {
	assert.True(t, IsValidEmail("test@example.com"))
	assert.False(t, IsValidEmail("invalid-email"))
}

func TestIsValidUsername(t *testing.T) {
	assert.True(t, IsValidUsername("username"))
	assert.False(t, IsValidUsername(""))
}

func TestIsValidPassword(t *testing.T) {
	assert.True(t, isValidPassword("StrongPassword1"))
	assert.False(t, isValidPassword("weak"))
}

func TestReturnSortOrderOption(t *testing.T) {
	assert.Equal(t, descSortOrderOption, ReturnSortOrderOption("invalid"))
	assert.Equal(t, ascSortOrderOption, ReturnSortOrderOption(ascSortOrderOption))
}

func TestGetContextRequestID(t *testing.T) {
	ctx := context.WithValue(context.Background(), RequestIDName, uuid.New())

	requestID, err := GetContextRequestID(ctx)
	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, requestID)
}

func TestAddMetadataRequestID(t *testing.T) {
	ctx := context.WithValue(context.Background(), RequestIDName, uuid.New())

	newCtx, err := AddMetadataRequestID(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, newCtx)
}

func TestValidateReviewRating(t *testing.T) {
	assert.True(t, ValidateReviewRating(model.Review{Rating: 3}))
	assert.False(t, ValidateReviewRating(model.Review{Rating: 0}))
	assert.False(t, ValidateReviewRating(model.Review{Rating: 6}))
}

func TestSetSessionCookie(t *testing.T) {
	cookie := SetSessionCookie("test-value")
	assert.Equal(t, "test-value", cookie.Value)
	assert.Equal(t, SessionName, cookie.Name)
}

func TestRemoveSessionCookie(t *testing.T) {
	cookie := RemoveSessionCookie()
	assert.Equal(t, deleteSessionLifetime, cookie.MaxAge)
}
