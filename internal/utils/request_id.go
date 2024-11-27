package utils

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
)

const RequestIDName = "request-id"

func GetContextRequestID(ctx context.Context) (uuid.UUID, error) {
	if requestID, ok := ctx.Value(RequestIDName).(uuid.UUID); ok {
		return requestID, nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		requestIDs := md.Get(RequestIDName)
		requestID, err := uuid.Parse(requestIDs[0])
		if err != nil {
			slog.Error("[GetContextRequestID] Failed to parse requestID from metadata")

			return [16]byte{}, errs.RequestIDNotFound
		}

		return requestID, nil
	}

	return [16]byte{}, errs.RequestIDNotFound
}

func AddMetadataRequestID(ctx context.Context) (context.Context, error) {
	requestID, err := GetContextRequestID(ctx)
	if err != nil {
		return nil, err
	}

	newCTX := metadata.AppendToOutgoingContext(ctx, RequestIDName, requestID.String())

	return newCTX, nil
}
