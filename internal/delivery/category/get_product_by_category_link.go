package category

import (
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/gorilla/mux"
)

func (cd *CategoriesDelivery) GetProductsByCategoryLink(w http.ResponseWriter, r *http.Request) {
	requestID, err := utils.GetContextRequestID(r.Context())
	if err != nil {
		cd.log.Error("[CategoriesDelivery.GetProductsByCategoryLink] No request ID")
		utils.WriteErrorJSONByError(w, err, cd.errResolver)

		return
	}

	cd.log.Info("[CategoriesDelivery.GetProductsByCategoryLink] Started executing", slog.Any("request-id", requestID))

	vars := mux.Vars(r)
	link := vars["link"]

	sortField := r.URL.Query().Get(utils.SearchFieldParam)
	sortOrder := r.URL.Query().Get(utils.SearchOrderParam)

	products, err := cd.repo.GetProductsByCategoryLink(r.Context(), link, sortField, sortOrder)
	if err != nil {
		err, i := cd.errResolver.Get(err)

		cd.log.Error("GetProductsByCategoryLink", "err", err, "response", i)
		utils.WriteErrorJSON(w, i, err)
		return
	}

	dtoProducts := make([]dtoProductCatalog, 0, len(products))

	for _, product := range products {
		dto := toDTOProductCatalogFromModel(product)
		dtoProducts = append(dtoProducts, dto)
	}

	utils.WriteJSON(w, http.StatusOK, dtoProducts)
}
