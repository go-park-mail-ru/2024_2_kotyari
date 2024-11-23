package reviews

import (
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/gorilla/mux"
)

func (h *ReviewsHandler) DeleteReview(w http.ResponseWriter, r *http.Request) {
	requestID, err := utils.GetContextRequestID(r.Context())
	if err != nil {
		h.log.Error("[ReviewsHandler.DeleteReview] No request ID")
		utils.WriteErrorJSONByError(w, err, h.errResolver)

		return
	}

	h.log.Info("[ReviewsHandler.DeleteReview] DeleteReview handler", slog.Any("request-id", requestID))

	vars := mux.Vars(r)
	productID, err := utils.StrToUint32(vars["id"])
	if err != nil {
		h.log.Error("[ReviewsHandler.DeleteReview] No ProductID in path")
		utils.WriteErrorJSON(w, http.StatusBadRequest, errs.ParsingURLArg)

		return
	}

	userID, ok := utils.GetContextSessionUserID(r.Context())
	if !ok {
		h.log.Error("[ReviewsHandler.DeleteReview] No UserID")
		utils.WriteErrorJSON(w, http.StatusUnauthorized, errs.UserNotAuthorized)

		return
	}

	err = h.reviewsManager.DeleteReview(r.Context(), productID, userID)
	if err != nil {
		utils.WriteErrorJSONByError(w, err, h.errResolver)

		return
	}

	utils.WriteJSON(w, http.StatusNoContent, nil)
}
