package db

import (
	"errors"
	"fmt"
	"testing"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()
	usersDB := InitUsersWithData()

	tests := []struct {
		email string
		User
		expectedError error
	}{
		{
			"test@test.com",
			User{
				Username: "user1",
				Password: "pass1",
			},
			nil,
		},
		{
			"test@test.com",
			User{
				Username: "user1",
				Password: "pass1",
			},
			errs.UserAlreadyExists,
		},
		{
			"test1@test.com",
			User{
				Username: "user1",
				Password: "pass1",
			},
			nil,
		},
		{
			"test2@test.com",
			User{
				Username: "user1",
				Password: "pass1",
			},
			nil,
		},
		{
			"test2@test.com",
			User{
				Username: "user1",
				Password: "pass1",
			},
			errs.UserAlreadyExists,
		},
	}

	for i, testUser := range tests {
		t.Run(fmt.Sprintf("Test %v: ", i), func(t *testing.T) {
			err := usersDB.CreateUser(testUser.email, testUser.User)

			if !errors.Is(err, testUser.expectedError) {
				t.Errorf("Test %v failed, ecpexted error: %v, got: %v", i, testUser.expectedError, err.Error())
			}
		})
	}
}
