package model

import "github.com/google/uuid"

type OrderState struct {
	ID    uuid.UUID
	State string
}
