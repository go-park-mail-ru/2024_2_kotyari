package morders

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (m *OrdersManager) GetNearestDeliveryDate(ctx context.Context) (time.Time, error) {
	userID := utils.GetContextSessionUserID(ctx)
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
