package csat

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type csatRepository interface {
	GetCSAT(ctx context.Context, csat model.CSAT) (model.CSAT, error)
	CreateCSAT(ctx context.Context, csat model.CSAT) error
}

type CSATService struct {
	repository csatRepository
	log        *slog.Logger
}

func NewCSATService(repository csatRepository, logger *slog.Logger) *CSATService {
	return &CSATService{
		repository: repository,
		log:        logger,
	}
}
