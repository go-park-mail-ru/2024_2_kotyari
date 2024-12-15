package search

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/pool"
	"log/slog"
)

type SearchStore struct {
	db  pool.DBPool
	log *slog.Logger
}

func NewSearchStore(db pool.DBPool, logger *slog.Logger) *SearchStore {
	return &SearchStore{
		db:  db,
		log: logger,
	}
}
