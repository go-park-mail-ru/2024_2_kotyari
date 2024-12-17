package reviews

import (
	"time"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type GetProductReviewsResponseDTO struct {
	TotalReviewCount uint32              `json:"total_review_count"`
	TotalRating      float32             `json:"total_review_rating"`
	UserReview       ReviewResponseDTO   `json:"user_review"`
	Reviews          []ReviewResponseDTO `json:"reviews"`
}

type ReviewResponseDTO struct {
	Username  string    `json:"username,omitempty"`
	AvatarURL string    `json:"avatar_url,omitempty"`
	Text      string    `json:"text"`
	Rating    uint8     `json:"rating"`
	IsPrivate bool      `json:"is_private"`
	CreatedAt time.Time `json:"created_at"`
}

type AddReviewRequestDTO struct {
	Rating    uint8  `json:"rating"`
	Text      string `json:"text"`
	IsPrivate bool   `json:"is_private"`
}

func (r AddReviewRequestDTO) ToModel() model.Review {
	return model.Review{
		Rating:    r.Rating,
		Text:      r.Text,
		IsPrivate: r.IsPrivate,
	}
}

type UpdateReviewRequestDTO struct {
	Rating    uint8  `json:"rating"`
	Text      string `json:"text"`
	IsPrivate bool   `json:"is_private"`
}

func (r UpdateReviewRequestDTO) ToModel() model.Review {
	return model.Review{
		Rating: r.Rating,
		Text:   r.Text,
	}
}

func reviewResponseFromModel(review model.Review) ReviewResponseDTO {
	return ReviewResponseDTO{
		Username:  review.Username,
		AvatarURL: review.AvatarURL,
		Text:      review.Text,
		Rating:    review.Rating,
		IsPrivate: review.IsPrivate,
		CreatedAt: review.CreatedAt,
	}
}

func productReviewsFromModel(r model.Reviews, reviews []ReviewResponseDTO) GetProductReviewsResponseDTO {
	return GetProductReviewsResponseDTO{
		TotalReviewCount: r.TotalReviewCount,
		TotalRating:      r.TotalRating,
		UserReview:       reviewResponseFromModel(r.UserReview),
		Reviews:          reviews,
	}
}
