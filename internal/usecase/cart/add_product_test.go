package cart

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"testing"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/cart/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var dbTestError = errors.New("ошибка базы данных")

func TestCartManager_AddProduct(t *testing.T) {
	t.Parallel()

	type want error

	tests := []struct {
		name      string
		productID uint32
		userID    uint32
		setupFunc func(ctrl *gomock.Controller) *CartManager
		want      want
	}{
		{
			name:      "Продукт успешно добавлен в корзину",
			productID: 1,
			userID:    1,
			setupFunc: func(ctrl *gomock.Controller) *CartManager {
				cartRepositoryMock := mocks.NewMockcartRepository(ctrl)

				cartRepositoryMock.EXPECT().GetCartProduct(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(model.CartProduct{}, errs.ProductNotInCart)

				cartRepositoryMock.EXPECT().AddProduct(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(nil)

				return &CartManager{
					cartRepository: cartRepositoryMock,
				}
			},
			want: nil,
		},
		{
			name:      "Продукт уже есть в корзине",
			productID: 1,
			userID:    1,
			setupFunc: func(ctrl *gomock.Controller) *CartManager {
				cartRepositoryMock := mocks.NewMockcartRepository(ctrl)

				cartRepositoryMock.EXPECT().GetCartProduct(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(model.CartProduct{
					IsDeleted: false,
				}, nil)

				return &CartManager{
					cartRepository: cartRepositoryMock,
				}
			},
			want: errs.ProductAlreadyInCart,
		},
		{
			name:      "Продукт уже есть в корзине",
			productID: 1,
			userID:    1,
			setupFunc: func(ctrl *gomock.Controller) *CartManager {
				cartRepositoryMock := mocks.NewMockcartRepository(ctrl)

				cartRepositoryMock.EXPECT().GetCartProduct(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(model.CartProduct{
					IsDeleted: false,
				}, nil)

				return &CartManager{
					cartRepository: cartRepositoryMock,
				}
			},
			want: errs.ProductAlreadyInCart,
		},
		{
			name:      "Продукт уже есть в корзине, но удален",
			productID: 1,
			userID:    1,
			setupFunc: func(ctrl *gomock.Controller) *CartManager {
				cartRepositoryMock := mocks.NewMockcartRepository(ctrl)

				cartRepositoryMock.EXPECT().GetCartProduct(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(model.CartProduct{
					IsDeleted: true,
				}, nil)

				cartRepositoryMock.EXPECT().ChangeCartProductDeletedState(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(nil)

				return &CartManager{
					cartRepository: cartRepositoryMock,
				}
			},
			want: nil,
		},
		{
			name:      "Произошла ошибка получения продукта из корзины",
			productID: 1,
			userID:    1,
			setupFunc: func(ctrl *gomock.Controller) *CartManager {
				cartRepositoryMock := mocks.NewMockcartRepository(ctrl)
				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				cartRepositoryMock.EXPECT().GetCartProduct(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(model.CartProduct{}, dbTestError)

				return &CartManager{
					cartRepository: cartRepositoryMock,
					log:            logger,
				}
			},
			want: dbTestError,
		},
		{
			name:      "Произошла ошибка добавления продукта в корзину",
			productID: 1,
			userID:    1,
			setupFunc: func(ctrl *gomock.Controller) *CartManager {
				cartRepositoryMock := mocks.NewMockcartRepository(ctrl)
				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				cartRepositoryMock.EXPECT().GetCartProduct(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(model.CartProduct{}, errs.ProductNotInCart)

				cartRepositoryMock.EXPECT().AddProduct(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(dbTestError)

				return &CartManager{
					cartRepository: cartRepositoryMock,
					log:            logger,
				}
			},
			want: dbTestError,
		},
		{
			name:      "Произошла ошибка изменения is_deleted у продукта",
			productID: 1,
			userID:    1,
			setupFunc: func(ctrl *gomock.Controller) *CartManager {
				cartRepositoryMock := mocks.NewMockcartRepository(ctrl)
				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				cartRepositoryMock.EXPECT().GetCartProduct(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(model.CartProduct{
					IsDeleted: true,
				}, nil)

				cartRepositoryMock.EXPECT().ChangeCartProductDeletedState(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(dbTestError)

				return &CartManager{
					cartRepository: cartRepositoryMock,
					log:            logger,
				}
			},
			want: dbTestError,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			resp := tt.setupFunc(ctrl).AddProduct(context.Background(), tt.productID, tt.userID)
			assert.Equal(t, tt.want, resp)
		})
	}
}
