package orders

import (
	"context"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/orders/mocks"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestOrdersManager_GetOrderById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mocks.NewMockOrdersRepo(ctrl)
	loggerMock := slog.New(slog.NewTextHandler(os.Stdout, nil))
	cartMock := mocks.NewMockcartGetter(ctrl)

	ordersManager := NewOrdersManager(repoMock, loggerMock, cartMock)
	ctx := context.Background()

	t.Run("successful retrieval", func(t *testing.T) {
		orderID := uuid.New()
		var userID uint32
		faker.FakeData(&userID)
		deliveryDate := time.Now().UTC().Truncate(time.Millisecond)

		var productId uint32
		faker.FakeData(&productId)
		var count uint32
		faker.FakeData(&count)

		var cost uint32
		faker.FakeData(&cost)
		var weight float32
		faker.FakeData(&weight)

		// Сгенерировать случайный заказ с помощью Faker
		generatedOrder := &model.Order{
			ID:           orderID,
			Recipient:    faker.Name(),
			Address:      faker.Word(),
			Status:       "delivered",
			DeliveryDate: deliveryDate,
			OrderDate:    time.Now(),
			Products: []model.ProductOrder{
				{
					ProductID: productId,
					Cost:      cost,
					Count:     count,
					ImageUrl:  faker.URL(),
					Weight:    weight,
					Name:      faker.Word(),
				},
			},
		}

		// Настроить возвращаемое значение мока
		repoMock.EXPECT().
			GetOrderById(ctx, orderID, userID, deliveryDate).
			Return(generatedOrder, nil).
			Times(1)

		order, err := ordersManager.GetOrderById(ctx, orderID, deliveryDate, userID)
		assert.NoError(t, err)
		assert.NotNil(t, order)
		assert.Equal(t, generatedOrder.ID, order.ID)
		assert.Equal(t, generatedOrder.Recipient, order.Recipient)
		assert.Equal(t, generatedOrder.Address, order.Address)
		assert.Equal(t, generatedOrder.Status, order.Status)
	})

	t.Run("order not found", func(t *testing.T) {
		orderID := uuid.New()
		var userID uint32
		faker.FakeData(&userID)
		deliveryDate := time.Now().UTC().Truncate(time.Millisecond)

		// Настроить мок так, чтобы он возвращал pgx.ErrNoRows
		repoMock.EXPECT().
			GetOrderById(ctx, orderID, userID, deliveryDate).
			Return(nil, pgx.ErrNoRows).
			Times(1)

		order, err := ordersManager.GetOrderById(ctx, orderID, deliveryDate, userID)
		assert.Error(t, err)
		assert.Nil(t, order)
		assert.Equal(t, pgx.ErrNoRows, err)
	})
}
