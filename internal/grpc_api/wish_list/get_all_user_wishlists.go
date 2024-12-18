package wish_list

import (
	"context"
	wishlistgrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/wish_list/gen"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

func (wlg *WishlistGrpc) GetAllUserWishlists(ctx context.Context, in *wishlistgrpc.GetAllWishlistsRequest) (*wishlistgrpc.GetAllWishlistsResponse, error) {
	wishlists, err := wlg.manager.GetALlUserWishlists(ctx, in.GetUserId())
	if err != nil {
		return nil, err
	}

	out := make([]*wishlistgrpc.Wishlist, len(wishlists))

	for i, wishlist := range wishlists {
		out[i] = &wishlistgrpc.Wishlist{
			Link:  wishlist.Link,
			Name:  wishlist.Name,
			Items: make([]*wishlistgrpc.WishlistItem, 0, len(wishlist.Items)),
		}

		for _, item := range wishlist.Items {
			out[i].Items = append(out[i].Items, &wishlistgrpc.WishlistItem{
				ProductId: item.ProductID,
				AddedAt:   timestamppb.New(item.AddedAt),
			})
		}
	}

	response := &wishlistgrpc.GetAllWishlistsResponse{
		Wishlists: out,
	}

	wlg.log.Info("[ WishlistGrpc.GetAllUserWishlists ]",
		slog.Any("response", response))

	return response, nil
}
