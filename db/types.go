package db

import "sync"

type UserDB struct {
	mu    sync.Mutex
	users map[string]User
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
