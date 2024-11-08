package orders

import (
	"context"
	order "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func (m *OrdersManager) GetOrderById(ctx context.Context, id uuid.UUID, userID uint32) (*order.Order, error) {
	orderById, err := m.repo.GetOrderById(ctx, id, userID)
	if err != nil {
		return nil, err
	}

	if orderById == nil {
		m.logger.Warn("[OrdersManager.GetOrderByID] orderById not found", slog.String("order_id", id.String()))
		return nil, pgx.ErrNoRows
	}

	m.logger.Info("[OrdersManager.GetOrderByID] GetOrderByID completed successfully", slog.String("order_id", id.String()))
	return orderById, nil
}
