package reviews

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/pool"
	"log/slog"
)

type ReviewsStore struct {
	db  pool.DBPool
	log *slog.Logger
}

func NewReviewsStore(pool pool.DBPool, logger *slog.Logger) *ReviewsStore {
	return &ReviewsStore{
		db:  pool,
		log: logger,
	}
}
