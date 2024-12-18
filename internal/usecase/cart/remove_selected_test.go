package cart

import (
	"context"
	"go.uber.org/mock/gomock"
	"io"
	"log/slog"
	"testing"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/cart/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCartManager_RemoveSelected(t *testing.T) {
	t.Parallel()

	type want error
	tests := []struct {
		name      string
		userID    uint32
		setupFunc func(ctrl *gomock.Controller) *CartManager
		want      want
	}{
		{
			name:   "Успешное удаление is_selected у продуктов",
			userID: 1,
			setupFunc: func(ctrl *gomock.Controller) *CartManager {
				cartRepositoryMock := mocks.NewMockcartRepository(ctrl)
				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				selectedProducts := []model.ProductOrder{
					{
						ID:    2,
						Count: 10,
					},
					{
						ID:    3,
						Count: 2,
					},
				}

				cartRepositoryMock.EXPECT().GetSelectedCartItems(
					gomock.Any(),
					uint32(1)).Return(selectedProducts, nil)

				for _, product := range selectedProducts {
					cartRepositoryMock.EXPECT().RemoveCartProduct(
						gomock.Any(),
						product.ID,
						int32(product.Count),
						uint32(1)).Return(nil)
				}

				return &CartManager{
					cartRepository: cartRepositoryMock,
					log:            logger,
				}
			},
			want: nil,
		},
		{
			name:   "Нет выбранных продуктов",
			userID: 1,
			setupFunc: func(ctrl *gomock.Controller) *CartManager {
				cartRepositoryMock := mocks.NewMockcartRepository(ctrl)
				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				cartRepositoryMock.EXPECT().GetSelectedCartItems(
					gomock.Any(),
					uint32(1)).Return([]model.ProductOrder{}, nil)

				return &CartManager{
					cartRepository: cartRepositoryMock,
					log:            logger,
				}
			},
			want: errs.NoSelectedProducts,
		},
		{
			name:   "Ошибка получения выбранных продуктов",
			userID: 1,
			setupFunc: func(ctrl *gomock.Controller) *CartManager {
				cartRepositoryMock := mocks.NewMockcartRepository(ctrl)
				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				cartRepositoryMock.EXPECT().GetSelectedCartItems(
					gomock.Any(),
					uint32(1)).Return([]model.ProductOrder{}, dbTestError)
				return &CartManager{
					cartRepository: cartRepositoryMock,
					log:            logger,
				}
			},
			want: dbTestError,
		},
		{
			name:   "Ошибка удаления is_selected продукта",
			userID: 1,
			setupFunc: func(ctrl *gomock.Controller) *CartManager {
				cartRepositoryMock := mocks.NewMockcartRepository(ctrl)
				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				selectedProducts := []model.ProductOrder{
					{
						ID:    2,
						Count: 10,
					},
					{
						ID:    3,
						Count: 2,
					},
				}

				cartRepositoryMock.EXPECT().GetSelectedCartItems(
					gomock.Any(),
					uint32(1)).Return(selectedProducts, nil)

				cartRepositoryMock.EXPECT().RemoveCartProduct(
					gomock.Any(),
					selectedProducts[0].ID,
					int32(selectedProducts[0].Count),
					uint32(1)).Return(nil)

				cartRepositoryMock.EXPECT().RemoveCartProduct(
					gomock.Any(),
					selectedProducts[1].ID,
					int32(selectedProducts[1].Count),
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

			resp := tt.setupFunc(ctrl).RemoveSelected(ctx, tt.userID)
			assert.Equal(t, tt.want, resp)
		})
	}
}
