package db

import "2024_2_kotyari/errs"

var userDB = UserDB{
	users: make(map[string]User),
}

func CreateUser(email string, user User) error {
	userDB.mu.Lock()
	defer userDB.mu.Unlock()

	if _, ok := userDB.users[email]; ok {
		return errs.UserAlreadyExists
	}

	userDB.users[email] = user
	return nil
}
