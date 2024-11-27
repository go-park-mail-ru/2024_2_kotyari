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

func TestOrdersManager_GetOrders(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	var userID uint32
	faker.FakeData(&userID)

	orderID := uuid.New()
	deliveryDate := time.Now().AddDate(0, 0, 5)
	orderDate := time.Now().AddDate(0, -1, 0)

	var productName string
	var imageUrl string
	faker.FakeData(&productName)
	faker.FakeData(&imageUrl)

	totalPriceSlice, err := faker.RandomInt(100, 2000)
	if err != nil {
		t.Fatalf("failed to generate random total price: %v", err)
	}
	totalPrice := uint32(totalPriceSlice[0])

	productIDSlice, err := faker.RandomInt(1, 100)
	if err != nil {
		t.Fatalf("failed to generate random product ID: %v", err)
	}
	productID := uint32(productIDSlice[0])

	expectedOrders := []model.Order{
		{
			ID:           orderID,
			OrderDate:    orderDate,
			DeliveryDate: deliveryDate,
			TotalPrice:   totalPrice,
			Status:       faker.Word(),
			Products: []model.ProductOrder{
				{ProductID: productID, Name: productName, ImageUrl: imageUrl},
			},
		},
	}

	tests := []struct {
		name           string
		setupMockFunc  func(repoMock *mocks.MockOrdersRepo)
		expectedOrders []model.Order
		expectedErr    error
	}{
		{
			name: "success",
			setupMockFunc: func(repoMock *mocks.MockOrdersRepo) {
				repoMock.EXPECT().
					GetOrders(ctx, userID).
					Return(expectedOrders, nil).
					Times(1)
			},
			expectedOrders: expectedOrders,
			expectedErr:    nil,
		},
		{
			name: "repo error",
			setupMockFunc: func(repoMock *mocks.MockOrdersRepo) {
				repoMock.EXPECT().
					GetOrders(ctx, userID).
					Return(nil, errors.New("database error")).
					Times(1)
			},
			expectedOrders: nil,
			expectedErr:    errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			repoMock := mocks.NewMockOrdersRepo(ctrl)
			loggerMock := slog.New(slog.NewTextHandler(os.Stdout, nil))
			cartMock := mocks.NewMockcartGetter(ctrl)
			ordersManager := NewOrdersManager(repoMock, loggerMock, cartMock)

			tt.setupMockFunc(repoMock)

			orders, err := ordersManager.GetOrders(ctx, userID)
			assert.Equal(t, tt.expectedOrders, orders)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
