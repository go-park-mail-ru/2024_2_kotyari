package promocodes

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/segmentio/kafka-go"
)

func (p *PromoCodesConsumer) Read() error {
	ctx, cancel := context.WithCancel(context.Background())

	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-sigChan
		cancel()
		err := p.reader.Close()
		if err != nil {
			p.log.Error("[PromoCodesConsumer.Read] Failed to close reader",
				slog.String("error", err.Error()))
		}
	}()

	for {
		///TODO: Разобраться с EOF
		message, _ := p.reader.ReadMessage(ctx)
		//if err != nil {
		//	p.log.Error("[PromoCodesConsumer.Read] Error reading message", slog.String("error", err.Error()))
		//
		//	err = p.reader.Close()
		//	if err != nil {
		//		p.log.Error("[PromoCodesConsumer.Read] Error closing reader", slog.String("error", err.Error()))
		//
		//		return err
		//	}
		//
		//	return err
		//}

		err := p.processMessage(message)
		if err != nil {
			p.log.Error("[PromoCodesConsumer.Read] Error processing message", slog.String("error", err.Error()))

			return err
		}
	}

}

func (p *PromoCodesConsumer) processMessage(promoMessage kafka.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var message utils.PromoMessage

	err := json.Unmarshal(promoMessage.Value, &message)
	if err != nil {
		p.log.Error("[PromoCodesConsumer.processMessage] Failed to unmarshal kafka msg",
			slog.String("error", err.Error()))

		return err
	}

	p.log.Info("[PromoCodesConsumer.processMessage] Started processing message",
		slog.Any("request-id", message.RequestID))

	ctx = context.WithValue(ctx, utils.RequestIDName, message.RequestID)

	switch utils.MessageType(promoMessage.Key) {
	case utils.AddPromo:
		err = p.repository.AddPromoCode(ctx, message.UserID, message.PromoID)
		if err != nil {
			p.log.Error("[PromoCodesConsumer.processMessage] Error adding promo",
				slog.String("error", err.Error()))

			return err
		}

		return nil

	case utils.DeletePromo:
		err = p.repository.DeletePromoCode(ctx, message.UserID, message.PromoID)
		if err != nil {
			p.log.Error("[PromoCodesConsumer.processMessage] Error adding promo",
				slog.String("error", err.Error()))

			return err
		}

		return nil

	default:
		p.log.Error("[PromoCodesConsumer.processMessage] Invalid message type")

		return nil
	}
}
