package tests

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"testing"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/address"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/address/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAddressService_GetAddressByProfileID(t *testing.T) {
	tests := []struct {
		name            string
		userID          uint32
		expectedAddress model.Addresses
		expectedError   error
		mockSetup       func(mockRepo *mocks.MockaddressRepository)
	}{
		{
			name:   "Success",
			userID: 2,
			expectedAddress: model.Addresses{
				Id:     2,
				City:   "Казань",
				Street: "Ломоносова",
				House:  "1",
				Flat:   func() *string { s := ""; return &s }(),
			},
			expectedError: nil,
			mockSetup: func(mockRepo *mocks.MockaddressRepository) {
				mockRepo.EXPECT().GetAddressByProfileID(gomock.Any(), uint32(2)).Return(model.Addresses{
					Id:     2,
					City:   "Казань",
					Street: "Ломоносова",
					House:  "1",
					Flat:   nil,
				}, nil)
			},
		},
		{
			name:            "Addresses not found",
			userID:          7,
			expectedAddress: model.Addresses{Id: 0, City: "", Street: "", House: "", Flat: func() *string { s := ""; return &s }()},
			expectedError:   nil,
			mockSetup: func(mockRepo *mocks.MockaddressRepository) {
				mockRepo.EXPECT().GetAddressByProfileID(gomock.Any(), uint32(7)).Return(model.Addresses{
					Id:     0,
					City:   "",
					Street: "",
					House:  "",
					Flat:   new(string),
				}, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mocks.NewMockaddressRepository(ctrl)
			logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

			tt.mockSetup(mockRepo)

			service := address.NewAddressService(mockRepo, logger)

			address, err := service.GetAddressByProfileID(context.Background(), tt.userID)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedAddress, address)
			}
		})
	}
}

func TestAddressService_UpdateUsersAddress(t *testing.T) {
	tests := []struct {
		name          string
		addressID     uint32
		newAddress    model.Addresses
		expectedError error
		mockSetup     func(mockRepo *mocks.MockaddressRepository)
	}{
		{
			name:      "Success",
			addressID: 2,
			newAddress: model.Addresses{
				Id:     2,
				City:   "Москва",
				Street: "Тверская",
				House:  "12",
				Flat:   func() *string { s := "34"; return &s }(),
			},
			expectedError: nil,
			mockSetup: func(mockRepo *mocks.MockaddressRepository) {
				mockRepo.EXPECT().UpdateUsersAddress(gomock.Any(), uint32(2), model.Addresses{
					Id:     2,
					City:   "Москва",
					Street: "Тверская",
					House:  "12",
					Flat:   func() *string { s := "34"; return &s }(),
				}).Return(nil)
			},
		},
		{
			name:      "RepositoryError",
			addressID: 3,
			newAddress: model.Addresses{
				Id:     3,
				City:   "Санкт-Петербург",
				Street: "Невский проспект",
				House:  "10",
				Flat:   func() *string { s := ""; return &s }(),
			},
			expectedError: errors.New("ошибка обновления адреса"),
			mockSetup: func(mockRepo *mocks.MockaddressRepository) {
				mockRepo.EXPECT().UpdateUsersAddress(gomock.Any(), uint32(3), model.Addresses{
					Id:     3,
					City:   "Санкт-Петербург",
					Street: "Невский проспект",
					House:  "10",
					Flat:   func() *string { s := ""; return &s }(),
				}).Return(errors.New("ошибка обновления адреса"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mocks.NewMockaddressRepository(ctrl)
			logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
			service := address.NewAddressService(mockRepo, logger)

			tt.mockSetup(mockRepo)

			err := service.UpdateUsersAddress(context.Background(), tt.addressID, tt.newAddress)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
