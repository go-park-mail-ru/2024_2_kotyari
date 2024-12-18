package reviews

import (
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
)

func (h *ReviewsHandler) UpdateReview(w http.ResponseWriter, r *http.Request) {
	requestID, err := utils.GetContextRequestID(r.Context())
	if err != nil {
		h.log.Error("[ReviewsHandler.UpdateReview] No request ID")
		utils.WriteErrorJSONByError(w, err, h.errResolver)

		return
	}

	h.log.Info("[ReviewsHandler.UpdateReview] UpdateReview handler", slog.Any("request-id", requestID))

	vars := mux.Vars(r)
	productID, err := utils.StrToUint32(vars["id"])
	if err != nil {
		h.log.Error("[ReviewsHandler.UpdateReview] No ProductID in path")
		utils.WriteErrorJSON(w, http.StatusBadRequest, errs.ParsingURLArg)

		return
	}

	userID, ok := utils.GetContextSessionUserID(r.Context())
	if !ok {
		h.log.Error("[ReviewsHandler.UpdateReview] No UserID")
		utils.WriteErrorJSON(w, http.StatusUnauthorized, errs.UserNotAuthorized)

		return
	}

	var req UpdateReviewRequestDTO
	err = easyjson.UnmarshalFromReader(r.Body, &req)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, errs.BadRequest)

		return
	}

	err = h.reviewsManager.UpdateReview(r.Context(), productID, userID, req.ToModel())
	if err != nil {
		utils.WriteErrorJSONByError(w, err, h.errResolver)

		return
	}

	utils.WriteJSON(w, http.StatusNoContent, nil)
}
