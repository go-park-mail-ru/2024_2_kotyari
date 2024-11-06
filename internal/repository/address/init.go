package address

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AddressStore struct {
	Db  *pgxpool.Pool
	Log *slog.Logger
}

func NewAddressRepo(db *pgxpool.Pool, logger *slog.Logger) *AddressStore {
	return &AddressStore{
		Db:  db,
		Log: logger,
	}
}
