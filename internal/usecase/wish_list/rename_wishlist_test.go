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

func TestWishListUsecase_RenameWishList(t *testing.T) {
	t.Parallel()

	type want error

	var tests = []struct {
		name      string
		userId    uint32
		link      string
		newName   string
		setupFunc func(ctrl *gomock.Controller) *WishListUsecase
		want      want
	}{
		{
			name:    "1",
			userId:  uint32(123),
			link:    "123",
			newName: "1234",
			setupFunc: func(ctrl *gomock.Controller) *WishListUsecase {
				wishListRep := mocks.NewMockwishListRepo(ctrl)
				wishListLinkRep := mocks.NewMockwishListLinkRepo(ctrl)

				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				wishListRep.EXPECT().RenameWishlist(
					gomock.Any(),
					uint32(123),
					"1234",
					"123",
				).Return(e)

				return &WishListUsecase{
					wishListRepo:     wishListRep,
					wishListLinkRepo: wishListLinkRep,
					log:              logger,
				}
			}, want: e,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.Background()

			err := tt.setupFunc(ctrl).RenameWishList(ctx, tt.userId, tt.newName, tt.link)
			assert.Equal(t, tt.want, err)
		})
	}

}
