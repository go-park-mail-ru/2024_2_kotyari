package user

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

type UsersStore struct {
	db  *pgxpool.Pool
	log *slog.Logger
}

func NewUsersStore(db *pgxpool.Pool, log *slog.Logger) *UsersStore {
	return &UsersStore{db: db,
		log: log}
}
