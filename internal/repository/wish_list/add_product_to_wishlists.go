package wish_list

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (wlr *WishListRepo) AddProductToWishlists(ctx context.Context, userID uint32, links []string, productID uint32) error {
	newItem := dtoWishlistItem{
		ProductID: productID,
		AddedAt:   time.Now(),
	}

	for _, link := range links {
		filter := bson.M{
			"user_id":        userID,
			"wishlists.link": link,
		}
		update := bson.M{
			"$push": bson.M{
				"wishlists.$.items": newItem,
			},
		}
		_, err := wlr.collection.UpdateOne(ctx, filter, update)

		if err != nil {
			return fmt.Errorf("failed to update MongoDB for link %s: %w", link, err)
		}
	}

	return nil
}
