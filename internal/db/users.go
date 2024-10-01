package db

import (
	"sync"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
)

type Users struct {
	mu    sync.RWMutex
	users map[string]User
}

func InitUsers() *Users {
	return &Users{
		users: make(map[string]User),
	}
}

func InitUsersWithData() *Users {
	return &Users{
		users: usersData,
	}
}

// GetUserByEmail возвращает пользователя по email
func (u *Users) GetUserByEmail(email string) (User, bool) {
	u.mu.Lock()
	defer u.mu.Unlock()

	user, exists := u.users[email]

	return user, exists
}

func (u *Users) CreateUser(email string, user User) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	if _, ok := u.users[email]; ok {
		return errs.UserAlreadyExists
	}

	u.users[email] = user

	return nil
}
