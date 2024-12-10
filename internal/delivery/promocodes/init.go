package promocodes

import (
	"fmt"
	"log/slog"

	promocodes "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/promocodes/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PromoCodesGRPC struct {
	client       promocodes.PromoCodesClient
	errorHandler errs.GetErrorCode
	log          *slog.Logger
}

func NewPromoCodesGRPC(config map[string]any, errorHandler errs.GetErrorCode, log *slog.Logger) (*PromoCodesGRPC, error) {
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
		client:       client,
		errorHandler: errorHandler,
		log:          log,
	}, nil
}
