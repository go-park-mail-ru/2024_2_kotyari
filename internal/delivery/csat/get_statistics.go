package csat

import (
	"encoding/json"
	grpc_gen "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/csat/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"log/slog"
	"net/http"
)

func (cd *CsatDelivery) GetStatistics(r *http.Request, w http.ResponseWriter) {
	var req GetStatisticsRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, errs.HTTPErrorResponse{
			ErrorMessage: errs.InvalidJSONFormat.Error(),
		})
		cd.log.Error("[ CsatDelivery.GetStatistics ] Ошибка при декодировании запроса", slog.String("error", err.Error()))
		return
	}

	_, ok := utils.GetContextSessionUserID(r.Context())
	if !ok {
		utils.WriteErrorJSON(w, http.StatusUnauthorized, errs.UserNotAuthorized)

		return
	}

	getStaticsResponse, err := cd.csatGrpcClient.GetStatistics(r.Context(), &grpc_gen.GetStatisticsRequest{Type: req.Type})
	if err != nil {
		err, code := cd.errResolver.Get(err)
		utils.WriteJSON(w, code, errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})

		cd.log.Error("[ CsatDelivery.GetStatistics ] Ошибка при передаче на grpc", slog.String("error", err.Error()))
		return
	}
	httpStats := convertGrpcStatsToHTTP(getStaticsResponse.GetStats())

	utils.WriteJSON(w, http.StatusOK, GetStatisticsResponse{
		Stats:   httpStats,
		Average: getStaticsResponse.GetAverage(),
	})
}
