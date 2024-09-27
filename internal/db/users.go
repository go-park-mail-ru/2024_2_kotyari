package db

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"sync"
)

type Users struct {
	mu    sync.RWMutex
	users map[string]User
}

func InitUsersWithData() *Users {
	return &Users{
		users: usersData,
	}
}

// GetUserByEmail возвращает пользователя по email
func (u *Users) GetUserByEmail(email string) (User, bool) {
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
