package wish_list

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (wlr *WishListRepo) CopyWishlist(ctx context.Context, sourceUserID uint32, sourceLink string, targetUserID uint32, newLink string) error {
	var doc model.UserWishLists

	err := wlr.collection.FindOne(ctx, bson.M{"user_id": sourceUserID}).Decode(&doc)
	if err != nil {
		return fmt.Errorf("failed to query MongoDB for source user_id %d: %w", sourceUserID, err)
	}

	var sourceWishlist *model.Wishlist
	for _, wl := range doc.Wishlists {
		if wl.Link == sourceLink {
			sourceWishlist = &wl
			break
		}
	}
	if sourceWishlist == nil {
		return fmt.Errorf("wishlist with link %s not found", sourceLink)
	}

	copiedWishlist := model.Wishlist{
		Name:  sourceWishlist.Name,
		Link:  newLink,
		Items: sourceWishlist.Items,
	}
	filter := bson.M{"user_id": targetUserID}
	update := bson.M{"$push": bson.M{"wishlists": copiedWishlist}}
	_, err = wlr.collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("failed to insert new wishlist into MongoDB: %w", err)
	}

	return nil
}
