package handler

import (
	"fmt"
	"encoding/json"

	manager "Go_ChatServer/chat_server/manager"
	message "Go_ChatServer/chat_server/message"
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


