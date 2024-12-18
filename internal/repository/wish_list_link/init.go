package wish_list_link

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/pool"
)

type WishListLinksRepo struct {
	db pool.DBPool
}

func NewWishListLinkRepo(db pool.DBPool) *WishListLinksRepo {
	return &WishListLinksRepo{
		db: db,
	}
}
