package cart

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/gorilla/mux"
)

func (ch *CartHandler) RemoveProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID, err := utils.StrToUint32(vars["id"])
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, errs.ParsingURLArg)

		return
	}

	userID := utils.GetContextSessionUserID(r.Context())

	err = ch.cartManager.RemoveProduct(r.Context(), productID, userID)
	if err != nil {
		err, code := ch.errResolver.Get(err)
		utils.WriteErrorJSON(w, code, err)

		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}