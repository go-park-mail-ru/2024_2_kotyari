package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

type UsecaseInterface interface {
	GetUserByEmail(userLoginRequest *model.UserLoginRequestDTO) (*model.User, bool)
	CreateUser(userSignupRequest *model.UserSignupRequestDTO) (*model.User, error)
}

type Delivery struct {
	uc UsecaseInterface
}

func NewUserDelivery(uc UsecaseInterface) *Delivery {
	return &Delivery{uc: uc}
}

func (d *Delivery) GetUserByEmail(r *http.Request) (*model.User, error) {
	var loginRequest model.UserLoginRequestDTO

	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		return nil, err
	}

	if err = utils.ValidateEmailAndPassword(loginRequest.Email, loginRequest.Password); err != nil {
		return nil, err
	}

	user, exists := d.uc.GetUserByEmail(&loginRequest)
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

	return d.uc.CreateUser(&signupRequest)
}
