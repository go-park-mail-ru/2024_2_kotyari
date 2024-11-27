package product

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

const queryGetOptions = `
    SELECT 
        po.values::jsonb
    FROM product_options po
    WHERE po.product_id = $1;
`

func (ps *ProductsStore) getProductOptions(ctx context.Context, productID uint64) (model.Options, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return model.Options{}, err
	}

	ps.log.Info("[ProductsStore.getProductOptions] Started executing", slog.Any("request-id", requestID))

	rows, err := ps.db.Query(ctx, queryGetOptions, productID)
	if err != nil {
		ps.log.Error("[ ProductsStore.getProductOptions ] ошибка выполнения запроса", slog.String("error", err.Error()))
		return model.Options{}, err
	}
	defer rows.Close()

	var options model.Options

	for rows.Next() {
		var optionValuesJSON []byte

		err = rows.Scan(&optionValuesJSON)
		if err != nil {
			ps.log.Error("[ ProductsStore.getProductOptions ] ошибка сканирования", "error", slog.String("error", err.Error()))
			return model.Options{}, err
		}

		var dtoOptions []dtoOptionBlock
		err = json.Unmarshal(optionValuesJSON, &dtoOptions)
		if err != nil {
			ps.log.Error("[ ProductsStore.getProductOptions ] ошибка unmarshal опции", slog.String("error", err.Error()))
			return model.Options{}, err
		}

		for _, dtoOpt := range dtoOptions {
			optionsBlock := dtoOpt.ToModel()
			options.Values = append(options.Values, optionsBlock)
		}
	}

	if len(options.Values) == 0 {
		ps.log.Warn("[ ProductsStore.getProductOptions ] не найдены опции ]")
		return model.Options{}, errs.OptionsDoesNotExists
	}

	return options, nil
}
