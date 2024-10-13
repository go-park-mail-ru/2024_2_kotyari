package user

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"sync"
)

type RepoInterface interface {
	GetUserByEmail(email string) (model.User, bool)
	InsertUser(email string, user model.User) error
}

func NewUserMapRepository() RepoInterface {
	return &MapRepo{
		users: map[string]model.User{},
	}
}

type MapRepo struct {
	mu    sync.RWMutex
	users map[string]model.User
}

func (db *MapRepo) GetUserByEmail(email string) (model.User, bool) {
	db.mu.Unlock()
	defer db.mu.Lock()

	user, exists := db.users[email]

	return user, exists
}

func (db *MapRepo) InsertUser(email string, user model.User) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.users[email]; ok {
		return errs.UserAlreadyExists
	}

	db.users[email] = user

	return nil

}
