package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

type UserService interface {
	LoginByEmail(userLoginRequest model.UserLoginRequestDTO) (model.User, error)
	CreateUser(userSignupRequest model.UserSignupRequestDTO) (model.User, error)
	GetSessionUser(email string) (model.User, error)
}

type UserDelivery interface {
	LoginByEmail(r *http.Request) (model.UserDTO, error)
	CreateUser(r *http.Request) (model.UserDTO, error)
	GetSessionUser(email string) (model.UserDTO, error)
}

type Delivery struct {
	Uc UserService
}

func NewUserDelivery(uc UserService) UserDelivery {
	return &Delivery{Uc: uc}
}

func (d *Delivery) LoginByEmail(r *http.Request) (model.UserDTO, error) {
	var loginRequest model.UserLoginRequestDTO

	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		return model.UserDTO{}, err
	}

	if err = utils.ValidateEmailAndPassword(loginRequest.Email, loginRequest.Password); err != nil {
		return model.UserDTO{}, err
	}

	user, err := d.Uc.LoginByEmail(loginRequest)
	if err != nil {
		return model.UserDTO{}, err
	}

	return *user.ToUserDTO(), nil
}

func (d *Delivery) CreateUser(r *http.Request) (model.UserDTO, error) {
	var signupRequest model.UserSignupRequestDTO

	err := json.NewDecoder(r.Body).Decode(&signupRequest)
	if err != nil {
		return model.UserDTO{}, err
	}

	if err = utils.ValidateRegistration(signupRequest); err != nil {
		return model.UserDTO{}, err
	}

	user, err := d.Uc.CreateUser(signupRequest)
	if err != nil {
		return model.UserDTO{}, err
	}

	return *user.ToUserDTO(), nil
}

func (d *Delivery) GetSessionUser(email string) (model.UserDTO, error) {
	user, err := d.Uc.GetSessionUser(email)
	if err != nil {
		return model.UserDTO{}, err
	}

	return *user.ToUserDTO(), nil
}
