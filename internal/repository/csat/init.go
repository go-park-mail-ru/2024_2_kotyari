package csat

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CSATStore struct {
	db  *pgxpool.Pool
	log *slog.Logger
}

func NewSurveyStore(pool *pgxpool.Pool, logger *slog.Logger) *CSATStore {
	return &CSATStore{
		db:  pool,
		log: logger,
	}
}
