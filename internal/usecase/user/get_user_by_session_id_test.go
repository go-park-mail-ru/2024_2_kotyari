package user

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/user/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestUsersService_GetUserBySessionID(t *testing.T) {
	t.Parallel()

	type want struct {
		user model.User
		err  error
	}

	tests := []struct {
		name      string
		sessionID string
		setupFunc func(ctrl *gomock.Controller) *UsersService
		want      want
	}{
		{
			name:      "Успешное получение пользователя по сессии",
			sessionID: "session-id",
			setupFunc: func(ctrl *gomock.Controller) *UsersService {
				sessionGetterMock := mocks.NewMocksessionGetter(ctrl)
				usersRepositoryMock := mocks.NewMockusersRepository(ctrl)

				sessionGetterMock.EXPECT().Get(
					gomock.Any(),
					"session-id").Return(model.Session{
					UserID: 1,
				}, nil)

				usersRepositoryMock.EXPECT().GetUserByUserID(
					gomock.Any(),
					uint32(1)).Return(model.User{

					ID: 1,
				}, nil)

				return &UsersService{
					userRepo:      usersRepositoryMock,
					sessionGetter: sessionGetterMock,
				}
			},
			want: want{
				user: model.User{
					ID: 1,
				},
				err: nil,
			},
		},
		{
			name:      "Ошибка получения сессии",
			sessionID: "session-id",
			setupFunc: func(ctrl *gomock.Controller) *UsersService {
				sessionGetterMock := mocks.NewMocksessionGetter(ctrl)
				usersRepositoryMock := mocks.NewMockusersRepository(ctrl)

				sessionGetterMock.EXPECT().Get(
					gomock.Any(),
					"session-id").Return(model.Session{}, testDBError)

				return &UsersService{
					userRepo:      usersRepositoryMock,
					sessionGetter: sessionGetterMock,
				}
			},
			want: want{
				user: model.User{},
				err:  testDBError,
			},
		},
		{
			name:      "Ошибка получения пользователя",
			sessionID: "session-id",
			setupFunc: func(ctrl *gomock.Controller) *UsersService {
				sessionGetterMock := mocks.NewMocksessionGetter(ctrl)
				usersRepositoryMock := mocks.NewMockusersRepository(ctrl)

				sessionGetterMock.EXPECT().Get(
					gomock.Any(),
					"session-id").Return(model.Session{
					UserID: 1,
				}, nil)

				usersRepositoryMock.EXPECT().GetUserByUserID(
					gomock.Any(),
					uint32(1)).Return(model.User{}, testDBError)

				return &UsersService{
					userRepo:      usersRepositoryMock,
					sessionGetter: sessionGetterMock,
				}
			},
			want: want{
				user: model.User{},
				err:  testDBError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			user, err := tt.setupFunc(ctrl).GetUserBySessionID(nil, tt.sessionID)

			assert.Equal(t, tt.want.user, user)
			assert.Equal(t, tt.want.err, err)
		})
	}
}
