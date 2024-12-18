package wish_list

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/wish_list/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"log/slog"
	"testing"
)

func TestWishListUsecase_RemoveFromWishlists(t *testing.T) {
	t.Parallel()

	type want error

	var tests = []struct {
		name      string
		userID    uint32
		links     []string
		productId uint32
		setupFunc func(ctrl *gomock.Controller) *WishListUsecase
		want      want
	}{
		{
			name:      "1",
			userID:    uint32(1),
			productId: uint32(1),
			links:     []string{"123"},
			setupFunc: func(ctrl *gomock.Controller) *WishListUsecase {
				wishListRep := mocks.NewMockwishListRepo(ctrl)
				wishListLinkRep := mocks.NewMockwishListLinkRepo(ctrl)

				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				wishListRep.EXPECT().RemoveProductFromWishlists(
					gomock.Any(),
					uint32(1),
					[]string{"123"},
					uint32(1),
				).Return(e)

				return &WishListUsecase{
					wishListRepo:     wishListRep,
					wishListLinkRepo: wishListLinkRep,
					log:              logger,
				}
			},
			want: e,
		},
		{
			name:      "1",
			userID:    uint32(1),
			productId: uint32(1),
			links:     []string{"123"},
			setupFunc: func(ctrl *gomock.Controller) *WishListUsecase {
				wishListRep := mocks.NewMockwishListRepo(ctrl)
				wishListLinkRep := mocks.NewMockwishListLinkRepo(ctrl)

				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				wishListRep.EXPECT().RemoveProductFromWishlists(
					gomock.Any(),
					uint32(1),
					[]string{"123"},
					uint32(1),
				).Return(nil)

				return &WishListUsecase{
					wishListRepo:     wishListRep,
					wishListLinkRepo: wishListLinkRep,
					log:              logger,
				}
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.Background()

			err := tt.setupFunc(ctrl).RemoveFromWishlists(ctx, tt.userID, tt.links, tt.productId)
			assert.Equal(t, tt.want, err)
		})
	}
}
