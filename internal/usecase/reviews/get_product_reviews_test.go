package reviews

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"io"
	"log/slog"
	"testing"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/reviews/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestReviewsService_GetProductReviewsNoLogin(t *testing.T) {
	t.Parallel()

	type args struct {
		productID uint32
		sortField string
		sortOrder string
	}

	type want struct {
		reviews model.Reviews
		err     error
	}

	tests := []struct {
		name      string
		args      args
		setupFunc func(ctrl *gomock.Controller) *ReviewsService
		want      want
	}{
		{
			name: "Успешное получение отзывов без авторизации",
			args: args{
				productID: 1,
				sortField: "date",
				sortOrder: "desc",
			},
			setupFunc: func(ctrl *gomock.Controller) *ReviewsService {
				reviewsRepositoryMock := mocks.NewMockreviewsRepo(ctrl)
				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				reviewsFromRepo := model.Reviews{
					Reviews: []model.Review{
						{Text: "Отличный продукт", Rating: 5, IsPrivate: false},
						{Text: "Хороший продукт", Rating: 4, IsPrivate: false},
					},
					TotalReviewCount: 2,
				}

				reviewsRepositoryMock.EXPECT().GetProductReviews(
					gomock.Any(),
					uint32(1),
					"date",
					"desc",
				).Return(reviewsFromRepo, nil)

				return &ReviewsService{
					reviewsRepo:     reviewsRepositoryMock,
					stringSanitizer: nil,
					log:             logger,
				}
			},
			want: want{
				reviews: model.Reviews{
					Reviews: []model.Review{
						{Text: "Отличный продукт", Rating: 5, IsPrivate: false},
						{Text: "Хороший продукт", Rating: 4, IsPrivate: false},
					},
					TotalRating:      4.5,
					TotalReviewCount: 2,
				},
				err: nil,
			},
		},
		{
			name: "Ошибка получения отзывов из репозитория",
			args: args{
				productID: 1,
				sortField: "date",
				sortOrder: "desc",
			},
			setupFunc: func(ctrl *gomock.Controller) *ReviewsService {
				reviewsRepositoryMock := mocks.NewMockreviewsRepo(ctrl)
				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				reviewsRepositoryMock.EXPECT().GetProductReviews(
					gomock.Any(),
					uint32(1),
					"date",
					"desc",
				).Return(model.Reviews{}, testDBError)

				return &ReviewsService{
					reviewsRepo:     reviewsRepositoryMock,
					stringSanitizer: nil,
					log:             logger,
				}
			},
			want: want{
				reviews: model.Reviews{},
				err:     testDBError,
			},
		},
		{
			name: "Получено нулевое количество отзывов",
			args: args{
				productID: 1,
				sortField: "date",
				sortOrder: "desc",
			},
			setupFunc: func(ctrl *gomock.Controller) *ReviewsService {
				reviewsRepositoryMock := mocks.NewMockreviewsRepo(ctrl)
				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				reviewsRepositoryMock.EXPECT().GetProductReviews(
					gomock.Any(),
					uint32(1),
					"date",
					"desc",
				).Return(model.Reviews{}, errs.NoReviewsForProduct)

				return &ReviewsService{
					reviewsRepo:     reviewsRepositoryMock,
					stringSanitizer: nil,
					log:             logger,
				}
			},
			want: want{
				reviews: model.Reviews{},
				err:     errs.NoReviewsForProduct,
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

			gotReviews, err := tt.setupFunc(ctrl).GetProductReviews(ctx, tt.args.productID, tt.args.sortField, tt.args.sortOrder)
			if tt.want.err != nil {
				assert.EqualError(t, err, tt.want.err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want.reviews.TotalRating, gotReviews.TotalRating)
				assert.Equal(t, tt.want.reviews.TotalReviewCount, gotReviews.TotalReviewCount)
				assert.Equal(t, tt.want.reviews.Reviews, gotReviews.Reviews)
			}
		})
	}
}
