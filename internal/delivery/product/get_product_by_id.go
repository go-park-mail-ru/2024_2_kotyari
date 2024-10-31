package product

import (
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/gorilla/mux"
)

func (pd *ProductsDelivery) GetProductById(w http.ResponseWriter, r *http.Request) {
	pd.log.Debug("[ ProductsDelivery.GetProductById ] is running ]")

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		utils.WriteErrorJSON(w, errs.InternalServerError, http.StatusInternalServerError)

		return
	}

	byID, err := pd.repo.GetProductByID(r.Context(), id)
	if err != nil {
		err, code := pd.errResolver.Get(err)
		utils.WriteErrorJSON(w, err, code)

		return
	}

	dtoProductByid := newDTOProductCardFromModel(byID)

	utils.WriteJSON(w, http.StatusOK, dtoProductByid)
}
