package user

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/user/mocks"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"io"
	"log/slog"
	"testing"
)

var (
	testContextRequestIDValue = uuid.New()
	dbTestError               = errors.New("ошибка базы данных")
)

const testContextRequestIDKey = "request-id"

func TestUsersService_CreateUser(t *testing.T) {
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
			name: "Успешное создание пользователя",
			user: model.User{
				Username: "testuser",
				Email:    "test@example.com",
				Password: "password",
			},
			setupFunc: func(ctrl *gomock.Controller) *UsersService {
				userRepo := mocks.NewMockusersRepository(ctrl)
				producer := mocks.NewMockpromoCodesMessageProducer(ctrl)

				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				salt := "some_salt"
				hashedPassword := utils.HashPassword("password", []byte(salt))

				userRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(model.User{
					ID:       1,
					Username: "testuser",
					Email:    "test@example.com",
					Password: hashedPassword,
				}, nil)

				producer.EXPECT().AddPromoCode(gomock.Any(), gomock.Any(), uint32(utils.AvailPromoTenID)).Return(nil)

				return &UsersService{
					userRepo:       userRepo,
					producer:       producer,
					log:            logger,
					inputValidator: utils.NewInputValidator(),
				}
			},
			want: want{
				user: model.User{
					ID:       1,
					Username: "testuser",
					Email:    "test@example.com",
					Password: "some",
				},
				err: nil,
			},
		},
		{
			name: "Ошибка при создании пользователя",
			user: model.User{
				Username: "testuser",
				Email:    "test@example.com",
				Password: "password",
			},
			setupFunc: func(ctrl *gomock.Controller) *UsersService {
				userRepo := mocks.NewMockusersRepository(ctrl)
				producer := mocks.NewMockpromoCodesMessageProducer(ctrl)

				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				userRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(model.User{}, errors.New("some error"))

				return &UsersService{
					userRepo:       userRepo,
					producer:       producer,
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
			name: "Ошибка при добавлении промо-кода",
			user: model.User{
				Username: "testuser",
				Email:    "test@example.com",
				Password: "password",
			},
			setupFunc: func(ctrl *gomock.Controller) *UsersService {
				userRepo := mocks.NewMockusersRepository(ctrl)
				producer := mocks.NewMockpromoCodesMessageProducer(ctrl)

				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				salt := "some_salt"
				hashedPassword := utils.HashPassword("password", []byte(salt))

				userRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(model.User{
					ID:       1,
					Username: "testuser",
					Email:    "test@example.com",
					Password: hashedPassword,
				}, nil)

				producer.EXPECT().AddPromoCode(gomock.Any(), gomock.Any(), uint32(utils.AvailPromoTenID)).Return(errors.New("some error"))

				return &UsersService{
					userRepo:       userRepo,
					producer:       producer,
					log:            logger,
					inputValidator: utils.NewInputValidator(),
				}
			},
			want: want{
				user: model.User{
					ID:       1,
					Username: "testuser",
					Email:    "test@example.com",
					Password: "hashedPassword",
				},
				err: nil,
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

			user, err := service.CreateUser(ctx, tt.user)

			assert.NotNil(t, tt.want.user, user)
			assert.Equal(t, tt.want.err, err)
		})
	}
}
