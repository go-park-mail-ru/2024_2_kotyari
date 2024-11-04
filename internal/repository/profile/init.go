package profile

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProfilesStore struct {
	db  *pgxpool.Pool
	log *slog.Logger
}

func NewProfileRepo(db *pgxpool.Pool, logger *slog.Logger) *ProfilesStore {
	return &ProfilesStore{
		db:  db,
		log: logger,
	}
}
