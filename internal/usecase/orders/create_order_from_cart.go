package orders

import (
	"context"
	"errors"
	"log/slog"
	"strconv"
	"time"

	order "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/google/uuid"
)

func (m *OrdersManager) CreateOrderFromCart(ctx context.Context, address string, userID uint32, promoName string) (*order.Order, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return nil, err
	}

	m.logger.Info("[OrdersManager.CreateOrderFromCart] Started executing", slog.Any("request-id", requestID))

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

	var totalPrice uint32
	productOrders := make([]order.ProductOrder, 0, len(cartItems))

	for _, item := range cartItems {
		totalPrice += item.Cost * item.Count
		productOrders = append(productOrders, item)
	}

	if promoName != "" {
		promoCode, err := m.promoCodesManager.GetPromoCode(ctx, userID, promoName)
		if err != nil {
			m.logger.Error("[OrdersManager.CreateOrderFromCart] Error getting promo code ", slog.Uint64("user_id", uint64(userID)))

			return nil, err
		}

		discountAmount := (totalPrice * promoCode.Bonus) / 100
		totalPrice -= discountAmount
		err = m.promoCodesManager.DeletePromoCode(ctx, userID, promoCode.ID)
		if err != nil {
			m.logger.Error("[OrdersManager.CreateOrderFromCart] Error deleting promo",
				slog.String("error", err.Error()))

			return nil, err
		}
	}

	orderData := &order.OrderFromCart{
		OrderID:      orderID,
		UserID:       userID,
		Address:      address,
		TotalPrice:   totalPrice,
		DeliveryDate: deliveryDate,
		Products:     cartItems,
	}

	orderFromCart, err := m.repo.CreateOrderFromCart(ctx, orderData)
	if err != nil {
		m.logger.Error("failed to create orderFromCart in repo", slog.String("error", err.Error()), slog.Uint64("user_id", uint64(userID)))
		return nil, err
	}

	m.logger.Info("CreateOrderFromCart completed successfully", slog.String("order_id", orderID.String()), slog.Uint64("user_id", uint64(userID)))
	return orderFromCart, nil
}
