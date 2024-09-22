package db

import (
	"errors"
)

var userDB = UserDB{
	users: make(map[string]User),
}

var ErrUserAlreadyExists = errors.New("user already exists")

func CreateUser(email string, user User) error {
	userDB.mu.Lock()
	defer userDB.mu.Unlock()

	if _, ok := userDB.users[email]; ok {
		return ErrUserAlreadyExists
	}

	userDB.users[email] = user
	return nil
}
