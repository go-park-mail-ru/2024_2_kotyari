package csat

import (
	"context"
	"google.golang.org/grpc"

	proto "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/csat/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"log/slog"
)

type csatManager interface {
	CreateCsat(ctx context.Context, csat model.CSAT) (model.CSAT, error)
}

type csatRepo interface {
	GetStatistics(ctx context.Context, typeOf string) (model.CSATStatistics, float64, error)
	GetCsat(ctx context.Context, typeOf string) (string, error)
}

type CsatsGrpc struct {
	proto.UnimplementedCsatServiceServer
	csatManager csatManager
	csatRepo    csatRepo
	log         *slog.Logger
}

func (cg *CsatsGrpc) Register(grpcServer *grpc.Server) {
	proto.RegisterCsatServiceServer(grpcServer, cg)
}

func NewCsatsGrpc(csatManager csatManager, csatRepo csatRepo, log *slog.Logger) *CsatsGrpc {
	return &CsatsGrpc{
		csatManager: csatManager,
		csatRepo:    csatRepo,
		log:         log,
	}
}