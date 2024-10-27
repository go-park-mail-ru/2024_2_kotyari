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
		ps.log.Error("[ ProductsStore.GetProductByID ] Error executing options query", slog.String("error", err.Error()))
		return model.Options{}, err
	}
	defer rowsOptions.Close()

	var options model.Options

	for rowsOptions.Next() {
		var optionValuesJSON []byte

		err = rowsOptions.Scan(&optionValuesJSON)
		if err != nil {
			ps.log.Error("[ ProductsStore.GetProductByID ] Error scanning option", "error", slog.String("error", err.Error()))
			return model.Options{}, err
		}

		var opts dtoOptionBlock
		err = json.Unmarshal(optionValuesJSON, &opts)
		if err != nil {
			ps.log.Error("[ ProductsStore.GetProductByID ] Error decoding options", slog.String("error", err.Error()))
			return model.Options{}, err
		}

		optionsBlock := model.OptionsBlock{
			Title: opts.Title,
			Type:  opts.Type,
		}

		for _, dtoOpt := range opts.Options {
			option := dtoOpt.ToModel()
			optionsBlock.Options = append(optionsBlock.Options, option)
		}

		options.Values = append(options.Values, optionsBlock)
	}

	if rowsOptions.Err() != nil {
		ps.log.Error("[ ProductsStore.GetProductByID ] Error iterating over options rows", slog.String("error", rowsOptions.Err().Error()))
		return model.Options{}, rowsOptions.Err()
	}

	return options, nil
}
