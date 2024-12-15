package user

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

type promoCodesMessageProducer interface {
	AddPromoCode(ctx context.Context, userID uint32, promoID uint32) error
}

type usersRepository interface {
	CreateUser(ctx context.Context, userModel model.User) (model.User, error)
	GetUserByEmail(ctx context.Context, userModel model.User) (model.User, error)
	GetUserByUserID(ctx context.Context, id uint32) (model.User, error)
}

type UsersService struct {
	userRepo       usersRepository
	producer       promoCodesMessageProducer
	inputValidator *utils.InputValidator
	log            *slog.Logger
}

func NewUserService(usersRepository usersRepository, producer promoCodesMessageProducer, inputValidator *utils.InputValidator, log *slog.Logger) *UsersService {
	return &UsersService{
		userRepo:       usersRepository,
		producer:       producer,
		inputValidator: inputValidator,
		log:            log,
	}
}
