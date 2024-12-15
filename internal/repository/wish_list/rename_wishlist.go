package wish_list

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

func (wlr *WishListRepo) RenameWishlist(ctx context.Context, userID uint32, newName string, link string) error {
	filter := bson.M{"user_id": userID, "wishlists.link": link}
	update := bson.M{"$set": bson.M{"wishlists.$.name": newName}}
	_, err := wlr.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update MongoDB: %w", err)
	}

	return nil
}
