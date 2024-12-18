package profile

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/profile/mocks"
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

func TestProfilesService_GetProfile(t *testing.T) {
	type want struct {
		profile model.Profile
		err     error
	}

	var tests = []struct {
		name      string
		id        uint32
		setupFunc func(ctrl *gomock.Controller) *ProfilesService
		want      want
	}{
		{
			name: "Успешное получение профиля",
			id:   1,
			setupFunc: func(ctrl *gomock.Controller) *ProfilesService {
				profileRepo := mocks.NewMockprofileRepository(ctrl)

				profileRepo.EXPECT().GetProfile(gomock.Any(), uint32(1)).Return(model.Profile{
					ID:       1,
					Username: "testuser",
					Email:    "test@example.com",
				}, nil)

				return &ProfilesService{
					profileRepo: profileRepo,
					log:         slog.New(slog.NewTextHandler(io.Discard, nil)),
				}
			},
			want: want{
				profile: model.Profile{
					ID:       1,
					Username: "testuser",
					Email:    "test@example.com",
				},
				err: nil,
			},
		},
		{
			name: "Ошибка при получении профиля",
			id:   1,
			setupFunc: func(ctrl *gomock.Controller) *ProfilesService {
				profileRepo := mocks.NewMockprofileRepository(ctrl)

				profileRepo.EXPECT().GetProfile(gomock.Any(), uint32(1)).Return(model.Profile{}, errors.New("some error"))

				return &ProfilesService{
					profileRepo: profileRepo,
					log:         slog.New(slog.NewTextHandler(io.Discard, nil)),
				}
			},
			want: want{
				profile: model.Profile{},
				err:     errors.New("some error"),
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

			profile, err := service.GetProfile(ctx, tt.id)

			assert.Equal(t, tt.want.profile, profile)
			assert.Equal(t, tt.want.err, err)
		})
	}
}
