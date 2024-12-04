package user

import (
	"context"
	"errors"
	"fmt"
	proto "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/user/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/postgres"
	userProducerLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/user"
	user2 "github.com/go-park-mail-ru/2024_2_kotyari/internal/grpc_api/user"
	userRepoLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/user"
	userServiceLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/user"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"google.golang.org/grpc"
	"log/slog"
)

type usersDelivery interface {
	Register(grpcServer *grpc.Server)
	LoginUser(ctx context.Context, in *proto.UsersLoginRequest) (*proto.UsersDefaultResponse, error)
	GetUserById(ctx context.Context, in *proto.GetUserByIdRequest) (*proto.UsersDefaultResponse, error)
	CreateUser(ctx context.Context, in *proto.UsersSignUpRequest) (*proto.UsersDefaultResponse, error)
}

type UsersApp struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	delivery   usersDelivery
	config     configs.ServiceViperConfig
}

func NewUsersApp(log *slog.Logger, grpcServer *grpc.Server, conf map[string]any) (*UsersApp, error) {
	c := configs.ParseServiceViperConfig(conf)
	if c.Address == "" || c.Port == "" {
		return nil, errors.New("config is empty")
	}

	dbPool, err := postgres.LoadPgxPool()
	if err != nil {
		return nil, fmt.Errorf("не инициализируется бд %v", err)
	}

	// todo добавить
	inputValidator := utils.NewInputValidator()

	userRepo := userRepoLib.NewUsersStore(dbPool, log)
	userProducer := userProducerLib.NewMessageProducer(log)
	userService := userServiceLib.NewUserService(userRepo, userProducer, inputValidator, log)

	delivery := user2.NewUsersGrpc(userService, userRepo, log)

	return &UsersApp{
		log:        log,
		gRPCServer: grpcServer,
		delivery:   delivery,
		config:     c,
	}, nil
}
