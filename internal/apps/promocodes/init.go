package promocodes

import (
	"context"
	"log/slog"

	promocodes "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/promocodes/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/logger"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/postgres"
	promocodesGRPCLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/grpc_api/promocodes"
	promocodesConsumerLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/kafka_api/promocodes"
	promocodesRepositoryLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/promocodes"
	"google.golang.org/grpc"
)

type promoCodesGRPCDelivery interface {
	Register(server *grpc.Server)
	GetUserPromoCodes(ctx context.Context, request *promocodes.GetUserPromoCodesRequest) (*promocodes.GetUserPromoCodesResponse, error)
}

type promoCodesReader interface {
	Read() error
}

type PromoCodesApp struct {
	reader       promoCodesReader
	grpcDelivery promoCodesGRPCDelivery
	server       *grpc.Server
	log          *slog.Logger
	grpcConf     configs.ServiceViperConfig
}

func NewPromoCodesApp(kafkaConf map[string]any, serviceConf map[string]any) (*PromoCodesApp, error) {
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

	serviceConfig, err := configs.ParseServiceViperConfig(serviceConf)
	if err != nil {
		slog.Error("[NewPromoCodesApp] Failed to service cfg",
			slog.String("error", err.Error()))

		return nil, err
	}

	log := logger.InitLogger()
	grpcServer := grpc.NewServer()

	promoCodesRepo := promocodesRepositoryLib.NewPromoCodesStore(pool, log)
	promoCodesConsumer := promocodesConsumerLib.NewPromoCodesConsumer(kafkaConfig, promoCodesRepo, log)
	promoCodesGRPC := promocodesGRPCLib.NewPromoCodesGRPC(promoCodesRepo, log)
	promoCodesGRPC.Register(grpcServer)

	return &PromoCodesApp{
		server:       grpcServer,
		grpcConf:     serviceConfig,
		reader:       promoCodesConsumer,
		grpcDelivery: promoCodesGRPC,
		log:          log,
	}, nil
}
