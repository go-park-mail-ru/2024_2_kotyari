package recommendations

import (
	productDTO "github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/product"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *RecDelivery) GetRecommendations(w http.ResponseWriter, r *http.Request) {
	requestID, err := utils.GetContextRequestID(r.Context())
	if err != nil {
		h.log.Error("[ RecDelivery.GetRecommendations] Can't get RequestId")
		utils.WriteErrorJSONByError(w, err, h.errResolver)

		return
	}

	h.log.Info("", slog.Any("request-id", requestID))

	vars := mux.Vars(r)
	productId, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusInternalServerError, errs.InternalServerError)

		return
	}

	products, err := h.repo.GetRecommendations(r.Context(), productId)
	if err != nil {
		h.log.Error("[ RecDelivery.GetRecommendations] Ошибка при получении товаров из уроня репозитория",
			slog.String("method", r.Method),
			slog.String("url", r.URL.String()),
			slog.String("error", err.Error()),
		)

		utils.WriteErrorJSON(w, http.StatusNotFound, errs.ProductsDoesNotExists)

		return
	}

	dtoProducts := make([]productDTO.DtoProductCatalog, len(products))
	for i, product := range products {
		dtoProducts[i] = productDTO.NewDTOProductCatalogFromModel(product)
	}

	utils.WriteJSON(w, http.StatusOK, dtoProducts)
}
