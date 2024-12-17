package sessions

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/go-redis/redismock/v9"
	"testing"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateSession(t *testing.T) {
	redisClient, mock := redismock.NewClientMock()
	sessionStore := &SessionStore{
		redis: redisClient,
	}

	tests := []struct {
		name        string
		session     model.Session
		mockSetup   func()
		expectedID  string
		expectedErr error
	}{
		{
			name: "Success",
			session: model.Session{
				SessionID: "test-session-id",
				UserID:    123,
			},
			mockSetup: func() {
				mock.ExpectSet("test-session-id", "123", utils.DefaultSessionLifetime).SetVal("OK")
			},
			expectedID:  "test-session-id",
			expectedErr: nil,
		},
		{
			name: "Redis error",
			session: model.Session{
				SessionID: "test-session-id",
				UserID:    123,
			},
			mockSetup: func() {
				mock.ExpectSet("test-session-id", "123", utils.DefaultSessionLifetime).SetErr(errs.SessionCreationError)
			},
			expectedID:  "",
			expectedErr: errs.SessionCreationError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			id, err := sessionStore.Create(context.Background(), tt.session)

			assert.Equal(t, tt.expectedID, id)
			assert.Equal(t, tt.expectedErr, err)
			assert.NoError(t, mock.ExpectationsWereMet()) // Проверка всех моков
		})
	}
}
