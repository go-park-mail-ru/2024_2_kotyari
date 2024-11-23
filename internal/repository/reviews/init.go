package reviews

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ReviewsStore struct {
	db  *pgxpool.Pool
	log *slog.Logger
}

func NewReviewsStore(pool *pgxpool.Pool, logger *slog.Logger) *ReviewsStore {
	return &ReviewsStore{
		db:  pool,
		log: logger,
	}
}
