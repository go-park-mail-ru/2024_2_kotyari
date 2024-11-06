package sessions

import (
	"errors"
	"testing"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/sessions/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var dbTestError = errors.New("ошибка базы данных")

func TestSessionService_Create(t *testing.T) {
	t.Parallel()

	type want error

	tests := []struct {
		name      string
		userID    uint32
		setupFunc func(ctrl *gomock.Controller) *SessionService
		want      want
	}{
		{
			name:   "Сессия успешно создана",
			userID: 1,
			setupFunc: func(ctrl *gomock.Controller) *SessionService {
				sessionRepositoryMock := mocks.NewMocksessionRepository(ctrl)

				sessionRepositoryMock.EXPECT().Create(
					gomock.Any(),
					gomock.Any()).Return("3cf7ddc3-0b02-4da1-b878-9a3ca5771e62", nil)

				return &SessionService{
					SessionRepo: sessionRepositoryMock,
				}
			},
			want: nil,
		},
		{
			name:   "Ошибка создания сессии",
			userID: 1,
			setupFunc: func(ctrl *gomock.Controller) *SessionService {
				sessionRepositoryMock := mocks.NewMocksessionRepository(ctrl)

				sessionRepositoryMock.EXPECT().Create(
					gomock.Any(),
					gomock.Any()).Return("", dbTestError)

				return &SessionService{
					SessionRepo: sessionRepositoryMock,
				}
			},
			want: dbTestError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			_, err := tt.setupFunc(ctrl).Create(nil, tt.userID)

			assert.Equal(t, tt.want, err)
		})
	}
}
