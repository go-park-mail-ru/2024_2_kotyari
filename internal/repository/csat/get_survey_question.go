package csat

import (
	"context"
	"errors"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/jackc/pgx/v5"
)

func (s *CSATStore) GetSurveyQuestion(ctx context.Context, statisticType model.CSATType) (model.SurveyQuestion, error) {
	const query = `
		select question_text 
		from survey_questions
		where type = $1;
	`

	var surveyQuestion model.SurveyQuestion

	err := s.db.QueryRow(ctx, query, statisticType).Scan(&surveyQuestion.QuestionText)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			s.log.Error("[CSATStore.GetSurveyQuestion] No questions", err.Error())

			return model.SurveyQuestion{}, errs.NoQuestionText
		}

		s.log.Error("[CSATStore.GetSurveyQuestion] Unexpected error", err.Error())

		return model.SurveyQuestion{}, err
	}

	return surveyQuestion, err
}
