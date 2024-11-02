package profile

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProfilesStore struct {
	db  *pgxpool.Pool
	log *slog.Logger
}

func NewProfileRepo(db *pgxpool.Pool, logger *slog.Logger) *ProfilesStore {
	return &ProfilesStore{
		db:  db,
		log: logger,
	}
}

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
