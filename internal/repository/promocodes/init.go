package promocodes

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PromoCodesStore struct {
	db  *pgxpool.Pool
	log *slog.Logger
}

func NewPromoCodesStore(pool *pgxpool.Pool, logger *slog.Logger) *PromoCodesStore {
	return &PromoCodesStore{
		db:  pool,
		log: logger,
	}
}
