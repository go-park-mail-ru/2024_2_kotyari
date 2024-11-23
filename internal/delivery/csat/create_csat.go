package csat

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"log/slog"
	"net/http"
)

func (cd *CsatDelivery) CreateCsat(r *http.Request, w http.ResponseWriter) {
	var req CreateCsatRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, errs.HTTPErrorResponse{
			ErrorMessage: errs.InvalidJSONFormat.Error(),
		})
		cd.log.Error("[ CsatDelivery.CreateCsat ] Ошибка при декодировании запроса", slog.String("error", err.Error()))
		return
	}

	_, ok := utils.GetContextSessionUserID(r.Context())
	if !ok {
		utils.WriteErrorJSON(w, http.StatusUnauthorized, errs.UserNotAuthorized)

		return
	}

	createCsatResponse, err := cd.csatGrpcClient.CreateCsat(r.Context(), req.ToGrpcCreateCsatRequest())
	if err != nil {
		err, code := cd.errResolver.Get(err)
		utils.WriteJSON(w, code, errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})

		cd.log.Error("[ CsatDelivery.CreateCsat ] Ошибка при передаче на grpc", slog.String("error", err.Error()))
		return
	}
	utils.WriteJSON(w, http.StatusOK, FromGrpcResponse(createCsatResponse))
}
