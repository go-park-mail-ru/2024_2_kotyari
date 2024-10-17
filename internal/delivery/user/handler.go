package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

type UsecaseInterface interface {
	LoginByEmail(userLoginRequest *model.UserLoginRequestDTO) (*model.User, bool)
	CreateUser(userSignupRequest *model.UserSignupRequestDTO) (*model.User, error)
	GetUserByEmail(email string) (*model.User, bool)
}

type Delivery struct {
	Uc UsecaseInterface
}

func NewUserDelivery(uc UsecaseInterface) *Delivery {
	return &Delivery{Uc: uc}
}

func (d *Delivery) LoginByEmail(r *http.Request) (*model.User, error) {
	var loginRequest model.UserLoginRequestDTO

	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		return nil, err
	}

	if err = utils.ValidateEmailAndPassword(loginRequest.Email, loginRequest.Password); err != nil {
		return nil, err
	}

	user, exists := d.Uc.LoginByEmail(&loginRequest)
	if !exists || !utils.VerifyPassword(user.Password, loginRequest.Password) {
		return nil, errs.WrongCredentials
	}

	return user, nil
}

func (d *Delivery) CreateUser(r *http.Request) (*model.User, error) {
	var signupRequest model.UserSignupRequestDTO

	err := json.NewDecoder(r.Body).Decode(&signupRequest)
	if err != nil {
		return nil, err
	}

	if err = utils.ValidateRegistration(signupRequest); err != nil {
		return nil, err
	}

	return d.Uc.CreateUser(&signupRequest)
}
