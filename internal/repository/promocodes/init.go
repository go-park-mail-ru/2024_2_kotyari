package promocodes

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/pool"
	"log/slog"
)

type PromoCodesStore struct {
	db  pool.DBPool
	log *slog.Logger
}

func NewPromoCodesStore(pool pool.DBPool, logger *slog.Logger) *PromoCodesStore {
	return &PromoCodesStore{
		db:  pool,
		log: logger,
	}
}
