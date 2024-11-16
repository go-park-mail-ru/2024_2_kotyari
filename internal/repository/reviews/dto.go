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
	ID        uint32
	Text      pgtype.Text
	Username  string
	AvatarURL string
	Rating    uint8
	IsPrivate bool
	UpdatedAt time.Time
	CreatedAt time.Time
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
