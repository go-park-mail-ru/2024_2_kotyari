package csat_service

import (
	"context"
	"errors"
	"fmt"
	proto "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/csat/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/postgres"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/grpc_api/csat"
	csatRepoLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/csat"
	csatServiceLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/csat"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type csatDelivery interface {
	Register(grpcServer *grpc.Server)
	CreateCsat(ctx context.Context, in *proto.CreateCsatRequest) (*empty.Empty, error)
	GetStatistics(ctx context.Context, in *proto.GetStatisticsRequest) (*proto.GetStatisticsResponse, error)
}

type config struct {
	Address string
	Port    string
}

type CSATApp struct {
	log        *slog.Logger
	gRPCServer *grpc.Server

	delivery csatDelivery
	config   config
}

func NewCsatApp(
	grpcServer *grpc.Server,
	conf map[string]any,
	slogLog *slog.Logger,
) (*CSATApp, error) {
	c := config{
		Address: conf[configs.KeyAddress].(string),
		Port:    conf[configs.KeyPort].(string),
	}

	if c.Address == "" || c.Port == "" {
		return nil, errors.New("config is empty")
	}

	dbPool, err := postgres.LoadPgxPool(postgres.CSATDBCFG)
	if err != nil {
		return nil, fmt.Errorf("не инициализируется бд %v", err)
	}

	csatRepo := csatRepoLib.NewSurveyStore(dbPool, slogLog)
	csatService := csatServiceLib.NewCSATService(csatRepo, slogLog)

	delivery := csat.NewCsatsGrpc(csatService, csatRepo, slogLog)

	return &CSATApp{
		log:        slogLog,
		gRPCServer: grpcServer,
		delivery:   delivery,
		config:     c,
	}, nil
}

func (app *CSATApp) Run() error {
	lis, err := net.Listen("tcp",
		app.config.Address+app.config.Port,
	)
	if err != nil {
		app.log.Error("[  ProfilesApp.Run ] ",
			slog.String("error", err.Error()),
		)

		return err
	}

	app.delivery.Register(app.gRPCServer)

	app.log.Info("[ ProfilesApp.Run ]",
		slog.String("address", app.config.Address+app.config.Port),
	)

	if err = app.gRPCServer.Serve(lis); err != nil {
		app.log.Error("[ ProfilesApp.Run ]",
			slog.String("error", err.Error()),
		)

		return err
	}

	return nil
}
