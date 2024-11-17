package reviews

import (
	"time"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/jackc/pgx/v5/pgtype"
)

const (
	anonymousUsername = "Аноним"
	anonymousAvatar   = "files/default.jpeg"
)

type ReviewDTO struct {
	ID        uint32      `db:"id"`
	ProductID uint32      `db:"product_id"`
	UserID    uint32      `db:"user_id"`
	Text      pgtype.Text `db:"text"`
	Username  string      `db:"username"`
	AvatarURL string      `db:"avatar_url"`
	Rating    uint8       `db:"rating"`
	IsPrivate bool        `db:"is_private"`
	UpdatedAt time.Time   `db:"updated_at"`
	CreatedAt time.Time   `db:"created_at"`
}

func (r ReviewDTO) ToModel() model.Review {
	if r.IsPrivate {
		r.Username = anonymousUsername
		r.AvatarURL = anonymousAvatar
	}

	return model.Review{
		ID:        r.ID,
		Username:  r.Username,
		IsPrivate: r.IsPrivate,
		Rating:    r.Rating,
		AvatarURL: r.AvatarURL,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
		Text:      r.Text.String,
	}
}

func ToReviewModelSlice(reviewsDTO []ReviewDTO) []model.Review {
	var reviewModelSlice []model.Review

	for _, review := range reviewsDTO {
		reviewModelSlice = append(reviewModelSlice, review.ToModel())
	}

	return reviewModelSlice
}
