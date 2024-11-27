package tests

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/image"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/profile"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/profile/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func setupTest(t *testing.T) (context.Context, *mocks.MockprofileRepository, *slog.Logger, uint32) {
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockprofileRepository(ctrl)
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	var profileID uint32
	err := faker.FakeData(&profileID)
	assert.NoError(t, err)

	return context.Background(), mockRepo, logger, profileID
}

func randomProfile(profileID uint32) model.Profile {
	ageValues, err := faker.RandomInt(18, 100)
	genders := []string{"Мужской", "Женский"}
	randomIndex, _ := faker.RandomInt(0, len(genders)-1)
	gender := genders[randomIndex[0]]

	if err != nil {
		panic("Failed to generate random age") // handle error in test setup
	}
	return model.Profile{
		ID:        profileID,
		Email:     faker.Email(),
		Username:  faker.Username(),
		Gender:    gender,
		Age:       uint8(ageValues[0]),
		AvatarURL: "files/default.jpeg",
	}
}

func TestProfilesService_GetProfile(t *testing.T) {
	ctx, mockRepo, logger, profileID := setupTest(t)
	service := profile.NewProfileService(&image.ImagesUsecase{}, mockRepo, logger)

	expectedData := randomProfile(profileID)

	tests := []struct {
		name            string
		profileID       uint32
		expectedProfile model.Profile
		expectedError   error
		mockSetup       func()
	}{
		{
			name:            "Success",
			profileID:       profileID,
			expectedProfile: expectedData,
			expectedError:   nil,
			mockSetup: func() {
				mockRepo.EXPECT().GetProfile(ctx, profileID).Return(expectedData, nil)
			},
		},
		{
			name:            "Repository error",
			profileID:       profileID,
			expectedProfile: model.Profile{},
			expectedError:   errs.UserDoesNotExist,
			mockSetup: func() {
				mockRepo.EXPECT().GetProfile(ctx, profileID).Return(model.Profile{}, errs.UserDoesNotExist)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			profile, err := service.GetProfile(ctx, tt.profileID)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedProfile, profile)
			}
		})
	}
}

func TestProfilesService_UpdateProfile(t *testing.T) {
	ctx, mockRepo, logger, profileID := setupTest(t)
	service := profile.NewProfileService(nil, mockRepo, logger)

	oldData := randomProfile(profileID)
	newData := model.Profile{
		Email:    faker.Email(),
		Username: faker.Username(),
		Gender:   "Мужской",
	}

	tests := []struct {
		name           string
		oldProfileData model.Profile
		newProfileData model.Profile
		expectedError  error
		mockSetup      func()
	}{
		{
			name:           "Update email and username successfully",
			oldProfileData: oldData,
			newProfileData: newData,
			expectedError:  nil,
			mockSetup: func() {
				mockRepo.EXPECT().UpdateProfile(ctx, profileID, gomock.Any()).Return(nil)
			},
		},
		{
			name:           "Invalid email format",
			oldProfileData: oldData,
			newProfileData: model.Profile{Email: "invalid-email"},
			expectedError:  errs.InvalidEmailFormat,
			mockSetup:      func() {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			err := service.UpdateProfile(ctx, tt.oldProfileData, tt.newProfileData)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestProfilesService_UpdateProfileAvatar(t *testing.T) {
	ctx, mockRepo, logger, profileID := setupTest(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockImageSaver := mocks.NewMockimageSaver(ctrl)
	service := profile.NewProfileService(mockImageSaver, mockRepo, logger)

	tests := []struct {
		name          string
		expectedError error
		mockSetup     func(expectedFilePath string)
	}{
		{
			name:          "Successful avatar update",
			expectedError: nil,
			mockSetup: func(expectedFilePath string) {
				mockImageSaver.EXPECT().SaveImage(gomock.Any(), gomock.Any()).Return(expectedFilePath, nil)
				mockRepo.EXPECT().UpdateProfileAvatar(gomock.Any(), profileID, expectedFilePath).Return(nil)
			},
		},
		{
			name:          "Error saving image",
			expectedError: errs.InternalServerError,
			mockSetup: func(_ string) {
				mockImageSaver.EXPECT().SaveImage(gomock.Any(), gomock.Any()).Return("", errs.InternalServerError)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempFile, err := os.CreateTemp("", "test_avatar_*.jpg")
			assert.NoError(t, err)
			defer os.Remove(tempFile.Name())
			defer tempFile.Close()

			expectedFilePath := "generated/path/to/avatar.jpg"
			tt.mockSetup(expectedFilePath)

			filepath, err := service.UpdateProfileAvatar(ctx, profileID, tempFile)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError.Error())
				assert.Empty(t, filepath)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, expectedFilePath, filepath)
			}
		})
	}
}
