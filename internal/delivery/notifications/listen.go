package notifications

import (
	"fmt"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (n *NotificationsHandler) Listen(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HELLO")

	conn, err := upgrader.Upgrade(w, r, nil)
	defer conn.Close()

	if err != nil {
		utils.WriteErrorJSONByError(w, err, n.errResolver)

		return
	}

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			utils.WriteErrorJSONByError(w, err, n.errResolver)

			return
		}

		conn.WriteMessage(messageType, message)
	}
}
