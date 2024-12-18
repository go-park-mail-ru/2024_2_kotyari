package address

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/pool"
	"log/slog"
)

type AddressStore struct {
	Db  pool.DBPool
	Log *slog.Logger
}

func NewAddressRepo(db pool.DBPool, logger *slog.Logger) *AddressStore {
	return &AddressStore{
		Db:  db,
		Log: logger,
	}
}
