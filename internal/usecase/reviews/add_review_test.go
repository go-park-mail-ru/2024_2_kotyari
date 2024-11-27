package reviews

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"testing"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/reviews/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const testContextRequestIDKey = "request-id"

var (
	testContextRequestIDValue    = uuid.New()
	testDBError                  = errors.New("тестовая ошибка базы")
	testProductNotFoundGRPCError = status.Error(codes.NotFound, errs.ProductNotFound.Error())
	testGRPCInternalServerError  = status.Error(codes.Internal, errs.InternalServerError.Error())
	testGRPCUnavailableError     = status.Error(codes.Unavailable, errs.InternalServerError.Error())
)

func TestReviewsService_AddReview(t *testing.T) {
	t.Parallel()

	type want error

	tests := []struct {
		name      string
		productID uint32
		userID    uint32
		review    model.Review
		setupFunc func(ctrl *gomock.Controller) *ReviewsService
		want      want
	}{
		{
			name:      "Отзыв успешно добавлен",
			productID: 1,
			userID:    1,
			review: model.Review{
				Text:      "Классный продукт",
				Rating:    5,
				IsPrivate: false,
			},
			setupFunc: func(ctrl *gomock.Controller) *ReviewsService {
				reviewsRepositoryMock := mocks.NewMockreviewsRepo(ctrl)
				ratingUpdaterMock := mocks.NewMockratingUpdater(ctrl)

				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				reviewsRepositoryMock.EXPECT().GetReview(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(model.Review{}, errs.ReviewNotFound)

				review := model.Review{
					Text:      "Классный продукт",
					Rating:    5,
					IsPrivate: false,
				}

				reviewsRepositoryMock.EXPECT().AddReview(
					gomock.Any(),
					uint32(1),
					uint32(1),
					review).Return(nil)

				ratingUpdaterMock.EXPECT().UpdateRating(
					gomock.Any(),
					uint32(1)).Return(nil)

				return &ReviewsService{
					reviewsRepo:     reviewsRepositoryMock,
					ratingUpdater:   ratingUpdaterMock,
					stringSanitizer: nil,
					log:             logger,
				}
			},
			want: nil,
		},
		{
			name:      "Отзыв уже существует",
			productID: 1,
			userID:    1,
			review: model.Review{
				Text:      "Классный продукт",
				Rating:    5,
				IsPrivate: false,
			},
			setupFunc: func(ctrl *gomock.Controller) *ReviewsService {
				reviewsRepositoryMock := mocks.NewMockreviewsRepo(ctrl)
				ratingUpdaterMock := mocks.NewMockratingUpdater(ctrl)

				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				reviewsRepositoryMock.EXPECT().GetReview(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(model.Review{
					Text:      "Классный продукт",
					Rating:    5,
					IsPrivate: false,
				}, nil)

				return &ReviewsService{
					reviewsRepo:     reviewsRepositoryMock,
					ratingUpdater:   ratingUpdaterMock,
					stringSanitizer: nil,
					log:             logger,
				}
			},
			want: errs.ReviewAlreadyExists,
		},
		{
			name:      "Ошибка при получении существующего отзыва",
			productID: 1,
			userID:    1,
			review: model.Review{
				Text:      "Классный продукт",
				Rating:    5,
				IsPrivate: false,
			},
			setupFunc: func(ctrl *gomock.Controller) *ReviewsService {
				reviewsRepositoryMock := mocks.NewMockreviewsRepo(ctrl)
				ratingUpdaterMock := mocks.NewMockratingUpdater(ctrl)

				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				reviewsRepositoryMock.EXPECT().GetReview(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(model.Review{}, testDBError)

				return &ReviewsService{
					reviewsRepo:     reviewsRepositoryMock,
					ratingUpdater:   ratingUpdaterMock,
					stringSanitizer: nil,
					log:             logger,
				}
			},
			want: testDBError,
		},
		{
			name:      "Ошибка при добавлении отзыва",
			productID: 1,
			userID:    1,
			review: model.Review{
				Text:      "Классный продукт",
				Rating:    5,
				IsPrivate: false,
			},
			setupFunc: func(ctrl *gomock.Controller) *ReviewsService {
				reviewsRepositoryMock := mocks.NewMockreviewsRepo(ctrl)
				ratingUpdaterMock := mocks.NewMockratingUpdater(ctrl)

				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				reviewsRepositoryMock.EXPECT().GetReview(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(model.Review{}, errs.ReviewNotFound)

				review := model.Review{
					Text:      "Классный продукт",
					Rating:    5,
					IsPrivate: false,
				}

				reviewsRepositoryMock.EXPECT().AddReview(
					gomock.Any(),
					uint32(1),
					uint32(1),
					review).Return(testDBError)

				return &ReviewsService{
					reviewsRepo:     reviewsRepositoryMock,
					ratingUpdater:   ratingUpdaterMock,
					stringSanitizer: nil,
					log:             logger,
				}
			},
			want: testDBError,
		},
		{
			name:      "Ошибка при обновлении рейтинга продукта - продукт не найден",
			productID: 1,
			userID:    1,
			review: model.Review{
				Text:      "Классный продукт",
				Rating:    5,
				IsPrivate: false,
			},
			setupFunc: func(ctrl *gomock.Controller) *ReviewsService {
				reviewsRepositoryMock := mocks.NewMockreviewsRepo(ctrl)
				ratingUpdaterMock := mocks.NewMockratingUpdater(ctrl)

				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				reviewsRepositoryMock.EXPECT().GetReview(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(model.Review{}, errs.ReviewNotFound)

				review := model.Review{
					Text:      "Классный продукт",
					Rating:    5,
					IsPrivate: false,
				}

				reviewsRepositoryMock.EXPECT().AddReview(
					gomock.Any(),
					uint32(1),
					uint32(1),
					review).Return(nil)

				ratingUpdaterMock.EXPECT().UpdateRating(
					gomock.Any(),
					uint32(1)).Return(testProductNotFoundGRPCError)

				return &ReviewsService{
					reviewsRepo:     reviewsRepositoryMock,
					ratingUpdater:   ratingUpdaterMock,
					stringSanitizer: nil,
					log:             logger,
				}
			},
			want: errs.FailedToChangeProductRating,
		},
		{
			name:      "Сервис обновлении рейтинга продукта недоступен",
			productID: 1,
			userID:    1,
			review: model.Review{
				Text:      "Классный продукт",
				Rating:    5,
				IsPrivate: false,
			},
			setupFunc: func(ctrl *gomock.Controller) *ReviewsService {
				reviewsRepositoryMock := mocks.NewMockreviewsRepo(ctrl)
				ratingUpdaterMock := mocks.NewMockratingUpdater(ctrl)

				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				reviewsRepositoryMock.EXPECT().GetReview(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(model.Review{}, errs.ReviewNotFound)

				review := model.Review{
					Text:      "Классный продукт",
					Rating:    5,
					IsPrivate: false,
				}

				reviewsRepositoryMock.EXPECT().AddReview(
					gomock.Any(),
					uint32(1),
					uint32(1),
					review).Return(nil)

				ratingUpdaterMock.EXPECT().UpdateRating(
					gomock.Any(),
					uint32(1)).Return(testGRPCUnavailableError)

				return &ReviewsService{
					reviewsRepo:     reviewsRepositoryMock,
					ratingUpdater:   ratingUpdaterMock,
					stringSanitizer: nil,
					log:             logger,
				}
			},
			want: nil,
		},
		{
			name:      "Ошибка при обновлении рейтинга продукта",
			productID: 1,
			userID:    1,
			review: model.Review{
				Text:      "Классный продукт",
				Rating:    5,
				IsPrivate: false,
			},
			setupFunc: func(ctrl *gomock.Controller) *ReviewsService {
				reviewsRepositoryMock := mocks.NewMockreviewsRepo(ctrl)
				ratingUpdaterMock := mocks.NewMockratingUpdater(ctrl)

				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				reviewsRepositoryMock.EXPECT().GetReview(
					gomock.Any(),
					uint32(1),
					uint32(1)).Return(model.Review{}, errs.ReviewNotFound)

				review := model.Review{
					Text:      "Классный продукт",
					Rating:    5,
					IsPrivate: false,
				}

				reviewsRepositoryMock.EXPECT().AddReview(
					gomock.Any(),
					uint32(1),
					uint32(1),
					review).Return(nil)

				ratingUpdaterMock.EXPECT().UpdateRating(
					gomock.Any(),
					uint32(1)).Return(testGRPCInternalServerError)

				return &ReviewsService{
					reviewsRepo:     reviewsRepositoryMock,
					ratingUpdater:   ratingUpdaterMock,
					stringSanitizer: nil,
					log:             logger,
				}
			},
			want: nil,
		},
		{
			name:      "Неправильная оценка в отзыве",
			productID: 1,
			userID:    1,
			review: model.Review{
				Text:      "Классный продукт",
				Rating:    10,
				IsPrivate: false,
			},
			setupFunc: func(ctrl *gomock.Controller) *ReviewsService {
				reviewsRepositoryMock := mocks.NewMockreviewsRepo(ctrl)
				ratingUpdaterMock := mocks.NewMockratingUpdater(ctrl)

				logger := slog.New(slog.NewTextHandler(io.Discard, nil))

				return &ReviewsService{
					reviewsRepo:     reviewsRepositoryMock,
					ratingUpdater:   ratingUpdaterMock,
					stringSanitizer: nil,
					log:             logger,
				}
			},
			want: errs.BadRequest,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.WithValue(context.Background(), testContextRequestIDKey, testContextRequestIDValue)

			resp := tt.setupFunc(ctrl).AddReview(ctx, tt.productID, tt.userID, tt.review)
			assert.Equal(t, tt.want, resp)
		})
	}
}
