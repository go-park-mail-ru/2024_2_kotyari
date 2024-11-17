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

func TestReviewsService_DeleteReview(t *testing.T) {
	t.Parallel()

	type want error

	tests := []struct {
		name      string
		productID uint32
		userID    uint32
		setupFunc func(ctrl *gomock.Controller) *ReviewsService
		want      want
	}{
		{
			name:      "Отзыв успешно удален",
			productID: 1,
			userID:    1,
			setupFunc: func(ctrl *gomock.Controller) *ReviewsService {
				reviewsRepositoryMock := mocks.NewMockreviewsRepo(ctrl)
				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				reviewsRepositoryMock.EXPECT().GetReview(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(model.Review{
					Text:      "Классный продукт",
					Rating:    5,
					IsPrivate: false,
				}, nil)

				reviewsRepositoryMock.EXPECT().DeleteReview(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(nil)

				return &ReviewsService{
					reviewsRepo:    reviewsRepositoryMock,
					inputValidator: nil,
					log:            logger,
				}
			},
			want: nil,
		},
		{
			name:      "Отзыв не найден",
			productID: 1,
			userID:    1,
			setupFunc: func(ctrl *gomock.Controller) *ReviewsService {
				reviewsRepositoryMock := mocks.NewMockreviewsRepo(ctrl)
				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				reviewsRepositoryMock.EXPECT().GetReview(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(model.Review{}, errs.ReviewNotFound)

				return &ReviewsService{
					reviewsRepo:    reviewsRepositoryMock,
					inputValidator: nil,
					log:            logger,
				}
			},
			want: errs.ReviewNotFound,
		},
		{
			name:      "Ошибка при получения отзыва",
			productID: 1,
			userID:    1,
			setupFunc: func(ctrl *gomock.Controller) *ReviewsService {
				reviewsRepositoryMock := mocks.NewMockreviewsRepo(ctrl)
				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				reviewsRepositoryMock.EXPECT().GetReview(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(model.Review{}, testDBError)

				return &ReviewsService{
					reviewsRepo:    reviewsRepositoryMock,
					inputValidator: nil,
					log:            logger,
				}
			},
			want: testDBError,
		},
		{
			name:      "Ошибка при удалении отзыва",
			productID: 1,
			userID:    1,
			setupFunc: func(ctrl *gomock.Controller) *ReviewsService {
				reviewsRepositoryMock := mocks.NewMockreviewsRepo(ctrl)
				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				reviewsRepositoryMock.EXPECT().GetReview(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(model.Review{
					Text:      "Классный продукт",
					Rating:    5,
					IsPrivate: false,
				}, nil)

				reviewsRepositoryMock.EXPECT().DeleteReview(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(testDBError)

				return &ReviewsService{
					reviewsRepo:    reviewsRepositoryMock,
					inputValidator: nil,
					log:            logger,
				}
			},
			want: testDBError,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.WithValue(context.Background(), testContextRequestIDKey, testContextRequestIDValue)

			resp := tt.setupFunc(ctrl).DeleteReview(ctx, tt.productID, tt.userID)
			assert.Equal(t, tt.want, resp)
		})
	}
}
