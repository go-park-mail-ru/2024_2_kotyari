package category

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoriesStore struct {
	db  *pgxpool.Pool
	log *slog.Logger
}

func NewCategoriesStore(db *pgxpool.Pool, log *slog.Logger) *CategoriesStore {
	return &CategoriesStore{
		db:  db,
		log: log,
	}
}
