package handlers

import "2024_2_kotyari/db"

type loginApiRequest struct {
	Email string `json:"email"`
	db.User
}
