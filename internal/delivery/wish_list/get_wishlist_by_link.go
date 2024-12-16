package wish_list

import (
	"fmt"
	wishlistgrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/wish_list/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/gorilla/mux"
	"google.golang.org/grpc/status"
	"log/slog"
	"net/http"
)

func (wld *WishlistDelivery) GetWishlistByLink(w http.ResponseWriter, r *http.Request) {
	requestID, err := utils.GetContextRequestID(r.Context())
	if err != nil {
		wld.log.Error("[WishlistDelivery.GetWishlistByLink] No request ID")
		utils.WriteErrorJSONByError(w, err, wld.errResolver)
		return
	}

	wld.log.Info("[WishlistDelivery.GetWishlistByLink] Started executing", slog.Any("request-id", requestID))
	vars := mux.Vars(r)
	link := vars["link"]

	if link == "" {
		utils.WriteErrorJSON(w, http.StatusBadRequest, fmt.Errorf("link is empty"))
		return
	}

	list, err := wld.client.GetWishlistByLink(r.Context(), &wishlistgrpc.GetWishlistByLinkRequest{
		Link: link,
	})
	if err != nil {
		grpcErr, okis := status.FromError(err)
		if okis {
			wld.log.Error("[WishlistDelivery.GetWishlistByLink] gRPC error",
				slog.String("func", "client.GetWishlistByLink"),
				slog.String("error", grpcErr.Message()),
			)
			utils.WriteErrorJSON(w, http.StatusInternalServerError, fmt.Errorf(grpcErr.Message()))
			return
		}

		wld.log.Error("[WishlistDelivery.GetWishlistByLink] Internal error",
			slog.String("func", "client.GetWishlistByLink"),
			slog.String("error", err.Error()),
		)
		utils.WriteErrorJSON(w, http.StatusInternalServerError, errs.InternalServerError)
		return
	}

	items := list.GetWishlist().GetItems()
	if items == nil {
		items = []*wishlistgrpc.WishlistItem{} // Инициализация пустого списка
	}

	ids := make([]uint32, len(items))
	for i, it := range items {
		ids[i] = it.GetProductId()
	}

	ds, err := wld.productsGetter.GetProductsByIDs(r.Context(), ids)
	if err != nil {
		wld.log.Error("[WishlistDelivery.GetWishlistByLink] Error getting products",
			slog.String("func", "productsGetter.GetProductsByIDs"),
			slog.String("error", err.Error()),
		)
		utils.WriteErrorJSON(w, http.StatusInternalServerError, errs.InternalServerError)
		return
	}

	isAuthor := false
	userID, ok := utils.GetContextSessionUserID(r.Context())
	if ok {
		wld.log.Info("[ проверка id ]",
			slog.Any("user-id", userID),
			slog.Any("creator", list.CreatorId))

		isAuthor = userID == list.CreatorId
	}

	responseItems := make([]item, len(ds))
	for i, d := range ds {
		responseItems[i] = item{
			Id:       d.ID,
			Title:    d.Title,
			ImageUrl: d.ImageURL,
			Price:    d.Price,
		}
	}

	resp := getWishlistByLinkResponse{
		IsAuthor: isAuthor,
		Items:    responseItems,
		Link:     link,
		Name:     list.GetWishlist().GetName(),
	}

	utils.WriteJSON(w, http.StatusOK, resp)
}
