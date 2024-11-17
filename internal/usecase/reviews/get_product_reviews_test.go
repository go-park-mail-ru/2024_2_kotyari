package reviews

import (
	"context"
	"io"
	"log/slog"
	"testing"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/reviews/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestReviewsService_GetProductReviews(t *testing.T) {
	t.Parallel()

	type want struct {
		reviews model.Reviews
		err     error
	}

	tests := []struct {
		name      string
		productID uint32
		sortField string
		sortOrder string
		setupFunc func(ctrl *gomock.Controller) *ReviewsService
		want      want
	}{
		{
			name:      "Отзывы успешно получены",
			productID: 1,
			sortField: "",
			sortOrder: "",
			setupFunc: func(ctrl *gomock.Controller) *ReviewsService {
				reviewsRepositoryMock := mocks.NewMockreviewsRepo(ctrl)
				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				reviewsRepositoryMock.EXPECT().GetProductReviews(
					gomock.Any(),
					uint32(1),
					"",
					"").Return(model.Reviews{
					TotalReviewCount: 0,
					TotalRating:      0,
					Reviews: []model.Review{
						{
							Rating: 2,
						},
						{
							Rating: 4,
						},
					},
				}, nil)

				return &ReviewsService{
					reviewsRepo:    reviewsRepositoryMock,
					inputValidator: nil,
					log:            logger,
				}
			},
			want: want{
				reviews: model.Reviews{
					TotalRating:      float32(3.0),
					TotalReviewCount: 2,
					Reviews: []model.Review{
						{
							Rating: 2,
						},
						{
							Rating: 4,
						},
					},
				},
				err: nil,
			},
		},
		{
			name:      "Отзывы для продукта не найдены",
			productID: 1,
			sortField: "",
			sortOrder: "",
			setupFunc: func(ctrl *gomock.Controller) *ReviewsService {
				reviewsRepositoryMock := mocks.NewMockreviewsRepo(ctrl)
				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				reviewsRepositoryMock.EXPECT().GetProductReviews(
					gomock.Any(),
					uint32(1),
					"",
					"",
				).Return(model.Reviews{}, errs.NoReviewsForProduct)

				return &ReviewsService{
					reviewsRepo:    reviewsRepositoryMock,
					inputValidator: nil,
					log:            logger,
				}
			},
			want: want{
				reviews: model.Reviews{},
				err:     errs.NoReviewsForProduct,
			},
		},
		{
			name:      "Ошибка при получении отзывов для продукта",
			productID: 1,
			sortField: "",
			sortOrder: "",
			setupFunc: func(ctrl *gomock.Controller) *ReviewsService {
				reviewsRepositoryMock := mocks.NewMockreviewsRepo(ctrl)
				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				reviewsRepositoryMock.EXPECT().GetProductReviews(
					gomock.Any(),
					uint32(1),
					"",
					"",
				).Return(model.Reviews{}, testDBError)

				return &ReviewsService{
					reviewsRepo:    reviewsRepositoryMock,
					inputValidator: nil,
					log:            logger,
				}
			},
			want: want{
				reviews: model.Reviews{},
				err:     testDBError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.WithValue(context.Background(), testContextRequestIDKey, testContextRequestIDValue)

			resp, err := tt.setupFunc(ctrl).GetProductReviews(ctx, tt.productID, tt.sortField, tt.sortOrder)
			assert.Equal(t, tt.want.reviews, resp)
			assert.Equal(t, tt.want.err, err)
		})
	}
}
