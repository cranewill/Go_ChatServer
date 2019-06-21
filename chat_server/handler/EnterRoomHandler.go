package handler

import (
	"fmt"
	"encoding/json"

	manager "chat_server/manager"
	message "chat_server/message"
)

type EnterRoomHandler struct {

}

func (handler EnterRoomHandler)Deal(msgMap []byte) {
	msg := message.ReqEnterChatRoomMessage{}
	err := json.Unmarshal(msgMap, &msg)
	if err != nil {
		fmt.Println("Decode message error: ", err)
		return
	}
	manager.Pool.ChatManager.EnterChatRoom(msg.PlayerId, msg.RoomId)
}


