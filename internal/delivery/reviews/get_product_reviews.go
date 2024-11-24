package reviews

import (
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/gorilla/mux"
)

func (h *ReviewsHandler) GetProductReviews(w http.ResponseWriter, r *http.Request) {
	requestID, err := utils.GetContextRequestID(r.Context())
	if err != nil {
		h.log.Error("[ReviewsHandler.GetProductReviews] No request ID")
		utils.WriteErrorJSONByError(w, err, h.errResolver)

		return
	}

	h.log.Info("[ReviewsHandler.GetProductReviews] GetProductReviews handler", slog.Any("request-id", requestID))

	vars := mux.Vars(r)
	productID, err := utils.StrToUint32(vars["id"])
	if err != nil {
		h.log.Error("[ReviewsHandler.GetProductReviews] No ProductID in path")
		utils.WriteErrorJSON(w, http.StatusBadRequest, errs.ParsingURLArg)

		return
	}

	sortField := r.URL.Query().Get(utils.SearchFieldParam)
	sortOrder := r.URL.Query().Get(utils.SearchOrderParam)

	var reviews model.Reviews

	userID, ok := utils.GetContextSessionUserID(r.Context())
	if !ok {
		reviews, err = h.reviewsGetter.GetProductReviewsNoLogin(r.Context(), productID, sortField, sortOrder)
		if err != nil {
			utils.WriteErrorJSONByError(w, err, h.errResolver)

			return
		}

	} else {
		reviews, err = h.reviewsGetter.GetProductReviewsWithLogin(r.Context(), productID, userID, sortField, sortOrder)
		if err != nil {
			utils.WriteErrorJSONByError(w, err, h.errResolver)

			return
		}
	}

	var resp GetProductReviewsResponseDTO
	reviewsResponse := make([]ReviewResponseDTO, 0, len(reviews.Reviews))
	for _, review := range reviews.Reviews {
		reviewsResponse = append(reviewsResponse, reviewResponseFromModel(review))
	}

	resp = productReviewsFromModel(reviews, reviewsResponse)

	utils.WriteJSON(w, http.StatusOK, resp)
}
