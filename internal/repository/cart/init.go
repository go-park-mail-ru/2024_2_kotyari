package cart

import "github.com/jackc/pgx/v5/pgxpool"

type CartsStore struct {
	db *pgxpool.Pool
}

func NewCartsStore(db *pgxpool.Pool) *CartsStore {
	return &CartsStore{
		db: db,
	}
}
