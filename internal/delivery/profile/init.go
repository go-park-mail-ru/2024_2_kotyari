package profile

import (
	"context"
	"log/slog"
	"os"

	profilegrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/profile/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type imageSaver interface {
	SaveImage(filename string, file *os.File) (string, error)
}

type addressGetter interface {
	GetAddressByProfileID(ctx context.Context, profileID uint32) (model.Addresses, error)
}

type ProfilesDelivery struct {
	log *slog.Logger

	addressGetter addressGetter
	imageSaver    imageSaver
	errResolver   errs.GetErrorCode

	client profilegrpc.ProfileClient
}

func NewProfilesHandler(
	client profilegrpc.ProfileClient,
	logger *slog.Logger,
	addressGetter addressGetter,
	imageSaver imageSaver,
	errResolver errs.GetErrorCode,
) *ProfilesDelivery {

	return &ProfilesDelivery{
		log:           logger,
		addressGetter: addressGetter,
		imageSaver:    imageSaver,
		errResolver:   errResolver,
		client:        client,
	}
}
