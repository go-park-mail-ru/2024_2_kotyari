package wish_list

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/wish_list/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"io"
	"log/slog"
	"testing"
)

func TestWishListUsecase_AddProductToWishlists(t *testing.T) {
	type want error

	var tests = []struct {
		name        string
		userSession uint32
		links       []string
		productID   uint32
		setupFunc   func(ctrl *gomock.Controller) *WishListUsecase
		want        want
	}{
		{
			name:        "Успешное добавление",
			userSession: 1,
			links:       []string{"123"},
			productID:   1,
			setupFunc: func(ctrl *gomock.Controller) *WishListUsecase {
				wishListRep := mocks.NewMockwishListRepo(ctrl)
				wishListLinkRep := mocks.NewMockwishListLinkRepo(ctrl)

				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				wishListLinkRep.EXPECT().GetUserIDFromLink(
					gomock.Any(), "123").Return(uint32(1), nil)

				wishListRep.EXPECT().AddProductToWishlists(
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
		{
			name:        "Успешное добавление",
			userSession: 1,
			links:       []string{"123"},
			productID:   1,
			setupFunc: func(ctrl *gomock.Controller) *WishListUsecase {
				wishListRep := mocks.NewMockwishListRepo(ctrl)
				wishListLinkRep := mocks.NewMockwishListLinkRepo(ctrl)

				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				wishListLinkRep.EXPECT().GetUserIDFromLink(
					gomock.Any(), "123").Return(uint32(1), nil)

				wishListRep.EXPECT().AddProductToWishlists(
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
			name:        "Успешное добавление",
			userSession: 1,
			links:       nil,
			productID:   1,
			setupFunc: func(ctrl *gomock.Controller) *WishListUsecase {
				wishListRep := mocks.NewMockwishListRepo(ctrl)
				wishListLinkRep := mocks.NewMockwishListLinkRepo(ctrl)

				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				return &WishListUsecase{
					wishListRepo:     wishListRep,
					wishListLinkRepo: wishListLinkRep,
					log:              logger,
				}
			},
			want: nil,
		},
		{
			name:        "Успешное добавление",
			userSession: 1,
			links:       []string{"123"},
			productID:   1,
			setupFunc: func(ctrl *gomock.Controller) *WishListUsecase {
				wishListRep := mocks.NewMockwishListRepo(ctrl)
				wishListLinkRep := mocks.NewMockwishListLinkRepo(ctrl)

				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				wishListLinkRep.EXPECT().GetUserIDFromLink(
					gomock.Any(), "123").Return(uint32(0), e)

				return &WishListUsecase{
					wishListRepo:     wishListRep,
					wishListLinkRepo: wishListLinkRep,
					log:              logger,
				}
			},
			want: e,
		},
		{
			name:        "Пользователь пытается добавить не в свой вишлист",
			userSession: 1,
			links:       []string{"123"},
			productID:   1,
			setupFunc: func(ctrl *gomock.Controller) *WishListUsecase {
				wishListRep := mocks.NewMockwishListRepo(ctrl)
				wishListLinkRep := mocks.NewMockwishListLinkRepo(ctrl)

				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				wishListLinkRep.EXPECT().GetUserIDFromLink(
					gomock.Any(), "123").Return(uint32(2), nil)

				return &WishListUsecase{
					wishListRepo:     wishListRep,
					wishListLinkRepo: wishListLinkRep,
					log:              logger,
				}
			},
			want: errs.ErrNotPermitted,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.Background()

			resp := tt.setupFunc(ctrl).AddProductToWishlists(ctx, tt.userSession, tt.links, tt.productID)
			assert.Equal(t, tt.want, resp)
		})
	}
}
