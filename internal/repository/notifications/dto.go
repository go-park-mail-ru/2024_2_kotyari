package notifications

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/google/uuid"
)

type OrderStateDTO struct {
	ID    uuid.UUID `db:"id"`
	State string    `db:"new_status"`
}

func (o OrderStateDTO) ToModel() model.OrderState {
	return model.OrderState{
		ID:    o.ID,
		State: o.State,
	}
}
