package db

import "errors"

var (
	users         = make(map[UserID]User)
	userID UserID = 0
)

func CreateUser(user User) error {
	for _, u := range users {
		if u.Username == user.Username {
			return errors.New("user already exists")
		}
	}

	users[userID] = user
	userID++
	return nil
}
