package csat

import (
	"context"
	"errors"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/jackc/pgx/v5"
)

func (s *CSATStore) GetCSAT(ctx context.Context, csat model.CSAT) (model.CSAT, error) {
	const query = `
		select user_id, text, rating, type
		from surveys
		where user_id = $1;
	`

	var csatDTO CSATDTO

	err := s.db.QueryRow(ctx, query, csat.UserID).Scan(&csatDTO.UserID, &csatDTO.Text, &csatDTO.Rating, &csatDTO.Type)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			s.log.Error("[CSATStore.GetCSAT] No csat for this user", err.Error())

			return model.CSAT{}, errs.NoUserCSAT
		}

		s.log.Error("[CSATStore.GetCSAT] Error happened executing query", err.Error())

		return model.CSAT{}, err
	}

	return csatDTO.ToModel(), nil
}
