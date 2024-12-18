package product

import (
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/gorilla/mux"
)

func (pd *ProductsDelivery) GetProductById(w http.ResponseWriter, r *http.Request) {
	requestID, err := utils.GetContextRequestID(r.Context())
	if err != nil {
		pd.log.Error("[ProductsDelivery.GetProductById] No request ID")
		utils.WriteErrorJSONByError(w, err, pd.errResolver)

		return
	}

	pd.log.Info("[ProductsDelivery.GetProductById] Started executing", slog.Any("request-id", requestID))

	vars := mux.Vars(r)
	id, err := utils.StrToUint32(vars["id"])
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusInternalServerError, errs.InternalServerError)

		return
	}

	userID, _ := utils.GetContextSessionUserID(r.Context())

	byID, err := pd.productByIdGetter.GetProductByID(r.Context(), userID, id)
	if err != nil {
		utils.WriteErrorJSONByError(w, err, pd.errResolver)

		return
	}

	dtoProductByID := newDTOProductCardFromModel(byID)

	utils.WriteJSON(w, http.StatusOK, dtoProductByID)
}
