package cart

import (
	"encoding/json"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/gorilla/mux"
)

func (ch *CartHandler) ChangeCartProductQuantity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID, err := utils.StrToUint32(vars["id"])
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, errs.ParsingURLArg)

		return
	}

	userID := utils.GetContextSessionUserID(r.Context())
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, err)

		return
	}

	var req ChangeCartProductCountRequest

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, err)

		return
	}

	err = ch.cartManager.ChangeCartProductCount(r.Context(), productID, req.ToModel(), userID)
	if err != nil {
		err, code := ch.errResolver.Get(err)
		utils.WriteErrorJSON(w, code, err)

		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}
