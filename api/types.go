package api

import "2024_2_kotyari/db"

type signupApiRequest struct {
	Email string
	db.User
}
