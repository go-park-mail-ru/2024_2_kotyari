package wish_list

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/wish_list/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"io"
	"log/slog"
	"testing"
)

func TestWishListUsecase_DeleteWishlist(t *testing.T) {
	t.Parallel()

	type want error

	var tests = []struct {
		name      string
		userId    uint32
		setupFunc func(ctrl *gomock.Controller) *WishListUsecase
		want      want
	}{
		{
			name:   "1",
			userId: 123,
			setupFunc: func(ctrl *gomock.Controller) *WishListUsecase {
				wishListRep := mocks.NewMockwishListRepo(ctrl)
				wishListLinkRep := mocks.NewMockwishListLinkRepo(ctrl)

				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				wishListRep.EXPECT().DeleteWishlist(
					gomock.Any(),
					uint32(123),
					gomock.Any(),
				).Return(e)

				return &WishListUsecase{
					wishListRepo:     wishListRep,
					wishListLinkRepo: wishListLinkRep,
					log:              logger,
				}
			}, want: e,
		},
		{
			name:   "2",
			userId: 123,
			setupFunc: func(ctrl *gomock.Controller) *WishListUsecase {
				wishListRep := mocks.NewMockwishListRepo(ctrl)
				wishListLinkRep := mocks.NewMockwishListLinkRepo(ctrl)

				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				wishListRep.EXPECT().DeleteWishlist(
					gomock.Any(),
					uint32(123),
					gomock.Any(),
				).Return(nil)

				wishListLinkRep.EXPECT().DeleteWishListLink(
					gomock.Any(),
					gomock.Any(),
				).Return(e)

				return &WishListUsecase{
					wishListRepo:     wishListRep,
					wishListLinkRepo: wishListLinkRep,
					log:              logger,
				}
			}, want: e,
		},
		{
			name:   "3",
			userId: 123,
			setupFunc: func(ctrl *gomock.Controller) *WishListUsecase {
				wishListRep := mocks.NewMockwishListRepo(ctrl)
				wishListLinkRep := mocks.NewMockwishListLinkRepo(ctrl)

				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				wishListRep.EXPECT().DeleteWishlist(
					gomock.Any(),
					uint32(123),
					gomock.Any(),
				).Return(nil)

				wishListLinkRep.EXPECT().DeleteWishListLink(
					gomock.Any(),
					gomock.Any(),
				).Return(nil)

				return &WishListUsecase{
					wishListRepo:     wishListRep,
					wishListLinkRepo: wishListLinkRep,
					log:              logger,
				}
			}, want: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.Background()

			resp := tt.setupFunc(ctrl).DeleteWishlist(ctx, tt.userId, "123")
			assert.Equal(t, tt.want, resp)
		})
	}

}
