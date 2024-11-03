package address

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
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
