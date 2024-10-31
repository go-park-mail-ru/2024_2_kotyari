package product

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

const queryGetOptions = `
    SELECT 
        po.values::jsonb
    FROM product_options po
    WHERE po.product_id = $1;
`

func (ps *ProductsStore) getProductOptions(ctx context.Context, productID uint64) (model.Options, error) {
	rowsOptions, err := ps.db.Query(ctx, queryGetOptions, productID)
	if err != nil {
		ps.log.Error("[ ProductsStore.getProductOptions ] Error executing options query", slog.String("error", err.Error()))
		return model.Options{}, err
	}
	defer rowsOptions.Close()

	var options model.Options

	for rowsOptions.Next() {
		var optionValuesJSON []byte

		err = rowsOptions.Scan(&optionValuesJSON)
		if err != nil {
			ps.log.Error("[ ProductsStore.getProductOptions  ] Error scanning option", "error", slog.String("error", err.Error()))
			return model.Options{}, err
		}

		var dtoOptions []dtoOptionBlock
		err = json.Unmarshal(optionValuesJSON, &dtoOptions)
		if err != nil {
			ps.log.Error("[ ProductsStore.getProductOptions ] Error decoding options", slog.String("error", err.Error()))
			return model.Options{}, err
		}

		// Convert DTO to model and append to options.Values
		for _, dtoOpt := range dtoOptions {
			optionsBlock := dtoOpt.ToModel()
			options.Values = append(options.Values, optionsBlock)
		}
	}

	if rowsOptions.Err() != nil {
		ps.log.Error("[ ProductsStore.getProductOptions ] Error iterating over options rows", slog.String("error", rowsOptions.Err().Error()))
		return model.Options{}, rowsOptions.Err()
	}

	return options, nil
}
