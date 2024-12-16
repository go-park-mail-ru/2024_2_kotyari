package wish_list_link

import "github.com/jackc/pgx/v5/pgxpool"

type WishListLinksRepo struct {
	db *pgxpool.Pool
}

func NewWishListLinkRepo(db *pgxpool.Pool) *WishListLinksRepo {
	return &WishListLinksRepo{
		db: db,
	}
}
