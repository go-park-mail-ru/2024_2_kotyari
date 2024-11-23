package csat

import (
	"encoding/json"
	grpc_gen "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/csat/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"log/slog"
	"net/http"
)

func (cd *CsatDelivery) GetCsat(r *http.Request, w http.ResponseWriter) {
	var req GetCsatsRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, errs.HTTPErrorResponse{
			ErrorMessage: errs.InvalidJSONFormat.Error(),
		})
		cd.log.Error("[ CsatDelivery.GetCsat ] Ошибка при декодировании запроса", slog.String("error", err.Error()))
		return
	}

	_, ok := utils.GetContextSessionUserID(r.Context())
	if !ok {
		utils.WriteErrorJSON(w, http.StatusUnauthorized, errs.UserNotAuthorized)

		return
	}

	getScatResponse, err := cd.csatGrpcClient.GetCsat(r.Context(), &grpc_gen.GetCsatRequest{Type: req.Type, Rating: req.Rating, Text: req.Text})
	if err != nil {
		err, code := cd.errResolver.Get(err)
		utils.WriteJSON(w, code, errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})

		cd.log.Error("[ CsatDelivery.GetCsat ] Ошибка при передаче на grpc", slog.String("error", err.Error()))
		return
	}

	utils.WriteJSON(w, http.StatusOK, GetCsatsResponse{
		Text: getScatResponse.Text,
	})
}
