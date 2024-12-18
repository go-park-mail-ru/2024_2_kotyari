package sessions

import (
	"context"
	"testing"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
)

func TestDeleteSession(t *testing.T) {
	redisClient, mock := redismock.NewClientMock()
	sessionStore := &SessionStore{
		redis: redisClient,
	}

	tests := []struct {
		name        string
		session     model.Session
		mockSetup   func()
		expectedErr error
	}{
		{
			name: "Success",
			session: model.Session{
				SessionID: "test-session-id",
			},
			mockSetup: func() {
				mock.ExpectDel("test-session-id").SetVal(1)
			},
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			err := sessionStore.Delete(context.Background(), tt.session)

			assert.Equal(t, tt.expectedErr, err)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
