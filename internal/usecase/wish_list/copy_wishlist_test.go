package wish_list

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/wish_list/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"io"
	"log/slog"
	"testing"
)

var (
	returnsUUID = uuid.NewString()
	e           = errors.New("ошибка")
)

func TestWishListUsecase_CopyWishList(t *testing.T) {
	t.Parallel()

	type want struct {
		err error
		str string
	}

	var tests = []struct {
		name         string
		link         string
		targetUserId uint32
		setupFunc    func(ctrl *gomock.Controller) *WishListUsecase
		want         want
	}{
		//{
		//	name:         "Успешное копирование",
		//	link:         "123",
		//	targetUserId: 2,
		//	setupFunc: func(ctrl *gomock.Controller) *WishListUsecase {
		//		wishListRep := mocks.NewMockwishListRepo(ctrl)
		//		wishListLinkRep := mocks.NewMockwishListLinkRepo(ctrl)
		//
		//		logger := slog.New(slog.NewTextHandler(io.Discard, nil))
		//
		//		wishListLinkRep.EXPECT().GetUserIDFromLink(
		//			gomock.Any(), "123").Return(uint32(1), nil)
		//
		//		newLink := "bd7f1c3d-4191-4a81-ad34-af49b3fb529e"
		//		wishListRep.EXPECT().CopyWishlist(
		//			gomock.Any(),
		//			uint32(1),
		//			"123",
		//			uint32(2),
		//			newLink,
		//		).Return(nil)
		//
		//		return &WishListUsecase{
		//			wishListRepo:     wishListRep,
		//			wishListLinkRepo: wishListLinkRep,
		//			log:              logger,
		//		}
		//	},
		//	want: want{
		//		err: nil,
		//		str: "bd7f1c3d-4191-4a81-ad34-af49b3fb529e",
		//	},
		//},
		{
			name:         "Успешное копирование",
			link:         "123",
			targetUserId: 2,
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
			want: want{
				err: e,
				str: "",
			},
		},
		{
			name:         "Успешное копирование",
			link:         "123",
			targetUserId: uint32(22),
			setupFunc: func(ctrl *gomock.Controller) *WishListUsecase {
				wishListRep := mocks.NewMockwishListRepo(ctrl)
				wishListLinkRep := mocks.NewMockwishListLinkRepo(ctrl)

				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				wishListLinkRep.EXPECT().GetUserIDFromLink(
					gomock.Any(), "123").Return(uint32(1), nil)

				wishListRep.EXPECT().CopyWishlist(
					gomock.Any(),
					uint32(1),
					"123",
					uint32(22),
					gomock.Any(),
				).Return(e)

				return &WishListUsecase{
					wishListRepo:     wishListRep,
					wishListLinkRepo: wishListLinkRep,
					log:              logger,
				}
			},
			want: want{
				err: e,
				str: "",
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

			resp, err := tt.setupFunc(ctrl).CopyWishList(ctx, tt.link, tt.targetUserId)

			assert.Equal(t, tt.want.err, err)
			assert.Equal(t, tt.want.str, resp)
		})
	}
}
