package product

import (
	"log/slog"
	"net/http"
	"strconv"

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
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusInternalServerError, errs.InternalServerError)

		return
	}

	byID, err := pd.repo.GetProductByID(r.Context(), id)
	if err != nil {
		err, code := pd.errResolver.Get(err)
		utils.WriteErrorJSON(w, code, err)

		return
	}

	dtoProductByid := newDTOProductCardFromModel(byID)
	userId, ok := utils.GetContextSessionUserID(r.Context())
	if ok {
		flag, err := pd.checker.ProductInCart(r.Context(), userId, uint32(id))
		if err != nil {
			return
		}

		dtoProductByid.InCart = flag
	}

	utils.WriteJSON(w, http.StatusOK, dtoProductByid)
}
