package db

import (
	"errors"
	"fmt"
	"testing"

	"2024_2_kotyari/errs"
)

func TestCreateUser(t *testing.T) {
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
			err := CreateUser(testUser.email, testUser.User)

			if !errors.Is(err, testUser.expectedError) {
				t.Errorf("Test %v failed, ecpexted error: %v, got: %v", i, testUser.expectedError, err.Error())
			}
		})
	}
}
