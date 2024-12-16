package wish_list

import (
	"context"
	"fmt"
	wishlistgrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/wish_list/gen"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

func (wlg *WishlistGrpc) GetWishlistByLink(ctx context.Context, in *wishlistgrpc.GetWishlistByLinkRequest) (*wishlistgrpc.GetWishlistByLinkResponse, error) {
	if in.GetLink() == "" {
		return nil, fmt.Errorf("link is empty")
	}

	link, userID, err := wlg.manager.GetWishListByLink(ctx, in.GetLink())
	if err != nil {
		return nil, err
	}

	list := make([]*wishlistgrpc.WishlistItem, len(link.Items))
	for i, item := range link.Items {
		wlg.log.Info("GetWishlistByLink", slog.Any("item", item))

		list[i] = &wishlistgrpc.WishlistItem{
			ProductId: item.ProductID,
			AddedAt:   timestamppb.New(item.AddedAt),
		}
	}

	wlg.log.Info("GetWishlistByLink", slog.Any("list", list))

	wish := &wishlistgrpc.Wishlist{
		Items: list,
		Name:  link.Name,
		Link:  link.Link,
	}

	return &wishlistgrpc.GetWishlistByLinkResponse{
		Wishlist:  wish,
		CreatorId: userID,
	}, nil
}
