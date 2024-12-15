package notifications

import (
	notifications "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/notifications/gen"
	"google.golang.org/grpc"
)

func (n *NotificationsGRPC) Register(server *grpc.Server) {
	notifications.RegisterNotificationsServer(server, n)
}
