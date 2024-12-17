package sessions

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"testing"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
)

func TestGetSession(t *testing.T) {
	redisClient, mock := redismock.NewClientMock()
	sessionStore := &SessionStore{
		redis: redisClient,
	}

	tests := []struct {
		name          string
		sessionID     string
		mockSetup     func()
		expectedError error
		expectedValue model.Session
	}{
		{
			name:      "Success",
			sessionID: "test-session-id",
			mockSetup: func() {
				mock.ExpectGet("test-session-id").SetVal("123")
			},
			expectedError: nil,
			expectedValue: model.Session{
				UserID:    123,
				SessionID: "test-session-id",
			},
		},
		{
			name:      "Session Not Found",
			sessionID: "test-session-id",
			mockSetup: func() {
				mock.ExpectGet("test-session-id").SetErr(redis.Nil)
			},
			expectedError: errs.SessionNotFound,
			expectedValue: model.Session{},
		},
		{
			name:      "Redis Error",
			sessionID: "test-session-id",
			mockSetup: func() {
				mock.ExpectGet("test-session-id").SetErr(errors.New("redis error"))
			},
			expectedError: errs.InternalServerError,
			expectedValue: model.Session{},
		},
		{
			name:      "Invalid Value Conversion",
			sessionID: "test-session-id",
			mockSetup: func() {
				mock.ExpectGet("test-session-id").SetVal("invalid_user_id")
			},
			expectedError: errs.InternalServerError,
			expectedValue: model.Session{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			result, err := sessionStore.Get(context.Background(), tt.sessionID)

			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedValue, result)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
