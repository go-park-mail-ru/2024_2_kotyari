package csat

import (
	"time"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/jackc/pgx/v5/pgtype"
)

type CSATDTO struct {
	ID        uint32
	UserID    uint32 `db:"user_id"`
	Rating    pgtype.Uint32
	Type      model.CSATType
	Text      pgtype.Text
	UpdatedAt time.Time
	CreatedAt time.Time
}

func (c CSATDTO) ToModel() model.CSAT {
	return model.CSAT{
		ID:     c.ID,
		UserID: c.UserID,
		Rating: c.Rating.Uint32,
		Type:   c.Type,
		Text:   c.Text.String,
	}
}
