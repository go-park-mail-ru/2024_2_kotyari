package db

import (
	"errors"
)

var (
	users = make(map[string]User)
)

func CreateUser(email string, user User) error {
	if _, ok := users[email]; ok {
		return errors.New("user already exists")
	}

	users[email] = user
	return nil
}
