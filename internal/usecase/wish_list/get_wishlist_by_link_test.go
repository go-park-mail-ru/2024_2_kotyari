package wish_list

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/wish_list/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"io"
	"log/slog"
	"testing"
	"time"
)

func TestWishListUsecase_GetWishListByLink(t *testing.T) {
	t.Parallel()

	type want struct {
		Wish   model.Wishlist
		UserId uint32
		err    error
	}

	var tests = []struct {
		name      string
		link      string
		setupFunc func(ctrl *gomock.Controller) *WishListUsecase
		want      want
	}{
		{
			name: "1",
			link: "123",
			setupFunc: func(ctrl *gomock.Controller) *WishListUsecase {
				wishListRep := mocks.NewMockwishListRepo(ctrl)
				wishListLinkRep := mocks.NewMockwishListLinkRepo(ctrl)

				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				wishListLinkRep.EXPECT().GetUserIDFromLink(
					gomock.Any(),
					"123",
				).Return(uint32(0), e)

				return &WishListUsecase{
					wishListRepo:     wishListRep,
					wishListLinkRepo: wishListLinkRep,
					log:              logger,
				}
			},
			want: want{
				err:    e,
				UserId: uint32(0),
				Wish:   model.Wishlist{},
			},
		},
		{
			name: "2",
			link: "123",
			setupFunc: func(ctrl *gomock.Controller) *WishListUsecase {
				wishListRep := mocks.NewMockwishListRepo(ctrl)
				wishListLinkRep := mocks.NewMockwishListLinkRepo(ctrl)

				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				wishListLinkRep.EXPECT().GetUserIDFromLink(
					gomock.Any(),
					"123",
				).Return(uint32(1), nil)

				wishListRep.EXPECT().GetWishListByLink(
					gomock.Any(),
					uint32(1),
					"123",
				).Return(model.Wishlist{}, e)

				return &WishListUsecase{
					wishListRepo:     wishListRep,
					wishListLinkRepo: wishListLinkRep,
					log:              logger,
				}
			}, want: want{
				err:    e,
				UserId: uint32(0),
				Wish:   model.Wishlist{},
			},
		},
		{
			name: "2",
			link: "123",
			setupFunc: func(ctrl *gomock.Controller) *WishListUsecase {
				wishListRep := mocks.NewMockwishListRepo(ctrl)
				wishListLinkRep := mocks.NewMockwishListLinkRepo(ctrl)

				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				wishListLinkRep.EXPECT().GetUserIDFromLink(
					gomock.Any(),
					"123",
				).Return(uint32(1), nil)

				wishListRep.EXPECT().GetWishListByLink(
					gomock.Any(),
					uint32(1),
					"123",
				).Return(model.Wishlist{
					Name: "1",
					Link: "123",
					Items: []model.WishlistItem{
						{
							ProductID: uint32(1),
							AddedAt:   time.Unix(0, 0),
						},
						{
							ProductID: uint32(1),
							AddedAt:   time.Unix(0, 0),
						},
					},
				}, nil)

				return &WishListUsecase{
					wishListRepo:     wishListRep,
					wishListLinkRepo: wishListLinkRep,
					log:              logger,
				}
			}, want: want{
				err:    nil,
				UserId: uint32(1),
				Wish: model.Wishlist{
					Name: "1",
					Link: "123",
					Items: []model.WishlistItem{
						{
							ProductID: uint32(1),
							AddedAt:   time.Unix(0, 0),
						},
						{
							ProductID: uint32(1),
							AddedAt:   time.Unix(0, 0),
						},
					},
				},
			},
		}}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.Background()

			wish, userId, err := tt.setupFunc(ctrl).GetWishListByLink(ctx, tt.link)
			assert.Equal(t, tt.want.Wish, wish)
			assert.Equal(t, tt.want.err, err)
			assert.Equal(t, tt.want.UserId, userId)
		})
	}
}
