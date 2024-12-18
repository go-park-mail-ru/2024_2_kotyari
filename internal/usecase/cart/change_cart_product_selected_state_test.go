package cart

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/cart/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"log/slog"
	"testing"
)

func TestCartManager_ChangeCartProductSelectedState(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		productID     uint32
		userID        uint32
		isSelected    bool
		setupFunc     func(ctrl *gomock.Controller) *CartManager
		expectedError error
	}{
		{
			name:       "Ошибка получения продукта",
			productID:  1,
			userID:     1,
			isSelected: true,
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
			expectedError: dbTestError,
		},
		{
			name:       "Продукт не найден в корзине",
			productID:  1,
			userID:     1,
			isSelected: true,
			setupFunc: func(ctrl *gomock.Controller) *CartManager {
				cartRepositoryMock := mocks.NewMockcartRepository(ctrl)
				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				cartRepositoryMock.EXPECT().GetCartProduct(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(model.CartProduct{IsDeleted: true}, nil)

				return &CartManager{
					cartRepository: cartRepositoryMock,
					log:            logger,
				}
			},
			expectedError: errs.ProductNotInCart,
		},
		{
			name:       "Ошибка изменения состояния is_selected продукта",
			productID:  1,
			userID:     1,
			isSelected: true,
			setupFunc: func(ctrl *gomock.Controller) *CartManager {
				cartRepositoryMock := mocks.NewMockcartRepository(ctrl)
				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				cartRepositoryMock.EXPECT().GetCartProduct(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(model.CartProduct{IsDeleted: false}, nil)

				cartRepositoryMock.EXPECT().ChangeCartProductSelectedState(
					gomock.Any(),
					uint32(1),
					uint32(1),
					true).Return(dbTestError)

				return &CartManager{
					cartRepository: cartRepositoryMock,
					log:            logger,
				}
			},
			expectedError: dbTestError,
		},
		{
			name:       "Успешное изменение состояния is_selected продукта",
			productID:  1,
			userID:     1,
			isSelected: true,
			setupFunc: func(ctrl *gomock.Controller) *CartManager {
				cartRepositoryMock := mocks.NewMockcartRepository(ctrl)
				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				cartRepositoryMock.EXPECT().GetCartProduct(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(model.CartProduct{IsDeleted: false}, nil)

				cartRepositoryMock.EXPECT().ChangeCartProductSelectedState(
					gomock.Any(),
					uint32(1),
					uint32(1),
					true).Return(nil)

				return &CartManager{
					cartRepository: cartRepositoryMock,
					log:            logger,
				}
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.WithValue(context.Background(), testContextRequestIDKey, testContextRequestIDValue)

			err := tt.setupFunc(ctrl).ChangeCartProductSelectedState(ctx, tt.productID, tt.userID, tt.isSelected)

			assert.Equal(t, tt.expectedError, err)
		})
	}
}
