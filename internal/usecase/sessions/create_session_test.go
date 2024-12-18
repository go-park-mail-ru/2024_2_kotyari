package sessions

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/sessions/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"io"
	"log/slog"
	"testing"
)

func TestSessionService_Create(t *testing.T) {
	type want struct {
		sessionID string
		err       error
	}

	var tests = []struct {
		name      string
		userID    uint32
		setupFunc func(ctrl *gomock.Controller) *SessionService
		want      want
	}{
		{
			name:   "Успешное создание сессии",
			userID: 1,
			setupFunc: func(ctrl *gomock.Controller) *SessionService {
				sessionRepo := mocks.NewMocksessionRepository(ctrl)

				sessionRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return("session_id", nil)

				return &SessionService{
					SessionRepo: sessionRepo,
					log:         slog.New(slog.NewTextHandler(io.Discard, nil)),
				}
			},
			want: want{
				sessionID: "session_id",
				err:       nil,
			},
		},
		{
			name:   "Ошибка при создании сессии",
			userID: 1,
			setupFunc: func(ctrl *gomock.Controller) *SessionService {
				sessionRepo := mocks.NewMocksessionRepository(ctrl)

				sessionRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return("", errors.New("some error"))

				return &SessionService{
					SessionRepo: sessionRepo,
					log:         slog.New(slog.NewTextHandler(io.Discard, nil)),
				}
			},
			want: want{
				sessionID: "",
				err:       errors.New("some error"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.Background()

			service := tt.setupFunc(ctrl)

			sessionID, err := service.Create(ctx, tt.userID)

			assert.Equal(t, tt.want.sessionID, sessionID)
			assert.Equal(t, tt.want.err, err)
		})
	}
}
