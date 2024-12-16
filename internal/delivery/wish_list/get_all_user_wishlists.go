package wish_list

import (
	wishlistgrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/wish_list/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"google.golang.org/grpc/status"
	"log/slog"
	"net/http"
)

func (wld *WishlistDelivery) GetAllUserWishlists(w http.ResponseWriter, r *http.Request) {
	requestID, err := utils.GetContextRequestID(r.Context())
	if err != nil {
		wld.log.Error("[WishlistDelivery.GetAllUserWishlists] No request ID")
		utils.WriteErrorJSON(w, http.StatusUnauthorized, err)

		return
	}

	wld.log.Info("[WishlistDelivery.GetAllUserWishlists] Started executing", slog.Any("request-id", requestID))

	userID, ok := utils.GetContextSessionUserID(r.Context())
	if !ok {
		utils.WriteErrorJSON(w, http.StatusUnauthorized, errs.UserNotAuthorized)

		return
	}

	wishlists, err := wld.client.GetAllUserWishlists(r.Context(), &wishlistgrpc.GetAllWishlistsRequest{
		UserId: userID,
	})

	if err != nil {
		wld.log.Error("[ WishlistDelivery.GetAllUserWishlists ] not ok",
			slog.String("func", "client.GetAllUserWishlists"),
			slog.String("error", err.Error()),
		)

		_, okis := status.FromError(err)
		if okis {

			utils.WriteErrorJSON(w, http.StatusInternalServerError, err)

			return
		}

		utils.WriteErrorJSON(w, http.StatusInternalServerError, errs.InternalServerError)

		return
	}

	wld.log.Info("[  ]")

	utils.WriteJSON(w, http.StatusOK, wishlists.Wishlists)
}
