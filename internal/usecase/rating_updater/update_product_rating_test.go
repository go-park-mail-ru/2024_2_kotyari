package rating_updater

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/rating_updater/mocks"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"io"
	"log/slog"
	"testing"
)

var (
	testContextRequestIDValue = uuid.New()
	dbTestError               = errors.New("ошибка базы данных")
)

const testContextRequestIDKey = "request-id"

func TestRatingUpdaterService_UpdateProductRating(t *testing.T) {
	type want struct {
		err error
	}

	var tests = []struct {
		name      string
		productID uint32
		setupFunc func(ctrl *gomock.Controller) *RatingUpdaterService
		want      want
	}{
		{
			name:      "Успешное обновление рейтинга",
			productID: 1,
			setupFunc: func(ctrl *gomock.Controller) *RatingUpdaterService {
				repository := mocks.NewMockRatingUpdaterRepository(ctrl)
				reviewsGetter := mocks.NewMockReviewsGetter(ctrl)

				reviewsGetter.EXPECT().GetProductReviewsNoLogin(gomock.Any(), uint32(1), utils.DefaultFieldParam, utils.DefaultOrderParam).Return(model.Reviews{
					Reviews: []model.Review{
						{Rating: 5},
						{Rating: 4},
						{Rating: 3},
					},
				}, nil)

				repository.EXPECT().UpdateProductRating(gomock.Any(), uint32(1), float32(4)).Return(nil)

				return &RatingUpdaterService{
					repository:    repository,
					reviewsGetter: reviewsGetter,
					log:           slog.New(slog.NewTextHandler(io.Discard, nil)),
				}
			},
			want: want{
				err: nil,
			},
		},
		{
			name:      "Ошибка при обновлении рейтинга",
			productID: 1,
			setupFunc: func(ctrl *gomock.Controller) *RatingUpdaterService {
				repository := mocks.NewMockRatingUpdaterRepository(ctrl)
				reviewsGetter := mocks.NewMockReviewsGetter(ctrl)

				reviewsGetter.EXPECT().GetProductReviewsNoLogin(gomock.Any(), uint32(1), utils.DefaultFieldParam, utils.DefaultOrderParam).Return(model.Reviews{
					Reviews: []model.Review{
						{Rating: 5},
						{Rating: 4},
						{Rating: 3},
					},
				}, nil)

				repository.EXPECT().UpdateProductRating(gomock.Any(), uint32(1), float32(4)).Return(errors.New("some error"))

				return &RatingUpdaterService{
					repository:    repository,
					reviewsGetter: reviewsGetter,
					log:           slog.New(slog.NewTextHandler(io.Discard, nil)),
				}
			},
			want: want{
				err: errors.New("some error"),
			},
		},
		{
			name:      "Нет отзывов для продукта",
			productID: 1,
			setupFunc: func(ctrl *gomock.Controller) *RatingUpdaterService {
				repository := mocks.NewMockRatingUpdaterRepository(ctrl)
				reviewsGetter := mocks.NewMockReviewsGetter(ctrl)

				reviewsGetter.EXPECT().GetProductReviewsNoLogin(gomock.Any(), uint32(1), utils.DefaultFieldParam, utils.DefaultOrderParam).Return(model.Reviews{}, errs.NoReviewsForProduct)

				repository.EXPECT().UpdateProductRating(gomock.Any(), uint32(1), float32(0)).Return(nil)

				return &RatingUpdaterService{
					repository:    repository,
					reviewsGetter: reviewsGetter,
					log:           slog.New(slog.NewTextHandler(io.Discard, nil)),
				}
			},
			want: want{
				err: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.WithValue(context.Background(), testContextRequestIDKey, testContextRequestIDValue)

			service := tt.setupFunc(ctrl)

			err := service.UpdateProductRating(ctx, tt.productID)

			assert.Equal(t, tt.want.err, err)
		})
	}
}
