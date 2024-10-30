package cart

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CartsStore struct {
	db  *pgxpool.Pool
	log *slog.Logger
}

func NewCartsStore(db *pgxpool.Pool, logger *slog.Logger) *CartsStore {
	return &CartsStore{
		db:  db,
		log: logger,
	}
}
