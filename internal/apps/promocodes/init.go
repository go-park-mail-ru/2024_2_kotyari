package promocodes

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs"
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

func NewPromoCodesApp(kafkaConf map[string]any, _ map[string]any) (*PromoCodesApp, error) {
	pool, err := postgres.LoadPgxPool()
	if err != nil {
		slog.Error("[NewPromoCodesApp] Failed to load pgxpool",
			slog.String("error", err.Error()))

		return nil, err
	}

	kafkaConfig, err := configs.ParseKafkaViperConfig(kafkaConf)
	if err != nil {
		slog.Error("[NewPromoCodesApp] Failed to kafka cfg",
			slog.String("error", err.Error()))

		return nil, err
	}

	log := logger.InitLogger()
	promoCodesRepo := promocodesRepositoryLib.NewPromoCodesStore(pool, log)
	promoCodesConsumer := promocodesDeliveryLib.NewPromoCodesConsumer(kafkaConfig, promoCodesRepo, log)

	return &PromoCodesApp{
		reader: promoCodesConsumer,
		log:    log,
	}, nil
}
