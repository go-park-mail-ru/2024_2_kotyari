package morders

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"log/slog"
	"time"

	order "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (m *OrdersManager) GetOrderById(ctx context.Context, id uuid.UUID, deliveryDate time.Time, userID uint32) (*order.Order, error) {
	order, err := m.repo.GetOrderById(ctx, id, userID, deliveryDate)
	if err != nil {
		return nil, err
	}

	if order == nil {
		m.logger.Warn("[OrdersManager.GetOrderByID] order not found", slog.String("order_id", id.String()))
		return nil, pgx.ErrNoRows
	}

	m.logger.Info("[OrdersManager.GetOrderByID] GetOrderByID completed successfully", slog.String("order_id", id.String()))
	return order, nil
}
