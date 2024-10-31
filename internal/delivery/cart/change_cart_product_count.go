package cart

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/gorilla/mux"
)

func (ch *CartHandler) ChangeCartProductQuantity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID, err := strconv.ParseUint(vars["id"], 10, 32)
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

	err = ch.cartManager.ChangeCartProductCount(r.Context(), uint32(productID), req.ToModel())
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, err)

		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}
