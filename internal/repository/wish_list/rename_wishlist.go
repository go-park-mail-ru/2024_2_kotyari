package wish_list

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (wlr *WishListRepo) RenameWishlist(ctx context.Context, userID uint32, newName string, link string) error {
	filter := bson.M{
		"user_id":        userID,
		"wishlists.link": link,
	}

	projection := bson.M{"wishlists.$": 1}

	var result struct {
		Wishlists []struct {
			Link string `bson:"link"`
			Name string `bson:"name"`
		} `bson:"wishlists"`
	}

	err := wlr.collection.FindOne(ctx, filter, options.FindOne().SetProjection(projection)).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fmt.Errorf("wishlist not found for link: %s", link)
		}
		return fmt.Errorf("failed to find wishlist: %w", err)
	}

	if len(result.Wishlists) == 0 {
		return fmt.Errorf("wishlist not found for link: %s", link)
	}

	update := bson.M{
		"$set": bson.M{
			"wishlists.$.name": newName,
		},
	}

	_, err = wlr.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update MongoDB: %w", err)
	}

	return nil
}
