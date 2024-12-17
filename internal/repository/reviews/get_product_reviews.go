package reviews

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/jackc/pgx/v5"
)

func (r *ReviewsStore) GetProductReviewsNoLogin(ctx context.Context, productID uint32, sortField string, sortOrder string) (model.Reviews, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return model.Reviews{}, err
	}

	r.log.Info("[ReviewsStore.GetProductReviewsNoLogin] Started executing", slog.Any("request-id", requestID))

	var reviews model.Reviews
	const countReviews = `
		select count(id)
		from reviews
		where product_id = $1;
	`

	err = r.db.QueryRow(ctx, countReviews, productID).Scan(&reviews.TotalReviewCount)
	if err != nil {
		r.log.Error("[ReviewsStore.GetProductReviewsNoLogin] Error getting total review count", slog.String("error", err.Error()))

		return model.Reviews{}, err
	}

	fieldSortOptions := map[string]string{
		"rating": "r.rating",
		"date":   "r.created_at",
	}

	field, ok := fieldSortOptions[sortField]
	if !ok {
		field = "r.created_at"
	}

	sortOrder = utils.ReturnSortOrderOption(sortOrder)

	query := fmt.Sprintf(`
		select r.text, r.rating, r.is_private, u.username, u.avatar_url, r.created_at
		from reviews r
		join users u on u.id = r.user_id
		where r.product_id = $1
		order by %s %s;
	`, field, sortOrder)

	rows, err := r.db.Query(ctx, query, productID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.log.Error("[ReviewsStore.GetProductReviewsNoLogin] No reviews")

			return model.Reviews{}, errs.NoReviewsForProduct
		}

		r.log.Info("[ReviewsStore.GetProductReviewsNoLogin] Unexpected error happened", slog.String("error", err.Error()))

		return model.Reviews{}, err
	}

	reviewsDTO, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[ReviewDTO])
	if err != nil {
		r.log.Error("[ReviewsStore.GetProductReviewsNoLogin] Error happened when scanning rows", slog.String("error", err.Error()))

		return model.Reviews{}, err
	}

	if len(reviewsDTO) == 0 {
		return model.Reviews{}, errs.NoReviewsForProduct
	}

	reviews.Reviews = ToReviewModelSlice(reviewsDTO)

	return reviews, nil
}

func (r *ReviewsStore) GetProductReviewsWithLogin(ctx context.Context, productID uint32, userID uint32, sortField string, sortOrder string) (model.Reviews, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return model.Reviews{}, err
	}

	r.log.Info("[ReviewsStore.GetProductReviewsWithLogin] Started executing", slog.Any("request-id", requestID))

	var reviews model.Reviews
	const countReviews = `
		select count(id)
		from reviews
		where product_id = $1;
	`

	err = r.db.QueryRow(ctx, countReviews, productID).Scan(&reviews.TotalReviewCount)
	if err != nil {
		r.log.Error("[ReviewsStore.GetProductReviewsWithLogin] Error getting total review count", slog.String("error", err.Error()))

		return model.Reviews{}, err
	}

	const userReviewQuery = `
		select r.text, r.rating, r.is_private, r.created_at
        from reviews r
        join users u ON u.id = r.user_id
        where r.product_id = $1 AND r.user_id = $2;
	`

	var userReviewDTO ReviewDTO

	err = r.db.QueryRow(ctx, userReviewQuery, productID, userID).Scan(
		&userReviewDTO.Text,
		&userReviewDTO.Rating,
		&userReviewDTO.IsPrivate,
		&userReviewDTO.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			reviews.UserReview = model.Review{}
		} else {
			r.log.Error("[ReviewsStore.GetProductReviewsWithLogin] Error getting user review", slog.String("error", err.Error()))

			return model.Reviews{}, err
		}
	}

	reviews.UserReview = userReviewDTO.ToModel()

	fieldSortOptions := map[string]string{
		"rating": "r.rating",
		"date":   "r.created_at",
	}

	field, ok := fieldSortOptions[sortField]
	if !ok {
		field = "r.created_at"
	}

	sortOrder = utils.ReturnSortOrderOption(sortOrder)

	query := fmt.Sprintf(`
		select r.text, r.rating, r.is_private, u.username, u.avatar_url, r.created_at
		from reviews r
		join users u on u.id = r.user_id
		where r.product_id = $1 and r.user_id != $2
		order by %s %s;
	`, field, sortOrder)

	rows, err := r.db.Query(ctx, query, productID, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.log.Error("[ReviewsStore.GetProductReviewsWithLogin] No reviews")

			return model.Reviews{}, errs.NoReviewsForProduct
		}

		r.log.Info("[ReviewsStore.GetProductReviewsWithLogin] Unexpected error happened", slog.String("error", err.Error()))

		return model.Reviews{}, err
	}

	reviewsDTO, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[ReviewDTO])
	if err != nil {
		r.log.Error("[ReviewsStore.GetProductReviewsWithLogin] Error happened when scanning rows", slog.String("error", err.Error()))

		return model.Reviews{}, err
	}

	if (len(reviewsDTO) == 0) && (userReviewDTO.Rating == 0) {
		return model.Reviews{}, errs.NoReviewsForProduct
	}

	reviews.Reviews = ToReviewModelSlice(reviewsDTO)

	return reviews, nil
}
