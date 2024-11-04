package morders

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"log/slog"
	"strconv"
	"time"

	order "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (m *OrdersManager) CreateOrderFromCart(ctx context.Context, address string, userID uint32) (*order.Order, error) {
	orderID := uuid.New()
	orderDate := time.Now()
	deliveryDate := orderDate.Add(72 * time.Hour)

	cartItems, err := m.cart.GetSelectedCartItems(ctx, userID)
	if err != nil {
		m.logger.Error("[OrdersManager.CreateOrderFromCart] failed to fetch selected cart items", slog.String("error", err.Error()), slog.Uint64("user_id", uint64(userID)))
		return nil, err
	}

	if len(cartItems) == 0 {
		m.logger.Error("[OrdersManager.CreateOrderFromCart] cart is empty for user: ", slog.Uint64("user_id", uint64(userID)))
		return nil, errors.New("cart is empty for user: " + strconv.Itoa(int(userID)))
	}

	var totalPrice uint16
	productOrders := make([]order.ProductOrder, 0, len(cartItems))

	for _, item := range cartItems {
		totalPrice += item.Cost * item.Count
		productOrders = append(productOrders, item)
	}

	orderData := &order.OrderFromCart{
		OrderID:      orderID,
		UserID:       userID,
		Address:      address,
		TotalPrice:   totalPrice,
		DeliveryDate: deliveryDate,
		Products:     cartItems,
	}

	order, err := m.repo.CreateOrderFromCart(ctx, orderData)
	if err != nil {
		m.logger.Error("failed to create order in repo", slog.String("error", err.Error()), slog.Uint64("user_id", uint64(userID)))
		return nil, err
	}

	m.logger.Info("CreateOrderFromCart completed successfully", slog.String("order_id", orderID.String()), slog.Uint64("user_id", uint64(userID)))
	return order, nil
}
