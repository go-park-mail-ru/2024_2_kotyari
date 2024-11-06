package product

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductsStore struct {
	db  *pgxpool.Pool
	log *slog.Logger
}

func NewProductsStore(db *pgxpool.Pool, log *slog.Logger) *ProductsStore {
	return &ProductsStore{
		db:  db,
		log: log,
	}
}
