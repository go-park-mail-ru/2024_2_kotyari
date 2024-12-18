package wish_list

//import (
//	"context"
//	"errors"
//	"testing"
//	"time"
//
//	"github.com/stretchr/testify/require"
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
//	"go.mongodb.org/mongo-driver/mongo/options"
//)
//
//func (wlr *WishListRepo) AddProductToWishlists(ctx context.Context, userID uint32, links []string, productID uint32) error {
//	newItem := dtoWishlistItem{
//		ProductID: productID,
//		AddedAt:   time.Now(),
//	}
//
//	for _, link := range links {
//		filter := bson.M{
//			"user_id":        userID,
//			"wishlists.link": link,
//		}
//		update := bson.M{
//			"$push": bson.M{
//				"wishlists.$.items": newItem,
//			},
//		}
//		_, err := wlr.collection.UpdateOne(ctx, filter, update)
//		if err != nil {
//			return errors.New("failed to update MongoDB: " + err.Error())
//		}
//	}
//
//	return nil
//}
//
//func (wlr *WishListRepo) CopyWishlist(ctx context.Context, sourceUserID uint32, sourceLink string, targetUserID uint32, newLink string) error {
//	var doc dtoUserWishLists
//
//	err := wlr.collection.FindOne(ctx, bson.M{"user_id": sourceUserID}).Decode(&doc)
//	if err != nil {
//		return errors.New("failed to query MongoDB for source user_id: " + err.Error())
//	}
//
//	var sourceWishlist *dtoWishlist
//	for _, wl := range doc.Wishlists {
//		if wl.Link == sourceLink {
//			sourceWishlist = &wl
//			break
//		}
//	}
//	if sourceWishlist == nil {
//		return errors.New("wishlist not found")
//	}
//
//	copiedWishlist := dtoWishlist{
//		Name:  sourceWishlist.Name,
//		Link:  newLink,
//		Items: sourceWishlist.Items,
//	}
//	filter := bson.M{"user_id": targetUserID}
//	update := bson.M{"$push": bson.M{"wishlists": copiedWishlist}}
//	_, err = wlr.collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
//	if err != nil {
//		return errors.New("failed to insert new wishlist into MongoDB: " + err.Error())
//	}
//
//	return nil
//}
//
//func (wlr *WishListRepo) CreateWishlist(ctx context.Context, userID uint32, name string, link string) error {
//	newWishlist := dtoWishlist{
//		Name:  name,
//		Link:  link,
//		Items: []dtoWishlistItem{},
//	}
//
//	filter := bson.M{"user_id": userID}
//	update := bson.M{"$push": bson.M{"wishlists": newWishlist}}
//	_, err := wlr.collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
//	if err != nil {
//		return errors.New("failed to insert into MongoDB: " + err.Error())
//	}
//
//	return nil
//}
//
//func (wlr *WishListRepo) DeleteWishlist(ctx context.Context, userID uint32, link string) error {
//	filter := bson.M{"user_id": userID}
//	update := bson.M{"$pull": bson.M{"wishlists": bson.M{"link": link}}}
//	_, err := wlr.collection.UpdateOne(ctx, filter, update)
//	if err != nil {
//		return errors.New("failed to update MongoDB: " + err.Error())
//	}
//
//	return nil
//}
//
//func TestAddProductToWishlists(t *testing.T) {
//	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
//
//	mt.Run("Add product to existing wishlist", func(mt *mtest.T) {
//		mt.AddMockResponses(mtest.CreateSuccessResponse())
//
//		repo := &WishListRepo{collection: mt.Coll}
//		ctx := context.Background()
//
//		err := repo.AddProductToWishlists(ctx, 1, []string{"link1", "link2"}, 1001)
//		require.NoError(t, err)
//	})
//
//	mt.Run("Error when MongoDB update fails", func(mt *mtest.T) {
//		mt.AddMockResponses(mtest.CreateWriteErrorResponse(errors.New("update error")))
//
//		repo := &WishListRepo{collection: mt.Coll}
//		ctx := context.Background()
//
//		err := repo.AddProductToWishlists(ctx, 1, []string{"link1"}, 1001)
//		require.Error(t, err)
//	})
//}
//
//func TestCopyWishlist(t *testing.T) {
//	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
//
//	mt.Run("Copy wishlist successfully", func(mt *mtest.T) {
//		mt.AddMockResponses(
//			mtest.CreateCursorResponse(1, "testdb.testcoll", mtest.FirstBatch, bson.D{
//				{Key: "user_id", Value: 1},
//				{Key: "wishlists", Value: []bson.M{
//					{"link": "sourceLink", "name": "Source Wishlist", "items": []bson.M{}},
//				}},
//			}),
//			mtest.CreateSuccessResponse(),
//		)
//
//		repo := &WishListRepo{collection: mt.Coll}
//		ctx := context.Background()
//
//		err := repo.CopyWishlist(ctx, 1, "sourceLink", 2, "newLink")
//		require.NoError(t, err)
//	})
//
//	mt.Run("Source wishlist not found", func(mt *mtest.T) {
//		mt.AddMockResponses(
//			mtest.CreateCursorResponse(1, "testdb.testcoll", mtest.FirstBatch, bson.D{
//				{Key: "user_id", Value: 1},
//				{Key: "wishlists", Value: []bson.M{}},
//			}),
//		)
//
//		repo := &WishListRepo{collection: mt.Coll}
//		ctx := context.Background()
//
//		err := repo.CopyWishlist(ctx, 1, "nonExistentLink", 2, "newLink")
//		require.Error(t, err)
//	})
//}
//
//func TestCreateWishlist(t *testing.T) {
//	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
//
//	mt.Run("Create new wishlist successfully", func(mt *mtest.T) {
//		mt.AddMockResponses(mtest.CreateSuccessResponse())
//
//		repo := &WishListRepo{collection: mt.Coll}
//		ctx := context.Background()
//
//		err := repo.CreateWishlist(ctx, 1, "New Wishlist", "newLink")
//		require.NoError(t, err)
//	})
//
//	mt.Run("Error when MongoDB update fails", func(mt *mtest.T) {
//		mt.AddMockResponses(mtest.CreateWriteErrorResponse(errors.New("update error")))
//
//		repo := &WishListRepo{collection: mt.Coll}
//		ctx := context.Background()
//
//		err := repo.CreateWishlist(ctx, 9999, "New Wishlist", "newLink")
//		require.Error(t, err)
//	})
//}
//
//func TestDeleteWishlist(t *testing.T) {
//	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
//
//	mt.Run("Delete existing wishlist successfully", func(mt *mtest.T) {
//		mt.AddMockResponses(mtest.CreateSuccessResponse())
//
//		repo := &WishListRepo{collection: mt.Coll}
//		ctx := context.Background()
//
//		err := repo.DeleteWishlist(ctx, 1, "existingLink")
//		require.NoError(t, err)
//	})
//
//	mt.Run("Error when MongoDB update fails", func(mt *mtest.T) {
//		mt.AddMockResponses(mtest.CreateWriteErrorResponse(errors.New("delete error")))
//
//		repo := &WishListRepo{collection: mt.Coll}
//		ctx := context.Background()
//
//		err := repo.DeleteWishlist(ctx, 9999, "nonExistingLink")
//		require.Error(t, err)
//	})
//}
