package product

import (
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (pd *ProductsDelivery) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	pd.log.Debug("[ ProductsDelivery.GetAllProducts ] is running ]")

	products, err := pd.repo.GetAllProducts(r.Context())
	if err != nil {
		pd.log.Error("[ ProductsDelivery.GetAllProducts ] no products ]",
			slog.String("method", r.Method),
			slog.String("url", r.URL.String()),
			slog.String("error", err.Error()),
		)

		utils.WriteErrorJSON(w, errs.ProductsDoesNotExists, http.StatusNotFound)

		return
	}

	dtoProducts := make([]dtoProductCatalog, len(products))
	for i, product := range products {
		dtoProducts[i] = newDTOProductCatalogFromModel(product)
	}

	utils.WriteJSON(w, http.StatusOK, dtoProducts)
}
