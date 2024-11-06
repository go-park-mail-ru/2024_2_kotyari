package morders

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/bxcodec/faker/v4"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/orders/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"log/slog"
)

func TestOrdersManager_GetNearestDeliveryDate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	var userID uint32
	faker.FakeData(&userID)

	repoMock := mocks.NewMockOrdersRepo(ctrl)
	loggerMock := slog.New(slog.NewTextHandler(os.Stdout, nil))

	ordersManager := NewOrdersManager(repoMock, loggerMock, nil)

	t.Run("success", func(t *testing.T) {
		deliveryDate := time.Now().Add(48 * time.Hour)

		repoMock.EXPECT().
			GetNearestDeliveryDate(ctx, userID).
			Return(deliveryDate, nil).
			Times(1)

		date, err := ordersManager.GetNearestDeliveryDate(ctx, userID)
		assert.NoError(t, err)
		assert.Equal(t, deliveryDate, date)
	})

	t.Run("no future deliveries error", func(t *testing.T) {
		repoMock.EXPECT().
			GetNearestDeliveryDate(ctx, userID).
			Return(time.Time{}, nil).
			Times(1)

		date, err := ordersManager.GetNearestDeliveryDate(ctx, userID)
		assert.Error(t, err)
		assert.Equal(t, "no future deliveries found", err.Error())
		assert.True(t, date.IsZero())
	})

	t.Run("repo error", func(t *testing.T) {
		repoMock.EXPECT().
			GetNearestDeliveryDate(ctx, userID).
			Return(time.Time{}, errors.New("database error")).
			Times(1)

		date, err := ordersManager.GetNearestDeliveryDate(ctx, userID)
		assert.Error(t, err)
		assert.True(t, date.IsZero())
		assert.Equal(t, "database error", err.Error())
	})
}
