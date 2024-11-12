package search

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"log/slog"
)

type SearchRepository interface {
	GetSearchTitleSuggestions(ctx context.Context, queryParam string) (model.SearchTitleSuggestions, error)
}

type SearchHandler struct {
	searchRepository SearchRepository
	errResolver      errs.GetErrorCode
	log              *slog.Logger
}

func NewSearchDelivery(repository SearchRepository, errCodeGetter errs.GetErrorCode, logger *slog.Logger) *SearchHandler {
	return &SearchHandler{
		searchRepository: repository,
		errResolver:      errCodeGetter,
		log:              logger,
	}
}
