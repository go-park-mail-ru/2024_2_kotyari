package db

import "sync"

// User представляет пользователя в системе
type User struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
}

type UserDB struct {
	mu    sync.Mutex
	users map[string]User
}

var userDB = UserDB{
	users: map[string]User{"user@example.com": {Username: "Goshanchik", PasswordHash: "Password123@"},
		"user1@example.com": {Username: "Igorechik", PasswordHash: "Password124@"}},
}

// GetUserByEmail возвращает пользователя по email
func GetUserByEmail(email string) (User, bool) {
	user, exists := userDB.users[email]

	return user, exists
}
