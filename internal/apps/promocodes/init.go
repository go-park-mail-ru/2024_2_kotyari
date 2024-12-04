package promocodes

import (
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/logger"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/postgres"
	promocodesDeliveryLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/kafka_api/promocodes"
	promocodesRepositoryLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/promocodes"
)

type promoCodesReader interface {
	Read() error
}

type PromoCodesApp struct {
	reader promoCodesReader
	log    *slog.Logger
}

func NewPromoCodesApp() (*PromoCodesApp, error) {
	pool, err := postgres.LoadPgxPool()
	if err != nil {
		slog.Error("[NewPromoCodesApp] Failed to load pgxpool",
			slog.String("error", err.Error()))
	}

	log := logger.InitLogger()
	promoCodesRepo := promocodesRepositoryLib.NewPromoCodesStore(pool, log)
	promoCodesConsumer := promocodesDeliveryLib.NewPromoCodesConsumer(promoCodesRepo, log)

	return &PromoCodesApp{
		reader: promoCodesConsumer,
		log:    log,
	}, nil
}
