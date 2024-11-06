package category

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (cd *CategoriesDelivery) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := cd.repo.GetAllCategories(r.Context())
	if err != nil {
		err, i := cd.errResolver.Get(err)

		utils.WriteErrorJSON(w, i, err)
	}

	dtoCategories := make([]dtoCategory, 0, len(categories))
	for _, category := range categories {
		dtoCategories = append(dtoCategories, dtoCategory(category))
	}

	utils.WriteJSON(w, http.StatusOK, dtoCategories)
}
