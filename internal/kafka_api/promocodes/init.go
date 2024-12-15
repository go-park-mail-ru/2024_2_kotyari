package promocodes

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/segmentio/kafka-go"
)

type promoCodesRepository interface {
	AddPromoCode(ctx context.Context, userID uint32, promoID uint32) error
	GetUserPromoCodes(ctx context.Context, userID uint32) ([]model.PromoCode, error)
}

type PromoCodesConsumer struct {
	repository promoCodesRepository
	reader     *kafka.Reader
	log        *slog.Logger
}

func NewPromoCodesConsumer(kafkaConfig configs.KafkaConfig, repository promoCodesRepository, logger *slog.Logger) *PromoCodesConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{fmt.Sprintf("%s:%s", kafkaConfig.Domain, kafkaConfig.Port)},
		Topic:   utils.PromoTopic,
	})

	return &PromoCodesConsumer{
		repository: repository,
		reader:     reader,
		log:        logger,
	}
}
