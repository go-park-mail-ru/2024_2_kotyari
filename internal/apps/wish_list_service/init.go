package wish_list_service

import (
	"context"
	"errors"
	"fmt"
	wishlistgrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/wish_list/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/logger"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/mongo_db"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/postgres"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/grpc_api/wish_list"
	rep "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/wish_list"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/wish_list_link"
	use "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/wish_list"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"log/slog"
)

type wishlistServer interface {
	Register(grpcServer *grpc.Server)
	AddProductToWishlists(ctx context.Context, in *wishlistgrpc.AddProductRequest) (*empty.Empty, error)
	CopyWishlist(ctx context.Context, in *wishlistgrpc.CopyWishlistRequest) (*wishlistgrpc.CopyWishlistResponse, error)
	CreateWishlist(ctx context.Context, in *wishlistgrpc.CreateWishlistRequest) (*empty.Empty, error)
	DeleteWishlist(ctx context.Context, in *wishlistgrpc.DeleteWishlistRequest) (*empty.Empty, error)
	GetAllUserWishlists(ctx context.Context, in *wishlistgrpc.GetAllWishlistsRequest) (*wishlistgrpc.GetAllWishlistsResponse, error)
	GetWishlistByLink(ctx context.Context, in *wishlistgrpc.GetWishlistByLinkRequest) (*wishlistgrpc.GetWishlistByLinkResponse, error)
	RemoveFromWishlists(ctx context.Context, in *wishlistgrpc.RemoveFromWishlistsRequest) (*empty.Empty, error)
	RenameWishlist(ctx context.Context, in *wishlistgrpc.RenameWishlistRequest) (*empty.Empty, error)
}

type WishlistApp struct {
	log        *slog.Logger
	gRPCServer *grpc.Server

	server wishlistServer
	config configs.ServiceViperConfig
}

func NewWishlistApp(conf map[string]any, server *grpc.Server) (*WishlistApp, error) {
	slogLog := logger.InitLogger()

	cfg, err := configs.ParseServiceViperConfig(conf)
	if err != nil {
		slogLog.Error("[NewWishlistApp] Failed to parse cfg")

		return nil, err
	}

	if cfg.Address == "" || cfg.Port == "" {
		return nil, errors.New("[ ERROR ] пустая конфигурация сервиса NewWishlistApp")
	}

	dbPool, err := postgres.LoadPgxPool()
	if err != nil {
		return nil, fmt.Errorf("[ ERROR ] не инициализируется бд %v", err)
	}

	mongoDb, err := mongo_db.Connect()
	if err != nil {
		return nil, fmt.Errorf("[ ERROR ] не инициализируется монго %v", err)
	}

	repo := rep.NewWishListRepo(mongoDb, "wishlists", slogLog)

	links := wish_list_link.NewWishListLinkRepo(dbPool)

	usecase := use.NewWushlistUsecase(repo, links, slogLog)

	api := wish_list.NewWishlistsGrpc(usecase, slogLog)

	return &WishlistApp{
		config:     cfg,
		log:        slogLog,
		gRPCServer: server,
		server:     api,
	}, nil
}
