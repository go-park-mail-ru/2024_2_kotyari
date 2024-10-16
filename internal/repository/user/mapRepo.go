package user

import (
	"sync"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	userU "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/user"
)

func NewUserMapRepository() userU.RepoInterface {
	return &MapRepo{
		users: map[string]model.User{},
	}
}

type MapRepo struct {
	mu    sync.RWMutex
	users map[string]model.User
}

func InitUsers() *MapRepo {
	return &MapRepo{
		users: make(map[string]model.User),
	}
}

// Fix
var usersData = map[string]model.User{
	"user@example.com": {
		Username: "Goshanchik",
		Password: "gbHWrVy4JEmoO06xZa4Z3h/LnkSFl0wzkJNtDXXLmq9pU8LRhOhRQRnZ79AdABaK",
	},
	"user1@example.com": {
		Username: "Igorechik",
		Password: "Password124@",
	},
}

func InitUsersWithData() *MapRepo {
	return &MapRepo{
		users: usersData,
	}
}

func (db *MapRepo) GetUserByEmail(email string) (*model.User, bool) {
	db.mu.Lock()
	defer db.mu.Unlock()

	user, exists := db.users[email]

	return &user, exists
}

func (db *MapRepo) InsertUser(user *model.User) (*model.User, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.users[user.Email]; ok {
		return nil, errs.UserAlreadyExists
	}

	db.users[user.Email] = *user

	return user, nil

}
