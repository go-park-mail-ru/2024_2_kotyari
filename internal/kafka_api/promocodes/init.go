package promocodes

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/segmentio/kafka-go"
)

type promoCodesRepository interface {
	AddPromoCode(ctx context.Context, userID uint32, promoID uint32) error
	GetUserPromoCodes(ctx context.Context, userID uint32) ([]model.PromoCode, error)
	DeletePromoCode(ctx context.Context, userID uint32, promoID uint32) error
}

type PromoCodesConsumer struct {
	repository promoCodesRepository
	reader     *kafka.Reader
	log        *slog.Logger
}

func NewPromoCodesConsumer(repository promoCodesRepository, logger *slog.Logger) *PromoCodesConsumer {
	//TODO: REFACTOR INIT
	topic := "promo_codes"
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"kafka:9092"},
		Topic:    topic,
		MaxBytes: 10e6,
	})

	return &PromoCodesConsumer{
		repository: repository,
		reader:     reader,
		log:        logger,
	}
}
