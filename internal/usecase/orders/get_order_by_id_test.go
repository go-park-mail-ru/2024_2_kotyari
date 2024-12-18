package orders

/*import (
	"context"
	"log/slog"

	order "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (m *OrdersManager) GetOrderById(ctx context.Context, id uuid.UUID, userID uint32) (*order.Order, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return nil, err
	}

	m.logger.Info("[OrdersManager.GetOrderById] Started executing", slog.Any("request-id", requestID))

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
*/
