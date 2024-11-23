package repository

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SurveyStore struct {
	db  *pgxpool.Pool
	log *slog.Logger
}

func NewSurveyStore(pool *pgxpool.Pool, logger *slog.Logger) *SurveyStore {
	return &SurveyStore{
		db:  pool,
		log: logger,
	}
}
