package wish_list

import (
	"encoding/json"
	wishlistgrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/wish_list/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"google.golang.org/grpc/status"
	"log/slog"
	"net/http"
)

func (wld *WishlistDelivery) CreateWishlist(w http.ResponseWriter, r *http.Request) {
	requestID, err := utils.GetContextRequestID(r.Context())
	if err != nil {
		wld.log.Error("[WishlistDelivery.CreateWishlist] No request ID")
		utils.WriteErrorJSONByError(w, err, wld.errResolver)

		return
	}

	wld.log.Info("[WishlistDelivery.CreateWishlist] Started executing", slog.Any("request-id", requestID))

	userID, ok := utils.GetContextSessionUserID(r.Context())
	if !ok {
		utils.WriteErrorJSON(w, http.StatusUnauthorized, errs.UserNotAuthorized)

		return
	}

	var request createWishlistRequest

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, err)

		return
	}

	_, err = wld.client.CreateWishlist(r.Context(), &wishlistgrpc.CreateWishlistRequest{
		UserId: userID,
		Name:   request.Name,
	})
	if err != nil {
		grpcErr, okis := status.FromError(err)
		if okis {
			switch grpcErr.Code() {

			}

			wld.log.Error("[ WishlistDelivery.CreateWishlist ] ok",
				slog.String("func", "client.CreateWishlist"),
				slog.String("error", err.Error()),
			)

			return
		}

		wld.log.Error("[ WishlistDelivery.CreateWishlist ] not ok",
			slog.String("func", "client.CreateWishlist"),
			slog.String("error", err.Error()),
		)

		utils.WriteErrorJSON(w, http.StatusInternalServerError, errs.InternalServerError)

		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}
