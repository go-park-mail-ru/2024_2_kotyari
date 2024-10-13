package user

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/user"
)

type UsecaseInterface interface {
	GetUserByEmail(email string) (model.User, bool)
	CreateUser(email string, user model.User) error
}

func NewUserUsecase(userRepo user.RepoInterface) UsecaseInterface {
	return &Usecase{repo: userRepo}
}

type Usecase struct {
	repo user.RepoInterface
}

func (u Usecase) GetUserByEmail(email string) (model.User, bool) {
	return u.repo.GetUserByEmail(email)
}

func (u Usecase) CreateUser(email string, user model.User) error {
	return u.repo.InsertUser(email, user)
}
