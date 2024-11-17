package utils

import "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"

func ValidateReviewRating(review model.Review) bool {
	return !(review.Rating < 1 || review.Rating > 5)
}
