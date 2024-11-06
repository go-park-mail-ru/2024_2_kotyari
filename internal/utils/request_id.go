package utils

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/google/uuid"
)

const RequestIDName = "request-id"

func GetContextRequestID(ctx context.Context) (uuid.UUID, error) {
	requestID, ok := ctx.Value(RequestIDName).(uuid.UUID)
	if !ok {
		return [16]byte{}, errs.RequestIDNotFound
	}

	return requestID, nil
}
