package cart

import (
	"context"
	"io"
	"log/slog"
	"testing"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/cart/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCartManager_RemoveProduct(t *testing.T) {
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
			name:      "Успешное удаление из корзины",
			productID: 1,
			userID:    1,
			setupFunc: func(ctrl *gomock.Controller) *CartManager {
				cartRepositoryMock := mocks.NewMockcartRepository(ctrl)
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

				cartRepositoryMock.EXPECT().RemoveCartProduct(
					gomock.Any(),
					uint32(1),
					-int32(cartProduct.Count),
					uint32(1)).Return(nil)

				return &CartManager{
					cartRepository: cartRepositoryMock,
					log:            logger,
				}
			},
			want: nil,
		},
		{
			name:      "Не удалось получить продукт",
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
			name:      "Ошибка удаления продукта",
			productID: 1,
			userID:    1,
			setupFunc: func(ctrl *gomock.Controller) *CartManager {
				cartRepositoryMock := mocks.NewMockcartRepository(ctrl)
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

				cartRepositoryMock.EXPECT().RemoveCartProduct(
					gomock.Any(),
					uint32(1),
					-int32(cartProduct.Count),
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

			ctx := context.WithValue(context.Background(), testContextRequestIDKey, testContextRequestIDValue)

			resp := tt.setupFunc(ctrl).RemoveProduct(ctx, tt.productID, tt.userID)
			assert.Equal(t, tt.want, resp)
		})
	}
}
