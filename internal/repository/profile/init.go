package profile

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/pool"
	"log/slog"
)

type ProfilesStore struct {
	db  pool.DBPool
	log *slog.Logger
}

func NewProfileRepo(db pool.DBPool, logger *slog.Logger) *ProfilesStore {
	return &ProfilesStore{
		db:  db,
		log: logger,
	}
}
