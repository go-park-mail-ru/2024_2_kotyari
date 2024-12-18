package profile

//
//import (
//	"context"
//	"errors"
//	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
//	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
//	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/profile/mocks"
//	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
//	"github.com/stretchr/testify/assert"
//	"go.uber.org/mock/gomock"
//	"io"
//	"log/slog"
//	"testing"
//)
//
//func TestProfilesService_ChangePassword(t *testing.T) {
//	type want struct {
//		err error
//	}
//
//	var tests = []struct {
//		name           string
//		userId         uint32
//		oldPassword    string
//		newPassword    string
//		repeatPassword string
//		setupFunc      func(ctrl *gomock.Controller) *ProfilesService
//		want           want
//	}{
//		{
//			name:           "Успешное изменение пароля",
//			userId:         1,
//			oldPassword:    "oldpassword",
//			newPassword:    "newpassword",
//			repeatPassword: "newpassword",
//			setupFunc: func(ctrl *gomock.Controller) *ProfilesService {
//				userRepo := mocks.NewMockuserStore(ctrl)
//
//				userRepo.EXPECT().GetUserByUserID(gomock.Any(), uint32(1)).Return(model.User{
//					ID:       1,
//					Username: "testuser",
//					Email:    "test@example.com",
//					Password: "oldpassword",
//				}, nil)
//
//				userRepo.EXPECT().GetUserByEmail(gomock.Any(), model.User{
//					ID:       1,
//					Username: "testuser",
//					Email:    "test@example.com",
//					Password: "oldpassword",
//				}).Return(model.User{
//					ID:       1,
//					Username: "testuser",
//					Email:    "test@example.com",
//					Password: "oldpassword",
//				}, nil)
//
//				salt := "some_big_salt_value"
//				hashedPassword := utils.HashPassword("newpassword", []byte(salt))
//
//				userRepo.EXPECT().ChangePassword(gomock.Any(), uint32(1), hashedPassword).Return(nil)
//
//				return &ProfilesService{
//					userRepo: userRepo,
//					log:      slog.New(slog.NewTextHandler(io.Discard, nil)),
//				}
//			},
//			want: want{
//				err: nil,
//			},
//		},
//		{
//			name:           "Ошибка при изменении пароля",
//			userId:         1,
//			oldPassword:    "oldpassword",
//			newPassword:    "newpassword",
//			repeatPassword: "newpassword",
//			setupFunc: func(ctrl *gomock.Controller) *ProfilesService {
//				userRepo := mocks.NewMockuserStore(ctrl)
//
//				userRepo.EXPECT().GetUserByUserID(gomock.Any(), uint32(1)).Return(model.User{
//					ID:       1,
//					Username: "testuser",
//					Email:    "test@example.com",
//					Password: "oldpassword",
//				}, nil)
//
//				userRepo.EXPECT().GetUserByEmail(gomock.Any(), model.User{
//					ID:       1,
//					Username: "testuser",
//					Email:    "test@example.com",
//					Password: "oldpassword",
//				}).Return(model.User{
//					ID:       1,
//					Username: "testuser",
//					Email:    "test@example.com",
//					Password: "oldpassword",
//				}, nil)
//
//				salt := "some_big_salt_value"
//				hashedPassword := utils.HashPassword("newpassword", []byte(salt))
//
//				userRepo.EXPECT().ChangePassword(gomock.Any(), uint32(1), hashedPassword).Return(errors.New("some error"))
//
//				return &ProfilesService{
//					userRepo: userRepo,
//					log:      slog.New(slog.NewTextHandler(io.Discard, nil)),
//				}
//			},
//			want: want{
//				err: errors.New("some error"),
//			},
//		},
//		{
//			name:           "Неправильный старый пароль",
//			userId:         1,
//			oldPassword:    "wrongpassword",
//			newPassword:    "newpassword",
//			repeatPassword: "newpassword",
//			setupFunc: func(ctrl *gomock.Controller) *ProfilesService {
//				userRepo := mocks.NewMockuserStore(ctrl)
//
//				userRepo.EXPECT().GetUserByUserID(gomock.Any(), uint32(1)).Return(model.User{
//					ID:       1,
//					Username: "testuser",
//					Email:    "test@example.com",
//					Password: "oldpassword",
//				}, nil)
//
//				userRepo.EXPECT().GetUserByEmail(gomock.Any(), model.User{
//					ID:       1,
//					Username: "testuser",
//					Email:    "test@example.com",
//					Password: "oldpassword",
//				}).Return(model.User{
//					ID:       1,
//					Username: "testuser",
//					Email:    "test@example.com",
//					Password: "oldpassword",
//				}, nil)
//
//				salt := "some_big_salt_value"
//				_ = utils.HashPassword("newpassword", []byte(salt))
//
//				return &ProfilesService{
//					userRepo: userRepo,
//					log:      slog.New(slog.NewTextHandler(io.Discard, nil)),
//				}
//			},
//			want: want{
//				err: errs.WrongPassword,
//			},
//		},
//		{
//			name:           "Неправильный повторный пароль",
//			userId:         1,
//			oldPassword:    "oldpassword",
//			newPassword:    "newpassword",
//			repeatPassword: "wrongpassword",
//			setupFunc: func(ctrl *gomock.Controller) *ProfilesService {
//				userRepo := mocks.NewMockuserStore(ctrl)
//
//				userRepo.EXPECT().GetUserByUserID(gomock.Any(), uint32(1)).Return(model.User{
//					ID:       1,
//					Username: "testuser",
//					Email:    "test@example.com",
//					Password: "oldpassword",
//				}, nil)
//
//				userRepo.EXPECT().GetUserByEmail(gomock.Any(), model.User{
//					ID:       1,
//					Username: "testuser",
//					Email:    "test@example.com",
//					Password: "oldpassword",
//				}).Return(model.User{
//					ID:       1,
//					Username: "testuser",
//					Email:    "test@example.com",
//					Password: "oldpassword",
//				}, nil)
//
//				return &ProfilesService{
//					userRepo: userRepo,
//					log:      slog.New(slog.NewTextHandler(io.Discard, nil)),
//				}
//			},
//			want: want{
//				err: errs.PasswordsDoNotMatch,
//			},
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			t.Parallel()
//
//			ctrl := gomock.NewController(t)
//			defer ctrl.Finish()
//
//			ctx := context.WithValue(context.Background(), testContextRequestIDKey, testContextRequestIDValue)
//
//			service := tt.setupFunc(ctrl)
//
//			err := service.ChangePassword(ctx, tt.userId, tt.oldPassword, tt.newPassword, tt.repeatPassword)
//
//			assert.Equal(t, tt.want.err, err)
//		})
//	}
//}
