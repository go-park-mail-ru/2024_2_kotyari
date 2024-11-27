package category

import (
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (cd *CategoriesDelivery) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	requestID, err := utils.GetContextRequestID(r.Context())
	if err != nil {
		cd.log.Error("[CategoriesDelivery.GetAllCategories] No request ID")
		utils.WriteErrorJSONByError(w, err, cd.errResolver)

		return
	}

	cd.log.Info("[CategoriesDelivery.GetAllCategories] Started executing", slog.Any("request-id", requestID))

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
