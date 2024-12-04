package user

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/api/protos/user/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

type sessionCreator interface {
	Create(ctx context.Context, userID uint32) (string, error)
}

type UsersDelivery struct {
	userClientGrpc grpc_gen.UserServiceClient
	inputValidator *utils.InputValidator
	sessionService sessionCreator
	errResolver    errs.GetErrorCode
	log            *slog.Logger
}

func NewUsersDelivery(userManager grpc_gen.UserServiceClient, inputValidator *utils.InputValidator, sessionService sessionCreator, errResolver errs.GetErrorCode, log *slog.Logger) *UsersDelivery {
	return &UsersDelivery{
		userClientGrpc: userManager,
		inputValidator: inputValidator,
		sessionService: sessionService,
		errResolver:    errResolver,
		log:            log,
	}
}

type MessageProducer struct {
	writer *kafka.Writer
	log    *slog.Logger
}

func NewMessageProducer(logger *slog.Logger) *MessageProducer {
	///TODO: Remove magic constants
	topic := "promo-topic"
	w := &kafka.Writer{
		Addr:     kafka.TCP("kafka:9092"),
		Topic:    topic,
		Balancer: &kafka.Hash{},
	}
	return &MessageProducer{
		writer: w,
		log:    logger,
	}
}
