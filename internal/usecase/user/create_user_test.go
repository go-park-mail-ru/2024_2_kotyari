package user

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/user/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

var testDBError = errors.New("ошибка базы данных")

func TestUsersService_CreateUser(t *testing.T) {
	t.Parallel()

	type want error

	tests := []struct {
		name      string
		user      model.User
		setupFunc func(ctrl *gomock.Controller) *UsersService
		want      want
	}{
		{
			name: "Создание пользователя",
			user: model.User{
				Email:    "test@test.com",
				Password: "Password123@",
				Username: "test",
			},
			setupFunc: func(ctrl *gomock.Controller) *UsersService {
				usersRepositoryMock := mocks.NewMockusersRepository(ctrl)
				sessionCreatorMock := mocks.NewMocksessionCreator(ctrl)

				usersRepositoryMock.EXPECT().CreateUser(
					gomock.Any(),
					gomock.Any()).DoAndReturn(func(ctx context.Context, user model.User) (model.User, error) {
					return model.User{
						ID:       1,
						Email:    user.Email,
						Username: user.Username,
						Password: user.Password,
					}, nil
				})

				sessionCreatorMock.EXPECT().Create(
					gomock.Any(),
					uint32(1)).Return("session-id", nil)

				return &UsersService{
					userRepo:       usersRepositoryMock,
					sessionCreator: sessionCreatorMock,
				}
			},
			want: nil,
		},
		{
			name: "Ошибка создания пользователя",
			user: model.User{
				Email:    "test@test.com",
				Username: "test",
			},
			setupFunc: func(ctrl *gomock.Controller) *UsersService {
				usersRepositoryMock := mocks.NewMockusersRepository(ctrl)

				usersRepositoryMock.EXPECT().CreateUser(
					gomock.Any(),
					gomock.Any()).DoAndReturn(func(ctx context.Context, user model.User) (model.User, error) {
					return model.User{}, testDBError
				})

				return &UsersService{
					userRepo: usersRepositoryMock,
				}
			},
			want: testDBError,
		},
		{
			name: "Ошибка создания сессии",
			user: model.User{
				Email:    "test@test.com",
				Password: "Password123@",
				Username: "test",
			},
			setupFunc: func(ctrl *gomock.Controller) *UsersService {
				usersRepositoryMock := mocks.NewMockusersRepository(ctrl)
				sessionCreatorMock := mocks.NewMocksessionCreator(ctrl)

				usersRepositoryMock.EXPECT().CreateUser(
					gomock.Any(),
					gomock.Any()).DoAndReturn(func(ctx context.Context, user model.User) (model.User, error) {
					return model.User{
						ID:       1,
						Email:    user.Email,
						Username: user.Username,
						Password: user.Password,
					}, nil
				})

				sessionCreatorMock.EXPECT().Create(
					gomock.Any(),
					uint32(1)).Return("", testDBError)

				return &UsersService{
					userRepo:       usersRepositoryMock,
					sessionCreator: sessionCreatorMock,
				}
			},
			want: testDBError,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sessionID, dbUser, err := tt.setupFunc(ctrl).CreateUser(context.Background(), tt.user)

			if tt.want != nil {
				assert.Equal(t, tt.want, err)
				assert.Empty(t, sessionID)
				assert.Equal(t, model.User{}, dbUser)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, sessionID)
				assert.Equal(t, uint32(1), dbUser.ID)
				assert.Equal(t, tt.user.Email, dbUser.Email)
				assert.Equal(t, tt.user.Username, dbUser.Username)
				assert.NotEqual(t, tt.user.Password, dbUser.Password)
			}
		})
	}
}
