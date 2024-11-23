package csat

import (
	grpc_gen "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/csat/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"log/slog"
)

type CsatDelivery struct {
	csatGrpcClient grpc_gen.CsatServiceClient
	errResolver    errs.GetErrorCode
	log            *slog.Logger
}

func NewCsatDelivery(csatGrpcClient grpc_gen.CsatServiceClient,
	errResolver errs.GetErrorCode,
	log *slog.Logger) *CsatDelivery {
	return &CsatDelivery{
		csatGrpcClient: csatGrpcClient,
		errResolver:    errResolver,
		log:            log,
	}
}
