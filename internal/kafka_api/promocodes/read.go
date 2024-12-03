package promocodes

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/segmentio/kafka-go"
)

func (p *PromoCodesConsumer) Read() error {
	for {
		message, err := p.reader.ReadMessage(context.Background())
		if err != nil {
			p.log.Error("[PromoCodesConsumer.Read] Error reading message", slog.String("error", err.Error()))

			err = p.reader.Close()
			if err != nil {
				p.log.Error("[PromoCodesConsumer.Read] Error closing reader", slog.String("error", err.Error()))

				return err
			}

			return err
		}
	}
}

func (p *PromoCodesConsumer) switchMessageType(promoMessage kafka.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var message utils.PromoMessage

	err := json.Unmarshal(promoMessage.Value, &message)
	if err != nil {
		p.log.Error("[PromoCodesConsumer.switchMessageType] Failed to unmarshal kafka msg",
			slog.String("error", err.Error()))

		return err
	}

	switch utils.MessageType(promoMessage.Key) {
	case utils.AddPromo:
		err = p.repository.AddPromoCode(ctx, message.UserID, message.PromoID)
		if err != nil {
			p.log.Error("[PromoCodesConsumer.switchMessageType] Error adding promo",
				slog.String("error", err.Error()))

			return err
		}

		return nil

	case utils.DeletePromo:
		err = p.repository.DeletePromoCode(ctx, message.UserID, message.PromoID)
		if err != nil {
			p.log.Error("[PromoCodesConsumer.switchMessageType] Error adding promo",
				slog.String("error", err.Error()))

			return err
		}

		return nil

	default:
		p.log.Error("[PromoCodesConsumer.switchMessageType] Invalid message type")

		return nil
	}
}
