package rating_updater

import (
	ratingUpdater "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/rating_updater/gen"
	"google.golang.org/grpc"
)

func (r *RatingUpdaterGRPC) Register(server *grpc.Server) {
	ratingUpdater.RegisterRatingUpdaterServer(server, r)
}
