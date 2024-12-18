package user

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/pool"
	"log/slog"
)

type UsersStore struct {
	db  pool.DBPool
	log *slog.Logger
}

func NewUsersStore(db pool.DBPool, log *slog.Logger) *UsersStore {
	return &UsersStore{db: db,
		log: log}
}
