package wish_list

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

func (wlr *WishListRepo) DeleteWishlist(ctx context.Context, userID uint32, link string) error {
	filter := bson.M{"user_id": userID}
	update := bson.M{"$pull": bson.M{"wishlists": bson.M{"link": link}}}
	_, err := wlr.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update MongoDB: %w", err)
	}

	return nil
}
