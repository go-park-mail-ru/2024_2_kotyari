package search

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

type SearchStore struct {
	db  *pgxpool.Pool
	log *slog.Logger
}

func NewSearchStore(pool *pgxpool.Pool, logger *slog.Logger) *SearchStore {
	return &SearchStore{
		db:  pool,
		log: logger,
	}
}
