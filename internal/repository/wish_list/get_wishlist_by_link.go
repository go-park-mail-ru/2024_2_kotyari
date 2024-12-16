package wish_list

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"log/slog"
)

func (wlr *WishListRepo) GetWishListByLink(ctx context.Context, userID uint32, link string) (model.Wishlist, error) {
	var doc dtoUserWishLists
	err := wlr.collection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&doc)
	if err != nil {
		return model.Wishlist{}, fmt.Errorf("failed to query MongoDB: %w", err)
	}

	wlr.log.Info("[ WishListRepo.GetWishListByLink ]",
		slog.Any("doc", doc))

	for _, wl := range doc.Wishlists {
		if wl.Link == link {

			wlr.log.Info("[ WishListRepo.GetWishListByLink ]",
				slog.Any("wl", wl))

			items := make([]model.WishlistItem, len(wl.Items))
			for i, item := range wl.Items {
				items[i] = model.WishlistItem{
					ProductID: item.ProductID,
					AddedAt:   item.AddedAt,
				}
			}

			return model.Wishlist{
				Link:  link,
				Name:  wl.Name,
				Items: items,
			}, nil
		}
	}

	return model.Wishlist{}, fmt.Errorf("wishlist not found in MongoDB")
}
