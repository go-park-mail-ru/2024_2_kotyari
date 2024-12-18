package profile

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/profile/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"log/slog"
	"testing"
)

func TestProfilesService_UpdateProfile(t *testing.T) {
	type want struct {
		err error
	}

	var tests = []struct {
		name       string
		oldProfile model.Profile
		newProfile model.Profile
		setupFunc  func(ctrl *gomock.Controller) *ProfilesService
		want       want
	}{
		{
			name: "Успешное обновление профиля",
			oldProfile: model.Profile{
				ID:       1,
				Username: "testuser",
				Email:    "test@example.com",
			},
			newProfile: model.Profile{
				Username: "newtestuser",
				Email:    "newtest@example.com",
			},
			setupFunc: func(ctrl *gomock.Controller) *ProfilesService {
				profileRepo := mocks.NewMockprofileRepository(ctrl)

				profileRepo.EXPECT().UpdateProfile(gomock.Any(), uint32(1), model.Profile{
					ID:       1,
					Username: "newtestuser",
					Email:    "newtest@example.com",
				}).Return(nil)

				return &ProfilesService{
					profileRepo: profileRepo,
					log:         slog.New(slog.NewTextHandler(io.Discard, nil)),
				}
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "Ошибка при обновлении профиля",
			oldProfile: model.Profile{
				ID:       1,
				Username: "testuser",
				Email:    "test@example.com",
			},
			newProfile: model.Profile{
				Username: "newtestuser",
				Email:    "newtest@example.com",
			},
			setupFunc: func(ctrl *gomock.Controller) *ProfilesService {
				profileRepo := mocks.NewMockprofileRepository(ctrl)

				profileRepo.EXPECT().UpdateProfile(gomock.Any(), uint32(1), model.Profile{
					ID:       1,
					Username: "newtestuser",
					Email:    "newtest@example.com",
				}).Return(errors.New("some error"))

				return &ProfilesService{
					profileRepo: profileRepo,
					log:         slog.New(slog.NewTextHandler(io.Discard, nil)),
				}
			},
			want: want{
				err: errors.New("some error"),
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

			err := service.UpdateProfile(ctx, tt.oldProfile, tt.newProfile)

			assert.Equal(t, tt.want.err, err)
		})
	}
}
