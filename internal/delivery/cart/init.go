package cart

import (
	"context"
	"fmt"
	promocodes "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/promocodes/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	"time"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type cartManager interface {
	ChangeCartProductCount(ctx context.Context, productID uint32, count int32, userID uint32) (uint32, error)
	AddProduct(ctx context.Context, productID uint32, userID uint32) error
	RemoveProduct(ctx context.Context, productID uint32, userID uint32) error
	ChangeCartProductSelectedState(ctx context.Context, productID uint32, userID uint32, isSelected bool) error
	GetSelectedFromCart(ctx context.Context, userID uint32, promoName string) (model.CartForOrder, error)
	RemoveSelected(ctx context.Context, userID uint32) error
}

type cartManip interface {
	GetCart(ctx context.Context, userID uint32, deliveryDate time.Time) (model.Cart, error)
	ChangeAllCartProductsState(ctx context.Context, userID uint32, isSelected bool) error
	UpdatePaymentMethod(ctx context.Context, userID uint32, method string) error
}

type CartHandler struct {
	cartManager cartManager
	cartManip   cartManip
	errResolver errs.GetErrorCode
	log         *slog.Logger
}

func NewCartHandler(manager cartManager, manip cartManip, errorHandler errs.GetErrorCode, logger *slog.Logger) *CartHandler {
	return &CartHandler{
		cartManager: manager,
		cartManip:   manip,
		errResolver: errorHandler,
		log:         logger,
	}
}

type PromoCodesGetterGRPC struct {
	client promocodes.PromoCodesClient
	log    *slog.Logger
}

func NewPromoCodesGetterGRPC(config map[string]any, log *slog.Logger) (*PromoCodesGetterGRPC, error) {
	cfg, err := configs.ParseServiceViperConfig(config)
	if err != nil {
		slog.Error("[NewPromoCodesGetterGRPC] Failed to parse cfg",
			slog.String("error", err.Error()))

		return nil, err
	}

	promoCodesConnection, err := grpc.NewClient(fmt.Sprintf("%s:%s", cfg.Domain, cfg.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("[NewPromoCodesGetterGRPC] Failed to establish gRPC connection",
			slog.String("error", err.Error()))

		return nil, err
	}

	client := promocodes.NewPromoCodesClient(promoCodesConnection)

	return &PromoCodesGetterGRPC{
		client: client,
		log:    log,
	}, nil
}
