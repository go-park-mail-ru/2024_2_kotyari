package product

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/pool"
	"log/slog"
)

type ProductsStore struct {
	db  pool.DBPool
	log *slog.Logger
}

func NewProductsStore(db pool.DBPool, log *slog.Logger) *ProductsStore {
	return &ProductsStore{
		db:  db,
		log: log,
	}
}
