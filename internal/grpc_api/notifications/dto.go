package notifications

import (
	notifications "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/notifications/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func orderUpdateToGRPC(state model.OrderState) *notifications.OrderUpdateMessage {
	return &notifications.OrderUpdateMessage{
		OrderId:   state.ID.String(),
		NewStatus: state.State,
	}
}
