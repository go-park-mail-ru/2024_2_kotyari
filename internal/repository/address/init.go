package address

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

type AddressStore struct {
	db  *pgxpool.Pool
	log *slog.Logger
}

func NewAddressRepo(db *pgxpool.Pool, logger *slog.Logger) *AddressStore {
	return &AddressStore{
		db:  db,
		log: logger,
	}
}
