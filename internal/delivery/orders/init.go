package orders

import (
	"context"
	"fmt"
	promocodes "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/promocodes/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	"time"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	order "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type ordersManager interface {
	GetOrders(ctx context.Context, userID uint32) ([]order.Order, error)
	GetOrderById(ctx context.Context, id uuid.UUID, userID uint32) (*order.Order, error)
	CreateOrderFromCart(ctx context.Context, address string, userID uint32, promoName string) (*order.Order, error)
	GetNearestDeliveryDate(ctx context.Context, userID uint32) (time.Time, error)
}

type OrdersHandler struct {
	ordersManager ordersManager
	logger        *slog.Logger
	errResolver   errs.GetErrorCode
}

func NewOrdersHandler(manager ordersManager, logger *slog.Logger, errResolver errs.GetErrorCode) *OrdersHandler {
	return &OrdersHandler{
		ordersManager: manager,
		logger:        logger,
		errResolver:   errResolver,
	}
}

type PromoCodesGRPC struct {
	client promocodes.PromoCodesClient
	log    *slog.Logger
}

func NewPromoCodesGRPC(config map[string]any, log *slog.Logger) (*PromoCodesGRPC, error) {
	cfg, err := configs.ParseServiceViperConfig(config)
	if err != nil {
		slog.Error("[NewPromoCodesGRPC] Failed to parse cfg",
			slog.String("error", err.Error()))

		return nil, err
	}

	promoCodesConnection, err := grpc.NewClient(fmt.Sprintf("%s:%s", cfg.Domain, cfg.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("[NewPromoCodesGRPC] Failed to establish gRPC connection",
			slog.String("error", err.Error()))

		return nil, err
	}

	client := promocodes.NewPromoCodesClient(promoCodesConnection)

	return &PromoCodesGRPC{
		client: client,
		log:    log,
	}, nil
}
