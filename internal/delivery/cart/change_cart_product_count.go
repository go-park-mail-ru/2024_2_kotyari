package cart

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/gorilla/mux"
)

func (ch *CartHandler) ChangeCartProductQuantity(w http.ResponseWriter, r *http.Request) {
	requestID, err := utils.GetContextRequestID(r.Context())
	if err != nil {
		ch.log.Error("[CartHandler.ChangeCartProductQuantity] No request ID")
		utils.WriteErrorJSONByError(w, err, ch.errResolver)

		return
	}

	ch.log.Info("[CartHandler.ChangeCartProductQuantity] Started executing", slog.Any("request-id", requestID))

	vars := mux.Vars(r)
	productID, err := utils.StrToUint32(vars["id"])
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, errs.ParsingURLArg)

		return
	}

	userID, ok := utils.GetContextSessionUserID(r.Context())
	if !ok {
		utils.WriteErrorJSON(w, http.StatusUnauthorized, errs.UserNotAuthorized)
	}

	var (
		req  ChangeCartProductCountRequest
		resp ChangeCartProductCountResponse
	)

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, err)

		return
	}

	cartProductCount, err := ch.cartManager.ChangeCartProductCount(r.Context(), productID, req.ToModel(), userID)
	if err != nil {
		err, code := ch.errResolver.Get(err)
		utils.WriteErrorJSON(w, code, err)

		return
	}

	resp.Count = cartProductCount

	utils.WriteJSON(w, http.StatusOK, resp)
}
