package wish_list

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

func (wlr *WishListRepo) RemoveProductFromWishlists(ctx context.Context, userID uint32, links []string, productID uint32) error {
	for _, link := range links {
		filter := bson.M{"user_id": userID, "wishlists.link": link}
		update := bson.M{"$pull": bson.M{"wishlists.$.items": bson.M{"product_id": productID}}}
		_, err := wlr.collection.UpdateOne(ctx, filter, update)
		if err != nil {
			return fmt.Errorf("failed to update MongoDB for link %s: %w", link, err)
		}
	}

	return nil
}
