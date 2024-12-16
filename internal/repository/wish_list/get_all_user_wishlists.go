package wish_list

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
)

func (wlr *WishListRepo) GetALlUserWishlists(ctx context.Context, userID uint32) ([]model.Wishlist, error) {
	wlr.log.Info("[ WishListRepo.GetALlUserWishlists ]", slog.Any("userId", userID))

	filter := bson.M{"user_id": userID}

	var result dtoUserWishLists

	err := wlr.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return []model.Wishlist{}, fmt.Errorf("отсутствуют вишлисты")
		}

		wlr.log.Error("[ WishListRepo.GetALlUserWishlists ] ",
			"func", "FindOne",
			slog.Any("err", err))

		return nil, fmt.Errorf("внутренняя ошибка сервера %w", err)
	}

	wlr.log.Info("[ WishListRepo.GetALlUserWishlists ]",
		slog.Any("userId", userID),
		slog.Any("result", result))

	if len(result.Wishlists) == 0 {
		return []model.Wishlist{}, nil
	}

	res := make([]model.Wishlist, len(result.Wishlists))

	for i, item := range result.Wishlists {
		res[i] = model.Wishlist{
			Link: item.Link,
			Name: item.Name,
		}

		items := make([]model.WishlistItem, len(item.Items))
		for j, it := range item.Items {
			items[j] = model.WishlistItem{
				ProductID: it.ProductID,
				AddedAt:   it.AddedAt,
			}
		}
		res[i].Items = items
	}

	wlr.log.Info("[ WishListRepo.GetALlUserWishlists ]",
		slog.Any("userId", userID),
		slog.Any("res", res))

	return res, nil
}
