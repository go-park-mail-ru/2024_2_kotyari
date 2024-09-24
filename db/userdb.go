package db

import "sync"

// User представляет пользователя в системе
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDB struct {
	mu    sync.Mutex
	users map[string]User
}

var userDB = UserDB{
	users: map[string]User{
		"user@example.com":  {Email: "user@example.com", Password: "Password123"},
		"user1@example.com": {Email: "user1@example.com", Password: "Password124"},
	},
}

// GetUserByEmail возвращает пользователя по email
func GetUserByEmail(email string) (User, bool) {
	user, exists := userDB.users[email]

	return user, exists
}
