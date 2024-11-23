package csat_service

import (
	"context"
	"errors"
	"fmt"
	proto "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/csat/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/postgres"
	profileRepoLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/csat"
	profileServiceLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/csat"
	"google.golang.org/grpc"
	"log/slog"
)

type csatDelivery interface {
	Register(grpcServer *grpc.Server)
	CreateCsat(ctx context.Context, in *proto.CreateCsatRequest) (*proto.CreateCsatResponse, error)
	GetStatistics(ctx context.Context, in *proto.GetStatisticsRequest) (*proto.GetStatisticsResponse, error)
}

type config struct {
	Address string
	Port    string
}

type ProfilesApp struct {
	log        *slog.Logger
	gRPCServer *grpc.Server

	delivery csatDelivery
	config   config
}

func NewCsatApp(
	grpcServer *grpc.Server,
	conf map[string]any,
	slogLog *slog.Logger,
) (*ProfilesApp, error) {
	c := config{
		Address: conf[configs.KeyAddress].(string),
		Port:    conf[configs.KeyPort].(string),
	}

	if c.Address == "" || c.Port == "" {
		return nil, errors.New("config is empty")
	}

	dbPool, err := postgres.LoadPgxPool()
	if err != nil {
		return nil, fmt.Errorf("не инициализируется бд %v", err)
	}

	csatRepo := profileRepoLib.NewCsatRepo(dbPool, slogLog)
	csatService := profileServiceLib.NewCsatService(csatRepo, slogLog)

	delivery := csat.NewCsatGrpc(slogLog)

	return &ProfilesApp{
		log:        slogLog,
		gRPCServer: grpcServer,
		delivery:   delivery,
		config:     c,
	}, nil
}
