package cart

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/pool"
	"log/slog"
)

type CartsStore struct {
	db  pool.DBPool
	log *slog.Logger
}

func NewCartsStore(db pool.DBPool, logger *slog.Logger) *CartsStore {
	return &CartsStore{
		db:  db,
		log: logger,
	}
}
