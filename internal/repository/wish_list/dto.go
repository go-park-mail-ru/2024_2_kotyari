package wish_list

import "time"

type dtoUserWishLists struct {
	UserID    uint32        `bson:"user_id"`
	Wishlists []dtoWishlist `bson:"wishlists"`
}

type dtoWishlist struct {
	Name  string            `bson:"name"`
	Link  string            `bson:"link"`
	Items []dtoWishlistItem `bson:"items"`
}

type dtoWishlistItem struct {
	ProductID uint32    `bson:"product_id"`
	AddedAt   time.Time `bson:"added_at"`
}
