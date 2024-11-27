package orders

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/orders/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestOrdersManager_CreateOrderFromCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	var userID uint32
	var address string
	faker.FakeData(&userID)
	faker.FakeData(&address)
	orderID := uuid.New()
	deliveryDate := time.Now().Add(72 * time.Hour)

	totalPriceSlice, err := faker.RandomInt(100, 2000)
	if err != nil {
		t.Fatalf("failed to generate random total price: %v", err)
	}
	totalPrice := uint32(totalPriceSlice[0])

	var optionID uint32
	faker.FakeData(&optionID)

	var ID1 uint32
	faker.FakeData(&ID1)
	var ID2 uint32
	faker.FakeData(&ID2)

	countSlice1, err := faker.RandomInt(100, 2000)
	if err != nil {
		t.Fatalf("failed to generate random count: %v", err)
	}
	count1 := uint32(countSlice1[0])

	countSlice2, err := faker.RandomInt(100, 2000)
	if err != nil {
		t.Fatalf("failed to generate random second count: %v", err)
	}
	count2 := uint32(countSlice2[0])

	costSlice1, err := faker.RandomInt(100, 2000)
	if err != nil {
		t.Fatalf("failed to generate random cost: %v", err)
	}
	cost1 := uint32(costSlice1[0])

	costSlice2, err := faker.RandomInt(100, 2000)
	if err != nil {
		t.Fatalf("failed to generate random second cost: %v", err)
	}
	cost2 := uint32(costSlice2[0])

	cartItems := []model.ProductOrder{
		{ID: ID1, Count: count1, Cost: cost1},
		{ID: ID2, OptionID: &optionID, Count: count2, Cost: cost2},
	}

	repoMock := mocks.NewMockOrdersRepo(ctrl)
	cartMock := mocks.NewMockcartGetter(ctrl)
	loggerMock := slog.New(slog.NewTextHandler(os.Stdout, nil))

	ordersManager := NewOrdersManager(repoMock, loggerMock, cartMock)

	t.Run("success", func(t *testing.T) {
		cartMock.EXPECT().
			GetSelectedCartItems(ctx, userID).
			Return(cartItems, nil).
			Times(1)

		repoMock.EXPECT().
			CreateOrderFromCart(gomock.Any(), gomock.Any()).
			DoAndReturn(func(_ context.Context, orderData *model.OrderFromCart) (*model.Order, error) {
				return &model.Order{
					ID:           orderID,
					Address:      address,
					Status:       "awaiting_payment",
					OrderDate:    time.Now(),
					DeliveryDate: deliveryDate,
					Products:     cartItems,
					TotalPrice:   totalPrice,
				}, nil
			}).
			Times(1)

		order, err := ordersManager.CreateOrderFromCart(ctx, address, userID)
		assert.NoError(t, err)
		assert.NotNil(t, order)
		assert.Equal(t, address, order.Address)
		assert.Equal(t, totalPrice, order.TotalPrice)
		assert.Equal(t, "awaiting_payment", order.Status)
	})

	t.Run("empty cart error", func(t *testing.T) {
		cartMock.EXPECT().
			GetSelectedCartItems(ctx, userID).
			Return([]model.ProductOrder{}, nil).
			Times(1)

		order, err := ordersManager.CreateOrderFromCart(ctx, address, userID)
		assert.Error(t, err)
		assert.Nil(t, order)
		assert.Contains(t, err.Error(), "cart is empty")
	})

	t.Run("repo error", func(t *testing.T) {
		cartMock.EXPECT().
			GetSelectedCartItems(ctx, userID).
			Return(cartItems, nil).
			Times(1)

		repoMock.EXPECT().
			CreateOrderFromCart(gomock.Any(), gomock.Any()).
			Return(nil, errors.New("database error")).
			Times(1)

		order, err := ordersManager.CreateOrderFromCart(ctx, address, userID)
		assert.Error(t, err)
		assert.Nil(t, order)
		assert.Contains(t, err.Error(), "database error")
	})
}
