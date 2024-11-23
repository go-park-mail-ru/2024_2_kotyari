package repository

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (s *CSATStore) CreateCSAT(ctx context.Context, csat model.CSAT) error {
	const query = `
		insert into surveys (user_id, text, rating, type)
		values ($1, $2, $3, $4)
	`

	commandTag, err := s.db.Exec(ctx, query, csat.UserID, csat.Text, csat.Rating, csat.Type)
	if err != nil {
		s.log.Error("[CSATStore.CreateCSAT] Error happened executing query", err.Error())

		return err
	}

	if commandTag.RowsAffected() != 1 {
		s.log.Error("[CSATStore.CreateCSAT] No Rows affected")

		return err
	}

	return nil
}
