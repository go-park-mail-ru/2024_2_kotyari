package user

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/segmentio/kafka-go"
)

func (m *MessageProducer) AddPromoCode(ctx context.Context, userID uint32, promoID uint32) error {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		m.log.Error("[MessageProducer.AddPromoCode] Failed to get request id",
			slog.String("error", err.Error()))

		return err
	}

	marshalled, err := json.Marshal(utils.PromoMessage{
		UserID:    userID,
		PromoID:   promoID,
		RequestID: requestID,
	})
	if err != nil {
		m.log.Error("[MessageProducer.AddPromoCode] Failed to marshal message struct",
			slog.String("error", err.Error()))

		return err
	}

	///TODO: Разобраться с контекстом
	//ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	//defer cancel()

	fmt.Println("АЛО АЛО", userID, promoID)

	err = m.writer.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte(utils.AddPromo),
			Value: marshalled,
		},
	)
	if err != nil {
		m.log.Error("[MessageProducer.AddPromoCode] Error sending message",
			slog.String("error", err.Error()))

		return err
	}

	return nil
}
