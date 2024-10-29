package user

import "github.com/jackc/pgx/v5/pgxpool"

type UsersStore struct {
	db *pgxpool.Pool
}

func NewUsersStore(db *pgxpool.Pool) *UsersStore {
	return &UsersStore{db: db}
}
