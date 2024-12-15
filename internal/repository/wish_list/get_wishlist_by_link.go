package wish_list

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (wlr *WishListRepo) GetWishListByLink(ctx context.Context, userID uint32, link string) (model.Wishlist, error) {
	var doc model.UserWishLists
	err := wlr.collection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&doc)
	if err != nil {
		return model.Wishlist{}, fmt.Errorf("failed to query MongoDB: %w", err)
	}

	for _, wl := range doc.Wishlists {
		if wl.Link == link {
			return wl, nil
		}
	}

	return model.Wishlist{}, fmt.Errorf("wishlist not found in MongoDB")
}
