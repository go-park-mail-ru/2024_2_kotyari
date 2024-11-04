package morders

import (
	"context"
	"errors"
	"log/slog"
	"time"
)

func (m *OrdersManager) GetNearestDeliveryDate(ctx context.Context, userID uint32) (time.Time, error) {
	m.logger.Info("GetNearestDeliveryDate called", slog.Uint64("user_id", uint64(userID)))

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
