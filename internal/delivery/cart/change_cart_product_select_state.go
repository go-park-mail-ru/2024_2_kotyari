package cart

import (
	"encoding/json"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/gorilla/mux"
)

func (ch *CartHandler) ChangeCartProductSelectedState(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID, err := utils.StrToUint32(vars["id"])
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, errs.ParsingURLArg)

		return
	}

	userID := utils.GetContextSessionUserID(r.Context())

	var req ChangeCartProductSelectedStateRequest

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, errs.InvalidJSONFormat)

		return
	}

	err = ch.cartManager.ChangeCartProductSelectedState(r.Context(), productID, userID, req.ToModel())
	if err != nil {
		err, code := ch.errResolver.Get(err)
		utils.WriteErrorJSON(w, code, err)

		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}
