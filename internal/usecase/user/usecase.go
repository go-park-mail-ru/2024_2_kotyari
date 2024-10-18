package user

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/user"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

type UserStorer interface {
	GetUserByEmail(email string) (model.User, bool)
	InsertUser(user model.User) (model.User, error)
}

func NewUserUseCase(userRepo UserStorer) user.UserService {
	return &UseCase{repo: userRepo}
}

type UseCase struct {
	repo UserStorer
}

func (u *UseCase) LoginByEmail(userLoginRequest model.UserLoginRequestDTO) (model.User, error) {
	userModel, exists := u.repo.GetUserByEmail(userLoginRequest.Email)
	if !exists || !utils.VerifyPassword(userModel.Password, userLoginRequest.Password) {
		return model.User{}, errs.WrongCredentials
	}

	return userModel, nil
}

func (u *UseCase) CreateUser(userSignupRequest model.UserSignupRequestDTO) (model.User, error) {
	salt, err := utils.GenerateSalt()
	if err != nil {
		return model.User{}, err
	}

	userSignupRequest.Password = utils.HashPassword(userSignupRequest.Password, salt)
	userModel := userSignupRequest.ToUserModel()

	return u.repo.InsertUser(*userModel)
}

func (u *UseCase) GetSessionUser(email string) (model.User, error) {
	userModel, exists := u.repo.GetUserByEmail(email)
	if !exists {
		return model.User{}, errs.WrongCredentials
	}

	return userModel, nil
}
