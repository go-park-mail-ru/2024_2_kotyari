package csat

import (
	"context"
	"errors"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/jackc/pgx/v5"
)

func (s *CSATStore) GetStatistics(ctx context.Context, statisticType model.CSATType) (model.CSATStatistics, error) {
	const query = `
		select count(*) as total, avg(rating) as avg_rating
        from surveys
        where rating between 1 and 10 and type = $1;
	`

	var stat model.CSATStatistics
	err := s.db.QueryRow(ctx, query, statisticType).Scan(&stat.Total, &stat.Avg)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			s.log.Error("[CSATStore.GetStatistics] No rows")

			return model.CSATStatistics{}, errs.NoStatistics
		}

		s.log.Error("[CSATStore.GetStatistics] Unexpected error")

		return model.CSATStatistics{}, err
	}

	ratings, err := s.getCSATRatings(ctx, statisticType)
	if err != nil {
		s.log.Error("[CSATStore.GetStatistics] Error getting csat ratings", slog.String("error", err.Error()))

		return model.CSATStatistics{}, err
	}

	stat.Ratings = ratings

	return stat, nil
}

func (s *CSATStore) getCSATRatings(ctx context.Context, statisticType model.CSATType) ([]model.CSATRating, error) {
	const query = `
		select s.rating, count(surveys.id) as count
        from generate_series(1,10) as s(rating)
        left join surveys on surveys.rating = s.rating
        where type = $1
        group by s.rating
        order by s.rating
	`

	rows, err := s.db.Query(ctx, query, statisticType)
	if err != nil {
		s.log.Error("[CSATStore.getCSATRatings] Error happened fetching rows", slog.String("error", err.Error()))

		return nil, err
	}

	ratings, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.CSATRating])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			s.log.Error("[CSATStore.getCSATRatings] No rows")

			return nil, errs.NoCSATReviews
		}

		s.log.Error("[CSATStore.getCSATRatings] Unexpected error")

		return nil, err
	}

	return ratings, nil
}
