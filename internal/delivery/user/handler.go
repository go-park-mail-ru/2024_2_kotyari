package user

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/user"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

type Delivery struct {
	uc user.UsecaseInterface
}

func NewUserDelivery(uc user.UsecaseInterface) *Delivery {
	return &Delivery{uc: uc}
}

func (d *Delivery) GetUserByEmail(uRequest model.UserApiRequest) (model.User, bool) {
	return d.uc.GetUserByEmail(uRequest.Email)
}

func (d *Delivery) CreateUser(uRequest model.UserApiRequest) error {
	salt, err := utils.GenerateSalt()

	if err != nil {
		return err
	}
	hashedPassword := utils.HashPassword(uRequest.Password, salt)

	newUser := model.User{
		Username: uRequest.Username,
		Password: hashedPassword,
	}

	return d.uc.CreateUser(uRequest.Email, newUser)
}
