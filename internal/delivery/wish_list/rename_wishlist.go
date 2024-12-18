package wish_list

import (
	"fmt"
	"log/slog"
	"net/http"

	wishlistgrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/wish_list/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/mailru/easyjson"
	"google.golang.org/grpc/status"
)

func (wld *WishlistDelivery) RenameWishlist(w http.ResponseWriter, r *http.Request) {
	requestID, err := utils.GetContextRequestID(r.Context())
	if err != nil {
		wld.log.Error("[WishlistDelivery.RenameWishlist] No request ID")
		utils.WriteErrorJSONByError(w, err, wld.errResolver)

		return
	}

	wld.log.Info("[WishlistDelivery.RenameWishlist] Started executing", slog.Any("request-id", requestID))

	userID, ok := utils.GetContextSessionUserID(r.Context())
	if !ok {
		utils.WriteErrorJSON(w, http.StatusUnauthorized, errs.UserNotAuthorized)

		return
	}

	var request renameWishlistRequest

	if err = easyjson.UnmarshalFromReader(r.Body, &request); err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, err)

		return
	}

	fmt.Println("[ test ] ", "link", request.Link)

	_, err = wld.client.RenameWishlist(r.Context(), &wishlistgrpc.RenameWishlistRequest{
		UserId:  userID,
		NewName: request.NewName,
		Link:    request.Link,
	})
	if err != nil {
		grpcErr, okis := status.FromError(err)
		if okis {
			switch grpcErr.Code() {

			}

			wld.log.Error("[ WishlistDelivery.RenameWishlist ] ok",
				slog.String("func", "client.RenameWishlist"),
				slog.String("error", err.Error()),
			)

			return
		}

		wld.log.Error("[ WishlistDelivery.RenameWishlist ] not ok",
			slog.String("func", "client.RenameWishlist"),
			slog.String("error", err.Error()),
		)

		utils.WriteErrorJSON(w, http.StatusInternalServerError, errs.InternalServerError)

		return
	}

	utils.WriteJSON(w, http.StatusNoContent, nil)
}
