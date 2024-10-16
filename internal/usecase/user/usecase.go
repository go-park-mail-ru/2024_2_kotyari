package user

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/user"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

type RepoInterface interface {
	GetUserByEmail(email string) (*model.User, bool)
	InsertUser(user *model.User) (*model.User, error)
}

func NewUserUsecase(userRepo RepoInterface) user.UsecaseInterface {
	return &Usecase{repo: userRepo}
}

type Usecase struct {
	repo RepoInterface
}

func (u *Usecase) GetUserByEmail(userLoginRequest *model.UserLoginRequestDTO) (*model.User, bool) {
	return u.repo.GetUserByEmail(userLoginRequest.Email)
}

func (u *Usecase) CreateUser(userSignupRequest *model.UserSignupRequestDTO) (*model.User, error) {
	salt, err := utils.GenerateSalt()
	if err != nil {
		return nil, err
	}

	userSignupRequest.Password = utils.HashPassword(userSignupRequest.Password, salt)
	userModel := userSignupRequest.ToUserModel()

	return u.repo.InsertUser(userModel)
}
