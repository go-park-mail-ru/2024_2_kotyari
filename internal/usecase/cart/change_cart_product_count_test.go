package cart

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/cart/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"io"
	"log/slog"
	"testing"
)

func TestCartManager_ChangeCartProductCount(t *testing.T) {
	t.Parallel()

	type want error

	tests := []struct {
		name      string
		productID uint32
		count     int32
		userID    uint32
		setupFunc func(ctrl *gomock.Controller) *CartManager
		want      want
	}{
		{
			name:      "Успешное изменение количества товара в корзине",
			productID: 1,
			count:     1,
			userID:    1,
			setupFunc: func(ctrl *gomock.Controller) *CartManager {
				cartRepositoryMock := mocks.NewMockcartRepository(ctrl)
				productCountGetterMock := mocks.NewMockproductCountGetter(ctrl)

				cartProduct := model.CartProduct{
					BaseProduct: model.BaseProduct{
						Count: 10,
					},
					IsDeleted: false,
				}

				cartRepositoryMock.EXPECT().GetCartProduct(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(cartProduct, nil)

				productCountGetterMock.EXPECT().GetProductCount(
					gomock.Any(),
					uint32(1)).Return(uint32(10), nil)

				cartRepositoryMock.EXPECT().ChangeCartProductCount(
					gomock.Any(),
					uint32(1),
					int32(1),
					uint32(1)).Return(nil)

				return &CartManager{
					cartRepository:     cartRepositoryMock,
					productCountGetter: productCountGetterMock,
				}
			},
			want: nil,
		},
		{
			name:      "Не удалось получить продукт",
			productID: 1,
			count:     1,
			userID:    1,
			setupFunc: func(ctrl *gomock.Controller) *CartManager {
				cartRepositoryMock := mocks.NewMockcartRepository(ctrl)
				productCountGetterMock := mocks.NewMockproductCountGetter(ctrl)
				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				cartRepositoryMock.EXPECT().GetCartProduct(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(model.CartProduct{}, dbTestError)

				return &CartManager{
					cartRepository:     cartRepositoryMock,
					productCountGetter: productCountGetterMock,
					log:                logger,
				}
			},
			want: dbTestError,
		},
		{
			name:      "Продукта нет в корзине",
			productID: 1,
			count:     1,
			userID:    1,
			setupFunc: func(ctrl *gomock.Controller) *CartManager {
				cartRepositoryMock := mocks.NewMockcartRepository(ctrl)
				productCountGetterMock := mocks.NewMockproductCountGetter(ctrl)

				cartProduct := model.CartProduct{
					BaseProduct: model.BaseProduct{
						Count: 10,
					},
					IsDeleted: true,
				}

				cartRepositoryMock.EXPECT().GetCartProduct(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(cartProduct, nil)

				return &CartManager{
					cartRepository:     cartRepositoryMock,
					productCountGetter: productCountGetterMock,
				}
			},
			want: errs.ProductNotInCart,
		},
		{
			name:      "Ошибка получения количества продукта",
			productID: 1,
			count:     1,
			userID:    1,
			setupFunc: func(ctrl *gomock.Controller) *CartManager {
				cartRepositoryMock := mocks.NewMockcartRepository(ctrl)
				productCountGetterMock := mocks.NewMockproductCountGetter(ctrl)
				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				cartProduct := model.CartProduct{
					BaseProduct: model.BaseProduct{
						Count: 10,
					},
					IsDeleted: false,
				}

				cartRepositoryMock.EXPECT().GetCartProduct(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(cartProduct, nil)

				productCountGetterMock.EXPECT().GetProductCount(
					gomock.Any(),
					uint32(1)).Return(uint32(10), dbTestError)

				return &CartManager{
					cartRepository:     cartRepositoryMock,
					productCountGetter: productCountGetterMock,
					log:                logger,
				}
			},
			want: dbTestError,
		},
		{
			name:      "Ошибка изменения количества продукта в корзине",
			productID: 1,
			count:     1,
			userID:    1,
			setupFunc: func(ctrl *gomock.Controller) *CartManager {
				cartRepositoryMock := mocks.NewMockcartRepository(ctrl)
				productCountGetterMock := mocks.NewMockproductCountGetter(ctrl)
				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				cartProduct := model.CartProduct{
					BaseProduct: model.BaseProduct{
						Count: 10,
					},
					IsDeleted: false,
				}

				cartRepositoryMock.EXPECT().GetCartProduct(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(cartProduct, nil)

				productCountGetterMock.EXPECT().GetProductCount(
					gomock.Any(),
					uint32(1)).Return(uint32(10), nil)

				cartRepositoryMock.EXPECT().ChangeCartProductCount(
					gomock.Any(),
					uint32(1),
					int32(1),
					uint32(1)).Return(dbTestError)

				return &CartManager{
					cartRepository:     cartRepositoryMock,
					productCountGetter: productCountGetterMock,
					log:                logger,
				}
			},
			want: dbTestError,
		},
		{
			name:      "Количество товара в корзине слишком мало",
			productID: 1,
			count:     11,
			userID:    1,
			setupFunc: func(ctrl *gomock.Controller) *CartManager {
				cartRepositoryMock := mocks.NewMockcartRepository(ctrl)
				productCountGetterMock := mocks.NewMockproductCountGetter(ctrl)
				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				cartProduct := model.CartProduct{
					BaseProduct: model.BaseProduct{
						Count: 10,
					},
					IsDeleted: false,
				}

				cartRepositoryMock.EXPECT().GetCartProduct(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(cartProduct, nil)

				productCountGetterMock.EXPECT().GetProductCount(
					gomock.Any(),
					uint32(1)).Return(uint32(10), nil)

				return &CartManager{
					cartRepository:     cartRepositoryMock,
					productCountGetter: productCountGetterMock,
					log:                logger,
				}
			},
			want: errs.ProductCountTooLow,
		},
		{
			name:      "Ошибка удаления продукта из корзины",
			productID: 1,
			count:     -1,
			userID:    1,
			setupFunc: func(ctrl *gomock.Controller) *CartManager {
				cartRepositoryMock := mocks.NewMockcartRepository(ctrl)
				productCountGetterMock := mocks.NewMockproductCountGetter(ctrl)
				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				cartProduct := model.CartProduct{
					BaseProduct: model.BaseProduct{
						Count: 1,
					},
					IsDeleted: false,
				}

				cartRepositoryMock.EXPECT().GetCartProduct(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(cartProduct, nil)

				productCountGetterMock.EXPECT().GetProductCount(
					gomock.Any(),
					uint32(1)).Return(uint32(10), nil)

				cartRepositoryMock.EXPECT().RemoveCartProduct(
					gomock.Any(),
					uint32(1),
					int32(-1),
					uint32(1)).Return(dbTestError)

				return &CartManager{
					cartRepository:     cartRepositoryMock,
					productCountGetter: productCountGetterMock,
					log:                logger,
				}
			},
			want: dbTestError,
		},
		{
			name:      "Ошибка изменения количества продукта",
			productID: 1,
			count:     -1,
			userID:    1,
			setupFunc: func(ctrl *gomock.Controller) *CartManager {
				cartRepositoryMock := mocks.NewMockcartRepository(ctrl)
				productCountGetterMock := mocks.NewMockproductCountGetter(ctrl)
				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				cartProduct := model.CartProduct{
					BaseProduct: model.BaseProduct{
						Count: 10,
					},
					IsDeleted: false,
				}

				cartRepositoryMock.EXPECT().GetCartProduct(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(cartProduct, nil)

				productCountGetterMock.EXPECT().GetProductCount(
					gomock.Any(),
					uint32(1)).Return(uint32(100), nil)

				cartRepositoryMock.EXPECT().ChangeCartProductCount(
					gomock.Any(),
					uint32(1),
					int32(-1),
					uint32(1)).Return(dbTestError)

				return &CartManager{
					cartRepository:     cartRepositoryMock,
					productCountGetter: productCountGetterMock,
					log:                logger,
				}
			},
			want: dbTestError,
		},
		{
			name:      "Неправильное изменение количества товара в корзине",
			productID: 1,
			count:     0,
			userID:    1,
			setupFunc: func(ctrl *gomock.Controller) *CartManager {
				cartRepositoryMock := mocks.NewMockcartRepository(ctrl)
				productCountGetterMock := mocks.NewMockproductCountGetter(ctrl)
				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				cartProduct := model.CartProduct{
					BaseProduct: model.BaseProduct{
						Count: 10,
					},
					IsDeleted: false,
				}

				cartRepositoryMock.EXPECT().GetCartProduct(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(cartProduct, nil)

				productCountGetterMock.EXPECT().GetProductCount(
					gomock.Any(),
					uint32(1)).Return(uint32(100), nil)

				return &CartManager{
					cartRepository:     cartRepositoryMock,
					productCountGetter: productCountGetterMock,
					log:                logger,
				}
			},
			want: errs.BadRequest,
		},
		{
			name:      "Успешное удаление товара из корзины",
			productID: 1,
			count:     -1,
			userID:    1,
			setupFunc: func(ctrl *gomock.Controller) *CartManager {
				cartRepositoryMock := mocks.NewMockcartRepository(ctrl)
				productCountGetterMock := mocks.NewMockproductCountGetter(ctrl)

				cartProduct := model.CartProduct{
					BaseProduct: model.BaseProduct{
						Count: 1,
					},
					IsDeleted: false,
				}

				cartRepositoryMock.EXPECT().GetCartProduct(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(cartProduct, nil)

				productCountGetterMock.EXPECT().GetProductCount(
					gomock.Any(),
					uint32(1)).Return(uint32(10), nil)

				cartRepositoryMock.EXPECT().RemoveCartProduct(
					gomock.Any(),
					uint32(1),
					int32(-1),
					uint32(1)).Return(nil)

				return &CartManager{
					cartRepository:     cartRepositoryMock,
					productCountGetter: productCountGetterMock,
				}
			},
			want: nil,
		},
		{
			name:      "Успешное изменение количества товара в корзине",
			productID: 1,
			count:     -1,
			userID:    1,
			setupFunc: func(ctrl *gomock.Controller) *CartManager {
				cartRepositoryMock := mocks.NewMockcartRepository(ctrl)
				productCountGetterMock := mocks.NewMockproductCountGetter(ctrl)

				cartProduct := model.CartProduct{
					BaseProduct: model.BaseProduct{
						Count: 10,
					},
					IsDeleted: false,
				}

				cartRepositoryMock.EXPECT().GetCartProduct(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(cartProduct, nil)

				productCountGetterMock.EXPECT().GetProductCount(
					gomock.Any(),
					uint32(1)).Return(uint32(10), nil)

				cartRepositoryMock.EXPECT().ChangeCartProductCount(
					gomock.Any(),
					uint32(1),
					int32(-1),
					uint32(1)).Return(nil)

				return &CartManager{
					cartRepository:     cartRepositoryMock,
					productCountGetter: productCountGetterMock,
				}
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			resp := tt.setupFunc(ctrl).ChangeCartProductCount(nil, tt.productID, tt.count, tt.userID)

			assert.Equal(t, tt.want, resp)
		})
	}
}
