package user

import "github.com/jackc/pgx/v5/pgxpool"

type UsersRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UsersRepo {
	return &UsersRepo{db: db}
}
