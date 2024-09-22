package db

import (
	"errors"
)

var users = make(map[string]User)

var ErrUserAlreadyExists = errors.New("user already exists")

func CreateUser(email string, user User) error {
	if _, ok := users[email]; ok {
		return ErrUserAlreadyExists
	}

	users[email] = user
	return nil
}
