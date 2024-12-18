package user

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/user/mocks"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"io"
	"log/slog"
	"testing"
)

func TestUsersService_LoginUser(t *testing.T) {
	type want struct {
		user model.User
		err  error
	}

	var tests = []struct {
		name      string
		user      model.User
		setupFunc func(ctrl *gomock.Controller) *UsersService
		want      want
	}{
		{
			name: "Успешный вход",
			user: model.User{
				Username: "testuser",
				Email:    "test@example.com",
				Password: "password",
			},
			setupFunc: func(ctrl *gomock.Controller) *UsersService {
				userRepo := mocks.NewMockusersRepository(ctrl)

				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				dbUser := model.User{
					ID:       1,
					Username: "testuser",
					Email:    "test@example.com",
					Password: utils.HashPassword("password", []byte("some_salt")),
				}

				userRepo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(dbUser, nil)

				return &UsersService{
					userRepo:       userRepo,
					log:            logger,
					inputValidator: utils.NewInputValidator(),
				}
			},
			want: want{
				user: model.User{
					ID:       1,
					Username: "testuser",
					Email:    "test@example.com",
					Password: utils.HashPassword("password", []byte("some_salt")),
				},
				err: errs.WrongCredentials,
			},
		},
		{
			name: "Не найден юзер",
			user: model.User{
				Username: "testuser",
				Email:    "test@example.com",
				Password: "password",
			},
			setupFunc: func(ctrl *gomock.Controller) *UsersService {
				userRepo := mocks.NewMockusersRepository(ctrl)

				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				userRepo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(model.User{}, errors.New("some error"))

				return &UsersService{
					userRepo:       userRepo,
					log:            logger,
					inputValidator: utils.NewInputValidator(),
				}
			},
			want: want{
				user: model.User{},
				err:  errors.New("some error"),
			},
		},
		{
			name: "Неправильный пароль",
			user: model.User{
				Username: "testuser",
				Email:    "test@example.com",
				Password: "wrong_password",
			},
			setupFunc: func(ctrl *gomock.Controller) *UsersService {
				userRepo := mocks.NewMockusersRepository(ctrl)

				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				dbUser := model.User{
					ID:       1,
					Username: "testuser",
					Email:    "test@example.com",
					Password: utils.HashPassword("password", []byte("some_salt")),
				}

				userRepo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(dbUser, nil)

				return &UsersService{
					userRepo:       userRepo,
					log:            logger,
					inputValidator: utils.NewInputValidator(),
				}
			},
			want: want{
				user: model.User{},
				err:  errs.WrongCredentials,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.WithValue(context.Background(), testContextRequestIDKey, testContextRequestIDValue)

			service := tt.setupFunc(ctrl)

			user, err := service.LoginUser(ctx, tt.user)

			assert.NotNil(t, tt.want.user, user)
			assert.Equal(t, tt.want.err, err)
		})
	}
}
