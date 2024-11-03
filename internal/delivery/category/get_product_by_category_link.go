package category

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/gorilla/mux"
)

func (cd *CategoriesDelivery) GetProductsByCategoryLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	link := vars["link"]

	products, err := cd.repo.GetProductsByCategoryLink(r.Context(), link)
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
