package orders

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (m *OrdersManager) GetNearestDeliveryDate(ctx context.Context, userID uint32) (time.Time, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return time.Time{}, err
	}

	m.logger.Info("[OrdersManager.GetNearestDeliveryDate] Started executing", slog.Any("request-id", requestID))

	deliveryDate, err := m.repo.GetNearestDeliveryDate(ctx, userID)
	if err != nil {
		m.logger.Error("[OrdersManager.GetNearestDeliveryDate] failed to get nearest delivery date", slog.String("error", err.Error()))
		return time.Time{}, err
	}

	if deliveryDate.IsZero() {
		m.logger.Error("[OrdersManager.GetNearestDeliveryDate] no future deliveries found")
		return time.Time{}, errors.New("no future deliveries found")
	}

	m.logger.Info("[OrdersManager.GetNearestDeliveryDate] Fetched nearest delivery date successfully")
	return deliveryDate, nil
}
