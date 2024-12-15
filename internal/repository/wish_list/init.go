package wish_list

import (
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
)

type WishListRepo struct {
	collection *mongo.Collection
	log        *slog.Logger
}

func NewWishListRepo(db *mongo.Database, collectionName string, log *slog.Logger) *WishListRepo {
	return &WishListRepo{
		collection: db.Collection(collectionName),
		log:        log,
	}
}
