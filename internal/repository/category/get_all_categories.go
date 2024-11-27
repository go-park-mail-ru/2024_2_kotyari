package category

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

func (cs *CategoriesStore) GetAllCategories(ctx context.Context) ([]model.Category, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return nil, err
	}

	cs.log.Info("[CategoriesStore.GetAllCategories] Started executing", slog.Any("request-id", requestID))

	const query = `SELECT name, link_to, picture
                   FROM categories
                   WHERE active=true`

	rows, err := cs.db.Query(ctx, query)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errs.CategoriesDoesNotExits
		}

		return nil, err
	}

	var categories []model.Category
	for rows.Next() {
		var category model.Category

		err = rows.Scan(&category.Name, &category.LinkTo, &category.Picture)
		if err != nil {
			cs.log.WarnContext(ctx, fmt.Sprintf("[ CategoriesStore.GetAllCategories ] error while scanning row: %s", err))
			return nil, err
		}

		categories = append(categories, category)
	}

	if len(categories) == 0 {
		return nil, errs.CategoriesDoesNotExits
	}

	return categories, nil
}
