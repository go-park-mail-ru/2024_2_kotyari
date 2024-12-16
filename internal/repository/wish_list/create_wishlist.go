package wish_list

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (wlr *WishListRepo) CreateWishlist(ctx context.Context, userID uint32, name string, link string) error {
	newWishlist := model.Wishlist{
		Name:  name,
		Link:  link,
		Items: []model.WishlistItem{},
	}

	filter := bson.M{"user_id": userID}
	update := bson.M{"$push": bson.M{"wishlists": newWishlist}}
	_, err := wlr.collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("failed to insert into MongoDB: %w", err)
	}

	return nil
}
