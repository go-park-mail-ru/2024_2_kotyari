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
	users: map[string]User{"user@example.com": {Username: "Goshanchik", PasswordHash: "ONpKOiynSJk23B6NPeGEqFd0PcC/jlW7ntuzbMRixiKnRQFDlyoFyfkEtzjhLPez"},
		"user1@example.com": {Username: "Igorechik", PasswordHash: "i0yRAkITO6Kr6fgpflRL3KYtAZhD1tGG+aY6ZwMbkIXQTauBiJ8hZVp5V7SZDcjO"}},
}

// GetUserByEmail возвращает пользователя по email
func GetUserByEmail(email string) (User, bool) {
	user, exists := userDB.users[email]

	return user, exists
}
