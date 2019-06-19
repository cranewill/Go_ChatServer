package handler

import (
	"fmt"
	"encoding/json"

	manager "chat_server/manager"
	message "chat_server/message"
)

type ShowRoomsHandler struct {

}

func (handler ShowRoomsHandler)Deal(msgMap []byte) {
	msg := message.ReqShowRoomsMessage{}
	err := json.Unmarshal(msgMap, &msg)
	if err != nil {
		fmt.Println("Decode message error: ", err)
		return
	}
	manager.Pool.ChatManager.ShowChatRooms(msg.PlayerId)
}


