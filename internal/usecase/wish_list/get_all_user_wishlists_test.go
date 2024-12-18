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
)

func TestWishListUsecase_GetALlUserWishlists(t *testing.T) {
	t.Parallel()

	type want struct {
		err   error
		model []model.Wishlist
	}

	var tests = []struct {
		name      string
		userId    uint32
		setupFunc func(ctrl *gomock.Controller) *WishListUsecase
		want      want
	}{
		{
			name:   "1",
			userId: uint32(123),
			setupFunc: func(ctrl *gomock.Controller) *WishListUsecase {
				wishListRep := mocks.NewMockwishListRepo(ctrl)
				wishListLinkRep := mocks.NewMockwishListLinkRepo(ctrl)

				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				wishListRep.EXPECT().GetALlUserWishlists(
					gomock.Any(),
					uint32(123),
				).Return([]model.Wishlist{}, e)

				return &WishListUsecase{
					wishListRepo:     wishListRep,
					wishListLinkRepo: wishListLinkRep,
					log:              logger,
				}
			}, want: want{
				err:   e,
				model: []model.Wishlist{},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.Background()

			resp, err := tt.setupFunc(ctrl).GetALlUserWishlists(ctx, tt.userId)
			assert.Equal(t, tt.want.model, resp)
			assert.Equal(t, tt.want.err, err)
		})
	}

}
